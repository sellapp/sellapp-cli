# customers

[All commands](../commands.md)

## create-customer

Create a customer

Create a customer with normalized unique email and immutable-per-store external_id. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **consequential**

```sh
sellapp customers create-customer --body '{"id":314,"email":"maya.chen@example.com","external_id":"crm_maya_314","name":"Maya Chen","locale":"en-GB","metadata":{"plan":"standard","seats":3}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | Yes | Value for email. |
| --external-id | No | Value for external id. |
| --name | No | Value for name. |
| --locale | No | Value for locale. |
| --metadata | No | Value for metadata. |

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

Inspect the complete schema: `sellapp commands customers create-customer --json`.

## get

Retrieve a customer

Retrieve a store-scoped customer profile and concise order, line-item, subscription, and revenue insights. Wallet insights are included only when the token and staff role also have wallet or store management permission. Requires the `invoice` Sanctum ability and matching store permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/retrieve-customer) · Effect: **read**

```sh
sellapp customers get 125
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer path parameter. |

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

Inspect the complete schema: `sellapp commands customers get --json`.

## get-customer-by-external-id

Retrieve a customer by external ID

Resolve by immutable store-local external ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **read**

```sh
sellapp customers get-customer-by-external-id 314
```

| Input | Required | Meaning |
| --- | --- | --- |
| externalId (positional) | Yes | The externalId identifier. |

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

Inspect the complete schema: `sellapp commands customers get-customer-by-external-id --json`.

## list

List customers

List customer profiles for the authenticated store. Requires the `invoice` Sanctum ability and matching store permission. Responses deliberately redact checkout IP, billing, location, and community identity data. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/list-customers) · Effect: **read**

```sh
sellapp customers list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching customers without pagination metadata. |

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

Inspect the complete schema: `sellapp commands customers list --json`.

## search

Search customers

Search customer email identities and compose supported ID or email filters and deterministic sorts. Requires the `invoice` Sanctum ability and matching store permission. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/search-customers) · Effect: **read**

```sh
sellapp customers search --body '{"filters":[{"field":"id","operator":"=","value":125}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching customers without pagination metadata. |
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

Inspect the complete schema: `sellapp commands customers search --json`.

## update-customer

Update a customer

Update name, BCP 47 locale, normalized email, or constrained scalar metadata. Assign external_id when it is unset; changing an existing external_id returns 409. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **consequential**

```sh
sellapp customers update-customer 314 --locale en-US --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customer (positional) | Yes | The customer identifier. |
| --email | No | Value for email. |
| --name | No | Value for name. |
| --locale | No | Value for locale. |
| --metadata | No | Value for metadata. |
| --external-id | No | May be assigned when unset; an existing value is immutable. |

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

Inspect the complete schema: `sellapp commands customers update-customer --json`.

## update-customer-by-external-id

Update a customer by external ID

Update the customer resolved by immutable external ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **consequential**

```sh
sellapp customers update-customer-by-external-id 314 --locale en-US --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| externalId (positional) | Yes | The externalId identifier. |
| --email | No | Value for email. |
| --name | No | Value for name. |
| --locale | No | Value for locale. |
| --metadata | No | Value for metadata. |
| --body-external-id | No | May be assigned when unset; an existing value is immutable. |

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

Inspect the complete schema: `sellapp commands customers update-customer-by-external-id --json`.

## upsert-by-external-id

Create or update a customer by external ID

Create a customer when the store-local external ID is new, or update the same customer when it already exists. Requires the invoice API-key ability and store invoice permission. Send email on every call; omitted name, locale, and metadata remain unchanged. Metadata replaces the complete metadata object when supplied. The external ID comes from the URL and cannot be changed in the body. An email already used by another customer returns 422; a concurrent email or external-ID collision returns 409. Retrieve the customer and check its identity before retrying. Accounts are never merged. Returns 201 on creation and 200 on update. Reusing the external ID preserves customer identity. Use the same optional Idempotency-Key and body to replay an identical response. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/customers/identity-and-entitlements) · Effect: **consequential**

```sh
sellapp customers upsert-by-external-id crm_maya_314 --email 'maya.chen@example.com' --name 'Maya Chen' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| externalId (positional) | Yes | Store-local external ID, 1–255 characters; whitespace-only values are invalid. URL-encode it as one path segment. |
| --email | Yes | Value for email. |
| --name | No | Value for name. |
| --locale | No | Value for locale. |
| --metadata | No | Value for metadata. |

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

Inspect the complete schema: `sellapp commands customers upsert-by-external-id --json`.

