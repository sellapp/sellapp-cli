# SellApp CLI usage

Configuration precedence is flags, invocation environment, selected profile, then safe defaults. Profiles contain nonsecret store, client, origin, and credential-reference metadata. New profiles save access locally and require no keyring setup. Run `sellapp auth status` to see the selected profile and storage backend.

A TTY may prompt for required schema fields. Non-TTY use never prompts: missing input is a JSON error with exit code 2. Consequential actions name the operation and store; automation needs `--yes`. `--dry-run` never reaches the network and redacts Authorization, X-STORE, and Idempotency-Key. Keys are generated only for operations declaring idempotency.

Outputs: `json`, `jsonl`, `yaml`, `csv`, `markdown`, `table`, and `raw`. The small expression surface is `--fields a,b`, `--filter field=value`, and `--template '{{field}}'`; no code executes. Bound pagination with `--max-items` and `--max-pages`.

## Browser login and saved access

Run `sellapp login`, sign in in the browser, and approve access across current and future accessible stores. Consent permits access to real store data within your current membership and role permissions. Owners have no separate CLI blocking control. Inspect or disconnect your personal connection in **CLI access**. One accessible store is selected automatically; multiple stores prompt in a terminal. Use `sellapp stores list` and `sellapp stores use SLUG` to change selection. Switching stores cannot expand access. A user-owned API key with account:read can also list stores; selected-store restrictions and current membership still limit the result.

Use `sellapp login --no-browser` when the browser cannot open. Open the printed authorization URL and keep the command running. It uses fresh PKCE S256 and state with a five-minute numeric loopback callback; the browser must reach that listener.

The CLI saves access and refresh tokens together and refreshes them automatically. One process may refresh a profile at a time. A lost or invalid-grant refresh response blocks further use of the old token; authorize again. An `invalid_client` or `invalid_scope` response requires checking the CLI build or server configuration; repeating the same login will not fix it. A store-specific denial preserves credentials for other stores. A crashed process can leave a lock: do not delete it while another process might still be running. `sellapp auth logout` revokes the OAuth connection before removing local credentials.

## Alternative authentication and storage

Supply an automation key through `SELLAPP_API_KEY` and its target through `SELLAPP_STORE`. The API key takes precedence over a saved OAuth profile. Manual-key profiles use `sellapp login --method api-key --store SLUG`; interactive input hides the key.

New profiles use encrypted files on macOS, Linux/WSL, and Windows. Select a backend deliberately with `SELLAPP_CLI_SECRET_STORE=file` or `SELLAPP_CLI_SECRET_STORE=keyring`. An unavailable selected keyring is an error, with no silent fallback. Linux/WSL keyring use requires `secret-tool` and a running, unlocked Secret Service. File storage does not require either. Existing profiles retain their selected backend.

V1 uses one public admin grant. It permits only reviewed operations and never bypasses current membership or role checks. A connection with zero stores can run `sellapp setup` to choose or explicitly create one. New store creation is limited to one per user every 60 seconds across credentials. On 429, honor Retry-After and reuse the original key and body; a successful replay returns the original store.

Credentials and signed download URLs are redacted from formatted results and MCP output. Never enable shell tracing when passing credentials. A customer-session token uses `--customer-session` or `SELLAPP_CUSTOMER_SESSION` and is isolated from merchant credentials.

## Event, export, and delivery workflows

`workflow events-poll --checkpoint ./events.cursor --handler /absolute/path/process-event --max-polls 5` passes each event as JSON on the handler's standard input. The handler must save all required work durably before exiting with status zero. Only then, after every event in that page succeeds, does the CLI save the next cursor. Handler failure or a crash can repeat previously processed events; make processing safe for duplicates. The checkpoint is bound to the origin, profile, store, and filters. A `410 cursor_expired` requires rebuilding your state from current resources; the CLI does not silently discard the old cursor.

`workflow export --type sales --format csv --parameters '{"from":"2026-08-01","to":"2026-08-31"}' --idempotency-key launch-sales-202608 --file ./sales.csv --yes` creates a real export job and polls at most five times, honoring the server's `poll_after_seconds` minimum. Reuse that key only for the same intended export and body. If still pending, resume with the returned `--export ID`. Downloads refuse existing files, cap output at 1 GiB, and send no API credentials to object storage. Only one export per store and two per account may be queued or running. On 429, honor Retry-After and poll an existing job. Omitted dates cover today and the previous 365 days. Reports exceeding 10,000 source rows, 5 MiB of source data, or 10 MiB output fail without a partial file; use a smaller date range. Inspect authorization failures, failed jobs, and expired exports before trying again.

`workflow webhook-delivery --delivery ID` inspects a stored delivery. Add `--replay --idempotency-key replay-delivery-001 --yes` to queue a real external request to its saved destination. Reuse that key only for the same delivery. A queued replay does not mean its destination received it; inspect the returned delivery ID afterward.

## Composed workflows

`store-readiness` reads store settings, payment methods, webhook channels, and custom domains. `orders-poll` and `support-poll` require an explicit finite `--max-polls`; `support-triage` combines the ticket list, ticket detail, and messages. `product-draft` creates a private draft and may then upload a deliverable or inventory serial file when explicit product and variant IDs are supplied. `webhook-test` rereads the channel before sending a named test event. Every write keeps the normal confirmation, dry-run, and audit controls.

Catalog export and diff are local. Catalog apply accepts a JSON plan with a `changes` array. Each change names `read_operation`, `write_operation`, optional `path`/`query`, the `expected` reread value, and the write `body`. Apply rereads every affected resource and refuses the whole plan on any mismatch before sending writes. A reread alone does not prevent concurrent writes. Supply the operation's documented concurrency fields in each write body; this local plan is not an atomic transaction.

The redacted JSONL audit lives in the OS config directory, mode 0600, rotates at 1 MiB, and retains five files. Disable it with `--no-audit` or `SELLAPP_AUDIT_DISABLED=1`. MCP uses the hosted bridge and its write/confirmation policy, with compact discovery by default. Run `sellapp workflow list` or `sellapp workflow show ID` for a focused guide. Install the bundled skill with an explicit project or global scope; it makes no AI calls.

Use the generated stores, catalog, and webhook delivery commands for their documented server workflows. Do not assume safe retries for writes without a documented guarantee.
