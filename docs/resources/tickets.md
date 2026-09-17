# tickets

[All commands](../commands.md)

## get

Retrieve specific ticket

Retrieve a support ticket by its ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp tickets get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets get --json`.

## list

List all tickets

List your store's support tickets, 15 per page by default.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp tickets list
```

| Input | Required | Meaning |
| --- | --- | --- |
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets list --json`.

## messages get

Retrieve specific ticket message

Retrieve one message using its ticket ID and message ID.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp tickets messages get --ticket 1 2
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ticket | Yes | The ticket path parameter. |
| message (positional) | Yes | The message path parameter. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets messages get --json`.

## messages list

List all ticket messages

Read a ticket's conversation, 15 messages per page by default.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp tickets messages list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets messages list --json`.

## messages reply

Reply to ticket

Add a seller reply to a support conversation.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp tickets messages reply 1 --content 'You can choose from the payment methods shown at checkout.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
| --content | Yes | Value for content. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets messages reply --json`.

## messages search

Search ticket messages

Search a ticket's messages using JSON body filters, search terms, includes, and sort instructions.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp tickets messages search 1 --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets messages search --json`.

## messages v-2-get-ticket-message

Retrieve specific ticket message

Retrieve one message using its ticket ID and message ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/retrieve-specific-ticket-message) · Effect: **read**

```sh
sellapp tickets messages v-2-get-ticket-message --ticket 1 2
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ticket | Yes | The ticket path parameter. |
| message (positional) | Yes | The message path parameter. |

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

Inspect the complete schema: `sellapp commands tickets messages v-2-get-ticket-message --json`.

## messages v-2-list-ticket-messages

List all ticket messages

Read a ticket's conversation, 15 messages per page by default. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/list-all-ticket-messages) · Effect: **read**

```sh
sellapp tickets messages v-2-list-ticket-messages 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
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

Inspect the complete schema: `sellapp commands tickets messages v-2-list-ticket-messages --json`.

## messages v-2-reply-to-ticket

Reply to ticket

Add a seller reply to a support conversation. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/reply-to-ticket) · Effect: **consequential**

```sh
sellapp tickets messages v-2-reply-to-ticket 1 --content 'You can choose from the payment methods shown at checkout.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
| --content | Yes | Value for content. |

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

Inspect the complete schema: `sellapp commands tickets messages v-2-reply-to-ticket --json`.

## messages v-2-search-ticket-messages

Search ticket messages

Search a ticket's messages using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/search-ticket-messages) · Effect: **read**

```sh
sellapp tickets messages v-2-search-ticket-messages 1 --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
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

Inspect the complete schema: `sellapp commands tickets messages v-2-search-ticket-messages --json`.

## search

Search tickets

Search tickets using JSON body filters, search terms, includes, and sort instructions.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp tickets search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands tickets search --json`.

## update

Close, reopen, or archive a ticket

Change ticket status or archive state in the selected store. Requires the ticket API-key ability and store ticket permission. Send status OPEN to reopen or CLOSED to close; closing sends the existing customer notification only when the status changes. Send archived true to archive or false to restore. You may send both fields together; omitted fields stay unchanged. Use the same optional Idempotency-Key and identical body after a lost response. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/retrieve-specific-ticket) · Effect: **consequential**

```sh
sellapp tickets update 42 --status CLOSED --archived true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | Value for ticket. |
| --status | No | Value for status. |
| --archived | No | Value for archived. |

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

Inspect the complete schema: `sellapp commands tickets update --json`.

## v-2-get-ticket

Retrieve specific ticket

Retrieve a support ticket by its ID. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/retrieve-specific-ticket) · Effect: **read**

```sh
sellapp tickets v-2-get-ticket 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |

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

Inspect the complete schema: `sellapp commands tickets v-2-get-ticket --json`.

## v-2-list-tickets

List all tickets

List your store's support tickets, 15 per page by default. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/list-all-tickets) · Effect: **read**

```sh
sellapp tickets v-2-list-tickets
```

| Input | Required | Meaning |
| --- | --- | --- |
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

Inspect the complete schema: `sellapp commands tickets v-2-list-tickets --json`.

## v-2-search-tickets

Search tickets

Search tickets using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/search-tickets) · Effect: **read**

```sh
sellapp tickets v-2-search-tickets --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
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

Inspect the complete schema: `sellapp commands tickets v-2-search-tickets --json`.

