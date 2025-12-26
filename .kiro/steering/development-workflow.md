---
inclusion: always
---

# Development Workflow & Commands

## Local Development Setup

```bash
# Install dependencies
go mod download
go mod vendor

# Run tests
go test -mod vendor -v ./statuspage/...

# Lint code
golint ./...

# Format code
gofmt -w .
```

## Build Commands

```bash
# Local build
go build -mod=vendor -o terraform-provider-statuspage

# Cross-platform builds (as done in CI)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-darwin-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-s -w" -a -o build/terraform-provider-statuspage-linux-amd64
```

## Docker Development

```bash
# Build and test using Docker
docker build .

# Run tests only
docker build --target test .

# Build binaries only (skip tests)
docker build --target build .
```

**Note:** The Docker build includes testing but not linting, as golint has compatibility issues with Go 1.12 in Alpine. Linting is handled by GitHub Actions using the native environment.

## Testing Provider Locally

1. Build the provider binary
2. Place in Terraform plugins directory or use local override
3. Create test Terraform configuration with statuspage provider
4. Set STATUSPAGE_API_KEY environment variable
5. Run `terraform init`, `terraform plan`, `terraform apply`

## Environment Variables

- `STATUSPAGE_API_KEY` - Required for provider authentication
- `TF_LOG=DEBUG` - Enable Terraform debug logging
- `TF_LOG_PROVIDER=DEBUG` - Enable provider-specific debug logging

## CI/CD Notes

- GitHub Actions workflow runs lint, test, and multi-arch builds
- All binaries get SHA256 checksums generated and uploaded as artifacts
- Uses vendored dependencies for reproducible builds
- Workflow triggers on pushes to main/master branches and pull requests
- Docker build is also available as an alternative build method
