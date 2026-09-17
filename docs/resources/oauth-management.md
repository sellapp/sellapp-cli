# oauth-management

[All commands](../commands.md)

## delete-oauth-installation

Disconnect your CLI connection

Disconnect your official CLI connection and revoke its authorization codes and token families. Run sellapp login to reconnect.

[API reference](https://sell.app/docs/api/oauth) · Effect: **consequential**

```sh
sellapp oauth-management delete-oauth-installation --yes
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth-management delete-oauth-installation --json`.

## get-oauth-installation

Read your CLI connection

Read your active official CLI connection, its account-wide grant model and granted scopes.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp oauth-management get-oauth-installation
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "oauthAccessToken",
      "scopes": [
        "admin"
      ]
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth-management get-oauth-installation --json`.

