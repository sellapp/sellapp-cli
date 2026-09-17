# affiliate-referral-sessions

[All commands](../commands.md)

## get

Retrieve an affiliate referral session

Retrieve one store-scoped referral attribution session. Requires `affiliate` and `affiliate:sensitive`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-referral-sessions) · Effect: **read**

```sh
sellapp affiliate-referral-sessions get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| referralSession (positional) | Yes | The referralSession path parameter. |

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

Inspect the complete schema: `sellapp commands affiliate-referral-sessions get --json`.

## list

List affiliate referral sessions

List store-scoped referral attribution sessions. Session identifiers are sensitive. Requires `affiliate` and `affiliate:sensitive`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-referral-sessions) · Effect: **read**

```sh
sellapp affiliate-referral-sessions list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --affiliate-id | No | Filter by affiliate ID. |
| --active | No | Return active or expired sessions. |
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

Inspect the complete schema: `sellapp commands affiliate-referral-sessions list --json`.

