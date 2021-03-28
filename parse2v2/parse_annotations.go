package parse2v2

import (
	"fmt"
	"strings"

	"github.com/spdx/yaml/model"
	spdx "github.com/spdx/yaml/spdx"
)

func ParseAnnotation(parser *model.SpdxYaml) (annotations []spdx.Annotation2_2) {
	//type annotations *[]spdx.Annotation2_2
	a := spdx.Annotation2_2{}

	for i := 0; i < len(parser.Annotations); i++ {

		//Annotation Date in 2v2
		a.AnnotationDate = (*parser).Annotations[i].AnnotationDate

		//Annotation Comment in 2v2
		a.AnnotationComment = (*parser).Annotations[i].Comment

		//Annotation Type in 2v2
		setAnnotationType((*parser).Annotations[i].AnnotationType, &a)

		//Annotator and Annotator Type in 2v2
		setAnnotatorFromString((*parser).Annotations[i].Annotator, &a)

		(annotations) = append((annotations), a)

	}
	return annotations
}

func setAnnotatorFromString(annotatorString string, ann *spdx.Annotation2_2) error {
	subkey, subvalue, err := ExtractSubs(annotatorString, ":")
	if err != nil {
		return err
	}
	if subkey == "Person" || subkey == "Organization" || subkey == "Tool" {
		ann.AnnotatorType = subkey
		ann.Annotator = subvalue
		return nil
	}
	return fmt.Errorf("unrecognized Annotator type %v while parsing annotation", subkey)
}
func ExtractSubs(value string, sep string) (string, string, error) {
	// parse the value to see if it's a valid subvalue format
	sp := strings.SplitN(value, sep, 2)
	if len(sp) == 1 {
		return "", "", fmt.Errorf("invalid subvalue format for %s (no %s found)", value, sep)
	}

	subkey := strings.TrimSpace(sp[0])
	subvalue := strings.TrimSpace(sp[1])

	return subkey, subvalue, nil
}

func setAnnotationType(annType string, ann *spdx.Annotation2_2) error {
	if annType == "OTHER" || annType == "REVIEW" {
		ann.AnnotationType = annType
	} else {
		return fmt.Errorf("unknown annotation type %s", annType)
	}
	return nil
}
