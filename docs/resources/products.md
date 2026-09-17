# products

[All commands](../commands.md)

## addons list

List a product's add-ons

Retrieve the ordered add-ons assigned to a product, course, or booking. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/list-product-add-ons) · Effect: **read**

```sh
sellapp products addons list 120
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
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

Inspect the complete schema: `sellapp commands products addons list --json`.

## addons replace

Replace a product's add-ons

Replace the complete ordered add-on assignment for a non-subscription product, course, or booking. Newly assigned add-ons must be published, public, and owned by the same store; already assigned non-public add-ons may be retained for safe editing. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/replace-product-add-ons) · Effect: **consequential**

```sh
sellapp products addons replace 120 --body '{"resources":[411,410]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
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

Inspect the complete schema: `sellapp commands products addons replace --json`.

## batch-create

Batch create products

Create multiple products in one request by sending a `resources` array. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/batch-create-products) · Effect: **consequential**

```sh
sellapp products batch-create --body '{"resources":[{"title":"Example product","description":"An example product created through the API.","visibility":"PUBLIC","type":"product"}]}' --yes
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

Inspect the complete schema: `sellapp commands products batch-create --json`.

## batch-delete

Batch delete products

Delete multiple products in one request by sending the product IDs in `resources`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/batch-delete-products) · Effect: **consequential**

```sh
sellapp products batch-delete --body '{"resources":[1,2]}' --yes
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

Inspect the complete schema: `sellapp commands products batch-delete --json`.

## batch-update

Batch update products

Update multiple products in one request by sending a `resources` object keyed by product ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/batch-update-products) · Effect: **consequential**

```sh
sellapp products batch-update --body '{"resources":{"1":{"title":"Updated product","visibility":"PUBLIC"}}}' --yes
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

Inspect the complete schema: `sellapp commands products batch-update --json`.

## create

Create a product

Create a product in the selected store. Requires a token with the `listing` ability. The same product rules apply as in the dashboard. Add a variant before offering the product at checkout. Include variants to create the product and its initial variants atomically. For a complete bundle, set type to bundle, supply one BUNDLE variant with empty deliverable.data, and include bundle_items. The response includes the created variant IDs in data.variants. Omit these fields to retain separate creation. These options apply to single-product creation, not batch creation. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/create-a-product) · Effect: **write**

```sh
sellapp products create --title 'Design kit' --description 'Templates for your next project.' --visibility HIDDEN
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Value for title. |
| --description | Yes | Value for description. |
| --visibility | Yes | Value for visibility. |
| --slug | No | Value for slug. |
| --type | No | Value for type. |
| --section | No | Value for section. |
| --additional-information | No | Value for additional information. |
| --other-settings | No | Value for other settings. |
| --variants | No | Optional initial variants, using the same fields as variant creation. A bundle accepts exactly one variant. A validation failure rolls back the product and every variant. |
| --bundle-items | No | Only for type bundle, with exactly one initial variant. Same-store standard product variants to include. Each variant ID must be unique. |

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

Inspect the complete schema: `sellapp commands products create --json`.

## delete

Delete a product

Soft-delete a product in the authenticated store. Products with active or past-due subscriptions cannot be deleted. Send `expected_updated_at` to reject a stale deletion. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/delete-a-product) · Effect: **consequential**

```sh
sellapp products delete 1 --expected-updated-at 2026-08-01T12:00:00Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot. Send the exact `updated_at` value from the product you last retrieved; the mutation is rejected if the product has changed. |

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

Inspect the complete schema: `sellapp commands products delete --json`.

## get

Retrieve a product

Retrieve a product by its ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/retrieve-a-product) · Effect: **read**

```sh
sellapp products get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --with-drafts | No | Include draft products alongside published products. |
| --only-drafts | No | Return only draft products. |

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

Inspect the complete schema: `sellapp commands products get --json`.

## list

List all products

List your store's products, 15 per page by default. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/list-all-products) · Effect: **read**

```sh
sellapp products list --limit 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --with-drafts | No | Include draft products alongside published products. |
| --only-drafts | No | Return only draft products. |

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

Inspect the complete schema: `sellapp commands products list --json`.

## replace

Update a product

Update a product in the authenticated store. Requires a token with the `listing` ability. Send `expected_updated_at` to reject stale writes instead of overwriting a newer change. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **write**

```sh
sellapp products replace 120 --title 'Design kit' --description 'Templates for your next project.' --visibility HIDDEN --expected-updated-at 2026-08-30T12:00:00.000000Z
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --title | No | Value for title. |
| --description | No | Value for description. |
| --visibility | No | Value for visibility. |
| --slug | No | Value for slug. |
| --section | No | Value for section. |
| --additional-information | No | Value for additional information. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot. Send the exact `updated_at` value from the product you last retrieved; the mutation is rejected if the product has changed. |

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

Inspect the complete schema: `sellapp commands products replace --json`.

## search

Search products

Search products using JSON body filters, search terms, includes, and sort instructions. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/search-products) · Effect: **read**

```sh
sellapp products search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
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

Inspect the complete schema: `sellapp commands products search --json`.

## update

Update a product

Update a product in the authenticated store. Requires a token with the `listing` ability. Send `expected_updated_at` to reject stale writes instead of overwriting a newer change. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/update-a-product) · Effect: **write**

```sh
sellapp products update 120 --title 'Design kit' --description 'Templates for your next project.' --visibility HIDDEN --expected-updated-at 2026-08-30T12:00:00.000000Z
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --title | No | Value for title. |
| --description | No | Value for description. |
| --visibility | No | Value for visibility. |
| --slug | No | Value for slug. |
| --section | No | Value for section. |
| --additional-information | No | Value for additional information. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot. Send the exact `updated_at` value from the product you last retrieved; the mutation is rejected if the product has changed. |

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

Inspect the complete schema: `sellapp commands products update --json`.

