package main

import (
	"fmt"
	"log"
	"net/http"

	oas3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

type options struct {
	//Version int    `short:"v" long:"version" description:"OAS Version" required:"optional"`
	File string `short:"f" long:"file" description:"Spec File" required:"true"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.File) == 0 {
		log.Fatal("E_NO_FILE")
	}

	loader := oas3.NewSwaggerLoader()
	spec, err := loader.LoadSwaggerFromFile(opts.File)
	if err != nil {
		log.Fatal(err)
	}

	paths := getEndpointsMetaFlat(spec)

	fmtutil.PrintJSON(paths)

	err = table.WriteCSVSimple([]string{"Tag", "Operation", "URL"}, paths, "api_paths.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")
}

func getEndpointsMetaFlat(spec *oas3.Swagger) [][]string {
	paths := [][]string{}
	for url, path := range spec.Paths {
		paths = appendPathsPathInfo(http.MethodGet, url, path.Get, paths)
	}
	return paths
}

func appendPathsPathInfo(method, url string, info *oas3.Operation, paths [][]string) [][]string {
	if info == nil {
		return paths
	}
	if len(info.Tags) > 0 {
		for _, tag := range info.Tags {
			paths = append(paths, []string{tag, method, url})
		}
	} else {
		paths = append(paths, []string{"", method, url})
	}
	return paths
}
