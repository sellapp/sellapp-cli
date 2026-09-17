# affiliate-referrals

[All commands](../commands.md)

## get

Retrieve an affiliate referral

Retrieve a referral and its attributed order and purchase line item. Requires the `affiliate` credential ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/manage-referral) · Effect: **read**

```sh
sellapp affiliate-referrals get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| referral (positional) | Yes | The referral path parameter. |

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

Inspect the complete schema: `sellapp commands affiliate-referrals get --json`.

## list

List affiliate referrals

List store referrals with their attributed order and purchase line item. Requires the `affiliate` credential ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-referrals) · Effect: **read**

```sh
sellapp affiliate-referrals list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --affiliate-id | No | Filter by store affiliate ID. |
| --order-id | No | Filter by attributed Order ID. |
| --status | No | Filter by referral status. |
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

Inspect the complete schema: `sellapp commands affiliate-referrals list --json`.

## update

Update referral status

Move an unassigned referral through a valid review, acceptance, or rejection transition. Referrals already assigned to a payout cannot be changed independently. Requires `affiliate` and `affiliate:write`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/manage-referral) · Effect: **consequential**

```sh
sellapp affiliate-referrals update 71 --status accepted --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| referral (positional) | Yes | The referral path parameter. |
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

Inspect the complete schema: `sellapp commands affiliate-referrals update --json`.

