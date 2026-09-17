# add-ons

[All commands](../commands.md)

## create

Create an add-on draft

Create an add-on, optionally including its variant and parent_product_ids in one transaction. Set is_draft false with variant to publish immediately. Without variant, new add-ons remain drafts. Requires the listing API-key ability and store listing permission. Invalid variants or parent products roll back the complete creation. This call does not have response-replay idempotency; check for the created add-on before retrying a request whose response was lost. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/create-an-add-on-draft) · Effect: **consequential**

```sh
sellapp add-ons create --body '{"title":"Customer support","description":"Priority support for launches scheduled suspiciously close to Friday.","visibility":"PUBLIC","parent_product_ids":[120,121]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Seller-facing title for the add-on. |
| --description | Yes | Description shown when the add-on is offered. |
| --visibility | Yes | Value for visibility. |
| --slug | No | Optional store-unique slug. When omitted during creation, SellApp generates one from the title. |
| --is-draft | No | Defaults to true on creation. Set false to publish. Publication requires exactly one published, fixed-price, single-payment variant; creation can include that variant in the same request. |
| --parent-product-ids | No | Ordered IDs of same-store, non-subscription products, courses, or bookings to assign. Supplying the field replaces the complete assignment set. |
| --variant | No | Value for variant. |

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

Inspect the complete schema: `sellapp commands add-ons create --json`.

## delete

Delete an add-on

Soft-delete an add-on and remove its product assignments. The deleted resource is returned for audit and reconciliation workflows. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/delete-an-add-on) · Effect: **consequential**

```sh
sellapp add-ons delete 410 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| addon (positional) | Yes | The addon path parameter. |

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

Inspect the complete schema: `sellapp commands add-ons delete --json`.

## get

Retrieve an add-on

Retrieve one add-on owned by the authenticated store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/retrieve-an-add-on) · Effect: **read**

```sh
sellapp add-ons get 410
```

| Input | Required | Meaning |
| --- | --- | --- |
| addon (positional) | Yes | The addon path parameter. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --with-drafts | No | Include draft add-ons alongside published add-ons. |
| --only-drafts | No | Return only draft add-ons. |

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

Inspect the complete schema: `sellapp commands add-ons get --json`.

## list

List add-ons

Retrieve the store's add-ons with compact variant references and ordered parent-product assignments. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/list-add-ons) · Effect: **read**

```sh
sellapp add-ons list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --with-drafts | No | Include draft add-ons alongside published add-ons. |
| --only-drafts | No | Return only draft add-ons. |

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

Inspect the complete schema: `sellapp commands add-ons list --json`.

## parent-products list

List an add-on's parent products

Retrieve the ordered products, courses, and bookings that currently offer this add-on. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/list-parent-products) · Effect: **read**

```sh
sellapp add-ons parent-products list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| addon (positional) | Yes | The addon path parameter. |
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

Inspect the complete schema: `sellapp commands add-ons parent-products list --json`.

## parent-products replace

Replace an add-on's parent products

Replace the complete ordered parent-product assignment. IDs must belong to the same store and may reference non-subscription products, courses, or bookings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/replace-parent-products) · Effect: **consequential**

```sh
sellapp add-ons parent-products replace 410 --body '{"resources":[121,120]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| addon (positional) | Yes | The addon path parameter. |
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

Inspect the complete schema: `sellapp commands add-ons parent-products replace --json`.

## replace

Update an add-on

Update an add-on and optionally replace its ordered parent-product assignments. Publication is rejected unless the add-on has exactly one published, fixed-price, single-payment variant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/replace-an-add-on) · Effect: **consequential**

```sh
sellapp add-ons replace 410 --description 'Priority email, chat, and launch-day reassurance.' --is-draft false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| addon (positional) | Yes | The addon path parameter. |
| --title | No | Seller-facing title for the add-on. |
| --slug | No | Optional store-unique slug. When omitted during creation, SellApp generates one from the title. |
| --description | No | Description shown when the add-on is offered. |
| --visibility | No | Value for visibility. |
| --is-draft | No | Defaults to true on creation. Set false to publish. Publication requires exactly one published, fixed-price, single-payment variant; creation can include that variant in the same request. |
| --parent-product-ids | No | Ordered IDs of same-store, non-subscription products, courses, or bookings to assign. Supplying the field replaces the complete assignment set. |

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

Inspect the complete schema: `sellapp commands add-ons replace --json`.

## search

Search add-ons

Search add-ons by title, slug, or description and optionally filter or sort the result set. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/search-add-ons) · Effect: **read**

```sh
sellapp add-ons search --body '{"filters":[{"field":"id","operator":"=","value":410}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --with-drafts | No | Include draft add-ons alongside published add-ons. |
| --only-drafts | No | Return only draft add-ons. |
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

Inspect the complete schema: `sellapp commands add-ons search --json`.

## update

Update an add-on

Update an add-on and optionally replace its ordered parent-product assignments. Publication is rejected unless the add-on has exactly one published, fixed-price, single-payment variant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/add-ons/update-an-add-on) · Effect: **consequential**

```sh
sellapp add-ons update 410 --description 'Priority email, chat, and launch-day reassurance.' --is-draft false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| addon (positional) | Yes | The addon path parameter. |
| --title | No | Seller-facing title for the add-on. |
| --slug | No | Optional store-unique slug. When omitted during creation, SellApp generates one from the title. |
| --description | No | Description shown when the add-on is offered. |
| --visibility | No | Value for visibility. |
| --is-draft | No | Defaults to true on creation. Set false to publish. Publication requires exactly one published, fixed-price, single-payment variant; creation can include that variant in the same request. |
| --parent-product-ids | No | Ordered IDs of same-store, non-subscription products, courses, or bookings to assign. Supplying the field replaces the complete assignment set. |

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

Inspect the complete schema: `sellapp commands add-ons update --json`.

