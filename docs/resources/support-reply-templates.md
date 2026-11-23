# support-reply-templates

[All commands](../commands.md)

## create

Create a support reply template

Create a reply template. Saving it sends nothing to anyone. `label` and `body` are required; `is_active` and `supports_manual` default to `true` and `supports_auto_respond` to `false`. Check a draft body with the reply template options endpoint before saving. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). Creating a template with `supports_auto_respond: true` also needs the `support.automation` ability and the Support Automation role, otherwise `403`. Not safe to retry blindly: a repeated request creates a second template, so list the templates first when a response is lost. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/create-a-support-reply-template) · Effect: **consequential**

```sh
sellapp support-reply-templates create --label 'Refund credited' --body-body 'Hi {customer_email}, we credited {refund_amount} to your {store_name} wallet for {product_title}. Your wallet balance is now {wallet_balance}.' --supports-auto-respond true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --label | Yes | Name your team sees when choosing the template. Surrounding spaces are trimmed. |
| --body-body | Yes | Reply text, trimmed. Only these placeholders are accepted: `{customer_email}`, `{ticket_id}`, `{order_id}`, `{product_title}`, `{refund_amount}`, `{wallet_balance}` and `{store_name}`. Any other placeholder returns `422` on `body` with `Unsupported placeholder: ...`. |
| --is-active | No | `false` pauses the template: it is no longer offered for replies, and automation rules that use it stop sending it. |
| --supports-manual | No | Whether your team can insert the template into replies they write. |
| --supports-auto-respond | No | Whether automation rules may send the template as an automatic reply. Needs the `support.automation` ability and the Support Automation role. |

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

Inspect the complete schema: `sellapp commands support-reply-templates create --json`.

## delete

Delete a support reply template

Delete a reply template and return it as it was. A template that a current automation rule uses is kept and returns `422` on `supportReplyTemplate`: point those rules at another template or archive them first. Archived rules that used the template are left without one. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). A template offered to automation rules (`supports_auto_respond`) or used by a rule also needs the `support.automation` ability and the Support Automation role; without them the change returns `403`. Deleting cannot be undone; repeating it returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/delete-a-support-reply-template) · Effect: **consequential**

```sh
sellapp support-reply-templates delete 33 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReplyTemplate (positional) | Yes | The reply template ID. |

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

Inspect the complete schema: `sellapp commands support-reply-templates delete --json`.

## get

Retrieve a support reply template

Retrieve one reply template, with `placeholders` listing the placeholders its body uses. Another store's template returns `404`. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/retrieve-a-support-reply-template) · Effect: **read**

```sh
sellapp support-reply-templates get 32
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReplyTemplate (positional) | Yes | The reply template ID. |

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

Inspect the complete schema: `sellapp commands support-reply-templates get --json`.

## get-options

Retrieve support reply template options

Get what you need to build a reply template editor: the defaults for a new template, whether the key's user may manage automation templates (`permissions.can_manage_automation`), template counts, the seven supported placeholders, the accepted list filters and the dashboard page. Send a draft `body`, and optionally more placeholder names in `requested_placeholders[]`, to have `placeholder_guidance` check them: it names each unsupported placeholder and a supported replacement where one exists, such as `{customer_email}` for `{customer_name}`. Nothing is saved. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/support-reply-template-options) · Effect: **read**

```sh
sellapp support-reply-templates get-options
```

| Input | Required | Meaning |
| --- | --- | --- |
| --body | No | A draft template body to check. |
| --requested-placeholders | No | More placeholder names to check, with or without braces. Repeat the parameter to send several. |

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

Inspect the complete schema: `sellapp commands support-reply-templates get-options --json`.

## list

List support reply templates

List your store's reply templates in the order your team sees them, 15 per page by default and at most 100 with `limit`; a larger `limit` returns `422`. A reply template is saved text with placeholders such as `{customer_email}` that your team inserts into ticket replies and that automation rules send as automatic replies. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/list-support-reply-templates) · Effect: **read**

```sh
sellapp support-reply-templates list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --active-state | No | `active` returns records that are switched on, `inactive` returns paused ones. Any other value returns `422` on `active_state`. |
| --availability | No | `manual` returns templates your team can insert into replies (`supports_manual`); `auto` returns templates automation rules may send (`supports_auto_respond`). Any other value returns `422` on `availability`. |
| --search | No | Matches part of the label or the body. `%` and `_` are matched as ordinary characters, not wildcards, and a whole number also matches the record with that `id`. |

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

Inspect the complete schema: `sellapp commands support-reply-templates list --json`.

## render

Render a support reply template

Fill a reply template for one of your tickets and get the text back, ready for you to send as a reply. Nothing is sent, saved or refunded. The ticket supplies the customer's email, the ticket and order numbers, the product title, the store name, the customer's wallet balance, and its latest wallet refund that is still standing as `refund_amount`. Type what the ticket does not hold in `values`. To write the note for a wallet refund you are about to issue, send `refund_amount_cents` (and `invoice_id` when the ticket's order has several items): `{refund_amount}` and `{wallet_balance}` then describe that refund. Only active templates with `supports_manual` can be filled; any other returns `422` on `supportReplyTemplate`. A missing template, or another store's template or ticket, returns `404`. Refusals: `values` and `values.{placeholder}` name the values still needed or invalid; `values.refund_amount` when a typed amount is not the ticket's wallet refund, or when the template shows a wallet balance for a ticket with no wallet refund; `invoice_id` when the item is not on the ticket's order or the order has several items and none was named; `refund_amount_cents` when the item has less than that left to refund; `ticket_id` when the ticket has no item left to refund. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). Safe to retry. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/render-a-support-reply-template) · Effect: **consequential**

```sh
sellapp support-reply-templates render 1 --ticket-id 4821 --invoice-id 7001 --refund-amount-cents 999 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReplyTemplate (positional) | Yes | The reply template ID. |
| --ticket-id | Yes | The ticket the reply is for. |
| --values | No | Value for values. |
| --invoice-id | No | The ticket's order item the refund note is about. Needed with `refund_amount_cents` when the order has several items. |
| --refund-amount-cents | No | A wallet refund you are about to issue, in integer USD cents (`999` means $9.99). It fills `{refund_amount}` and adds to `{wallet_balance}`, and cannot exceed what is left to refund on the item. |

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

Inspect the complete schema: `sellapp commands support-reply-templates render --json`.

## reorder

Reorder support reply templates

Set the order your team sees the templates in. Send every template of the store exactly once in `ordered_template_ids`, first one first; the response lists them all with `sort_order` counted from `0`. A list that repeats, adds or leaves out a template returns `422` on `ordered_template_ids` and changes nothing: list the templates again and send the full list. This order does not change which automation rule answers a ticket. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). When any template is offered to automation rules or used by one, reordering also needs the `support.automation` ability and the Support Automation role, otherwise `403`. Repeating the same order is safe. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/reorder-support-reply-templates) · Effect: **consequential**

```sh
sellapp support-reply-templates reorder --body '{"ordered_template_ids":[32,31]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ordered-template-ids | Yes | Every reply template ID of the store, each once, in the new order. |

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

Inspect the complete schema: `sellapp commands support-reply-templates reorder --json`.

## replace

Update a support reply template

Update a reply template. Only the fields you send change; `PUT` and `PATCH` behave the same. A body with nothing to change returns `200` with the template as it is and saves nothing. Pausing a template (`is_active: false`) or turning off `supports_auto_respond` takes it away from the automation rules that use it: a rule that only replies stops answering and lets the next rule try, while a rule that also refunds or closes still takes the tickets it matches but does nothing for them. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). A template offered to automation rules (`supports_auto_respond`) or used by a rule also needs the `support.automation` ability and the Support Automation role; without them the change returns `403`. Repeating the same change is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the template again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/update-a-support-reply-template) · Effect: **consequential**

```sh
sellapp support-reply-templates replace 31 --is-active false --expected-updated-at 2026-07-03T09:12:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReplyTemplate (positional) | Yes | The reply template ID. |
| --label | No | Name your team sees when choosing the template. Surrounding spaces are trimmed. |
| --body-body | No | Reply text, trimmed. Only these placeholders are accepted: `{customer_email}`, `{ticket_id}`, `{order_id}`, `{product_title}`, `{refund_amount}`, `{wallet_balance}` and `{store_name}`. Any other placeholder returns `422` on `body` with `Unsupported placeholder: ...`. |
| --is-active | No | `false` pauses the template: it is no longer offered for replies, and automation rules that use it stop sending it. |
| --supports-manual | No | Whether your team can insert the template into replies they write. |
| --supports-auto-respond | No | Whether automation rules may send the template as an automatic reply. Needs the `support.automation` ability and the Support Automation role. |
| --expected-updated-at | No | The `updated_at` you read. When the record was saved by someone else since, the change is refused with `422` on `expected_updated_at` and nothing is saved, so you never overwrite an edit you have not seen. |

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

Inspect the complete schema: `sellapp commands support-reply-templates replace --json`.

## update

Update a support reply template

Update a reply template. Only the fields you send change; `PUT` and `PATCH` behave the same. A body with nothing to change returns `200` with the template as it is and saves nothing. Pausing a template (`is_active: false`) or turning off `supports_auto_respond` takes it away from the automation rules that use it: a rule that only replies stops answering and lets the next rule try, while a rule that also refunds or closes still takes the tickets it matches but does nothing for them. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). A template offered to automation rules (`supports_auto_respond`) or used by a rule also needs the `support.automation` ability and the Support Automation role; without them the change returns `403`. Repeating the same change is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the template again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reply-templates/update-a-support-reply-template) · Effect: **consequential**

```sh
sellapp support-reply-templates update 31 --is-active false --expected-updated-at 2026-07-03T09:12:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReplyTemplate (positional) | Yes | The reply template ID. |
| --label | No | Name your team sees when choosing the template. Surrounding spaces are trimmed. |
| --body-body | No | Reply text, trimmed. Only these placeholders are accepted: `{customer_email}`, `{ticket_id}`, `{order_id}`, `{product_title}`, `{refund_amount}`, `{wallet_balance}` and `{store_name}`. Any other placeholder returns `422` on `body` with `Unsupported placeholder: ...`. |
| --is-active | No | `false` pauses the template: it is no longer offered for replies, and automation rules that use it stop sending it. |
| --supports-manual | No | Whether your team can insert the template into replies they write. |
| --supports-auto-respond | No | Whether automation rules may send the template as an automatic reply. Needs the `support.automation` ability and the Support Automation role. |
| --expected-updated-at | No | The `updated_at` you read. When the record was saved by someone else since, the change is refused with `422` on `expected_updated_at` and nothing is saved, so you never overwrite an edit you have not seen. |

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

Inspect the complete schema: `sellapp commands support-reply-templates update --json`.

