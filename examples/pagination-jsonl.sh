#!/usr/bin/env sh
set -eu
./sellapp --api-key dummy --store example --base-url http://127.0.0.1:9 --dry-run --max-retries 0 --output jsonl products list --all --max-items 20 --max-pages 2
