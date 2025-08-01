package config

type Config struct {
	// Common fields for both modes
	Version        string
	ReadOnly       bool
	Toolsets       []string
	EnabledModules []string
	LogFilePath    string
	Debug          bool

	Internal bool

	// Only used for external mode
	BaseURL          string
	AccountID        string
	DefaultOrgID     string
	DefaultProjectID string
	APIKey           string

	// Only used for internal mode
	BearerToken             string
	PipelineSvcBaseURL      string
	PipelineSvcSecret       string
	NgManagerBaseURL        string
	NgManagerSecret         string
	ChatbotBaseURL          string
	ChatbotSecret           string
	GenaiBaseURL            string
	GenaiSecret             string
	ArtifactRegistryBaseURL string
	ArtifactRegistrySecret  string
	NextgenCEBaseURL        string
	NextgenCESecret         string
	IDPSvcBaseURL           string
	IDPSvcSecret            string
	McpSvcSecret            string
	ChaosManagerSvcBaseURL  string
	ChaosManagerSvcSecret   string
	TemplateSvcBaseURL      string
	TemplateSvcSecret       string
	IntelligenceSvcBaseURL  string
	IntelligenceSvcSecret   string
	CodeSvcBaseURL          string
	CodeSvcSecret           string
	LogSvcBaseURL           string
	LogSvcSecret            string
	DashboardSvcBaseURL     string
	DashboardSvcSecret      string
	SCSSvcSecret            string // Added for SCS toolset
	SCSSvcBaseURL           string // Added for SCS toolset
	STOSvcSecret            string // Added for STO toolset
	STOSvcBaseURL           string // Added for STO toolset
	AuditSvcBaseURL         string
	AuditSvcSecret          string
	DBOpsSvcBaseURL         string
	DBOpsSvcSecret          string
	RBACSvcBaseURL          string
	RBACSvcSecret           string
}