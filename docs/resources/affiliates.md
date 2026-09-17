# affiliates

[All commands](../commands.md)

## get

Retrieve an affiliate

Retrieve one store-scoped affiliate, profile, balances, and counts. Requires the `affiliate` ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/retrieve-affiliate) · Effect: **read**

```sh
sellapp affiliates get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate path parameter. |

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

Inspect the complete schema: `sellapp commands affiliates get --json`.

## list

List affiliates

List store affiliate memberships, approved profile details, referral counts, and integer USD-cent balances. Requires the `affiliate` ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-affiliates) · Effect: **read**

```sh
sellapp affiliates list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --status | No | Filter by affiliate program status. |
| --search | No | Search affiliate email addresses and profile identifiers. |
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

Inspect the complete schema: `sellapp commands affiliates list --json`.

## update

Update affiliate status

Approve, disable, re-enable, or reject an affiliate through its valid state transition. Requires `affiliate` and `affiliate:write`; pending invitations cannot be approved until an application is submitted. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/update-affiliate-status) · Effect: **consequential**

```sh
sellapp affiliates update 42 --status active --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate path parameter. |
| --status | Yes | Value for status. |

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

Inspect the complete schema: `sellapp commands affiliates update --json`.

