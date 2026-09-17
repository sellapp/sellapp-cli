# Create and retrieve exports

Queue an export, poll within finite bounds and retrieve the completed file.

Contract: `fe25fd396af7918889c08bfa13e4285af1bb53f0e9b2a0cadba604cdf7f33547`; workflow: `20a92459d3c39dddcce05c1f6de46904ccb042b126c5a3408bbaa16009409ef1`. Source: `content/docs/api/(resources)/exports/index.mdx`.

## Prerequisites

The selected store and permission for the requested export category.

## Preparation

Use sales, tax, customers or affiliate_payouts with CSV or JSON. Date ranges are limited to 366 inclusive days; omitted dates cover today and the preceding 365 days. Only one export per store and two per account may be queued or running. On HTTP 429, respect Retry-After and poll the existing export. Reports over 10,000 source rows, 5 MiB of source data or 10 MiB of output fail without a partial download; use a smaller date range.

### 1. execution

Create one intended export with a stable Idempotency-Key and retain data.id.

Operation: `createExport`. Effect: consequential. Target: store.

Inspect `sellapp exports create-export --schema`; execute `sellapp exports create-export` with the required inputs. Use the explicit intended store.

MCP: `sellapp_create_export` in full mode; inspect `sellapp_get_operation` with operation `createExport`, then use `sellapp_write` in compact mode. Supply a stable idempotency_key. Consequential writes first return a preview; once authorized repeat the exact arguments with its confirmation token.

Canonical request body example (replace fixture values with the intended inputs and retained IDs):

```json
{
  "format": "csv",
  "parameters": {
    "from": "2026-08-01",
    "to": "2026-08-31"
  },
  "type": "sales"
}
```

Recovery for this operation: Retain the same Idempotency-Key and exact request for the same intended action; follow the operation retention and conflict rules.

### 2. verification

Poll pending/running work with finite limits and respect meta.poll_after_seconds. Inspect failures; only completed work is downloadable.

Operation: `getExport`. Effect: read. Target: store.

Inspect `sellapp exports get-export --schema`; execute `sellapp exports get-export` with the required inputs. Use the explicit intended store.

MCP: `sellapp_get_export` in full mode; inspect `sellapp_get_operation` with operation `getExport`, then use `sellapp_read` in compact mode.

Required parameters: `path.export`.

Use returned values: `path.export` from `createExport:data.id`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

### 3. execution

Use the existing CLI export download or SDK file path after completion. Signed download URLs expire after five minutes; never forward API credentials to the storage host. This step is unavailable through hosted MCP.

Operation: `downloadExport`. Effect: read. Target: store.

Inspect `sellapp exports download-export --schema`; execute `sellapp exports download-export` with the required inputs. Use the explicit intended store.

This step requires the CLI or SDK file transfer path. Hosted MCP cannot read local paths or arbitrary file contents.

Required parameters: `path.export`, `query.expires`, `query.signature`.

Use returned values: `path.export` from `createExport:data.id`.

Recovery for this operation: Read again within finite pagination/polling bounds; respect Retry-After.

## Verify the result

Verify a completed export and successful file retrieval; queued work is not a downloaded report.

## Recover from partial completion

Inspect a failed export before creating a new attempt with a new key. An expired export can still report completed; download may return 410.

Reusing the same key retrieves the same operation rather than restarting it. Keep signed URLs and exported customer data out of logs.
