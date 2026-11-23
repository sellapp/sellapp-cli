# support-automation-rules

[All commands](../commands.md)

## create

Create a support automation rule

Create an automation rule, paused unless you send `is_active: true`. Once switched on, a rule answers new matching tickets without anyone on your team: it posts its template as a reply to the customer, can credit a refund to the customer's store wallet, and can close the ticket. Create it paused, test it with the dry run, then switch it on with the state endpoint. `name` and `response_template_id` are required. Conditions you leave out or send as `null` match every ticket. Amounts are integer USD cents (`1999` means $19.99), and `action_refund_percent_basis` is in hundredths of a percent (`5000` is 50%). A field the rule does not take returns `422` naming it. The new rule is checked after your existing rules. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Not safe to retry blindly: a repeated request creates a second rule, so list the rules first when a response is lost. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/create-a-support-automation-rule) · Effect: **consequential**

```sh
sellapp support-automation-rules create --body '{"name":"Expired download links","response_template_id":32,"match_reason_source":"store","match_reason_code":"download-link-expired","match_product_ids":[120],"match_max_order_total_cents":5000,"max_prior_support_refund_count":0,"action_refund_kind":"partial","action_refund_percent_basis":5000,"action_refund_cap_cents":1000,"action_close_outcome":"refunded_partial"}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Name your team sees in the rule list. |
| --response-template-id | Yes | The reply template the rule sends. It must be one of your active templates with `supports_auto_respond`; another returns `422` on this field. If the template is paused or taken off automation later, a rule that only replies stops answering and lets the next rule try, and a rule that also refunds or closes takes the tickets it matches but does nothing for them. |
| --match-reason-source | No | With `match_reason_code`, the ticket reason the rule answers: `platform` for a built-in reason, `store` for one of yours. Send both or neither; `null` on both answers any reason. A store reason must be active and offered for every product. Read the codes from the support reason options endpoint. |
| --match-reason-code | No | A built-in reason code such as `not_received`, or your reason's `slug`. An unknown or unusable reason returns `422` on this field. |
| --match-product-ids | No | Products whose tickets the rule answers; a ticket about another product, or without a purchased item, does not match. `[]` answers every product. At most 200 distinct IDs of your store's products; a repeated, unknown or deleted product returns `422` on `match_product_ids.{index}`, except a product the rule already names. |
| --match-customer-filter | No | `first_purchase` answers customers with at most one completed order within `history_window`, `repeat` answers customers with more. |
| --history-window | No | How far back the customer-history conditions look: the last 7, 30 or 90 days, or `all_time`, which is the default. |
| --match-min-order-total-cents | No | Lowest value of the total of the ticket's purchased item, in integer USD cents (`1999` means $19.99). A ticket without a purchased item does not match when either is set. |
| --match-max-order-total-cents | No | Highest value of the total of the ticket's purchased item, in integer USD cents (`1999` means $19.99). A ticket without a purchased item does not match when either is set. Lower than the minimum returns `422` on this field. |
| --min-completed-orders | No | Lowest value of how many completed orders the customer has with your store within `history_window`. |
| --max-completed-orders | No | Highest value of how many completed orders the customer has with your store within `history_window`. Lower than the minimum returns `422` on this field. |
| --min-lifetime-spend-cents | No | Lowest value of what the customer spent on completed orders within `history_window`, in integer USD cents. |
| --max-lifetime-spend-cents | No | Highest value of what the customer spent on completed orders within `history_window`, in integer USD cents. Lower than the minimum returns `422` on this field. |
| --min-prior-ticket-count | No | Lowest value of how many other tickets the customer opened within `history_window`. |
| --max-prior-ticket-count | No | Highest value of how many other tickets the customer opened within `history_window`. Lower than the minimum returns `422` on this field. |
| --min-prior-support-refund-count | No | Lowest value of how many support wallet refunds the customer received on other tickets within `history_window`, not counting reversed ones. |
| --max-prior-support-refund-count | No | Highest value of how many support wallet refunds the customer received on other tickets within `history_window`, not counting reversed ones. Lower than the minimum returns `422` on this field. |
| --min-prior-support-refund-cents | No | Lowest value of the total of those support wallet refunds, in integer USD cents. |
| --max-prior-support-refund-cents | No | Highest value of the total of those support wallet refunds, in integer USD cents. Lower than the minimum returns `422` on this field. |
| --customer-age-days-min | No | Lowest value of days since the customer's first completed order within `history_window`. |
| --customer-age-days-max | No | Highest value of days since the customer's first completed order within `history_window`. Lower than the minimum returns `422` on this field. |
| --action-refund-kind | No | The wallet refund the rule credits to the customer: `full` refunds what is left to refund on the ticket's item, `partial` a share of it set by `action_refund_percent_basis`, `fixed` the amount in `action_refund_fixed_cents`. `null` means the rule does not refund. |
| --action-refund-percent-basis | No | For a `partial` refund only: the share in hundredths of a percent, so `5000` is 50% and `750` is 7.5%. The refund is rounded down to the cent. Sent with another kind, it returns `422`. |
| --action-refund-fixed-cents | No | For a `fixed` refund only: the amount in integer USD cents (`500` means $5.00). Sent with another kind, it returns `422`. When the item has less left to refund, the refund fails and the rule does nothing for that ticket. |
| --action-refund-cap-cents | No | Most the rule refunds automatically for one order, in integer USD cents, counted across all tickets about that order. A refund is reduced to what is left under it; when nothing is left, the rule does nothing for that ticket. Needs `action_refund_kind`. |
| --action-close-outcome | No | When set, the rule closes the ticket with this outcome after replying. `null` leaves the ticket open. |
| --is-active | No | Leave it out, or send `false`, to create the rule paused, then switch it on with the state endpoint after testing. `true` switches it on at once; for a refunding rule that returns `422` on `is_active` while support automation is on and no daily refund cap is set. |
| --id | No | Read-only. Accepted so a read can be sent back. |
| --archived | No | Accepted only as `false`. |
| --currency | No | Read-only. Accepted only as `USD`. |
| --created-at | No | Read-only. Accepted so a read can be sent back. |
| --updated-at | No | Read-only and not checked. Send `expected_updated_at` to refuse a stale change. |

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

Inspect the complete schema: `sellapp commands support-automation-rules create --json`.

## dry-run

Dry run support automation rules

See what your automation would do with one of your tickets if it were opened now: which rule would answer it, the reply it would send, the refund in USD cents and the close outcome. Nothing is sent, refunded, closed or recorded. Only switched-on rules take part, in their current order, so a paused rule shows up only after you switch it on. With support automation off, the answer is `matched: false` with `skip_reason: auto_respond_disabled`. For a refund rule, the dry run checks that the ticket's order is paid (`PAID`, `PARTIAL` or `COMPLETED`) and was not already refunded; otherwise it answers `would_execute: false` with `skip_reason: refund_unavailable`. It does not check the daily refund cap, whether a refund would be more than nothing (no ticket item, nothing left to refund, or the per-order cap used up), or whether the item is a subscription: those are checked when the rule runs on a new ticket, and a refund that fails them makes the rule do nothing for that ticket. Another store's ticket returns `404`. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/dry-run-support-automation-rules) · Effect: **read**

```sh
sellapp support-automation-rules dry-run --ticket-id 4821
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ticket-id | Yes | The ticket to test the rules against. |

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

Inspect the complete schema: `sellapp commands support-automation-rules dry-run --json`.

## get

Retrieve a support automation rule

Retrieve one automation rule, archived ones included. Keep its `updated_at` to send as `expected_updated_at` when you change the rule or switch it on. Another store's rule returns `404`. Amounts are integer USD cents (`1999` means $19.99), and `action_refund_percent_basis` is in hundredths of a percent (`5000` is 50%). Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/retrieve-a-support-automation-rule) · Effect: **read**

```sh
sellapp support-automation-rules get 7
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportAutomationRule (positional) | Yes | The automation rule ID. |

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

Inspect the complete schema: `sellapp commands support-automation-rules get --json`.

## list

List support automation rules

List your store's automation rules in the order they are checked, lowest `priority` first and the newest first on a tie, 15 per page by default and at most 100 with `limit`; a larger `limit` returns `422`. Archived rules are left out unless you ask for them with `archive_state`. Amounts are integer USD cents (`1999` means $19.99), and `action_refund_percent_basis` is in hundredths of a percent (`5000` is 50%). Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/list-support-automation-rules) · Effect: **read**

```sh
sellapp support-automation-rules list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --active-state | No | `active` returns records that are switched on, `inactive` returns paused ones. Any other value returns `422` on `active_state`. |
| --archive-state | No | `current` returns rules that are not archived, `archived` only archived ones, `all` both. Any other value returns `422` on `archive_state`. |
| --search | No | Matches part of the rule name or of its reply template's label. `%` and `_` are matched as ordinary characters, not wildcards, and a whole number also matches the record with that `id`. |

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

Inspect the complete schema: `sellapp commands support-automation-rules list --json`.

## reorder

Reorder support automation rules

Set the order rules are checked in. For each new ticket, the first switched-on rule that matches answers it and no rule after it runs, so put narrow rules before broad ones. Send every rule that is not archived exactly once in `ordered_rule_ids`, first checked first; the response lists them with `priority` counted from `0`. A list that repeats, adds or leaves out a rule returns `422` on `ordered_rule_ids` and changes nothing: list the rules again and send the full list. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Repeating the same order is safe. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/reorder-support-automation-rules) · Effect: **consequential**

```sh
sellapp support-automation-rules reorder --body '{"ordered_rule_ids":[7]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ordered-rule-ids | Yes | Every rule ID of the store that is not archived, each once, in the new order. |

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

Inspect the complete schema: `sellapp commands support-automation-rules reorder --json`.

## replace

Update a support automation rule

Update an automation rule. Only the conditions and actions you send change; `PUT` and `PATCH` behave the same. A field sent as `null` keeps its value; name it in `clear` to remove it. `match_product_ids` replaces the whole list. A change to a switched-on rule applies to the next new ticket. Switching the rule on or off, archiving and the order have their own endpoints, so `is_active`, `archived` and `priority` are accepted only with the values the rule already has; a read can be sent back unchanged. A template or reason the rule already has is kept even if it was paused since; a new one has to be usable. Adding a refund to a switched-on rule while support automation is on and no daily refund cap is set returns `422` on `action_refund_kind`. An archived rule returns `422` on `supportAutomationRule`: restore it first. Amounts are integer USD cents (`1999` means $19.99), and `action_refund_percent_basis` is in hundredths of a percent (`5000` is 50%). Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Repeating the same change is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the rule again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/update-a-support-automation-rule) · Effect: **consequential**

```sh
sellapp support-automation-rules replace 7 --body '{"name":"Expired download links, first refund only","clear":["match_max_order_total_cents"],"expected_updated_at":"2026-07-03T09:20:00.000000Z"}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportAutomationRule (positional) | Yes | The automation rule ID. |
| --name | No | Name your team sees in the rule list. |
| --match-reason-source | No | With `match_reason_code`, the ticket reason the rule answers: `platform` for a built-in reason, `store` for one of yours. Send both or neither; `null` on both answers any reason. A store reason must be active and offered for every product. Read the codes from the support reason options endpoint. |
| --match-reason-code | No | A built-in reason code such as `not_received`, or your reason's `slug`. An unknown or unusable reason returns `422` on this field. |
| --match-product-ids | No | Products whose tickets the rule answers; a ticket about another product, or without a purchased item, does not match. `[]` answers every product. At most 200 distinct IDs of your store's products; a repeated, unknown or deleted product returns `422` on `match_product_ids.{index}`, except a product the rule already names. |
| --match-customer-filter | No | `first_purchase` answers customers with at most one completed order within `history_window`, `repeat` answers customers with more. |
| --history-window | No | How far back the customer-history conditions look: the last 7, 30 or 90 days, or `all_time`, which is the default. |
| --match-min-order-total-cents | No | Lowest value of the total of the ticket's purchased item, in integer USD cents (`1999` means $19.99). A ticket without a purchased item does not match when either is set. |
| --match-max-order-total-cents | No | Highest value of the total of the ticket's purchased item, in integer USD cents (`1999` means $19.99). A ticket without a purchased item does not match when either is set. Lower than the minimum returns `422` on this field. |
| --min-completed-orders | No | Lowest value of how many completed orders the customer has with your store within `history_window`. |
| --max-completed-orders | No | Highest value of how many completed orders the customer has with your store within `history_window`. Lower than the minimum returns `422` on this field. |
| --min-lifetime-spend-cents | No | Lowest value of what the customer spent on completed orders within `history_window`, in integer USD cents. |
| --max-lifetime-spend-cents | No | Highest value of what the customer spent on completed orders within `history_window`, in integer USD cents. Lower than the minimum returns `422` on this field. |
| --min-prior-ticket-count | No | Lowest value of how many other tickets the customer opened within `history_window`. |
| --max-prior-ticket-count | No | Highest value of how many other tickets the customer opened within `history_window`. Lower than the minimum returns `422` on this field. |
| --min-prior-support-refund-count | No | Lowest value of how many support wallet refunds the customer received on other tickets within `history_window`, not counting reversed ones. |
| --max-prior-support-refund-count | No | Highest value of how many support wallet refunds the customer received on other tickets within `history_window`, not counting reversed ones. Lower than the minimum returns `422` on this field. |
| --min-prior-support-refund-cents | No | Lowest value of the total of those support wallet refunds, in integer USD cents. |
| --max-prior-support-refund-cents | No | Highest value of the total of those support wallet refunds, in integer USD cents. Lower than the minimum returns `422` on this field. |
| --customer-age-days-min | No | Lowest value of days since the customer's first completed order within `history_window`. |
| --customer-age-days-max | No | Highest value of days since the customer's first completed order within `history_window`. Lower than the minimum returns `422` on this field. |
| --response-template-id | No | The reply template the rule sends. It must be one of your active templates with `supports_auto_respond`; another returns `422` on this field. If the template is paused or taken off automation later, a rule that only replies stops answering and lets the next rule try, and a rule that also refunds or closes takes the tickets it matches but does nothing for them. |
| --action-refund-kind | No | The wallet refund the rule credits to the customer: `full` refunds what is left to refund on the ticket's item, `partial` a share of it set by `action_refund_percent_basis`, `fixed` the amount in `action_refund_fixed_cents`. `null` means the rule does not refund. |
| --action-refund-percent-basis | No | For a `partial` refund only: the share in hundredths of a percent, so `5000` is 50% and `750` is 7.5%. The refund is rounded down to the cent. Sent with another kind, it returns `422`. |
| --action-refund-fixed-cents | No | For a `fixed` refund only: the amount in integer USD cents (`500` means $5.00). Sent with another kind, it returns `422`. When the item has less left to refund, the refund fails and the rule does nothing for that ticket. |
| --action-refund-cap-cents | No | Most the rule refunds automatically for one order, in integer USD cents, counted across all tickets about that order. A refund is reduced to what is left under it; when nothing is left, the rule does nothing for that ticket. Needs `action_refund_kind`. |
| --action-close-outcome | No | When set, the rule closes the ticket with this outcome after replying. `null` leaves the ticket open. |
| --is-active | No | Accepted only with the value the rule already has, so a read can be sent back. Switch the rule on or off with the state endpoint; another value returns `422` on `is_active`. |
| --clear | No | Fields to remove, since `null` keeps a value. `match_reason` removes the reason, `action_refund` removes the whole refund, `action_refund_cap_cents` only the per-order cap, `action_close_outcome` stops the rule closing tickets. Clearing a field you also change returns `422` on `clear`. |
| --expected-updated-at | No | The `updated_at` you read. When the record was saved by someone else since, the change is refused with `422` on `expected_updated_at` and nothing is saved, so you never overwrite an edit you have not seen. |
| --id | No | Read-only. Accepted so a read can be sent back. |
| --archived | No | Accepted only with the value the rule already has. Archive or restore the rule with the state endpoint; another value returns `422` on `archived`. |
| --priority | No | Accepted only with the value the rule already has. Change the order with the reorder endpoint; another value returns `422` on `priority`. |
| --currency | No | Read-only. Accepted only as `USD`. |
| --created-at | No | Read-only. Accepted so a read can be sent back. |
| --updated-at | No | Read-only and not checked. Send `expected_updated_at` to refuse a stale change. |

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

Inspect the complete schema: `sellapp commands support-automation-rules replace --json`.

## update

Update a support automation rule

Update an automation rule. Only the conditions and actions you send change; `PUT` and `PATCH` behave the same. A field sent as `null` keeps its value; name it in `clear` to remove it. `match_product_ids` replaces the whole list. A change to a switched-on rule applies to the next new ticket. Switching the rule on or off, archiving and the order have their own endpoints, so `is_active`, `archived` and `priority` are accepted only with the values the rule already has; a read can be sent back unchanged. A template or reason the rule already has is kept even if it was paused since; a new one has to be usable. Adding a refund to a switched-on rule while support automation is on and no daily refund cap is set returns `422` on `action_refund_kind`. An archived rule returns `422` on `supportAutomationRule`: restore it first. Amounts are integer USD cents (`1999` means $19.99), and `action_refund_percent_basis` is in hundredths of a percent (`5000` is 50%). Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Repeating the same change is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the rule again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/update-a-support-automation-rule) · Effect: **consequential**

```sh
sellapp support-automation-rules update 7 --body '{"name":"Expired download links, first refund only","clear":["match_max_order_total_cents"],"expected_updated_at":"2026-07-03T09:20:00.000000Z"}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportAutomationRule (positional) | Yes | The automation rule ID. |
| --name | No | Name your team sees in the rule list. |
| --match-reason-source | No | With `match_reason_code`, the ticket reason the rule answers: `platform` for a built-in reason, `store` for one of yours. Send both or neither; `null` on both answers any reason. A store reason must be active and offered for every product. Read the codes from the support reason options endpoint. |
| --match-reason-code | No | A built-in reason code such as `not_received`, or your reason's `slug`. An unknown or unusable reason returns `422` on this field. |
| --match-product-ids | No | Products whose tickets the rule answers; a ticket about another product, or without a purchased item, does not match. `[]` answers every product. At most 200 distinct IDs of your store's products; a repeated, unknown or deleted product returns `422` on `match_product_ids.{index}`, except a product the rule already names. |
| --match-customer-filter | No | `first_purchase` answers customers with at most one completed order within `history_window`, `repeat` answers customers with more. |
| --history-window | No | How far back the customer-history conditions look: the last 7, 30 or 90 days, or `all_time`, which is the default. |
| --match-min-order-total-cents | No | Lowest value of the total of the ticket's purchased item, in integer USD cents (`1999` means $19.99). A ticket without a purchased item does not match when either is set. |
| --match-max-order-total-cents | No | Highest value of the total of the ticket's purchased item, in integer USD cents (`1999` means $19.99). A ticket without a purchased item does not match when either is set. Lower than the minimum returns `422` on this field. |
| --min-completed-orders | No | Lowest value of how many completed orders the customer has with your store within `history_window`. |
| --max-completed-orders | No | Highest value of how many completed orders the customer has with your store within `history_window`. Lower than the minimum returns `422` on this field. |
| --min-lifetime-spend-cents | No | Lowest value of what the customer spent on completed orders within `history_window`, in integer USD cents. |
| --max-lifetime-spend-cents | No | Highest value of what the customer spent on completed orders within `history_window`, in integer USD cents. Lower than the minimum returns `422` on this field. |
| --min-prior-ticket-count | No | Lowest value of how many other tickets the customer opened within `history_window`. |
| --max-prior-ticket-count | No | Highest value of how many other tickets the customer opened within `history_window`. Lower than the minimum returns `422` on this field. |
| --min-prior-support-refund-count | No | Lowest value of how many support wallet refunds the customer received on other tickets within `history_window`, not counting reversed ones. |
| --max-prior-support-refund-count | No | Highest value of how many support wallet refunds the customer received on other tickets within `history_window`, not counting reversed ones. Lower than the minimum returns `422` on this field. |
| --min-prior-support-refund-cents | No | Lowest value of the total of those support wallet refunds, in integer USD cents. |
| --max-prior-support-refund-cents | No | Highest value of the total of those support wallet refunds, in integer USD cents. Lower than the minimum returns `422` on this field. |
| --customer-age-days-min | No | Lowest value of days since the customer's first completed order within `history_window`. |
| --customer-age-days-max | No | Highest value of days since the customer's first completed order within `history_window`. Lower than the minimum returns `422` on this field. |
| --response-template-id | No | The reply template the rule sends. It must be one of your active templates with `supports_auto_respond`; another returns `422` on this field. If the template is paused or taken off automation later, a rule that only replies stops answering and lets the next rule try, and a rule that also refunds or closes takes the tickets it matches but does nothing for them. |
| --action-refund-kind | No | The wallet refund the rule credits to the customer: `full` refunds what is left to refund on the ticket's item, `partial` a share of it set by `action_refund_percent_basis`, `fixed` the amount in `action_refund_fixed_cents`. `null` means the rule does not refund. |
| --action-refund-percent-basis | No | For a `partial` refund only: the share in hundredths of a percent, so `5000` is 50% and `750` is 7.5%. The refund is rounded down to the cent. Sent with another kind, it returns `422`. |
| --action-refund-fixed-cents | No | For a `fixed` refund only: the amount in integer USD cents (`500` means $5.00). Sent with another kind, it returns `422`. When the item has less left to refund, the refund fails and the rule does nothing for that ticket. |
| --action-refund-cap-cents | No | Most the rule refunds automatically for one order, in integer USD cents, counted across all tickets about that order. A refund is reduced to what is left under it; when nothing is left, the rule does nothing for that ticket. Needs `action_refund_kind`. |
| --action-close-outcome | No | When set, the rule closes the ticket with this outcome after replying. `null` leaves the ticket open. |
| --is-active | No | Accepted only with the value the rule already has, so a read can be sent back. Switch the rule on or off with the state endpoint; another value returns `422` on `is_active`. |
| --clear | No | Fields to remove, since `null` keeps a value. `match_reason` removes the reason, `action_refund` removes the whole refund, `action_refund_cap_cents` only the per-order cap, `action_close_outcome` stops the rule closing tickets. Clearing a field you also change returns `422` on `clear`. |
| --expected-updated-at | No | The `updated_at` you read. When the record was saved by someone else since, the change is refused with `422` on `expected_updated_at` and nothing is saved, so you never overwrite an edit you have not seen. |
| --id | No | Read-only. Accepted so a read can be sent back. |
| --archived | No | Accepted only with the value the rule already has. Archive or restore the rule with the state endpoint; another value returns `422` on `archived`. |
| --priority | No | Accepted only with the value the rule already has. Change the order with the reorder endpoint; another value returns `422` on `priority`. |
| --currency | No | Read-only. Accepted only as `USD`. |
| --created-at | No | Read-only. Accepted so a read can be sent back. |
| --updated-at | No | Read-only and not checked. Send `expected_updated_at` to refuse a stale change. |

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

Inspect the complete schema: `sellapp commands support-automation-rules update --json`.

## update-state

Update a support automation rule state

Switch a rule on (`active`) or off (`paused`), archives it (`archived`), or brings an archived rule back (`restored`) in the on or off state it had. A rule you switch on answers new matching tickets straight away, without anyone on your team, while support automation is on. A rule that refunds can be switched on or restored only with the `expected_updated_at` you read, so you have seen its refund (otherwise `422` on `expected_updated_at`), and, while automation is on, only when a daily refund cap is set (otherwise `422` on `state`). An archived rule accepts only `restored`; another state returns `422` on `state`. A rule already in the requested state is returned unchanged. Archived rules never run. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Repeating the request is safe, except with an old `expected_updated_at`, which no longer matches after the first change: read the rule again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-rules/change-a-support-automation-rule-state) · Effect: **consequential**

```sh
sellapp support-automation-rules update-state 1 --state active --expected-updated-at 2026-07-03T09:20:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| supportAutomationRule (positional) | Yes | The automation rule ID. |
| --state | Yes | The state to move the rule to. |
| --expected-updated-at | No | The `updated_at` you read. Required to switch on or restore a rule that refunds. When the rule changed since, nothing changes and `422` is returned on `expected_updated_at`. |

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

Inspect the complete schema: `sellapp commands support-automation-rules update-state --json`.

