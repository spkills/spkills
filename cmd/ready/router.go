package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
)

var (
	file = flag.String("file", "route.conf", "Path to route.conf file")
)

func main() {
	cli := &Ready{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (r *Ready) Run(args []string) int {

	// オプション引数のパース
	var infile string
	flags := flag.NewFlagSet("awesome-cli", flag.ContinueOnError)
	flags.SetOutput(r.errStream)
	flags.StringVar(&infile, "file", "route.conf", "Path to route.conf file")

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
		r.createHandler(filename)

	}
	r.createRegister("register", files)

	return ExitCodeOK
}

/*
func regist(mux *http.ServeMux) {
	mux.HandleFunc("/sandbox", sandboxHandler)
	mux.HandleFunc("/error", errorHandler)
	mux.HandleFunc("/hogehoge", hogehogeHandler)
	mux.HandleFunc("/test", testHandler)
}
*/
func (r *Ready) createRegister(filename string, funcs []string) {
	outfile := "controller/" + filename + ".go"
	tmpfile := outfile + ".tmp"

	outf, err := os.Create(tmpfile)
	if err != nil {
		fmt.Sprintf("cannot createHandler file %q: %s", tmpfile, err)
		panic(err)
	}

	fmt.Printf("Compiling %q to %q...", filename, outfile)
	fmt.Fprint(outf, "package controller\n")
	fmt.Fprint(outf, "import(\"github.com/buaazp/fasthttprouter\")\n")
	fmt.Fprint(outf, "func Regist(router *fasthttprouter.Router) {\n")

	for _, v := range funcs {
		fmt.Fprintf(outf, "router.GET(\"/%s\", %sHandler)\n", v, v)
	}

	fmt.Fprint(outf, "}")
	if err = outf.Close(); err != nil {
		fmt.Printf("error when closing file %q: %s", tmpfile, err)
		panic(err)
	}

	// prettify the output file
	uglyCode, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		fmt.Printf("cannot read file %q: %s", tmpfile, err)
	}

	prettyCode, err := format.Source(uglyCode)
	if err != nil {
		fmt.Printf("error when formatting compiled code for %q: %s. See %q for details", filename, err, tmpfile)
		panic(err)
	}

	if err = ioutil.WriteFile(outfile, prettyCode, 0666); err != nil {
		fmt.Printf("error when writing file %q: %s", outfile, err)
		panic(err)
	}
	if err = os.Remove(tmpfile); err != nil {
		fmt.Printf("error when removing file %q: %s", tmpfile, err)
		panic(err)
	}
}

func (r *Ready) createHandler(filename string) {

	outfile := "controller/" + filename + ".go"
	fmt.Printf("Compiling %q to %q...", filename, outfile)

	if r.fileExists(outfile) {
		fmt.Println("file exists")
		return
	}

	tmpfile := outfile + ".tmp"
	outf, err := os.Create(tmpfile)
	if err != nil {
		fmt.Sprintf("cannot createHandler file %q: %s", tmpfile, err)
		panic(err)
	}

	fmt.Fprint(outf, "package controller\n")
	fmt.Fprint(outf, "import(\"fmt\"\n\"github.com/valyala/fasthttp\")\n")
	fmt.Println(filename)
	fmt.Fprintf(outf, "func %sHandler(ctx *fasthttp.RequestCtx){\nfmt.Println(\"test\")\n}\n", filename)

	if err = outf.Close(); err != nil {
		fmt.Printf("error when closing file %q: %s", tmpfile, err)
		panic(err)
	}

	// prettify the output file
	uglyCode, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		fmt.Printf("cannot read file %q: %s", tmpfile, err)
	}

	prettyCode, err := format.Source(uglyCode)
	if err != nil {
		fmt.Printf("error when formatting compiled code for %q: %s. See %q for details", filename, err, tmpfile)
		panic(err)
	}

	if err = ioutil.WriteFile(outfile, prettyCode, 0666); err != nil {
		fmt.Printf("error when writing file %q: %s", outfile, err)
		panic(err)
	}
	if err = os.Remove(tmpfile); err != nil {
		fmt.Printf("error when removing file %q: %s", tmpfile, err)
		panic(err)
	}
}

func (r *Ready) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
