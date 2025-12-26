---
inclusion: always
---

# Go Coding Standards for Terraform Provider

## Code Style & Formatting

- Use `gofmt` for consistent formatting
- Follow standard Go naming conventions (PascalCase for exported, camelCase for unexported)
- Use meaningful variable names, especially for Terraform schema fields
- Keep functions focused and single-purpose

## Terraform Provider Patterns

- Resource functions follow CRUD pattern: Create, Read, Update, Delete
- Always implement `resourceRead` to refresh state from API
- Use `d.SetId()` to set resource identifier
- Handle API errors gracefully with descriptive messages
- Use `schema.ResourceData.Set()` for all state updates

## Error Handling

- Return descriptive errors with context: `fmt.Errorf("Failed to update page %q: %s", pageID, err)`
- Log operations with structured logging using logrus
- Include resource ID in log fields for debugging

## Testing

- Write unit tests for all resource operations
- Use Terraform's testing framework for acceptance tests
- Mock external API calls in unit tests
- Test error conditions and edge cases

## API Client Usage

- Always use context for API calls: `context.TODO()` or proper context
- Handle rate limiting and retries appropriately
- Validate required fields before API calls
- Use typed parameters for API operations
