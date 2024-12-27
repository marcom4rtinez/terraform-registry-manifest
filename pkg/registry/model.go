package registry

type TemplateData struct {
	Version        string
	Protocols      []string
	Platforms      []Platform
	ProviderName   string
	BaseURL        string
	KeyID          string
	AsciiArmor     string
	TrustSignature string
	Source         string
	SourceURL      string
}

type Platform struct {
	OS           string
	Arch         string
	PlatformName string
}
