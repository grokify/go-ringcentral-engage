package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	oas3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/io/ioutilmore"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/grokify/spectrum/openapi3/springopenapi3"
	log "github.com/sirupsen/logrus"
)

func ReadOas3Spec(file string, validate bool) (*oas3.Swagger, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errorsutil.Wrap(err, fmt.Sprintf("Filename [%v]", file))
	}
	swag := &oas3.Swagger{}
	err = swag.UnmarshalJSON(bytes)
	if err != nil {
		return nil, errorsutil.Wrap(err, fmt.Sprintf("Filename [%v]", file))
	}
	if validate {
		err = swag.Validate(context.Background())
		return swag, errorsutil.Wrap(err, fmt.Sprintf("Filename [%v]", file))
	}
	return swag, nil
}

func main() {
	if 1 == 1 {
		file := "../openapi-spec_agents.json"

		swag, err := ReadOas3Spec(file, true)
		if err != nil {
			log.Fatal(err)
		}
		fmtutil.PrintJSON(swag)
		panic("Z")
	}

	bytes, err := ioutil.ReadFile("_class_agent.java")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
	lines := strings.Split(string(bytes), "\n")
	fmtutil.PrintJSON(lines)
	lines = stringsutil.SliceLinesTrim(lines, " \t", true)
	fmtutil.PrintJSON(lines)
	//panic("Z")

	columnsRaw := springopenapi3.ParseSpringCodeColumnsRaw(lines)
	fmtutil.PrintJSON(columnsRaw)

	schema := oas3.Schema{Properties: map[string]*oas3.SchemaRef{}}
	fmtutil.PrintJSON(schema)

	if 1 == 0 {
		line := "private Boolean userManagedByRC;"
		line = "private Boolean userManagedByRC = false;"
		name, prop, err := fromspring.ParseSpringLineToSchema(line)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("NAME [%v]\n", name)
		fmtutil.PrintJSON(prop)
		panic("Z")
	}

	mss, err := springopenapi3.ParseSpringPropertyLinesSliceToSchema(columnsRaw)
	if err != nil {
		log.Info("S1")
		log.Fatal(err)
	}
	fmtutil.PrintJSON(mss)

	err = ioutilmore.WriteFileJSON("_schema_agent.json", mss, 644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE [%v]\n", "_schema_agent.json")

	swag2 := oas3.Swagger{
		Components: oas3.Components{
			Schemas: map[string]*oas3.SchemaRef{
				"Agent": &oas3.SchemaRef{
					Value: &oas3.Schema{
						Properties: mss,
					},
				},
			},
		},
	}
	fmtutil.PrintJSON(swag2)
	swagFile := "_openapi-spec_agents_models.json"
	err = ioutilmore.WriteFileJSON(swagFile, swag2, 0644, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE [%v]\n", swagFile)

	fmt.Println("DONE")
}
