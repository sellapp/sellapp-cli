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

## notify

Notify variant buyers

Email your customers that this variant has new content: every completed order that delivered its downloadable files gets one email. The emails are queued, so `202` means they were queued, not delivered. It sends only for a published variant with at least one file, at least one such order, and no upload in progress. The buyers of one variant can be emailed this way at most once an hour, whichever API key or team member sends it, and each team member can send at most 3 times a minute per variant. Requires the `listing` API-key ability and the `listing` permission in your team role. Send an `Idempotency-Key` header; requests without one return `400`. A retry with the same key and an empty body returns the first response with `Idempotent-Replayed: true` and sends nothing new. Use a new key for every send you intend, including a new attempt after a `422`. The same key while the first request is still running returns `409`. Every refusal is `422` on `variant`: a draft variant, a variant without files, a variant nobody has bought, an upload still in progress, the hourly limit (the message says how many minutes are left), the per-minute limit, or another send for the same variant at the same moment. Check `GET` on the same path first to see the numbers. A variant of another product returns `404`. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/notify-variant-buyers) · Effect: **write**

```sh
sellapp variant-deliverables notify --product 1 2
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

Inspect the complete schema: `sellapp commands variant-deliverables notify --json`.

## preview

Preview variant buyer notification

See who would be emailed if you told this variant's buyers about new downloadable files, without sending anything. Use it before `POST` on the same path to check the numbers. Works for any variant, drafts included. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/product-variants/notify-variant-buyers) · Effect: **read**

```sh
sellapp variant-deliverables preview --product 1 2
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

Inspect the complete schema: `sellapp commands variant-deliverables preview --json`.

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

