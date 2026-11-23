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

## community-access replace

Replace variant community access

Give buyers of a normal product's variant access to a community your store has connected: a Discord server with optional roles, a Telegram group, a Slack channel, or a WhatsApp group. It replaces the variant's access on the platform you send and keeps its access on the other platforms. `community.role_ids` is for Discord only: for other platforms leave it out or send an empty list. Discord roles must belong to the chosen server. Read the saved access on the variant (`GET /v2/products/{product}/variants/{variant}`) under `other_settings.community_servers.<platform>`, as `server_id`, `roles` (the Discord role IDs you sent as `role_ids`), `required`, and `remove_on_expire`. For Discord the same values are also in `other_settings.discord_server`, `discord_roles`, `discord_data_required`, and `discord_kick_on_expire`. Requires the `listing` API-key ability and the `listing` permission in your team role. Send the variant's `updated_at` as `expected_updated_at` to refuse the change when the variant was edited since you read it. Repeating the same body without `expected_updated_at` gives the same result. A server, group, channel, or role your store has not connected returns `422` on `community.server_id` or `community.role_ids`, and so do roles for a platform other than Discord. A stale `expected_updated_at` returns `422` on `expected_updated_at`. A variant of a listing that is not a normal product returns `422` on `community`. A variant of another product returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/replace-variant-community-access) · Effect: **write**

```sh
sellapp product-variants community-access replace --product 1 2 --body '{"community":{"platform":"discord","server_id":"1187040134912155648","role_ids":["1187040299781210112"],"required":true,"remove_on_expire":true},"expected_updated_at":"2024-01-10T18:47:22.000000Z"}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --community | Yes | Value for community. |
| --expected-updated-at | No | The variant's `updated_at` from your last read. The change is refused with `422` when the variant changed since. |

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

Inspect the complete schema: `sellapp commands product-variants community-access replace --json`.

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
| --payment-methods | Yes | Complete gateway selection for this variant. Selecting CUSTOM_PAYMENT_METHOD also requires at least one custom_payment_method_ids entry. |
| --minimum-purchase-quantity | No | Value for minimum purchase quantity. |
| --maximum-purchase-quantity | No | Value for maximum purchase quantity. |
| --bulk-discount | No | Value for bulk discount. |
| --custom-payment-method-ids | No | Enabled, usable custom payment method ULIDs belonging to the selected store. Supply at least one when payment_methods includes CUSTOM_PAYMENT_METHOD. Include CUSTOM_PAYMENT_METHOD in payment_methods whenever sending nonempty IDs. On updates, omitting both selections preserves the current methods. Nonempty IDs are rejected when custom payments are not selected. |
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

Soft-delete a variant in the authenticated store. Variants referenced by non-terminal subscriptions cannot be deleted. Send expected_updated_at from the latest variant response to reject a stale deletion with 422. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/delete-a-product-variant) · Effect: **consequential**

```sh
sellapp product-variants delete --product 1 2 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot. Send the exact `updated_at` value from the product you last retrieved; the mutation is rejected if the product has changed. |

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

## duplicate

Duplicate a product variant

Copy a variant of a normal product onto the same product and return the copy. The copy keeps every setting, the title, and the draft state, so a copy of a published variant is on sale straight away. Downloadable files and their folders are copied too. Serial codes and issued licence keys are not copied: a copy that delivers serial codes gets a stock of `0`, so it shows as sold out until you add codes. Subscription variants cannot be duplicated. Requires the `listing` API-key ability and the `listing` permission in your team role. Not safe to retry blindly: every call creates another variant. If a response is lost, list the product's variants, with `with_drafts=true` for a draft, before trying again. A variant of a bundle, add-on, credits product, booking product, or course returns `422` on `product`. A subscription variant returns `422` on `variant`. A variant of another product returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/duplicate-a-product-variant) · Effect: **write**

```sh
sellapp product-variants duplicate --product 1 2
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

Inspect the complete schema: `sellapp commands product-variants duplicate --json`.

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

Replace a variant's pricing and payment-method selection. Omitted `compare_at_price` and `bulk_discount` values are cleared. An omitted `trial` is kept. A `trial` is accepted only on `SUBSCRIPTION` pricing and is dropped when the pricing becomes `SINGLE_PAYMENT`. Send `trial: null` to remove it. An empty payment-method array is accepted only when the resulting published variant is genuinely free, including having no paid credit-rate tiers. Bundle products only support one-time pricing. A bundle variant whose resulting pricing type is `SUBSCRIPTION` is refused with a `422` carrying that sentence on `pricing.type`. Requires the `listing` Sanctum ability. Relationship IDs are scoped to the current store. Use the ISO 8601 string `expected_updated_at` to reject stale writes. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

Update only the supplied pricing, subscription cadence, quantity pricing tier, or payment-method fields; omitted fields retain their current values. A `trial` is accepted only on `SUBSCRIPTION` pricing and is dropped when the pricing becomes `SINGLE_PAYMENT`. Send `trial: null` to remove it. An empty payment-method array is accepted only when the resulting published variant is genuinely free, including having no paid credit-rate tiers. Bundle products only support one-time pricing. A bundle variant whose resulting pricing type is `SUBSCRIPTION` is refused with a `422` carrying that sentence on `pricing.type`. Requires the `listing` Sanctum ability. Relationship IDs are scoped to the current store. Use the ISO 8601 string `expected_updated_at` to reject stale writes. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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

## publish

Publish a product variant

Publish one draft option of a booking product so customers can book it. It is checked the way the booking editor checks it: a title, a description of at least 5 characters, finished booking settings, a price, and payment methods enabled for the store that can take its currency. On a published booking product the option can be booked straight away; a draft booking product still needs `POST /v2/products/{product}/publish`, which needs at least one published option when the product has several. Only booking options are published one at a time. A variant of any other kind of product returns `422` on `variant`: its variants are published with the product through `POST /v2/products/{product}/publish`. Requires the `listing` API-key ability and the `listing` permission in your team role. Safe to repeat: once the option is published, another call changes nothing and returns `422` with `This booking option is already published.` on `variant`. Every other refusal is also `422` on `variant`, with one sentence per thing the option still needs. A variant of another product, or a product of another store, returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/publish-a-product-variant) · Effect: **consequential**

```sh
sellapp product-variants publish --product 11 12 --yes
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

Inspect the complete schema: `sellapp commands product-variants publish --json`.

## reorder

Reorder product variants

Set the order buyers see a normal product's published variants in, first to last. Send every published variant ID exactly once; draft variants are not part of the order and are refused. The response lists the product's published variants in their new order, without pagination. Requires the `listing` API-key ability and the `listing` permission in your team role. Safe to repeat: the same list gives the same order. A list that leaves out a published variant, or names a draft or a variant of another product, returns `422` on `variant_ids`; read the variants again and send the full list. An ID sent twice returns `422` on `variant_ids.{i}`. A listing that is not a normal product returns `422` on `product`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/reorder-product-variants) · Effect: **consequential**

```sh
sellapp product-variants reorder 1 --body '{"variant_ids":[2,1]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --variant-ids | Yes | Every published variant ID of the product, first to last. |

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

Inspect the complete schema: `sellapp commands product-variants reorder --json`.

## replace

Update a product variant with PUT

Update a product variant using the PUT compatibility method. Fields remain partial for backward compatibility. Bundle products only support one-time pricing. Adding the bundle deliverable type to a subscription variant is refused with a `422` carrying that sentence on `deliverable.types`. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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
| --payment-methods | No | Complete gateway selection for this variant. Selecting CUSTOM_PAYMENT_METHOD also requires at least one custom_payment_method_ids entry. |
| --custom-payment-method-ids | No | Enabled, usable custom payment method ULIDs belonging to the selected store. Supply at least one when payment_methods includes CUSTOM_PAYMENT_METHOD. Include CUSTOM_PAYMENT_METHOD in payment_methods whenever sending nonempty IDs. On updates, omitting both selections preserves the current methods. Nonempty IDs are rejected when custom payments are not selected. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot for updates. Use the variant's latest updated_at value; stale snapshots receive a 422 validation response. |
| --stock | No | Set an existing variant's stock count while preserving delivery instructions and uploaded files. Send 0 for out of stock or null for unlimited stock. Omit deliverable for a stock-only change. TEXT variants manage stock through their serial inventory and reject this field. |

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

Search a product's variants using JSON body filters, search terms, includes, and sort instructions, 15 per page by default. Draft variants are left out unless you send `with_drafts` or `only_drafts`. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/search-product-variants) · Effect: **read**

```sh
sellapp product-variants search 1 --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| product (positional) | Yes | The product path parameter. |
| --limit | No | Number of items to return per page. |
| --page | No | Page number to return. |
| --with-drafts | No | Include draft variants alongside published variants. |
| --only-drafts | No | Return only draft variants. |
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

## subscription-settings get

Retrieve variant subscription settings

Get how buyers of a published subscription variant may pause (`pause_policy`) and move their billing date (`billing_date_policy`), whether its delivery is repeated on each renewal (`renewal_policy`), and the plan changes offered to them (`change_paths`). A setting that was never saved shows the default the dashboard would save. `state` identifies these values; send it back as `expected_state` when you update them. Requires the `listing` API-key ability and the `listing` permission in your team role. A one-time variant or a draft variant has no such settings and returns `404`, as does a variant of another product. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/variant-subscription-settings) · Effect: **read**

```sh
sellapp product-variants subscription-settings get --product 1 4
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

Inspect the complete schema: `sellapp commands product-variants subscription-settings get --json`.

## subscription-settings update

Update variant subscription settings

Update the pause, billing date, and renewal settings of a published subscription variant. Send only the policies you change: a policy you leave out is not written, and a field you leave out inside a policy keeps its current value. The result is checked as a whole, as the dashboard checks it. Plan changes (`change_paths`) cannot be changed here. Requires the `listing` API-key ability and the `listing` permission in your team role. Send the `state` from your last read as `expected_state` to refuse the update when someone changed the settings in the meantime; that refusal is `422` on `expected_state`. Repeating the same body without `expected_state` gives the same result. A one-time variant or a draft variant returns `422` on `variant_id`. A body without any policy returns `422` on `pause_policy`, an unknown field inside a policy returns `422` on that policy, and a value out of range returns `422` on the field, such as `billing_date_policy.allowed_day_max`. A variant of another product returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/variant-subscription-settings) · Effect: **write**

```sh
sellapp product-variants subscription-settings update --product 1 4 --body '{"pause_policy":{"max_pause_days":30},"renewal_policy":{"redeliver_on_renewal":false},"expected_state":"3f9a1c0e7b2d4a6f8e1c5b9d0a7f3e2c4b6d8a0f1e3c5a7b9d2f4e6a8c0b1d3e"}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product path parameter. |
| variant (positional) | Yes | The variant path parameter. |
| --pause-policy | No | Value for pause policy. |
| --billing-date-policy | No | Value for billing date policy. |
| --renewal-policy | No | Value for renewal policy. |
| --expected-state | No | The `state` from your last read. The update is refused when the settings changed since. |

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

Inspect the complete schema: `sellapp commands product-variants subscription-settings update --json`.

## update

Update a product variant

Updates a product variant's details. Bundle products only support one-time pricing. Adding the bundle deliverable type to a subscription variant is refused with a `422` carrying that sentence on `deliverable.types`. Requires a token with the listing ability. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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
| --payment-methods | No | Complete gateway selection for this variant. Selecting CUSTOM_PAYMENT_METHOD also requires at least one custom_payment_method_ids entry. |
| --custom-payment-method-ids | No | Enabled, usable custom payment method ULIDs belonging to the selected store. Supply at least one when payment_methods includes CUSTOM_PAYMENT_METHOD. Include CUSTOM_PAYMENT_METHOD in payment_methods whenever sending nonempty IDs. On updates, omitting both selections preserves the current methods. Nonempty IDs are rejected when custom payments are not selected. |
| --other-settings | No | Value for other settings. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot for updates. Use the variant's latest updated_at value; stale snapshots receive a 422 validation response. |
| --stock | No | Set an existing variant's stock count while preserving delivery instructions and uploaded files. Send 0 for out of stock or null for unlimited stock. Omit deliverable for a stock-only change. TEXT variants manage stock through their serial inventory and reject this field. |

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

