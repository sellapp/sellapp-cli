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

## modifier get

Retrieve a payment method modifier

Get the discount or fee customers get at checkout when they pay with this payment method: a percentage of the order total, a fixed amount, or both. `kind` is `none` when there is neither. `condition` and `threshold` show when it applies only to checkouts above or below a total, and are `null` when every checkout gets it. Requires the `payment-method` or `store` API-key ability, with the same permission in your team role. An unknown payment method returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/payment-method-modifier) · Effect: **read**

```sh
sellapp payment-methods modifier get STRIPE
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

Inspect the complete schema: `sellapp commands payment-methods modifier get --json`.

## modifier replace

Update a payment method modifier

Set the discount or fee customers get at checkout when they pay with this payment method. Send `kind` with positive amounts: `fee` adds them to the price and `discount` takes them off. A discount or fee needs `percentage`, `fixed_amount` or both; send `kind: none` with both `null` to remove it. Both amount fields are required and may be `null`, and the amounts you send replace the current ones. The condition works differently: leave out both `condition` and `threshold` to keep the condition the method has now, even one changed in the dashboard while your request was in flight; send `condition: null` to remove it so every checkout gets the discount or fee; send `condition` with `threshold` to set it. A `threshold` without a `condition`, or a `condition` with `kind: none`, returns `422` on `condition`. The method must be turned on for checkout, otherwise `422` on `payment_method`. A percentage over 100 returns `422` on `percentage`, and a fixed amount over 100000 returns `422` on `fixed_amount`. Sending the modifier the method already has returns `200` and changes nothing, so a retry is safe; the `data` of `GET /v2/payment-methods/{paymentMethod}/modifier` without `payment_method` is such a request. Requires the `store` API-key ability and the `store` permission in your team role. An unknown payment method returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/payment-methods/payment-method-modifier) · Effect: **consequential**

```sh
sellapp payment-methods modifier replace STRIPE --kind fee --percentage 3.00 --fixed-amount 0.30 --condition below --threshold 20.00 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| paymentMethod (positional) | Yes | Uppercase payment method identifier returned by the collection endpoint. |
| --kind | Yes | `discount` lowers the price customers pay with this method, `fee` raises it, and `none` means neither. |
| --percentage | Yes | Percentage of the order total, as a positive decimal string with at most two decimals: `"3.00"` is 3%. At most 100. `null` when not used. |
| --fixed-amount | Yes | Fixed amount added or taken off once per checkout, as a positive decimal string in the checkout's currency with at most two decimals: `"0.30"` on a USD checkout is $0.30. At most 100000. `null` when not used. |
| --condition | No | Optional. `above` or `below` gives the discount or fee only to checkouts whose total is more than or less than `threshold`. `null` removes the condition, so every checkout gets it. Leave out both `condition` and `threshold` to keep the condition the method has now. A `condition` with `kind: none` returns `422` on `condition`. |
| --threshold | No | Required when `condition` is `above` or `below`: the checkout total to compare against, in the checkout's own currency, as a number or decimal string with at most two decimals from `0.01` to `99999999`. `20` and `"20.00"` mean the same. Leave it out or send `null` otherwise: a `threshold` without a `condition` returns `422` on `condition`. |

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

Inspect the complete schema: `sellapp commands payment-methods modifier replace --json`.

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

