# highlights

[All commands](../commands.md)

## create

Create a highlight

Create and publish a highlight with one to 30 image or video files. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **write**

```sh
sellapp highlights create --title test_title --hidden true --upload 'files=example-file'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | Value for title. |
| --hidden | Yes | Value for hidden. |
| --upload files=PATH | Yes | Value for files. |

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

Inspect the complete schema: `sellapp commands highlights create --json`.

## delete

Delete a highlight

Delete a highlight in your store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **consequential**

```sh
sellapp highlights delete 42 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |

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

Inspect the complete schema: `sellapp commands highlights delete --json`.

## get

Retrieve a highlight

Retrieve a highlight in your store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **read**

```sh
sellapp highlights get 42
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |

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

Inspect the complete schema: `sellapp commands highlights get --json`.

## list

List highlights

Retrieve a paginated list of highlights for your store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **read**

```sh
sellapp highlights list
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

Inspect the complete schema: `sellapp commands highlights list --json`.

## media add

Add highlight media

Upload an image or video and append it to a highlight, up to 30 media items in total. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **write**

```sh
sellapp highlights media add 1 --upload 'file=example-file'
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |
| --upload file=PATH | Yes | Value for file. |
| --cta-title | No | Value for cta title. |
| --product-id | No | Value for product id. |

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

Inspect the complete schema: `sellapp commands highlights media add --json`.

## media delete

Delete highlight media

Delete highlight media for a highlight. Deleting the last media item unpublishes the highlight. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **consequential**

```sh
sellapp highlights media delete --highlight 1 84 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --highlight | Yes | The highlight path parameter. |
| media (positional) | Yes | The media path parameter. |

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

Inspect the complete schema: `sellapp commands highlights media delete --json`.

## media list

List highlight media

List media attached to a highlight in display order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **read**

```sh
sellapp highlights media list 1
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |

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

Inspect the complete schema: `sellapp commands highlights media list --json`.

## media reorder

Reorder highlight media

Replace the complete media order for a highlight with the supplied media IDs. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **consequential**

```sh
sellapp highlights media reorder 1 --body '{"resources":[42,41]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |
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

Inspect the complete schema: `sellapp commands highlights media reorder --json`.

## media replace

Replace highlight media

Replace highlight media for a highlight. Deleting the last media item unpublishes the highlight. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **write**

```sh
sellapp highlights media replace --highlight 1 84 --upload 'file=example-file'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --highlight | Yes | The highlight path parameter. |
| media (positional) | Yes | The media path parameter. |
| --upload file=PATH | Yes | Value for file. |
| --cta-title | No | Value for cta title. |
| --product-id | No | Value for product id. |

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

Inspect the complete schema: `sellapp commands highlights media replace --json`.

## media replace-metadata

Update highlight media

Update highlight media for a highlight. Deleting the last media item unpublishes the highlight. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **write**

```sh
sellapp highlights media replace-metadata --highlight 1 84 --cta-title 'View product' --product-id 123
```

| Input | Required | Meaning |
| --- | --- | --- |
| --highlight | Yes | The highlight path parameter. |
| media (positional) | Yes | The media path parameter. |
| --cta-title | No | Value for cta title. |
| --product-id | No | Value for product id. |

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

Inspect the complete schema: `sellapp commands highlights media replace-metadata --json`.

## media update

Update highlight media

Update highlight media for a highlight. Deleting the last media item unpublishes the highlight. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlight-media) · Effect: **write**

```sh
sellapp highlights media update --highlight 1 84 --cta-title 'View product' --product-id 123
```

| Input | Required | Meaning |
| --- | --- | --- |
| --highlight | Yes | The highlight path parameter. |
| media (positional) | Yes | The media path parameter. |
| --cta-title | No | Value for cta title. |
| --product-id | No | Value for product id. |

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

Inspect the complete schema: `sellapp commands highlights media update --json`.

## reorder

Reorder highlights

Replace the complete published highlight order with the supplied highlight IDs. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **consequential**

```sh
sellapp highlights reorder --body '{"resources":[42,41]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
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

Inspect the complete schema: `sellapp commands highlights reorder --json`.

## replace

Update a highlight

Update a highlight in your store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **write**

```sh
sellapp highlights replace 42 --title 'Stealth-mode launch' --hidden false
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |
| --title | No | Value for title. |
| --hidden | No | Value for hidden. |
| --published | No | Set to true to publish. Publishing is one-way through this field. |

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

Inspect the complete schema: `sellapp commands highlights replace --json`.

## search

Search highlights

Search highlights using JSON body filters, search terms, includes, and sort instructions. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **read**

```sh
sellapp highlights search --body '{"filters":[{"field":"id","operator":"=","value":42}],"sort":[{"field":"created_at","direction":"desc"}]}'
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

Inspect the complete schema: `sellapp commands highlights search --json`.

## update

Update a highlight

Update a highlight in your store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/highlights/manage-highlights) · Effect: **write**

```sh
sellapp highlights update 42 --title 'Stealth-mode launch' --hidden false
```

| Input | Required | Meaning |
| --- | --- | --- |
| highlight (positional) | Yes | The highlight path parameter. |
| --title | No | Value for title. |
| --hidden | No | Value for hidden. |
| --published | No | Set to true to publish. Publishing is one-way through this field. |

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

Inspect the complete schema: `sellapp commands highlights update --json`.

