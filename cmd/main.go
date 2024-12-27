package main

import (
	"flag"
	"fmt"

	"github.com/marcom4rtinez/terraform-registry-manifest/pkg/hash"
	"github.com/marcom4rtinez/terraform-registry-manifest/pkg/registry"
)

func main() {
	version := flag.String("version", "", "Version of the Provider")
	protocols := flag.String("protocols", "", "Comma-separated list of protocols")
	os := flag.String("os", "", "Operating system")
	arch := flag.String("arch", "", "Architecture")
	providerName := flag.String("provider_name", "", "Name of the provider")
	baseURL := flag.String("base_url", "", "Download URL for the artifact")
	keyID := flag.String("key_id", "", "GPG key ID")
	asciiArmor := flag.String("ascii_armor", "", "GPG public key in ASCII armor format")
	source := flag.String("source", "", "Source of the GPG key")
	sourceURL := flag.String("source_url", "", "Source URL of the GPG key")
	manifestFile := flag.String("manifest", "", "Manifest file path")
	replaceSHAhashes := flag.Bool("hashes", false, "Set this to true if you want to replace the hashes in a already written file")

	flag.Parse()

	if *replaceSHAhashes {
		fmt.Println(*manifestFile)
		hash.ReplaceHashes(*manifestFile)
		return
	}

	fmt.Println(registry.GenerateManifest(*version, *protocols, *os, *arch, *providerName, *baseURL, *keyID, *asciiArmor, *source, *sourceURL))
}
