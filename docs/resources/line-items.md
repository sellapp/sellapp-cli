# line-items

[All commands](../commands.md)

## get

Retrieve an order line item

Retrieve one product-level sale record for the selected store. Requires the `invoice` token ability. Records from another store or whose order does not match the store return `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/order-line-items/retrieve-an-order-line-item) · Effect: **read**

```sh
sellapp line-items get 9001
```

| Input | Required | Meaning |
| --- | --- | --- |
| lineItem (positional) | Yes | The purchase line-item ID. |

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

Inspect the complete schema: `sellapp commands line-items get --json`.

## list

List order line items

List purchase line items for the selected store. Requires the `invoice` API-key ability or the corresponding OAuth scope and current store permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/order-line-items/list-order-line-items) · Effect: **read**

```sh
sellapp line-items list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Maximum number of line items per page. |
| --page | No | One-based result page. |

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

Inspect the complete schema: `sellapp commands line-items list --json`.

## search

Search order line items

Search product-level sale records with filters and sorting. Requires the `invoice` token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/order-line-items/search-order-line-items) · Effect: **read**

```sh
sellapp line-items search
```

| Input | Required | Meaning |
| --- | --- | --- |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --pagination | No | Value for pagination. |

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

Inspect the complete schema: `sellapp commands line-items search --json`.

