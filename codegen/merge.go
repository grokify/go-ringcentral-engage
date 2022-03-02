package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/io/ioutilmore"
	"github.com/grokify/mogo/path/filepathutil"
	"github.com/grokify/spectrum/openapi2"
	"github.com/grokify/spectrum/openapi3"
	"github.com/jessevdk/go-flags"

	oas3 "github.com/getkin/kin-openapi/openapi3"
)

type options struct {
	Directory string `short:"d" long:"directory" description:"OAS Directory" required:"true"`
	Version   int    `short:"v" long:"version" description:"OAS Version" required:"true"`
}

func MergeOAS2(dir, outfile string) error {
	spec, err := openapi2.MergeDirectory(dir)
	if err != nil {
		return errorsutil.Wrap(err, "E_MERGE_FAILED")
	}

	err = ioutilmore.WriteFileJSON(outfile, spec, 0644, "", "  ")
	if err != nil {
		return errorsutil.Wrap(err, "E_WRITE_FAILED")
	}
	fmt.Printf("WROTE [%v]\n", outfile)
	return nil
}

func MergeOAS3(dir, outfile string) error {
	fmt.Println("START_OAS3\n")
	spec, err := openapi3.MergeDirectory(dir)
	if err != nil {
		return errorsutil.Wrap(err, "E_MERGE_DIRECTORY_FAILED")
	}
	if 1 == 0 {
		schemas := []string{"ReferenceObject", "Agent"}
		for i, sch := range schemas {
			has := openapi3.SpecHasComponentSchema(spec, sch, false)
			hasStr := "YES"
			if !has {
				hasStr = "NO"
			}
			fmt.Printf("[%v] [%v] [%v]\n", i, sch, hasStr)
		}

		if 1 == 0 {
			err = spec.Validate(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
		bytes, err := json.Marshal(spec)
		if err != nil {
			log.Fatal(err)
		}
		loader := oas3.NewSwaggerLoader()
		spec2, err := loader.LoadSwaggerFromData(bytes)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(spec2)
		err = spec2.Validate(loader.Context)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("VALIDATED")
		//fmtutil.PrintJSON
		panic("Z")
	}
	bytes, err := spec.MarshalJSON()
	if err != nil {
		return errorsutil.Wrap(err, "E_JSON_ENCODING_FAILED")
	}

	err = ioutil.WriteFile(outfile, bytes, 0644)
	if err != nil {
		return errorsutil.Wrap(err, "E_WRITE_FAILED")
	}
	fmt.Printf("WROTE [%v]\n", outfile)
	return nil
}

func main() {
	if 1 == 0 {
		file := "specs-voice_v3.0.0/openapi-spec_campaigns.json"
		file = "specs-voice_v3.0.0/openapi-spec_countries.json"
		file = "specs-example_v3.0.0.json"
		file = "specs-voice_v3.0.0/openapi-spec_agent-groups.json"
		file = "specs-voice_v3.0.0/openapi-spec_agents.json"
		if 1 == 1 {
			spec, err := openapi3.ReadFile(file, false)
			if err != nil {
				fmt.Println("TEST_UNMARSHAL")
				log.Fatal(err)
			}
			fmtutil.PrintJSON(spec)
			name := "Country"
			fmtutil.PrintJSON(spec.Components.Schemas[name])

			fmt.Printf("HAS [%v][%v]\n", name, openapi3.SpecHasComponentSchema(spec, name, false))
			err = spec.Validate(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			panic("READ_FILE_NEW")
		}
		spec, err := openapi3.ReadFileLoader(file)
		if err != nil {
			fmt.Println("TEST_LOADER")
			log.Fatal(err)
		}
		fmtutil.PrintJSON(spec)
		err = spec.Validate(context.Background())
		if err != nil {
			fmt.Println("TEST")
			log.Fatal(err)
		}
		fmtutil.PrintJSON(spec)

		fmt.Println("DONE")
		panic("Z")
	}
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	version := 2
	if opts.Version == 2 || opts.Version == 3 {
		version = opts.Version
	}
	dir := opts.Directory

	dir = filepathutil.TrimRight(dir)

	_, leaf := filepath.Split(dir)

	outfile := leaf + ".json"

	//outfile := opts.Directory + ".json"
	//outfile := fmt.Sprintf("openapi-spec_v%d.0.0.json", version)
	//dir := fmt.Sprintf("partial-specs_v%d.0.0", version)

	switch version {
	case 2:
		err = MergeOAS2(dir, outfile)
	case 3:
		err = MergeOAS3(dir, outfile)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
