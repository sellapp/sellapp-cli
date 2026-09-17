# product-variants

[All commands](../commands.md)

## batch-create

Batch create product variants

Create multiple variants for a product in one request by sending a `resources` array. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/batch-create-product-variants) · Effect: **consequential**

```sh
sellapp product-variants batch-create 1 --body '{"resources":[{"title":"Default","description":"Default product variant.","deliverable":{"types":["TEXT"],"data":{"serials":["SERIAL-001"],"parsingMode":"NEW_LINE","removeDuplicate":true}},"pricing":{"humble":false,"price":{"price":1000,"currency":"USD"}},"payment_methods":["PAYPAL"]}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands product-variants batch-create --json`.

## batch-delete

Batch delete product variants

Delete multiple product variants in one request by sending the variant IDs in `resources`. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/batch-delete-product-variants) · Effect: **consequential**

```sh
sellapp product-variants batch-delete 1 --body '{"resources":[1,2]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands product-variants batch-delete --json`.

## batch-update

Batch update product variants

Update multiple product variants in one request by sending a `resources` object keyed by variant ID. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/batch-update-product-variants) · Effect: **consequential**

```sh
sellapp product-variants batch-update 1 --body '{"resources":{"1":{"title":"Updated variant"}}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --resources | Yes | Value for resources. |

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

Inspect the complete schema: `sellapp commands product-variants batch-update --json`.

## booking create-hold

Create a booking hold

Temporarily reserve a booking slot before checkout. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/create-booking-hold) · Effect: **consequential**

```sh
sellapp product-variants booking create-hold --product 1 1 --body '{"slot_start_at":"2026-06-22T14:00:00+00:00","quantity":1,"customer_key":"visitor-session-123","meta":{"customer_timezone":"America/New_York"}}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --slot-start-at | Yes | Value for slot start at. |
| --quantity | No | Value for quantity. |
| --customer-key | No | Value for customer key. |
| --meta | No | Value for meta. |

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

Inspect the complete schema: `sellapp commands product-variants booking create-hold --json`.

## booking get

Retrieve booking configuration

Retrieve creator-facing availability, capacity, calendar, meeting, and reminder configuration for a booking variant. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/retrieve-booking-configuration) · Effect: **read**

```sh
sellapp product-variants booking get --product 41 73
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The booking product ID or slug. |
| variant (positional) | Yes | The variant path parameter. |

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

Inspect the complete schema: `sellapp commands product-variants booking get --json`.

## booking list-availability

List booking availability

Retrieve available appointment slots for a booking product variant. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/list-booking-availability) · Effect: **read**

```sh
sellapp product-variants booking list-availability --product 1 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --from | No | Start date for the availability window. |
| --to | No | End date for the availability window. |
| --quantity | No | Number of seats to reserve for each returned slot. |

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

Inspect the complete schema: `sellapp commands product-variants booking list-availability --json`.

## booking release-hold

Release a booking hold

Release a previously created booking hold. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/release-booking-hold) · Effect: **consequential**

```sh
sellapp product-variants booking release-hold --product 1 test_hold --customer-key visitor-session-123 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| hold (positional) | Yes | The hold path parameter. |
| --customer-key | No | Value for customer key. |

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

Inspect the complete schema: `sellapp commands product-variants booking release-hold --json`.

## booking replace

Update booking configuration

Partially update a booking variant's configuration. Provider connection IDs must belong to the selected store; secret credentials are never accepted or returned here. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/update-booking-configuration) · Effect: **consequential**

```sh
sellapp product-variants booking replace --product 41 73 --timezone Europe/London --duration-minutes 60 --capacity-per-slot 1 --min-notice-minutes 1440 --max-advance-days 60 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The booking product ID or slug. |
| variant (positional) | Yes | The variant path parameter. |
| --mode | No | Value for mode. |
| --conflict-scope | No | Value for conflict scope. |
| --timezone | No | Value for timezone. |
| --duration-minutes | No | Value for duration minutes. |
| --capacity-per-slot | No | Value for capacity per slot. |
| --min-notice-minutes | No | Value for min notice minutes. |
| --max-advance-days | No | Value for max advance days. |
| --buffer-before-minutes | No | Value for buffer before minutes. |
| --buffer-after-minutes | No | Value for buffer after minutes. |
| --availability | No | Value for availability. |
| --provider-connection-ids | No | Value for provider connection ids. |
| --video-provider | No | Value for video provider. |
| --video-provider-connection-id | No | Value for video provider connection id. |
| --reminders-enabled | No | Value for reminders enabled. |
| --reminder-offset-value | No | Value for reminder offset value. |
| --reminder-offset-unit | No | Value for reminder offset unit. |
| --meta | No | Value for meta. |

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

Inspect the complete schema: `sellapp commands product-variants booking replace --json`.

## booking update

Update booking configuration

Partially update a booking variant's configuration. Provider connection IDs must belong to the selected store; secret credentials are never accepted or returned here. Requires the `listing` Sanctum ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/bookings/update-booking-configuration) · Effect: **consequential**

```sh
sellapp product-variants booking update --product 41 73 --timezone Europe/London --duration-minutes 60 --capacity-per-slot 1 --min-notice-minutes 1440 --max-advance-days 60 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The booking product ID or slug. |
| variant (positional) | Yes | The variant path parameter. |
| --mode | No | Value for mode. |
| --conflict-scope | No | Value for conflict scope. |
| --timezone | No | Value for timezone. |
| --duration-minutes | No | Value for duration minutes. |
| --capacity-per-slot | No | Value for capacity per slot. |
| --min-notice-minutes | No | Value for min notice minutes. |
| --max-advance-days | No | Value for max advance days. |
| --buffer-before-minutes | No | Value for buffer before minutes. |
| --buffer-after-minutes | No | Value for buffer after minutes. |
| --availability | No | Value for availability. |
| --provider-connection-ids | No | Value for provider connection ids. |
| --video-provider | No | Value for video provider. |
| --video-provider-connection-id | No | Value for video provider connection id. |
| --reminders-enabled | No | Value for reminders enabled. |
| --reminder-offset-value | No | Value for reminder offset value. |
| --reminder-offset-unit | No | Value for reminder offset unit. |
| --meta | No | Value for meta. |

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

Inspect the complete schema: `sellapp commands product-variants booking update --json`.

## create

Create a product variant

Create a product variant for the authenticated store. Course, credits, and add-on products permit only one variant. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/create-a-product-variant) · Effect: **write**

```sh
sellapp product-variants create 120 --body '{"title":"Monthly membership","description":"One operating memo each month; access is provisioned by our team.","deliverable":{"types":["MANUAL"],"data":{"stock":null,"comment":"We will send your reading-room invitation."}},"pricing":{"humble":false,"price":{"price":1999,"currency":"USD"}},"payment_methods":["STRIPE"]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --title | Yes | Value for title. |
| --description | Yes | Value for description. |
| --deliverable | Yes | Value for deliverable. |
| --pricing | Yes | Value for pricing. |
| --payment-methods | Yes | Value for payment methods. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --bulk-discount | No | Value for bulk discount. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot for updates. Use the variant's latest updated_at value; stale snapshots receive a 422 validation response. |

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

Inspect the complete schema: `sellapp commands product-variants create --json`.

## delete

Delete a product variant

Deletes a product variant. This will permanently delete the product variant and all of its details, including the sensitive deliverables, will no longer be retrievable. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/delete-a-product-variant) · Effect: **consequential**

```sh
sellapp product-variants delete --product 1 2 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |

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

Inspect the complete schema: `sellapp commands product-variants delete --json`.

## get

Retrieve a product variant

Retrieve a variant using its product ID and variant ID. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/retrieve-a-product-variant) · Effect: **read**

```sh
sellapp product-variants get --product 1 2
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --with-drafts | No | Include draft variants alongside published variants. |
| --only-drafts | No | Return only draft variants. |

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

Inspect the complete schema: `sellapp commands product-variants get --json`.

## list

List all product variants

List a product's variants, 15 per page by default. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/list-all-product-variants) · Effect: **read**

```sh
sellapp product-variants list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --with-drafts | No | Include draft variants alongside published variants. |
| --only-drafts | No | Return only draft variants. |

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

Inspect the complete schema: `sellapp commands product-variants list --json`.

## pricing replace

Replace product variant pricing

Replace a variant's pricing and payment-method selection. Omitted `compare_at_price` and `bulk_discount` values are cleared. An empty payment-method array is accepted only when the resulting published variant is genuinely free, including having no paid credit-rate tiers. Requires the `listing` Sanctum ability. Relationship IDs are scoped to the current store. Use the ISO 8601 string `expected_updated_at` to reject stale writes. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp product-variants pricing replace --product 120 4321 --body '{"pricing":{"type":"SUBSCRIPTION","humble":false,"price":{"price":1999,"currency":"USD"},"frequency":{"value":1,"interval":"MONTH"}},"payment_methods":["STRIPE"]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --pricing | Yes | Value for pricing. |
| --payment-methods | Yes | Complete payment-method selection. May be empty only when the resulting published variant is genuinely free, including having no paid credit-rate tiers. |
| --custom-payment-method-ids | No | Enabled custom payment method ULIDs owned by the current store. Required when CUSTOM_PAYMENT_METHOD is selected. |
| --bulk-discount | No | Quantity pricing tiers for one-time variants. Subscription variants cannot use these tiers. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot. The token advances for pricing and custom payment-method selection changes; a stale value returns 422 without changing either selection. |

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

Inspect the complete schema: `sellapp commands product-variants pricing replace --json`.

## pricing update

Partially update product variant pricing

Update only the supplied pricing, subscription cadence, quantity pricing tier, or payment-method fields; omitted fields retain their current values. An empty payment-method array is accepted only when the resulting published variant is genuinely free, including having no paid credit-rate tiers. Requires the `listing` Sanctum ability. Relationship IDs are scoped to the current store. Use the ISO 8601 string `expected_updated_at` to reject stale writes. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/update-product-variant-pricing) · Effect: **write**

```sh
sellapp product-variants pricing update --product 120 4321 --body '{"pricing":{"price":{"price":2499,"currency":"USD"}}}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --pricing | No | Value for pricing. |
| --payment-methods | No | Complete payment-method selection. May be empty only when the resulting published variant is genuinely free, including having no paid credit-rate tiers. |
| --custom-payment-method-ids | No | Enabled custom payment method ULIDs owned by the current store. Required when CUSTOM_PAYMENT_METHOD is selected. |
| --bulk-discount | No | Quantity pricing tiers for one-time variants. Subscription variants cannot use these tiers. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot. The token advances for pricing and custom payment-method selection changes; a stale value returns 422 without changing either selection. |

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

Inspect the complete schema: `sellapp commands product-variants pricing update --json`.

## replace

Update a product variant with PUT

Update a product variant using the PUT compatibility method. Fields remain partial for backward compatibility. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp product-variants replace --product 120 4321 --title 'Monthly membership plus' --description 'One annotated operating memo and a monthly founder discussion.'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --title | No | Value for title. |
| --description | No | Value for description. |
| --deliverable | No | Value for deliverable. |
| --pricing | No | Value for pricing. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --bulk-discount | No | Value for bulk discount. |
| --payment-methods | No | Value for payment methods. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot for updates. Use the variant's latest updated_at value; stale snapshots receive a 422 validation response. |

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

Inspect the complete schema: `sellapp commands product-variants replace --json`.

## search

Search product variants

Search a product's variants using JSON body filters, search terms, includes, and sort instructions. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/search-product-variants) · Effect: **read**

```sh
sellapp product-variants search 1 --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
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

Inspect the complete schema: `sellapp commands product-variants search --json`.

## update

Update a product variant

Updates a product variant's details. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/update-a-product-variant) · Effect: **write**

```sh
sellapp product-variants update --product 120 4321 --title 'Monthly membership plus' --description 'One annotated operating memo and a monthly founder discussion.'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --title | No | Value for title. |
| --description | No | Value for description. |
| --deliverable | No | Value for deliverable. |
| --pricing | No | Value for pricing. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --bulk-discount | No | Value for bulk discount. |
| --payment-methods | No | Value for payment methods. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot for updates. Use the variant's latest updated_at value; stale snapshots receive a 422 validation response. |

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

Inspect the complete schema: `sellapp commands product-variants update --json`.

