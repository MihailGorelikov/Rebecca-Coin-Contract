package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

// TemplateData holds the data to be inserted into the template.
type TemplateData struct {
	Filename       string
	PackageName    string
	ContractSource string
}

//go:embed templates/*
var templates embed.FS

func main() {
	contractSource, err := os.ReadFile("./artifacts/contracts/RebeccaCoin.sol/RebeccaCoin.json")
	if err != nil {
		panicf("failed to open contract ABI file: %v", err)
	}

	templateData := TemplateData{
		Filename:       "rebecca_coin_contract.go",
		PackageName:    "rebecca_coin_contract",
		ContractSource: string(contractSource),
	}

	parsedTemplate, err := template.ParseFS(templates, "templates/rebecca_coin_contract.gotmpl")
	if err != nil {
		panicf("failed to parse template: %v", err)
	}

	file, err := os.Create(templateData.Filename)
	if err != nil {
		panicf("failed to create file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panicf("failed to close file: %v", err)
		}
	}()

	err = parsedTemplate.Execute(file, templateData)
	if err != nil {
		panicf("failed to execute template: %v", err)
	}
}

func panicf(format string, args ...any) string {
	panic(fmt.Sprintf(format, args...))
}
