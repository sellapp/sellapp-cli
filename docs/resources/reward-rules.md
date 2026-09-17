# reward-rules

[All commands](../commands.md)

## create

Create a reward rule

Create reward rule configuration. Grant issuance remains a separate operation. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-rules/create-a-reward-rule) · Effect: **consequential**

```sh
sellapp reward-rules create --body '{"name":"Launch Lab regular","is_active":false,"trigger_type":"PURCHASE_COUNT","trigger_threshold":3,"outputs":[{"type":"BADGE","label":"Launch Lab regular","color":"violet"}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Value for name. |
| --is-active | Yes | Value for is active. |
| --trigger-type | Yes | Value for trigger type. |
| --trigger-threshold | Yes | Value for trigger threshold. |
| --outputs | Yes | Value for outputs. |
| --description | No | Value for description. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands reward-rules create --json`.

## get

Retrieve a reward rule

Retrieve one store-owned reward rule. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-rules/retrieve-a-reward-rule) · Effect: **read**

```sh
sellapp reward-rules get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| rewardRule (positional) | Yes | The rewardRule path parameter. |

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

Inspect the complete schema: `sellapp commands reward-rules get --json`.

## list

List reward rules

List reward rules owned by the authenticated store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-rules/list-reward-rules) · Effect: **read**

```sh
sellapp reward-rules list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |

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

Inspect the complete schema: `sellapp commands reward-rules list --json`.

## replace

Update a reward rule

Update reward rule configuration without directly issuing a customer grant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-rules/update-a-reward-rule) · Effect: **consequential**

```sh
sellapp reward-rules replace 1 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| rewardRule (positional) | Yes | The rewardRule path parameter. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --is-active | No | Value for is active. |
| --trigger-type | No | Value for trigger type. |
| --trigger-threshold | No | Value for trigger threshold. |
| --outputs | No | Value for outputs. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands reward-rules replace --json`.

## search

Search reward rules

Search and filter reward rules owned by the authenticated store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-rules/search-reward-rules) · Effect: **read**

```sh
sellapp reward-rules search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --filters | No | Value for filters. |
| --sort | No | Value for sort. |
| --search | No | Value for search. |
| --includes | No | Value for includes. |

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

Inspect the complete schema: `sellapp commands reward-rules search --json`.

## update

Update a reward rule

Update reward rule configuration without directly issuing a customer grant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-rules/update-a-reward-rule) · Effect: **consequential**

```sh
sellapp reward-rules update 1 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| rewardRule (positional) | Yes | The rewardRule path parameter. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --is-active | No | Value for is active. |
| --trigger-type | No | Value for trigger type. |
| --trigger-threshold | No | Value for trigger threshold. |
| --outputs | No | Value for outputs. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands reward-rules update --json`.

