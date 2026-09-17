# licenses

[All commands](../commands.md)

## activate

Activate a license key

Activate a license for a named device or installation. Save the returned instance ID to validate that activation later. Activation fails if the key is missing, inactive, expired, or at its activation limit. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/licenses/activate-a-license-key) · Effect: **consequential**

```sh
sellapp licenses activate --license-key 01965f1d-f038-7116-b57f-9e7ecb4e7b8f --instance-name 'Grace Wilson' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --license-key | Yes | Value for license key. |
| --instance-name | Yes | Value for instance name. |

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

Inspect the complete schema: `sellapp commands licenses activate --json`.

## deactivate-license

Deactivate a license instance

Deactivate one installation. The response never exposes the raw license key. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/licenses/deactivate-a-license) · Effect: **consequential**

```sh
sellapp licenses deactivate-license --license-key SELL-LICENSE-REDACTED --instance-id laptop-maya --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --license-key | Yes | Value for license key. |
| --instance-id | Yes | Value for instance id. |

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

Inspect the complete schema: `sellapp commands licenses deactivate-license --json`.

## get

Retrieve a license key

Retrieve a license record by its ID, not the license key string a customer enters. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/licenses/retrieve-a-license-key) · Effect: **read**

```sh
sellapp licenses get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| license_key (positional) | Yes | The license key path parameter. |

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

Inspect the complete schema: `sellapp commands licenses get --json`.

## instances get

Retrieve a license instance

Retrieve an activation using the license key ID and the instance UUID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/license-instances/retrieve-a-license-instance) · Effect: **read**

```sh
sellapp licenses instances get --license-key 1 9ebd37af-2077-42f9-9f88-d96cfc6ef1a8
```

| Input | Required | Meaning |
| --- | --- | --- |
| --license-key | Yes | The license key path parameter. |
| instance (positional) | Yes | The instance path parameter. |

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

Inspect the complete schema: `sellapp commands licenses instances get --json`.

## instances list

List license instances

List the activations for a license key, one page at a time. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/license-instances/list-license-instances) · Effect: **read**

```sh
sellapp licenses instances list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| license_key (positional) | Yes | The license key path parameter. |
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

Inspect the complete schema: `sellapp commands licenses instances list --json`.

## list

List all license keys

List your store's license keys. Use `limit` and `page` to read one page at a time. The response schema describes each license. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/licenses/list-all-license-keys) · Effect: **read**

```sh
sellapp licenses list
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

Inspect the complete schema: `sellapp commands licenses list --json`.

## update

Update a license key

Change a license's activation limit, expiry date, or active status. If `limit` is lower than `instances_count`, the oldest instances are deleted until the count fits the new limit. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/licenses/update-a-license-key) · Effect: **consequential**

```sh
sellapp licenses update 1 --limit 10 --active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| license_key (positional) | Yes | The license key path parameter. |
| --limit | No | Value for limit. |
| --expires-at | No | Value for expires at. |
| --active | No | Value for active. |

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

Inspect the complete schema: `sellapp commands licenses update --json`.

## validate

Validate a license key

Checks whether a license key is currently valid. You may include `instance_id` to validate a specific activation; if omitted, the API validates the license key alone. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/licenses/validate-a-license-key) · Effect: **consequential**

```sh
sellapp licenses validate --license-key 01965f1d-f038-7116-b57f-9e7ecb4e7b8f --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --license-key | Yes | Value for license key. |
| --instance-id | No | Value for instance id. |

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

Inspect the complete schema: `sellapp commands licenses validate --json`.

