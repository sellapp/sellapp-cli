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

List merchant-managed affiliate payout records, newest first, 15 per page by default. These record commission payments arranged by the merchant; SellApp does not transfer or hold the funds. Requires the `affiliate` credential ability. A `status` or `payout_method` value the API does not know returns `422`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-payouts) · Effect: **read**

```sh
sellapp affiliate-payouts list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --affiliate-id | No | Filter by affiliate ID. |
| --status | No | Filter by payout status. |
| --payout-method | No | List only payouts recorded for this payout method. Uppercase, as returned in `payout_method`; `paypal` returns `422`. |
| --search | No | The start of the affiliate's email address, such as `alex` or `alex.morgan@`. It matches from the first character only, and `%` and `_` match themselves. |
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

## mark-paid

Mark affiliate payouts paid

Record up to 20 due affiliate payouts as paid in one request, after you have paid the affiliates yourself. No money moves: SellApp only records the payment. Each payout marked paid also marks its referrals paid and emails that affiliate that the payout was paid. Send the payout IDs and `expected_total_usd_cents`, the sum of their `amount_usd_cents` as you read them, in integer USD cents (`4250` means $42.50). Every payout is checked first, and nothing changes when one of them is unknown in your store, is not `due`, or has referrals that no longer add up to it or can no longer be paid (`422` on `payout_ids`, one message per payout), or when the amounts add up to another total (`422` on `expected_total_usd_cents`). Then each payout is marked in the order sent, and `data` reports one `outcome` per payout: `paid`, `skipped`, or `failed`. A payout that changes between the check and its turn is skipped or fails without affecting the others. Requires the `affiliate` and `affiliate:payout` API-key abilities, and the staff member who owns the key needs the affiliate and affiliate payout permissions. There is no idempotency key, but sending the same batch again is refused with `422` on `payout_ids`, because those payouts are no longer due, so nobody is emailed twice. After a lost response, list the payouts with `GET /v2/affiliate-payouts?status=due` and send only those still due, with their new total. The emails cannot be taken back; to reopen a payout recorded by mistake, set it back to `due` with `PATCH /v2/affiliate-payouts/{payout}/status`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/mark-payouts-paid) · Effect: **consequential**

```sh
sellapp affiliate-payouts mark-paid --body '{"payout_ids":[91,92],"expected_total_usd_cents":4250}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --payout-ids | Yes | The due payouts to record as paid, at most 20, each once. They are marked in this order. |
| --expected-total-usd-cents | Yes | The sum of the payouts' `amount_usd_cents`, in integer USD cents, as you read them. If they add up to anything else, nothing is marked. |

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

Inspect the complete schema: `sellapp commands affiliate-payouts mark-paid --json`.

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

