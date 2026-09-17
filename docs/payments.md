# Charges, wallet top-ups, and ticket shortcuts

After `sellapp login`, select your store with `sellapp stores use SLUG` and check it with `sellapp products list --limit 1`. The commands below create real records or change balances and ticket state. Add `--dry-run` to inspect a request without sending it. Add `--yes` only when you intend to confirm a write in a script.

## Create a charge

```bash
sellapp charges create 25.00 USD \
  --email maya.chen@example.com \
  --return-url https://example.com/payment-return \
  --payment-methods '["STRIPE","PAYPAL"]'
```

The positional amount uses major units: `25.00 USD` sends `total: 2500`. JPY accepts whole amounts; supported currencies such as HUF and TWD use two ISO minor-unit places. Extra decimal places are rejected without rounding. The existing `--total 2500 --currency USD` and JSON-body forms remain available. Do not combine the amount shorthand with `--total`, `--currency`, `--body`, or `--set`.

Use `--use-all-payment-methods` to offer all usable methods, or `--payment-method STRIPE` to preselect one. For a custom method, pass `--payment-method CUSTOM_PAYMENT_METHOD --custom-payment-method-id METHOD_ID`. Use `--custom-payment-method-ids '["METHOD_ID"]'` with a custom-method choice to restrict the allowed methods. IDs must come from `sellapp payment-methods custom list`; availability depends on your store and currency.

Creation returns the charge and its `data.url`. It does not prove payment. A repeated charge creation can create a second charge: inspect the original result or list charges after an uncertain failure before trying again.

```bash
sellapp charges get CHARGE_ID
sellapp charges open CHARGE_ID
sellapp charges wait CHARGE_ID --timeout 5m
```

Use the returned charge ID in these commands. `open` retrieves and opens its returned HTTPS checkout URL. `wait` polls every two seconds, for at most 300 reads and the chosen timeout. Use `--interval` and `--max-polls` to set smaller bounds. Only `COMPLETED` produces `completed: true`. `REVIEW` keeps waiting for independent merchant review; this command never approves payment. Failed, voided, reversed, refunded, or disputed states, unknown states, and timeout return a nonzero exit code with the last charge when available. A failed read stops polling.

## Give a customer a wallet top-up payment link

```bash
sellapp wallet topups create CUSTOMER_ID 25.00 \
  --payment-method STRIPE \
  --idempotency-key topup-customer-001
```

Replace `CUSTOMER_ID` with the existing customer's ID. The command returns `checkout_url`, the link to share with that customer, and `charge_id`, which you can inspect with `charges get` or `charges wait`. Both the HTTP API and CLI JSON wrap these fields in `data`: read `data.checkout_url` and `data.charge_id`.

Wallet top-ups use USD only. Wallets must be enabled and unfrozen, the amount must fit your store's top-up limits, and the selected payment method must be enabled for wallet top-ups. You need both wallet and charge permissions. Custom methods also require `--custom-payment-method-id METHOD_ID`.

Creating the link leaves the wallet balance unchanged. Successful payment credits the deposit and any bonus saved at creation once. Custom-method modifiers do not change the deposit. Submitted payment proof and browser returns do not prove payment; verify funds independently before approving a custom payment. Return and cancel links lead to the store's customer portal wallet.

Retain the same idempotency key, customer, store, amount, and payment method to recover a response within 24 hours. Changed input returns `409`. After an uncertain provider outcome, inspect charges before starting a new top-up. For integer cents or nested JSON, use the original form: `wallet topups create CUSTOMER_ID --amount-cents 2500 --payment-method STRIPE --idempotency-key KEY`.

## Adjust a wallet directly

```bash
sellapp wallet credit CUSTOMER_ID 25.00 --note "Goodwill credit" --idempotency-key goodwill-001
sellapp wallet debit CUSTOMER_ID 5.00 --note "Correct an earlier credit" --idempotency-key correction-001
```

These commands change spendable USD funds immediately without collecting payment. Both require a note and an explicit, stable idempotency key. Supply a positive amount; the command applies the direction. Debits cannot exceed available funds after holds. Reuse the same key and identical input for a retry; use a new key for a new adjustment. These keys belong to the saved wallet transaction and do not expire like the 24-hour response receipts. `wallet adjust` retains the full integer-cents request interface.

## Update a ticket

```bash
sellapp tickets close TICKET_ID --idempotency-key ticket-close-001
sellapp tickets reopen TICKET_ID
sellapp tickets archive TICKET_ID
```

Closing an open ticket notifies the customer. Close and reopen change only status; archive changes only the archive flag. Use `tickets update TICKET_ID --archived=false` to unarchive. For response replay after a lost connection, retain the optional idempotency key and inputs. Current ticket permissions and write confirmation still apply.
