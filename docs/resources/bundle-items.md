# bundle-items

[All commands](../commands.md)

## attach

Attach bundle items

Attach one or more product variants to a bundle with quantities. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **consequential**

```sh
sellapp bundle-items attach 1 --body '{"resources":{"1":{"quantity":1}}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| bundle (positional) | Yes | The bundle path parameter. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands bundle-items attach --json`.

## detach

Detach bundle items

Detach one or more product variants from a bundle. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **consequential**

```sh
sellapp bundle-items detach 1 --body '{"resources":[1]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| bundle (positional) | Yes | The bundle path parameter. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands bundle-items detach --json`.

## get

Retrieve a bundle item

Retrieve a specific item from a bundle by its product variant ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **read**

```sh
sellapp bundle-items get --bundle 1 2
```

| Input | Required | Meaning |
| --- | --- | --- |
| --bundle | Yes | The bundle path parameter. |
| item (positional) | Yes | The item path parameter. |

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

Inspect the complete schema: `sellapp commands bundle-items get --json`.

## list

List bundle items

Retrieve a paginated list of product variants attached to a bundle. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/products) · Effect: **read**

```sh
sellapp bundle-items list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| bundle (positional) | Yes | The bundle path parameter. |
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

Inspect the complete schema: `sellapp commands bundle-items list --json`.

