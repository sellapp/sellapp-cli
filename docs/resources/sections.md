# sections

[All commands](../commands.md)

## batch-create

Batch create sections

Create multiple sections in one request by sending a `resources` array.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections batch-create --body '{"resources":[{"title":"Featured","hidden":false}]}' --yes
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

Inspect the complete schema: `sellapp commands sections batch-create --json`.

## batch-delete

Batch delete sections

Delete multiple sections in one request by sending the section IDs in `resources`.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections batch-delete --body '{"resources":[1,2]}' --yes
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

Inspect the complete schema: `sellapp commands sections batch-delete --json`.

## batch-update

Batch update sections

Update multiple sections in one request by sending a `resources` object keyed by section ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections batch-update --body '{"resources":{"1":{"title":"Featured","hidden":false}}}' --yes
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

Inspect the complete schema: `sellapp commands sections batch-update --json`.

## create

Create a section

Create a section to organize your storefront.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections create --title 'Founder resources' --hidden false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Value for title. |
| --hidden | Yes | Value for hidden. |

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

Inspect the complete schema: `sellapp commands sections create --json`.

## delete

Delete a section

Deletes a section. This will permanently delete the section and its details.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections delete 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |

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

Inspect the complete schema: `sellapp commands sections delete --json`.

## get

Retrieve a section

Retrieve a section by its ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp sections get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |

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

Inspect the complete schema: `sellapp commands sections get --json`.

## list

List all sections

List your storefront sections, 15 per page by default.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp sections list
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands sections list --json`.

## replace

Update a section

Update a storefront section's details.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections replace 1 --title 'Founder resources' --hidden false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
| --title | No | Value for title. |
| --hidden | No | Value for hidden. |

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

Inspect the complete schema: `sellapp commands sections replace --json`.

## replace-groups

Replace section groups

Replace the groups assigned to a section. Group IDs are stored in the supplied order; omitted groups are detached and unknown IDs are rejected.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections replace-groups 1 --body '{"resources":[3,1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
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

Inspect the complete schema: `sellapp commands sections replace-groups --json`.

## replace-order

Replace section order

Replace the storefront section order with every section ID in the store. IDs are applied in array order; incomplete lists and unknown IDs are rejected.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections replace-order --body '{"resources":[3,1,2]}' --yes
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

Inspect the complete schema: `sellapp commands sections replace-order --json`.

## replace-products

Replace section products

Replace the products assigned to a section. Product IDs are stored in the supplied order; omitted products are detached and unknown IDs are rejected.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections replace-products 1 --body '{"resources":[3,1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
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

Inspect the complete schema: `sellapp commands sections replace-products --json`.

## search

Search sections

Search sections using JSON body filters, search terms, includes, and sort instructions.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp sections search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands sections search --json`.

## update

Update a section

Update a storefront section's details.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp sections update 1 --title 'Founder resources' --hidden false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
| --title | No | Value for title. |
| --hidden | No | Value for hidden. |

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

Inspect the complete schema: `sellapp commands sections update --json`.

## v-2-batch-create-sections

Batch create sections

Create multiple sections in one request by sending a `resources` array. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/batch-create-sections) · Effect: **consequential**

```sh
sellapp sections v-2-batch-create-sections --body '{"resources":[{"title":"Featured","hidden":false}]}' --yes
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

Inspect the complete schema: `sellapp commands sections v-2-batch-create-sections --json`.

## v-2-batch-delete-sections

Batch delete sections

Delete multiple sections in one request by sending the section IDs in `resources`. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/batch-delete-sections) · Effect: **consequential**

```sh
sellapp sections v-2-batch-delete-sections --body '{"resources":[1,2]}' --yes
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

Inspect the complete schema: `sellapp commands sections v-2-batch-delete-sections --json`.

## v-2-batch-update-sections

Batch update sections

Update multiple sections in one request by sending a `resources` object keyed by section ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/batch-update-sections) · Effect: **consequential**

```sh
sellapp sections v-2-batch-update-sections --body '{"resources":{"1":{"title":"Featured","hidden":false}}}' --yes
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

Inspect the complete schema: `sellapp commands sections v-2-batch-update-sections --json`.

## v-2-create-section

Create a section

Create a section to organize your storefront. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/create-a-section) · Effect: **write**

```sh
sellapp sections v-2-create-section --title 'Founder resources' --hidden false
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Value for title. |
| --hidden | Yes | Value for hidden. |

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

Inspect the complete schema: `sellapp commands sections v-2-create-section --json`.

## v-2-delete-section

Delete a section

Deletes a section. This will permanently delete the section and its details. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/delete-a-section) · Effect: **consequential**

```sh
sellapp sections v-2-delete-section 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |

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

Inspect the complete schema: `sellapp commands sections v-2-delete-section --json`.

## v-2-get-section

Retrieve a section

Retrieve a section by its ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/retrieve-a-section) · Effect: **read**

```sh
sellapp sections v-2-get-section 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |

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

Inspect the complete schema: `sellapp commands sections v-2-get-section --json`.

## v-2-list-sections

List all sections

List your storefront sections, 15 per page by default. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/list-all-sections) · Effect: **read**

```sh
sellapp sections v-2-list-sections
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

Inspect the complete schema: `sellapp commands sections v-2-list-sections --json`.

## v-2-replace-section

Update a section

Update a storefront section's details. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/update-a-section) · Effect: **write**

```sh
sellapp sections v-2-replace-section 1 --title 'Founder resources' --hidden false
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
| --title | No | Value for title. |
| --hidden | No | Value for hidden. |

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

Inspect the complete schema: `sellapp commands sections v-2-replace-section --json`.

## v-2-replace-section-groups

Replace section groups

Replace the groups assigned to a section. Group IDs are stored in the supplied order; omitted groups are detached and unknown IDs are rejected. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/replace-section-groups) · Effect: **consequential**

```sh
sellapp sections v-2-replace-section-groups 1 --body '{"resources":[3,1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
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

Inspect the complete schema: `sellapp commands sections v-2-replace-section-groups --json`.

## v-2-replace-section-order

Replace section order

Replace the storefront section order with every section ID in the store. IDs are applied in array order; incomplete lists and unknown IDs are rejected. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/order-sections) · Effect: **consequential**

```sh
sellapp sections v-2-replace-section-order --body '{"resources":[3,1,2]}' --yes
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

Inspect the complete schema: `sellapp commands sections v-2-replace-section-order --json`.

## v-2-replace-section-products

Replace section products

Replace the products assigned to a section. Product IDs are stored in the supplied order; omitted products are detached and unknown IDs are rejected. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/replace-section-products) · Effect: **consequential**

```sh
sellapp sections v-2-replace-section-products 1 --body '{"resources":[3,1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
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

Inspect the complete schema: `sellapp commands sections v-2-replace-section-products --json`.

## v-2-search-sections

Search sections

Search sections using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/search-sections) · Effect: **read**

```sh
sellapp sections v-2-search-sections --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
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

Inspect the complete schema: `sellapp commands sections v-2-search-sections --json`.

## v-2-update-section

Update a section

Update a storefront section's details. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/sections/update-a-section) · Effect: **write**

```sh
sellapp sections v-2-update-section 1 --title 'Founder resources' --hidden false
```

| Input | Required | Meaning |
| --- | --- | --- |
| section (positional) | Yes | The section path parameter. |
| --title | No | Value for title. |
| --hidden | No | Value for hidden. |

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

Inspect the complete schema: `sellapp commands sections v-2-update-section --json`.

