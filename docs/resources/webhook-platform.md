# webhook-platform

[All commands](../commands.md)

## get-webhook-delivery

Retrieve a webhook delivery

Retrieve one redacted delivery. A disabled delivery with zero attempts was retained without contacting the endpoint. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/events) · Effect: **read**

```sh
sellapp webhook-platform get-webhook-delivery 01992b31-c8bd-75b5-b02d-6ae0aa418940
```

| Input | Required | Meaning |
| --- | --- | --- |
| delivery (positional) | Yes | The delivery identifier. |

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

Inspect the complete schema: `sellapp commands webhook-platform get-webhook-delivery --json`.

## list-webhook-deliveries

List webhook deliveries

List your store's outbound webhook deliveries, newest first, 25 per page by default. Each delivery shows only the destination origin; response bodies are redacted and capped. Filter by `event_type` or `status`; both must match exactly, and a value that matches nothing returns an empty page rather than an error. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/events) · Effect: **read**

```sh
sellapp webhook-platform list-webhook-deliveries
```

| Input | Required | Meaning |
| --- | --- | --- |
| --page | No | Page number, starting at 1. |
| --limit | No | Deliveries per page. Values below 1 are read as 1 and values above 100 as 100. |
| --event-type | No | Keep deliveries of this event type, as listed by `GET /v2/webhook-event-types`. |
| --status | No | Keep deliveries in this state: `pending`, `retrying`, `succeeded`, `failed`, or `disabled`. |

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

Inspect the complete schema: `sellapp commands webhook-platform list-webhook-deliveries --json`.

## list-webhook-event-types

List webhook event types

List versioned outbound event types. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/events) · Effect: **read**

```sh
sellapp webhook-platform list-webhook-event-types
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

Inspect the complete schema: `sellapp commands webhook-platform list-webhook-event-types --json`.

## replay-webhook-delivery

Replay a webhook delivery

Accept no URL, payload, or query overrides. Resolve the current channel server-side, create a new delivery ID, retain the integration-event ID, and sign with a new timestamp. A disabled channel returns 422: fix the endpoint and save the channel before replaying. A successful test alone does not re-enable delivery. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/events) · Effect: **consequential**

```sh
sellapp webhook-platform replay-webhook-delivery delivery_01K4CUSTOMER --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| delivery (positional) | Yes | The delivery identifier. |

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

Inspect the complete schema: `sellapp commands webhook-platform replay-webhook-delivery --json`.

