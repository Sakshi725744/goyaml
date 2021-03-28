package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spdx/yaml/model"
	parse "github.com/spdx/yaml/parse_Annotation"
	"gopkg.in/yaml.v2"
)

func readConf(filename string) (*model.SpdxYaml, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &model.SpdxYaml{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func main() {
	c, err := readConf("/home/sakshi7parikh/SPDX-YAML/examples/spdxYaml.yaml")
	if err != nil {
		log.Fatal(err)
	}

	annotations2v2 := parse.ParseAnnotation2v2(c)

	for i := 0; i < len(annotations2v2); i++ {
		fmt.Println((annotations2v2)[i].AnnotationComment)
		fmt.Println((annotations2v2)[i].AnnotationType)
		fmt.Println((annotations2v2)[i].AnnotatorType)
		fmt.Println((annotations2v2)[i].Annotator)
		fmt.Println((annotations2v2)[i].AnnotationDate)
	}
	annotations2v1 := parse.ParseAnnotation2v1(c)

	for i := 0; i < len(annotations2v1); i++ {
		fmt.Println((annotations2v1)[i].AnnotationComment)
		fmt.Println((annotations2v1)[i].AnnotationType)
		fmt.Println((annotations2v1)[i].AnnotatorType)
		fmt.Println((annotations2v1)[i].Annotator)
		fmt.Println((annotations2v1)[i].AnnotationDate)
	}
}
