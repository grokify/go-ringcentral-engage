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
		log.Fatal()
	}

	err = ioutilmore.WriteFileJSON(outfile, spec, 0644, "", "  ")
	if err != nil {
		log.Fatal()
	}

	fmt.Printf("WROTE [%v]\n", outfile)
	fmt.Println("DONE")
}
