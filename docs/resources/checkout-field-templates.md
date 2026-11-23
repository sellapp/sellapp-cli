# checkout-field-templates

[All commands](../commands.md)

## create

Create a checkout field template

Save a named set of checkout fields that you can reuse on products. Creating a template does not change any product. Each successful call creates another template, but a second template with the same name returns `422` on `name`, so a retry after a lost response cannot create a duplicate. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/create-a-checkout-field-template) · Effect: **consequential**

```sh
sellapp checkout-field-templates create --body '{"name":"Discord access","fields":[{"type":"text","label":"Discord username","key":"discord_username","required":true,"placeholder":"name#0001","description":"We add this account to the community after purchase."},{"type":"radio","label":"Experience level","key":"experience_level","required":false,"options":["Beginner","Intermediate","Advanced"],"variant":"cards","variant_options":{"vertical":true,"indicator":true}}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --name | Yes | The template name, 2 to 100 characters and unique in the store. Surrounding spaces are removed. |
| --body-fields | Yes | The checkout fields, 1 to 25, in the order customers see them. Properties a checkout field does not have are dropped. An invalid field returns `422` on its position, such as `fields.0.options`. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates create --json`.

## delete

Delete a checkout field template

Delete a template permanently and return it as it was. Products that use its fields keep them. A template that is already gone returns `404`, so a retry after a lost response cannot delete anything else. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/delete-a-checkout-field-template) · Effect: **consequential**

```sh
sellapp checkout-field-templates delete 14 --expected-updated-at 2026-09-04T09:00:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| checkoutFieldTemplate (positional) | Yes | The template ID. |
| --expected-updated-at | No | The `updated_at` you last read. If the template changed since, the request returns `422` on `expected_updated_at` and changes nothing. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates delete --json`.

## duplicate

Duplicate a checkout field template

Copy a template's fields into a new template with the name you give. The original does not change. A name already used in the store returns `422` on `name`, so retrying after a lost response cannot create a second copy. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/duplicate-a-checkout-field-template) · Effect: **consequential**

```sh
sellapp checkout-field-templates duplicate 14 --name 'Discord access for courses' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| checkoutFieldTemplate (positional) | Yes | The template ID. |
| --name | Yes | The template name, 2 to 100 characters and unique in the store. Surrounding spaces are removed. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates duplicate --json`.

## get

Retrieve a checkout field template

Retrieve one checkout field template with its fields. A template from another store returns `404`. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/retrieve-a-checkout-field-template) · Effect: **read**

```sh
sellapp checkout-field-templates get 14
```

| Input | Required | Meaning |
| --- | --- | --- |
| checkoutFieldTemplate (positional) | Yes | The template ID. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates get --json`.

## list

List checkout field templates

List your store's checkout field templates in one response, sorted by name. The list is not paginated. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/list-checkout-field-templates) · Effect: **read**

```sh
sellapp checkout-field-templates list
```

| Input | Required | Meaning |
| --- | --- | --- |
| --search | No | Return only templates whose name contains this text, ignoring case. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates list --json`.

## replace

Update a checkout field template

Rename a template, replace its fields, or both. Fields you send replace the whole list; leave `fields` out to keep them. PUT behaves exactly like PATCH. Products that already use these fields keep their own copy and do not change. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/update-a-checkout-field-template) · Effect: **consequential**

```sh
sellapp checkout-field-templates replace 14 --name 'Discord onboarding' --expected-updated-at 2026-09-04T09:00:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| checkoutFieldTemplate (positional) | Yes | The template ID. |
| --name | No | The template name, 2 to 100 characters and unique in the store. Surrounding spaces are removed. |
| --body-fields | No | The checkout fields, 1 to 25, in the order customers see them. Properties a checkout field does not have are dropped. An invalid field returns `422` on its position, such as `fields.0.options`. |
| --expected-updated-at | No | The `updated_at` you last read. If the template changed since, the request returns `422` on `expected_updated_at` and changes nothing. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates replace --json`.

## update

Update a checkout field template

Rename a template, replace its fields, or both. Fields you send replace the whole list; leave `fields` out to keep them. PUT behaves exactly like PATCH. Products that already use these fields keep their own copy and do not change. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/checkout-field-templates/update-a-checkout-field-template) · Effect: **consequential**

```sh
sellapp checkout-field-templates update 14 --name 'Discord onboarding' --expected-updated-at 2026-09-04T09:00:00.000000Z --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| checkoutFieldTemplate (positional) | Yes | The template ID. |
| --name | No | The template name, 2 to 100 characters and unique in the store. Surrounding spaces are removed. |
| --body-fields | No | The checkout fields, 1 to 25, in the order customers see them. Properties a checkout field does not have are dropped. An invalid field returns `422` on its position, such as `fields.0.options`. |
| --expected-updated-at | No | The `updated_at` you last read. If the template changed since, the request returns `422` on `expected_updated_at` and changes nothing. |

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

Inspect the complete schema: `sellapp commands checkout-field-templates update --json`.

