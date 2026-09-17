# entitlements

[All commands](../commands.md)

## list-customer-entitlements

List customer entitlements

Unified access without raw keys, serials, files, dynamic payloads, or seller-only metadata. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **read**

```sh
sellapp entitlements list-customer-entitlements 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer identifier. |

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

Inspect the complete schema: `sellapp commands entitlements list-customer-entitlements --json`.

## list-customer-entitlements-by-external-id

List customer entitlements

Unified access without raw keys, serials, files, dynamic payloads, or seller-only metadata. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **read**

```sh
sellapp entitlements list-customer-entitlements-by-external-id externalId_01K4CUSTOMER
```

| Input | Required | Meaning |
| --- | --- | --- |
| externalId (positional) | Yes | The externalId identifier. |

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

Inspect the complete schema: `sellapp commands entitlements list-customer-entitlements-by-external-id --json`.

