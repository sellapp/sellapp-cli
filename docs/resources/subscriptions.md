# subscriptions

[All commands](../commands.md)

## cancel

Cancel a subscription

Cancel an active subscription now or at period end. cancel_at_period_end is required. Refund flags are valid only when it is false; period-end cancellation does not refund the last payment. Refunds are off by default. refund_last_payment requests a real refund; pro_rated_refund requires refund_last_payment and defaults to false. Cancellation can succeed even if the requested refund fails. A 200 response or confirmed cancellation does not prove a refund, and the action response does not expose the provider refund result. Check the order and payment-provider refund records before retrying; do not replay cancellation as a refund-recovery step. Reuse an idempotency key only with identical action input and actor. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/cancel-a-subscription-immediately-with-a-refund) · Effect: **consequential**

```sh
sellapp subscriptions cancel 55 --cancel-at-period-end true --body-idempotency-key design-kit-subscription-55-cancel-v1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| subscription (positional) | Yes | The subscription path parameter. |
| --cancel-at-period-end | Yes | Value for cancel at period end. |
| --refund-last-payment | No | Request a real refund of the latest payment during immediate cancellation. Cancellation success does not prove refund success; verify provider/order refund records. |
| --pro-rated-refund | No | Request a prorated refund. Requires refund_last_payment=true. False requests a full latest-payment refund when refund_last_payment is enabled. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |

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

Inspect the complete schema: `sellapp commands subscriptions cancel --json`.

## cancel-at-period-end

Cancel a subscription at period end

Schedule cancellation at the end of the current billing period without requesting a refund. Inspect capabilities first. The response records the cancellation action; retrieve the subscription to observe its current state. Reuse an idempotency key only with identical action input and actor. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/cancel-at-period-end) · Effect: **consequential**

```sh
sellapp subscriptions cancel-at-period-end 55 --reason 'Customer request' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --reason | No | Value for reason. |

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

Inspect the complete schema: `sellapp commands subscriptions cancel-at-period-end --json`.

## cancel-immediately

Cancel a subscription immediately

Request immediate subscription cancellation. Inspect capabilities for provider support. This ends future billing and can change customer access. Refunds are off by default. refund_last_payment requests a real refund; pro_rated_refund requires refund_last_payment and defaults to false. Cancellation can succeed even if the requested refund fails. A 200 response or confirmed cancellation does not prove a refund, and the action response does not expose the provider refund result. Check the order and payment-provider refund records before retrying; do not replay cancellation as a refund-recovery step. Reuse an idempotency key only with identical action input and actor. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/cancel-immediately) · Effect: **consequential**

```sh
sellapp subscriptions cancel-immediately 55 --reason 'Customer request' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --reason | No | Value for reason. |
| --refund-last-payment | No | Request a real refund of the latest payment during immediate cancellation. Cancellation success does not prove refund success; verify provider/order refund records. |
| --pro-rated-refund | No | Request a prorated refund. Requires refund_last_payment=true. False requests a full latest-payment refund when refund_last_payment is enabled. |

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

Inspect the complete schema: `sellapp commands subscriptions cancel-immediately --json`.

## confirm-plan-change

Confirm a subscription plan change

This action belongs to the customer. Seller API requests return a 422 validation error. Check subscription capabilities and use the customer flow; never expose a seller API key to work around this restriction. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/confirm-plan-change) · Effect: **consequential**

```sh
sellapp subscriptions confirm-plan-change 1 --target-variant-id 4321 --preview-token subprev_9c4b2f --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --target-variant-id | Yes | Value for target variant id. |
| --preview-token | Yes | Value for preview token. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --effective-timing | No | Value for effective timing. |
| --proration-behavior | No | Value for proration behavior. |
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

Inspect the complete schema: `sellapp commands subscriptions confirm-plan-change --json`.

## confirm-renewal-date

Confirm a subscription renewal date change

Confirm a subscription renewal date change. A preview token is required when the variant's billing-date policy requires previews. The current Stripe and PayPal provider flows do not support renewal-date changes and return 422. Check shift_billing_date capabilities before sending a request. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/confirm-renewal-date-change) · Effect: **consequential**

```sh
sellapp subscriptions confirm-renewal-date 1 --renewal-date 2026-10-01T12:00:00Z --preview-token subprev_project_library_55 --reason 'Align Maya'\''s membership with the monthly reading circle.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --renewal-date | Yes | Value for renewal date. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --preview-token | No | Value for preview token. |
| --reason | No | Value for reason. |
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

Inspect the complete schema: `sellapp commands subscriptions confirm-renewal-date --json`.

## get-capabilities

View subscription capabilities

Return the lifecycle actions currently supported for a subscription. Capability status controls whether an actor can pause, resume, cancel, update the payment method, change plan, or shift the billing date. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/view-subscription-capabilities) · Effect: **read**

```sh
sellapp subscriptions get-capabilities 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |

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

Inspect the complete schema: `sellapp commands subscriptions get-capabilities --json`.

## get-subscription

Retrieve a subscription

Retrieve by SellApp subscription record ID. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/read-subscriptions) · Effect: **read**

```sh
sellapp subscriptions get-subscription 991
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

Inspect the complete schema: `sellapp commands subscriptions get-subscription --json`.

## list-subscriptions

List subscriptions

List page-based subscriptions with canonical customer and deprecated provider aliases. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/read-subscriptions) · Effect: **read**

```sh
sellapp subscriptions list-subscriptions
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

Inspect the complete schema: `sellapp commands subscriptions list-subscriptions --json`.

## pause

Pause a subscription

Pause a subscription after checking capabilities. Automatic resume dates are supported only by the current Stripe flow; use a future resume_at when supplied. Other provider pauses may need an explicit resume action. Customer-originated pauses are also subject to store pause policies. Reuse an idempotency key only with identical input and actor, and retrieve the subscription after the action. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/pause-a-subscription) · Effect: **consequential**

```sh
sellapp subscriptions pause 55 --resume-at 2026-10-10T12:00:00Z --reason 'Customer request' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --resume-at | No | Use a future date-time at execution time. Automatic resume dates are supported only for Stripe; the shown date is illustrative. |
| --reason | No | Value for reason. |

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

Inspect the complete schema: `sellapp commands subscriptions pause --json`.

## preview-plan-change

Preview a subscription plan change

This action belongs to the customer. Seller API requests return a 422 validation error. Check subscription capabilities and use the customer flow; never expose a seller API key to work around this restriction. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/preview-plan-change) · Effect: **consequential**

```sh
sellapp subscriptions preview-plan-change 1 --target-variant-id 4321 --effective-timing immediate --proration-behavior provider_default --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --target-variant-id | Yes | Value for target variant id. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --effective-timing | No | Value for effective timing. |
| --proration-behavior | No | Value for proration behavior. |
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

Inspect the complete schema: `sellapp commands subscriptions preview-plan-change --json`.

## preview-renewal-date

Preview a subscription renewal date change

Preview a new renewal date before confirmation when the subscription's billing-date policy requires a preview. The current Stripe and PayPal provider flows do not support renewal-date changes and return 422. Check shift_billing_date capabilities before sending a request. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/preview-renewal-date-change) · Effect: **consequential**

```sh
sellapp subscriptions preview-renewal-date 1 --renewal-date 2026-10-01T12:00:00Z --reason 'Align Maya'\''s membership with the monthly reading circle.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --renewal-date | Yes | Value for renewal date. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --reason | No | Value for reason. |
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

Inspect the complete schema: `sellapp commands subscriptions preview-renewal-date --json`.

## resume

Resume a subscription

Resume a paused subscription and restore related entitlements where possible. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/resume-a-subscription) · Effect: **consequential**

```sh
sellapp subscriptions resume 55 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |

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

Inspect the complete schema: `sellapp commands subscriptions resume --json`.

## search-subscriptions

Search subscriptions

Search while preserving page-based pagination. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/read-subscriptions) · Effect: **read**

```sh
sellapp subscriptions search-subscriptions --search 'maya.chen@example.com' --status active
```

| Input | Required | Meaning |
| --- | --- | --- |
| --search | No | Value for search. |
| --status | No | Value for status. |

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

Inspect the complete schema: `sellapp commands subscriptions search-subscriptions --json`.

## update-payment-method

Update a subscription payment method

This action belongs to the customer. Seller API requests return a 422 validation error. Check subscription capabilities and use the customer flow; never expose a seller API key to work around this restriction. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/update-payment-method) · Effect: **consequential**

```sh
sellapp subscriptions update-payment-method 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |

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

Inspect the complete schema: `sellapp commands subscriptions update-payment-method --json`.

