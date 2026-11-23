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

## duplicate

Duplicate a product

Copy a normal product, bundle, add-on, credits product, or booking product within the same store and return the copy. Courses cannot be duplicated. The copy is titled `(Copy)`, `(Copy 2)`, and so on, with a matching `-copy` slug, and keeps every other product setting, including its images, `visibility`, and draft state. A copy of a published product is therefore on sale straight away. It gets no place in the product order or in a section. A normal product's variants are copied with their settings and downloadable files, except subscription variants. A bundle's variant is copied with the products in it, an add-on's with its downloadable files and the products it is offered with, a credits product's with its rate tiers, and a booking product's options, except subscription options, with their schedules. Never copied: sales, reviews, appointments, serial codes, issued licence keys, and group membership. A copied variant that delivers serial codes gets a stock of `0`, so it shows as sold out until you add codes. Requires the `listing` API-key ability and the `listing` permission in your team role. Not safe to retry blindly: every call creates another copy. If a response is lost, list your products and look for the copy before trying again. A course, or a product with nothing to copy (only subscription variants), returns `422` on `product`. A product that changed while the copy was being made also returns `422` on `product`; nothing was copied. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/duplicate-a-product) · Effect: **write**

```sh
sellapp products duplicate 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |

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

Inspect the complete schema: `sellapp commands products duplicate --json`.

## featured-variant get

Retrieve product featured variant

Get the variant highlighted on the product page and its label, or nulls when no variant is featured. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/product-featured-variant) · Effect: **read**

```sh
sellapp products featured-variant get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |

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

Inspect the complete schema: `sellapp commands products featured-variant get --json`.

## featured-variant replace

Replace product featured variant

Highlight one of the product's published variants on its product page with a label, or remove the highlight by sending `variant_id` as null. Normal products, booking products, and credits products that have been published can feature a variant once they have at least two variants. Requires the `listing` API-key ability and the `listing` permission in your team role. Safe to repeat: the same body gives the same result. A course, bundle, or add-on, a draft that was never published, or a product with fewer than two variants returns `422` on `product`. A draft variant or a variant of another product returns `422` on `variant_id`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/product-featured-variant) · Effect: **write**

```sh
sellapp products featured-variant replace 1 --variant-id 2 --highlight-label best_value
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --variant-id | Yes | A published variant of this product, or null to stop featuring a variant. Always send it, even as null. |
| --highlight-label | No | The label shown on the featured variant. Defaults to `popular`. |

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

Inspect the complete schema: `sellapp commands products featured-variant replace --json`.

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
| --min-price | No | Only products whose every priced variant costs at least this much, in integer minor units of that variant's own currency: `999` means $9.99 for a USD variant. Draft variants count. A product without a priced variant never matches a price filter. Must not be greater than `max_price`. |
| --max-price | No | Only products whose every priced variant costs at most this much, in integer minor units of that variant's own currency: `2500` means $25.00 for a USD variant. Draft variants count. A product without a priced variant never matches a price filter. |
| --stock | No | Filter by stock, draft variants included. `in_stock`: at least one variant can be bought (unlimited stock, or at least its minimum purchase quantity left). `sold_out`: no variant can be bought. `any_variant_sold_out`: at least one variant with a stock count is sold out. `limited`: at least one variant has a stock count. `unlimited`: no variant has a stock count. |
| --visibility | No | Only products with one of these visibilities. Repeat the parameter to select several. A single `visibility=HIDDEN` also works. |
| --has-sales | No | `true` for products with at least one completed sale, `false` for products without one. Only purchase line items with status `COMPLETED` count, so pending or refunded ones do not. |

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

## move

Move a product

Move a normal product to a position in your store's product order, shifting the products in between by one. Positions run from `1` to the number of normal products, drafts included; a product's position is its `order`. Gaps in the stored order are closed first. Requires the `listing` API-key ability and the `listing` permission in your team role. Safe to repeat: moving a product to the position it already has changes nothing. Send `expected_position` with the position you last read to refuse the move when someone else moved the product in the meantime. A position past the end returns `422` on `position` with the allowed range, a product that is no longer at `expected_position` returns `422` on `expected_position` with its current position, and a listing that is not a normal product returns `422` on `product`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/move-a-product) · Effect: **write**

```sh
sellapp products move 1 --position 1 --expected-position 3
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --position | Yes | The new position, where `1` is first. At most the number of normal products in the store, drafts included. |
| --expected-position | No | The position you last read for this product (its `order`). The move is refused when the product is no longer there. |

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

Inspect the complete schema: `sellapp commands products move --json`.

## publish

Publish a product

Publish a draft product so customers can buy it. Works for normal products, bundles, add-ons, credits products, booking products, and courses. Each is checked as saved, the way its dashboard editor checks it before publishing, and nothing else about the product changes. A normal product needs a description and at least one published variant, and every variant that will be on sale needs a title, a description of at least 5 characters, a working delivery (for example an uploaded file for a download), a price, and payment methods enabled for the store that can take that currency. With exactly one variant, that variant is published with the product. With several, send the draft variants to publish with it in `variant_ids`; the others stay drafts. Publishing a subscription variant also creates its billing plans with the payment providers. A booking product publishes its one option with it; with several options, publish at least one with `POST /v2/products/{product}/variants/{variant}/publish` first, and the other options keep their own state. A bundle needs at least one product in it. A course is published with `visibility`, or with its saved visibility when you leave it out; `PUBLIC` also needs a ready promo video and at least one published lesson. A subscription-priced course cannot be published here. Publishing does not change a normal product's `visibility`: a product that is not `PUBLIC` is published but stays off the storefront listing. Requires the `listing` API-key ability and the `listing` permission in your team role. Safe to repeat: once the product is published, another call changes nothing and returns `422` with `The product is already published.` on `product`. Every refusal returns `422`. `visibility` is accepted only for a course and `variant_ids` only for a normal product; each is otherwise refused on that field, and `variant_ids.{i}` names an ID that is not a variant of this product. Everything else is on `product`, with one sentence per thing the product still needs. One `product` refusal comes after the product was published: when a payment provider refuses to create a subscription billing plan, the product and its variants are already published and the message is the provider's reason. Retrieve the product before trying again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/publish-a-product) · Effect: **consequential**

```sh
sellapp products publish 1 --body '{"variant_ids":[2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --visibility | No | Value for visibility. |
| --variant-ids | No | Normal products only. IDs of this product's draft variants to publish with it. Needed when the product has several variants and none is published yet; a single variant is published anyway. Draft variants you leave out stay drafts. |

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

Inspect the complete schema: `sellapp commands products publish --json`.

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

Search your store's products by title and narrow them by price, stock, visibility, or sales, 15 per page by default. Send the filters in the JSON body. Products support no `filters` or `sort` entries: any entry returns `422`. A filter value that is not valid, such as an unknown `stock` state or a `min_price` above `max_price`, returns `422` on that field. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/search-products) · Effect: **read**

```sh
sellapp products search --body '{"search":{"value":"launch"},"max_price":2500,"stock":"in_stock","visibility":["PUBLIC"],"has_sales":true}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --filters | No | Not supported for products: no field can be filtered this way, so any entry returns `422` on `filters`. Use the price, stock, visibility, and sales fields instead. |
| --sort | No | Not supported for products: any entry returns `422` on `sort`. |
| --search | No | Value for search. |
| --includes | No | Only `variants` is accepted, and variants are always included. |
| --min-price | No | Only products whose every priced variant costs at least this much, in integer minor units of that variant's own currency: `999` means $9.99 for a USD variant. Draft variants count. A product without a priced variant never matches a price filter. Must not be greater than `max_price`. |
| --max-price | No | Only products whose every priced variant costs at most this much, in integer minor units of that variant's own currency: `2500` means $25.00 for a USD variant. Draft variants count. A product without a priced variant never matches a price filter. |
| --stock | No | Filter by stock, draft variants included. `in_stock`: at least one variant can be bought (unlimited stock, or at least its minimum purchase quantity left). `sold_out`: no variant can be bought. `any_variant_sold_out`: at least one variant with a stock count is sold out. `limited`: at least one variant has a stock count. `unlimited`: no variant has a stock count. |
| --visibility | No | Only products with one of these visibilities. A single string such as `"HIDDEN"` is also accepted. |
| --has-sales | No | `true` for products with at least one completed sale, `false` for products without one. Only purchase line items with status `COMPLETED` count, so pending or refunded ones do not. |
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

Inspect the complete schema: `sellapp commands products search --json`.

## similar-products list

List product similar products

List the variants of other products recommended on this product's page, in the order shown. Deleted products are left out. The list is not paginated. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/product-similar-products) · Effect: **read**

```sh
sellapp products similar-products list 30
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |

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

Inspect the complete schema: `sellapp commands products similar-products list --json`.

## similar-products replace

Replace product similar products

Replace the variants recommended on this product's page with the list you send, in that order. Send an empty list to remove them all. A variant you add must be one the product editor offers: an in-stock, one-time-priced variant of another `PUBLIC` normal product in this product's currency. A variant already on the list may stay even if it no longer qualifies. Requires the `listing` API-key ability and the `listing` permission in your team role. Safe to repeat: the same list gives the same result. An add-on, a draft that was never published, or a product with a subscription-priced variant returns `422` on `product`. A variant that does not qualify, is named twice, or is not in your store returns `422` on `variant_ids.{i}`, where `{i}` is its position in your list. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products/product-similar-products) · Effect: **write**

```sh
sellapp products similar-products replace 30 --body '{"variant_ids":[31,41]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --variant-ids | Yes | Variant IDs to recommend, first to last. At most 100. Always send it, even as an empty list. |

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

Inspect the complete schema: `sellapp commands products similar-products replace --json`.

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

