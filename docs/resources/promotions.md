# promotions

[All commands](../commands.md)

## create

Create a promotion

Create a promotion and its initial ordered phase set. `applies_to` chooses whether it discounts one-time purchases or the first `subscription_payments` payments of a subscription plan, and a subscription promotion can be limited to picked products and plans with `store_wide`, `products`, and `product_variants`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/create-promotion) · Effect: **consequential**

```sh
sellapp promotions create --body '{"name":"Ship Week","status":"active","starts_at":"2026-08-01T00:00:00Z","ends_at":"2026-08-08T00:00:00Z","priority":1,"is_stackable":false,"max_redemptions":500,"phases":[{"discount_type":"percentage","discount_value":"20","ends_at":"2026-08-04T00:00:00Z","max_redemptions":200,"minimum_amount":"10"},{"discount_type":"fixed","discount_value":"5","ends_at":null,"max_redemptions":null,"minimum_amount":"25"}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | Value for name. |
| --status | Yes | Value for status. |
| --priority | Yes | Value for priority. |
| --is-stackable | Yes | Value for is stackable. |
| --phases | Yes | The complete ordered phase set. Array order becomes phase position. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | A date and time after starts_at and after the current time, or null for no end date. |
| --max-redemptions | No | Value for max redemptions. |
| --applies-to | No | What the promotion discounts: `one_time` for one-time purchases, or `subscriptions` for the first payments of a subscription plan. A new promotion defaults to `one_time`, and an update that leaves this out keeps the current kind. The promotion only discounts order lines of its kind. |
| --subscription-payments | No | How many payments a `subscriptions` promotion discounts, counting the signup payment as the first, so `3` discounts the signup and the next two renewals. A `subscriptions` promotion needs one, and an update that leaves it out keeps the stored number. Missing on a `subscriptions` promotion, or outside 1 to 36, it returns `422` on `subscription_payments` with `The number of discounted payments must be between 1 and 36.` A `one_time` promotion always stores null here. |
| --store-wide | No | Send `false` with `products` to limit a `subscriptions` promotion to picked products and plans. Required whenever the request sends `products` or `product_variants`. A new promotion defaults to `true`, an update that leaves it out keeps the current scope, and switching a promotion to `one_time` makes it store-wide again. A `one_time` promotion with `false` returns `422` on `store_wide` with `A promotion for one-time purchases applies to the whole store.` |
| --products | No | Product IDs a narrowed subscription promotion reaches. Each one must sell a subscription, or the request returns `422` on `products` with `Pick only products that sell a subscription.` A narrowed promotion left with no products returns `422` with `Pick at least one subscription product for this promotion.` An update replaces the picks with this list, or keeps them when it is left out. |
| --product-variants | No | Subscription plan IDs to limit the picked products to. Each one must be a subscription plan of a product in `products`, or the request returns `422` on `product_variants` with `Pick only subscription plans of the picked products.` A picked product with none of its plans listed takes every plan, including plans added later. An update that leaves this out keeps the listed plans of products that stay picked. |

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

Inspect the complete schema: `sellapp commands promotions create --json`.

## delete

Delete a promotion

Delete a promotion. Unreferenced phases are purged; phases referenced by order history remain soft-deleted. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/delete-promotion) · Effect: **consequential**

```sh
sellapp promotions delete 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |

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

Inspect the complete schema: `sellapp commands promotions delete --json`.

## get

Retrieve a promotion

Retrieve a promotion and its ordered phases. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/retrieve-promotion) · Effect: **read**

```sh
sellapp promotions get 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |

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

Inspect the complete schema: `sellapp commands promotions get --json`.

## list

List promotions

List promotions for the authenticated store, including their ordered phases. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/list-promotions) · Effect: **read**

```sh
sellapp promotions list
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

Inspect the complete schema: `sellapp commands promotions list --json`.

## phases list

List promotion phases

List a promotion's active phases in evaluation order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/replace-promotion-phases) · Effect: **read**

```sh
sellapp promotions phases list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |

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

Inspect the complete schema: `sellapp commands promotions phases list --json`.

## phases replace

Replace promotion phases

Replace the complete phase set. Array order becomes phase position; at least one phase is required. Historical referenced phases remain soft-deleted. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/replace-promotion-phases) · Effect: **consequential**

```sh
sellapp promotions phases replace 1 --body '{"phases":[{"discount_type":"percentage","discount_value":"20","ends_at":"2026-08-04T00:00:00Z","max_redemptions":200,"minimum_amount":"10"},{"discount_type":"fixed","discount_value":"5","ends_at":null,"max_redemptions":null,"minimum_amount":"25"}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |
| --phases | Yes | Value for phases. |

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

Inspect the complete schema: `sellapp commands promotions phases replace --json`.

## replace

Update a promotion

Update promotion fields. Include `phases` to replace the complete ordered phase set; omit it to preserve current phases. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/update-promotion) · Effect: **consequential**

```sh
sellapp promotions replace 1 --name 'One More Sprint' --is-stackable true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |
| --name | No | Value for name. |
| --status | No | Value for status. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | A date and time after starts_at and after the current time, or null for no end date. |
| --priority | No | Value for priority. |
| --is-stackable | No | Value for is stackable. |
| --max-redemptions | No | Value for max redemptions. |
| --applies-to | No | What the promotion discounts: `one_time` for one-time purchases, or `subscriptions` for the first payments of a subscription plan. A new promotion defaults to `one_time`, and an update that leaves this out keeps the current kind. The promotion only discounts order lines of its kind. |
| --subscription-payments | No | How many payments a `subscriptions` promotion discounts, counting the signup payment as the first, so `3` discounts the signup and the next two renewals. A `subscriptions` promotion needs one, and an update that leaves it out keeps the stored number. Missing on a `subscriptions` promotion, or outside 1 to 36, it returns `422` on `subscription_payments` with `The number of discounted payments must be between 1 and 36.` A `one_time` promotion always stores null here. |
| --store-wide | No | Send `false` with `products` to limit a `subscriptions` promotion to picked products and plans. Required whenever the request sends `products` or `product_variants`. A new promotion defaults to `true`, an update that leaves it out keeps the current scope, and switching a promotion to `one_time` makes it store-wide again. A `one_time` promotion with `false` returns `422` on `store_wide` with `A promotion for one-time purchases applies to the whole store.` |
| --products | No | Product IDs a narrowed subscription promotion reaches. Each one must sell a subscription, or the request returns `422` on `products` with `Pick only products that sell a subscription.` A narrowed promotion left with no products returns `422` with `Pick at least one subscription product for this promotion.` An update replaces the picks with this list, or keeps them when it is left out. |
| --product-variants | No | Subscription plan IDs to limit the picked products to. Each one must be a subscription plan of a product in `products`, or the request returns `422` on `product_variants` with `Pick only subscription plans of the picked products.` A picked product with none of its plans listed takes every plan, including plans added later. An update that leaves this out keeps the listed plans of products that stay picked. |
| --phases | No | The complete ordered phase set. Array order becomes phase position. |

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

Inspect the complete schema: `sellapp commands promotions replace --json`.

## restore

Restore a promotion

Restore a soft-deleted promotion with a complete replacement phase set. Historical phases referenced by order history remain soft-deleted and are not reactivated. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/restore-promotion) · Effect: **consequential**

```sh
sellapp promotions restore 1 --body '{"name":"Ship Week","status":"active","starts_at":"2026-08-01T00:00:00Z","ends_at":"2026-08-08T00:00:00Z","priority":1,"is_stackable":false,"max_redemptions":500,"phases":[{"discount_type":"percentage","discount_value":"20","ends_at":"2026-08-04T00:00:00Z","max_redemptions":200,"minimum_amount":"10"},{"discount_type":"fixed","discount_value":"5","ends_at":null,"max_redemptions":null,"minimum_amount":"25"}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |
| --name | Yes | Value for name. |
| --status | Yes | Value for status. |
| --priority | Yes | Value for priority. |
| --is-stackable | Yes | Value for is stackable. |
| --phases | Yes | The complete ordered phase set. Array order becomes phase position. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | A date and time after starts_at and after the current time, or null for no end date. |
| --max-redemptions | No | Value for max redemptions. |
| --applies-to | No | What the promotion discounts: `one_time` for one-time purchases, or `subscriptions` for the first payments of a subscription plan. A new promotion defaults to `one_time`, and an update that leaves this out keeps the current kind. The promotion only discounts order lines of its kind. |
| --subscription-payments | No | How many payments a `subscriptions` promotion discounts, counting the signup payment as the first, so `3` discounts the signup and the next two renewals. A `subscriptions` promotion needs one, and an update that leaves it out keeps the stored number. Missing on a `subscriptions` promotion, or outside 1 to 36, it returns `422` on `subscription_payments` with `The number of discounted payments must be between 1 and 36.` A `one_time` promotion always stores null here. |
| --store-wide | No | Send `false` with `products` to limit a `subscriptions` promotion to picked products and plans. Required whenever the request sends `products` or `product_variants`. A new promotion defaults to `true`, an update that leaves it out keeps the current scope, and switching a promotion to `one_time` makes it store-wide again. A `one_time` promotion with `false` returns `422` on `store_wide` with `A promotion for one-time purchases applies to the whole store.` |
| --products | No | Product IDs a narrowed subscription promotion reaches. Each one must sell a subscription, or the request returns `422` on `products` with `Pick only products that sell a subscription.` A narrowed promotion left with no products returns `422` with `Pick at least one subscription product for this promotion.` An update replaces the picks with this list, or keeps them when it is left out. |
| --product-variants | No | Subscription plan IDs to limit the picked products to. Each one must be a subscription plan of a product in `products`, or the request returns `422` on `product_variants` with `Pick only subscription plans of the picked products.` A picked product with none of its plans listed takes every plan, including plans added later. An update that leaves this out keeps the listed plans of products that stay picked. |

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

Inspect the complete schema: `sellapp commands promotions restore --json`.

## search

Search promotions

Search promotions by name and compose supported filters and sort instructions. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/search-promotions) · Effect: **read**

```sh
sellapp promotions search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
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

Inspect the complete schema: `sellapp commands promotions search --json`.

## update

Update a promotion

Update promotion fields. Include `phases` to replace the complete ordered phase set; omit it to preserve current phases. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/promotions/update-promotion) · Effect: **consequential**

```sh
sellapp promotions update 1 --name 'One More Sprint' --is-stackable true --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| promotion (positional) | Yes | The promotion path parameter. |
| --name | No | Value for name. |
| --status | No | Value for status. |
| --starts-at | No | Value for starts at. |
| --ends-at | No | A date and time after starts_at and after the current time, or null for no end date. |
| --priority | No | Value for priority. |
| --is-stackable | No | Value for is stackable. |
| --max-redemptions | No | Value for max redemptions. |
| --applies-to | No | What the promotion discounts: `one_time` for one-time purchases, or `subscriptions` for the first payments of a subscription plan. A new promotion defaults to `one_time`, and an update that leaves this out keeps the current kind. The promotion only discounts order lines of its kind. |
| --subscription-payments | No | How many payments a `subscriptions` promotion discounts, counting the signup payment as the first, so `3` discounts the signup and the next two renewals. A `subscriptions` promotion needs one, and an update that leaves it out keeps the stored number. Missing on a `subscriptions` promotion, or outside 1 to 36, it returns `422` on `subscription_payments` with `The number of discounted payments must be between 1 and 36.` A `one_time` promotion always stores null here. |
| --store-wide | No | Send `false` with `products` to limit a `subscriptions` promotion to picked products and plans. Required whenever the request sends `products` or `product_variants`. A new promotion defaults to `true`, an update that leaves it out keeps the current scope, and switching a promotion to `one_time` makes it store-wide again. A `one_time` promotion with `false` returns `422` on `store_wide` with `A promotion for one-time purchases applies to the whole store.` |
| --products | No | Product IDs a narrowed subscription promotion reaches. Each one must sell a subscription, or the request returns `422` on `products` with `Pick only products that sell a subscription.` A narrowed promotion left with no products returns `422` with `Pick at least one subscription product for this promotion.` An update replaces the picks with this list, or keeps them when it is left out. |
| --product-variants | No | Subscription plan IDs to limit the picked products to. Each one must be a subscription plan of a product in `products`, or the request returns `422` on `product_variants` with `Pick only subscription plans of the picked products.` A picked product with none of its plans listed takes every plan, including plans added later. An update that leaves this out keeps the listed plans of products that stay picked. |
| --phases | No | The complete ordered phase set. Array order becomes phase position. |

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

Inspect the complete schema: `sellapp commands promotions update --json`.

