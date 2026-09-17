# Create checkout and verify order state

Create an order and distinguish its checkout URL from confirmed payment.

Contract: `fe25fd396af7918889c08bfa13e4285af1bb53f0e9b2a0cadba604cdf7f33547`; workflow: `20a92459d3c39dddcce05c1f6de46904ccb042b126c5a3408bbaa16009409ef1`. Source: `content/docs/api/(guides)/create-a-checkout.mdx`.

## Prerequisites

A valid product variant, configured payment method and customer details permitted by the schema.

## Preparation

Configure and test webhook handling before relying on payment notifications. Creation can affect stock; free and wallet orders can complete immediately. Inspect the operation effect and obtain any missing authorization for consequential changes.

### 1. preparation

Inspect current variant availability, price and allowed payment methods.

Operation: `getProductVariant`. Effect: read. Target: store.

Inspect `sellapp product-variants get --schema`; execute `sellapp product-variants get` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_product_variant` in full mode; inspect `sellapp_get_operation` with operation `getProductVariant`, then use `sellapp_read` in compact mode.

Required parameters: `path.product`, `path.variant`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 2. execution

Create checkout with explicit customer, variant and payment details. Preserve the supported Idempotency-Key, order ID and returned checkout URL.

Operation: `createOrder`. Effect: consequential. Target: store.

Inspect `sellapp orders create --schema`; execute `sellapp orders create` with the required inputs. Use the explicit intended store.

MCP: `sellapp_create_order` in full mode; inspect `sellapp_get_operation` with operation `createOrder`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Required parameters: `header.Idempotency-Key`.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "customer_email": "maya@example.com",
  "payment_method": "STRIPE",
  "product_variants": {
    "4321": {
      "quantity": 1
    }
  }
}
```

Recovery for this operation: Retain the same Idempotency-Key and exact request for the same intended action; follow the operation retention and conflict rules.

### 3. verification

Use the retained order ID to inspect payment and fulfillment state.

Operation: `getOrder`. Effect: read. Target: store.

Inspect `sellapp orders get --schema`; execute `sellapp orders get` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_order` in full mode; inspect `sellapp_get_operation` with operation `getOrder`, then use `sellapp_read` in compact mode.

Required parameters: `path.order`.

Use returned values: `path.order` from `createOrder:data.id`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

## Verify the result

A browser return or checkout URL is not payment success. Verify current Order status and sale-line fulfillment independently.

## Recover from partial completion

After an uncertain response, reuse the same key and body only under the operation’s documented idempotency rules. Inspect the order before creating another.

Handle duplicate and out-of-order webhooks, and reconcile by reading current order state.
