package main

import "os"

//var (
//	file = flag.String("file", "route.conf", "Path to route.conf file")
//)
/*
func regist(mux *http.ServeMux) {
	mux.HandleFunc("/sandbox", sandboxHandler)
	mux.HandleFunc("/error", errorHandler)
	mux.HandleFunc("/hogehoge", hogehogeHandler)
	mux.HandleFunc("/test", testHandler)
}
*/
func main() {
	cli := &Ready{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

/*
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

	tmpfile := outfile + ".tmp"
	outf, err := os.Create(tmpfile)
	if err != nil {
		fmt.Sprintf("cannot createHandler file %q: %s", tmpfile, err)
		panic(err)
	}

	fmt.Fprint(outf, "package main\n")
	fmt.Fprint(outf, "import(\"fmt\"\n\"net/http\")\n")

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		funcname := scanner.Text()
		fmt.Println(funcname)
		fmt.Fprintf(outf, "func %sHandler(w http.ResponseWriter, r *http.Request){\nfmt.Println(\"test\")\n}\n", funcname)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

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
*/
