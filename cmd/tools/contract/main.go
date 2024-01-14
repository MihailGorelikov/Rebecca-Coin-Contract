package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

// SolidityContract represents the structure of the JSON data.
type SolidityContract struct {
	Format                 string       `json:"_format"`
	ContractName           string       `json:"contractName"`
	SourceName             string       `json:"sourceName"`
	ABI                    []ABIElement `json:"abi"`
	Bytecode               string       `json:"bytecode"`
	DeployedBytecode       string       `json:"deployedBytecode"`
	LinkReferences         any          `json:"linkReferences"`
	DeployedLinkReferences any          `json:"deployedLinkReferences"`
}

// ABIElement represents an element in the ABI array.
type ABIElement struct {
	Inputs          []InputOutput `json:"inputs"`
	Name            string        `json:"name,omitempty"`
	StateMutability string        `json:"stateMutability,omitempty"`
	Type            string        `json:"type"`
	Outputs         []InputOutput `json:"outputs,omitempty"`
	Anonymous       bool          `json:"anonymous,omitempty"`
}

// InputOutput represents an input or output in the ABIElement.
type InputOutput struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

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

	var contract SolidityContract
	err = json.Unmarshal(contractSource, &contract)
	if err != nil {
		panicf("failed to unmarshal contract ABI file: %v", err)
	}

	contractABI, err := json.MarshalIndent(contract.ABI, "", "\t")
	if err != nil {
		panicf("failed to marshal contract ABI: %v", err)
	}

	templateData := TemplateData{
		Filename:       "rebecca_coin_contract.go",
		PackageName:    "rebecca_coin_contract",
		ContractSource: string(contractABI),
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
