package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	openapi "github.com/deepfence/golang_deepfence_sdk/client"
	rhttp "github.com/hashicorp/go-retryablehttp"
)

const (
	maxIdleConnsPerHost = 1024
	auth_field          = "Authorization"
	bearer_format       = "Bearer %s"
)

var (
	AuthError = errors.New("Authentication error")
)

type OpenapiHttpClient struct {
	client *openapi.APIClient
	// Refresher works under the assumption that `tokens.mu` is acquired
	refresher *openapi.APIClient
	api_token *string
	tokens    *ThreadSafeTokens
	done      chan struct{}
}

func (ohc *OpenapiHttpClient) Close() {
	ohc.done <- struct{}{}
}

func (client *OpenapiHttpClient) Client() *openapi.APIClient {
	return client.client
}

type ThreadSafeTokens struct {
	mu                   sync.RWMutex
	access_header_value  string
	refresh_header_value string
	access_token         string
	refresh_token        string
	last_update          time.Time
}

func (tst *ThreadSafeTokens) HasExpired() bool {
	// TODO: extract expiry from JWT token instead of assuming
	return tst.last_update.Before(time.Now().Local().Add(-5 * time.Minute))
}

func NewThreadSafeTokens() *ThreadSafeTokens {
	return &ThreadSafeTokens{
		mu:                   sync.RWMutex{},
		access_header_value:  "",
		refresh_header_value: "",
		access_token:         "",
		refresh_token:        "",
	}
}

type AccessInjectorTransport struct {
	tokens            *ThreadSafeTokens
	originalTransport *http.Transport
}

type RefreshInjectorTransport struct {
	tokens            *ThreadSafeTokens
	originalTransport *http.Transport
}

func NewInjectorTransport[T AccessInjectorTransport | RefreshInjectorTransport](shared_header *ThreadSafeTokens, org *http.Transport) *T {
	return &T{
		tokens:            shared_header,
		originalTransport: org,
	}
}

func (c *AccessInjectorTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	c.tokens.mu.RLock()
	r.Header.Set(auth_field, c.tokens.access_header_value)
	c.tokens.mu.RUnlock()
	return c.originalTransport.RoundTrip(r)
}

func (c *RefreshInjectorTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// No need for lock here as refresher only happens under lock already
	r.Header.Set(auth_field, c.tokens.refresh_header_value)
	return c.originalTransport.RoundTrip(r)
}

func buildHttpsRefresherClient(tokens *ThreadSafeTokens) *http.Client {
	// Set up our own certificate pool
	tlsConfig := &tls.Config{RootCAs: x509.NewCertPool(), InsecureSkipVerify: true}
	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 30 * time.Second,
		TLSClientConfig:     tlsConfig,
	}
	//transport := http.DefaultTransport.(*http.Transport).Clone()
	//transport.TLSClientConfig = tlsConfig
	transport.DisableKeepAlives = false
	transport.DisableCompression = false
	client := &http.Client{
		Transport: NewInjectorTransport[RefreshInjectorTransport](tokens, transport),
	}
	return client
}

// Client is not thread safe.
func NewHttpsConsoleClient(url, port string) *OpenapiHttpClient {

	tlsConfig := &tls.Config{RootCAs: x509.NewCertPool(), InsecureSkipVerify: true}
	transport := &http.Transport{
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 30 * time.Second,
		TLSClientConfig:     tlsConfig,
	}
	transport.DisableKeepAlives = false
	transport.DisableCompression = false

	auth_tokens := NewThreadSafeTokens()

	//transport := http.DefaultTransport.(*http.Transport).Clone()
	//transport.TLSClientConfig = tlsConfig
	rhc := rhttp.NewClient()
	//rhc.HTTPClient.Timeout = 10 * time.Second
	rhc.RetryMax = 3
	rhc.RetryWaitMin = 1 * time.Second
	rhc.RetryWaitMax = 10 * time.Second
	//rhc.Logger = log.NewStdLoggerWithLevel(zerolog.DebugLevel)
	rhc.HTTPClient = &http.Client{
		Transport: NewInjectorTransport[AccessInjectorTransport](auth_tokens, transport),
	}

	servers := openapi.ServerConfigurations{
		{
			URL:         fmt.Sprintf("https://%s:%s", url, port),
			Description: "deepfence_server",
		},
	}

	cfg := openapi.NewConfiguration()
	cfg.HTTPClient = rhc.StandardClient()
	cfg.Servers = servers
	client := openapi.NewAPIClient(cfg)

	cfg2 := openapi.NewConfiguration()
	cfg2.HTTPClient = buildHttpsRefresherClient(auth_tokens)
	cfg2.Servers = servers
	refresher := openapi.NewAPIClient(cfg2)

	done := make(chan struct{})
	unique_client := &OpenapiHttpClient{
		client:    client,
		refresher: refresher,
		done:      done,
		tokens:    auth_tokens,
	}

	rhc.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if err != nil || resp == nil {
			return false, err
		}
		if resp.StatusCode == http.StatusUnauthorized {
			auth_tokens.mu.Lock()
			if auth_tokens.HasExpired() {
				err = unique_client.refreshToken()
			} else {
				err = nil
			}
			auth_tokens.mu.Unlock()
			return err == nil, err
		} else if resp.StatusCode == http.StatusServiceUnavailable {
			return false, err
		}
		return rhttp.DefaultRetryPolicy(ctx, resp, err)
	}

	return unique_client
}

func (cl *OpenapiHttpClient) APITokenAuthenticate(api_token string) error {
	req := cl.client.AuthenticationAPI.AuthToken(context.Background()).ModelAPIAuthRequest(openapi.ModelAPIAuthRequest{
		ApiToken: api_token,
	})
	resp, _, err := cl.client.AuthenticationAPI.AuthTokenExecute(req)
	if err != nil {
		return AuthError
	}

	cl.tokens.mu.Lock()
	defer cl.tokens.mu.Unlock()
	cl.api_token = &api_token

	return cl.updateHeaders(*resp)
}

func (cl *OpenapiHttpClient) refreshToken() error {

	req := cl.refresher.AuthenticationAPI.AuthTokenRefresh(context.Background())

	resp, http_resp, err := cl.refresher.AuthenticationAPI.AuthTokenRefreshExecute(req)
	if err != nil {
		if http_resp != nil && http_resp.StatusCode == http.StatusUnauthorized && cl.api_token != nil {
			req := cl.refresher.AuthenticationAPI.AuthToken(context.Background()).ModelAPIAuthRequest(openapi.ModelAPIAuthRequest{
				ApiToken: *cl.api_token,
			})
			resp, http_resp, err = cl.refresher.AuthenticationAPI.AuthTokenExecute(req)
			if err != nil {
				return AuthError
			}

			if http_resp.StatusCode == http.StatusUnauthorized {
				return AuthError
			}
		} else {
			return err
		}
	}

	return cl.updateHeaders(*resp)
}

func (cl *OpenapiHttpClient) updateHeaders(tokens openapi.ModelResponseAccessToken) error {
	accessToken := tokens.AccessToken
	refreshToken := tokens.RefreshToken
	if accessToken == "" || refreshToken == "" {
		return AuthError
	}

	cl.tokens.access_token = accessToken
	cl.tokens.refresh_token = refreshToken

	cl.tokens.access_header_value = fmt.Sprintf(bearer_format, accessToken)
	cl.tokens.refresh_header_value = fmt.Sprintf(bearer_format, refreshToken)

	cl.tokens.last_update = time.Now()

	return nil
}

func (cl *OpenapiHttpClient) DumpTokens() (access string, refresh string) {
	cl.tokens.mu.RLock()
	defer cl.tokens.mu.RUnlock()
	return cl.tokens.access_token, cl.tokens.refresh_token
}

func (cl *OpenapiHttpClient) SetTokens(access string, refresh string) {
	cl.tokens.mu.Lock()
	defer cl.tokens.mu.Unlock()
	cl.tokens.access_token = access
	cl.tokens.refresh_token = refresh

	cl.tokens.access_header_value = fmt.Sprintf(bearer_format, access)
	cl.tokens.refresh_header_value = fmt.Sprintf(bearer_format, refresh)

	cl.tokens.last_update = time.Now()
}
