package registry

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func GenerateManifest(version string, protocols string, os string, arch string, providerName string, baseURL string, keyID string, asciiArmor string, source string, sourceURL string) string {
	// Create Platform combinations
	platforms := []Platform{}
	for _, current_os := range strings.Split(os, ",") {
		for _, current_arch := range strings.Split(arch, ",") {
			if current_os == "darwin" && (current_arch == "386" || current_arch == "arm") {
				continue
			}
			current_platform := Platform{
				OS:           current_os,
				Arch:         current_arch,
				PlatformName: fmt.Sprintf("%s_%s", current_os, current_arch),
			}
			platforms = append(platforms, current_platform)
		}
	}

	// Create the data structure to pass to the template
	templateData := TemplateData{
		Version:        version,
		Protocols:      strings.Split(protocols, ","),
		Platforms:      platforms,
		ProviderName:   providerName,
		BaseURL:        baseURL,
		KeyID:          keyID,
		AsciiArmor:     strings.ReplaceAll(strings.ReplaceAll(asciiArmor, "\n", "\\n"), "-----BEGIN PGP PUBLIC KEY BLOCK-----", "-----BEGIN PGP PUBLIC KEY BLOCK-----\\nVersion: GnuPG v1"),
		TrustSignature: "",
		Source:         source,
		SourceURL:      sourceURL,
	}

	tmpl, err := template.New("registryTemplate").Parse(registryTemplateContent)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
	return buf.String()
}
