package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"time"

	openapi "github.com/deepfence/golang_deepfence_sdk/client"
	rhttp "github.com/hashicorp/go-retryablehttp"
)

func IsConsoleAgent(url string) bool {
	return len(os.Getenv("DEEPFENCE_CONSOLE_AGENT")) > 0
}

func GetConsoleApiToken(console, port string) (string, error) {
	// setup http client
	rhc := rhttp.NewClient()
	rhc.HTTPClient.Timeout = 10 * time.Second
	rhc.RetryMax = 3
	rhc.RetryWaitMin = 1 * time.Second
	rhc.RetryWaitMax = 10 * time.Second

	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		RootCAs:            x509.NewCertPool(),
		InsecureSkipVerify: true,
	}
	tr.DisableKeepAlives = false

	rhc.HTTPClient = &http.Client{Transport: tr}

	servers := openapi.ServerConfigurations{
		{
			URL:         fmt.Sprintf("http://%s:%s", console, port),
			Description: "deepfence_server_internal",
		},
	}

	cfg := openapi.NewConfiguration()
	cfg.HTTPClient = rhc.StandardClient()
	cfg.Servers = servers
	client := openapi.NewAPIClient(cfg)

	token, resp, err := client.InternalAPI.GetConsoleApiTokenExecute(openapi.ApiGetConsoleApiTokenRequest{})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(resp.Status)
	}

	return token.GetApiToken(), nil
}

type StreamingConfig struct {
	Addition    bool
	Deletion    bool
	EntityTypes []string //Supported types: Node, Container, ContainerImage
}

func (cfg StreamingConfig) validateConfig() error {
	if !cfg.Addition && !cfg.Deletion {
		return fmt.Errorf("Invalid StreamingConfig, should have atleast one event set")
	}

	if len(cfg.EntityTypes) == 0 {
		return fmt.Errorf("Invalid StreamingConfig, should have atleast one entity")
	}

	validEntityTypes := make(map[string]bool)
	//We only support streaming for entities of type
	//Host, Container and ContainerImage
	validEntityTypes["Node"] = true
	validEntityTypes["Container"] = true
	validEntityTypes["ContainerImage"] = true

	for _, entity := range cfg.EntityTypes {
		if _, ok := validEntityTypes[entity]; !ok {
			return fmt.Errorf("Invalid entity in request:" + entity)
		}
	}

	return nil
}

func StartStreamingDelta(ctx context.Context, cfg StreamingConfig,
	url, port, apiToken string,
	callback func(bool, []openapi.ModelNodeIdentifier),
	onErrorCallback func(error)) error {
	//First validate the config
	err := cfg.validateConfig()
	if err != nil {
		return err
	}

	apiClient := NewHttpsConsoleClient(url, port)
	if apiClient == nil {
		return fmt.Errorf("Failed to create console client")
	}

	err = apiClient.APITokenAuthenticate(apiToken)
	if err != nil {
		return err
	}

	//Client is authenticated and we have valid config
	//Start the streaming loop
	ticker := time.NewTicker(30 * time.Second)
	timestamp := time.Now()

	topologyAPI := apiClient.Client().TopologyAPI
	req := topologyAPI.GetTopologyDelta(ctx)
	deltaReq := openapi.ModelTopologyDeltaReq{}
	deltaReq.SetAddition(cfg.Addition)
	deltaReq.SetDeletion(cfg.Deletion)
	deltaReq.SetEntityTypes(cfg.EntityTypes)
	deltaReq.SetAdditionTimestamp(timestamp.UnixMilli())
	deltaReq.SetDeletionTimestamp(timestamp.UnixMilli())

streamloop:
	for {
		select {
		case <-ctx.Done():
			break streamloop
		case <-ticker.C:
			req = req.ModelTopologyDeltaReq(deltaReq)
			deltaRes, rh, err := topologyAPI.GetTopologyDeltaExecute(req)
			if err != nil {
				onErrorCallback(err)
				continue
			}

			if rh.StatusCode != http.StatusOK {
				//We will continue on getting any error
				//Not sure if we should keep a counter to limit
				//this retries.
				onErrorCallback(fmt.Errorf("Got non-200 status code:%d", rh.StatusCode))
				continue
			}

			if deltaRes.Additons != nil && len(deltaRes.Additons) > 0 {
				callback(true, deltaRes.Additons)
			}

			if deltaRes.Deletions != nil && len(deltaRes.Deletions) > 0 {
				callback(false, deltaRes.Deletions)
			}

			if deltaReq.Addition && (deltaRes.AdditionTimestamp != nil) {
				deltaReq.SetAdditionTimestamp(*deltaRes.AdditionTimestamp)
			}

			if deltaReq.Deletion && (deltaRes.DeletionTimestamp != nil) {
				deltaReq.SetDeletionTimestamp(*deltaRes.DeletionTimestamp)
			}
		}
	}

	return nil
}
