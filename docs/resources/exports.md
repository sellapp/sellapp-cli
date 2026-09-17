# exports

[All commands](../commands.md)

## create-export

Create an export

Create a queued sales, tax, customer or affiliate payout report. One export may be queued or running per store, and two per requesting account across all credentials and stores. When capacity is full, HTTP 429 includes Retry-After: 60; poll the existing export and wait before creating another. Date ranges include both end dates and may span at most 366 calendar days. Omitted or null to defaults to today; omitted or null from defaults to 365 days before to. Sales and tax use sale update dates; customers and affiliate payouts use creation dates. Reports are limited to 10,000 source rows, 5 MiB of serialized source data and a 10 MiB output file. Tax summaries count source sales toward these limits. A report that exceeds a limit fails without publishing a partial file; request a smaller date range. Jobs may wait up to 15 minutes to start and run for at most 60 seconds. The requesting actor is recorded but is not returned in this response. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request. Sales, tax and customer reports require invoice store permission. Affiliate payout reports require affiliate and affiliate:payout permissions. The category selects the applicable permission checks; listing returns only authorized categories.

[API reference](https://sell.app/docs/api/exports) · Effect: **consequential**

```sh
sellapp exports create-export --body '{"type":"sales","format":"csv","parameters":{"from":"2026-08-01","to":"2026-08-31"}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --type | Yes | Value for type. |
| --format | Yes | Value for format. |
| --parameters | No | Value for parameters. |

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

Inspect the complete schema: `sellapp commands exports create-export --json`.

## download-export

Download an export

Validate the five-minute signed API URL, then redirect to a short-lived private object-storage URL. The API does not stream the file. Pending or running exports return 409; expired exports return 410. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request. Sales, tax and customer reports require invoice store permission. Affiliate payout reports require affiliate and affiliate:payout permissions. The category selects the applicable permission checks; listing returns only authorized categories.

[API reference](https://sell.app/docs/api/exports) · Effect: **read**

```sh
sellapp exports download-export 01992a65-e064-71ba-b38f-902b7966a6be --expires 1788513423 --signature 2c91df645a086ec399153a932b741f809d2b85c69740eaf3612384ebfb913a65
```

| Input | Required | Meaning |
| --- | --- | --- |
| export (positional) | Yes | The export identifier. |
| --expires | Yes | Expiry timestamp from the generated download URL. |
| --signature | Yes | Copy the signature from the response-supplied download_url. The example is illustrative and cannot authorize a download. |

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

Inspect the complete schema: `sellapp commands exports download-export --json`.

## get-export

Retrieve an export

Poll until completed, failed, or expired. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request. Sales, tax and customer reports require invoice store permission. Affiliate payout reports require affiliate and affiliate:payout permissions. The category selects the applicable permission checks; listing returns only authorized categories.

[API reference](https://sell.app/docs/api/exports) · Effect: **read**

```sh
sellapp exports get-export 01992a65-e064-71ba-b38f-902b7966a6be
```

| Input | Required | Meaning |
| --- | --- | --- |
| export (positional) | Yes | The export identifier. |

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

Inspect the complete schema: `sellapp commands exports get-export --json`.

## list-exports

List exports

List permitted exports, 50 per page. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request. Sales, tax and customer reports require invoice store permission. Affiliate payout reports require affiliate and affiliate:payout permissions. The category selects the applicable permission checks; listing returns only authorized categories.

[API reference](https://sell.app/docs/api/exports) · Effect: **read**

```sh
sellapp exports list-exports
```

| Input | Required | Meaning |
| --- | --- | --- |

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

Inspect the complete schema: `sellapp commands exports list-exports --json`.

