# TAP Go Coding Standard

This document defines the coding standard for all Go code in this repository
(server and CLI client). It is inspired by the 42 Norm (v4), adapted to Go:
rules that conflict with `gofmt` or with Go idioms are dropped or replaced,
and Go-specific rules are added.

Most rules are enforced automatically by `golangci-lint` (configuration:
[`server/.golangci.yml`](server/.golangci.yml)) and by extra checks in the
`Makefile`. Rules marked **manual** are checked during peer review.

## Running the checks

```sh
make install   # once: downloads Go deps and golangci-lint v2.15.0
make lint      # runs golangci-lint + extra Makefile checks
```

To fix formatting automatically:

```sh
cd server && ~/go/bin/golangci-lint fmt
```

`make build` and `make run-server` run `make lint` first: code that breaks
the standard does not build.

---

## 1. Naming

| Rule | Checked by |
|---|---|
| Identifiers use `MixedCaps` / `mixedCaps`, never underscores (`maxLength`, not `MAX_LENGTH`). | `revive` var-naming, staticcheck ST1003 |
| A leading capital letter means exported; lowercase means package-private. | Go language |
| Initialisms keep a consistent case: `roomID`, `ServeHTTP`, not `roomId`. | staticcheck ST1003 |
| Identifiers contain ASCII characters only. String literals may contain UTF-8. | `asciicheck` |
| Identifiers are in English and explicit. Short names only in very small scopes. | `varnamelen`, **manual** |
| Do not shadow built-in identifiers (`len`, `new`, `copy`, ...). | `predeclared` |
| Error types end with `Error`; sentinel error variables start with `Err`. | `errname`, staticcheck ST1012 |
| Method receivers are 1-2 letters, identical across all methods of a type (`g *Game`). | staticcheck ST1006, ST1016, `revive` receiver-naming |
| File names use lowercase letters, digits and `_` only (`game_loop.go`). | **manual** |
| Package names are a single short lowercase word, matching the directory. | **manual** |
| Getters have no `Get` prefix: `Owner()`, not `GetOwner()`. | **manual** |
| Global variables are allowed. | — |
| `init()` functions are forbidden. | `gochecknoinits` |

## 2. Formatting

| Rule | Checked by |
|---|---|
| Code is formatted with `gofmt` and `goimports` (tabs, spacing, alignment, braces on the same line). | `gofmt`, `goimports` |
| Tabs are counted as 4 columns. | `lll` (`tab-width: 4`) |
| A line is at most **80 columns**, comments included. | `lll` (`line-length: 80`) |
| A function body is at most **25 lines** (comments not counted). | `funlen` (`lines: 25`) |
| No empty line at the start or end of a block. | `whitespace` |
| Nested `if` complexity is at most 3. | `nestif` |
| Control structures are nested at most 3 levels deep. | `revive` max-control-nesting |
| Cyclomatic complexity is at most 10; cognitive complexity at most 15. | `gocyclo`, `gocognit` |

## 3. Functions

| Rule | Checked by |
|---|---|
| A function takes at most **4 parameters**. | `revive` argument-limit |
| A function returns at most **2 values**. | `revive` function-result-limit |
| `error` is always the last return value. | staticcheck ST1008, `revive` error-return |
| At most **5 variables** declared per block. | **manual** |
| No naked `return`. | `nakedret` (`max-func-lines: 0`) |
| No unused parameters. | `unparam`, `revive` unused-parameter |
| Prefer early returns; no `else` after a `return`. | `revive` early-return, indent-error-flow |
| `os.Exit` and `log.Fatal` are only allowed in `main`. | `revive` deep-exit |

## 4. Errors

| Rule | Checked by |
|---|---|
| Every returned error is checked; never discard it with `_`. | `errcheck` (`check-blank: true`) |
| Type assertions check the `ok` value: `v, ok := x.(T)`. | `errcheck` (`check-type-assertions: true`) |
| Error strings start lowercase and have no final punctuation: `"room not found"`. | staticcheck ST1005, `revive` error-strings |
| Compare errors with `errors.Is` / `errors.As`; wrap them with `%w`. | `errorlint` |
| `panic` is forbidden: return an `error` instead. | `forbidigo` |

## 5. Types and constants

| Rule | Checked by |
|---|---|
| One main type per file, in the package it belongs to. | **manual** |
| Enums are `const` blocks with `iota`. A `switch` on an enum handles every value or has a `default`. | `exhaustive` |
| No magic numbers: use named constants (`MaxLineLen = 1024`). | `mnd` |
| Repeated string literals become constants (command names, error codes). | `goconst` |
| No unnecessary type conversions. | `unconvert` |

## 6. Imports

| Rule | Checked by |
|---|---|
| No unused imports. | Go compiler |
| Imports are grouped, standard library first. | `goimports` |
| No dot imports (`import . "fmt"`). | staticcheck ST1001 |
| No import cycles. | Go compiler |

## 7. Forbidden

| Rule | Checked by |
|---|---|
| `goto` | `make lint` (grep) |
| `panic` | `forbidigo` |
| `init()` | `gochecknoinits` |
| `//nolint` without a specific linter and a justification. Required form: `//nolint:linter // reason` | `nolintlint` |

Not forbidden, unlike the 42 Norm: `for` (the only loop in Go), `switch`
(idiomatic for command dispatch), `fmt.Print*`. Go has no ternary operator.

## 8. Comments

| Rule | Checked by |
|---|---|
| Every exported name has a doc comment starting with its name: `// ParseCommand parses ...`. | `revive` exported, staticcheck ST1020-ST1022 |
| Declaration comments end with a period. | `godot` |
| Comments are in English and spelled correctly (US locale). | `misspell`, `dupword` |
| No comments inside function bodies. | **manual** |
| A package doc comment is not required. | — |

## 9. Files

| Rule | Checked by |
|---|---|
| At most **5 function definitions** per `.go` file, methods included. `_test.go` files are exempt. | `make lint` (grep + awk) |

## 10. Tests

Test files (`_test.go`) are exempt from `funlen`, `goconst`, `mnd`,
`varnamelen` and `lll`, so that table-driven tests stay readable.

## 11. Dropped 42 Norm rules

| 42 rule | Reason |
|---|---|
| `s_`, `t_`, `e_`, `g_` prefixes and snake_case identifiers | Go uses `MixedCaps`; case controls visibility. |
| Opening brace on its own line | Invalid in Go (automatic semicolon insertion). |
| Declarations at the top of the function, no declaration + initialization | `x := f()` and `v, err := f()` are idiomatic Go. |
| `return (x)` | Not idiomatic Go. |
| `for`, `switch`, `case` forbidden | `for` is Go's only loop; `switch` is idiomatic. |
| 42 header | Not used in Go projects. |
| Headers / include guards | Go has packages and imports instead. |
