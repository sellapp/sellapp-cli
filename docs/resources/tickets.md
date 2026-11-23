# tickets

[All commands](../commands.md)

## analytics get

Retrieve support analytics

Read your store's support report for the last 7, 30 or 90 days, counted back from now: tickets received and resolved, closures by people and by automation rules, wallet credit issued from support, each automation rule's results, and the latest resolution activity. Wallet credit counts only credit that was not reversed, in integer USD cents: `2500` means $25.00. Requires the `ticket` API-key ability and the ticket permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/retrieve-support-analytics) · Effect: **read**

```sh
sellapp tickets analytics get
```

| Input | Required | Meaning |
| --- | --- | --- |
| --range | No | How many days back from now the report covers: `7d`, `30d` (default) or `90d`. Any other value returns `422` on `range`. |
| --rule-limit | No | How many automation rules to return, most closures first. Default 10; a value outside 1 to 20 is brought into that range. |
| --activity-limit | No | How many recent activity entries to return. Default 10; a value outside 1 to 25 is brought into that range. |

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

Inspect the complete schema: `sellapp commands tickets analytics get --json`.

## batch-update

Update several tickets

Close, reopen, archive, restore or mark as unread up to 50 tickets in one request, and get one result per ticket in the order you sent them. Every ID must be a ticket of the selected store: one unknown ID returns `404` and changes nothing. Each ticket is then changed on its own, so one refusal does not undo the others. `close` needs an `outcome`, which is recorded on every ticket it closes, and emails each customer that the ticket was closed, at most once an hour per ticket. `mark_unread` marks the ticket unread for your store (`read_status.store` becomes `false`). Requires the `ticket` API-key ability and the ticket permission in your team role. Repeating the same request is safe: a ticket already in the requested state comes back as `skipped`. Send an optional `Idempotency-Key` with an identical body to get the original response back after a lost connection. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/update-several-tickets) · Effect: **consequential**

```sh
sellapp tickets batch-update --body '{"action":"close","outcome":"delivered","resources":[311,313]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --action | Yes | What to do with every listed ticket. `restore` takes a ticket out of the archive. |
| --resources | Yes | Between 1 and 50 different ticket IDs. |
| --outcome | No | Why the ticket was closed: `refunded`, `refunded_partial` (partly refunded), `delivered`, `no_action`, `duplicate` or `other`. Required for `close` and refused for every other action. |

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

Inspect the complete schema: `sellapp commands tickets batch-update --json`.

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

## get-context

Retrieve a ticket's customer and refund context

Read what you need beside a ticket's conversation in one call: the resolution, the customer's order history in this store, the linked order, what can still be refunded to the customer's wallet, and the latest 20 support activity entries, newest first. `refund_options` has the same items as the ticket's wallet refund options. The customer figures count the store's completed orders whose customer or order email is the ticket's email. Requires the `ticket` API-key ability and the ticket permission in your team role. The response includes the customer's email, order count and lifetime spend, so give a key with this ability only to people and tools that may see them. Amounts are integer USD cents: `1999` means $19.99. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/retrieve-ticket-context) · Effect: **read**

```sh
sellapp tickets get-context 312
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket ID in the selected store. |

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

Inspect the complete schema: `sellapp commands tickets get-context --json`.

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

Search a ticket's messages using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. Page through results with `limit` (1 to 100, default 15; a larger value returns `422`) and `page`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/search-ticket-messages) · Effect: **read**

```sh
sellapp tickets messages v-2-search-ticket-messages 1 --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket path parameter. |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
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

Close, reopen, archive or restore a ticket in the selected store. Send `status` `CLOSED` to close or `OPEN` to reopen, and `archived` `true` to archive or `false` to restore; you may send both, and omitted fields stay unchanged. When you close a ticket you can record why with `outcome`, which is accepted only together with `status` `CLOSED`. Closing emails the customer that the ticket was closed, at most once an hour per ticket: reopening and closing again within the hour sends no second email. A closed ticket keeps its outcome. Sending the outcome it already has changes nothing; a different outcome, or an outcome for a ticket that was closed without one, returns `422` on `outcome` and nothing in the request is applied, so reopen the ticket first. Requires the `ticket` API-key ability and the ticket permission in your team role. Use the same optional Idempotency-Key and identical body after a lost response. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/retrieve-specific-ticket) · Effect: **consequential**

```sh
sellapp tickets update 42 --status CLOSED --outcome delivered --archived true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | Value for ticket. |
| --status | No | Value for status. |
| --archived | No | Value for archived. |
| --outcome | No | Why the ticket was closed: `refunded`, `refunded_partial` (partly refunded), `delivered`, `no_action`, `duplicate` or `other`. Only with `status` `CLOSED`. |

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

List your store's support tickets, 15 per page by default. This v2 route is store-scoped. Narrow the results with `waiting_on`, `reason` and `customer_email`; they combine with each other and with any other filters. An unknown `waiting_on` or `reason`, or a `customer_email` longer than 255 characters, returns `422`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/list-all-tickets) · Effect: **read**

```sh
sellapp tickets v-2-list-tickets
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --waiting-on | No | `store`: the latest message is from the customer, so the ticket waits on you. `customer`: the latest message is yours. A ticket without messages matches neither. |
| --reason | No | A ticket reason by its label (as tickets show it in `reason`), its key such as `platform:refund_request`, or its code; case does not matter. A reason the store does not have returns `422` on `reason`, listing the store's reasons. |
| --customer-email | No | Part of the customer's email, case-insensitive. `%` and `_` match those characters literally. |

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

Search tickets using JSON body filters, search terms, includes, and sort instructions. This v2 route is store-scoped. Narrow the results with `waiting_on`, `reason` and `customer_email`; they combine with each other and with any other filters. An unknown `waiting_on` or `reason`, or a `customer_email` longer than 255 characters, returns `422`. Page through results with `limit` (1 to 100, default 15; a larger value returns `422`) and `page`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/search-tickets) · Effect: **read**

```sh
sellapp tickets v-2-search-tickets --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --search | No | Value for search. |
| --includes | No | Value for includes. |
| --waiting-on | No | `store`: the latest message is from the customer, so the ticket waits on you. `customer`: the latest message is yours. A ticket without messages matches neither. |
| --reason | No | A ticket reason by its label (as tickets show it in `reason`), its key such as `platform:refund_request`, or its code; case does not matter. A reason the store does not have returns `422` on `reason`, listing the store's reasons. |
| --customer-email | No | Part of the customer's email, case-insensitive. `%` and `_` match those characters literally. |

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

## wallet-refunds create

Refund a ticket to the customer's wallet

Credit the customer's store wallet for one item of the ticket's order and send `customer_note` to the customer. This moves money: the credit is store credit in USD that the customer can spend in your store, and the note is posted as your reply on the ticket, which emails it to the customer. The ticket stays open. Read the ticket's wallet refund options first, then send the item's `invoice_id`, the `amount_cents` to credit (at most its `refundable_cents`) and `expected_refundable_cents` equal to the `refundable_cents` you read. If what is left to refund changed since you read it, the request returns `422` on `expected_refundable_cents` and credits nothing: read the options again. Amounts are integer USD cents: `500` means $5.00. Requires both the `ticket` and `support.refund` API-key abilities, and the store member needs both the ticket and the support refund permissions. Send an `Idempotency-Key` header; requests without one return `400`. Retrying with the same key and identical body within 24 hours returns the original refund with `Idempotent-Replayed: true` and does not credit or email the customer again. Never retry under a new key: the first refund already lowered what is left, so the retry returns `422` on `expected_refundable_cents`, and a request with the same amount and note as an earlier refund of this item returns `422` on `customer_note`. Change the note if you mean to credit the customer a second time. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/refund-ticket-to-wallet) · Effect: **consequential**

```sh
sellapp tickets wallet-refunds create 312 --invoice-id 5120 --amount-cents 500 --expected-refundable-cents 1999 --customer-note 'We added $5.00 to your store wallet for the trouble with your download link.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket ID in the selected store. |
| --invoice-id | Yes | The item to refund, from the ticket's wallet refund options. It must be an item of the ticket's order that was not refunded. |
| --amount-cents | Yes | Credit to add, in USD cents: `500` means $5.00. At most the item's `refundable_cents`. |
| --expected-refundable-cents | Yes | The item's `refundable_cents` as you read it. It protects against refunding the same money twice when someone else refunded in between. |
| --customer-note | Yes | Message to the customer, at least 6 characters after trimming. It is posted on the ticket and emailed. A leftover placeholder such as `{refund_amount}` is refused. |

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

Inspect the complete schema: `sellapp commands tickets wallet-refunds create --json`.

## wallet-refunds list

List a ticket's wallet refunds

List every wallet refund issued from a ticket, oldest first, whether a person or an automation rule issued it, and whether it was reversed. A ticket without refunds returns an empty `data` array. The list is not paginated. Amounts are integer USD cents. Requires the `ticket` API-key ability and the ticket permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/list-ticket-wallet-refunds) · Effect: **read**

```sh
sellapp tickets wallet-refunds list 312
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket ID in the selected store. |

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

Inspect the complete schema: `sellapp commands tickets wallet-refunds list --json`.

## wallet-refunds list-options

List a ticket's wallet refund options

List the items of the ticket's order that you can refund to the customer's store wallet, oldest first, with what is left to refund for each. Read this before refunding: send an item's `invoice_id`, an amount up to its `refundable_cents`, and that `refundable_cents` as `expected_refundable_cents`. A ticket without an order returns an empty `data` array. The list is not paginated. Amounts are integer USD cents: `1999` means $19.99. Requires the `ticket` API-key ability and the ticket permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/list-ticket-wallet-refund-options) · Effect: **read**

```sh
sellapp tickets wallet-refunds list-options 312
```

| Input | Required | Meaning |
| --- | --- | --- |
| ticket (positional) | Yes | The ticket ID in the selected store. |

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

Inspect the complete schema: `sellapp commands tickets wallet-refunds list-options --json`.

## wallet-refunds reverse

Reverse a ticket's wallet refund

Take a wallet refund issued from this ticket back out of the customer's wallet, and return the refund with `reversed` `true` and its `reversal`. The reversal removes the full refunded amount from the wallet balance, records your `reason`, and makes that amount refundable again for the item. It does not message the customer. If the wallet no longer holds the full amount because the customer already spent some of it, the request returns `422` on `walletRefund` and nothing changes: there is no partial reversal. A refund that was already reversed also returns `422` on `walletRefund`, so the money is never taken back twice. An ID that is not a wallet refund of this ticket returns `404`. Amounts are integer USD cents. Requires both the `ticket` and `support.refund` API-key abilities, and the store member needs both the ticket and the support refund permissions. To recover a lost response, retry with the same optional `Idempotency-Key` and identical body, or list the ticket's wallet refunds. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/tickets/reverse-ticket-wallet-refund) · Effect: **consequential**

```sh
sellapp tickets wallet-refunds reverse --ticket 312 7001 --reason 'Issued on the wrong ticket.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ticket | Yes | The ticket ID in the selected store. |
| walletRefund (positional) | Yes | The wallet refund's `id`, from the ticket's wallet refunds. |
| --reason | Yes | Why the refund is reversed, 6 to 5000 characters. Kept with the reversal; not sent to the customer. |

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

Inspect the complete schema: `sellapp commands tickets wallet-refunds reverse --json`.

