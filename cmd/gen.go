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
	flag.Parse()

	if len(*file) > 0 {
		compileFile(*file)
		return
	}
}

func compileSingleFile(filename string) {
	var fp *os.File
	var err error

	fp, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	compileFile(filename)
}

func compileFile(infile string) {
	outfile := infile + ".go"
	fmt.Printf("Compiling %q to %q...", infile, outfile)

	var fp *os.File
	var err error

	fp, err = os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	inf, err := os.Open(infile)
	if err != nil {
		fmt.Printf("cannot open file %q: %s", infile, err)
		panic(err)
	}

	tmpfile := outfile + ".tmp"
	outf, err := os.Create(tmpfile)
	if err != nil {
		fmt.Sprintf("cannot create file %q: %s", tmpfile, err)
		panic(err)
	}

	fmt.Fprint(outf, "package main\n")
	fmt.Fprint(outf, "import(\"fmt\")\n")

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		funcname := scanner.Text()
		fmt.Println(funcname)
		fmt.Fprintf(outf, "func %sHandler(){\nfmt.Println(\"test\")\n}\n", funcname)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err = outf.Close(); err != nil {
		fmt.Printf("error when closing file %q: %s", tmpfile, err)
		panic(err)
	}
	if err = inf.Close(); err != nil {
		fmt.Printf("error when closing file %q: %s", infile, err)
		panic(err)
	}

	// prettify the output file
	uglyCode, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		fmt.Printf("cannot read file %q: %s", tmpfile, err)
	}

	prettyCode, err := format.Source(uglyCode)
	if err != nil {
		fmt.Printf("error when formatting compiled code for %q: %s. See %q for details", infile, err, tmpfile)
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
