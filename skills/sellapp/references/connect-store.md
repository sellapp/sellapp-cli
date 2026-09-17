# Connect and choose or create a store

Establish a connection and an explicit target for commerce operations.

Contract: `fe25fd396af7918889c08bfa13e4285af1bb53f0e9b2a0cadba604cdf7f33547`; workflow: `20a92459d3c39dddcce05c1f6de46904ccb042b126c5a3408bbaa16009409ef1`. Source: `content/docs/api/(guides)/oauth-authorization.mdx`.

## Prerequisites

A SellApp user account. SDK/REST account reads require account:read; creation requires stores:create on an unrestricted API key. CLI and MCP retain browser OAuth.

## Preparation

Run sellapp setup, or sellapp login followed by store discovery. Browser consent grants admin across current and future accessible stores; current membership and role permissions still apply. A connection with zero stores is successful authentication.

Use sellapp login --no-browser when the browser cannot open. Never extract tokens or register a customer OAuth application. Ordinary SDK integrations use user-owned API keys; account calls omit X-STORE and store calls select one slug.

### 1. preparation

Inspect accessible stores and retain their IDs and slugs. An empty list is successful. SDK/REST callers use account:read without a configured store; selected-store restrictions filter the list.

Operation: `listStores`. Effect: read. Target: account.

Inspect `sellapp stores list --schema`; execute `sellapp stores list` with the required inputs. Omit the store argument and header.

MCP: `sellapp_list_stores` in full mode.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 2. execution (when needed)

Only when creation is intended: provide name and slug, retain the same Idempotency-Key and body across retries, and keep the returned ID and slug. Visibility defaults to PUBLIC.

Operation: `createStore`. Effect: write. Target: account.

Inspect `sellapp stores create --schema`; execute `sellapp stores create` with the required inputs. Omit the store argument and header.

MCP: `sellapp_create_store` in full mode; inspect `sellapp_get_operation` with operation `createStore`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Required parameters: `header.Idempotency-Key`.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "name": "Launch Lab",
  "slug": "launchlab"
}
```

Recovery for this operation: Retain the same Idempotency-Key and exact request for the same intended action; follow the operation retention and conflict rules.

### 3. verification

For a selected store, request a bounded first page. In CLI use sellapp stores use SLUG, or stores create --use for explicit local selection. MCP business calls carry the store slug; account calls omit it.

Operation: `listProducts`. Effect: read. Target: store.

Inspect `sellapp products list --schema`; execute `sellapp products list` with the required inputs. Use the explicit intended store.

MCP: `sellapp_list_products` in full mode; inspect `sellapp_get_operation` with operation `listProducts`, then use `sellapp_read` in compact mode.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

## Verify the result

Read the selected store successfully. Creating through the API does not change dashboard selection; CLI creation without --use preserves the selected profile store.

## Recover from partial completion

Declined consent is cancellation. Invalid client or scope indicates configuration failure. For revoked grants use login; for store permission failure inspect membership and role rather than repeating login.

If local selection fails after creation, retain the successful creation result and run sellapp stores use RETURNED_SLUG. Do not create a second store.

New store creation is limited to one per user every 60 seconds across credentials. On 429, honor Retry-After and keep the same key and body. A successful same-key replay returns the original store.
