# cashback-rules

[All commands](../commands.md)

## archive

Archive a cashback rule

Archive a cashback rule. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/cashback-rules) · Effect: **consequential**

```sh
sellapp cashback-rules archive 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| cashbackRule (positional) | Yes | The cashbackRule path parameter. |

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

Inspect the complete schema: `sellapp commands cashback-rules archive --json`.

## create

Create a cashback rule

Create a cashback rule. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/cashback-rules) · Effect: **consequential**

```sh
sellapp cashback-rules create --body '{"percent_basis":500,"maximum_cashback_cents":1000,"is_active":false,"product_ids":[120]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --percent-basis | Yes | Value for percent basis. |
| --maximum-cashback-cents | Yes | Value for maximum cashback cents. |
| --is-active | Yes | Value for is active. |
| --product-ids | Yes | Value for product ids. |

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

Inspect the complete schema: `sellapp commands cashback-rules create --json`.

## list

List cashback rules

List current and archived store records. Requires the `wallet` or `wallet:write` API ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/cashback-rules) · Effect: **read**

```sh
sellapp cashback-rules list
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

Inspect the complete schema: `sellapp commands cashback-rules list --json`.

## replace

Update a cashback rule

Update a cashback rule. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/cashback-rules) · Effect: **consequential**

```sh
sellapp cashback-rules replace 1 --body '{"percent_basis":500,"maximum_cashback_cents":1000,"is_active":false,"product_ids":[120]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| cashbackRule (positional) | Yes | The cashbackRule path parameter. |
| --percent-basis | Yes | Value for percent basis. |
| --maximum-cashback-cents | Yes | Value for maximum cashback cents. |
| --is-active | Yes | Value for is active. |
| --product-ids | Yes | Value for product ids. |

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

Inspect the complete schema: `sellapp commands cashback-rules replace --json`.

## restore

Restore a cashback rule

Restore a cashback rule. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/cashback-rules) · Effect: **consequential**

```sh
sellapp cashback-rules restore 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| cashbackRule (positional) | Yes | The cashbackRule path parameter. |

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

Inspect the complete schema: `sellapp commands cashback-rules restore --json`.

## update

Update a cashback rule

Update a cashback rule. Requires the `wallet:write` API ability and the store wallet permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/wallet/cashback-rules) · Effect: **consequential**

```sh
sellapp cashback-rules update 1 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| cashbackRule (positional) | Yes | The cashbackRule path parameter. |
| --percent-basis | No | Value for percent basis. |
| --maximum-cashback-cents | No | Value for maximum cashback cents. |
| --is-active | No | Value for is active. |
| --product-ids | No | Value for product ids. |

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

Inspect the complete schema: `sellapp commands cashback-rules update --json`.

