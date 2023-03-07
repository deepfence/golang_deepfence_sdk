package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func newMinioClient() (*minio.Client, error) {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("deepfence", "deepfence", ""),
		Secure: false,
	})
	return minioClient, err
}

const (
	ConsoleDiagnosisFileServerPrefix = "/diagnosis/console-diagnosis/"
	AgentDiagnosisFileServerPrefix   = "/diagnosis/agent-diagnosis/"
	DiagnosisLinkExpiry              = 5 * time.Minute
	namespace                        = "default"
	consoleIp                        = "localhost"
)

type DiagnosticLogsResponse struct {
	UrlLink   string `json:"url_link"`
	Label     string `json:"label"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type GetDiagnosticLogsResponse struct {
	ConsoleLogs []DiagnosticLogsResponse `json:"console_logs"`
	AgentLogs   []DiagnosticLogsResponse `json:"agent_logs"`
}

func GetDiagnosticLogs(ctx context.Context) (*GetDiagnosticLogsResponse, error) {
	mc, err := newMinioClient()
	if err != nil {
		return nil, err
	}
	diagnosticLogs := GetDiagnosticLogsResponse{
		ConsoleLogs: getDiagnosticLogsHelper(ctx, mc, ConsoleDiagnosisFileServerPrefix),
		AgentLogs:   getDiagnosticLogsHelper(ctx, mc, AgentDiagnosisFileServerPrefix),
	}
	return &diagnosticLogs, err
}

func ExposeFile(mc *minio.Client, filePath string) (string, error) {
	headers := http.Header{}
	//headers.Add("Host", consoleIp)
	urlLink, err := mc.PresignHeader(context.Background(),
		"GET",
		namespace,
		filePath,
		DiagnosisLinkExpiry,
		url.Values{},
		headers)
	if err != nil {
		return "", err
	}
	fmt.Println(filePath, urlLink.String())
	//exposedUrl := strings.ReplaceAll(urlLink.String(), "deepfence-file-server:9000", fmt.Sprintf("%s/file-server", consoleIp))
	//exposedUrl = strings.ReplaceAll(exposedUrl, "https://", "http://")
	return urlLink.String(), nil
}

func getDiagnosticLogsHelper(ctx context.Context, mc *minio.Client, pathPrefix string) []DiagnosticLogsResponse {
	fmt.Println(pathPrefix)
	objects := mc.ListObjects(ctx, namespace, minio.ListObjectsOptions{
		WithVersions: false,
		WithMetadata: false,
		Prefix:       pathPrefix,
		Recursive:    false,
		MaxKeys:      0,
		StartAfter:   "",
		UseV1:        false,
	})
	diagnosticLogsResponse := make([]DiagnosticLogsResponse, 0)
	for obj := range objects {
		message := ""
		urlLink, err := ExposeFile(mc, obj.Key)
		if err != nil {
			message = err.Error()
		}
		diagnosticLogsResponse = append(diagnosticLogsResponse, DiagnosticLogsResponse{
			UrlLink:   urlLink,
			Label:     filepath.Base(obj.Key),
			Message:   message,
			CreatedAt: obj.LastModified.Format(time.DateTime),
		})
	}
	return diagnosticLogsResponse
}

func main() {
	d, err := GetDiagnosticLogs(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	a, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(a))
}

func main2() {
	minioClient, err := newMinioClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	objects := minioClient.ListObjects(ctx, namespace, minio.ListObjectsOptions{
		WithVersions: false,
		WithMetadata: false,
		Prefix:       "ramanan/df/",
		Recursive:    false,
		MaxKeys:      0,
		StartAfter:   "",
		UseV1:        false,
	})
	for obj := range objects {
		a, _ := json.MarshalIndent(obj, "", "\t")
		fmt.Println(string(a))
		fmt.Println("")
		//fmt.Printf("%+v \n", obj)
	}
}
