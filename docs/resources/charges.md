# charges

[All commands](../commands.md)

## create

Create a charge

Create a standalone charge for payment or a free claim. Requires the charge API ability and store permission. The total is an integer in the currency's minor units: 1999 USD means $19.99. Supplying currency requires total. For a positive paid charge, choose enabled payment_method/payment_methods or use_all_payment_methods; availability depends on the store and currency. CUSTOM_PAYMENT_METHOD is supported: preselection requires custom_payment_method_id. custom_payment_method_ids optionally restricts the available custom methods; omission snapshots the currently usable IDs for this store, without product assignments. Customer proof or confirmation changes custom charges to REVIEW, not COMPLETED. Creating a charge does not prove payment. A repeated create may create another charge; check the original result before retrying. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/charges/create-a-charge) · Effect: **consequential**

```sh
sellapp charges create --email 'sofia.rivera@example.com' --return-url https://example.com/launch-complete --currency USD --total 10000 --payment-method PAYPAL --reference 'One more thing launch' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --email | Yes | Value for email. |
| --return-url | Yes | Value for return url. |
| --cancel-url | No | Value for cancel url. |
| --webhook | No | Optional public webhook destination validated for safe delivery. A browser return is not payment confirmation. |
| --reference | No | Value for reference. |
| --description | No | Value for description. |
| --currency | No | Supported three-letter currency code, normalized to uppercase. Defaults to USD when omitted. When supplied, total is required; supply both for a paid charge. |
| --total | No | Integer minor units in currency; 1999 USD is $19.99. Required when currency is supplied. Maximum is the minor-unit equivalent of 99,999,999 major currency units. Omit total and currency together for a free USD claim. |
| --payment-method | No | Value for payment method. |
| --payment-methods | No | Value for payment methods. |
| --custom-payment-method-id | No | Required when payment_method is CUSTOM_PAYMENT_METHOD; otherwise omit it. The ULID must identify an enabled, usable method in this store and belong to custom_payment_method_ids when provided. |
| --custom-payment-method-ids | No | Optional subset of enabled, usable custom method ULIDs for this charge. Include CUSTOM_PAYMENT_METHOD among the allowed gateways. Omit to snapshot all currently usable methods; newly created methods are not added later. Product assignments do not apply. Null and an empty array are invalid. |
| --use-all-payment-methods | No | Value for use all payment methods. |
| --deliverable | No | Value for deliverable. |
| --metadata | No | Value for metadata. |
| --coupon-code | No | Value for coupon code. |

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

Inspect the complete schema: `sellapp commands charges create --json`.

## get

Retrieve a charge

Retrieve a charge by its ID to check its details and payment state. The response schema describes the returned fields. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/charges/retrieve-a-charge) · Effect: **read**

```sh
sellapp charges get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| charge (positional) | Yes | The charge path parameter. |

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

Inspect the complete schema: `sellapp commands charges get --json`.

## list

List all charges

Retrieve a paginated list of charges for your store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/charges/list-all-charges) · Effect: **read**

```sh
sellapp charges list
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

Inspect the complete schema: `sellapp commands charges list --json`.

## mark-completed

Mark pending charge completed

Manually complete a PENDING or VOIDED ordinary charge, or a custom-payment charge in REVIEW, after independently verifying receipt of funds. This does not collect, capture, or verify payment. Initialized custom-payment wallet top-ups may be approved from PENDING or REVIEW and credit the deposit plus any snapshotted bonus exactly once. Other wallet gateways cannot be completed manually, and voided custom top-ups cannot be revived. Completion emits charge.completed webhooks and store notifications and records applicable platform fees. Requires the charge API ability and update permission. A charge already completed returns 422; retrieve its state before retrying. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/charges/mark-pending-charge-completed) · Effect: **consequential**

```sh
sellapp charges mark-completed 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| charge_id (positional) | Yes | The charge id path parameter. |

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

Inspect the complete schema: `sellapp commands charges mark-completed --json`.

## mark-voided

Mark pending charge voided

Void a PENDING charge or a custom-payment charge in REVIEW. This cancels the internal charge, releases an applicable redeemed reward coupon, and emits charge.voided; it does not refund a payment. Requires the charge API ability and update permission. Other states return 422, including a repeat after successful voiding; retrieve its state before retrying. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/charges/mark-pending-charge-voided) · Effect: **consequential**

```sh
sellapp charges mark-voided 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| charge_id (positional) | Yes | The charge id path parameter. |

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

Inspect the complete schema: `sellapp commands charges mark-voided --json`.

