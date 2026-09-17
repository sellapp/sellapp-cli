# affiliate-program

[All commands](../commands.md)

## get

Retrieve affiliate program configuration

Retrieve the authenticated store's affiliate-program enablement, commissions, tracking settings, payout preferences, and selected product variants. Requires the `affiliate` token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliate-program/retrieve-affiliate-program) · Effect: **read**

```sh
sellapp affiliate-program get
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

Inspect the complete schema: `sellapp commands affiliate-program get --json`.

## invite

Invite an affiliate

Create a pending affiliate invitation and email a single-use onboarding link. Duplicate invitations are rejected and invites are rate limited per store. The affiliate program must be enabled. Requires the `affiliate` token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliate-program/invite-an-affiliate) · Effect: **consequential**

```sh
sellapp affiliate-program invite --email 'alex.morgan@example.com' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | Yes | Value for email. |

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

Inspect the complete schema: `sellapp commands affiliate-program invite --json`.

## list

List pending affiliate invitations

List pending invitations that have not completed affiliate onboarding. The one-time onboarding link is delivered by email and is never returned by the API. Requires the `affiliate` token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliate-program/list-affiliate-invitations) · Effect: **read**

```sh
sellapp affiliate-program list
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

Inspect the complete schema: `sellapp commands affiliate-program list --json`.

## replace

Replace affiliate program configuration

Replace the complete affiliate-program configuration. The products array is authoritative: enabled entries replace existing store-level product restrictions, and every variant ID must belong to the authenticated store. Requires the `affiliate` token ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/affiliate-program/replace-affiliate-program) · Effect: **consequential**

```sh
sellapp affiliate-program replace --body '{"enabled":true,"settings":{"auto_approve_affiliates":false,"minimum_payout":"25","commission":{"type":"percentage","amount":"20"},"referrer_type":"first_referrer","tracking_length":30,"subscription_commission":true,"enabled_specific_products":true,"payout_methods":["PAYPAL"],"enable_hub":false},"products":[{"id":42,"enabled":true,"commission":{"type":"percentage","percentage":"25"}}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --enabled | Yes | Value for enabled. |
| --settings | Yes | Value for settings. |
| --products | Yes | Value for products. |

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

Inspect the complete schema: `sellapp commands affiliate-program replace --json`.

