---
inclusion: always
---

# Terraform Provider for Statuspage - Project Overview

This is a Terraform provider for managing Statuspage.io resources. The project follows Terraform provider conventions and uses the HashiCorp Terraform SDK.

## Project Structure

- `main.go` - Entry point that serves the Terraform plugin
- `statuspage/` - Core provider implementation
  - `provider.go` - Provider definition and configuration
  - `config.go` - API client configuration
  - `resource_page.go` - Statuspage page resource implementation
  - `provider_test.go` - Provider tests

## Key Dependencies

- HashiCorp Terraform SDK v0.11.13 (legacy version)
- Custom Statuspage Go client: `github.com/nagelflorian/statuspage-go`
- Logrus for structured logging

## API Limitations

The Statuspage API has known limitations:

- No endpoint to create new pages (only read/update existing ones)
- No endpoint to delete pages
- Provider mainly manages existing page configurations

## Build Process

- Uses Docker for consistent builds across platforms
- Builds for multiple architectures: darwin-amd64, linux-386, linux-amd64, linux-arm
- Includes linting with golint and testing before build
- Generates SHA256 checksums for all binaries
