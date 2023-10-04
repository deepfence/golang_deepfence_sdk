package main

import (
	"context"
	"encoding/json"
	"strings"

	openapi "github.com/deepfence/golang_deepfence_sdk/client"
	oahttp "github.com/deepfence/golang_deepfence_sdk/utils/http"
	"github.com/rs/zerolog/log"
)

var sampleApiClient *oahttp.OpenapiHttpClient
var ctxCancelFunc context.CancelFunc
var scanConfig = []openapi.ModelVulnerabilityScanConfigLanguage{
	{"base"},
	{"java"},
	{"javascript"},
	{"golang"},
	{"ruby"},
	{"python"},
	{"php"},
	{"dotnet"},
	{"golang-binary"},
	{"rust-binary"},
}

func SampleStreamer(console_host, console_port,
	console_api_token string) {

	//This sampleApiClient is required to make the api calls like
	//start secret scan, start malware scan, start vulnerability scan etc
	sampleApiClient = oahttp.NewHttpsConsoleClient(console_host, console_port)
	err := sampleApiClient.APITokenAuthenticate(console_api_token)
	if err != nil {
		log.Info().Msgf("API Client Autentication failed with error: %v", err)
		return
	}

	streamConfig := oahttp.StreamingConfig{}
	streamConfig.Addition = true
	streamConfig.Deletion = true
	streamConfig.EntityTypes = []string{"Node",
		"Container",
		"ContainerImage"}

	ctx, cancelFunc := context.WithCancel(context.Background())
	ctxCancelFunc = cancelFunc

	err = oahttp.StartStreamingDelta(ctx, streamConfig, console_host,
		console_port, console_api_token, StreamConsumerFunc, OnErrorCallback)

	if err != nil {
		log.Info().Msgf("Error in StartStreamingDelta: %v", err)
	}
}

// This is a callback function called by streamer when we receive the data.
// Ideally the logic inside this function should be non-blocking to
// avoid blocking the streamer.
func StreamConsumerFunc(isAdd bool, deltas []openapi.ModelNodeIdentifier) {
	if isAdd {
		go handleAddition(deltas)
	} else {
		go handleDeletion(deltas)
	}
}

// This is a callback function called by streamer when we encounter any error
// during streaming.
func OnErrorCallback(err error) {
	log.Error().Msgf("Error occured during the streaming:%v", err)
	//Implementations can either choose to cancel the context and stop streaming
	//or we can just log the error and continue.
	ctxCancelFunc()
}

// Sample handler for addition deltas. This starts scans on
// the received deltas
func handleAddition(deltas []openapi.ModelNodeIdentifier) {
	log.Info().Msgf("Addition Delta received: %v", deltas)

	stdDeltas := make([]openapi.ModelNodeIdentifier, 0)
	complianceDeltas := make([]openapi.ModelNodeIdentifier, 0)

	for _, entity := range deltas {
		switch strings.ToLower(entity.NodeType) {
		case "container", "containerimage":
			stdDeltas = append(stdDeltas, entity)
		case "node":
			stdDeltas = append(stdDeltas, entity)
			complianceDeltas = append(complianceDeltas, entity)
		}
	}
	if len(stdDeltas) > 0 {
		startSecretScan(stdDeltas)
		startMalwareScan(stdDeltas)
		startVulnerabilityScan(stdDeltas)
	}

	if len(complianceDeltas) > 0 {
		startComplianceScan(complianceDeltas)
	}
}

func handleDeletion(deltas []openapi.ModelNodeIdentifier) {
	log.Info().Msgf("Deletion Deltas received: %v", deltas)
	//Add the code to handle the deletion deltas in this function
}

func startSecretScan(deltas []openapi.ModelNodeIdentifier) {
	secretScanApi := sampleApiClient.Client().SecretScanAPI

	secretScanReq := openapi.ModelSecretScanTriggerReq{}
	secretScanReq.Filters = *(openapi.NewModelScanFilterWithDefaults())
	secretScanReq.NodeIds = deltas

	req := secretScanApi.StartSecretScan(context.Background())
	req = req.ModelSecretScanTriggerReq(secretScanReq)
	res, _, err := secretScanApi.StartSecretScanExecute(req)
	if err != nil {
		log.Info().Msgf("Error in starting secret scan: %v", err)
		return
	}

	respStr, err := json.Marshal(res)
	if err != nil {
		log.Info().Msgf("Failed to json marshal response, error:%v", err)
		return
	}
	log.Info().Msgf("Successfully started scan for deltas: %v", deltas)
	log.Info().Msgf("StartSecretScan Response: %s", respStr)
}

func startMalwareScan(deltas []openapi.ModelNodeIdentifier) {
	malwareScanApi := sampleApiClient.Client().MalwareScanAPI
	malwareScanReq := openapi.ModelMalwareScanTriggerReq{}

	malwareScanReq.Filters = *(openapi.NewModelScanFilterWithDefaults())
	malwareScanReq.NodeIds = deltas
	req := malwareScanApi.StartMalwareScan(context.Background())
	req = req.ModelMalwareScanTriggerReq(malwareScanReq)
	res, _, err := malwareScanApi.StartMalwareScanExecute(req)
	if err != nil {
		log.Info().Msgf("Error in starting malware scan: %v", err)
		return
	}

	respStr, err := json.Marshal(res)
	if err != nil {
		log.Info().Msgf("Failed to json marshal response, error:%v", err)
		return
	}

	log.Info().Msgf("Successfully started Malware scan for deltas: %v", deltas)
	log.Info().Msgf("StartMalwareScan Response: %s", respStr)

}

func startVulnerabilityScan(deltas []openapi.ModelNodeIdentifier) {
	vulnerabilityScanApi := sampleApiClient.Client().VulnerabilityAPI
	vulnerabilityScanReq := openapi.ModelVulnerabilityScanTriggerReq{}

	vulnerabilityScanReq.Filters = *(openapi.NewModelScanFilterWithDefaults())
	vulnerabilityScanReq.NodeIds = deltas

	vulnerabilityScanReq.ScanConfig = scanConfig

	req := vulnerabilityScanApi.StartVulnerabilityScan(context.Background())
	req = req.ModelVulnerabilityScanTriggerReq(vulnerabilityScanReq)
	res, _, err := vulnerabilityScanApi.StartVulnerabilityScanExecute(req)
	if err != nil {
		log.Info().Msgf("Error in starting vulnerability scan: %v", err)
		return
	}

	respStr, err := json.Marshal(res)
	if err != nil {
		log.Info().Msgf("Failed to json marshal response, error:%v", err)
		return
	}

	log.Info().Msgf("Successfully started Vulnerability scan for deltas: %v", deltas)
	log.Info().Msgf("StartVulnerabilityScan Response: %s", respStr)
}

func startComplianceScan(deltas []openapi.ModelNodeIdentifier) {
	complianceScanApi := sampleApiClient.Client().ComplianceAPI
	complianceScanReq := openapi.ModelComplianceScanTriggerReq{}

	complianceScanReq.Filters = *(openapi.NewModelScanFilterWithDefaults())
	complianceScanReq.NodeIds = deltas
	req := complianceScanApi.StartComplianceScan(context.Background())
	req = req.ModelComplianceScanTriggerReq(complianceScanReq)
	res, _, err := complianceScanApi.StartComplianceScanExecute(req)
	if err != nil {
		log.Info().Msgf("Error in starting compliance scan: %v", err)
		return
	}

	respStr, err := json.Marshal(res)
	if err != nil {
		log.Info().Msgf("Failed to json marshal response, error:%v", err)
		return
	}

	log.Info().Msgf("Successfully started Compliance scan for deltas: %v", deltas)
	log.Info().Msgf("StartComplianceScan Response: %s", respStr)
}
