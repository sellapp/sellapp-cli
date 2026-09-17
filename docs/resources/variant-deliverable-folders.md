# variant-deliverable-folders

[All commands](../commands.md)

## create

Create a variant deliverable folder

Create a variant deliverable folder. Requires the listing Sanctum ability. Parent folders are validated against the same variant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverable-folders create --product test_product 1 --name 'Design kit' --description 'Files included with your purchase.'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| variant (positional) | Yes | The product variant identifier. |
| --name | Yes | Value for name. |
| --description | No | Value for description. |
| --parent-id | No | Value for parent id. |
| --sort-order | No | Value for sort order. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-folders create --json`.

## delete

Delete a variant deliverable folder

Delete a variant deliverable folder. Requires the listing Sanctum ability. Folder and parent IDs are scoped through the authenticated store, product, and variant; cycles are rejected. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **consequential**

```sh
sellapp variant-deliverable-folders delete --product test_product --variant 1 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| folder (positional) | Yes | The deliverable folder identifier. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-folders delete --json`.

## get

Retrieve a variant deliverable folder

Retrieve a variant deliverable folder. Requires the listing Sanctum ability. Folder and parent IDs are scoped through the authenticated store, product, and variant; cycles are rejected. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **read**

```sh
sellapp variant-deliverable-folders get --product test_product --variant 1 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| folder (positional) | Yes | The deliverable folder identifier. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-folders get --json`.

## list

List variant deliverable folders

List variant deliverable folders. Requires the listing Sanctum ability. Returns folders ordered by sort order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **read**

```sh
sellapp variant-deliverable-folders list --product test_product 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| variant (positional) | Yes | The product variant identifier. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-folders list --json`.

## replace

Replace variant deliverable folder settings

Replace variant deliverable folder settings. Requires the listing Sanctum ability. Folder and parent IDs are scoped through the authenticated store, product, and variant; cycles are rejected. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverable-folders replace --product test_product --variant 1 1 --name 'Design kit'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| folder (positional) | Yes | The deliverable folder identifier. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --parent-id | No | Value for parent id. |
| --sort-order | No | Value for sort order. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-folders replace --json`.

## update

Update a variant deliverable folder

Update a variant deliverable folder. Requires the listing Sanctum ability. Folder and parent IDs are scoped through the authenticated store, product, and variant; cycles are rejected. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverable-folders update --product test_product --variant 1 1 --name 'Design kit'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| folder (positional) | Yes | The deliverable folder identifier. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --parent-id | No | Value for parent id. |
| --sort-order | No | Value for sort order. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-folders update --json`.

