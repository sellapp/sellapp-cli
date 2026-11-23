# refunds

[All commands](../commands.md)

## get-refund

Retrieve a refund

Retrieve one store-scoped refund. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/refunds) · Effect: **read**

```sh
sellapp refunds get-refund 781
```

| Input | Required | Meaning |
| --- | --- | --- |
| refund (positional) | Yes | The refund identifier. |

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

Inspect the complete schema: `sellapp commands refunds get-refund --json`.

## list-refunds

List refunds

List refunds for orders and standalone charges in the selected store, newest first, 15 per page by default. Filter by `status`, `provider`, `source`, or the order or charge refunded; filters combine. To see every refund of one order, send `refundable_type=order` and `refundable_id` set to the order ID. A value outside a parameter's accepted values, a `limit` above 100, or a query parameter the endpoint does not take returns `422` on that parameter. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/refunds) · Effect: **read**

```sh
sellapp refunds list-refunds
```

| Input | Required | Meaning |
| --- | --- | --- |
| --page | No | Page number, starting at 1. |
| --limit | No | Refunds per page, from 1 to 100. |
| --status | No | Keep refunds in this state: `pending`, `effective`, `failed`, or `cancelled`. |
| --provider | No | Keep refunds made through this payment provider: `manual`, `mercado_pago`, `nmi`, `mollie`, `paypal`, `razorpay`, `square`, or `stripe`. |
| --source | No | Keep refunds by how they were recorded: `manual`, `provider_webhook`, `provider_api`, or `system`. For example, `provider_api` is a refund SellApp requested from the provider and `provider_webhook` one the provider reported. |
| --refundable-type | No | Keep refunds of orders (`order`) or of standalone charges (`charge`). |
| --refundable-id | No | Keep refunds of the order or charge with this ID. Send `refundable_type` with it, because an order and a charge can share an ID. |

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

Inspect the complete schema: `sellapp commands refunds list-refunds --json`.

