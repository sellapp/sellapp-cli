# feedback

[All commands](../commands.md)

## get

Retrieve specific feedback

Retrieve a customer's feedback by its ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp feedback get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| feedback (positional) | Yes | The feedback path parameter. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback get --json`.

## list

List all feedback

List customer feedback for your store, 15 entries per page by default.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp feedback list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback list --json`.

## reply

Reply to feedback

Publish a seller reply to a customer's feedback.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp feedback reply 1 --reply 'Please contact support if you need help with your download.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| feedback (positional) | Yes | The feedback path parameter. |
| --reply | Yes | Value for reply. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback reply --json`.

## search

Search feedback

Search feedback using JSON body filters, search terms, includes, and sort instructions.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp feedback search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --search | No | Value for search. |
| --includes | No | Value for includes. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback search --json`.

## v-2-get-feedback

Retrieve specific feedback

Retrieve a customer's feedback by its ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/feedback/retrieve-specific-feedback) · Effect: **read**

```sh
sellapp feedback v-2-get-feedback 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| feedback (positional) | Yes | The feedback path parameter. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    }
  ],
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback v-2-get-feedback --json`.

## v-2-list-feedback

List all feedback

List customer feedback for your store, 15 entries per page by default. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/feedback/list-all-feedback) · Effect: **read**

```sh
sellapp feedback v-2-list-feedback
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    }
  ],
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback v-2-list-feedback --json`.

## v-2-replace-feedback

Reply to feedback

Publish a seller reply to a customer's feedback. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/feedback) · Effect: **consequential**

```sh
sellapp feedback v-2-replace-feedback 1 --reply 'Please contact support if you need help with your download.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| feedback (positional) | Yes | The feedback path parameter. |
| --reply | Yes | Value for reply. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    }
  ],
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback v-2-replace-feedback --json`.

## v-2-reply-to-feedback

Reply to feedback

Publish a seller reply to a customer's feedback. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/feedback/reply-to-feedback) · Effect: **consequential**

```sh
sellapp feedback v-2-reply-to-feedback 1 --reply 'Please contact support if you need help with your download.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| feedback (positional) | Yes | The feedback path parameter. |
| --reply | Yes | Value for reply. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    }
  ],
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback v-2-reply-to-feedback --json`.

## v-2-search-feedback

Search feedback

Search feedback using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/feedback/search-feedback) · Effect: **read**

```sh
sellapp feedback v-2-search-feedback --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --search | No | Value for search. |
| --includes | No | Value for includes. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    }
  ],
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands feedback v-2-search-feedback --json`.

