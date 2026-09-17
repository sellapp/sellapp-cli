# events

[All commands](../commands.md)

## list-integration-events

List integration events

Poll the 30-day journal with signed, filter-bound cursors. Use client-controlled backoff between ordinary requests. An expired position returns 410 cursor_expired with cursor_polling guidance. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/events) · Effect: **read**

```sh
sellapp events list-integration-events
```

| Input | Required | Meaning |
| --- | --- | --- |
| --cursor | No | Signed cursor bound to this filter set. |
| --limit | No | Maximum events to return. |
| --type | No | Return only this event type. |
| --subject-type | No | Return events for this subject type. |
| --subject-id | No | Return events for this subject identifier. |

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

Inspect the complete schema: `sellapp commands events list-integration-events --json`.

## list-order-events

List order events

Read retained events for this order. Use the returned cursor for the next request and wait between polls. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/events) · Effect: **read**

```sh
sellapp events list-order-events 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |
| --cursor | No | Signed cursor bound to this filter set. |
| --limit | No | Maximum events to return. |

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

Inspect the complete schema: `sellapp commands events list-order-events --json`.

