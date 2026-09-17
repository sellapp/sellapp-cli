# variant-serials

[All commands](../commands.md)

## append

Append variant serial inventory

Atomically appends unsold serials and recalculates variant stock. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **consequential**

```sh
sellapp variant-serials append --product 1 1 --body '{"serials":["LICENSE-KEY-001","LICENSE-KEY-002"],"remove_duplicates":true}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --serials | Yes | Serial values to append or use as the replacement inventory. |
| --remove-duplicates | No | Remove duplicate values from this submitted payload. |

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

Inspect the complete schema: `sellapp commands variant-serials append --json`.

## delete

Delete a variant serial

Deletes one unsold serial scoped to the current store, product, and variant, then recalculates stock. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **consequential**

```sh
sellapp variant-serials delete --product 1 --variant 1 test_serial --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| --variant | Yes | The variant path parameter. |
| serial (positional) | Yes | The unsold serial entry ULID. |

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

Inspect the complete schema: `sellapp commands variant-serials delete --json`.

## list

List variant serial inventory

Lists unsold serial inventory for a variant in the current store. Requires the `listing` Sanctum ability. Serial metadata is not exposed. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **read**

```sh
sellapp variant-serials list --product 1 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --search | No | Filter inventory by a partial serial value. |
| --limit | No | Number of inventory entries per page. |
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

Inspect the complete schema: `sellapp commands variant-serials list --json`.

## queue

Queue a variant serial import

Uploads a text file to application-owned temporary storage and queues an atomic serial import. Use this endpoint for large imports. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **consequential**

```sh
sellapp variant-serials queue --product 1 1 --upload 'file=example-file' --parsing-mode COMMA --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --upload file=PATH | Yes | A text file up to 100 MiB. |
| --parsing-mode | Yes | How the uploaded file is split into serials. CUSTOM requires custom_delimiter; sending CUSTOM without it fails validation with a 422. |
| --custom-delimiter | No | The delimiter to split on. Required when parsing_mode is CUSTOM (omitting it fails validation with a 422); ignored otherwise. |
| --remove-duplicates | No | Value for remove duplicates. |
| --mode | No | Value for mode. |

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

Inspect the complete schema: `sellapp commands variant-serials queue --json`.

## replace

Replace variant serial inventory

Atomically replaces all unsold serials and recalculates variant stock. Concurrent inventory mutations are serialized. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **consequential**

```sh
sellapp variant-serials replace --product 1 1 --body '{"serials":["LICENSE-KEY-001","LICENSE-KEY-002"],"remove_duplicates":true}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --serials | Yes | Serial values to append or use as the replacement inventory. |
| --remove-duplicates | No | Remove duplicate values from this submitted payload. |

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

Inspect the complete schema: `sellapp commands variant-serials replace --json`.

