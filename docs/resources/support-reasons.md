# support-reasons

[All commands](../commands.md)

## create

Create a support reason

Add a ticket reason customers can choose when they open a ticket. `label` is required. SellApp makes the `slug` from the label, adding `-1`, `-2` and so on when it is taken, and never changes it afterwards. `is_active` defaults to `true`, and without `product_ids` the reason is offered for every product. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). Not safe to retry blindly: a repeated request creates a second reason with another slug, so list the reasons first when a response is lost. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/create-a-support-reason) · Effect: **consequential**

```sh
sellapp support-reasons create --label 'Download link expired' --description 'The download link in the order email no longer works.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --label | Yes | What customers see. It must contain at least one letter or number. |
| --description | No | Optional explanation shown under the label. |
| --is-active | No | `false` stops offering the reason to customers. Tickets already filed under it keep it. |
| --product-ids | No | Your store's products the reason is offered for; an empty list offers it for every product. At most 200 distinct IDs. A repeated, unknown or deleted product returns `422` on `product_ids.{index}`, except a product the reason already names, which stays accepted after it was deleted so a read can be sent back. |

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

Inspect the complete schema: `sellapp commands support-reasons create --json`.

## delete

Delete a support reason

Delete one of your ticket reasons and return it as it was. Customers can no longer choose it; tickets already filed under it keep their reason. Automation rules that match on the reason are not changed and do not stop the delete, so first look for rules whose `match_reason_code` is this reason's `slug`. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). Deleting cannot be undone; repeating it returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/delete-a-support-reason) · Effect: **consequential**

```sh
sellapp support-reasons delete 15 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReason (positional) | Yes | The support reason ID. |

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

Inspect the complete schema: `sellapp commands support-reasons delete --json`.

## get

Retrieve a support reason

Retrieve one of your ticket reasons. Another store's reason returns `404`. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/retrieve-a-support-reason) · Effect: **read**

```sh
sellapp support-reasons get 14
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReason (positional) | Yes | The support reason ID. |

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

Inspect the complete schema: `sellapp commands support-reasons get --json`.

## get-options

Retrieve support reason options

Get every reason a ticket can carry and what you need to build a reason editor: the seven built-in reasons (`platform_reasons`), your own reasons with their products (`custom_reasons`), the reasons offered to a customer before they pick a product (`customer_visible_reasons`: the built-in ones and your active reasons offered for every product), counts, defaults, the accepted list filters and the dashboard page. A reason's `source` and `code` are what an automation rule matches on. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/support-reason-options) · Effect: **read**

```sh
sellapp support-reasons get-options
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

Inspect the complete schema: `sellapp commands support-reasons get-options --json`.

## list

List support reasons

List your store's own ticket reasons, 15 per page by default and at most 100 with `limit`; a larger `limit` returns `422`. Customers choose a reason when they open a ticket, from SellApp's seven built-in reasons and the ones you add here. The built-in reasons are not in this list; read them from the support reason options endpoint. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/list-support-reasons) · Effect: **read**

```sh
sellapp support-reasons list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --active-state | No | `active` returns records that are switched on, `inactive` returns paused ones. Any other value returns `422` on `active_state`. |
| --product-scope | No | `all_products` returns reasons offered for every product, `selected_products` returns reasons limited to some products. Any other value returns `422` on `product_scope`. |
| --search | No | Matches part of the label, slug or description. `%` and `_` are matched as ordinary characters, not wildcards, and a whole number also matches the record with that `id`. |

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

Inspect the complete schema: `sellapp commands support-reasons list --json`.

## replace

Update a support reason

Update one of your ticket reasons. Only the fields you send change; `PUT` and `PATCH` behave the same. `product_ids` replaces the whole list, and `[]` offers the reason for every product. The `slug` stays the same when the label changes. A body with nothing to change returns `200` with the reason as it is and saves nothing. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). Repeating the same change is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the reason again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/update-a-support-reason) · Effect: **consequential**

```sh
sellapp support-reasons replace 14 --description 'The download link in the order email no longer opens.' --expected-updated-at 2026-07-03T09:10:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReason (positional) | Yes | The support reason ID. |
| --label | No | What customers see. It must contain at least one letter or number. |
| --description | No | Optional explanation shown under the label. |
| --is-active | No | `false` stops offering the reason to customers. Tickets already filed under it keep it. |
| --product-ids | No | Your store's products the reason is offered for; an empty list offers it for every product. At most 200 distinct IDs. A repeated, unknown or deleted product returns `422` on `product_ids.{index}`, except a product the reason already names, which stays accepted after it was deleted so a read can be sent back. |
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

Inspect the complete schema: `sellapp commands support-reasons replace --json`.

## update

Update a support reason

Update one of your ticket reasons. Only the fields you send change; `PUT` and `PATCH` behave the same. `product_ids` replaces the whole list, and `[]` offers the reason for every product. The `slug` stays the same when the label changes. A body with nothing to change returns `200` with the reason as it is and saves nothing. Requires the `ticket` API-key ability, and the key's user must be allowed to manage tickets in the store (the Resolution Center role, or owner or administrator). Repeating the same change is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the reason again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-reasons/update-a-support-reason) · Effect: **consequential**

```sh
sellapp support-reasons update 14 --description 'The download link in the order email no longer opens.' --expected-updated-at 2026-07-03T09:10:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportReason (positional) | Yes | The support reason ID. |
| --label | No | What customers see. It must contain at least one letter or number. |
| --description | No | Optional explanation shown under the label. |
| --is-active | No | `false` stops offering the reason to customers. Tickets already filed under it keep it. |
| --product-ids | No | Your store's products the reason is offered for; an empty list offers it for every product. At most 200 distinct IDs. A repeated, unknown or deleted product returns `422` on `product_ids.{index}`, except a product the reason already names, which stays accepted after it was deleted so a read can be sent back. |
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

Inspect the complete schema: `sellapp commands support-reasons update --json`.

