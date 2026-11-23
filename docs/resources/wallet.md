# wallet

[All commands](../commands.md)

## adjust

Adjust a customer wallet

Add money to or take money from a customer's wallet and save the change as a `MANUAL_ADJUSTMENT` entry with your note. Amounts are integer USD cents: `2500` adds $25.00 and `-2500` takes $25.00 away. A first write returns `201`. The customer is not emailed. Adding money opens a wallet for a customer who has none yet. When the store's wallet settings have `expiration_days`, money you add expires after that many days, shown in the entry's `expires_at`. Requires the `wallet:write` API-key ability and the wallet permission in your team role. Send `expected_balance_cents` with the balance you last read to apply the change only while the wallet still holds exactly that balance. If the balance changed, the request returns `422` on `expected_balance_cents` and nothing is written; read the wallet again before deciding. A customer without a wallet counts as a balance of `0`. For safe retries, `idempotency_key` identifies one adjustment. Repeating the same key with the same customer, amount, and note (compared without leading and trailing spaces), made by the same staff member, returns `200` with the original entry and writes nothing, even when the balance has moved on since (the expected balance is not compared again). The same key with a different amount, note, or customer returns `422` on `wallet` and changes nothing. Keys are unique across all stores, are kept permanently, and cannot start with a reserved internal prefix such as `cashback:` or `topup_bonus:` (`422` on `idempotency_key`), so use a random value such as a UUID. Other refusals return `422` on `wallet` without writing: wallets are turned off for the store, the wallet is frozen, a debit is larger than the balance left after checkout holds, or a debit targets a customer who has no wallet yet. To undo an adjustment, reverse it with `POST /v2/wallets/{customer}/transactions/{transaction}/reverse`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/adjust-customer-wallet) · Effect: **consequential**

```sh
sellapp wallet adjust 1 --amount-cents 2500 --body-idempotency-key wallet-adjustment-01992a65 --note 'Launch-day account credit' --expected-balance-cents 0 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer path parameter. |
| --amount-cents | Yes | The signed change in integer USD cents: positive adds money, negative takes it away. Must not be zero. |
| --body-idempotency-key | Yes | Your key for this one adjustment, in printable ASCII characters and unique across all stores. Reuse it only to retry the same customer, amount, and note as the same staff member. |
| --note | Yes | Why the balance changed, with at least one character that is not a space. Leading and trailing spaces are removed. Saved on the entry and shown in the wallet history. |
| --expected-balance-cents | No | Optional. The wallet balance in integer USD cents that you read before deciding on this change. If the balance is different now, nothing is written and the request returns `422` on `expected_balance_cents`. Use `0` for a customer who has no wallet yet. |

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

Inspect the complete schema: `sellapp commands wallet adjust --json`.

## get

Retrieve a customer wallet

Retrieve one current-store customer wallet by customer ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/retrieve-customer-wallet) · Effect: **read**

```sh
sellapp wallet get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer path parameter. |

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

Inspect the complete schema: `sellapp commands wallet get --json`.

## ledger list

List store wallet transactions

List the wallet entries of every customer in your store, newest first, 15 per page by default, with totals for the whole filtered set in `meta.summary`. Amounts are signed integer USD cents: `2500` is $25.00 added and `-2500` is $25.00 taken away. Requires the `wallet` or `wallet:write` API-key ability and the wallet permission in your team role. `period` defaults to the current month so far (`mtd`), and `meta.period` shows the exact window. Narrow the list to one entry `type`, and to one customer with `customer_id` or `email`; when you send both, the ID is tried first. A customer of your store who has no wallet returns an empty `data` array and zero totals. A customer ID or email that is not a customer of your store returns `422` on the filter you sent. A `limit` above 100 returns `422`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/list-store-wallet-transactions) · Effect: **read**

```sh
sellapp wallet ledger list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --period | No | The time window: `7d` is the last 7 days, `30d` the last 30 days, `mtd` the current calendar month so far (from the first of the month, 00:00 UTC), and `all` everything. Defaults to `mtd`. |
| --type | No | Return only entries of this type. |
| --customer-id | No | Return only this customer's entries. It must be a customer of your store. |
| --email | No | Return only the entries of the customer with this exact email address in your store. |
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

Inspect the complete schema: `sellapp commands wallet ledger list --json`.

## list

List customer wallets

List customer wallet balances for the authenticated store. Requires the `wallet` or `wallet:write` API-key ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/list-customer-wallets) · Effect: **read**

```sh
sellapp wallet list
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

Inspect the complete schema: `sellapp commands wallet list --json`.

## overview get

Retrieve the wallet overview

Read your store's wallet health in one call: settings, what customers hold and can spend now, money moved and credit given away in a period, active incentive rules, the largest wallets, and the newest entries. Every amount is in integer USD cents: `2500` means $25.00. Requires the `wallet` or `wallet:write` API-key ability and the wallet permission in your team role. `health` is always the current state; `activity`, `incentive_spend`, and `recent_activity` follow `period`. `wallet_limit` and `activity_limit` take 1 to 20 and default to 5; a value outside that range returns `422`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/wallet-overview) · Effect: **read**

```sh
sellapp wallet overview get
```

| Input | Required | Meaning |
| --- | --- | --- |
| --period | No | The time window: `7d` is the last 7 days, `30d` the last 30 days, `mtd` the current calendar month so far (from the first of the month, 00:00 UTC), and `all` everything. Defaults to `mtd`. |
| --wallet-limit | No | How many of the largest wallets to return in `top_wallets`. |
| --activity-limit | No | How many of the newest entries to return in `recent_activity`. |

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

Inspect the complete schema: `sellapp commands wallet overview get --json`.

## settings get

Retrieve wallet settings

Retrieve wallet settings for the authenticated store. Requires the `wallet` or `wallet:write` API-key ability. Payment methods are returned as their effective configurable values when no explicit restriction is stored. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/wallet-settings) · Effect: **read**

```sh
sellapp wallet settings get
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

Inspect the complete schema: `sellapp commands wallet settings get --json`.

## settings replace

Update wallet settings

Update wallet settings. Requires the `wallet:write` API-key ability and the wallet permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/wallet-settings) · Effect: **consequential**

```sh
sellapp wallet settings replace --body '{"enabled":false,"minimum_top_up_cents":null,"maximum_top_up_cents":null,"expiration_days":null,"payment_methods":[]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --enabled | Yes | Value for enabled. |
| --minimum-top-up-cents | Yes | Value for minimum top up cents. |
| --maximum-top-up-cents | Yes | Value for maximum top up cents. |
| --expiration-days | Yes | Value for expiration days. |
| --payment-methods | Yes | Effective top-up gateway groups. CUSTOM_PAYMENT_METHOD enables all usable custom methods; customers choose an individual method in the portal. With no saved restriction, every currently configurable group is returned, including custom methods. An empty array disables all methods. Custom top-ups charge and credit the exact USD deposit amount, plus any wallet bonus, without custom-method discounts or fees. |

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

Inspect the complete schema: `sellapp commands wallet settings replace --json`.

## topups create

Create a wallet top-up payment link

Create a paid USD top-up for an existing customer and return data.checkout_url for the customer to pay. Creating a link leaves the wallet balance unchanged. Requires both wallet:write and charge API-key abilities and both store wallet and charge permissions. The wallet must be enabled and unfrozen, and the amount and selected method must satisfy the store's wallet settings. Payment completion credits the deposit and any bonus saved at creation exactly once; custom-method modifiers do not change the deposit. Browser returns and submitted payment proof do not prove payment. The return and cancel destination is the store's customer portal wallet. Currency, metadata, and redirect overrides are unsupported. Send an Idempotency-Key and reuse the same key and request for response replay within 24 hours. Changed input returns 409. For an uncertain provider outcome, inspect the charge before starting another top-up. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/create-wallet-top-up) · Effect: **consequential**

```sh
sellapp wallet topups create 42 --amount-cents 2500 --payment-method STRIPE --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | Value for customer. |
| --amount-cents | Yes | Deposit in integer USD cents: 2500 means $25.00. The store's minimum and maximum top-up amounts also apply. |
| --payment-method | Yes | One enabled payment gateway allowed by the store's wallet settings for USD. CUSTOM_PAYMENT_METHOD also requires custom_payment_method_id. |
| --custom-payment-method-id | No | Enabled custom method in this store. Required for CUSTOM_PAYMENT_METHOD; omit or use null for other gateways. |

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

Inspect the complete schema: `sellapp commands wallet topups create --json`.

## transactions list

List wallet transactions

Read the complete wallet history for one customer in the selected store, newest first. Requires the wallet or wallet:write API-key ability and store wallet permission. Amounts and balances are integer USD cents: 2500 means $25.00. A wallet with no entries returns an empty data array; an unknown wallet returns 404. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/retrieve-customer-wallet) · Effect: **read**

```sh
sellapp wallet transactions list 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | Value for customer. |
| --page | No | Value for page. |
| --limit | No | Value for limit. |

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

Inspect the complete schema: `sellapp commands wallet transactions list --json`.

## transactions reverse

Reverse a wallet transaction

Undo one top-up or manual adjustment in a customer's wallet. SellApp writes a `REVERSAL` entry for the opposite amount, in integer USD cents, and returns it with `201`. Reversing a `TOP_UP` also reverses its `TOP_UP_BONUS`, so the balance drops by the top-up plus its bonus; the returned entry reverses the top-up itself. Reversing a negative `MANUAL_ADJUSTMENT` gives the money back, except credit that the debit had spent and that has since been refunded: a clawback entry takes that part back. The customer is not emailed, and no payment is refunded: reversing a paid top-up does not return money to the customer's card. Requires the `wallet:write` API-key ability and the wallet permission in your team role. Only `TOP_UP` and `MANUAL_ADJUSTMENT` entries can be reversed, once each, and a reversal cannot itself be reversed. Every refusal returns `422` on `transaction` and writes nothing: the entry was already reversed, is a reversal, is an expired credit, or is a type only the system reverses (`TOP_UP_BONUS` goes with its top-up, `CASHBACK` and `REWARD` are reversed by the system, a `REFUND` on its support ticket, a `PURCHASE_DEBIT` through its order); wallets are turned off for the store; or the money is reserved by a checkout or already spent. A customer of another store, a customer without a wallet, or an entry from another customer's wallet returns `404`. Send `expected_balance_cents` with the balance you read to reverse only while the wallet still holds it; otherwise the request returns `422` on `expected_balance_cents` and nothing changes. Repeating the request never takes the money twice: without a key, the repeat returns `422` because the entry is already reversed. To get the original `201` response back instead, send an `Idempotency-Key` header and repeat the same key with the same body within 24 hours; the replay carries `Idempotent-Replayed: true`. The same key with a different body returns `409`. To put money back after a reversal, make a new manual adjustment. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/reverse-wallet-transaction) · Effect: **consequential**

```sh
sellapp wallet transactions reverse --customer 42 501 --note 'Credit went to the wrong customer' --expected-balance-cents 2500 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer | Yes | The customer ID whose wallet holds the entry. |
| transaction (positional) | Yes | The ID of the wallet entry to reverse, as listed by `GET /v2/wallets/{customer}/transactions` or `GET /v2/wallet/transactions`. |
| --note | Yes | Why the entry is reversed. It must contain more than spaces and is saved on the reversal entry. |
| --expected-balance-cents | No | Optional. The wallet balance in integer USD cents that you read before deciding. If the balance is different now, nothing changes and the request returns `422` on `expected_balance_cents`. |

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

Inspect the complete schema: `sellapp commands wallet transactions reverse --json`.

## update-status

Update a customer wallet status

Freeze or activate a customer's existing wallet. A frozen wallet cannot be adjusted. Activating a frozen wallet first expires any credit whose expiry date has passed, so the returned `balance_cents` can be lower than before. Sending the status the wallet already has changes nothing, so a retry is safe. Requires the `wallet:write` API-key ability and the wallet permission in your team role. Refusals return `422` on `wallet` without writing: wallets are turned off for the store, or the customer has no wallet yet. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/update-wallet-status) · Effect: **consequential**

```sh
sellapp wallet update-status 1 --status frozen --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer path parameter. |
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

Inspect the complete schema: `sellapp commands wallet update-status --json`.

