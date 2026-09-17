# reward-coupon-templates

[All commands](../commands.md)

## create

Create a reward coupon template

Create reward coupon template configuration. Grant issuance remains a separate operation. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-coupon-templates/create-a-reward-coupon-template) · Effect: **consequential**

```sh
sellapp reward-coupon-templates create --body '{"name":"Launch Lab thank you","type":"PERCENTAGE","discount":"10.00","store_wide":true,"redemption_mode":"customer_locked","is_active":false,"listing_ids":[]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Value for name. |
| --type | Yes | Value for type. |
| --discount | Yes | Value for discount. |
| --store-wide | Yes | Value for store wide. |
| --redemption-mode | Yes | Value for redemption mode. |
| --is-active | Yes | Value for is active. |
| --listing-ids | Yes | Value for listing ids. |
| --minimum-amount | No | Value for minimum amount. |
| --maximum-discount-amount | No | Value for maximum discount amount. |
| --expires-at | No | Value for expires at. |
| --expires-after-days | No | Value for expires after days. |
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

Inspect the complete schema: `sellapp commands reward-coupon-templates create --json`.

## get

Retrieve a reward coupon template

Retrieve one store-owned reward coupon template. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-coupon-templates/retrieve-a-reward-coupon-template) · Effect: **read**

```sh
sellapp reward-coupon-templates get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| rewardCouponTemplate (positional) | Yes | The rewardCouponTemplate path parameter. |

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

Inspect the complete schema: `sellapp commands reward-coupon-templates get --json`.

## list

List reward coupon templates

List reward coupon templates owned by the authenticated store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-coupon-templates/list-reward-coupon-templates) · Effect: **read**

```sh
sellapp reward-coupon-templates list
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

Inspect the complete schema: `sellapp commands reward-coupon-templates list --json`.

## replace

Update a reward coupon template

Update reward coupon template configuration without directly issuing a customer grant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-coupon-templates/update-a-reward-coupon-template) · Effect: **consequential**

```sh
sellapp reward-coupon-templates replace 1 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| rewardCouponTemplate (positional) | Yes | The rewardCouponTemplate path parameter. |
| --name | No | Value for name. |
| --type | No | Value for type. |
| --discount | No | Value for discount. |
| --store-wide | No | Value for store wide. |
| --minimum-amount | No | Value for minimum amount. |
| --maximum-discount-amount | No | Value for maximum discount amount. |
| --expires-at | No | Value for expires at. |
| --expires-after-days | No | Value for expires after days. |
| --redemption-mode | No | Value for redemption mode. |
| --is-active | No | Value for is active. |
| --listing-ids | No | Value for listing ids. |
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

Inspect the complete schema: `sellapp commands reward-coupon-templates replace --json`.

## search

Search reward coupon templates

Search and filter reward coupon templates owned by the authenticated store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-coupon-templates/search-reward-coupon-templates) · Effect: **read**

```sh
sellapp reward-coupon-templates search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
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

Inspect the complete schema: `sellapp commands reward-coupon-templates search --json`.

## update

Update a reward coupon template

Update reward coupon template configuration without directly issuing a customer grant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/reward-coupon-templates/update-a-reward-coupon-template) · Effect: **consequential**

```sh
sellapp reward-coupon-templates update 1 --is-active false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| rewardCouponTemplate (positional) | Yes | The rewardCouponTemplate path parameter. |
| --name | No | Value for name. |
| --type | No | Value for type. |
| --discount | No | Value for discount. |
| --store-wide | No | Value for store wide. |
| --minimum-amount | No | Value for minimum amount. |
| --maximum-discount-amount | No | Value for maximum discount amount. |
| --expires-at | No | Value for expires at. |
| --expires-after-days | No | Value for expires after days. |
| --redemption-mode | No | Value for redemption mode. |
| --is-active | No | Value for is active. |
| --listing-ids | No | Value for listing ids. |
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

Inspect the complete schema: `sellapp commands reward-coupon-templates update --json`.

