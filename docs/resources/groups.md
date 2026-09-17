# groups

[All commands](../commands.md)

## create

Create a group

Create a group for related products. Supply product_ids and section_id to save the group, ordered product membership, and section together. The whole change rolls back if any selected product is invalid. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/create-a-group) · Effect: **write**

```sh
sellapp groups create --title 'Design kit' --unlisted true
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Value for title. |
| --unlisted | Yes | Value for unlisted. |
| --order | No | Value for order. |
| --product-ids | No | Ordered IDs of published standard products in this store. Supply an empty array to remove every product. Missing, draft, or foreign products are rejected. |
| --section-id | No | Section in this store. Omit to preserve it; null removes the section. |

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

Inspect the complete schema: `sellapp commands groups create --json`.

## delete

Delete a group

Deletes a group. This will permanently delete the group and its details. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/delete-a-group) · Effect: **consequential**

```sh
sellapp groups delete 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |

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

Inspect the complete schema: `sellapp commands groups delete --json`.

## get

Retrieve a group

Retrieve a group by its ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/retrieve-a-group) · Effect: **read**

```sh
sellapp groups get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |

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

Inspect the complete schema: `sellapp commands groups get --json`.

## list

List all groups

List your store's groups, 15 per page by default. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/list-all-groups) · Effect: **read**

```sh
sellapp groups list
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

Inspect the complete schema: `sellapp commands groups list --json`.

## products add

Add products to group

Add existing products to a group. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/add-products-to-group) · Effect: **consequential**

```sh
sellapp groups products add 1 --body '{"resources":[1]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |
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

Inspect the complete schema: `sellapp commands groups products add --json`.

## products get

List specific product within group

Retrieve one product from a group. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/list-specific-product-within-group) · Effect: **read**

```sh
sellapp groups products get --group 1 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --group | Yes | The group path parameter. |
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

Inspect the complete schema: `sellapp commands groups products get --json`.

## products list

List all products within group

List the products in a group, 15 per page by default. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/list-all-products-within-group) · Effect: **read**

```sh
sellapp groups products list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |
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

Inspect the complete schema: `sellapp commands groups products list --json`.

## products remove

Remove products from group

Remove products from a group without deleting the products. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/remove-products-from-group) · Effect: **consequential**

```sh
sellapp groups products remove 1 --body '{"resources":[1]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |
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

Inspect the complete schema: `sellapp commands groups products remove --json`.

## products replace

Replace ordered group products

Replace a group's complete product membership in one transaction. Requires the group API-key ability and store group permission. The order of product_ids becomes display order. An empty array clears the group. Missing or draft products return 422; products from another store are rejected. Retrying the same list has the same result. Private products remain assigned but are omitted from the returned group summary. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/list-all-products-within-group) · Effect: **consequential**

```sh
sellapp groups products replace 42 --body '{"product_ids":[120,121]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | Value for group. |
| --product-ids | Yes | Ordered IDs of published standard products in this store. Supply an empty array to remove every product. Missing, draft, or foreign products are rejected. |

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

Inspect the complete schema: `sellapp commands groups products replace --json`.

## products search

Search products within group

Search the products in a group using JSON body filters, search terms, includes, and sort instructions. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/search-products-within-group) · Effect: **read**

```sh
sellapp groups products search 1 --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |
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

Inspect the complete schema: `sellapp commands groups products search --json`.

## search

Search groups

Search groups using JSON body filters, search terms, includes, and sort instructions. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/search-groups) · Effect: **read**

```sh
sellapp groups search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
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

Inspect the complete schema: `sellapp commands groups search --json`.

## update

Update a group

Update a group's details. Supply product_ids and section_id to save the group, ordered product membership, and section together. The whole change rolls back if any selected product is invalid. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/groups/update-a-group) · Effect: **consequential**

```sh
sellapp groups update 1 --title 'Founder reading room' --unlisted true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| group (positional) | Yes | The group path parameter. |
| --title | No | Value for title. |
| --unlisted | No | Value for unlisted. |
| --order | No | Value for order. |
| --product-ids | No | Ordered IDs of published standard products in this store. Supply an empty array to remove every product. Missing, draft, or foreign products are rejected. |
| --section-id | No | Section in this store. Omit to preserve it; null removes the section. |

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

Inspect the complete schema: `sellapp commands groups update --json`.

