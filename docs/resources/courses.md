# courses

[All commands](../commands.md)

## get

Retrieve a course

Retrieve one store-scoped course by listing ID or slug with its ordered curriculum. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/retrieve-course) · Effect: **read**

```sh
sellapp courses get test_course
```

| Input | Required | Meaning |
| --- | --- | --- |
| course (positional) | Yes | The course path parameter. |

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

Inspect the complete schema: `sellapp commands courses get --json`.

## lessons create

Create a course lesson

Append a lesson, quiz, or assignment to a section. Rich text is sanitized before storage. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-lessons) · Effect: **write**

```sh
sellapp courses lessons create --course test_course 1 --title Welcome --type text --content 'Welcome to Launch Lab.' --is-published false
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| section (positional) | Yes | The section path parameter. |
| --title | Yes | Value for title. |
| --type | Yes | Value for type. |
| --content | No | Value for content. |
| --is-preview | No | Preview is available only for video lessons. |
| --is-published | No | Value for is published. |
| --assignment | No | Value for assignment. |
| --questions | No | Value for questions. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses lessons create --json`.

## lessons delete

Delete a course lesson

Delete a lesson and its structured questions, assignment, attachments, and video reference, then normalize lesson order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-lessons) · Effect: **consequential**

```sh
sellapp courses lessons delete --course test_course 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| lesson (positional) | Yes | The lesson path parameter. |

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

Inspect the complete schema: `sellapp commands courses lessons delete --json`.

## lessons reorder

Reorder course lessons

Replace the complete lesson order and section placement. Send every lesson ID exactly once with a destination section from the same course. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/reorder-course-lessons) · Effect: **consequential**

```sh
sellapp courses lessons reorder test_course --body '{"resources":[{"id":601,"section_id":501},{"id":602,"section_id":501}]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| course (positional) | Yes | The course path parameter. |
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

Inspect the complete schema: `sellapp commands courses lessons reorder --json`.

## lessons replace

Update a course lesson

Update lesson content, draft state, preview state, or replace the structured quiz or assignment body. Item type is immutable except while choosing video or text for a new lecture placeholder. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-lessons) · Effect: **write**

```sh
sellapp courses lessons replace --course test_course 1 --title Welcome --is-published false
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| lesson (positional) | Yes | The lesson path parameter. |
| --title | No | Value for title. |
| --type | No | Value for type. |
| --content | No | Value for content. |
| --is-preview | No | Preview is available only for video lessons. |
| --is-published | No | Value for is published. |
| --assignment | No | Value for assignment. |
| --questions | No | Value for questions. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses lessons replace --json`.

## lessons update

Update a course lesson

Update lesson content, draft state, preview state, or replace the structured quiz or assignment body. Item type is immutable except while choosing video or text for a new lecture placeholder. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-lessons) · Effect: **write**

```sh
sellapp courses lessons update --course test_course 1 --title Welcome --is-published false
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| lesson (positional) | Yes | The lesson path parameter. |
| --title | No | Value for title. |
| --type | No | Value for type. |
| --content | No | Value for content. |
| --is-preview | No | Preview is available only for video lessons. |
| --is-published | No | Value for is published. |
| --assignment | No | Value for assignment. |
| --questions | No | Value for questions. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses lessons update --json`.

## list

List courses

List course products and their complete ordered curriculum for the authenticated store. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/list-courses) · Effect: **read**

```sh
sellapp courses list
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

Inspect the complete schema: `sellapp commands courses list --json`.

## replace

Update a course

Update course metadata, access configuration, delivery copy, or publication visibility. Public visibility requires a ready promotional video and at least one published lesson. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/update-course) · Effect: **write**

```sh
sellapp courses replace test_course --level beginner --visibility HIDDEN
```

| Input | Required | Meaning |
| --- | --- | --- |
| course (positional) | Yes | The course path parameter. |
| --category | No | Value for category. |
| --level | No | Value for level. |
| --language | No | Value for language. |
| --subtitle | No | Value for subtitle. |
| --author | No | Value for author. |
| --subcategory | No | Value for subcategory. |
| --what-you-learn | No | Value for what you learn. |
| --requirements | No | Value for requirements. |
| --certificate-enabled | No | Value for certificate enabled. |
| --access-type | No | Value for access type. |
| --access-duration-days | No | Required when access_type is limited; cleared otherwise. |
| --enrollment-limit | No | Value for enrollment limit. |
| --delivery-text | No | Value for delivery text. |
| --visibility | No | Value for visibility. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses replace --json`.

## search

Search courses

Search courses by title, slug, or description and compose supported filters and sort instructions. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/search-courses) · Effect: **read**

```sh
sellapp courses search --body '{"filters":[{"field":"id","operator":"=","value":1}],"sort":[{"field":"created_at","direction":"desc"}]}'
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

Inspect the complete schema: `sellapp commands courses search --json`.

## sections create

Create a course section

Append a section to a course curriculum. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-sections) · Effect: **write**

```sh
sellapp courses sections create test_course --title 'Getting started'
```

| Input | Required | Meaning |
| --- | --- | --- |
| course (positional) | Yes | The course path parameter. |
| --title | Yes | Value for title. |
| --description | No | Value for description. |

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

Inspect the complete schema: `sellapp commands courses sections create --json`.

## sections delete

Delete a course section

Delete a section, its lessons, structured questions, attachments, and video references, then normalize remaining section order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-sections) · Effect: **consequential**

```sh
sellapp courses sections delete --course test_course 1 --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| section (positional) | Yes | The section path parameter. |

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

Inspect the complete schema: `sellapp commands courses sections delete --json`.

## sections reorder

Reorder course sections

Replace the complete section order. Send every section ID exactly once in the desired order. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/reorder-course-sections) · Effect: **consequential**

```sh
sellapp courses sections reorder test_course --body '{"resources":[501,502]}' --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| course (positional) | Yes | The course path parameter. |
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

Inspect the complete schema: `sellapp commands courses sections reorder --json`.

## sections replace

Update a course section

Update a section title or sanitized rich-text description. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-sections) · Effect: **write**

```sh
sellapp courses sections replace --course test_course 1 --title 'Getting started'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| section (positional) | Yes | The section path parameter. |
| --title | No | Value for title. |
| --description | No | Value for description. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses sections replace --json`.

## sections update

Update a course section

Update a section title or sanitized rich-text description. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/manage-course-sections) · Effect: **write**

```sh
sellapp courses sections update --course test_course 1 --title 'Getting started'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --course | Yes | The course path parameter. |
| section (positional) | Yes | The section path parameter. |
| --title | No | Value for title. |
| --description | No | Value for description. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses sections update --json`.

## update

Update a course

Update course metadata, access configuration, delivery copy, or publication visibility. Public visibility requires a ready promotional video and at least one published lesson. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/update-course) · Effect: **write**

```sh
sellapp courses update test_course --level beginner --visibility HIDDEN
```

| Input | Required | Meaning |
| --- | --- | --- |
| course (positional) | Yes | The course path parameter. |
| --category | No | Value for category. |
| --level | No | Value for level. |
| --language | No | Value for language. |
| --subtitle | No | Value for subtitle. |
| --author | No | Value for author. |
| --subcategory | No | Value for subcategory. |
| --what-you-learn | No | Value for what you learn. |
| --requirements | No | Value for requirements. |
| --certificate-enabled | No | Value for certificate enabled. |
| --access-type | No | Value for access type. |
| --access-duration-days | No | Required when access_type is limited; cleared otherwise. |
| --enrollment-limit | No | Value for enrollment limit. |
| --delivery-text | No | Value for delivery text. |
| --visibility | No | Value for visibility. |
| --expected-updated-at | No | Value for expected updated at. |

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

Inspect the complete schema: `sellapp commands courses update --json`.

