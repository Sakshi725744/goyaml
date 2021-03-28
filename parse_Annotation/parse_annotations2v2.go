package parse

import (
	"fmt"

	"github.com/spdx/yaml/model"
	spdx "github.com/spdx/yaml/spdx"
)

func ParseAnnotation2v2(parser *model.SpdxYaml) (annotations []spdx.Annotation2_2) {
	//type annotations *[]spdx.Annotation2_2
	a := spdx.Annotation2_2{}

	for i := 0; i < len(parser.Annotations); i++ {

		//Annotation Date in 2v2
		a.AnnotationDate = (*parser).Annotations[i].AnnotationDate

		//Annotation Comment in 2v2
		a.AnnotationComment = (*parser).Annotations[i].Comment

		//Annotation Type in 2v2
		setAnnotationType2v2((*parser).Annotations[i].AnnotationType, &a)

		//Annotator and Annotator Type in 2v2
		setAnnotatorFromString2v2((*parser).Annotations[i].Annotator, &a)

		(annotations) = append((annotations), a)

	}
	return annotations
}

func setAnnotatorFromString2v2(annotatorString string, ann *spdx.Annotation2_2) error {
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

func setAnnotationType2v2(annType string, ann *spdx.Annotation2_2) error {
	if annType == "OTHER" || annType == "REVIEW" {
		ann.AnnotationType = annType
	} else {
		return fmt.Errorf("unknown annotation type %s", annType)
	}
	return nil
}
