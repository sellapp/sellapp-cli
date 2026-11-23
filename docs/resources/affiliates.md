# affiliates

[All commands](../commands.md)

## coupon disable

Disable an affiliate coupon

Turn off the affiliate's coupon so customers can no longer use it. The coupon is kept, still tied to the affiliate, and its code stays reserved; the response shows `status: disabled`. To turn it back on, send `PATCH /v2/affiliates/{affiliate}/coupon` (an empty body works). When the affiliate's membership is active, they are emailed that their coupon is disabled. Requires the `affiliate`, `affiliate:write`, and `coupon` API-key abilities, and the staff member who owns the key needs the affiliate, coupon, and `store` permissions. Safe to repeat: disabling a coupon that is already disabled returns `200`, changes nothing, and sends no email. An affiliate with no coupon returns `422` on `affiliate`; an affiliate of another store returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/affiliate-coupon) · Effect: **consequential**

```sh
sellapp affiliates coupon disable 88 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate ID in your store. |

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

Inspect the complete schema: `sellapp commands affiliates coupon disable --json`.

## coupon get

Retrieve an affiliate coupon

Read the coupon tied to one affiliate, the code they share with customers, and whether it is enabled. An affiliate without a coupon returns `200` with `status: not_configured` and `coupon: null`. Requires the `affiliate` API-key ability and the staff affiliate permission. An affiliate of another store returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/affiliate-coupon) · Effect: **read**

```sh
sellapp affiliates coupon get 88
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate ID in your store. |

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

Inspect the complete schema: `sellapp commands affiliates coupon get --json`.

## coupon save

Save an affiliate coupon

Create the affiliate's coupon, or change it, in one call. When the affiliate has no coupon, send `code`, `type`, `discount`, and `store_wide`; each one you leave out returns `422` on that field. When a coupon exists, send only the fields to change; the others stay as they are. Any request also enables a coupon that was disabled, so `{}` re-enables it unchanged. A store-wide coupon returns `products: []`, and you can send a read coupon back unchanged. Coupon rules apply as for `PATCH /v2/coupons/{coupon}`: the code must be free in your store, `products` and `product_variants` must be yours, `expires_at` must be in the future, and a percentage `discount` cannot exceed 100; each refusal is keyed by that field. Amounts are decimals in the store currency's major unit. Re-enabling a disabled coupon emails the affiliate, when their membership is active, that their coupon is enabled again; creating or changing a coupon sends no email. Requires the `affiliate`, `affiliate:write`, and `coupon` API-key abilities, and the staff member who owns the key needs the affiliate, coupon, and `store` permissions. Safe to repeat: after a lost response, the same request leaves the same coupon, sends no second email, and never creates a second coupon. An affiliate of another store returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/affiliate-coupon) · Effect: **consequential**

```sh
sellapp affiliates coupon save 88 --code ALEX15 --type PERCENTAGE --discount 15 --store-wide true --maximum-discount-amount 20 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate ID in your store. |
| --code | No | Value for code. |
| --type | No | Value for type. |
| --discount | No | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. A `subscriptions` coupon cannot take 100% off: that returns `422` on `discount`, so offer a free trial instead. |
| --store-wide | No | Value for store wide. |
| --products | No | Product IDs the coupon applies to. Required with at least one ID when `store_wide` is false; send `[]` (or leave it out) for a store-wide coupon. |
| --product-variants | No | Optional variant restrictions. Every variant must belong to a selected product. Products without listed variants remain eligible on all variants. Each variant you add must be of the coupon's kind, a subscription plan or a one-time variant; another returns `422` on `product_variants`. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum amount in the store currency's major unit that the order lines of the coupon's kind must reach, or null for no minimum. |
| --applies-to | No | What the coupon discounts: `one_time` for one-time purchases, or `subscriptions` for the first payments of a subscription plan. A new coupon defaults to `one_time`, and an update that leaves this out keeps the current kind. The coupon only discounts order lines of its kind. |
| --subscription-payments | No | How many payments a `subscriptions` coupon discounts, counting the signup payment as the first, so `3` discounts the signup and the next two renewals. A `subscriptions` coupon needs one, and an update that leaves it out keeps the stored number. Missing on a `subscriptions` coupon, or outside 1 to 36, it returns `422` on `subscription_payments` with `The number of discounted payments must be between 1 and 36.` A `one_time` coupon always stores null here. |
| --maximum-discount-amount | No | Optional cap on the discount of a percentage coupon, as a decimal string of at least `"1"` in the store currency's major unit, such as `"20"`. `null` removes the cap. Sending it for an `AMOUNT` coupon returns `422`, and changing a coupon to `AMOUNT` removes its cap. |

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

Inspect the complete schema: `sellapp commands affiliates coupon save --json`.

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

List store affiliate memberships, approved profile details, referral counts, and integer USD-cent balances, newest first, 15 per page by default. Requires the `affiliate` ability. Filter by `status`, `search`, `custom_rates_only`, and `payout_method`, and send `sort=-earnings` to list the top earners first. A `status`, `custom_rates_only`, or `payout_method` value the API does not know returns `422`; an unknown `sort` value is ignored. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/list-affiliates) · Effect: **read**

```sh
sellapp affiliates list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --status | No | Filter by affiliate program status. |
| --search | No | Search affiliate email addresses and profile identifiers. |
| --custom-rates-only | No | Send `true` to list only affiliates with custom per-product commission rates (see `GET /v2/affiliates/{affiliate}/products`). `false` lists everyone. Accepts `true`, `false`, `1`, or `0`. |
| --payout-method | No | List only affiliates who chose this payout method. Uppercase, as returned in `payout_method`. |
| --sort | No | Send `-earnings` to list the affiliates who earned the most first: available, pending, and paid balances added together, largest first. Without it, or with any other value, the newest affiliates come first. |
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

## products get

Retrieve affiliate products

See whether one affiliate has custom commission rates and, if so, the product variants they earn on and the rate for each. When `enable_specific_products` is false, `products` is empty and the program's rates from `GET /v2/affiliate-program` apply. Fixed commissions are integer USD cents; percentages are decimal strings. Requires the `affiliate` API-key ability and the staff affiliate permission. An affiliate of another store returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/affiliate-products) · Effect: **read**

```sh
sellapp affiliates products get 88
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate ID in your store. |

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

Inspect the complete schema: `sellapp commands affiliates products get --json`.

## products update

Update affiliate products

Give one affiliate custom commission rates on chosen product variants, change those rates, or switch back to the program's rates. Send at least one of `enable_specific_products`, `products`, or `replace_products`; a field you leave out keeps what the affiliate has. With custom rates on, each entry in `products` sets the rate for one variant (`enabled: true` with a `commission`) or removes the variant (`enabled: false`); variants you do not name keep their rates. Add `replace_products: true` to make `products` the complete list instead. Turning custom rates on without `replace_products` keeps the program's current rate on every variant you do not name. `enable_specific_products: false` removes every custom rate, so the program's rates apply again; sending `products` with it returns `422`. Turning custom rates on for an affiliate who did not have them emails the affiliate that they have custom rates and emails your store a confirmation; other changes send nothing. Requires the `affiliate` and `affiliate:write` API-key abilities, and the staff member who owns the key needs both the affiliate and the `store` permission. Safe to repeat: the same request leaves the same rates and does not email again. Errors are keyed `products` (a variant that is not a published variant of your store, a repeated variant, or rates while custom rates stay off), `products.{i}.id`, `products.{i}.commission` (missing for an enabled variant), and `products.{i}.commission.type`, `products.{i}.commission.percentage` (0.01 to 100), or `products.{i}.commission.amount_usd_cents` (at least 1). An affiliate of another store returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliates/affiliate-products) · Effect: **consequential**

```sh
sellapp affiliates products update 88 --body '{"enable_specific_products":true,"replace_products":true,"products":[{"id":42,"enabled":true,"commission":{"type":"percentage","percentage":"25"}},{"id":43,"enabled":true,"commission":{"type":"fixed","amount_usd_cents":550}}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| affiliate (positional) | Yes | The affiliate ID in your store. |
| --enable-specific-products | No | `true` gives the affiliate custom rates; `false` removes them all and the program's rates apply. Leave it out to keep the current setting. |
| --products | No | Rates to set or remove, one entry per product variant. Each variant may appear once. |
| --replace-products | No | `true` removes every custom rate not in `products`, making `products` the complete list. |

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

Inspect the complete schema: `sellapp commands affiliates products update --json`.

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

