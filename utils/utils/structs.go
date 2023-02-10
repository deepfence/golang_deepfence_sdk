package utils

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
