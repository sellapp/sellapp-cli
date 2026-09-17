# bookings

[All commands](../commands.md)

## calendar-events list

List booking date overrides

List store-wide or variant-specific booking date overrides. Date boundaries are stored in UTC and interpreted in the booking timezone. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/list-booking-date-overrides) · Effect: **read**

```sh
sellapp bookings calendar-events list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product-variant-id | No | Return the effective override hierarchy for a same-store booking variant, including inherited store-wide and product-wide rows plus exact variant rows. |
| --from | No | Only return overrides ending after this instant. |
| --to | No | Only return overrides starting before this instant. |
| --status | No | Filter by override status. |
| --pagination | No | Set to false to return at most 100 matching overrides in one data array without pagination links or metadata. |
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

Inspect the complete schema: `sellapp commands bookings calendar-events list --json`.

## calendar-events set

Set booking date availability

Atomically mark one or more local calendar dates unavailable or available for the store or one same-store booking variant. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/set-booking-date-availability) · Effect: **consequential**

```sh
sellapp bookings calendar-events set --body '{"product_variant_id":73,"dates":["2028-03-26","2028-03-27"],"available":false}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --dates | Yes | Value for dates. |
| --available | Yes | Set false to block the dates and true to clear the exact block or open an inherited block for the variant. |
| --product-variant-id | No | Value for product variant id. |

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

Inspect the complete schema: `sellapp commands bookings calendar-events set --json`.

## cancel

Cancel an appointment

Cancel a store-owned appointment. Cancellation is the only supported administrative status transition. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/update-an-appointment) · Effect: **consequential**

```sh
sellapp bookings cancel 018f61d6-1c46-7b42-8a94-522bc6b5c53f --status cancelled --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| booking (positional) | Yes | The appointment UUID or numeric ID. |
| --status | Yes | Value for status. |

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

Inspect the complete schema: `sellapp commands bookings cancel --json`.

## get

Retrieve an appointment

Retrieve a store-scoped appointment by UUID or numeric ID. Provider error details are intentionally not exposed. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/retrieve-an-appointment) · Effect: **read**

```sh
sellapp bookings get 018f61d6-1c46-7b42-8a94-522bc6b5c53f
```

| Input | Required | Meaning |
| --- | --- | --- |
| booking (positional) | Yes | The appointment UUID or numeric ID. |

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

Inspect the complete schema: `sellapp commands bookings get --json`.

## list

List appointments

List appointments for the selected store. Requires the `invoice` Sanctum ability and invoice permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/list-appointments) · Effect: **read**

```sh
sellapp bookings list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |

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

Inspect the complete schema: `sellapp commands bookings list --json`.

## search

Search appointments

Search appointment UUIDs or customer email addresses and compose store-scoped filters and sorting. The public `id` filter targets the UUID returned as `id`; `uuid` remains an equivalent compatibility filter. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/search-appointments) · Effect: **read**

```sh
sellapp bookings search --body '{"filters":[{"field":"id","operator":"=","value":"018f61d6-1c46-7b42-8a94-522bc6b5c53f"}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --search | No | Value for search. |
| --includes | No | Value for includes. |

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

Inspect the complete schema: `sellapp commands bookings search --json`.

## update

Update an appointment

Reschedule an appointment atomically. Current availability is verified under the same conflict lock used by checkout holds. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/update-an-appointment) · Effect: **consequential**

```sh
sellapp bookings update 018f61d6-1c46-7b42-8a94-522bc6b5c53f --slot-start-at '2028-03-27T10:00:00+01:00' --timezone Europe/London --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| booking (positional) | Yes | The appointment UUID or numeric ID. |
| --slot-start-at | Yes | A future slot start including an explicit UTC offset. |
| --timezone | No | IANA timezone used for appointment display and availability rules. |

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

Inspect the complete schema: `sellapp commands bookings update --json`.

