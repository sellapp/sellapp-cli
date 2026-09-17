# payment-methods

[All commands](../commands.md)

## connect

Create a payment connection handoff

Return a provider-hosted connect URL and a status URL to poll. Currently supported for Stripe and PayPal; unsupported methods return a validation error. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods connect STRIPE --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| paymentMethod (positional) | Yes | Uppercase payment method identifier returned by the collection endpoint. |

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

Inspect the complete schema: `sellapp commands payment-methods connect --json`.

## custom create

Create a custom payment method

Create a store-scoped custom payment method. Requires payment-method management permission. The effective redirect type needs a public redirect_url; instructions type needs nonblank instructions or a nonempty step. Switching to instructions clears redirect options; switching to redirect clears instructions, steps and proof-of-payment requirements. This API does not assign the method to every product. Inspect existing checkout configuration before changing an enabled method; creates have no idempotency guarantee. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-custom-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods custom create --body '{"type":"instructions","name":"Manual payment","instructions":"Contact Launch Lab before sending a payment.","enabled":false,"modifier":{"percentage":"-2.50","fixed":"-1.00"}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --type | Yes | Value for type. |
| --name | Yes | Value for name. |
| --description | No | Value for description. |
| --instructions | No | For effective type instructions, provide nonblank instructions or at least one nonblank step. Switching to redirect clears this value. |
| --steps | No | Value for steps. |
| --redirect-url | No | Required for effective type redirect. Must pass public destination validation; existing unchanged destinations retain their established value. |
| --skip-interstitial-page | No | Value for skip interstitial page. |
| --show-processing-status-page | No | Value for show processing status page. |
| --require-proof-of-payment | No | Value for require proof of payment. |
| --enabled | No | Value for enabled. |
| --sort-order | No | Value for sort order. |
| --modifier | No | Value for modifier. |

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

Inspect the complete schema: `sellapp commands payment-methods custom create --json`.

## custom delete

Delete a custom payment method

Delete one store-scoped custom payment method and remove its product assignments. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-custom-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods custom delete 00000000000000000000000000 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customPaymentMethod (positional) | Yes | Store-scoped custom payment method ULID. |

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

Inspect the complete schema: `sellapp commands payment-methods custom delete --json`.

## custom get

Retrieve a custom payment method

Retrieve one custom payment method owned by the current store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-custom-payment-methods) · Effect: **read**

```sh
sellapp payment-methods custom get 00000000000000000000000000
```

| Input | Required | Meaning |
| --- | --- | --- |
| customPaymentMethod (positional) | Yes | Store-scoped custom payment method ULID. |

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

Inspect the complete schema: `sellapp commands payment-methods custom get --json`.

## custom list

List custom payment methods

List the current store's instruction and redirect payment methods. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-custom-payment-methods) · Effect: **read**

```sh
sellapp payment-methods custom list
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

Inspect the complete schema: `sellapp commands payment-methods custom list --json`.

## custom replace

Replace a custom payment method

Update a custom payment method with required type and name. Omitted optional fields retain their stored values; this PUT does not reset every omitted field. Requires payment-method management permission. The effective redirect type needs a public redirect_url; instructions type needs nonblank instructions or a nonempty step. Switching to instructions clears redirect options; switching to redirect clears instructions, steps and proof-of-payment requirements. This API does not assign the method to every product. Inspect existing checkout configuration before changing an enabled method; creates have no idempotency guarantee. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-custom-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods custom replace 00000000000000000000000000 --type instructions --name 'Manual payment' --instructions 'Contact Launch Lab before sending a payment.' --enabled false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customPaymentMethod (positional) | Yes | Store-scoped custom payment method ULID. |
| --type | Yes | Value for type. |
| --name | Yes | Value for name. |
| --description | No | Value for description. |
| --instructions | No | For effective type instructions, provide nonblank instructions or at least one nonblank step. Switching to redirect clears this value. |
| --steps | No | Value for steps. |
| --redirect-url | No | Required for effective type redirect. Must pass public destination validation; existing unchanged destinations retain their established value. |
| --skip-interstitial-page | No | Value for skip interstitial page. |
| --show-processing-status-page | No | Value for show processing status page. |
| --require-proof-of-payment | No | Value for require proof of payment. |
| --enabled | No | Value for enabled. |
| --sort-order | No | Value for sort order. |
| --modifier | No | Value for modifier. |

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

Inspect the complete schema: `sellapp commands payment-methods custom replace --json`.

## custom update

Update a custom payment method

Partially update a store-scoped custom payment method; omitted fields retain their values. Requires payment-method management permission. The effective redirect type needs a public redirect_url; instructions type needs nonblank instructions or a nonempty step. Switching to instructions clears redirect options; switching to redirect clears instructions, steps and proof-of-payment requirements. This API does not assign the method to every product. Inspect existing checkout configuration before changing an enabled method; creates have no idempotency guarantee. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-custom-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods custom update 00000000000000000000000000 --enabled false --modifier null --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| customPaymentMethod (positional) | Yes | Store-scoped custom payment method ULID. |
| --type | No | Value for type. |
| --name | No | Value for name. |
| --description | No | Value for description. |
| --instructions | No | For effective type instructions, provide nonblank instructions or at least one nonblank step. Switching to redirect clears this value. |
| --steps | No | Value for steps. |
| --redirect-url | No | Required for effective type redirect. Must pass public destination validation; existing unchanged destinations retain their established value. |
| --skip-interstitial-page | No | Value for skip interstitial page. |
| --show-processing-status-page | No | Value for show processing status page. |
| --require-proof-of-payment | No | Value for require proof of payment. |
| --enabled | No | Value for enabled. |
| --sort-order | No | Value for sort order. |
| --modifier | No | Value for modifier. |

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

Inspect the complete schema: `sellapp commands payment-methods custom update --json`.

## enable

Enable or disable a payment method

Change checkout availability after setup has succeeded. Enabling a disconnected method returns an actionable validation error. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods enable STRIPE --enabled false --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| paymentMethod (positional) | Yes | Uppercase payment method identifier returned by the collection endpoint. |
| --enabled | Yes | Value for enabled. |

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

Inspect the complete schema: `sellapp commands payment-methods enable --json`.

## get

Retrieve payment method status

Poll one payment method's connection, health, required-input, and checkout-enabled state. No stored secret is returned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-payment-methods) · Effect: **read**

```sh
sellapp payment-methods get STRIPE
```

| Input | Required | Meaning |
| --- | --- | --- |
| paymentMethod (positional) | Yes | Uppercase payment method identifier returned by the collection endpoint. |

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

Inspect the complete schema: `sellapp commands payment-methods get --json`.

## list

List payment methods

List safe, pollable setup and checkout status for every payment method available to the store. Provider credentials and wallet addresses are never serialized. Requires the `payment-method` ability or the stronger legacy `store` ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-payment-methods) · Effect: **read**

```sh
sellapp payment-methods list
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

Inspect the complete schema: `sellapp commands payment-methods list --json`.

## validate

Validate and save payment method configuration

Validate provider credentials or crypto wallet input with the same setup services used by the dashboard, then save the configuration. Secret fields are write-only and never appear in responses or audit payloads. Supported methods are NMI, Mercado Pago, Razorpay, Mollie, LiFi, and Bitcart-backed crypto wallets. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/manage-payment-methods) · Effect: **consequential**

```sh
sellapp payment-methods validate NMI --body '{"merchant_secure_key":"replace-with-nmi-secure-key","merchant_tokenization_key":"replace-with-nmi-tokenization-key","signing_key":"replace-with-nmi-signing-key","currencies":["USD"]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| paymentMethod (positional) | Yes | Uppercase payment method identifier returned by the collection endpoint. |
| --merchant-secure-key | No | Write-only secret. SellApp never returns the stored value. |
| --merchant-tokenization-key | No | Write-only secret. SellApp never returns the stored value. |
| --signing-key | No | Write-only secret. SellApp never returns the stored value. |
| --currencies | No | Value for currencies. |
| --public-key | No | Write-only secret. SellApp never returns the stored value. |
| --body-access-token | No | Write-only secret. SellApp never returns the stored value. |
| --webhook-secret | No | Write-only secret. SellApp never returns the stored value. |
| --merchant-id | No | Value for merchant id. |
| --key-id | No | Value for key id. |
| --key-secret | No | Write-only secret. SellApp never returns the stored value. |
| --primary-currency | No | Value for primary currency. |
| --international-payments-enabled | No | Value for international payments enabled. |
| --supported-currencies | No | Value for supported currencies. |
| --body-api-key | No | Write-only secret. SellApp never returns the stored value. |
| --address | No | Write-only secret. SellApp never returns the stored value. |
| --private-view-key | No | Write-only secret. SellApp never returns the stored value. |
| --min-withdraw-amount | No | Value for min withdraw amount. |
| --destination-gateway | No | Value for destination gateway. |

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

Inspect the complete schema: `sellapp commands payment-methods validate --json`.

