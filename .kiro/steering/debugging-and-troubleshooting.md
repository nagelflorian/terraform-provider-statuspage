---
inclusion: manual
---

# Debugging & Troubleshooting Guide

## Common Issues & Solutions

### Provider Not Found

```bash
# Ensure provider is built and in correct location
go build -mod=vendor -o terraform-provider-statuspage
# Place in ~/.terraform.d/plugins/ or use local override
```

### API Authentication Errors

```bash
# Verify API key is set
echo $STATUSPAGE_API_KEY
# Enable debug logging
export TF_LOG=DEBUG
export TF_LOG_PROVIDER=DEBUG
```

### Resource State Issues

- Check if page_id exists in Statuspage
- Verify API permissions for the key
- Use `terraform refresh` to sync state
- Check Terraform state file for corruption

### Build Failures

```bash
# Clean and rebuild vendor
rm -rf vendor/
go mod vendor
go mod tidy

# Check for missing dependencies
go mod download
```

## Debug Logging Patterns

The provider uses logrus for structured logging:

```go
log.WithFields(log.Fields{
    "id": d.Id(),
    "operation": "read",
}).Debug("Processing resource")
```

## Testing API Connectivity

```bash
# Test API key manually
curl -H "Authorization: OAuth YOUR_API_KEY" \
     https://api.statuspage.io/v1/pages/YOUR_PAGE_ID
```

## Terraform Debug Commands

```bash
# Plan with debug output
TF_LOG=DEBUG terraform plan

# Show current state
terraform show

# Import existing resource
terraform import statuspage_page.example YOUR_PAGE_ID
```

## Known Limitations to Remember

- Cannot create new pages via API
- Cannot delete pages via API
- Provider only manages existing page configurations
- Some fields may be read-only depending on Statuspage plan
