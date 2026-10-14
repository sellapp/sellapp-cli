# Manage subscriptions with capability checks

Apply only the subscription action currently supported for its provider.

Contract: `23a0b45478edca19e5ab59d7c5b7bfffe2298eb336e937126ee04dcbbf13c3d0`; workflow: `8b2d02e5e08f2d6b53bf9ae26b890671afd472d72346df57b4e6821f335e6fbf`. Source: `content/docs/api/(guides)/handle-subscriptions.mdx`.

## Prerequisites

An integer SellApp subscription ID from the order’s sale line, not the payment provider’s subscription ID.

## Preparation

The SellApp subscription record may appear after checkout. Read current capability information for the particular subscription before changing it.

### 1. preparation

Inspect current status and provider.

Operation: `getSubscription`. Effect: read. Target: store.

Inspect `sellapp subscriptions get-subscription --schema`; execute `sellapp subscriptions get-subscription` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_subscription` in full mode; inspect `sellapp_get_operation` with operation `getSubscription`, then use `sellapp_read` in compact mode.

Required parameters: `path.productSubscription`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 2. preparation

Check the exact action. Stop for unsupported. A customer_only plan change or payment method update can still be requested on the customer's behalf; any other customer_only action goes to the customer through the documented path.

Operation: `getSubscriptionCapabilities`. Effect: read. Target: store.

Inspect `sellapp subscriptions get-capabilities --schema`; execute `sellapp subscriptions get-capabilities` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_subscription_capabilities` in full mode; inspect `sellapp_get_operation` with operation `getSubscriptionCapabilities`, then use `sellapp_read` in compact mode.

Required parameters: `path.productSubscription`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 3. execution (when needed)

Only when requested and supported: schedule cancellation at period end. Read consequential preview/confirmation requirements.

Operation: `cancelSubscriptionAtPeriodEnd`. Effect: consequential. Target: store.

Inspect `sellapp subscriptions cancel-at-period-end --schema`; execute `sellapp subscriptions cancel-at-period-end` with the required inputs. Use the explicit intended store.

MCP: `sellapp_cancel_subscription_at_period_end` in full mode; inspect `sellapp_get_operation` with operation `cancelSubscriptionAtPeriodEnd`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Required parameters: `path.productSubscription`.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "reason": "Customer request"
}
```

Recovery for this operation: Retain the same Idempotency-Key and exact request for the same intended action; follow the operation retention and conflict rules.

### 4. verification

Verify cancellation timing and state; a scheduled cancellation is not immediate termination.

Operation: `getSubscription`. Effect: read. Target: store.

Inspect `sellapp subscriptions get-subscription --schema`; execute `sellapp subscriptions get-subscription` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_subscription` in full mode; inspect `sellapp_get_operation` with operation `getSubscription`, then use `sellapp_read` in compact mode.

Required parameters: `path.productSubscription`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

## Verify the result

Report the server’s current capability and resulting state without inventing provider support.

## Recover from partial completion

Re-read capabilities if state changes. Do not turn unsupported operations into repeated retries or promise plan/date changes the API rejects.

For an uncertain action, inspect current subscription state before attempting it again.
