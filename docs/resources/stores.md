# stores

[All commands](../commands.md)

## create

Create a store

Create a store owned by the authenticated user, with the same defaults as dashboard creation. Requires stores:create on an unrestricted user-owned API key, or the official CLI OAuth admin grant. Customer and browser sessions and direct MCP bearer tokens are rejected; MCP uses reviewed server delegation. Selected-store keys cannot create stores. No existing store or X-STORE header is needed; Name and slug are required; slug is normalized to lowercase. Visibility defaults to PUBLIC. This does not change your dashboard's selected store. Idempotency-Key is required and retained for 24 hours: reuse the same key and identical body after a lost response. A conflicting request returns 409; invalid or duplicate slugs return 422. Retries are bound to the user and operation across token rotation and supported credential types; current authorization is checked before replay. Store defaults and the successful retry receipt commit together. You can create one new store per user every 60 seconds, across all credentials. A new creation during that interval returns 429 with Retry-After and rate-limit headers. Wait for that delay and retry the same key and body. Replaying a successful original key recovers its store without creating another or consuming another creation allowance.

[API reference](https://sell.app/docs/api/oauth) · Effect: **write**

```sh
sellapp stores create --name 'Launch Lab' --slug launchlab
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Value for name. |
| --slug | Yes | Lowercased before validation. Reserved infrastructure names such as api, www and cdn are unavailable. |
| --visibility | No | Value for visibility. |

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
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands stores create --json`.

## list

List accessible stores

List all stores you can currently access, including stores created or joined after login. API keys require account:read; selected-store keys see only the intersection of current membership and their fixed store list. An empty list is a successful result. Use a returned slug in X-STORE for store operations.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp stores list
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

Inspect the complete schema: `sellapp commands stores list --json`.

