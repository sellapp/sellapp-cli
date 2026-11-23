# community-connections

[All commands](../commands.md)

## complete

Complete a community connection

Complete an authenticated WhatsApp QR connection by selecting a group returned by the status operation. The encrypted status token binds the session to the authenticated store and initiating user, and the server ID is revalidated against the provider response. Requires the `community` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections) · Effect: **consequential**

```sh
sellapp community-connections complete whatsapp --status-token replace-with-token-from-connection-start --server-id replace-with-returned-server-id --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| platform (positional) | Yes | The community platform. |
| --status-token | Yes | Value for status token. |
| --server-id | Yes | Value for server id. |

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

Inspect the complete schema: `sellapp commands community-connections complete --json`.

## disconnect

Disconnect a community platform

Idempotently disconnect the selected platform for the authenticated store and remove only that store's associated servers, channels, groups, or account links. Requires the `community` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections) · Effect: **consequential**

```sh
sellapp community-connections disconnect discord --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| platform (positional) | Yes | The community platform. |

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

Inspect the complete schema: `sellapp commands community-connections disconnect --json`.

## list

List community connections

List Discord, Telegram, Slack, and WhatsApp connection and persisted health status for the authenticated store without contacting providers. Credentials and provider secrets are never returned. Requires the `community` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections) · Effect: **read**

```sh
sellapp community-connections list
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

Inspect the complete schema: `sellapp commands community-connections list --json`.

## list-failed-grant-reasons

List community failed grant reasons

List why community access grants failed across all platforms, with how many grants failed for each reason and what you can do about it, most frequent first. A grant is the access a purchase gives a customer, such as a Discord role or a place in a Telegram group. Only reasons that still need attention or are retried are counted; failures nobody can act on, such as a customer who left the server, are not. Only SellApp's records are read; no platform is contacted. An empty `data` array means no grant is failing. Requires the `community` API-key ability and the `community` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/list-failed-grant-reasons) · Effect: **read**

```sh
sellapp community-connections list-failed-grant-reasons
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

Inspect the complete schema: `sellapp commands community-connections list-failed-grant-reasons --json`.

## poll

Poll a community connection

Poll a pending Telegram verification or WhatsApp QR connection using the encrypted status token returned by the connect operation. Tokens are bound to the authenticated store and initiating user. Requires the `community` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections) · Effect: **read**

```sh
sellapp community-connections poll discord --status-token test_status_token
```

| Input | Required | Meaning |
| --- | --- | --- |
| platform (positional) | Yes | The community platform. |
| --status-token | Yes | Value for status token. |

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

Inspect the complete schema: `sellapp commands community-connections poll --json`.

## servers get

Retrieve a community server

Retrieve one Discord server, Slack channel, Telegram group or channel, or WhatsApp group connected to the store, with the product variants that give access through it. Read it before removing it: the removal needs exactly these variant IDs. `is_last` tells you it is the platform's last one. `404` when the store has no such server on that platform. Requires the `community` API-key ability and both the `store` and `community` permissions in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/community-server) · Effect: **read**

```sh
sellapp community-connections servers get --platform discord server_ship_room
```

| Input | Required | Meaning |
| --- | --- | --- |
| --platform | Yes | The community platform. |
| server (positional) | Yes | The platform's ID of the Discord server, Slack channel, Telegram group or channel, or WhatsApp group, as returned in `servers[].id` by `GET /v2/community-connections`. |

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

Inspect the complete schema: `sellapp commands community-connections servers get --json`.

## servers remove

Remove a community server

Remove one Discord server, Slack channel, Telegram group or channel, or WhatsApp group from the store, and clear it from every product variant that gives access through it, including deleted ones. Buyers of those variants get no access through it from then on. SellApp does not contact the platform: members, roles and the bot stay as they are there. Removing the last Telegram or WhatsApp group also disconnects that platform. Confirm the variants first: read them with `GET` on the same path and send their IDs as `affected_variant_ids` (in any order). When they differ from the variants that use it now, the request returns `422` on `affected_variant_ids` and nothing is removed. Leave `affected_variant_ids` out, or send an empty list, when no variant uses it. Clients that cannot send a body with `DELETE` can send the IDs in the query string instead, as `affected_variant_ids[]=902`. A repeated request returns `404` once it is removed. Requires the `community` API-key ability and both the `store` and `community` permissions in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/community-server) · Effect: **consequential**

```sh
sellapp community-connections servers remove --platform discord server_ship_room --body '{"affected_variant_ids":[902]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --platform | Yes | The community platform. |
| server (positional) | Yes | The platform's ID of the Discord server, Slack channel, Telegram group or channel, or WhatsApp group, as returned in `servers[].id` by `GET /v2/community-connections`. |
| --affected-variant-ids | No | The IDs of every variant that gives access through it, from `affected_variants[].variant_id` of the read. Defaults to none. |

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

Inspect the complete schema: `sellapp commands community-connections servers remove --json`.

## slack add-channel

Add a Slack channel

Add a private Slack channel to the store so product variants can give customers access to it. The workspace needs an invite link first, so customers can join it: set it with `PUT /v2/community-connections/slack/invite-link`. SellApp checks with Slack that its app is a member of the channel. Adding a channel the store already has updates its name and returns `201` again, so a retry is safe. `422` on `platform` when Slack is not connected, on `workspace_invite_link` when the invite link is missing, and on `channel_id` when the app cannot see the channel. Requires the `community` API-key ability and the `community` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/slack-channels) · Effect: **consequential**

```sh
sellapp community-connections slack add-channel --channel-id C07LAUNCHVIP --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --channel-id | Yes | The Slack channel ID, from `channels[].id` of `GET /v2/community-connections/slack/channels`. |

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

Inspect the complete schema: `sellapp commands community-connections slack add-channel --json`.

## slack list-channels

List Slack available channels

List the private channels of the connected Slack workspace that SellApp's Slack app is a member of and that the store has not added yet. Add one with `POST` on the same path so products can give access to it. A channel is missing when the app is not in it: invite the app to the channel in Slack first. SellApp asks Slack for the list; when Slack does not answer, `channels` is empty. `422` on `platform` when Slack is not connected. Requires the `community` API-key ability and the `community` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/slack-channels) · Effect: **read**

```sh
sellapp community-connections slack list-channels
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

Inspect the complete schema: `sellapp commands community-connections slack list-channels --json`.

## slack replace-invite-link

Update Slack workspace invite link

Set or remove the invite link customers use to join the store's Slack workspace before they get access to a channel. Adding channels needs it. Send it as `invite_link`, or as `workspace_invite_link`, the name the response uses, so a response can be sent back unchanged; when you send both they must be equal. A link whose host is not `join.slack.com` returns `422` on the field you sent. `422` on `platform` when Slack is not connected. Sending the same link again is safe. Requires the `community` API-key ability and the `community` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/slack-invite-link) · Effect: **consequential**

```sh
sellapp community-connections slack replace-invite-link --invite-link https://join.slack.com/t/launch-lab/shared_invite/zt-2abc123 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --invite-link | No | A Slack invite link on `join.slack.com`, such as `https://join.slack.com/t/launch-lab/shared_invite/zt-2abc123`. `null` removes it. |
| --workspace-invite-link | No | Another name for `invite_link`. One of the two is required. |

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

Inspect the complete schema: `sellapp commands community-connections slack replace-invite-link --json`.

## start

Start a community connection

Start the provider-specific connection flow. OAuth providers return a connect URL containing encrypted state bound to the authenticated actor, store, platform, and purpose; that state expires after 15 minutes and securely authorizes the provider callback without forwarding the API bearer token or requiring a SellApp dashboard session. Open the URL exactly as issued, then poll the connection list with the original bearer token after the provider redirects. Do not append credentials, reconstruct, or reuse the URL. Legacy unsigned Slack URLs are rejected, so clients must call this endpoint again to restart those flows. Telegram returns a short-lived verification token, and WhatsApp returns a pollable QR session. The provider application must be configured, and Discord server connections require a connected Discord account. Requires the `community` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections) · Effect: **consequential**

```sh
sellapp community-connections start discord --mode official_bot --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| platform (positional) | Yes | The community platform. |
| --mode | No | Discord-only connection mode. |

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

Inspect the complete schema: `sellapp commands community-connections start --json`.

## sync-discord-roles

Sync Discord roles

Read the role list of the store's Discord servers from Discord again and check the bot's role position, which tells whether the bot can give the roles your products sell. Syncs every connected server, or only `server_id`. A server Discord refuses is reported on its own row with `synced: false`, and the others still sync; `failed_count` counts those rows. Nothing is given to or taken from a member, so repeating the request is safe. `422` on `platform` when no Discord server is connected, and on `server_id` when that server is not connected to this store. Requires the `community` API-key ability and the `community` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections/sync-discord-roles) · Effect: **consequential**

```sh
sellapp community-connections sync-discord-roles --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --server-id | No | The Discord server to sync. Leave out or send `null` to sync every connected server. |

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

Inspect the complete schema: `sellapp commands community-connections sync-discord-roles --json`.

## verify

Verify a community connection

Recheck provider and server health, then safely requeue eligible failed community grants. Requires the `community` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/community-connections) · Effect: **consequential**

```sh
sellapp community-connections verify discord --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| platform (positional) | Yes | The community platform. |

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

Inspect the complete schema: `sellapp commands community-connections verify --json`.

