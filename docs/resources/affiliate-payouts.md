# affiliate-payouts

[All commands](../commands.md)

## create

Create an affiliate payout

Create a payout record for all eligible accepted referrals of one active affiliate. The merchant arranges payment separately. Replaying the same Idempotency-Key returns the same record. Requires `affiliate`, `affiliate:payout`, and enough eligible commissions to meet the store's minimum payout amount. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/create-payout) · Effect: **consequential**

```sh
sellapp affiliate-payouts create 1 --yes
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

Inspect the complete schema: `sellapp commands affiliate-payouts create --json`.

## get

Retrieve an affiliate payout

Retrieve one merchant-managed affiliate payout record with a masked destination and referral count. Requires the `affiliate` credential ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-payouts) · Effect: **read**

```sh
sellapp affiliate-payouts get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| payout (positional) | Yes | The payout path parameter. |

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

Inspect the complete schema: `sellapp commands affiliate-payouts get --json`.

## list

List affiliate payouts

List merchant-managed affiliate payout records. These record commission payments arranged by the merchant; SellApp does not transfer or hold the funds. Requires the `affiliate` credential ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-payouts) · Effect: **read**

```sh
sellapp affiliate-payouts list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --affiliate-id | No | Filter by affiliate ID. |
| --status | No | Filter by payout status. |
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

Inspect the complete schema: `sellapp commands affiliate-payouts list --json`.

## update

Update affiliate payout status

Mark an affiliate payout record due or paid and update its assigned referrals. Marking paid records a payment arranged by the merchant; it does not transfer money. Identical retries make no further change. Requires `affiliate` and `affiliate:payout`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/update-payout-status) · Effect: **consequential**

```sh
sellapp affiliate-payouts update 1 --status paid --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| payout (positional) | Yes | The payout path parameter. |
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

Inspect the complete schema: `sellapp commands affiliate-payouts update --json`.

