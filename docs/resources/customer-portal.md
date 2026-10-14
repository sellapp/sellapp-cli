# customer-portal

[All commands](../commands.md)

## cancel-customer-subscription-at-period-end

Cancel at period end

Schedule the subscription to end when the current period ends. The customer keeps access until then. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal cancel-customer-subscription-at-period-end 42 --reason 'Taking a break over the summer.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
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

End the subscription now. A customer session cannot request a refund. It is the one action left on a Stripe or PayPal subscription with a payment still owed, and one of two left on a Razorpay one alongside a payment method update. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal cancel-customer-subscription-immediately 42 --reason 'No longer needed.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
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

Confirm a previewed plan change, sending back its `preview_token`. When the change still needs the customer to approve it with the provider or pay for it, the action is `pending` and `redirect_url` says where. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal confirm-customer-subscription-plan-change 42 --target-variant-id 97 --preview-token subprev_portal_991 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --target-variant-id | Yes | Value for target variant id. |
| --preview-token | Yes | Value for preview token. |
| --effective-timing | No | Value for effective timing. |
| --proration-behavior | No | Value for proration behavior. |
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

Inspect the complete schema: `sellapp commands customer-portal confirm-customer-subscription-plan-change --json`.

## confirm-customer-subscription-renewal-date-change

Confirm a renewal-date change

Confirm a renewal-date change, sending the `preview_token` when the plan requires a preview. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal confirm-customer-subscription-renewal-date-change 42 --renewal-date 2026-10-15T00:00:00Z --preview-token subprev_portal_renewal_991 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --renewal-date | Yes | Value for renewal date. |
| --preview-token | No | Value for preview token. |
| --reason | No | Value for reason. |
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

Read immediately before displaying lifecycle controls. Each of the nine actions (`pause`, `resume`, `cancel_at_period_end`, `cancel_immediately`, `keep_subscription`, `update_payment_method`, `change_plan`, `shift_billing_date`, and `set_auto_renew`) reports a status and the reason to show the customer when it is not available. While a Stripe, PayPal, or Razorpay subscription has a payment still owed, every action except `cancel_immediately` and, on Razorpay, `update_payment_method`, is `unsupported`. `resume` is `unsupported` on a subscription on hold for an unpaid renewal, which the customer pays and resumes on the store's customer portal page. The example is abbreviated to three of the actions.

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

List subscriptions belonging to the customer session, newest first, 15 per page. `trial_ends_at` is only set while a subscription is `TRIALING`.

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

Pause billing, optionally until `resume_at`. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal pause-customer-subscription 42 --resume-at 2026-11-01T00:00:00Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --resume-at | No | Value for resume at. |
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

Preview a plan change before confirming it. On Stripe and in-house subscriptions, `preview.payment_required` says whether confirming charges `amount_due` today or leaves any amount to the next invoice, and `preview.renewal_date` says when the subscription renews next. A plan change during a trial never charges today. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal preview-customer-subscription-plan-change 42 --target-variant-id 97 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --target-variant-id | Yes | Value for target variant id. |
| --effective-timing | No | Value for effective timing. |
| --proration-behavior | No | Value for proration behavior. |
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

Inspect the complete schema: `sellapp commands customer-portal preview-customer-subscription-plan-change --json`.

## preview-customer-subscription-renewal-date-change

Preview a renewal-date change

Preview moving the next renewal to `renewal_date`. `preview` is the provider's own preview, so its keys differ per provider. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal preview-customer-subscription-renewal-date-change 42 --renewal-date 2026-10-15T00:00:00Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --renewal-date | Yes | Value for renewal date. |
| --reason | No | Value for reason. |
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

Inspect the complete schema: `sellapp commands customer-portal preview-customer-subscription-renewal-date-change --json`.

## resume-customer-subscription

Resume a subscription

Resume a paused subscription. A subscription the store bills in house that is on hold for an unpaid renewal comes back only when the customer pays that renewal on the store's customer portal page, so it is refused. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal resume-customer-subscription 42 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |

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

Start a payment method update. Send the customer to `redirect_url` to finish it with the provider. Read capabilities first. A refusal returns 422 with the sentence to show the customer under `errors.subscription`.

[API reference](https://sell.app/docs/api/customer-portal) · Effect: **consequential**

```sh
sellapp customer-portal update-customer-subscription-payment-method 42 --return-url https://acme.sell.app/customer-portal/subscriptions/991 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription identifier. |
| --return-url | No | Value for return url. |

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

