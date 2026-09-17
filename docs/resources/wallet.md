# wallet

[All commands](../commands.md)

## adjust

Adjust a customer wallet

Add or subtract a nonzero amount from a customer wallet. Requires the wallet:write API ability and store wallet permission. Use signed integer cents and a nonblank note. Positive adjustments can create a wallet and follow the store's credit-expiration policy; negative adjustments need an existing eligible wallet and sufficient balance. Wallets must be enabled. Send idempotency_key in the body. Replaying the same key returns the original entry only for the same store, customer, staff actor, amount and trimmed note; changed input returns 422. Keep that identity unchanged across retries. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/adjust-customer-wallet) · Effect: **consequential**

```sh
sellapp wallet adjust 1 --amount-cents 2500 --body-idempotency-key wallet-adjustment-01992a65 --note 'Launch-day account credit' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer path parameter. |
| --amount-cents | Yes | The signed adjustment amount in cents. Must not be zero. |
| --body-idempotency-key | Yes | Required body key, at most 128 characters. Keep the customer, amount, staff actor and trimmed note unchanged when retrying this intended adjustment. |
| --note | Yes | Value for note. |

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

## list

List customer wallets

List customer wallet balances for the authenticated store. Requires the `wallet` or `wallet:write` API ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

## settings get

Retrieve wallet settings

Retrieve wallet settings for the authenticated store. Requires the `wallet` or `wallet:write` API ability. Payment methods are returned as their effective configurable values when no explicit restriction is stored. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Update wallet settings. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

## update-status

Update a customer wallet status

Freeze or activate an existing customer wallet. Requires wallet:write and the store wallet permission, with wallets enabled. Activating a frozen wallet expires eligible credits before returning, so its balance may decrease. Read balance_cents from the returned wallet. Repeating the current status leaves the wallet unchanged. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

