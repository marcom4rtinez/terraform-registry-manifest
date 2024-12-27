Terraform Registry Manifest
---

Go Tool to create a manifest that can be used in a Terraform Registry. Such as https://github.com/marcom4rtinez/terraform-registry

> [!NOTE]
> This Tool has two usage modes.

1. Use Tool to generate the manifest file without hashes

```bash
go run github.com/marcom4rtinez/terraform-registry-manifest/cmd/manifest@latest "${args[@]}" > registry-manifest.json
```

2. Fill hashes into a pre-generated manifest file

```bash
cat hashes | go run github.com/marcom4rtinez/terraform-registry-manifest/cmd/manifest@latest --hashes --manifest registry-manifest.json
```