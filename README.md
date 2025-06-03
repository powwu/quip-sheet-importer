# quip-sheet-importer
Import a spreadsheet into Quip. Currently only supports CSV format.

If any of this poses you trouble, feel free to shoot me an email at `support@powwu.sh`.

## Setup
To use a custom endpoint, set the `QUIP_ENDPOINT` environment variable.
- Go to https://quip.com/api/personal-token and generate a Personal Access Token (PAT).
- Set the `QUIP_TOKEN` environment variable to your PAT. 
- Run as follows: `go run main.go /path/to/file.csv`

