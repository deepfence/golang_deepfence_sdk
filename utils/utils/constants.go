package utils

const (
	ErrorUserNotFound = "user not found"
)

// kafka topics
const (
	AUDIT_LOGS                   = "audit-logs"
	VULNERABILITY_SCAN           = "vulnerability-scan"
	VULNERABILITY_SCAN_STATUS    = "vulnerability-scan-status"
	SECRET_SCAN                  = "secret-scan"
	SECRET_SCAN_STATUS           = "secret-scan-status"
	MALWARE_SCAN                 = "malware-scan"
	MALWARE_SCAN_STATUS          = "malware-scan-status"
	SBOM_ARTIFACTS               = "sbom-artifact"
	SBOM_CVE_SCAN                = "sbom-cve-scan"
	CLOUD_COMPLIANCE_SCAN        = "cloud-compliance-scan"
	CLOUD_COMPLIANCE_SCAN_STATUS = "cloud-compliance-scan-status"
	COMPLIANCE_SCAN              = "compliance-scan"
	COMPLIANCE_SCAN_STATUS       = "compliance-scan-status"
	CLOUD_TRAIL_ALERTS           = "cloudtrail-alert"
)

// task names
const (
	SetUpGraphDBTask          = "set_up_graph_db"
	CleanUpGraphDBTask        = "clean_up_graph_db"
	CleanUpPostgresqlTask     = "clean_up_postgresql"
	CleanupDiagnosisLogs      = "clean_up_diagnosis_logs"
	RetryFailedScansTask      = "retry_failed_scans"
	RetryFailedUpgradesTask   = "retry_failed_upgrades"
	ScanSBOMTask              = "tasks_scan_sbom"
	GenerateSBOMTask          = "tasks_generate_sbom"
	CheckAgentUpgradeTask     = "tasks_check_agent_upgrade"
	SyncRegistryTask          = "task_sync_registry"
	TriggerConsoleActionsTask = "trigger_console_actions"
	SecretScanTask            = "task_secret_scan"
	MalwareScanTask           = "task_malware_scan"
	ReportGeneratorTask       = "tasks_generate_report"
	ComputeThreatTask         = "compute_threat"
	SendNotificationTask      = "tasks_send_notification"
	CloudComplianceTask       = "cloud_compliance"
)

const (
	SCAN_STATUS_SUCCESS    = "COMPLETE"
	SCAN_STATUS_STARTING   = "STARTING"
	SCAN_STATUS_INPROGRESS = "IN_PROGRESS"
	SCAN_STATUS_FAILED     = "ERROR"
)

type Neo4jScanType string

const (
	NEO4J_SECRET_SCAN           Neo4jScanType = "SecretScan"
	NEO4J_VULNERABILITY_SCAN    Neo4jScanType = "VulnerabilityScan"
	NEO4J_MALWARE_SCAN          Neo4jScanType = "MalwareScan"
	NEO4J_COMPLIANCE_SCAN       Neo4jScanType = "ComplianceScan"
	NEO4J_CLOUD_COMPLIANCE_SCAN Neo4jScanType = "CloudComplianceScan"
)

var (
	ScanTypeDetectedNode = map[Neo4jScanType]string{
		NEO4J_VULNERABILITY_SCAN:    "Vulnerability",
		NEO4J_SECRET_SCAN:           "Secret",
		NEO4J_MALWARE_SCAN:          "Malware",
		NEO4J_COMPLIANCE_SCAN:       "Compliance",
		NEO4J_CLOUD_COMPLIANCE_SCAN: "CloudCompliance",
	}
	DetectedNodeScanType = map[string]Neo4jScanType{
		"Vulnerability":   NEO4J_VULNERABILITY_SCAN,
		"Secret":          NEO4J_SECRET_SCAN,
		"Malware":         NEO4J_MALWARE_SCAN,
		"Compliance":      NEO4J_COMPLIANCE_SCAN,
		"CloudCompliance": NEO4J_CLOUD_COMPLIANCE_SCAN,
	}
)

type CloudProvider int

const (
	AWS CloudProvider = iota
	GCP
	Azure
	DO
	AWSOrg
)

func StringToCloudProvider(s string) CloudProvider {
	switch s {
	case "aws":
		return AWS
	case "aws_org":
		return AWSOrg
	case "gcp":
		return GCP
	case "azure":
		return Azure
	case "do":
		return DO
	}
	return -1
}

func ResourceTypeToNeo4jLabel(t CloudProvider) string {
	switch t {
	case AWS:
		return "AWS"
	case GCP:
		return "GCP"
	case Azure:
		return "Azure"
	case DO:
		return "DO"
	}
	return ""
}

var Topics = []string{
	VULNERABILITY_SCAN, VULNERABILITY_SCAN_STATUS,
	SECRET_SCAN, SECRET_SCAN_STATUS,
	MALWARE_SCAN, MALWARE_SCAN_STATUS,
	SBOM_ARTIFACTS, SBOM_CVE_SCAN,
	CLOUD_COMPLIANCE_SCAN, CLOUD_COMPLIANCE_SCAN_STATUS,
	COMPLIANCE_SCAN, COMPLIANCE_SCAN_STATUS,
	CLOUD_TRAIL_ALERTS,
	AUDIT_LOGS,
}

// list of task names to create topics
var Tasks = []string{
	SetUpGraphDBTask,
	CleanUpGraphDBTask,
	CleanUpPostgresqlTask,
	CleanupDiagnosisLogs,
	RetryFailedScansTask,
	RetryFailedUpgradesTask,
	ScanSBOMTask,
	GenerateSBOMTask,
	CheckAgentUpgradeTask,
	SyncRegistryTask,
	TriggerConsoleActionsTask,
	ReportGeneratorTask,
	SecretScanTask,
	MalwareScanTask,
	ComputeThreatTask,
	SendNotificationTask,
	CloudComplianceTask,
}

type ReportType string

const (
	ReportXLSX ReportType = "xlsx"
	ReportPDF  ReportType = "pdf"
)
