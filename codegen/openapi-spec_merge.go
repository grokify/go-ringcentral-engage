package main

import (
	"fmt"
	"log"

	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/swaggman/swagger2"
)

func main() {
	outfile := "openapi-spec.json"
	dir := "partial-specs"
	spec, err := swagger2.MergeDirectory(dir)
	if err != nil {
		fmt.Println("MERGE_ERROR")
		log.Fatal(err)
	}

	err = ioutilmore.WriteFileJSON(outfile, spec, 0644, "", "  ")
	if err != nil {
		fmt.Println("WRITE_ERROR")
		log.Fatal(err)
	}

	fmt.Printf("WROTE [%v]\n", outfile)
	fmt.Println("DONE")
}
