# Create a product and variant

Create a catalog product with a purchasable variant and verify both records.

Contract: `fe25fd396af7918889c08bfa13e4285af1bb53f0e9b2a0cadba604cdf7f33547`; workflow: `20a92459d3c39dddcce05c1f6de46904ccb042b126c5a3408bbaa16009409ef1`. Source: `content/docs/api/(guides)/sell-a-product.mdx`.

## Prerequisites

A selected store and current catalog write permission. Payment methods must already be configured.

## Preparation

Read the operation schemas and inspect the current catalog before creating duplicates. Choose deliverable and pricing types deliberately. Amounts use documented minor units.

HIDDEN controls listing visibility; it is not an access-control boundary. MANUAL delivery requires later fulfillment. Binary uploads use the existing CLI --upload or SDK paths; MCP does not read local files.

Alternatively, createProduct can include initial variants atomically; retain their returned IDs and skip separate variant creation. A bundle uses one BUNDLE variant and bundle_items. The separate-step walkthrough remains valid.

### 1. preparation

Inspect a bounded page for an existing matching product.

Operation: `listProducts`. Effect: read. Target: store.

Inspect `sellapp products list --schema`; execute `sellapp products list` with the required inputs. Use the explicit intended store.

MCP: `sellapp_list_products` in full mode; inspect `sellapp_get_operation` with operation `listProducts`, then use `sellapp_read` in compact mode.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 2. execution

Create the product and retain data.id. Do not infer checkout readiness from creation.

Operation: `createProduct`. Effect: write. Target: store.

Inspect `sellapp products create --schema`; execute `sellapp products create` with the required inputs. Use the explicit intended store.

MCP: `sellapp_create_product` in full mode; inspect `sellapp_get_operation` with operation `createProduct`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "description": "Templates for your next project.",
  "title": "Design kit",
  "visibility": "HIDDEN"
}
```

Recovery for this operation: Do not automatically repeat an uncertain write. Inspect current state first.

### 3. execution

Create a variant under the returned product ID, using a configured payment method and valid pricing/deliverable data. Retain the variant ID.

Operation: `createProductVariant`. Effect: write. Target: store.

Inspect `sellapp product-variants create --schema`; execute `sellapp product-variants create` with the required inputs. Use the explicit intended store.

MCP: `sellapp_create_product_variant` in full mode; inspect `sellapp_get_operation` with operation `createProductVariant`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Required parameters: `path.product`.

Use returned values: `path.product` from `createProduct:data.id`.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "deliverable": {
    "data": {
      "comment": "We will send your reading-room invitation.",
      "stock": null
    },
    "types": [
      "MANUAL"
    ]
  },
  "description": "One operating memo each month; access is provisioned by our team.",
  "payment_methods": [
    "STRIPE"
  ],
  "pricing": {
    "humble": false,
    "price": {
      "currency": "USD",
      "price": 1999
    }
  },
  "title": "Monthly membership"
}
```

Recovery for this operation: Do not automatically repeat an uncertain write. Inspect current state first.

### 4. verification

Read the created product using its retained ID.

Operation: `getProduct`. Effect: read. Target: store.

Inspect `sellapp products get --schema`; execute `sellapp products get` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_product` in full mode; inspect `sellapp_get_operation` with operation `getProduct`, then use `sellapp_read` in compact mode.

Required parameters: `path.product`.

Use returned values: `path.product` from `createProduct:data.id`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 5. verification

Read the created variant and verify pricing, deliverable and payment methods.

Operation: `getProductVariant`. Effect: read. Target: store.

Inspect `sellapp product-variants get --schema`; execute `sellapp product-variants get` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_product_variant` in full mode; inspect `sellapp_get_operation` with operation `getProductVariant`, then use `sellapp_read` in compact mode.

Required parameters: `path.product`, `path.variant`.

Use returned values: `path.variant` from `createProductVariant:data.id`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

## Verify the result

Confirm the product and variant belong to the intended store and expose the requested configuration.

## Recover from partial completion

Product and variant creation are separate writes. If the second step fails, inspect the retained product and resume there.

When a non-idempotent creation has an uncertain response, inspect existing records before retrying. Use expected_updated_at where update schemas require it.
