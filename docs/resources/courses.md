# courses

[All commands](../commands.md)

## create

Create a course

Create a course as a draft that customers cannot buy yet, whatever `visibility` you send: `visibility`, `HIDDEN` by default, applies once the course is published. It starts with no sections and one draft option named `Default`, priced at 0 USD: set the price with the product variant endpoints, add sections and lessons, then publish it with `POST /v2/products/{product}/publish`, where `{product}` is the course `id`. Each successful call creates another course, so do not retry blindly after a lost response: search courses for the title first. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/create-a-course) · Effect: **write**

```sh
sellapp courses create --body '{"title":"Brand identity basics","description":"<p>Design a logo, color palette, and type system for your next project.</p>","category":"design","level":"beginner","language":"English","what_you_learn":["Sketch a logo concept.","Choose a color palette."],"certificate_enabled":true,"delivery_text":"Open your course from the link in your receipt."}'
```

| Input | Required | Meaning |
| --- | --- | --- |
| --title | Yes | The course name customers see. |
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
| --slug | No | The last part of the course's storefront URL: lowercase letters and numbers in words joined by single hyphens, such as `design-foundations`. It must be unused by every other product in the store, or the request returns `422` on `slug`. When you leave it out, SellApp makes one from the title. |
| --description | No | The course description as HTML, at least 5 characters long and at most 2000 characters of visible text. Leave it out, or send null, to start without one; a course needs a description before it is published. |
| --visibility | No | Value for visibility. |

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

Inspect the complete schema: `sellapp commands courses create --json`.

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

## get-student-summary

Retrieve a course student summary

Summarize who bought your courses in a date range: totals for students, purchases, units, and orders, plus the top students and the top courses. A purchase counts when it is paid or completed and was made within the range; refunded and pending purchases do not count. A student is a customer, or the order email when a purchase has no customer record. The response contains customer email addresses, so handle it as personal data. Requires the `invoice` API-key ability and the invoice permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/courses/retrieve-course-student-summary) · Effect: **read**

```sh
sellapp courses get-student-summary --start-date 2026-08-01 --end-date 2026-08-31
```

| Input | Required | Meaning |
| --- | --- | --- |
| --start-date | Yes | The first day to include, as `YYYY-MM-DD` in the store's timezone. |
| --end-date | Yes | The last day to include, as `YYYY-MM-DD` in the store's timezone. It must not be before `start_date`, and the range can cover at most 366 days; otherwise the request returns `422` on `end_date`. |
| --course-id | No | Count one course only: the course `id` from `/v2/courses`. Another product, a deleted course, or another store's course returns `422` on `course_id`. |
| --limit | No | How many rows `students` and `courses` each return: 10 by default, at most 20. A larger value returns `422` on `limit`. The `summary` totals always cover everything. |

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

Inspect the complete schema: `sellapp commands courses get-student-summary --json`.

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

Update a course's title, URL slug, description, details, access rules, delivery message, or visibility. Only the fields you send change, and PUT behaves exactly like PATCH. Sending the course's current `visibility` changes nothing; a different `visibility` also publishes a draft course. `PUBLIC` first checks that the course has a ready promotional video and at least one published lesson, and otherwise returns `422` on `promo_video` or `curriculum` without changing anything. Every successful update moves `updated_at` forward, even one that changes nothing, so retrying with the same `expected_updated_at` after a success returns `422`: retrieve the course again first. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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
| --expected-updated-at | No | The `updated_at` you last read. If the course changed since, the request returns `422` on `expected_updated_at` and changes nothing. It is checked even when the request changes nothing else. |
| --title | No | The course name customers see. |
| --slug | No | The last part of the course's storefront URL: lowercase letters and numbers in words joined by single hyphens, such as `design-foundations`. It must be unused by every other product in the store, or the request returns `422` on `slug`. |
| --description | No | The course description as HTML, at least 5 characters long and at most 2000 characters of visible text. `null` keeps the current description. To clear the description of a draft course, send an empty element such as `<p></p>`; a published course must keep one, so clearing it returns `422` on `description`. |

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

Update a course's title, URL slug, description, details, access rules, delivery message, or visibility. Only the fields you send change, and PUT behaves exactly like PATCH. Sending the course's current `visibility` changes nothing; a different `visibility` also publishes a draft course. `PUBLIC` first checks that the course has a ready promotional video and at least one published lesson, and otherwise returns `422` on `promo_video` or `curriculum` without changing anything. Every successful update moves `updated_at` forward, even one that changes nothing, so retrying with the same `expected_updated_at` after a success returns `422`: retrieve the course again first. Requires the `listing` API-key ability and the `listing` permission in your team role. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

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
| --expected-updated-at | No | The `updated_at` you last read. If the course changed since, the request returns `422` on `expected_updated_at` and changes nothing. It is checked even when the request changes nothing else. |
| --title | No | The course name customers see. |
| --slug | No | The last part of the course's storefront URL: lowercase letters and numbers in words joined by single hyphens, such as `design-foundations`. It must be unused by every other product in the store, or the request returns `422` on `slug`. |
| --description | No | The course description as HTML, at least 5 characters long and at most 2000 characters of visible text. `null` keeps the current description. To clear the description of a draft course, send an empty element such as `<p></p>`; a published course must keep one, so clearing it returns `422` on `description`. |

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

