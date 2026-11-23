# support-automation-settings

[All commands](../commands.md)

## get

Retrieve support automation settings

Read the store's support automation switch, its daily cap on automatic refunds and how many rules are switched on. While `automation_enabled` is `false`, no rule acts on any ticket. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-settings/retrieve-support-automation-settings) · Effect: **read**

```sh
sellapp support-automation-settings get
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

Inspect the complete schema: `sellapp commands support-automation-settings get --json`.

## replace

Update support automation settings

Turn support automation on or off and set the daily refund cap; only the fields you send change, and `PUT` and `PATCH` behave the same. Turning automation on lets every switched-on rule act on new matching tickets at once: reply to the customer, credit wallet refunds and close tickets. `daily_refund_cap_cents` of `0` or `null` removes the cap. While a switched-on rule refunds, automation cannot run without a cap: turning it on without one returns `422` on `automation_enabled`, and removing the cap returns `422` on `daily_refund_cap_cents`. To refuse the change when someone changed the settings since you read them, send `expected_automation_enabled` and `expected_daily_refund_cap_cents` together, as you read them; a mismatch returns `422` on both and changes nothing. A read can be sent back unchanged. `enabled` and `expected_enabled` are refused with `422`. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Repeating the same change is safe, except with the old expected pair, which no longer matches after the first change: read the settings again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-settings/update-support-automation-settings) · Effect: **consequential**

```sh
sellapp support-automation-settings replace --automation-enabled true --daily-refund-cap-cents 5000 --expected-automation-enabled false --expected-daily-refund-cap-cents null --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --automation-enabled | No | `true` lets switched-on rules act on new tickets; `false` stops all of them. |
| --daily-refund-cap-cents | No | Most that rules may credit as automatic refunds in one day, in integer USD cents (`5000` means $50.00). `0` or `null` removes the cap. |
| --expected-automation-enabled | No | `automation_enabled` as you read it. Send it together with `expected_daily_refund_cap_cents`. |
| --expected-daily-refund-cap-cents | No | `daily_refund_cap_cents` as you read it (`null` or `0` for no cap). Send it together with `expected_automation_enabled`. |
| --currency | No | Read-only. Accepted only as `USD`. |
| --active-rules-count | No | Read-only and ignored. |
| --active-refund-rules-count | No | Read-only and ignored. |

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

Inspect the complete schema: `sellapp commands support-automation-settings replace --json`.

## update

Update support automation settings

Turn support automation on or off and set the daily refund cap; only the fields you send change, and `PUT` and `PATCH` behave the same. Turning automation on lets every switched-on rule act on new matching tickets at once: reply to the customer, credit wallet refunds and close tickets. `daily_refund_cap_cents` of `0` or `null` removes the cap. While a switched-on rule refunds, automation cannot run without a cap: turning it on without one returns `422` on `automation_enabled`, and removing the cap returns `422` on `daily_refund_cap_cents`. To refuse the change when someone changed the settings since you read them, send `expected_automation_enabled` and `expected_daily_refund_cap_cents` together, as you read them; a mismatch returns `422` on both and changes nothing. A read can be sent back unchanged. `enabled` and `expected_enabled` are refused with `422`. Requires the `support.automation` API-key ability, and the key's user needs the Support Automation role in the store (owners and administrators have it). Repeating the same change is safe, except with the old expected pair, which no longer matches after the first change: read the settings again. OAuth callers use the admin grant and select a store with X-STORE. Current membership and role permissions apply to every request.

[API reference](https://sell.app/docs/api/support-automation-settings/update-support-automation-settings) · Effect: **consequential**

```sh
sellapp support-automation-settings update --automation-enabled true --daily-refund-cap-cents 5000 --expected-automation-enabled false --expected-daily-refund-cap-cents null --yes
```

| Input | Required | Meaning |
| --- | --- | --- |
| --automation-enabled | No | `true` lets switched-on rules act on new tickets; `false` stops all of them. |
| --daily-refund-cap-cents | No | Most that rules may credit as automatic refunds in one day, in integer USD cents (`5000` means $50.00). `0` or `null` removes the cap. |
| --expected-automation-enabled | No | `automation_enabled` as you read it. Send it together with `expected_daily_refund_cap_cents`. |
| --expected-daily-refund-cap-cents | No | `daily_refund_cap_cents` as you read it (`null` or `0` for no cap). Send it together with `expected_automation_enabled`. |
| --currency | No | Read-only. Accepted only as `USD`. |
| --active-rules-count | No | Read-only and ignored. |
| --active-refund-rules-count | No | Read-only and ignored. |

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

Inspect the complete schema: `sellapp commands support-automation-settings update --json`.

