# orders

[All commands](../commands.md)

## create

Create an order

Create an order with one or more product-variant line items and start checkout. This can create provider payment sessions, reserve stock or wallet funds, and begin fulfillment for a free or fully wallet-funded order. Requires the invoice ability or orders:write OAuth scope and current store permission. Use the same Idempotency-Key and identical body after a lost response; inspect the order before creating a new key. The payment.checkout_url is a handoff, not proof of payment. A subscription plan that offers a trial is priced at the trial, nothing for a free trial and the trial price for a paid one, unless the customer already had a trial of that product, matched by `customer_email` or a recent `customer_ip`; then it is priced at the plan price. Two requests racing for the same trial can be refused with a `422` on `invoice`. A cart carries at most one subscription plan and nothing else; a plan mixed with other products or with a second plan returns `422` on `product_variants.{variant_id}`, and a plan's `quantity` other than 1 returns `422` on `quantity`. A customer who already holds the plan, matched by `customer_email`, is refused with `422` on `invoice` while that subscription is active, trialing, past due, or unpaid, whichever payment method bills it. A subscription scheduled to end still blocks while it can be kept, and the message then tells the customer to keep it from the customer portal. One that can no longer be kept, such as a PayPal subscription cancelled at period end, does not block. For an in-house subscription on hold for an unpaid renewal, the message tells the customer to pay and resume it from the customer portal. A `coupon` discounts only the lines of its kind, set by the coupon's `applies_to`. A `subscriptions` coupon discounts the plan's first `subscription_payments` payments when the `payment_method` can carry the discount: Stripe, which carries more than one discounted payment only on a monthly or yearly plan, or a method the store bills in house. On PayPal, Razorpay, or a weekly or daily Stripe plan discounted for more than one payment, the plan sells at its full price and the coupon is left off the order without an error. A coupon for the other kind, a `subscriptions` coupon on a plan that starts with a trial, and a fixed amount that would make the plan free each return `422` on `coupon`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create --body '{"customer_email":"maya@example.com","payment_method":"STRIPE","product_variants":{"4321":{"quantity":1}}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer-email | Yes | Value for customer email. |
| --payment-method | Yes | Value for payment method. |
| --product-variants | Yes | Value for product variants. |
| --customer-ip | No | The IP address the customer is buying from. Send it when you sell a plan with a trial: a trial of the same product taken from this address in the last 30 days counts as the customer's. Without it, only `customer_email` is checked, so a customer could take another trial with a new email address. |
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

Start or reuse checkout for a PENDING order. A new payment session returns 201; if the order already has one, that session is returned with 200. An order that costs nothing, or that the customer wallet covers in full, is marked as paid and returns 200 with a null payment.checkout_url. A free trial sold through Stripe, PayPal, or Razorpay is the exception: the customer still has to authorise the subscription with that provider, so checkout returns its payment.checkout_url and the order stays PENDING until they do. An unpaid order that sold a trial stops being payable after 72 hours, or on those three providers once the trial it sold is about to end; checkout then returns 422 on order.checkout, the order stays PENDING, and the customer has to start a new checkout. An order that is no longer PENDING also returns 422 on order.checkout. Other failures can mark the order FAILED and release wallet reservations. Do not infer payment from a checkout URL or browser return. The optional expected_status field is validated but is not a checkout concurrency precondition. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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
| --customer-ip | No | The IP address the customer is buying from. Send it when you sell a plan with a trial: a trial of the same product taken from this address in the last 30 days counts as the customer's. Without it, only `customer_email` is checked, so a customer could take another trial with a new email address. |
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

Request a real refund for a completed, nonsubscription order with refundable balance and a supported provider. amount is a decimal string in the order's `totals.refund_currency`: "5.00" means $5.00 in USD, not 500 cents. Omitting amount, sending null, or an empty string requests the FULL remaining balance. Do not send amount_cents: it is not a request field. Inspect refund.status (pending/effective/failed/cancelled); a 200 response does not guarantee money has reached the customer. Retrieve refund state before retrying an uncertain provider outcome. Send the order's `totals.refunded_total_cents` unchanged as `expected_refunded_total_cents` so that a retry under a new Idempotency-Key, or a refund someone else made meanwhile, is refused instead of refunding twice: the refund goes out only while the order's refunded total still equals it. A changed total returns `422` on `expected_refunded_total_cents`; read the order again before deciding whether to refund. An invalid `amount` returns `422` on `amount`; other refusals return `422` on `refund`, including, when you send `expected_refunded_total_cents`, a refund still pending on the order. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders create-refund 9001 --amount 5.00 --expected-refunded-total-cents 0 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --amount | No | Decimal major-unit amount in the order's `totals.refund_currency`, positive and no greater than remaining balance. Precision follows that currency. Omission/null/empty requests the full remaining balance. |
| --expected-refunded-total-cents | No | The refunded total you expect, in integer minor units of the order's `totals.refund_currency`, which can differ from the order currency: `500` means $5.00 in USD. Read it from `totals.refunded_total_cents` on the order and send it unchanged; it is `0` before the first refund. After a refund, `data.order.totals.refunded_total_cents` in the response is the value to send with the next one. With this field, the refund goes out only while the refunded total still equals it and no other refund on the order is pending. Omit it or send null to refund without the check. |

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

Create a replacement order and begin fulfillment for selected purchased variants. Supply a variant ID, an array of variant IDs, or a map of variant IDs to replacement quantities (null means original quantity). Variants must belong to this order; replacement quantities cannot exceed purchased quantity or available stock. Already replaced deliveries cannot be replaced again. The returned ID belongs to the new replacement order. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

List seller-visible fulfillment records, expanding purchased bundle snapshots into their component variants. File download links expire after one hour. A delivered record carries `delivered_product_id`, which `POST /v2/orders/{order}/dynamic-delivery-retries` takes to retry a dynamic delivery. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

## regenerate-telegram-invite

Create a Telegram invite link

Create a new Telegram invite link for a customer whose order granted access to a Telegram group or channel, for example after their first link expired or was used. Each link admits one person and expires 24 hours after it is created. SellApp saves the link on the order but does not send it: pass `invite_link` to the customer yourself. Send `variant_id` to choose the product when the order grants Telegram access for several; omit it when only one does. At most three links per product of an order are created in 24 hours, and a failed attempt does not count; a fourth returns `422` on `telegram_invite`, as do Telegram refusals such as an unhealthy bot connection or missing invite data. An order that is no longer paid, for example one refunded, disputed or voided, returns `422` on `telegram_invite` with `This order no longer grants Telegram access.` and creates no link. An order without granted Telegram access returns `422` on `order`; a product without it, or a missing `variant_id` when several products qualify, returns `422` on `variant_id`. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-a-telegram-invite-link) · Effect: **consequential**

```sh
sellapp orders regenerate-telegram-invite 9001 --variant-id 4321 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --variant-id | No | The purchased product variant whose Telegram access needs a new link, as in `line_items[].product_variant_id`. Required only when the order grants Telegram access for more than one product. |

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

Inspect the complete schema: `sellapp commands orders regenerate-telegram-invite --json`.

## retry-dynamic-delivery

Retry dynamic delivery

Queue another dynamic delivery attempt for one delivered product of a COMPLETED order. The delivered product, line item and currently dynamic variant must all belong to this order and store, with a configured webhook. 202 and meta.queued acknowledge the retry request, not successful external delivery. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders retry-dynamic-delivery 9001 --delivered-product-id 81 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --delivered-product-id | Yes | Delivered product record ID, not a catalog product or order line-item ID. Read it from `delivered_product_id` in `GET /v2/orders/{order}/deliverables`. |

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

Reset fulfillment effects for a PAID or PARTIAL order and retry delivery. The optional email overrides the recipient. meta.effects_reset counts reset effects; it does not prove fulfillment finished. Other order states and a busy lifecycle lock return 422. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Find your store's orders by transaction, customer, product, delivery, payment, status, currency, placement time, or total, newest first unless `sort` says otherwise. Every filter must match. Combine a `gte` and an `lte` filter on `created_at` or `total` for a range; a `gte` value above the `lte` value returns `422` on `filters`. An unknown field, an operator the field does not take, or a value of the wrong type returns `422` on its position, such as `filters.0.operator`. Requires the invoice token ability. Read-only and safe to repeat. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/search-orders) · Effect: **read**

```sh
sellapp orders search --body '{"filters":[{"field":"created_at","operator":"gte","value":"2026-07-01T00:00:00Z"},{"field":"created_at","operator":"lte","value":"2026-07-31T23:59:59Z"},{"field":"total","operator":"gte","value":1000},{"field":"currency","operator":"eq","value":"USD"},{"field":"status","operator":"in","value":["COMPLETED"]}],"sort":[{"field":"created_at","direction":"desc"}],"pagination":{"page":1,"limit":25}}'
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

Send delivery notifications for each selected purchased variant. Omit product_variant_ids to select all variants. A supplied list must contain distinct positive IDs belonging to this order and store. Noncompleted orders or a missing recipient return meta.notifications_sent=0. The count records notifications requested, not email delivery confirmation. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Change order lifecycle state. COMPLETED finalizes fulfillment of an already paid/partial/disputing order; it does not capture a pending payment. VOIDED can schedule provider side effects. REVIEW and DISPUTING have state restrictions. Send expected_status to prevent overwriting a changed state; an invalid transition, stale status, or busy lifecycle lock returns 422. To record a payment the store received outside SellApp, send `status: COMPLETED` with `payment_received: true` and `expected_status` set to the status you last read; this is the dashboard's manual completion. A `PENDING`, `REVIEW`, or `VOIDED` order then becomes `PAID` and delivery to the customer starts; a `PAID` or `PARTIAL` order has its delivery finished and becomes `COMPLETED`. An order can be marked paid by hand only when its total is at most 100000 minor units of its currency ($1,000.00 in USD), unless it is a custom payment method order in `REVIEW`. A larger order, or an order in any other status, returns `422` on `status`; a status that changed since you read it returns `422` on `expected_status`. Without `payment_received`, a `PENDING` order cannot be completed. Requires the `invoice` API-key ability or the `orders:write` OAuth scope, and current store permission. `Idempotency-Key` is required. Reuse the same key and identical request after a lost response; changed input or an in-progress request returns `409`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/orders/create-and-operate) · Effect: **consequential**

```sh
sellapp orders update-status 9001 --status COMPLETED --expected-status PAID --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --status | Yes | Value for status. |
| --expected-status | No | Expected current status. A status update checks this under the order lock and returns 422 if the order has changed. |
| --payment-received | No | Send `true` to confirm that the store received the payment for this order outside SellApp, which can release purchased products to the customer. `true` is accepted only with `status: COMPLETED` (otherwise `422` on `payment_received`) and requires `expected_status` (otherwise `422` on `expected_status`). `false` is accepted with any status and changes nothing. |

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

