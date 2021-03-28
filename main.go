package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spdx/yaml/model"
	"github.com/spdx/yaml/parse2v2"
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

	annotations := parse2v2.ParseAnnotation(c)

	for i := 0; i < len(annotations); i++ {
		fmt.Println((annotations)[i].AnnotationComment)
		fmt.Println((annotations)[i].AnnotationType)
		fmt.Println((annotations)[i].AnnotatorType)
		fmt.Println((annotations)[i].Annotator)
		fmt.Println((annotations)[i].AnnotationDate)
	}
}
