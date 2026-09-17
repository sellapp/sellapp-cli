# upsell-offers

[All commands](../commands.md)

## create

Create an upsell offer

Create an ordered post-purchase offer. Every product and variant ID is checked against the selected store and expected parent product. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **consequential**

```sh
sellapp upsell-offers create --body '{"name":"One More Feature","description":"Offer the advanced package with additional resources.","is_active":true,"source_listing_id":120,"source_variant_id":880,"minimum_order_total_usd_cents":1000,"maximum_order_total_usd_cents":25000,"available_for_days":14,"max_accepts_per_customer":1,"items":[{"target_listing_id":121,"target_variant_id":881,"headline":"Upgrade today","description":"Unlock dark mode, webhooks, and the premium launch checklist.","discount_type":"percentage","discount_value":"15.00","maximum_discount_amount":"25.00"}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Value for name. |
| --is-active | Yes | Value for is active. |
| --source-listing-id | Yes | Value for source listing id. |
| --items | Yes | Complete ordered target-item set. Target variant IDs must be unique and belong to their target product in this store. |
| --description | No | Value for description. |
| --source-variant-id | No | Value for source variant id. |
| --minimum-order-total-usd-cents | No | Integer USD cents, for example 1000 for USD 10.00. |
| --maximum-order-total-usd-cents | No | Integer USD cents. When present, must be greater than or equal to minimum_order_total_usd_cents. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | Must be after or equal to starts_at. |
| --available-for-days | No | Value for available for days. |
| --max-accepts-per-customer | No | Value for max accepts per customer. |
| --expected-version | No | Required optimistic concurrency token for updates. Use the version from the latest response; stale versions are rejected with 422. |

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

Inspect the complete schema: `sellapp commands upsell-offers create --json`.

## delete

Delete an upsell offer

Soft-delete an upsell offer while preserving checkout and reporting attribution. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **consequential**

```sh
sellapp upsell-offers delete 71 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| upsellOffer (positional) | Yes | The upsellOffer path parameter. |

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

Inspect the complete schema: `sellapp commands upsell-offers delete --json`.

## get

Retrieve an upsell offer

Retrieve one store-scoped upsell offer and its ordered target items. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **read**

```sh
sellapp upsell-offers get 71
```

| Input | Required | Meaning |
| --- | --- | --- |
| upsellOffer (positional) | Yes | The upsellOffer path parameter. |

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

Inspect the complete schema: `sellapp commands upsell-offers get --json`.

## list

List upsell offers

List the authenticated store's ordered post-purchase offers. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **read**

```sh
sellapp upsell-offers list
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

Inspect the complete schema: `sellapp commands upsell-offers list --json`.

## replace

Update an upsell offer

Update supplied fields. When items are supplied, they replace the complete ordered target-item set. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **consequential**

```sh
sellapp upsell-offers replace 71 --name 'One More Feature 2.0' --is-active false --expected-version 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| upsellOffer (positional) | Yes | The upsellOffer path parameter. |
| --expected-version | Yes | Required optimistic concurrency token for updates. Use the version from the latest response; stale versions are rejected with 422. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --is-active | No | Value for is active. |
| --source-listing-id | No | Value for source listing id. |
| --source-variant-id | No | Value for source variant id. |
| --minimum-order-total-usd-cents | No | Integer USD cents, for example 1000 for USD 10.00. |
| --maximum-order-total-usd-cents | No | Integer USD cents. When present, must be greater than or equal to minimum_order_total_usd_cents. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | Must be after or equal to starts_at. |
| --available-for-days | No | Value for available for days. |
| --max-accepts-per-customer | No | Value for max accepts per customer. |
| --items | No | Complete ordered target-item set. Target variant IDs must be unique and belong to their target product in this store. |

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

Inspect the complete schema: `sellapp commands upsell-offers replace --json`.

## search

Search upsell offers

Search offer names and descriptions and compose supported filters and sort instructions. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **read**

```sh
sellapp upsell-offers search --body '{"filters":[{"field":"id","operator":"=","value":71}],"sort":[{"field":"created_at","direction":"desc"}]}'
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

Inspect the complete schema: `sellapp commands upsell-offers search --json`.

## update

Update an upsell offer

Update supplied fields. When items are supplied, they replace the complete ordered target-item set. Requires a Sanctum token with the `upsell` ability and the Upsells role for the store selected by `X-STORE`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **consequential**

```sh
sellapp upsell-offers update 71 --name 'One More Feature 2.0' --is-active false --expected-version 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| upsellOffer (positional) | Yes | The upsellOffer path parameter. |
| --expected-version | Yes | Required optimistic concurrency token for updates. Use the version from the latest response; stale versions are rejected with 422. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --is-active | No | Value for is active. |
| --source-listing-id | No | Value for source listing id. |
| --source-variant-id | No | Value for source variant id. |
| --minimum-order-total-usd-cents | No | Integer USD cents, for example 1000 for USD 10.00. |
| --maximum-order-total-usd-cents | No | Integer USD cents. When present, must be greater than or equal to minimum_order_total_usd_cents. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | Must be after or equal to starts_at. |
| --available-for-days | No | Value for available for days. |
| --max-accepts-per-customer | No | Value for max accepts per customer. |
| --items | No | Complete ordered target-item set. Target variant IDs must be unique and belong to their target product in this store. |

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

Inspect the complete schema: `sellapp commands upsell-offers update --json`.

