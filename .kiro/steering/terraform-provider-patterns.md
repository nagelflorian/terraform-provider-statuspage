---
inclusion: fileMatch
fileMatchPattern: "statuspage/*.go"
---

# Terraform Provider Development Patterns

## Resource Implementation Checklist

When working on Terraform resources, ensure:

- [ ] Schema defines all available fields with correct types
- [ ] Required vs Optional fields are properly marked
- [ ] Default values use `DefaultFunc` for environment variables
- [ ] All CRUD operations are implemented (even if some are no-ops)
- [ ] Resource ID is set using `d.SetId()` in Create and Read
- [ ] All schema fields are populated in Read operation
- [ ] Update operation calls Read after successful API update
- [ ] Error messages include resource context and IDs

## Schema Field Patterns

```go
// Environment variable with fallback
"api_key": {
    Type:        schema.TypeString,
    Required:    true,
    DefaultFunc: schema.EnvDefaultFunc("STATUSPAGE_API_KEY", nil),
    Description: "The Statuspage API Key",
},

// Optional boolean with proper pointer handling
"hidden_from_search": {
    Type:     schema.TypeBool,
    Optional: true,
},
```

## API Client Integration

- Always pass context to API calls
- Handle API errors with resource context
- Use structured logging with resource IDs
- Validate required parameters before API calls

## State Management

- Use `d.Set()` for all state updates in Read
- Handle API response field mapping carefully
- Consider field name transformations (snake_case vs camelCase)
- Set resource ID to empty string in Delete operations

## Legacy SDK Considerations

This project uses Terraform SDK v0.11.x patterns:

- Uses `terraform.ResourceProvider` interface
- Schema validation is manual
- No built-in retry mechanisms
- Context handling is basic
