package utils

import "encoding/json"

type ScanSbomRequest struct {
	SbomParameters
	SbomBody
}

type SbomParameters struct {
	ImageName             string `json:"image_name"`
	ImageId               string `json:"image_id"`
	ScanId                string `json:"scan_id" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	HostName              string `json:"host_name"`
	NodeId                string `json:"node_id"`
	NodeType              string `json:"node_type"`
	ScanType              string `json:"scan_type"`
	ContainerName         string `json:"container_name"`
	SBOMFilePath          string `json:"sbom_file_path"`
	Mode                  string `json:"mode,omitempty"`
	RegistryId            string `json:"registry_id,omitempty"`
}

type SbomBody struct {
	SBOM string `json:"sbom" required:"true"`
}

type SecretScanParameters struct {
	ImageName             string `json:"image_name"`
	ImageId               string `json:"image_id"`
	ScanId                string `json:"scan_id" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	HostName              string `json:"host_name"`
	NodeId                string `json:"node_id"`
	NodeType              string `json:"node_type"`
	ScanType              string `json:"scan_type"`
	ContainerName         string `json:"container_name"`
	Mode                  string `json:"mode,omitempty"`
	RegistryId            string `json:"registry_id,omitempty"`
}

type MalwareScanParameters struct {
	ImageName             string `json:"image_name"`
	ImageId               string `json:"image_id"`
	ScanId                string `json:"scan_id" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	HostName              string `json:"host_name"`
	NodeId                string `json:"node_id"`
	NodeType              string `json:"node_type"`
	ScanType              string `json:"scan_type"`
	ContainerName         string `json:"container_name"`
	Mode                  string `json:"mode,omitempty"`
	RegistryId            string `json:"registry_id,omitempty"`
}

type ReportParams struct {
	ReportID   string        `json:"report_id"`
	ReportType string        `json:"report_type"`
	Duration   int           `json:"duration"`
	Filters    ReportFilters `json:"filters"`
}

type ReportFilters struct {
	ScanType            string   `json:"scan_type" validate:"required" required:"true" enum:"vulnerability,secret,malware,compliance,cloud_compliance"`
	NodeType            string   `json:"node_type" validate:"required" required:"true" enum:"host,container,container_image,linux,aws,gcp,azure"`
	SeverityOrCheckType []string `json:"severity_or_check_type" enum:"critical,high,medium,low,cis,gdpr,nist,hipaa,pci,soc2"`
}

func (r ReportFilters) String() string {
	if b, err := json.Marshal(r); err != nil {
		return ""
	} else {
		return string(b)
	}
}
