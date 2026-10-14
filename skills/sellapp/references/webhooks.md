# Configure, test and inspect webhooks

Create an event subscription and verify delivery and processing separately.

Contract: `b9eb7046aa1e9bcd12c19188a0f2f4f4377ee48aa813bddaf984c3c58ea4e92e`; workflow: `8b2d02e5e08f2d6b53bf9ae26b890671afd472d72346df57b4e6821f335e6fbf`. Source: `content/docs/api/(guides)/webhooks.mdx`.

## Prerequisites

A reachable endpoint with signature verification and durable processing.

## Preparation

Verify the raw request bytes with the documented HMAC and constant-time comparison. Store the verified payload and recoverable work before returning a successful response. Handle duplicates and out-of-order events.

### 1. preparation

Choose supported event names from the current API.

Operation: `listWebhookEventTypes`. Effect: read. Target: store.

Inspect `sellapp webhook-platform list-webhook-event-types --schema`; execute `sellapp webhook-platform list-webhook-event-types` with the required inputs. Use the explicit intended store.

MCP: `sellapp_list_webhook_event_types` in full mode; inspect `sellapp_get_operation` with operation `listWebhookEventTypes`, then use `sellapp_read` in compact mode.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 2. execution

Create the intended channel with explicit URL and events. Keep its ID; treat signing secrets as secrets.

Operation: `createWebhookChannel`. Effect: consequential. Target: store.

Inspect `sellapp webhook-channels create --schema`; execute `sellapp webhook-channels create` with the required inputs. Use the explicit intended store.

MCP: `sellapp_create_webhook_channel` in full mode; inspect `sellapp_get_operation` with operation `createWebhookChannel`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "allowed_notifications": [
    "order.created",
    "order.paid"
  ],
  "name": "Ship It webhook",
  "url": "https://example.com/webhooks/ship-it"
}
```

Recovery for this operation: Do not automatically repeat an uncertain write. Inspect current state first.

### 3. execution

Send a test to the created channel and inspect the returned channel, event and status. Then discover any recorded delivery through delivery listing.

Operation: `sendTestWebhook`. Effect: consequential. Target: store.

Inspect `sellapp webhook-channels send --schema`; execute `sellapp webhook-channels send` with the required inputs. Use the explicit intended store.

MCP: `sellapp_send_test_webhook` in full mode; inspect `sellapp_get_operation` with operation `sendTestWebhook`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Required parameters: `path.webhookChannel`.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "event": "order.created"
}
```

Recovery for this operation: Do not automatically repeat an uncertain write. Inspect current state first.

### 4. verification

Inspect a bounded delivery page for the channel.

Operation: `listWebhookDeliveries`. Effect: read. Target: store.

Inspect `sellapp webhook-platform list-webhook-deliveries --schema`; execute `sellapp webhook-platform list-webhook-deliveries` with the required inputs. Use the explicit intended store.

MCP: `sellapp_list_webhook_deliveries` in full mode; inspect `sellapp_get_operation` with operation `listWebhookDeliveries`, then use `sellapp_read` in compact mode.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 5. verification

Inspect response and delivery state; also verify processing in the receiver.

Operation: `getWebhookDelivery`. Effect: read. Target: store.

Inspect `sellapp webhook-platform get-webhook-delivery --schema`; execute `sellapp webhook-platform get-webhook-delivery` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_webhook_delivery` in full mode; inspect `sellapp_get_operation` with operation `getWebhookDelivery`, then use `sellapp_read` in compact mode.

Required parameters: `path.delivery`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

## Verify the result

A successful delivery does not prove the receiver completed its business work.

## Recover from partial completion

Inspect failures before replaying. Replay is a consequential operation and can repeat external effects.

Do not log signing secrets, authorization headers or customer data unnecessarily. Use existing webhook inspection and bounded event-polling CLI helpers.
