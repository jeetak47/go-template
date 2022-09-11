package main

import (
	"encoding/json"
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

func addFunction() template.FuncMap {
	return template.FuncMap{
		"seq": func(count float64) []uint64 {
			var i uint64
			var Items []uint64
			//cnt, err := strconv.Atoi(count)
			//check(err)
			for i = 0; i < uint64(count); i++ {
				Items = append(Items, i)
			}
			return Items
		},
	}
}

func main() {
	var args struct {
		Input        string `arg:"positional" help:"Imput json file"`
		Output       string `arg:"positional" help:"Output json file"`
		Verbose      bool   `arg:"-v,--verbose" help:"verbosity level"`
		TemplateFile string `arg:"-t,--template" help:"Template file"`
	}
	arg.MustParse(&args)
	jsonFile, err := os.Open(args.Input)
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	tmpl, err := template.New(args.TemplateFile).Funcs(addFunction()).ParseFiles(args.TemplateFile)
	check(err)
	f, err := os.Create(args.Output)
	check(err)
	defer f.Close()
	tmpl.Execute(f, result)

}
