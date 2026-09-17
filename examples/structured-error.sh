#!/usr/bin/env sh
set -eu
./sellapp --api-key dummy --store example products get || test "$?" -eq 2
