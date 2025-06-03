package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/powwu/go-quip"
)

func main() {
	// Checks
	token := os.Getenv("QUIP_TOKEN")
	if token == "" {
		panic("Set your quip token using the QUIP_TOKEN environment variable.")
	}

	if len(os.Args) != 2 {
		panic("Please specify the file to import from.")
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("Could not open file: %v", err))
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Errorf("Could not read CSV: %v", err))
	}

	if len(records) == 0 {
		panic("No data in provided CSV.")
	}

	// Convert CSV to HTML table
	var content strings.Builder
	content.WriteString("<table border='1'>\n")
	content.WriteString("  <thead>\n    <tr>\n")

	for _, header := range records[0] {
		content.WriteString("      <th>" + template.HTMLEscapeString(header) + "</th>\n")
	}
	content.WriteString("    </tr>\n  </thead>\n")
	content.WriteString("  <tbody>\n")

	for _, row := range records[1:] {
		content.WriteString("    <tr>\n")
		for _, cell := range row {
			content.WriteString("      <td>" + template.HTMLEscapeString(cell) + "</td>\n")
		}
		content.WriteString("    </tr>\n")
	}

	content.WriteString("  </tbody>\n")
	content.WriteString("</table>")

	// Perform API request
	q := quip.NewClient(token)
	newSheetParams := quip.NewDocumentParams{
		Content: content.String(),
		Title:   filepath.Base(fileName),
		Type:    "spreadsheet",
	}

	sheet := q.NewDocument(&newSheetParams)
	if len(sheet.UserIds) == 0 {
		panic("Couldn't verify document creation. Make sure you're using the correct authorization token, and that your custom endpoint (where applicable) does not end in `/`.")
	}

	fmt.Println("Done! The result is in your Private folder.")
}
