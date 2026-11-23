# disputes

[All commands](../commands.md)

## get-dispute

Retrieve a dispute

Retrieve one redacted dispute. No dispute mutations are exposed. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/disputes) · Effect: **read**

```sh
sellapp disputes get-dispute 01992b10-a7d2-7b91-9822-32e8d7454e7c
```

| Input | Required | Meaning |
| --- | --- | --- |
| dispute (positional) | Yes | The dispute identifier. |

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

Inspect the complete schema: `sellapp commands disputes get-dispute --json`.

## list-disputes

List disputes

List your store's payment disputes, most recently opened first, 50 per page. Each dispute is read-only and redacted, and may reference an order or a standalone charge. Filter by `status`, `provider`, or the day a dispute was opened; filters combine, and the date filters read whole days in the store's time zone and include both ends. An unknown `status`, a date that cannot be read, or an `opened_before` earlier than `opened_after` returns `422` on that parameter. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/disputes) · Effect: **read**

```sh
sellapp disputes list-disputes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --page | No | Page number, starting at 1. |
| --status | No | Keep disputes in this state: `open`, `under_review`, `won`, `lost`, or `closed`. |
| --provider | No | Keep disputes raised through this payment provider, such as `stripe`. Any letter case works. |
| --opened-after | No | Keep disputes opened on or after this moment. A date such as `2026-08-01` starts at midnight of that day in the store's time zone; a date-time without an offset is also read in the store's time zone. Inclusive. |
| --opened-before | No | Keep disputes opened on or before this moment. A date such as `2026-08-31` runs to the last second of that day in the store's time zone; a date-time without an offset is also read in the store's time zone. Inclusive. |

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

Inspect the complete schema: `sellapp commands disputes list-disputes --json`.

