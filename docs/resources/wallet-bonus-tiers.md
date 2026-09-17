# wallet-bonus-tiers

[All commands](../commands.md)

## archive

Archive a wallet bonus tier

Archive a wallet bonus tier. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/bonus-tiers) · Effect: **consequential**

```sh
sellapp wallet-bonus-tiers archive 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| bonusTier (positional) | Yes | The bonusTier path parameter. |

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

Inspect the complete schema: `sellapp commands wallet-bonus-tiers archive --json`.

## create

Create a wallet bonus tier

Create a wallet bonus tier. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/bonus-tiers) · Effect: **consequential**

```sh
sellapp wallet-bonus-tiers create --minimum-top-up-cents 10000 --bonus-kind fixed --fixed-bonus-cents 500 --percent-basis null --maximum-bonus-cents null --priority 0 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --minimum-top-up-cents | Yes | Value for minimum top up cents. |
| --bonus-kind | Yes | Value for bonus kind. |
| --fixed-bonus-cents | Yes | Value for fixed bonus cents. |
| --percent-basis | Yes | Value for percent basis. |
| --maximum-bonus-cents | Yes | Value for maximum bonus cents. |
| --priority | Yes | Value for priority. |
| --is-active | Yes | Value for is active. |

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

Inspect the complete schema: `sellapp commands wallet-bonus-tiers create --json`.

## list

List wallet bonus tiers

List current and archived store records. Requires the `wallet` or `wallet:write` API ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/bonus-tiers) · Effect: **read**

```sh
sellapp wallet-bonus-tiers list
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

Inspect the complete schema: `sellapp commands wallet-bonus-tiers list --json`.

## replace

Update a wallet bonus tier

Update a wallet bonus tier. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/bonus-tiers) · Effect: **consequential**

```sh
sellapp wallet-bonus-tiers replace 1 --minimum-top-up-cents 10000 --bonus-kind fixed --fixed-bonus-cents 500 --percent-basis null --maximum-bonus-cents null --priority 0 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| bonusTier (positional) | Yes | The bonusTier path parameter. |
| --minimum-top-up-cents | Yes | Value for minimum top up cents. |
| --bonus-kind | Yes | Value for bonus kind. |
| --fixed-bonus-cents | Yes | Value for fixed bonus cents. |
| --percent-basis | Yes | Value for percent basis. |
| --maximum-bonus-cents | Yes | Value for maximum bonus cents. |
| --priority | Yes | Value for priority. |
| --is-active | Yes | Value for is active. |

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

Inspect the complete schema: `sellapp commands wallet-bonus-tiers replace --json`.

## restore

Restore a wallet bonus tier

Restore a wallet bonus tier. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/bonus-tiers) · Effect: **consequential**

```sh
sellapp wallet-bonus-tiers restore 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| bonusTier (positional) | Yes | The bonusTier path parameter. |

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

Inspect the complete schema: `sellapp commands wallet-bonus-tiers restore --json`.

## update

Update a wallet bonus tier

Update a wallet bonus tier. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/bonus-tiers) · Effect: **consequential**

```sh
sellapp wallet-bonus-tiers update 1 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| bonusTier (positional) | Yes | The bonusTier path parameter. |
| --minimum-top-up-cents | No | Value for minimum top up cents. |
| --bonus-kind | No | Value for bonus kind. |
| --fixed-bonus-cents | No | Value for fixed bonus cents. |
| --percent-basis | No | Value for percent basis. |
| --maximum-bonus-cents | No | Value for maximum bonus cents. |
| --priority | No | Value for priority. |
| --is-active | No | Value for is active. |

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

Inspect the complete schema: `sellapp commands wallet-bonus-tiers update --json`.

