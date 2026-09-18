# SellApp CLI

Read products, manage orders, and work with your stores from the terminal.
Discover exact command schemas for scripts and agents.

## Install

Install with Node.js 22 or newer:

```sh
npm install -g @sell.app/cli
sellapp --version
```

npm installation needs Node.js 22 or newer. The package launches the included Go
executable; it does not download binaries during installation or first use.
Linux, macOS, and Windows support x64 and arm64.

Standalone archives require neither Node.js nor Go. Extract the matching archive
from [GitHub Releases](https://github.com/sellapp/sellapp-cli/releases/latest)
and put `sellapp` (`sellapp.exe` on Windows) on your `PATH`. On macOS, Linux,
and WSL, download and run the installer:

```sh
curl -fsSL https://sell.app/docs/cli/install.sh -o install.sh
sh install.sh
```

The installer checks the archive checksum and installs to `$XDG_BIN_HOME` or
`~/.local/bin`, without `sudo`. Set `SELLAPP_INSTALL_DIR` to choose a directory,
`SELLAPP_INSTALL_VERSION` to pin a version, or `SELLAPP_INSTALL_NO_MODIFY_PATH=1`
to leave shell profiles alone. Follow its PATH message; a new shell may be needed.

## Make your first request

```sh
sellapp login
```

Sign in in the browser and approve account-wide access to current and future accessible stores.
Consent never adds store permissions: the CLI uses your current membership and
role. One available store is selected automatically; choose one in the terminal
when several are available. To inspect or change your target:

```sh
sellapp stores list
sellapp stores use YOUR-STORE-SLUG
sellapp products list --limit 1 --output json
```

Replace `YOUR-STORE-SLUG` with a slug from the list. Consent uses the admin grant;
current membership and role determine access. `stores use` selects the target.
Zero stores is successful login: run `sellapp setup` to choose or create one.
Use `sellapp stores create --name NAME --slug SLUG --use` for explicit creation and selection.
The product request reads real data without changing it. A JSON array with one
product is a result; `[]` with exit code `0` means the request worked and there
are no products. For a store `403`, check the selected store, granted scopes,
and current membership and role. Keep the request ID for support.

The CLI remembers your login and refreshes access automatically. Setup needs no
keyring service, client secret, or manual token handling.

### Browserless login and errors

```sh
sellapp login --no-browser
```

Open the printed authorization URL, keep the command running, and approve within
five minutes. The browser must reach the CLI's temporary local callback. WSL can
use a Windows browser when that loopback connection reaches WSL.

For `invalid_client` or `invalid_scope`, check that you have the official release
and contact support if the error persists. Repeating the same rejected login
will not help. After `invalid_grant` or an uncertain refresh, authorize again.
Declined consent needs a new attempt only if you want to approve.

## Work with an agent

Install the bundled SellApp skill in one explicit scope:

```sh
sellapp skills add --project
# Or make it available across projects:
sellapp skills add --global
```

The default installs under `.agents/skills/sellapp`. Choose another client with
`sellapp skills add --client claude-code --project` (`.claude/skills`) or
`sellapp skills add --client cursor --project` (`.cursor/skills`). Use
`--global` instead of `--project` for your user. Start a new agent session
after installation and select the `sellapp` skill.

The skill teaches command discovery, store selection, bounded reads, dry runs,
and safe retries. Installation preserves local edits; use `--force` only when
you intend to replace them. See [agent guidance](docs/agent-skill.md).

Register the hosted MCP bridge with a supported client:

```sh
sellapp mcp add --client codex --global
sellapp mcp doctor --client codex --global
```

Claude Code and Cursor are also supported. Use `--help` for their scope choices.
The bridge uses your saved CLI login without another browser consent. Hosted
permissions apply to reads and writes; consequential actions require confirmation.
Registration does not save tokens in client configuration. Bare `sellapp mcp`
starts the stdio bridge.

## Preview the next change

```sh
sellapp products get 120
sellapp commands products create --json
sellapp products create --title "Design kit" --description "Templates for your next project." --visibility HIDDEN --dry-run
```

Replace `120` with a returned product ID. The create command is an offline
preview. Removing `--dry-run` creates a real hidden product in the selected store
when your role permits it. Your existing connection is used. Complete ordinary
catalog edits run directly; consequential actions require confirmation, or
explicit `--yes` in automation. Inspect the command schema before changing data.

## Saved access and disconnect

```sh
sellapp auth status
sellapp auth logout
```

Status reports the profile, selected store, and storage backend without secrets.
New profiles use encrypted credential files on every OS. Explicit vault use is
available through `SELLAPP_CLI_SECRET_STORE=keyring`; see
[alternative authentication and storage](docs/usage.md#alternative-authentication-and-storage).

Logout revokes the connection before removing local credentials. You can also
inspect and disconnect through **CLI access** in your dashboard, including after
leaving every store. Membership and role changes apply to subsequent requests.

## API keys for automation

Create a server-side key with the `listing` ability in your store's Developer
settings. Supply it through your automation secret storage as `SELLAPP_API_KEY`,
and set `SELLAPP_STORE` to the store slug. Then run
`sellapp products list --limit 1 --output json`.

An environment API key takes precedence over saved OAuth credentials. Remove
`SELLAPP_API_KEY` from the environment to use your saved CLI login. Never commit
a key or paste it into a shared command or log.

## Updates

For an npm installation:

```sh
npm install -g @sell.app/cli@latest
```

For a standalone installation, rerun the installer or replace the executable
from a verified release archive. The CLI does not update in the background.

## Build from source

With Go 1.25 or newer, run `go mod download` and
`go build -o sellapp ./cmd/sellapp` in the source checkout. Use `-o sellapp.exe`
on Windows. Add that executable to your PATH or use `./sellapp`.

A plain source build supports offline command discovery, dry runs, and
[API-key authentication](#api-keys-for-automation). It does not include the
official browser-login registration. Use an official release for `sellapp login`
and the MCP bridge.

## Reference

- [Command reference](docs/commands.md): positional IDs, flags, examples, and scopes.
- [Usage](docs/usage.md): credentials, output, pagination, errors, and workflows.
- [Executable examples](examples/README.md) and [SellApp skill](skills/sellapp/SKILL.md).
- `sellapp search TERM`, `sellapp docs products create`, and `sellapp products create --help`.

Output is readable in a terminal and JSON when redirected; `--output` overrides
that choice. Data goes to stdout and diagnostics to stderr. One page is the
default; bound `--all` with `--max-items` and `--max-pages`.
`sellapp mcp` and `sellapp --mcp` bridge to hosted MCP using the saved OAuth
login. Hosted permissions and consequential confirmation apply. Compact mode is
the default; select `--mode catalog`, `--mode full`, or `--mode retrieval` when needed.
Search returns up to 10 summaries; use `--limit` (1–25) and `--offset` to page
through matches. JSON includes `results`, `total`, and `next_offset` (null on the
last page). Each result points to a leaf `--schema` with inputs, authentication,
effects, pagination, and retry rules. `sellapp --llms` lists commands and schema
lookups without repeating those details; `--llms-full` retains the full metadata.
Discovery works without login. Use `sellapp workflow list` /
`sellapp workflow show ID` for guided tasks.

[API documentation](https://sell.app/docs/api) ·
[Support](https://github.com/sellapp/sellapp-cli/issues) ·
[License](LICENSE.txt) · [Contributor checks](docs/contributing.md)

## Payment and ticket shortcuts

Create charges with `sellapp charges create 25.00 USD`, return a customer payment link with `sellapp wallet topups create CUSTOMER_ID 25.00 --payment-method STRIPE`, or use wallet credit/debit and ticket close/reopen/archive. See [complete inputs, effects, and retry rules](docs/payments.md).
