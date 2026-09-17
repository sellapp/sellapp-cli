# variant-deliverable-files

[All commands](../commands.md)

## delete

Delete a variant deliverable file

Delete a variant deliverable file. Requires the listing Sanctum ability. The file is scoped through the authenticated store, product, and variant. Internal storage keys are immutable and never returned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **consequential**

```sh
sellapp variant-deliverable-files delete --product test_product --variant 1 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| file (positional) | Yes | The managed deliverable file identifier. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-files delete --json`.

## get

Retrieve a variant deliverable file

Retrieve a variant deliverable file. Requires the listing Sanctum ability. The file is scoped through the authenticated store, product, and variant. Internal storage keys are immutable and never returned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **read**

```sh
sellapp variant-deliverable-files get --product test_product --variant 1 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| file (positional) | Yes | The managed deliverable file identifier. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-files get --json`.

## list

List variant deliverable files

List variant deliverable files. Requires the listing Sanctum ability. Lists managed files without revealing storage keys. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **read**

```sh
sellapp variant-deliverable-files list --product test_product 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| variant (positional) | Yes | The product variant identifier. |
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

Inspect the complete schema: `sellapp commands variant-deliverable-files list --json`.

## replace

Replace variant deliverable file settings

Replace variant deliverable file settings. Requires the listing Sanctum ability. The file is scoped through the authenticated store, product, and variant. Internal storage keys are immutable and never returned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverable-files replace --product test_product --variant 1 1 --custom-name 'Design kit.zip'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| file (positional) | Yes | The managed deliverable file identifier. |
| --custom-name | No | Value for custom name. |
| --folder-id | No | Value for folder id. |
| --watermark | No | Value for watermark. |
| --max-downloads | No | Value for max downloads. |
| --limit-to-purchase-ip | No | Value for limit to purchase ip. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-files replace --json`.

## update

Update a variant deliverable file

Update a variant deliverable file. Requires the listing Sanctum ability. The file is scoped through the authenticated store, product, and variant. Internal storage keys are immutable and never returned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverable-files update --product test_product --variant 1 1 --custom-name 'Design kit.zip'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| --variant | Yes | The product variant identifier. |
| file (positional) | Yes | The managed deliverable file identifier. |
| --custom-name | No | Value for custom name. |
| --folder-id | No | Value for folder id. |
| --watermark | No | Value for watermark. |
| --max-downloads | No | Value for max downloads. |
| --limit-to-purchase-ip | No | Value for limit to purchase ip. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-files update --json`.

## upload

Upload a variant deliverable file

Upload a variant deliverable file. Requires the listing Sanctum ability. Accepts an uploaded file, stores it under a server-generated key, and queues the durable move after commit. Uploads may be staged before DOWNLOADABLE is activated; inactive pending files are retained temporarily and scheduled for cleanup if activation is abandoned. Client-supplied storage paths are not accepted. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverable-files upload --product test_product 1 --upload 'file=example-file'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| variant (positional) | Yes | The product variant identifier. |
| --upload file=PATH | Yes | Value for file. |
| --folder-id | No | Value for folder id. |

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

Inspect the complete schema: `sellapp commands variant-deliverable-files upload --json`.

