#!/usr/bin/env sh
set -eu
# After sellapp login, register sellapp mcp with your agent client.
# This local example inspects the bridge interface without starting a connection.
./sellapp mcp --help
./sellapp --llms >/dev/null
./sellapp workflow show connect-store
