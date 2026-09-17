# oauth

[All commands](../commands.md)

## approve-oauth-authorization

Approve CLI access

Submit the official CLI consent form in the same logged-in browser session. The admin grant covers every store you can access now or later, including when you currently have no stores. It also permits store creation. Local agents connected through the CLI share this connection. Approval never adds store role permissions. Current client eligibility, membership and permissions are rechecked, including for stale consent.

[API reference](https://sell.app/docs/api/oauth) · Effect: **consequential**

```sh
sellapp oauth approve-oauth-authorization --auth-token CONSENT_AUTH_TOKEN --client-id 01992a65-e064-71ba-b38f-902b7966a6be --state RANDOM_STATE --token CSRF_TOKEN --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --auth-token | Yes | Value for auth token. |
| --client-id | Yes | Value for client id. |
| --state | Yes | Value for state. |
| --token | Yes | Value for  token. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "oauthBrowserSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth approve-oauth-authorization --json`.

## deny-oauth-authorization

Deny CLI access

Submit the denial form from the same browser session. The callback receives access_denied and the original state.

[API reference](https://sell.app/docs/api/oauth) · Effect: **consequential**

```sh
sellapp oauth deny-oauth-authorization --auth-token CONSENT_AUTH_TOKEN --token CSRF_TOKEN --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --auth-token | Yes | Value for auth token. |
| --token | Yes | Value for  token. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "oauthBrowserSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth deny-oauth-authorization --json`.

## exchange-oauth-token

Exchange or refresh OAuth tokens

Exchange an official CLI authorization code with the original redirect URI and PKCE verifier, or rotate a refresh token. Send the configured client_id without a secret. Client eligibility and current grant validity are rechecked. Access tokens last 60 minutes; refresh token families last 90 days. Atomically replace rotated refresh tokens. Reuse may revoke the family; reauthorize after invalid_grant. Fix registration or scope configuration for invalid_client or invalid_scope instead of repeating login.

[API reference](https://sell.app/docs/api/oauth) · Effect: **consequential**

```sh
sellapp oauth exchange-oauth-token --client-id 01992a65-e064-71ba-b38f-902b7966a6be --grant-type authorization_code --code AUTHORIZATION_CODE --redirect-uri http://127.0.0.1:49152/callback --code-verifier dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --grant-type | Yes | Value for grant type. |
| --client-id | No | The configured official CLI public client ID. It identifies the registration, not the executable. |
| --client-secret | No | Legacy SDK field retained for compatibility. Omit for the official CLI; it has no client secret. |
| --code | No | Value for code. |
| --redirect-uri | No | Value for redirect uri. |
| --code-verifier | No | Value for code verifier. |
| --refresh-token | No | Value for refresh token. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [],
  [
    {
      "schemeName": "oauthClientBasic",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth exchange-oauth-token --json`.

## get-oauth-authorization-request

Review CLI authorization

Open in the user's browser to authorize the official CLI using existing store permissions. Only the configured public CLI registration is supported. PKCE S256 and non-empty state are required. The callback may vary only the port of the registered numeric loopback URL; localhost, wildcard hosts, fragments and non-loopback HTTP are rejected. Unsupported clients receive invalid_client and unsupported scopes receive invalid_scope.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp oauth get-oauth-authorization-request --response-type code --client-id 01992a65-e064-71ba-b38f-902b7966a6be --redirect-uri http://127.0.0.1:49152/callback --state RANDOM_STATE --code-challenge E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM --code-challenge-method S256
```

| Input | Required | Meaning |
| --- | --- | --- |
| --response-type | Yes | Value for response type. |
| --client-id | Yes | Value for client id. |
| --redirect-uri | Yes | Value for redirect uri. |
| --state | Yes | Value for state. |
| --code-challenge | Yes | Value for code challenge. |
| --code-challenge-method | Yes | Value for code challenge method. |
| --scope | No | Value for scope. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth get-oauth-authorization-request --json`.

## get-oauth-authorization-server-metadata

Read OAuth server metadata

Discover the official CLI authorization, token and revocation endpoints and supported scopes. The CLI uses public client authentication without a secret.

[API reference](https://sell.app/docs/api/oauth) · Effect: **read**

```sh
sellapp oauth get-oauth-authorization-server-metadata
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth get-oauth-authorization-server-metadata --json`.

## revoke-oauth-token

Revoke an OAuth token

Revoke the matching personal CLI installation and all its token families. Send the configured public client_id without a secret. Unknown tokens also succeed without disclosing whether a token existed.

[API reference](https://sell.app/docs/api/oauth) · Effect: **consequential**

```sh
sellapp oauth revoke-oauth-token --client-id 01992a65-e064-71ba-b38f-902b7966a6be --token REFRESH_TOKEN --token-type-hint refresh_token --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --token | Yes | Value for token. |
| --client-id | No | The configured official CLI public client ID. It identifies the registration, not the executable. |
| --client-secret | No | Legacy SDK field retained for compatibility. Omit for the official CLI; it has no client secret. |
| --token-type-hint | No | Value for token type hint. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [],
  [
    {
      "schemeName": "oauthClientBasic",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands oauth revoke-oauth-token --json`.

