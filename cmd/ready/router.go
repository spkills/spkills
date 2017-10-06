package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
)

var (
	file          = flag.String("file", "route.conf", "Path to route.conf file")
	templatesPath = flag.String("templatesPath", "templates", "Path to route.conf file")
)

type TemplateData struct {
	Name        string
	CapitalName string
}

func main() {
	cli := &Ready{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Create(os.Args))
}

func (r *Ready) Create(args []string) int {
	// オプション引数のパース
	var infile, inDir string
	flags := flag.NewFlagSet("awesome-cli", flag.ContinueOnError)
	flags.SetOutput(r.errStream)
	flag.Parse()
	flags.StringVar(&infile, "file", *file, "Path to route.conf file")
	flags.StringVar(&inDir, "templatesPath", *templatesPath, "Path to route.conf file")

	fmt.Println(inDir)
	fp, err := os.Open(infile)
	if err != nil {
		fmt.Println(err)
		return ExitCodeParseFlagError
	}

	files := make([]string, 0, 10)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		filename := scanner.Text()
		files = append(files, filename)
		r.createHandler(filename, inDir)

	}
	r.createRoutingFile(inDir, files)

	return ExitCodeOK
}

func (r *Ready) createRoutingFile(inDir string, funcs []string) {
	outfile := "controller/routing.go"

	//execute template
	outf, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("cannot createHandler file %q: %s\n", outfile, err)
		panic(err)
	}

	tpl := template.Must(template.ParseFiles(inDir + "/routing.tmpl"))

	var data []TemplateData
	for _, v := range funcs {

		data = append(data,
			TemplateData{
				Name:        v,
				CapitalName: strings.Title(v),
			},
		)
	}

	err = tpl.Execute(outf, data)
	if err != nil {
		panic(err)
	}
}

func (r *Ready) createHandler(filename, inDir string) {

	outfile := "controller/" + filename + ".go"
	fmt.Printf("Compiling %q to %q...\n", filename, outfile)

	if r.fileExists(outfile) {
		fmt.Println("file exists")
		return
	}

	outf, err := os.Create(outfile)
	if err != nil {
		fmt.Printf("cannot createHandler file %q: %s\n", outfile, err)
		panic(err)
	}

	//execute template
	tpl := template.Must(template.ParseFiles(inDir + "/handler.tmpl"))

	data := TemplateData{
		Name:        filename,
		CapitalName: strings.Title(filename),
	}

	err = tpl.Execute(outf, data)
	if err != nil {
		panic(err)
	}
}

func (r *Ready) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
