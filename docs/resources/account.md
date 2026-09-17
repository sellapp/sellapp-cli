# account

[All commands](../commands.md)

## get-account

Read your profile

Read your account identity and credential type with account:read on an API key or the official CLI OAuth grant. No store or X-STORE is needed, including before creating your first store. Browser and customer sessions are rejected.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp account get-account
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
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands account get-account --json`.

## get-store

Read an accessible store

Read a currently accessible store by numeric ID with account:read on an API key or official CLI OAuth. No X-STORE is needed. Keys with selected stores must include this ID. Inaccessible stores return 404.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp account get-store 12
```

| Input | Required | Meaning |
| --- | --- | --- |
| store (positional) | Yes | Value for store. |

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
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands account get-store --json`.

## get-store-permissions

Inspect effective permissions

Inspect authorization for the explicit X-STORE slug. API keys require account:read and return key_abilities, store_permissions, effective_abilities and store_access. OAuth returns its granted scopes and reviewed permission categories. The credential_type identifies the typed response. These checks do not establish record-state, provider or financial eligibility.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp account get-store-permissions
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
    },
    {
      "schemeName": "storeAuth",
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

Inspect the complete schema: `sellapp commands account get-store-permissions --json`.

