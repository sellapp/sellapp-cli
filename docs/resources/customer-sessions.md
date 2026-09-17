# customer-sessions

[All commands](../commands.md)

## create-customer-session

Create a customer session

Return a 15-minute bearer token once plus a separate single-use hosted redemption URL. Opening the URL signs the customer into this store's customer portal only. It does not sign them into the dashboard or other stores. Browser access and wallet credentials issued from it stop working when the session expires or is revoked. This response is never cached or replayed. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customer-sessions) · Effect: **consequential**

```sh
sellapp customer-sessions create-customer-session --external-customer-id crm_maya_314 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --customer-id | No | Value for customer id. |
| --external-customer-id | No | Value for external customer id. |

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

Inspect the complete schema: `sellapp commands customer-sessions create-customer-session --json`.

## revoke-customer-session

Revoke a customer session

Revoke immediately. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customer-sessions) · Effect: **consequential**

```sh
sellapp customer-sessions revoke-customer-session session_01K4CUSTOMER --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| session (positional) | Yes | The session identifier. |

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

Inspect the complete schema: `sellapp commands customer-sessions revoke-customer-session --json`.

