# Types of API Changes

## Breaking Changes

- Adding a new required field in the request body of an existing endpoint
- Adding a new update-able field to response
- Renaming an existing field in the request body of an existing endpoint
- Renaming an API endpoint (changing the URL structure)
- Removing an API endpoint
- Removing a field from the request or response body of an existing endpoint
- Remove a query parameter from an existing endpoint
- Changing the value type of an existing field in the request body of an existing endpoint. (Eg: from integer to string)
- Modifying the allowed values of an enum field
- Converting an optional field in the request body of an existing endpoint to required
- Adding/Modifying validation to an existing field of an existing endpoint
- Modifying existing response error codes, response error messages
- Modifying existing fields (keys/value types) of the response body
- Changing the response status code (Eg: HTTP 200 to 202)
- A significant change in functionality of an existing endpoint that results in surprising new behavior

## Backwards Compatible Changes

- Adding a new API endpoint
- Adding a new optional field in the request body of an existing endpoint
- Adding a new query parameter for an existing endpoint
- Adding a new read-only field to the response body
- Bug fixes to existing API
