# store

[All commands](../commands.md)

## custom-domains connect

Connect a custom domain

Connect a hostname through SellApp's managed Cloudflare custom-hostname flow. IP addresses, wildcard domains, credentials, ports, queries, and fragments are rejected. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/connect-custom-domain) · Effect: **consequential**

```sh
sellapp store custom-domains connect --domain example.com --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --domain | Yes | The custom domain as a hostname or URL. |

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

Inspect the complete schema: `sellapp commands store custom-domains connect --json`.

## custom-domains disconnect

Disconnect a custom domain

Disconnect a custom domain. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/manage-custom-domain) · Effect: **consequential**

```sh
sellapp store custom-domains disconnect 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customDomain (positional) | Yes | The customDomain path parameter. |

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

Inspect the complete schema: `sellapp commands store custom-domains disconnect --json`.

## custom-domains get

Retrieve a custom domain

Retrieve a custom domain. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/manage-custom-domain) · Effect: **read**

```sh
sellapp store custom-domains get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| customDomain (positional) | Yes | The customDomain path parameter. |

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

Inspect the complete schema: `sellapp commands store custom-domains get --json`.

## custom-domains list

List custom domains

List the current store's custom-domain connection states and DNS verification records. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/list-custom-domains) · Effect: **read**

```sh
sellapp store custom-domains list
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

Inspect the complete schema: `sellapp commands store custom-domains list --json`.

## custom-domains refresh

Refresh custom domain status

Ask Cloudflare to re-run custom-hostname validation and persist the latest DNS and SSL state. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/refresh-custom-domain) · Effect: **consequential**

```sh
sellapp store custom-domains refresh 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customDomain (positional) | Yes | The customDomain path parameter. |

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

Inspect the complete schema: `sellapp commands store custom-domains refresh --json`.

## notification-channels create

Create a notification channel

Create an email or Discord notification destination. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/create-notification-channel) · Effect: **consequential**

```sh
sellapp store notification-channels create --body '{"channel":{"type":"email","email":"maya@example.com","allowed_notifications":[]}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --channel | Yes | Value for channel. |

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

Inspect the complete schema: `sellapp commands store notification-channels create --json`.

## notification-channels delete

Delete a notification channel

Manage one store-scoped notification channel by its identifier. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/manage-notification-channel) · Effect: **consequential**

```sh
sellapp store notification-channels delete test_notificationChannel --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| notificationChannel (positional) | Yes | The UUID of the notification channel. |

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

Inspect the complete schema: `sellapp commands store notification-channels delete --json`.

## notification-channels get

Retrieve a notification channel

Manage one store-scoped notification channel by its identifier. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/manage-notification-channel) · Effect: **read**

```sh
sellapp store notification-channels get test_notificationChannel
```

| Input | Required | Meaning |
| --- | --- | --- |
| notificationChannel (positional) | Yes | The UUID of the notification channel. |

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

Inspect the complete schema: `sellapp commands store notification-channels get --json`.

## notification-channels list

List notification channels

List dashboard notification destinations without exposing Discord webhook URLs. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/list-notification-channels) · Effect: **read**

```sh
sellapp store notification-channels list
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

Inspect the complete schema: `sellapp commands store notification-channels list --json`.

## notification-channels replace

Update a notification channel

Manage one store-scoped notification channel by its identifier. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/manage-notification-channel) · Effect: **consequential**

```sh
sellapp store notification-channels replace test_notificationChannel --body '{"channel":{"type":"email","email":"maya@example.com","allowed_notifications":[]}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| notificationChannel (positional) | Yes | The UUID of the notification channel. |
| --channel | Yes | Value for channel. |

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

Inspect the complete schema: `sellapp commands store notification-channels replace --json`.

## notification-channels update

Update a notification channel

Manage one store-scoped notification channel by its identifier. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/manage-notification-channel) · Effect: **consequential**

```sh
sellapp store notification-channels update test_notificationChannel --body '{"channel":{"allowed_notifications":[]}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| notificationChannel (positional) | Yes | The UUID of the notification channel. |
| --channel | Yes | Value for channel. |

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

Inspect the complete schema: `sellapp commands store notification-channels update --json`.

## settings get

Retrieve store settings

Retrieve the current store configuration and status flags. Secret analytics credentials are represented only by configured booleans. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/retrieve-store-settings) · Effect: **read**

```sh
sellapp store settings get
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

Inspect the complete schema: `sellapp commands store settings get --json`.

## settings replace-analytics

Update analytics settings

Update analytics settings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/update-analytics-settings) · Effect: **consequential**

```sh
sellapp store settings replace-analytics --ga-4-measurement-id null --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ga-4-measurement-id | No | Value for ga4 measurement id. |
| --meta-pixel-id | No | Value for meta pixel id. |
| --tiktok-pixel-id | No | Value for tiktok pixel id. |
| --ga-4-api-secret | No | Value for ga4 api secret. |
| --meta-access-token | No | Value for meta access token. |
| --tiktok-access-token | No | Value for tiktok access token. |

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

Inspect the complete schema: `sellapp commands store settings replace-analytics --json`.

## settings replace-general

Update general store settings

Update general store settings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/update-general-settings) · Effect: **consequential**

```sh
sellapp store settings replace-general --name 'Launch Lab' --visibility HIDDEN --timezone Europe/London --currency USD --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Value for name. |
| --visibility | Yes | Value for visibility. |
| --timezone | Yes | Value for timezone. |
| --currency | Yes | Value for currency. |
| --dark-mode | No | Value for dark mode. |

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

Inspect the complete schema: `sellapp commands store settings replace-general --json`.

## settings replace-marketing

Update marketing settings

Update marketing settings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/update-marketing-settings) · Effect: **consequential**

```sh
sellapp store settings replace-marketing --body '{"abandoned_cart":{"enabled":false}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --abandoned-cart | Yes | Value for abandoned cart. |

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

Inspect the complete schema: `sellapp commands store settings replace-marketing --json`.

## settings update-analytics

Update analytics settings

Update analytics settings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/update-analytics-settings) · Effect: **consequential**

```sh
sellapp store settings update-analytics --ga-4-measurement-id null --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --ga-4-measurement-id | No | Value for ga4 measurement id. |
| --meta-pixel-id | No | Value for meta pixel id. |
| --tiktok-pixel-id | No | Value for tiktok pixel id. |
| --ga-4-api-secret | No | Value for ga4 api secret. |
| --meta-access-token | No | Value for meta access token. |
| --tiktok-access-token | No | Value for tiktok access token. |

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

Inspect the complete schema: `sellapp commands store settings update-analytics --json`.

## settings update-general

Update general store settings

Update general store settings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/update-general-settings) · Effect: **consequential**

```sh
sellapp store settings update-general --name 'Launch Lab' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | No | Value for name. |
| --visibility | No | Value for visibility. |
| --timezone | No | Value for timezone. |
| --currency | No | Value for currency. |
| --dark-mode | No | Value for dark mode. |

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

Inspect the complete schema: `sellapp commands store settings update-general --json`.

## settings update-marketing

Update marketing settings

Update marketing settings. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/update-marketing-settings) · Effect: **consequential**

```sh
sellapp store settings update-marketing --body '{"abandoned_cart":{"enabled":false}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --abandoned-cart | Yes | Value for abandoned cart. |

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

Inspect the complete schema: `sellapp commands store settings update-marketing --json`.

