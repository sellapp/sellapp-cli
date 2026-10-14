# invoices

[All commands](../commands.md)

## create

Create an invoice

Create an invoice for one or more product variants. Credit products are direct-only: send a credits variant by itself, use a quantity covered by its credit rate tiers, and respect the variant quantity increment rules. A subscription plan that offers a trial is priced at the trial, nothing for a free trial and the trial price for a paid one, unless the customer already had a trial of that product; then it is priced at the plan price. Two requests racing for the same trial can be refused with a `422` on `invoice`. A cart carries at most one subscription plan and nothing else; a plan mixed with other products or with a second plan returns `422` on `product_variants.{variant_id}`, and a plan's `quantity` other than 1 returns `422` on `quantity`. A customer who already holds the plan, matched by `customer_email`, is refused with `422` on `invoice` while that subscription is active, trialing, past due, or unpaid, whichever payment method bills it. A subscription scheduled to end still blocks while it can be kept, and the message then tells the customer to keep it from the customer portal. One that can no longer be kept, such as a PayPal subscription cancelled at period end, does not block. For an in-house subscription on hold for an unpaid renewal, the message tells the customer to pay and resume it from the customer portal. A `coupon` discounts only the lines of its kind, set by the coupon's `applies_to`. A `subscriptions` coupon discounts the plan's first `subscription_payments` payments when the `payment_method` can carry the discount: Stripe, which carries more than one discounted payment only on a monthly or yearly plan, or a method the store bills in house. On PayPal, Razorpay, or a weekly or daily Stripe plan discounted for more than one payment, the plan sells at its full price and the coupon is left off the order without an error. A coupon for the other kind, a `subscriptions` coupon on a plan that starts with a trial, and a fixed amount that would make the plan free each return `422` on `coupon`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/create-an-invoice) · Effect: **consequential**

```sh
sellapp invoices create --body '{"customer_email":"maya.chen@example.com","payment_method":"STRIPE","product_variants":{"4321":{"quantity":1}}}' --yes
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

Inspect the complete schema: `sellapp commands invoices create --json`.

## create-refund

Create invoice refund

Request a full or partial provider refund using the same idempotent refund ledger as the dashboard. Amounts are decimal currency strings, never floating-point numbers. Requires the `invoice` token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/refund-an-invoice) · Effect: **consequential**

```sh
sellapp invoices create-refund 1 --amount 12.50 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --amount | No | Optional partial refund amount in major currency units. Omit for the full remaining refundable amount. |

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

Inspect the complete schema: `sellapp commands invoices create-refund --json`.

## get

Retrieve an invoice

Retrieve an invoice by its ID to check its current payment and delivery state. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/retrieve-an-invoice) · Effect: **read**

```sh
sellapp invoices get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |

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

Inspect the complete schema: `sellapp commands invoices get --json`.

## get-deliverables

View invoice deliverables

Retrieve the deliverables sent to the customer, including each product in a multi-product purchase. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/view-invoice-deliverables) · Effect: **read**

```sh
sellapp invoices get-deliverables 1234
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |

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

Inspect the complete schema: `sellapp commands invoices get-deliverables --json`.

## go-to-checkout

Create a checkout session

Create or reuse a payment session for a pending invoice. Create the invoice first, then send the customer to the returned top-level `payment_url`. A new session answers `201`; if the invoice already has one, that session is returned with `200`. An invoice that costs nothing, or that the customer wallet covers in full, is marked as paid and answers `200` without a `payment_url`; delivery may still be queued. A free trial sold through Stripe, PayPal, or Razorpay is the exception: the customer still has to authorise the subscription with that provider, so it answers with a `payment_url` and the invoice stays pending until they do. An unpaid trial invoice stops being payable after 72 hours, or on those three providers once the trial it sold is about to end, and then answers `422` on `invoice.checkout`. Check for a URL before redirecting and verify payment through signed events or a fresh read. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/create-a-checkout-session) · Effect: **consequential**

```sh
sellapp invoices go-to-checkout 9001 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |

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

Inspect the complete schema: `sellapp commands invoices go-to-checkout --json`.

## issue-replacement

Issue replacement for completed invoice

Issue replacement deliverables for a completed purchase. This creates a new invoice and marks it completed, then starts the normal delivery and sales-notification flow. It is not just an edit to the old delivery. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/issue-replacement-for-completed-invoice) · Effect: **consequential**

```sh
sellapp invoices issue-replacement 1 --body '{"product_variants":[117214]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
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

Inspect the complete schema: `sellapp commands invoices issue-replacement --json`.

## list

List all invoices

List your store's invoices, 15 per page by default. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/list-all-invoices) · Effect: **read**

```sh
sellapp invoices list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --search | No | Free-text search term. |
| --search-by | No | Which invoice field the search term should be matched against. |
| --id | No | Filter by invoice ID. |
| --email | No | Filter by customer email. |
| --transaction-id | No | Filter by payment transaction ID. |
| --serial-code | No | Filter by delivered serial. |
| --additional-info | No | Filter by customer-provided additional information. |
| --product-name | No | Filter by product or variant title. |
| --discord-data | No | Filter by attached Discord data. |
| --crypto-txid | No | Filter by crypto TXID. |
| --crypto-address | No | Filter by crypto payment address. |
| --coupon-code | No | Filter by coupon code. |
| --status | No | Filter by one or more invoice statuses. Repeat the parameter to select several. |
| --payment-methods | No | Filter by one or more payment methods. Repeat the parameter to select several. |
| --sort | No | Sort order for the result set. |

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

Inspect the complete schema: `sellapp commands invoices list --json`.

## mark-completed

Mark pending invoice completed

Mark a pending invoice completed and start the normal delivery flow. SellApp normally handles completion after payment is confirmed. Use this only after independently confirming payment or deliberately authorizing delivery without it; the action can release purchased products. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/mark-pending-invoice-completed) · Effect: **consequential**

```sh
sellapp invoices mark-completed 1 --expected-status PENDING --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --expected-status | No | Optional optimistic-concurrency guard. The mutation fails when the current order status no longer matches this value. |

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

Inspect the complete schema: `sellapp commands invoices mark-completed --json`.

## mark-voided

Mark pending invoice voided

Void a pending invoice so it does not proceed to product delivery. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/mark-pending-invoice-voided) · Effect: **consequential**

```sh
sellapp invoices mark-voided 1 --expected-status PENDING --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --expected-status | No | Optional optimistic-concurrency guard. The mutation fails when the current order status no longer matches this value. |

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

Inspect the complete schema: `sellapp commands invoices mark-voided --json`.

## notify-fulfillment

Create fulfillment notifications

Resend fulfillment notifications for all or selected same-store product variants on a completed invoice. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/resend-invoice-deliverables) · Effect: **consequential**

```sh
sellapp invoices notify-fulfillment 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --email | No | Value for email. |
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

Inspect the complete schema: `sellapp commands invoices notify-fulfillment --json`.

## retry-dynamic-delivery

Create dynamic delivery retry

Retry one same-store dynamic delivered product after validating order, line-item, variant, and webhook ownership. Callers should avoid concurrent retries for the same delivered product. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/retry-dynamic-delivery) · Effect: **consequential**

```sh
sellapp invoices retry-dynamic-delivery 1 --delivered-product-id 42 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --delivered-product-id | Yes | Value for delivered product id. |

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

Inspect the complete schema: `sellapp commands invoices retry-dynamic-delivery --json`.

## retry-fulfillment

Create fulfillment retry

Retry failed delivery steps for a paid or partially fulfilled purchase. Repeated requests retry the same failed steps without creating duplicate work. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/retry-invoice-fulfillment) · Effect: **consequential**

```sh
sellapp invoices retry-fulfillment 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --email | No | Optional replacement delivery email. |

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

Inspect the complete schema: `sellapp commands invoices retry-fulfillment --json`.

## search

Search invoices

Search invoices with the same filters as the list endpoint, sent in a JSON body instead of query parameters. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/search-invoices) · Effect: **read**

```sh
sellapp invoices search --id 1 --sort -created_at
```

| Input | Required | Meaning |
| --- | --- | --- |
| --search | No | Value for search. |
| --search-by | No | Value for search by. |
| --id | No | Value for id. |
| --email | No | Value for email. |
| --transaction-id | No | Value for transaction id. |
| --serial-code | No | Value for serial code. |
| --additional-info | No | Value for additional info. |
| --product-name | No | Value for product name. |
| --discord-data | No | Value for discord data. |
| --crypto-txid | No | Value for crypto txid. |
| --crypto-address | No | Value for crypto address. |
| --coupon-code | No | Value for coupon code. |
| --status | No | Value for status. |
| --payment-methods | No | Value for payment methods. |
| --sort | No | Value for sort. |

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

Inspect the complete schema: `sellapp commands invoices search --json`.

## update-status

Update invoice status

Change a purchase through the invoice compatibility operation. Set status to COMPLETED, VOIDED, REVIEW, or DISPUTING. Dispute transitions schedule community-access revocation and a dispute webhook. Requires the `invoice` credential ability and store ownership. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/invoices/update-invoice-status) · Effect: **consequential**

```sh
sellapp invoices update-status 1 --status REVIEW --expected-status PENDING --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| invoice (positional) | Yes | The invoice path parameter. |
| --status | Yes | Value for status. |
| --expected-status | No | Optional optimistic-concurrency guard. The mutation fails when the current order status no longer matches this value. |

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

Inspect the complete schema: `sellapp commands invoices update-status --json`.

