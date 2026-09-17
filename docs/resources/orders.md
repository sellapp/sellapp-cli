# orders

[All commands](../commands.md)

## create

Create an order

Create an order with one or more product-variant line items and start checkout. This can create provider payment sessions, reserve stock or wallet funds, and begin fulfillment for a free or fully wallet-funded order. Requires the invoice ability or orders:write OAuth scope and current store permission. Use the same Idempotency-Key and identical body after a lost response; inspect the order before creating a new key. The payment.checkout_url is a handoff, not proof of payment. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create --body '{"customer_email":"maya@example.com","payment_method":"STRIPE","product_variants":{"4321":{"quantity":1}}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer-email | Yes | Value for customer email. |
| --payment-method | Yes | Value for payment method. |
| --product-variants | Yes | Value for product variants. |
| --customer-ip | No | Value for customer ip. |
| --coupon | No | Value for coupon. |
| --vat-id | No | Value for vat id. |
| --country | No | Value for country. |
| --affiliate | No | Value for affiliate. |
| --extra | No | Value for extra. |
| --custom-payment-method-id | No | Required when payment_method is CUSTOM_PAYMENT_METHOD. ULID of an enabled custom payment method belonging to this store. |

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

Inspect the complete schema: `sellapp commands orders create --json`.

## create-checkout

Create order checkout

Start or reuse checkout for a PENDING order. A new nonfree checkout returns 201; an existing payment session or free order returns 200. Free or fully wallet-funded checkout can mark the order paid. Failures can mark the order FAILED and release wallet reservations. Do not infer payment from a checkout URL or browser return. The optional expected_status field is validated but is not a checkout concurrency precondition. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create-checkout 9001 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --expected-status | No | Validated status value; checkout does not compare it with current status. Retrieve the order before retrying. |

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

Inspect the complete schema: `sellapp commands orders create-checkout --json`.

## create-from-wallet

Create and pay an order from a wallet

Create an order and debit the customer's wallet in one atomic call. This spends real wallet funds and queues normal fulfillment; PAID means payment was captured, while COMPLETED depends on fulfillment. Requires both invoice and wallet:write API-key abilities and both store permissions. Wallets must be enabled, the customer must already exist in this store, and an active wallet must cover the entire positive total after discounts and taxes. Subscriptions and partial wallet payments are not supported. Omit payment_method and custom_payment_method_id. An insufficient balance or invalid cart returns 422 without creating an order or debit. Send the required Idempotency-Key; retry an identical request with the same key after a lost response. Reusing it with different input returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create-from-wallet --body '{"customer_email":"maya.chen@example.com","product_variants":{"4321":{"quantity":1}},"country":"US"}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer-email | Yes | Value for customer email. |
| --product-variants | Yes | Value for product variants. |
| --customer-ip | No | Value for customer ip. |
| --coupon | No | Value for coupon. |
| --vat-id | No | Value for vat id. |
| --country | No | Value for country. |
| --affiliate | No | Value for affiliate. |
| --extra | No | Value for extra. |

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

Inspect the complete schema: `sellapp commands orders create-from-wallet --json`.

## create-refund

Refund an order

Request a real refund for a completed, nonsubscription order with refundable balance and a supported provider. amount is a decimal string in the order currency: "5.00" means $5.00 in USD, not 500 cents. Omitting amount, sending null, or an empty string requests the FULL remaining balance. Do not send amount_cents: it is not a request field. Inspect refund.status (pending/effective/failed/cancelled); a 200 response does not guarantee money has reached the customer. Retrieve refund state before retrying an uncertain provider outcome. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create-refund 9001 --amount 5.00 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --amount | No | Decimal major-unit amount in the order currency, positive and no greater than remaining balance. Precision follows that currency. Omission/null/empty requests the full remaining balance. |

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

Inspect the complete schema: `sellapp commands orders create-refund --json`.

## create-replacement

Issue replacements

Create a replacement order and begin fulfillment for selected purchased variants. Supply a variant ID, an array of variant IDs, or a map of variant IDs to replacement quantities (null means original quantity). Variants must belong to this order; replacement quantities cannot exceed purchased quantity or available stock. Already replaced deliveries cannot be replaced again. The returned ID belongs to the new replacement order. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create-replacement 9001 --body '{"product_variants":[4321]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --product-variants | Yes | Value for product variants. |

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

Inspect the complete schema: `sellapp commands orders create-replacement --json`.

## get

Retrieve an order

Retrieve an order and its product-level line-item summaries from the same store. Requires the invoice token ability. An ID from another store returns not found. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/retrieve-an-order) · Effect: **read**

```sh
sellapp orders get 1042
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order path parameter. |

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

Inspect the complete schema: `sellapp commands orders get --json`.

## line-items get

Retrieve an order's line item

Retrieve one product-level sale record for the selected store. Requires the `invoice` token ability. Records from another store or whose order does not match the store return `404`. The order path parameter scopes the result and must match the line item's parent order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/order-line-items/retrieve-an-orders-line-item) · Effect: **read**

```sh
sellapp orders line-items get --order 4001 9001
```

| Input | Required | Meaning |
| --- | --- | --- |
| --order | Yes | The parent order ID. |
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

Inspect the complete schema: `sellapp commands orders line-items get --json`.

## line-items list

List an order's line items

List purchase line items for the selected store. Requires the `invoice` API-key ability or the corresponding OAuth scope and current store permission. The order path parameter scopes the result and must match the line item's parent order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/order-line-items/list-an-orders-line-items) · Effect: **read**

```sh
sellapp orders line-items list 4001
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The parent order ID. |
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

Inspect the complete schema: `sellapp commands orders line-items list --json`.

## line-items search

Search an order's line items

Search product-level sale records with filters and sorting. Requires the `invoice` token ability. The order path parameter scopes the result and must match the line item's parent order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/order-line-items/search-an-orders-line-items) · Effect: **read**

```sh
sellapp orders line-items search 4001
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The parent order ID. |
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

Inspect the complete schema: `sellapp commands orders line-items search --json`.

## list

List orders

List orders for the selected store. Requires the invoice token ability. Results include customer and payment-provider summaries without sensitive data, totals in integer minor units, status timelines, and product-level line-item summaries from the same store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/list-orders) · Effect: **read**

```sh
sellapp orders list
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

Inspect the complete schema: `sellapp commands orders list --json`.

## list-deliverables

List order deliverables

List seller-visible fulfillment records, expanding purchased bundle snapshots into their component variants. File download links expire after one hour. This payload includes line-item and variant IDs but does not expose the delivered_product_id required for dynamic delivery retries. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **read**

```sh
sellapp orders list-deliverables 9001
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |

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

Inspect the complete schema: `sellapp commands orders list-deliverables --json`.

## pay-from-wallet

Pay an existing order from its customer wallet

Debit the order's customer wallet for the complete stored order total and queue normal fulfillment. Requires both invoice and wallet:write API-key abilities and both store permissions. Only PENDING orders with pending lines, no subscription, and no existing payment method or checkout session are eligible. Active provider checkouts cannot be replaced. The store must enable wallets and the customer's active wallet must cover the full amount. The order ID determines the customer and amount; neither can be overridden. Send the required Idempotency-Key and reuse it with identical input after a lost response. A changed expected_status returns 409; an ineligible order or insufficient balance returns 422 without a debit. PAID does not guarantee fulfillment has completed. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders pay-from-wallet 42 --expected-status PENDING --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | Value for order. |
| --expected-status | No | Value for expected status. |

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

Inspect the complete schema: `sellapp commands orders pay-from-wallet --json`.

## retry-dynamic-delivery

Retry dynamic delivery

Queue another dynamic delivery attempt for one delivered product of a COMPLETED order. The delivered product, line item and currently dynamic variant must all belong to this order and store, with a configured webhook. 202 and meta.queued acknowledge the retry request, not successful external delivery. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders retry-dynamic-delivery 9001 --delivered-product-id 81 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --delivered-product-id | Yes | Delivered product record ID, not a catalog product or order line-item ID. The public order-deliverables response does not currently expose this ID; obtain it from an authorized integration context before requesting a retry. |

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

Inspect the complete schema: `sellapp commands orders retry-dynamic-delivery --json`.

## retry-fulfillment

Retry fulfillment

Reset fulfillment effects for a PAID or PARTIAL order and retry delivery. The optional email overrides the recipient. meta.effects_reset counts reset effects; it does not prove fulfillment finished. Other order states and a busy lifecycle lock return 422. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders retry-fulfillment 9001 --email 'maya@example.com' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --email | No | Optional delivery recipient override. Omit or send null to use the order delivery email. |

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

Inspect the complete schema: `sellapp commands orders retry-fulfillment --json`.

## search

Search orders

Find orders using the same filters as the list endpoint. Requires the invoice token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/search-orders) · Effect: **read**

```sh
sellapp orders search --body '{"filters":[{"field":"transaction_id","operator":"contains","value":"pi_3Example"},{"field":"status","operator":"in","value":["COMPLETED"]}],"sort":[{"field":"created_at","direction":"desc"}],"pagination":{"page":1,"limit":25}}'
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

Inspect the complete schema: `sellapp commands orders search --json`.

## send-fulfillment-notifications

Send fulfillment notifications

Send delivery notifications for each selected purchased variant. Omit product_variant_ids to select all variants. A supplied list must contain distinct positive IDs belonging to this order and store. Noncompleted orders or a missing recipient return meta.notifications_sent=0. The count records notifications requested, not email delivery confirmation. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders send-fulfillment-notifications 9001 --body '{"email":"maya@example.com","product_variant_ids":[4321]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --email | No | Optional delivery recipient override. Omit or send null to use the order delivery email. |
| --product-variant-ids | No | Value for product variant ids. |

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

Inspect the complete schema: `sellapp commands orders send-fulfillment-notifications --json`.

## update-status

Update order status

Change order lifecycle state. COMPLETED finalizes fulfillment of an already paid/partial/disputing order; it does not capture a pending payment. VOIDED can schedule provider side effects. REVIEW and DISPUTING have state restrictions. Send expected_status to prevent overwriting a changed state; an invalid transition, stale status, or busy lifecycle lock returns 422. Requires the invoice ability or orders:write OAuth scope and current store permission. Idempotency-Key is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders update-status 9001 --status COMPLETED --expected-status PAID --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --status | Yes | Value for status. |
| --expected-status | No | Expected current status. A status update checks this under the order lock and returns 422 if the order has changed. |

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

Inspect the complete schema: `sellapp commands orders update-status --json`.

