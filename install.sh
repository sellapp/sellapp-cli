#!/bin/sh
# Install the official SellApp CLI from an immutable GitHub release.
set -eu
umask 077
repository='sellapp/sellapp-cli'
die() { printf '%s\n' "sellapp installer: $*" >&2; exit 1; }
say() { printf '%s\n' "$*" >&2; }
for tool in curl tar awk sed grep mktemp chmod mv mkdir uname dirname rm; do
  command -v "$tool" >/dev/null 2>&1 || die "Required program missing: $tool. Install it and rerun this installer."
done
case "$(uname -s)" in Linux) os=linux ;; Darwin) os=darwin ;; *) die 'Supported systems are macOS, Linux, and WSL. Use the npm package or a Windows release archive on Windows.' ;; esac
case "$(uname -m)" in x86_64|amd64) arch=amd64 ;; aarch64|arm64) arch=arm64 ;; *) die "Unsupported architecture: $(uname -m)." ;; esac
if command -v sha256sum >/dev/null 2>&1; then
  digest() { sha256sum < "$1" | awk '{print $1}'; }
elif command -v shasum >/dev/null 2>&1; then
  digest() { shasum -a 256 < "$1" | awk '{print $1}'; }
else die 'SHA-256 verification requires sha256sum or shasum.'; fi
download() { curl --fail --silent --show-error --location --proto '=https' --proto-redir '=https' --tlsv1.2 --connect-timeout 15 --max-time 300 --retry 2 --output "$2" "$1"; }
version=${SELLAPP_INSTALL_VERSION:-latest}
if [ "$version" = latest ]; then
  resolved=$(curl --fail --silent --show-error --location --proto '=https' --proto-redir '=https' --tlsv1.2 --connect-timeout 15 --max-time 60 --output /dev/null --write-out '%{url_effective}' "https://github.com/$repository/releases/latest") || die 'Could not resolve the latest release. Set SELLAPP_INSTALL_VERSION to a release version to retry a pinned installation.'
  prefix="https://github.com/$repository/releases/tag/v"
  case "$resolved" in "$prefix"*) version=${resolved#"$prefix"} ;; *) die 'The latest release did not resolve to an immutable version tag.' ;; esac
fi
version=${version#v}
printf '%s\n' "$version" | LC_ALL=C grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+(-[A-Za-z0-9]+([.-][A-Za-z0-9]+)*)?$' || die 'SELLAPP_INSTALL_VERSION must be latest or a version such as 0.1.0.'
destination=${SELLAPP_INSTALL_DIR:-${XDG_BIN_HOME:-"$HOME/.local/bin"}}
case "$destination" in /*) ;; *) die 'The installation directory must be an absolute path.' ;; esac
case "$destination" in *:*) die 'The installation directory must not contain a colon because PATH uses colons as separators.' ;; esac
case "$destination" in *'
'*|*"$(printf '\r')"*) die 'The installation directory must not contain line breaks.' ;; esac
# Refuse symlinked path components and parent traversal before creating anything.
safe_directory() {
  checked=''
  previous_ifs=$IFS; IFS=/
  set -f
  for component in $1; do
    [ -n "$component" ] || continue
    case "$component" in .|..) IFS=$previous_ifs; set +f; return 1 ;; esac
    checked="$checked/$component"
    if [ -L "$checked" ] || { [ -e "$checked" ] && [ ! -d "$checked" ]; }; then IFS=$previous_ifs; set +f; return 1; fi
  done
  IFS=$previous_ifs; set +f
}
safe_directory "$destination" || die 'The installation directory contains an unsafe path or symbolic link.'
mkdir -p "$destination" || die "Could not create $destination. Choose a writable SELLAPP_INSTALL_DIR; sudo is not needed."
[ ! -L "$destination/sellapp" ] || die 'Refusing to replace a symbolic link named sellapp.'
[ ! -e "$destination/sellapp" ] || [ -f "$destination/sellapp" ] || die 'The existing sellapp path is not a regular file.'
temporary=$(mktemp -d "$destination/.sellapp-install.XXXXXXXX") || die 'Could not create an installation staging directory.'
profile_temporary=''
cleanup() { [ -z "$profile_temporary" ] || rm -f "$profile_temporary"; rm -rf "$temporary"; }
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM
trap 'exit 129' HUP
asset="sellapp_${version}_${os}_${arch}.tar.gz"
base="https://github.com/$repository/releases/download/v$version"
say "Downloading SellApp $version for $os/$arch..."
download "$base/$asset" "$temporary/archive.tar.gz" || die 'The archive download failed. Your existing installation was preserved.'
download "$base/SHA256SUMS" "$temporary/SHA256SUMS" || die 'The checksum download failed. Your existing installation was preserved.'
expected=$(awk -v name="$asset" '$2 == name { count++; value=$1 } END { if (count != 1 || length(value) != 64 || value ~ /[^0-9a-f]/) exit 1; print value }' "$temporary/SHA256SUMS") || die 'The checksum manifest does not contain exactly one valid archive checksum.'
[ "$(digest "$temporary/archive.tar.gz")" = "$expected" ] || die 'The archive checksum did not match. Your existing installation was preserved.'
tar -tzf "$temporary/archive.tar.gz" > "$temporary/names" || die 'The archive is malformed.'
LC_ALL=C awk '
  /^\// || /\\/ || /(^|\/)\.\.?($|\/)/ || /[[:cntrl:]]/ || /\/$/ { exit 1 }
  { if (seen[$0]++) exit 1; if ($0 == "sellapp") binary++ }
  END { if (binary != 1) exit 1 }
' "$temporary/names" || die 'The archive contains unsafe paths, duplicate paths, or no unique sellapp executable.'
tar -tvzf "$temporary/archive.tar.gz" > "$temporary/types" || die 'Could not inspect archive entries.'
LC_ALL=C awk 'substr($0,1,1) != "-" { exit 1 }' "$temporary/types" || die 'The archive contains links or non-regular files.'
# Extract only executable bytes to a private staging file; archive paths are never materialized.
tar -xOzf "$temporary/archive.tar.gz" sellapp > "$temporary/sellapp" || die 'Could not read the executable from the archive.'
[ -s "$temporary/sellapp" ] || die 'The archive executable is empty.'
chmod 755 "$temporary/sellapp"
actual=$("$temporary/sellapp" --version) || die 'The downloaded executable could not start. Your existing installation was preserved.'
[ "$actual" = "sellapp version $version" ] || die "The executable version does not match release $version."
# Rename within the destination filesystem is atomic, including when an earlier CLI is running.
mv -f "$temporary/sellapp" "$destination/sellapp" || die 'Could not replace the executable. Your existing installation was preserved.'
say "Installed SellApp $version at $destination/sellapp."
quote_sh() { printf "'%s'" "$(printf '%s' "$1" | sed "s/'/'\\\\''/g")"; }
quote_fish() { printf "'%s'" "$(printf '%s' "$1" | sed "s/\\\\/\\\\\\\\/g; s/'/\\\\'/g")"; }
case ":$PATH:" in *":$destination:"*) say 'Run sellapp login to connect, approve stores, and select your store.'; exit 0 ;; esac
instruction="export PATH=$(quote_sh "$destination"):\"\$PATH\""
profile=''
case "${SHELL:-}" in
  */bash) if [ "$os" = darwin ]; then profile="$HOME/.bash_profile"; else profile="$HOME/.bashrc"; fi ;;
  */zsh) profile="${ZDOTDIR:-$HOME}/.zshrc" ;;
  */fish) profile="${XDG_CONFIG_HOME:-$HOME/.config}/fish/config.fish"; instruction="fish_add_path -- $(quote_fish "$destination")" ;;
esac
if [ "${SELLAPP_INSTALL_NO_MODIFY_PATH:-0}" = 1 ] || [ -z "$profile" ]; then
  say 'Add the installation directory to PATH, then run sellapp login:'
  say "  $instruction"
  exit 0
fi
if [ -L "$profile" ] || { [ -e "$profile" ] && [ ! -f "$profile" ]; } || ! safe_directory "$(dirname "$profile")"; then
  say "Shell profile $profile is unsafe to edit. Add this line yourself: $instruction"
  exit 0
fi
# An exact existing block needs no rewrite, preserving profile bytes and mode.
if [ -f "$profile" ] && SELLAPP_PATH_LINE="$instruction" awk '
  BEGIN { expected=ENVIRON["SELLAPP_PATH_LINE"] }
  /^# >>> SellApp CLI >>>$/ { starts++; if (getline <= 0 || $0 != expected) bad=1; if (getline <= 0 || $0 != "# <<< SellApp CLI <<<") bad=1 }
  END { exit !(starts == 1 && !bad) }
' "$profile"; then
  say "PATH is already configured in $profile. Open a new shell, then run sellapp login."
  exit 0
fi
mkdir -p "$(dirname "$profile")" || { say "Could not create the shell profile directory. Add this line yourself: $instruction"; exit 0; }
profile_temporary=$(mktemp "$(dirname "$profile")/.sellapp-path.XXXXXXXX") || { say "Could not update the shell profile. Add this line yourself: $instruction"; exit 0; }
if [ -f "$profile" ]; then
  awk '/^# >>> SellApp CLI >>>$/ { if (inside || done) exit 1; inside=1; next } /^# <<< SellApp CLI <<<$/{ if (!inside) exit 1; inside=0; done=1; next } !inside { print } END { if (inside) exit 1 }' "$profile" > "$profile_temporary" || { say "The SellApp PATH block in $profile needs manual review. Add: $instruction"; exit 0; }
fi
printf '\n# >>> SellApp CLI >>>\n%s\n# <<< SellApp CLI <<<\n' "$instruction" >> "$profile_temporary"
mv -f "$profile_temporary" "$profile" || { say "Could not save $profile. Add this line yourself: $instruction"; exit 0; }
profile_temporary=''
say "Updated $profile. Open a new shell, then run sellapp login."
