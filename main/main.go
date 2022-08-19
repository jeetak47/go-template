package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/alexflint/go-arg"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var args struct {
		Input        string `arg:"positional" help:"imput json file"`
		Output       string `arg:"positional" help:"output json file"`
		Verbose      bool   `arg:"-v,--verbose" help:"verbosity level"`
		Optimize     int    `arg:"-O" help:"optimization level"`
		TemplateFile string `arg:"-t,--template" help:"template file"`
	}
	arg.MustParse(&args)
	jsonFile, err := os.Open(args.Input)
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println("json file", result["data"])

	tmpl, err := template.ParseFiles(args.TemplateFile)
	check(err)
	f, err := os.Create(args.Output)
	check(err)
	defer f.Close()
	tmpl.Execute(f, result)

}
