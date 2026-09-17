# customer-portal

[All commands](../commands.md)

## cancel-customer-subscription-at-period-end

Cancel at period end

Cancel at period end. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal cancel-customer-subscription-at-period-end 42 --reason 'Customer requested this change' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal cancel-customer-subscription-at-period-end --json`.

## cancel-customer-subscription-immediately

Cancel immediately

Cancel immediately. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal cancel-customer-subscription-immediately 42 --reason 'Customer requested this change' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal cancel-customer-subscription-immediately --json`.

## confirm-customer-subscription-plan-change

Confirm a plan change

Confirm a plan change. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal confirm-customer-subscription-plan-change 42 --preview-id preview_01K4 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal confirm-customer-subscription-plan-change --json`.

## confirm-customer-subscription-renewal-date-change

Confirm a renewal-date change

Confirm a renewal-date change. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal confirm-customer-subscription-renewal-date-change 42 --preview-id preview_01K4 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal confirm-customer-subscription-renewal-date-change --json`.

## get-customer-portal-order

Retrieve a customer order

Retrieve a purchase belonging to the customer session.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal get-customer-portal-order 9001
```

| Input | Required | Meaning |
| --- | --- | --- |
| order (positional) | Yes | The order identifier. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal get-customer-portal-order --json`.

## get-customer-portal-profile

Retrieve the signed-in customer

Customer-session bearer authentication; resource is redacted.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal get-customer-portal-profile
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal get-customer-portal-profile --json`.

## get-customer-portal-subscription

Retrieve a customer subscription

Retrieve a subscription belonging to the customer session.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal get-customer-portal-subscription 991
```

| Input | Required | Meaning |
| --- | --- | --- |
| subscription (positional) | Yes | The subscription identifier. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal get-customer-portal-subscription --json`.

## get-customer-portal-subscription-capabilities

Retrieve subscription capabilities

Read immediately before displaying lifecycle controls.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal get-customer-portal-subscription-capabilities 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| subscription (positional) | Yes | The subscription identifier. |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal get-customer-portal-subscription-capabilities --json`.

## list-customer-portal-entitlements

List customer entitlements

List the customer's purchase access summary without fulfillment secrets.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal list-customer-portal-entitlements
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal list-customer-portal-entitlements --json`.

## list-customer-portal-orders

List customer orders

List purchases belonging to the customer session.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal list-customer-portal-orders
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal list-customer-portal-orders --json`.

## list-customer-portal-subscriptions

List customer subscriptions

List subscriptions belonging to the customer session.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **read**

```sh
sellapp customer-portal list-customer-portal-subscriptions
```

| Input | Required | Meaning |
| --- | --- | --- |

A request body is optional or not used by this operation.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal list-customer-portal-subscriptions --json`.

## pause-customer-subscription

Pause a subscription

Pause a subscription. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal pause-customer-subscription 42 --reason 'Customer requested this change' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal pause-customer-subscription --json`.

## preview-customer-subscription-plan-change

Preview a plan change

Preview a plan change. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal preview-customer-subscription-plan-change 42 --product-variant-id 84 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal preview-customer-subscription-plan-change --json`.

## preview-customer-subscription-renewal-date-change

Preview a renewal-date change

Preview a renewal-date change. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal preview-customer-subscription-renewal-date-change 42 --renewal-date 2026-10-15 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal preview-customer-subscription-renewal-date-change --json`.

## resume-customer-subscription

Resume a subscription

Resume a subscription. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal resume-customer-subscription 42 --reason 'Customer requested this change' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal resume-customer-subscription --json`.

## update-customer-portal-profile

Update the signed-in customer

Update customer-safe profile fields.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal update-customer-portal-profile --locale en-US --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | No | Value for email. |
| --name | No | Value for name. |
| --locale | No | Value for locale. |
| --metadata | No | Value for metadata. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal update-customer-portal-profile --json`.

## update-customer-subscription-payment-method

Update payment method

Update payment method. Read capabilities first; state conflicts return 409.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal update-customer-subscription-payment-method 42 --reason 'Customer requested this change' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --preview-id | No | Value for preview id. |
| --product-variant-id | No | Value for product variant id. |
| --renewal-date | No | Value for renewal date. |
| --return-url | No | Value for return url. |
| --reason | No | Value for reason. |

The request body is required.

Credential alternatives and scopes:

```json
[
  [
    {
      "schemeName": "customerSession",
      "scopes": []
    }
  ]
]
```

Response data is printed to stdout; diagnostics go to stderr. Use `--output json` for automation. [Output, errors, and transport controls](../usage.md).

Inspect the complete schema: `sellapp commands customer-portal update-customer-subscription-payment-method --json`.

