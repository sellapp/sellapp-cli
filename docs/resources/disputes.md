# disputes

[All commands](../commands.md)

## get-dispute

Retrieve a dispute

Retrieve one redacted dispute. No dispute mutations are exposed. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/disputes) · Effect: **read**

```sh
sellapp disputes get-dispute 01992b10-a7d2-7b91-9822-32e8d7454e7c
```

| Input | Required | Meaning |
| --- | --- | --- |
| dispute (positional) | Yes | The dispute identifier. |

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

Inspect the complete schema: `sellapp commands disputes get-dispute --json`.

## list-disputes

List disputes

Read-only, redacted, store-scoped disputes that may reference an Order or Charge. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/disputes) · Effect: **read**

```sh
sellapp disputes list-disputes
```

| Input | Required | Meaning |
| --- | --- | --- |

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

Inspect the complete schema: `sellapp commands disputes list-disputes --json`.

