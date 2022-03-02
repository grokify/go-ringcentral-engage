package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	oas3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/gocharts/data/table"
	"github.com/grokify/mogo/fmt/fmtutil"
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

	loader := oas3.NewLoader()
	spec, err := loader.LoadFromFile(opts.File)
	if err != nil {
		log.Fatal(err)
	}

	paths := getEndpointsMetaFlat(spec)
	SortSoS(2, paths)

	fmtutil.PrintJSON(paths)

	err = table.WriteCSVSimple([]string{"Tag", "Operation", "URL", "Summary"}, paths, "api_paths.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")
}

func getEndpointsMetaFlat(spec *oas3.T) [][]string {
	paths := [][]string{}
	for url, path := range spec.Paths {
		paths = appendPathsPathInfo(http.MethodConnect, url, path.Connect, paths)
		paths = appendPathsPathInfo(http.MethodGet, url, path.Get, paths)
		paths = appendPathsPathInfo(http.MethodOptions, url, path.Options, paths)
		paths = appendPathsPathInfo(http.MethodPatch, url, path.Patch, paths)
		paths = appendPathsPathInfo(http.MethodPost, url, path.Post, paths)
		paths = appendPathsPathInfo(http.MethodPut, url, path.Put, paths)
		paths = appendPathsPathInfo(http.MethodDelete, url, path.Delete, paths)
		paths = appendPathsPathInfo(http.MethodTrace, url, path.Trace, paths)
	}
	return paths
}

func appendPathsPathInfo(method, url string, op *oas3.Operation, paths [][]string) [][]string {
	if op == nil {
		return paths
	}
	if len(op.Tags) > 0 {
		for _, tag := range op.Tags {
			paths = append(paths, []string{tag, method, url, op.Summary})
		}
	} else {
		paths = append(paths, []string{"", method, url, op.Summary})
	}
	return paths
}

func SortSoS(idx int, s [][]string) {
	sort.Slice(s, func(i, j int) bool {
		// edge cases
		if len(s[i][idx]) == 0 && len(s[j][idx]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(s[i][idx]) == 0 || len(s[j][idx]) == 0 {
			return len(s[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}

		// both slices len() > 0, so can test this now:
		return s[i][idx][0] < s[j][idx][0]
	})
}
