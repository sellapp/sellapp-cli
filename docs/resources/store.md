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

## integrations disconnect-google-calendar

Disconnect Google Calendar

Disconnect the store's Google Calendar. Booking products that made Google Meet links with it switch to no video link; `meta.products_without_video_link` lists their titles. SellApp forgets the connection and does not contact Google, so remove SellApp's access in the Google account too if you want it revoked there. Send the `connection_id` you read from `GET /v2/store/settings/integrations`: a different ID, for example because the calendar was connected again since, returns `422` on `connection_id` and nothing changes. `404` when no calendar is connected, which is also what a repeated request gets. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/integrations) · Effect: **consequential**

```sh
sellapp store integrations disconnect-google-calendar --connection-id 7 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --connection-id | Yes | The `google_calendar.connection_id` you read. Required. It confirms you disconnect the connection you saw. |

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

Inspect the complete schema: `sellapp commands store integrations disconnect-google-calendar --json`.

## integrations get

Retrieve store integrations

Get the store's Google Calendar connection, which booking products use for availability and video links, and whether a Crisp live chat is set up for the storefront. No credentials, tokens or chat IDs are returned. Only SellApp's records are read; Google is not contacted. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/integrations) · Effect: **read**

```sh
sellapp store integrations get
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

Inspect the complete schema: `sellapp commands store integrations get --json`.

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

## sender-email delete

Delete the sender email

Remove the store's custom sender email: SellApp deletes the address at its email provider, then clears it and the sender name, so customer emails come from `default_from` again. To use the address again later, start over with a new verification email. Pass the sender you expect to remove as `email`; a different sender returns `422` on `email` and nothing is removed. `422` on `email` also when the store has no custom sender email, so a repeated request after a success is refused. `503` when the email provider fails: nothing was changed, try again later. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/sender-email) · Effect: **consequential**

```sh
sellapp store sender-email delete --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | No | The sender email you expect to remove. Protects against removing an address someone else set up after you read it. |

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

Inspect the complete schema: `sellapp commands store sender-email delete --json`.

## sender-email get

Retrieve the sender email

Get the address your customer emails come from: the store's custom sender email when it is set and verified, otherwise SellApp's default address. While you set up an address, pass it as `email` to see in `pending` whether its owner has opened the verification link yet. You only learn this for an address this store sent a verification email to in the last 7 days. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/sender-email) · Effect: **read**

```sh
sellapp store sender-email get
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | No | An address you are setting up as the sender email. When it differs from the current sender, the response reports it in `pending`. |

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

Inspect the complete schema: `sellapp commands store sender-email get --json`.

## sender-email send-verification

Send a sender email verification

Ask SellApp's email provider, Amazon SES, to email a verification link to the address: the first step of using your own address as the sender of customer emails. The store's sender does not change yet, so `email` stays `null` and customer emails keep coming from `default_from`. The address must be valid and its domain must accept email, and no other store may use it as its sender. The store must not have a sender email already: remove it first with `DELETE /v2/store/settings/sender-email`. Limits: one verification email per address and five per store in 24 hours; a refused request returns `422` on `email` with when the next one can be sent, and does not count. `422` on `email` too when another store is setting up the address. `503` when the email provider does not accept the request: nothing changed, try again later. Do not repeat a successful request: open the link in the email already sent. The example omits `email`, which is always `null` in this response. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/sender-email) · Effect: **consequential**

```sh
sellapp store sender-email send-verification --email 'orders@launch-lab.example.com' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | Yes | The address customer emails should come from. |

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

Inspect the complete schema: `sellapp commands store sender-email send-verification --json`.

## sender-email set

Set the sender email

Make a verified address the store's sender: customer emails then come from it, with the sender name from the personalization settings. Send the address you sent a verification email to with `POST /v2/store/settings/sender-email/verification` after its owner opened the link. Refused with `422` on `email` when the link was not opened yet, when this store did not send the address a verification email in the last 7 days, or when the store already has another sender (remove it first). Sending the current sender again returns `200` with the current state and changes nothing, so a retry is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/sender-email) · Effect: **consequential**

```sh
sellapp store sender-email set --email 'orders@launch-lab.example.com' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | Yes | The address customer emails should come from. |

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

Inspect the complete schema: `sellapp commands store sender-email set --json`.

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

## settings get-checkout

Retrieve checkout settings

Get your store's checkout settings: the shopping cart, whether checkout asks for an email, fraud controls, the minimum order total, the billing descriptor, the checkout rate limits, and the crypto checkout values when the store takes crypto through SellApp's crypto processor. The response body can be sent back to `PUT /v2/store/settings/checkout` unchanged. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/checkout-settings) · Effect: **read**

```sh
sellapp store settings get-checkout
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

Inspect the complete schema: `sellapp commands store settings get-checkout --json`.

## settings get-customer-emails

Retrieve customer email settings

Get which automatic emails your customers get: the delivery email after they pay, and the email asking them to pay when a subscription payment cannot be collected. The address these emails come from is managed with `/v2/store/settings/sender-email`. The response body can be sent back to `PUT /v2/store/settings/customer-emails` unchanged. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/customer-email-settings) · Effect: **read**

```sh
sellapp store settings get-customer-emails
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

Inspect the complete schema: `sellapp commands store settings get-customer-emails --json`.

## settings get-personalization

Retrieve personalization settings

Get the details customers see from your store: the support and Trustpilot addresses, the sender name and custom sender email of customer emails, the business details printed on invoices, and the Back button link of product and checkout pages. The response body can be sent back to `PUT /v2/store/settings/personalization` unchanged. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/personalization-settings) · Effect: **read**

```sh
sellapp store settings get-personalization
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

Inspect the complete schema: `sellapp commands store settings get-personalization --json`.

## settings get-store-alerts

Retrieve store alert settings

Get whether the store owner gets an email for each sale. Other alerts go to notification channels (`/v2/store/notification-channels`) and to the dashboard notifications (`/v2/notifications`). The response body can be sent back to `PUT /v2/store/settings/store-alerts` unchanged. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/store-alert-settings) · Effect: **read**

```sh
sellapp store settings get-store-alerts
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

Inspect the complete schema: `sellapp commands store settings get-store-alerts --json`.

## settings get-taxes

Retrieve tax settings

Get whether checkout collects VAT, whether your prices include it, and whether checkout asks for billing details. The response body can be sent back to `PUT /v2/store/settings/taxes` unchanged. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/tax-settings) · Effect: **read**

```sh
sellapp store settings get-taxes
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

Inspect the complete schema: `sellapp commands store settings get-taxes --json`.

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

## settings replace-checkout

Update checkout settings

Replace all eleven checkout settings. Every field except `crypto` is required; send `null` to remove a nullable one. A missing field returns `422` and nothing is saved. The body of `GET /v2/store/settings/checkout`, including `crypto`, can be sent back unchanged. Sending the same body again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/checkout-settings) · Effect: **consequential**

```sh
sellapp store settings replace-checkout --body '{"shopping_cart_enabled":true,"customer_email_required":true,"allow_vpn_purchases":true,"blacklist_on_dispute":true,"minimum_spend":"5.00","billing_descriptor":"LAUNCH LAB","attempts_per_customer":5,"attempts_per_ip":20,"attempts_per_email":null,"attempts_per_fingerprint":null,"rate_limit_window_seconds":3600,"crypto":{"underpaid_percentage":2,"minimum_confirmations":2}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --shopping-cart-enabled | Yes | Whether customers can add several products to a cart and buy them in one checkout. `false` turns the storefront cart off. |
| --customer-email-required | Yes | Whether checkout asks customers for an email address. With `false`, checkout does not ask, and delivery and payment emails cannot reach a customer who left none. |
| --allow-vpn-purchases | Yes | Whether customers who use a VPN or a proxy can buy. |
| --blacklist-on-dispute | Yes | Whether a customer's email address and IP address are added to your blacklist when they open a payment dispute. |
| --minimum-spend | Yes | The smallest order total, VAT included, that customers can check out with, as a decimal string in the order's currency: `"12.50"` on a USD order means $12.50. At least `1`; digits with an optional decimal part, no sign or exponent. Send `null` or `"0"` to remove the minimum; it then reads `null`. |
| --billing-descriptor | Yes | The purchase name customers see while they check out, up to 20 characters. `null` uses the default, `SellApp Order`. |
| --attempts-per-customer | Yes | How many checkout attempts one customer can make within the window, counted per IP address and email pair, or per checkout session fingerprint when checkout does not ask for an email. 1 to 100000; the default is `5`. This limit cannot be removed. |
| --attempts-per-ip | Yes | How many checkout attempts can come from one IP address within the window. 1 to 100000; `null` turns this limit off. |
| --attempts-per-email | Yes | How many checkout attempts can use one customer email address within the window. 1 to 100000; `null` turns this limit off. |
| --attempts-per-fingerprint | Yes | How many checkout attempts can come from one checkout session fingerprint (a browser identifier) within the window. 1 to 100000; `null` turns this limit off. |
| --rate-limit-window-seconds | Yes | How long each checkout limit counts attempts before allowing them again, in seconds: 60 (one minute) to 604800 (seven days). The default is `3600` (one hour). Attempts are counted per store. |
| --crypto | No | Value for crypto. |

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

Inspect the complete schema: `sellapp commands store settings replace-checkout --json`.

## settings replace-crypto-checkout

Update crypto checkout settings

Replace both crypto checkout values, for a store that takes crypto through SellApp's crypto processor; `underpaid_percentage` and `minimum_confirmations` are required. SellApp sends the changed values to the crypto processor and saves them only once it accepts them; an unchanged value is not sent. Returns the checkout settings with the new `crypto` values. `422` on `crypto` when the store does not take crypto through SellApp's crypto processor, or when the processor refuses the values; nothing is saved then. `422` on `underpaid_percentage` or `minimum_confirmations` for a value out of range. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/crypto-checkout-settings) · Effect: **consequential**

```sh
sellapp store settings replace-crypto-checkout --underpaid-percentage 1 --minimum-confirmations 3 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --underpaid-percentage | Yes | How far short of the amount due, in percent, a crypto payment can be and still count as paid: `2` accepts a payment up to 2% short. 0 to 100. |
| --minimum-confirmations | Yes | How many blockchain confirmations a crypto payment needs before the order is processed. 1 to 25. |

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

Inspect the complete schema: `sellapp commands store settings replace-crypto-checkout --json`.

## settings replace-customer-emails

Update customer email settings

Set both customer emails; both fields are required. A missing field returns `422` and nothing is saved. Sending the same body again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/customer-email-settings) · Effect: **consequential**

```sh
sellapp store settings replace-customer-emails --send-delivery-email true --send-payment-failure-email false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --send-delivery-email | Yes | Whether customers automatically get an email with what they bought after paying. Turn it off when you deliver to customers yourself. |
| --send-payment-failure-email | Yes | Whether customers get an email with a way to pay when your payment processor cannot collect a subscription payment. Turn it off when your processor already emails them. |

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

Inspect the complete schema: `sellapp commands store settings replace-customer-emails --json`.

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

## settings replace-personalization

Update personalization settings

Replace every personalization setting. All fields except `sender_email` are required; send `null` to remove one. A missing field returns `422` and nothing is saved. The body of `GET /v2/store/settings/personalization`, including `sender_email`, can be sent back unchanged. Sending the same body again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/personalization-settings) · Effect: **consequential**

```sh
sellapp store settings replace-personalization --support-email 'support@launch-lab.example.com' --trustpilot-email null --sender-name 'Launch Lab' --sender-email 'orders@launch-lab.example.com' --store-address '651 N Broad St, Suite 206' --store-city Middletown --store-state Delaware --store-zip 19709 --store-country 'United States' --store-tax 'EIN 98-7654321' --back-link-label 'Back to Launch Lab' --back-link-url https://launch-lab.example.com --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --support-email | Yes | The address customers can write to for support. |
| --trustpilot-email | Yes | Your Trustpilot invitation address. It receives a blind copy of every delivery email, including the link to what the customer bought, so Trustpilot can invite the customer to leave a review. |
| --sender-name | Yes | The name customer emails are sent from: letters, digits, spaces and `. - _ +`, up to 75 characters. It is used only while the store sends from its own verified sender email. A store without a custom sender email cannot keep one: sent with other changes it is ignored, and sent alone it returns `422` on `sender_name`. |
| --store-address | Yes | Street address printed on invoices. |
| --store-city | Yes | City printed on invoices. |
| --store-state | Yes | State or region printed on invoices. |
| --store-zip | Yes | Postal code printed on invoices. |
| --store-country | Yes | Country printed on invoices, as text. |
| --store-tax | Yes | Tax or VAT details printed on invoices, such as `VAT GB123456789`. |
| --back-link-label | Yes | Text of the Back button on product and checkout pages. A return link needs both a label and a URL: a change that leaves one of them empty returns `422` on the empty one. Send `null` for both to remove the link; the Back button then goes to the store's home page. |
| --back-link-url | Yes | Where the Back button on product and checkout pages goes, as a full URL. Needs `back_link_label`. |
| --sender-email | No | The store's custom sender email, or `null` when it has none. Read-only here: send it back unchanged or leave it out. A different value returns `422` on `sender_email`; set it up or remove it with `/v2/store/settings/sender-email`. |

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

Inspect the complete schema: `sellapp commands store settings replace-personalization --json`.

## settings replace-store-alerts

Update store alert settings

Set the sale emails to the store owner; `sale_emails_enabled` is required. Sending the same body again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/store-alert-settings) · Effect: **consequential**

```sh
sellapp store settings replace-store-alerts --sale-emails-enabled false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --sale-emails-enabled | Yes | Whether the store owner gets an email for each sale. |

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

Inspect the complete schema: `sellapp commands store settings replace-store-alerts --json`.

## settings replace-taxes

Update tax settings

Replace all three tax settings; every field is required. A missing field returns `422` and nothing is saved. `vat_enabled: true` while Stripe Tax is on returns `422` on `vat_enabled` and saves nothing: turn Stripe Tax off in the dashboard first. When Stripe Tax was turned on but Stripe is now turned off as a payment method, turning VAT on also turns Stripe Tax off. Sending the same body again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/tax-settings) · Effect: **consequential**

```sh
sellapp store settings replace-taxes --vat-enabled true --vat-inclusive-pricing true --collect-billing-details false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --vat-enabled | Yes | Whether checkout collects VAT from customers. VAT and Stripe Tax cannot both be on: turning VAT on while Stripe Tax is on returns `422` on `vat_enabled` with `Disable Stripe Tax before enabling Sell VAT.` and saves nothing. Turn Stripe Tax off in **Settings > Taxes** in the dashboard first. If Stripe Tax was turned on but Stripe is now turned off as a payment method, turning VAT on succeeds and turns Stripe Tax off as well. |
| --vat-inclusive-pricing | Yes | Whether your product prices already include VAT instead of VAT being added on top. Always `false` while VAT is off, and turning `vat_enabled` off turns it off too. `true` while VAT stays off returns `422` on `vat_inclusive_pricing`: send `"vat_enabled": true` in the same request. |
| --collect-billing-details | Yes | Whether checkout asks customers for their billing details. |

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

Inspect the complete schema: `sellapp commands store settings replace-taxes --json`.

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

## settings update-checkout

Update checkout settings

Change only the checkout settings you send and keep the others, including one someone else saved while your request was in flight. Send `null` to remove `minimum_spend`, `billing_descriptor`, `attempts_per_ip`, `attempts_per_email` or `attempts_per_fingerprint`. `crypto` is read-only here; change it with `PATCH /v2/store/settings/checkout/crypto`. A refused value returns `422` on the field name of this endpoint, such as `rate_limit_window_seconds`, and nothing is saved. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/checkout-settings) · Effect: **consequential**

```sh
sellapp store settings update-checkout --minimum-spend 10.00 --attempts-per-email 10 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --shopping-cart-enabled | No | Whether customers can add several products to a cart and buy them in one checkout. `false` turns the storefront cart off. |
| --customer-email-required | No | Whether checkout asks customers for an email address. With `false`, checkout does not ask, and delivery and payment emails cannot reach a customer who left none. |
| --allow-vpn-purchases | No | Whether customers who use a VPN or a proxy can buy. |
| --blacklist-on-dispute | No | Whether a customer's email address and IP address are added to your blacklist when they open a payment dispute. |
| --minimum-spend | No | The smallest order total, VAT included, that customers can check out with, as a decimal string in the order's currency: `"12.50"` on a USD order means $12.50. At least `1`; digits with an optional decimal part, no sign or exponent. Send `null` or `"0"` to remove the minimum; it then reads `null`. |
| --billing-descriptor | No | The purchase name customers see while they check out, up to 20 characters. `null` uses the default, `SellApp Order`. |
| --attempts-per-customer | No | How many checkout attempts one customer can make within the window, counted per IP address and email pair, or per checkout session fingerprint when checkout does not ask for an email. 1 to 100000; the default is `5`. This limit cannot be removed. |
| --attempts-per-ip | No | How many checkout attempts can come from one IP address within the window. 1 to 100000; `null` turns this limit off. |
| --attempts-per-email | No | How many checkout attempts can use one customer email address within the window. 1 to 100000; `null` turns this limit off. |
| --attempts-per-fingerprint | No | How many checkout attempts can come from one checkout session fingerprint (a browser identifier) within the window. 1 to 100000; `null` turns this limit off. |
| --rate-limit-window-seconds | No | How long each checkout limit counts attempts before allowing them again, in seconds: 60 (one minute) to 604800 (seven days). The default is `3600` (one hour). Attempts are counted per store. |
| --crypto | No | Value for crypto. |

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

Inspect the complete schema: `sellapp commands store settings update-checkout --json`.

## settings update-crypto-checkout

Update crypto checkout settings

Change the crypto checkout values you send, for a store that takes crypto through SellApp's crypto processor. SellApp sends the changed values to the crypto processor and saves them only once it accepts them; an unchanged value is not sent. Returns the checkout settings with the new `crypto` values. `422` on `crypto` when the store does not take crypto through SellApp's crypto processor, or when the processor refuses the values; nothing is saved then. `422` on `underpaid_percentage` or `minimum_confirmations` for a value out of range. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/crypto-checkout-settings) · Effect: **consequential**

```sh
sellapp store settings update-crypto-checkout --minimum-confirmations 3 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --underpaid-percentage | No | How far short of the amount due, in percent, a crypto payment can be and still count as paid: `2` accepts a payment up to 2% short. 0 to 100. |
| --minimum-confirmations | No | How many blockchain confirmations a crypto payment needs before the order is processed. 1 to 25. |

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

Inspect the complete schema: `sellapp commands store settings update-crypto-checkout --json`.

## settings update-customer-emails

Update customer email settings

Turn the customer emails you send on or off and keep the other one. Takes effect for emails sent after the change. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/customer-email-settings) · Effect: **consequential**

```sh
sellapp store settings update-customer-emails --send-payment-failure-email false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --send-delivery-email | No | Whether customers automatically get an email with what they bought after paying. Turn it off when you deliver to customers yourself. |
| --send-payment-failure-email | No | Whether customers get an email with a way to pay when your payment processor cannot collect a subscription payment. Turn it off when your processor already emails them. |

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

Inspect the complete schema: `sellapp commands store settings update-customer-emails --json`.

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

## settings update-personalization

Update personalization settings

Change only the personalization settings you send and keep the others. Send `null` to remove one. A request with no field this endpoint knows returns `200` with the current settings and saves nothing. A return link needs both `back_link_label` and `back_link_url`: a change that leaves one empty returns `422` on that field. `sender_email` is read-only here. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/personalization-settings) · Effect: **consequential**

```sh
sellapp store settings update-personalization --back-link-label 'Return to the shop' --back-link-url https://launch-lab.example.com/shop --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --support-email | No | The address customers can write to for support. |
| --trustpilot-email | No | Your Trustpilot invitation address. It receives a blind copy of every delivery email, including the link to what the customer bought, so Trustpilot can invite the customer to leave a review. |
| --sender-name | No | The name customer emails are sent from: letters, digits, spaces and `. - _ +`, up to 75 characters. It is used only while the store sends from its own verified sender email. A store without a custom sender email cannot keep one: sent with other changes it is ignored, and sent alone it returns `422` on `sender_name`. |
| --sender-email | No | The store's custom sender email, or `null` when it has none. Read-only here: send it back unchanged or leave it out. A different value returns `422` on `sender_email`; set it up or remove it with `/v2/store/settings/sender-email`. |
| --store-address | No | Street address printed on invoices. |
| --store-city | No | City printed on invoices. |
| --store-state | No | State or region printed on invoices. |
| --store-zip | No | Postal code printed on invoices. |
| --store-country | No | Country printed on invoices, as text. |
| --store-tax | No | Tax or VAT details printed on invoices, such as `VAT GB123456789`. |
| --back-link-label | No | Text of the Back button on product and checkout pages. A return link needs both a label and a URL: a change that leaves one of them empty returns `422` on the empty one. Send `null` for both to remove the link; the Back button then goes to the store's home page. |
| --back-link-url | No | Where the Back button on product and checkout pages goes, as a full URL. Needs `back_link_label`. |

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

Inspect the complete schema: `sellapp commands store settings update-personalization --json`.

## settings update-store-alerts

Update store alert settings

Turn the sale emails to the store owner on or off. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/store-alert-settings) · Effect: **consequential**

```sh
sellapp store settings update-store-alerts --sale-emails-enabled false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --sale-emails-enabled | No | Whether the store owner gets an email for each sale. |

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

Inspect the complete schema: `sellapp commands store settings update-store-alerts --json`.

## settings update-taxes

Update tax settings

Change only the tax settings you send and keep the others. Turning VAT off also turns VAT-inclusive pricing off. Asking for VAT-inclusive pricing while VAT stays off returns `422` on `vat_inclusive_pricing` and saves nothing. Turning VAT on while Stripe Tax is on returns `422` on `vat_enabled` and saves nothing: turn Stripe Tax off in the dashboard first. When Stripe Tax was turned on but Stripe is now turned off as a payment method, turning VAT on also turns Stripe Tax off. Sending the same change again is safe. Requires the `store` API-key ability and the `store` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/store-settings/tax-settings) · Effect: **consequential**

```sh
sellapp store settings update-taxes --vat-inclusive-pricing true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --vat-enabled | No | Whether checkout collects VAT from customers. VAT and Stripe Tax cannot both be on: turning VAT on while Stripe Tax is on returns `422` on `vat_enabled` with `Disable Stripe Tax before enabling Sell VAT.` and saves nothing. Turn Stripe Tax off in **Settings > Taxes** in the dashboard first. If Stripe Tax was turned on but Stripe is now turned off as a payment method, turning VAT on succeeds and turns Stripe Tax off as well. |
| --vat-inclusive-pricing | No | Whether your product prices already include VAT instead of VAT being added on top. Always `false` while VAT is off, and turning `vat_enabled` off turns it off too. `true` while VAT stays off returns `422` on `vat_inclusive_pricing`: send `"vat_enabled": true` in the same request. |
| --collect-billing-details | No | Whether checkout asks customers for their billing details. |

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

Inspect the complete schema: `sellapp commands store settings update-taxes --json`.

