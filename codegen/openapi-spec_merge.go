package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/swaggman/swagger2"
)

func main() {
	outfile := "openapi-spec_v2.0.0.json"
	dir := "partial-specs_v2.0.0"
	spec, err := swagger2.MergeDirectory(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "MERGE_ERROR")
		log.Fatal(err)
	}

	err = ioutilmore.WriteFileJSON(outfile, spec, 0644, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "WRITE_ERROR")
		log.Fatal(err)
	}

	fmt.Printf("WROTE [%v]\n", outfile)
	fmt.Println("DONE")
}
