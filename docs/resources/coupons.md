# coupons

[All commands](../commands.md)

## batch-create

Batch create coupons

Create multiple coupons in one request by sending a `resources` array.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons batch-create --body '{"resources":[{"code":"STARTER10","type":"PERCENTAGE","discount":10,"store_wide":false,"products":[123],"product_variants":[1001]}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --resources | Yes | Value for resources. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons batch-create --json`.

## batch-delete

Batch delete coupons

Soft-delete multiple coupons by sending their IDs in `resources`.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons batch-delete --body '{"resources":[1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --resources | Yes | Value for resources. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons batch-delete --json`.

## batch-update

Batch update coupons

Update multiple coupons in one request by sending a `resources` object keyed by coupon ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons batch-update --body '{"resources":{"1":{"store_wide":false,"products":[123],"product_variants":[1001,1002]}}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --resources | Yes | Value for resources. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons batch-update --json`.

## create

Create a coupon

Create a discount code for your store.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons create --body '{"code":"PLAN10","type":"PERCENTAGE","discount":10,"store_wide":false,"products":[123,456],"product_variants":[1001,1002]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --code | Yes | Value for code. |
| --type | Yes | Value for type. |
| --discount | Yes | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. |
| --store-wide | Yes | Value for store wide. |
| --products | No | Product IDs the coupon applies to when store_wide is false. |
| --product-variants | No | Optional variant restrictions. Every variant must belong to a selected product. Products without listed variants remain eligible on all variants. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum order amount in the store currency's major unit, or null for no minimum. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons create --json`.

## delete

Delete a coupon

Soft-delete a coupon. The returned coupon includes its deleted_at timestamp; use with_trashed or only_trashed to retrieve soft-deleted coupons later.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons delete 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons delete --json`.

## get

Retrieve a coupon

Retrieve a coupon by its ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp coupons get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons get --json`.

## list

List all coupons

List your store's coupons, 15 per page by default.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp coupons list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons list --json`.

## replace

Update a coupon

Change a coupon's discount, limits, or eligible products.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons replace 1 --body '{"store_wide":false,"products":[123],"product_variants":[1001,1002]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --code | No | Value for code. |
| --type | No | Value for type. |
| --discount | No | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. |
| --store-wide | No | Value for store wide. |
| --products | No | Product IDs the coupon applies to when store_wide is false. |
| --product-variants | No | Optional variant restrictions. Omit to preserve existing restrictions or send an empty array to allow every variant of the selected products. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum order amount in the store currency's major unit, or null for no minimum. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons replace --json`.

## search

Search coupons

Search coupons using JSON body filters, search terms, includes, and sort instructions.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp coupons search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons search --json`.

## update

Update a coupon

Change a coupon's discount, limits, or eligible products.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp coupons update 1 --body '{"store_wide":false,"products":[123],"product_variants":[1001,1002]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --code | No | Value for code. |
| --type | No | Value for type. |
| --discount | No | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. |
| --store-wide | No | Value for store wide. |
| --products | No | Product IDs the coupon applies to when store_wide is false. |
| --product-variants | No | Optional variant restrictions. Omit to preserve existing restrictions or send an empty array to allow every variant of the selected products. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum order amount in the store currency's major unit, or null for no minimum. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands coupons update --json`.

## v-2-batch-create-coupons

Batch create coupons

Create multiple coupons in one request by sending a `resources` array. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/batch-create-coupons) · Effect: **consequential**

```sh
sellapp coupons v-2-batch-create-coupons --body '{"resources":[{"code":"STARTER10","type":"PERCENTAGE","discount":10,"store_wide":false,"products":[123],"product_variants":[1001]}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands coupons v-2-batch-create-coupons --json`.

## v-2-batch-delete-coupons

Batch delete coupons

Soft-delete multiple coupons by sending their IDs in `resources`. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/batch-delete-coupons) · Effect: **consequential**

```sh
sellapp coupons v-2-batch-delete-coupons --body '{"resources":[1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands coupons v-2-batch-delete-coupons --json`.

## v-2-batch-update-coupons

Batch update coupons

Update multiple coupons in one request by sending a `resources` object keyed by coupon ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/batch-update-coupons) · Effect: **consequential**

```sh
sellapp coupons v-2-batch-update-coupons --body '{"resources":{"1":{"store_wide":false,"products":[123],"product_variants":[1001,1002]}}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands coupons v-2-batch-update-coupons --json`.

## v-2-create-coupon

Create a coupon

Create a discount code for your store. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/create-a-coupon) · Effect: **consequential**

```sh
sellapp coupons v-2-create-coupon --body '{"code":"PLAN10","type":"PERCENTAGE","discount":10,"store_wide":false,"products":[123,456],"product_variants":[1001,1002]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --code | Yes | Value for code. |
| --type | Yes | Value for type. |
| --discount | Yes | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. |
| --store-wide | Yes | Value for store wide. |
| --products | No | Product IDs the coupon applies to when store_wide is false. |
| --product-variants | No | Optional variant restrictions. Every variant must belong to a selected product. Products without listed variants remain eligible on all variants. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum order amount in the store currency's major unit, or null for no minimum. |

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

Inspect the complete schema: `sellapp commands coupons v-2-create-coupon --json`.

## v-2-delete-coupon

Delete a coupon

Soft-delete a coupon. The returned coupon includes its deleted_at timestamp; use with_trashed or only_trashed to retrieve soft-deleted coupons later. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/delete-a-coupon) · Effect: **consequential**

```sh
sellapp coupons v-2-delete-coupon 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |

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

Inspect the complete schema: `sellapp commands coupons v-2-delete-coupon --json`.

## v-2-get-coupon

Retrieve a coupon

Retrieve a coupon by its ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/retrieve-a-coupon) · Effect: **read**

```sh
sellapp coupons v-2-get-coupon 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |

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

Inspect the complete schema: `sellapp commands coupons v-2-get-coupon --json`.

## v-2-list-coupons

List all coupons

List your store's coupons, 15 per page by default. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/list-all-coupons) · Effect: **read**

```sh
sellapp coupons v-2-list-coupons
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |

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

Inspect the complete schema: `sellapp commands coupons v-2-list-coupons --json`.

## v-2-replace-coupon

Update a coupon

Change a coupon's discount, limits, or eligible products. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/update-a-coupon) · Effect: **consequential**

```sh
sellapp coupons v-2-replace-coupon 1 --body '{"store_wide":false,"products":[123],"product_variants":[1001,1002]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --code | No | Value for code. |
| --type | No | Value for type. |
| --discount | No | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. |
| --store-wide | No | Value for store wide. |
| --products | No | Product IDs the coupon applies to when store_wide is false. |
| --product-variants | No | Optional variant restrictions. Omit to preserve existing restrictions or send an empty array to allow every variant of the selected products. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum order amount in the store currency's major unit, or null for no minimum. |

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

Inspect the complete schema: `sellapp commands coupons v-2-replace-coupon --json`.

## v-2-search-coupons

Search coupons

Search coupons using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/search-coupons) · Effect: **read**

```sh
sellapp coupons v-2-search-coupons --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
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

Inspect the complete schema: `sellapp commands coupons v-2-search-coupons --json`.

## v-2-update-coupon

Update a coupon

Change a coupon's discount, limits, or eligible products. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/coupons/update-a-coupon) · Effect: **consequential**

```sh
sellapp coupons v-2-update-coupon 1 --body '{"store_wide":false,"products":[123],"product_variants":[1001,1002]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| coupon (positional) | Yes | The coupon path parameter. |
| --with-trashed | No | Include soft-deleted resources in the result. |
| --only-trashed | No | Return only soft-deleted resources. Ignored when with_trashed is true. |
| --code | No | Value for code. |
| --type | No | Value for type. |
| --discount | No | A decimal discount value of at least 1. Percentage discounts cannot exceed 100; amount discounts use the store currency's major unit. |
| --store-wide | No | Value for store wide. |
| --products | No | Product IDs the coupon applies to when store_wide is false. |
| --product-variants | No | Optional variant restrictions. Omit to preserve existing restrictions or send an empty array to allow every variant of the selected products. |
| --limit | No | Value for limit. |
| --expires-at | No | A future date and time, or null for no expiry. |
| --minimum-amount | No | A decimal minimum order amount in the store currency's major unit, or null for no minimum. |

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

Inspect the complete schema: `sellapp commands coupons v-2-update-coupon --json`.

