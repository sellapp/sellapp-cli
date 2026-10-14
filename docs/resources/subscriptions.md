# subscriptions

[All commands](../commands.md)

## actions dismiss

Dismiss a subscription action

Dismiss a pending lifecycle action that the provider left unresolved, which frees the subscription for its next change. Only a pending action can be dismissed, and the required `note` is stored on the action as its `error` so the decision stays on the record. An action in any other state returns `422` with the refusal sentence under `errors.subscription`, while a missing or over-long `note` returns an ordinary field error. This operation does not replay completed dismissals: idempotency keys are ignored. After a lost response, retrieve the action before repeating the request; an already-dismissed action returns `422`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/dismiss-a-subscription-action) · Effect: **consequential**

```sh
sellapp subscriptions actions dismiss --product-subscription 124 4207 --note 'Confirmed with the customer that they no longer want the upgrade.' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product-subscription | Yes | The productSubscription path parameter. |
| action (positional) | Yes | The action path parameter. |
| --note | Yes | Why the action is being dismissed. The note is kept on the dismissed action as its error. |
| --body-idempotency-key | No | Accepted for compatibility but ignored by this operation. Retrieve the action after an uncertain response; repeating a completed dismissal returns 422. |

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

Inspect the complete schema: `sellapp commands subscriptions actions dismiss --json`.

## actions get

Retrieve a subscription action

Retrieve one recorded lifecycle action. An action is only ever reachable through the subscription that recorded it, so an action belonging to another subscription of the same store is not found either. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/retrieve-a-subscription-action) · Effect: **read**

```sh
sellapp subscriptions actions get --product-subscription 124 4192
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product-subscription | Yes | The productSubscription path parameter. |
| action (positional) | Yes | The action path parameter. |

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

Inspect the complete schema: `sellapp commands subscriptions actions get --json`.

## actions list

List subscription actions

List the lifecycle actions recorded against one subscription, ordered by identifier descending so the newest action comes first. The order is fixed and cannot be steered from the request. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/list-subscription-actions) · Effect: **read**

```sh
sellapp subscriptions actions list 124
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
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

Inspect the complete schema: `sellapp commands subscriptions actions list --json`.

## cancel

Cancel a subscription

Cancel an active subscription now or at period end, or one with a payment still owed now: a past due one the store bills in house, or a past due or unpaid Stripe, PayPal, or Razorpay subscription. cancel_at_period_end is required. Refund flags are valid only when it is false; period-end cancellation does not refund the last payment. New integrations should use the cancel-at-period-end and cancel-immediately lifecycle actions instead. A provider refusal, or a state check such as a change already in progress or a subscription that is no longer active, returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. Refunds are off by default. refund_last_payment requests a real refund; pro_rated_refund requires refund_last_payment and defaults to false. While a Stripe, PayPal, or Razorpay subscription is past due or unpaid, its current period was never paid for, so a pro-rated refund returns `422`; a full refund or no refund still works. Cancellation can succeed even if the requested refund fails. A 200 response or confirmed cancellation does not prove a refund, and the action response does not expose the provider refund result. Check the order and payment-provider refund records before retrying; do not replay cancellation as a refund-recovery step. Reuse an idempotency key only with identical action input and actor. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/cancel-a-subscription-immediately-with-a-refund) · Effect: **consequential**

```sh
sellapp subscriptions cancel 55 --cancel-at-period-end true --body-idempotency-key design-kit-subscription-55-cancel-v1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| subscription (positional) | Yes | The subscription path parameter. |
| --cancel-at-period-end | Yes | Value for cancel at period end. |
| --refund-last-payment | No | Request a real refund of the latest payment during immediate cancellation. Cancellation success does not prove refund success; verify provider/order refund records. |
| --pro-rated-refund | No | Request a prorated refund. Requires refund_last_payment=true. False requests a full latest-payment refund when refund_last_payment is enabled. Refused with `422` while a Stripe, PayPal, or Razorpay subscription is past due or unpaid. |
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

Schedule cancellation at the end of the current billing period without requesting a refund. Inspect capabilities first. PayPal has no cancellation at period end, so on a PayPal subscription this cancels it with PayPal straight away and keeps the customer's access until the period ends; it can't be kept afterwards. The response records the cancellation action; retrieve the subscription to observe its current state. Reuse an idempotency key only with identical action input and actor. A provider refusal, or a state check such as a change already in progress or a subscription that is no longer active, returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Request immediate subscription cancellation. Inspect capabilities for provider support. This ends future billing and can change customer access. It is the one action left on a Stripe or PayPal subscription that is past due or unpaid, and one of two left on a Razorpay one alongside a payment method update. A provider refusal, or a state check such as a change already in progress or a subscription that is no longer active, returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. Refunds are off by default. refund_last_payment requests a real refund; pro_rated_refund requires refund_last_payment and defaults to false. While a Stripe, PayPal, or Razorpay subscription is past due or unpaid, its current period was never paid for, so a pro-rated refund returns `422`; a full refund or no refund still works. Cancellation can succeed even if the requested refund fails. A 200 response or confirmed cancellation does not prove a refund, and the action response does not expose the provider refund result. Check the order and payment-provider refund records before retrying; do not replay cancellation as a refund-recovery step. Reuse an idempotency key only with identical action input and actor. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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
| --pro-rated-refund | No | Request a prorated refund. Requires refund_last_payment=true. False requests a full latest-payment refund when refund_last_payment is enabled. Refused with `422` while a Stripe, PayPal, or Razorpay subscription is past due or unpaid. |

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

Confirm a previously previewed subscription plan change, sending back the `preview_token` for a preview of the same change. A plan change belongs to the customer, and a store token confirms it on their behalf; a provider refusal, or a state check such as a change already in progress or a subscription that is no longer active, returns `422` with the refusal sentence under `errors.subscription`. `approval_url` accompanies a `pending` action whenever the customer still has to approve the revision at the provider or pay the proration order `proration_order_id` names, and is left out entirely when the provider applied the change itself. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/confirm-plan-change) · Effect: **consequential**

```sh
sellapp subscriptions confirm-plan-change 55 --target-variant-id 4321 --preview-token subprev_9c4b2f --yes
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

Confirm a subscription renewal date change, sending a `preview_token` when the variant's billing-date policy requires previews. No payment method supports renewal-date changes today: Stripe, PayPal, Razorpay, and subscriptions the store bills in house all report `shift_billing_date` as `unsupported`, so this returns `422` with the refusal sentence under `errors.subscription`. Read the capabilities response first. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Return the lifecycle actions currently supported for a subscription, each with the reason the customer and the seller are shown. A `customer_only` status names who owns the action rather than who may call the API: a store token carries a plan change and a payment method update on the customer's behalf, while a `customer_only` auto-renewal switch stays with the customer portal. `set_auto_renew` has no API route at all. While a Stripe, PayPal, or Razorpay subscription is past due or unpaid, every action except `cancel_immediately` and, on Razorpay, `update_payment_method`, is `unsupported`, and a subscription the store bills in house that is on hold for an unpaid renewal reports `resume` as `customer_only`, because only the customer can pay to resume it. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Retrieve one subscription by its SellApp subscription record identifier. Requires the `invoice` Sanctum ability. Foreign-store identifiers return not found. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/read-subscriptions) · Effect: **read**

```sh
sellapp subscriptions get-subscription 123
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

Inspect the complete schema: `sellapp commands subscriptions get-subscription --json`.

## keep

Keep a subscription

Stop a scheduled cancellation so the subscription keeps renewing, which is the undo for cancelling at period end. The subscription has to still be active, still be scheduled to cancel, and its period must not have ended yet. Any of those, a provider that cannot undo its own schedule (a PayPal subscription or a Razorpay trial, both already cancelled with the provider), or a change already in progress returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/keep-a-subscription) · Effect: **consequential**

```sh
sellapp subscriptions keep 55 --yes
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

Inspect the complete schema: `sellapp commands subscriptions keep --json`.

## list-subscriptions

List subscriptions

List the subscriptions belonging to the authenticated store, ordered by identifier ascending. Requires the `invoice` Sanctum ability. Use `status[]` to keep only the states you care about. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/read-subscriptions) · Effect: **read**

```sh
sellapp subscriptions list-subscriptions
```

| Input | Required | Meaning |
| --- | --- | --- |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --pagination | No | Set to false to return at most 100 matching items in one data array without pagination links or metadata. |
| --status | No | Filter by one or more subscription statuses. Repeat the parameter to select several; a single `status=ACTIVE` works too. A status outside the list is rejected with a validation error. |

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

Pause a subscription after checking capabilities. Automatic resume dates are supported by Stripe and by subscriptions the store bills in house; use a future resume_at when supplied. Other provider pauses may need an explicit resume action. Customer-originated pauses are also subject to store pause policies. Reuse an idempotency key only with identical input and actor, and retrieve the subscription after the action. A provider refusal, or a state check such as a change already in progress or a subscription that is no longer active, returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/pause-a-subscription) · Effect: **consequential**

```sh
sellapp subscriptions pause 55 --resume-at 2026-10-10T12:00:00Z --reason 'Customer request' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --resume-at | No | Use a future date-time at execution time. Automatic resume dates are supported by Stripe and by subscriptions the store bills in house; the shown date is illustrative. |
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

Preview a subscription plan change before confirming it. The preview stores the change it was taken for and returns a `preview_token` the confirmation must send back, alongside the provider's own pricing preview. On Stripe and in-house subscriptions, `preview_payload.payment_required` says whether confirming charges `amount_due` today or leaves any amount to the next invoice, and `preview_payload.renewal_date` says when the subscription renews next. A plan change belongs to the customer, and a store token previews it on their behalf; a provider refusal, or a state check such as a change already in progress or a subscription that is no longer active, returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Preview a new renewal date before confirmation when the subscription's billing-date policy requires one. No payment method supports renewal-date changes today: Stripe, PayPal, Razorpay, and subscriptions the store bills in house all report `shift_billing_date` as `unsupported`, so this returns `422` with the refusal sentence under `errors.subscription`. Read the capabilities response first. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Resume a paused subscription and restore related entitlements where possible. A subscription that is not paused, a subscription the store bills in house that is on hold for an unpaid renewal (only the customer can resume it, by paying that renewal from the customer portal), a provider refusal, or a change already in progress returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Search the store's subscriptions, ordered by identifier ascending unless `sort` says otherwise. `search.value` matches the provider subscription ID, the provider customer ID, and the customer email. Filter on `id`, `customer_id` (the SellApp customer ID, as in `customer.id`), `provider`, `status`, or `product_variant_id`, and sort on `id`, `status`, `current_period_start`, `current_period_end`, `created_at`, or `updated_at`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/read-subscriptions) · Effect: **read**

```sh
sellapp subscriptions search-subscriptions --body '{"search":{"value":"maya.chen@example.com"},"filters":[{"field":"status","operator":"=","value":"ACTIVE"}]}'
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

Inspect the complete schema: `sellapp commands subscriptions search-subscriptions --json`.

## update-payment-method

Update a subscription payment method

Request a subscription payment method update. A payment method belongs to the customer, and a store token requests the update on their behalf. The customer still enters the new card themselves, so the action stays `pending` and `redirect_url` is where they finish it: Stripe's hosted page for a Stripe subscription, or a SellApp page for a Razorpay one. `return_url` chooses where Stripe sends the customer afterwards; Razorpay ignores it. PayPal subscriptions and subscriptions the store bills in house don't support the update. Read `update_payment_method` in the capabilities response before calling: an unsupported payment method, or a state check such as a change already in progress, a subscription that is no longer active, or a Stripe subscription with a payment still owed, returns `422` with the refusal sentence under `errors.subscription`. Requires the `invoice` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/subscriptions/update-payment-method) · Effect: **consequential**

```sh
sellapp subscriptions update-payment-method 55 --return-url https://store.example.com/billing/updated --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| productSubscription (positional) | Yes | The productSubscription path parameter. |
| --body-idempotency-key | No | Optional idempotency key. You may also send this as the Idempotency-Key header. |
| --return-url | No | Where Stripe returns the customer once its hosted page finishes. Defaults to the store's customer portal page for this subscription. Razorpay subscriptions ignore it. |

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

