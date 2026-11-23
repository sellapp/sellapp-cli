# notifications

[All commands](../commands.md)

## list

List notifications

List the store's latest dashboard notifications (the dashboard's bell), newest first: completed orders, disputes, stock running out, subscriptions that need attention and similar events. At most 10 are returned, with no further pages. `meta.unseen_count` counts every unseen notification of the store. Notifications belong to the store, not to a team member: every member sees the same list. Requires the `invoice` or `store` API-key ability, with the same permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/notifications/list-notifications) · Effect: **read**

```sh
sellapp notifications list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --unseen-only | No | Return only notifications nobody has marked seen. Accepts `true`, `false`, `1` and `0`. The filter applies before `limit`. |
| --limit | No | How many notifications to return, 1 to 10. A larger value returns `422` on `limit`. |

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

Inspect the complete schema: `sellapp commands notifications list --json`.

## mark

Mark notifications seen

Mark dashboard notifications seen for the whole team. Send `notification_ids` to mark those; IDs that belong to another store or are already seen are left alone. Leave `notification_ids` out, or send `null`, to mark every unseen notification of the store. An empty list marks nothing. Returns the latest 10 notifications, `meta.marked_seen` (how many changed) and `meta.unseen_count`. Repeating a request is safe: already seen notifications are not counted again. Requires the `invoice` or `store` API-key ability, with the same permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/notifications/mark-notifications-seen) · Effect: **consequential**

```sh
sellapp notifications mark --body '{"notification_ids":[9182]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --notification-ids | No | Up to 50 notification IDs to mark seen. Leave out or send `null` to mark every unseen notification. |

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

Inspect the complete schema: `sellapp commands notifications mark --json`.

