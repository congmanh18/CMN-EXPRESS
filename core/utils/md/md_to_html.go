package md

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
)

func ConvertMdToHtml(filePath string) (string, error) {
	// Kiểm tra file Markdown tồn tại
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("markdown file not found: %s", filePath)
	}

	// Đọc nội dung file Markdown
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading Markdown file: %s", err)
	}

	// Chuyển Markdown thành HTML
	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		return "", fmt.Errorf("error converting Markdown to HTML: %s", err)
	}

	// Bọc nội dung HTML với CSS
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Markdown Documentation</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
				padding: 0;
				line-height: 1.6;
			}
			table {
				border-collapse: collapse;
				width: 100%;
				margin-bottom: 20px;
			}
			table, th, td {
				border: 1px solid #ddd;
			}
			th, td {
				text-align: left;
				padding: 8px;
			}
			th {
				background-color: #f4f4f4;
			}
			tr:nth-child(even) {
				background-color: #f9f9f9;
			}
			h1, h2 {
				color: #333;
			}
		</style>
	</head>
	<body>
		` + buf.String() + `
	</body>
	</html>`
	return html, nil
}
