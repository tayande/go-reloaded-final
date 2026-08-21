# cli text correction tool

A text-transformation engine written in Go, now wrapped in a simple web interface so you can run every supported conversion from the browser instead of the command line.

## Overview

Go Reloaded reads raw text and applies a set of inline transformation rules, number base conversions, case changes, punctuation cleanup, and grammar fixes, producing a correctly formatted output string. The original version worked as a CLI file-processing pipeline; this version adds an HTTP server and HTML form so the same engine can be used interactively.

## Features / Supported Conversions

| Rule | Example Input | Example Output | Description |
|---|---|---|---|
| Hex → Decimal | `1E (hex)` | `30` | Converts a hexadecimal number to its decimal equivalent |
| Binary → Decimal | `1010 (bin)` | `10` | Converts a binary number to its decimal equivalent |
| Uppercase | `hello (up)` | `HELLO` | Converts the preceding word to uppercase |
| Lowercase | `WORLD (low)` | `world` | Converts the preceding word to lowercase |
| Capitalize | `go (cap)` | `Go` | Capitalizes the first letter of the preceding word |
| Multi-word modifiers | `ready set go (up, 2)` | `READY SET GO` | Applies up/low/cap to the preceding N words instead of just one |
| Punctuation spacing | `Hello , world !` | `Hello, world!` | Removes incorrect spacing before punctuation marks (`.` `,` `!` `?` `:` `;`) |
| Quote tightening | `' well done '` | `'well done'` | Removes extra spaces inside single quotes |
| Article correction | `a apple` | `an apple` | Fixes "a" → "an" when followed by a word starting with a vowel sound |

## Web Interface

The web UI puts a browser front end on top of the same transformation engine used by the CLI:

- A single-page form where you paste or type input text
- Submit button sends the text to the server for processing
- The transformed result is rendered back on the page
- No page reload required beyond the initial form submission (standard `net/http` + `html/template` handling)

### Under the hood
- **Backend:** Go `net/http` server handling form submissions (`r.ParseForm()` / `r.FormValue()`)
- **Templating:** `html/template` for rendering the input form and result page
- **Core logic:** unchanged transformation engine from the original CLI version (string parsing, regex-based rule matching, sequential rule application)
- **Routing:** separate routes for displaying the form (`GET`) and processing submitted text (`POST`)


## Running the Project

```bash
git clone <repo-url>
cd cli-text-correction-tool
go run .
```

Then open `http://localhost:8080` (or the configured port) in your browser, paste in your text, and submit to see the transformed output.

## Rule Application Order

Rules are applied in a single left-to-right pass over the text, so modifiers always act on the word(s) immediately preceding them, and multiple modifiers can be chained in the same input string.

## Notes

- Input parsing uses regex-based matching (`regexp` / `ReplaceAllString`) to detect modifier tags like `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`.
- Malformed or unsupported modifier tags are left untouched in the output rather than causing a crash.
