# webhook-channels

[All commands](../commands.md)

## create

Create a webhook channel

Create an outbound webhook destination with a stable UUID and an explicit event filter. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/create-a-webhook-channel) · Effect: **consequential**

```sh
sellapp webhook-channels create --body '{"name":"Ship It webhook","url":"https://example.com/webhooks/ship-it","allowed_notifications":["order.created","order.paid"]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --url | Yes | Public HTTP or HTTPS destination. The same canonical URL cannot be configured twice for one store. |
| --allowed-notifications | Yes | Complete set of public webhook event names enabled for this destination. |
| --name | No | Optional seller-facing reference name. |

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

Inspect the complete schema: `sellapp commands webhook-channels create --json`.

## delete

Delete a webhook channel

Delete a destination. Other channel UUIDs remain unchanged and the deleted UUID is not reassigned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/delete-a-webhook-channel) · Effect: **consequential**

```sh
sellapp webhook-channels delete 0f33d01f-f9f8-45e8-80c8-7734d057196d --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| webhookChannel (positional) | Yes | The stable webhook channel UUID. |

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

Inspect the complete schema: `sellapp commands webhook-channels delete --json`.

## get

Retrieve a webhook channel

Retrieve one outbound webhook destination by its stable UUID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/retrieve-a-webhook-channel) · Effect: **read**

```sh
sellapp webhook-channels get 0f33d01f-f9f8-45e8-80c8-7734d057196d
```

| Input | Required | Meaning |
| --- | --- | --- |
| webhookChannel (positional) | Yes | The stable webhook channel UUID. |

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

Inspect the complete schema: `sellapp commands webhook-channels get --json`.

## list

List webhook channels

Retrieve the authenticated store's outbound developer webhook destinations. Signing secrets and delivery headers are never included. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/list-webhook-channels) · Effect: **read**

```sh
sellapp webhook-channels list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --page | No | Page number. |
| --limit | No | Number of webhook channels to return per page. |

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

Inspect the complete schema: `sellapp commands webhook-channels list --json`.

## replace

Replace a webhook channel

Replace a destination with a complete representation. name, url, and allowed_notifications must all be present. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/replace-a-webhook-channel) · Effect: **consequential**

```sh
sellapp webhook-channels replace 0f33d01f-f9f8-45e8-80c8-7734d057196d --body '{"name":"Primary Ship It webhook","url":"https://example.com/webhooks/ship-it","allowed_notifications":["order.paid"]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| webhookChannel (positional) | Yes | The stable webhook channel UUID. |
| --name | Yes | Optional seller-facing reference name. |
| --url | Yes | Public HTTP or HTTPS destination. The same canonical URL cannot be configured twice for one store. |
| --allowed-notifications | Yes | Complete set of public webhook event names enabled for this destination. |

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

Inspect the complete schema: `sellapp commands webhook-channels replace --json`.

## rotate

Rotate the webhook signing secret

Replace the store-wide outbound webhook signing secret. The value is write-only and is never returned or written to delivery logs. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/rotate-the-signing-secret) · Effect: **consequential**

```sh
sellapp webhook-channels rotate --signing-secret replace-with-a-random-secret-at-least-32-characters-long --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --signing-secret | Yes | Value for signing secret. |

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

Inspect the complete schema: `sellapp commands webhook-channels rotate --json`.

## search

Search webhook channels

Search destinations by name or URL and optionally filter by a subscribed event. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/search-webhook-channels) · Effect: **read**

```sh
sellapp webhook-channels search --body '{"search":{"value":"orders"},"event":"order.paid"}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --search | No | Value for search. |
| --event | No | Public event name delivered in webhook payloads. SellApp maps it to the corresponding allowed_notifications filter. |
| --page | No | Value for page. |
| --limit | No | Value for limit. |

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

Inspect the complete schema: `sellapp commands webhook-channels search --json`.

## send

Send a test webhook

Synchronously send one mock delivery to the stored destination. `event` is the public delivery event name and the same name must already be enabled in `allowed_notifications`. The test payload uses the same top-level event/version contract as production. Test deliveries do not retry or follow redirects and never return response bodies or signing headers. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/send-a-test-webhook) · Effect: **consequential**

```sh
sellapp webhook-channels send 0f33d01f-f9f8-45e8-80c8-7734d057196d --event order.created --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| webhookChannel (positional) | Yes | The stable webhook channel UUID. |
| --event | Yes | Public event name used by channel configuration, search, tests, and delivered webhook payloads. |

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

Inspect the complete schema: `sellapp commands webhook-channels send --json`.

## update

Update a webhook channel

Partially update a destination. Supplying allowed_notifications replaces the complete event filter; omitted fields retain their current values. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/webhook-channels/update-a-webhook-channel) · Effect: **consequential**

```sh
sellapp webhook-channels update 0f33d01f-f9f8-45e8-80c8-7734d057196d --body '{"name":"Primary Ship It webhook","allowed_notifications":["order.paid"]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| webhookChannel (positional) | Yes | The stable webhook channel UUID. |
| --name | No | Optional seller-facing reference name. |
| --url | No | Public HTTP or HTTPS destination. The same canonical URL cannot be configured twice for one store. |
| --allowed-notifications | No | Complete set of public webhook event names enabled for this destination. |

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

Inspect the complete schema: `sellapp commands webhook-channels update --json`.

