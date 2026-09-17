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

