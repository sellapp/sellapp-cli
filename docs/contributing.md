# Contributing and local releases

Run `go test ./...` and `go vet ./...` before building. Change the owning generator templates and regenerate source.

Build local distributions with `go run ./tools/release -output dist` in a new or empty directory, then verify with `go run ./tools/release -output dist -verify`. This prepares six platform archives and local npm packages without publishing. Metadata and checksums describe local builds; they are not signed provenance.

Publication candidates must be built and verified with `-production`. This requires the official production browser-login configuration; release artifacts always target `https://sell.app`. Successful build checks do not establish that a real browser login works.
