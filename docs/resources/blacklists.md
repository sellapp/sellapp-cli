# blacklists

[All commands](../commands.md)

## create

Create a blacklist rule

Create a new blacklist rule for your store.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp blacklists create --type ASN --data test_data --description test_description --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --type | Yes | Value for type. |
| --data | Yes | The value to blacklist. |
| --description | Yes | Why this rule is being created. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands blacklists create --json`.

## delete

Delete a blacklist rule

Permanently delete a blacklist rule from your store.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp blacklists delete 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The unique identifier of the blacklist rule. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands blacklists delete --json`.

## get

Retrieve a blacklist rule

Retrieve a specific blacklist rule by its unique identifier.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp blacklists get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The unique identifier of the blacklist rule. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands blacklists get --json`.

## list

List all blacklist rules

Retrieve a paginated list of blacklist rules for your store. By default, a maximum of fifteen blacklist rules are returned per page.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **read**

```sh
sellapp blacklists list
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
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands blacklists list --json`.

## update

Update a blacklist rule

Update one or more attributes on an existing blacklist rule.

[API reference](https://sell.app/docs/api/legacy-v1) · Effect: **consequential**

```sh
sellapp blacklists update 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The unique identifier of the blacklist rule. |
| --type | No | Value for type. |
| --data | No | The updated value to blacklist. |
| --description | No | The updated reason for this rule. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "bearerAuth",
      "scopes": []
    },
    {
      "schemeName": "storeAuth",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands blacklists update --json`.

## v-2-create-blacklist

Create a blacklist rule

Create a store-scoped blacklist rule. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/blacklists/create-blacklist) · Effect: **consequential**

```sh
sellapp blacklists v-2-create-blacklist --type EMAIL --data 'blocked@example.com' --description 'Blocked after a verified fraud report.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --type | Yes | Value for type. |
| --data | Yes | The value to blacklist. |
| --description | Yes | Why this rule is being created. |

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

Inspect the complete schema: `sellapp commands blacklists v-2-create-blacklist --json`.

## v-2-delete-blacklist

Delete a blacklist rule

Delete a blacklist rule in the selected store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/blacklists/delete-blacklist) · Effect: **consequential**

```sh
sellapp blacklists v-2-delete-blacklist 42 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The blacklist identifier. |

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

Inspect the complete schema: `sellapp commands blacklists v-2-delete-blacklist --json`.

## v-2-get-blacklist

Retrieve a blacklist rule

Retrieve a blacklist rule in the selected store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/blacklists/retrieve-blacklist) · Effect: **read**

```sh
sellapp blacklists v-2-get-blacklist 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The blacklist identifier. |

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

Inspect the complete schema: `sellapp commands blacklists v-2-get-blacklist --json`.

## v-2-list-blacklists

List blacklist rules

List store-scoped blacklist rules. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/blacklists/list-blacklists) · Effect: **read**

```sh
sellapp blacklists v-2-list-blacklists
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

Inspect the complete schema: `sellapp commands blacklists v-2-list-blacklists --json`.

## v-2-replace-blacklist

Replace a blacklist rule

Replace a blacklist rule in the selected store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/blacklists) · Effect: **consequential**

```sh
sellapp blacklists v-2-replace-blacklist 42 --description 'Blocked after a verified fraud report.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The blacklist identifier. |
| --type | No | Value for type. |
| --data | No | The updated value to blacklist. |
| --description | No | The updated reason for this rule. |

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

Inspect the complete schema: `sellapp commands blacklists v-2-replace-blacklist --json`.

## v-2-update-blacklist

Update a blacklist rule

Update a blacklist rule in the selected store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/blacklists/update-blacklist) · Effect: **consequential**

```sh
sellapp blacklists v-2-update-blacklist 42 --description 'Blocked after a verified fraud report.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| blacklist (positional) | Yes | The blacklist identifier. |
| --type | No | Value for type. |
| --data | No | The updated value to blacklist. |
| --description | No | The updated reason for this rule. |

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

Inspect the complete schema: `sellapp commands blacklists v-2-update-blacklist --json`.

