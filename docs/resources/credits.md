# credits

[All commands](../commands.md)

## balances get

Retrieve a credit balance

Retrieve one customer and credits-product balance with up to 100 recent ledger entries. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/retrieve-a-credit-balance) · Effect: **read**

```sh
sellapp credits balances get --customer 1 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer | Yes | The customer path parameter. |
| creditProduct (positional) | Yes | The creditProduct path parameter. |

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

Inspect the complete schema: `sellapp commands credits balances get --json`.

## balances list

List credit balances

List store-scoped customer balances for dedicated credits products. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/list-credit-balances) · Effect: **read**

```sh
sellapp credits balances list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --customer-id | No | Filter by a store customer ID. |
| --product-id | No | Filter by a credits product ID. |

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

Inspect the complete schema: `sellapp commands credits balances list --json`.

## balances transactions list

List credit balance transactions

Read all credit transactions for one customer and credits product in the selected store, newest first. Requires the credit API-key ability and store credit permission. Amounts are integer credit units. Use links.next to continue beyond the 100 entries embedded in balance details. An existing balance with no entries returns an empty data array; an unknown balance returns 404. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/retrieve-a-credit-balance) · Effect: **read**

```sh
sellapp credits balances transactions list --customer 42 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer | Yes | Value for customer. |
| creditProduct (positional) | Yes | Value for creditProduct. |
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

Inspect the complete schema: `sellapp commands credits balances transactions list --json`.

## products create

Create a credits product

Create a credits product with one default CREDITS variant. Published products require at least one ordered, non-overlapping rate tier. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/create-a-credit-product) · Effect: **consequential**

```sh
sellapp credits products create --title 'Design credits' --visibility HIDDEN --price-cents 1999 --currency USD --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Value for title. |
| --visibility | Yes | Value for visibility. |
| --slug | No | Value for slug. |
| --description | No | Value for description. |
| --section-id | No | Value for section id. |
| --is-draft | No | Value for is draft. |
| --price-cents | No | Base price in integer minor currency units. |
| --currency | No | Value for currency. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --quantity-increment | No | Value for quantity increment. |
| --stock | No | Value for stock. |
| --payment-methods | No | Value for payment methods. |
| --rate-tiers | No | Value for rate tiers. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands credits products create --json`.

## products delete

Delete a credits product

Soft-delete a store-scoped credits product. Optional expected_updated_at protects against stale writes (422). Active subscriptions, any credit balance history, or pending purchase fulfillment prevent deletion (422). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/delete-a-credit-product) · Effect: **consequential**

```sh
sellapp credits products delete 1 --expected-updated-at 2026-08-24T10:00:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| creditProduct (positional) | Yes | The creditProduct path parameter. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands credits products delete --json`.

## products get

Retrieve a credits product

Retrieve a store-scoped credits product. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/retrieve-a-credit-product) · Effect: **read**

```sh
sellapp credits products get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| creditProduct (positional) | Yes | The creditProduct path parameter. |

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

Inspect the complete schema: `sellapp commands credits products get --json`.

## products list

List credits products

List dedicated credits products and their single default variants. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/list-credit-products) · Effect: **read**

```sh
sellapp credits products list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |

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

Inspect the complete schema: `sellapp commands credits products list --json`.

## products replace

Replace a credits product

Update the credits product and its single default variant atomically. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/replace-a-credit-product) · Effect: **consequential**

```sh
sellapp credits products replace 1 --title 'Design credits' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| creditProduct (positional) | Yes | The creditProduct path parameter. |
| --title | No | Value for title. |
| --slug | No | Value for slug. |
| --description | No | Value for description. |
| --visibility | No | Value for visibility. |
| --section-id | No | Value for section id. |
| --is-draft | No | Value for is draft. |
| --price-cents | No | Base price in integer minor currency units. |
| --currency | No | Value for currency. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --quantity-increment | No | Value for quantity increment. |
| --stock | No | Value for stock. |
| --payment-methods | No | Value for payment methods. |
| --rate-tiers | No | Value for rate tiers. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands credits products replace --json`.

## products search

Search credits products

Search credits products by title, slug, or description. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/search-credit-products) · Effect: **read**

```sh
sellapp credits products search
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --search | No | Value for search. |
| --includes | No | Value for includes. |

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

Inspect the complete schema: `sellapp commands credits products search --json`.

## products update

Update a credits product

Update the credits product and its single default variant atomically. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/update-a-credit-product) · Effect: **consequential**

```sh
sellapp credits products update 1 --title 'Design credits' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| creditProduct (positional) | Yes | The creditProduct path parameter. |
| --title | No | Value for title. |
| --slug | No | Value for slug. |
| --description | No | Value for description. |
| --visibility | No | Value for visibility. |
| --section-id | No | Value for section id. |
| --is-draft | No | Value for is draft. |
| --price-cents | No | Base price in integer minor currency units. |
| --currency | No | Value for currency. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --quantity-increment | No | Value for quantity increment. |
| --stock | No | Value for stock. |
| --payment-methods | No | Value for payment methods. |
| --rate-tiers | No | Value for rate tiers. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands credits products update --json`.

## record

Record a credit transaction

Atomically grant, consume, or administratively adjust integer credit units. Reusing the same idempotency key with the same transaction is a no-op; conflicting reuse is rejected. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/credits/record-a-credit-transaction) · Effect: **consequential**

```sh
sellapp credits record --customer-id 125 --product-id 120 --kind grant --amount-units 1000 --body-idempotency-key credits-grant-01992a65 --reason 'Launch cohort allocation' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer-id | Yes | Value for customer id. |
| --product-id | Yes | Value for product id. |
| --kind | Yes | Value for kind. |
| --amount-units | Yes | The signed amount of credit units. Must not be zero. |
| --body-idempotency-key | Yes | Value for idempotency key. |
| --reason | No | Value for reason. |
| --source-type | No | Value for source type. |
| --source-id | No | Value for source id. |
| --metadata | No | Value for metadata. |

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

Inspect the complete schema: `sellapp commands credits record --json`.

