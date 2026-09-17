# variant-deliverables

[All commands](../commands.md)

## get

Retrieve variant deliverable configuration

Retrieve variant deliverable configuration. Requires the listing Sanctum ability. Returns delivery types and safe configuration for a variant in the authenticated store. Dynamic endpoint URLs and inventory data are not returned. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **read**

```sh
sellapp variant-deliverables get --product test_product 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| variant (positional) | Yes | The product variant identifier. |

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

Inspect the complete schema: `sellapp commands variant-deliverables get --json`.

## replace

Replace variant deliverable configuration

Replace variant deliverable configuration. Requires the listing Sanctum ability. Replaces this resource's delivery types atomically. Pass expected_updated_at from the product variant resource to reject stale updates with 422. TEXT, BOOKING, CREDITS, and BUNDLE must remain unchanged. Removing DOWNLOADABLE deletes managed files and folders after commit; removing a type clears its private configuration. Dynamic endpoints are write-only. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants) · Effect: **write**

```sh
sellapp variant-deliverables replace --product test_product 1 --body '{"types":["MANUAL"],"data":{"comment":"Delivery is arranged by Launch Lab."}}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --product | Yes | The product identifier or slug. The product must belong to the authenticated store. |
| variant (positional) | Yes | The product variant identifier. |
| --types | Yes | The complete type list. TEXT, BOOKING, CREDITS, and BUNDLE are owned by their dedicated resources and must remain unchanged here. |
| --data | Yes | Value for data. |
| --expected-updated-at | No | Optional optimistic-concurrency snapshot from the deliverable resource's updated_at field. The request returns 422 if the variant changed after this timestamp was loaded. |

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

Inspect the complete schema: `sellapp commands variant-deliverables replace --json`.

