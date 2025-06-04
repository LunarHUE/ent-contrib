package entproto

import (
	"fmt"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema"
	"github.com/go-viper/mapstructure/v2"
)

// const FieldAnnotation = "ProtoField"
const AutoFieldAnnotation = "ProtoAutoField"

// type FieldOption func(*pbfield)
type AutoFieldOption func(*pbfield)

func AutoField(num int, options ...FieldOption) schema.Annotation {
	f := pbfield{Number: num}
	for _, apply := range options {
		apply(&f)
	}
	return f
}

func extractAutoFieldAnnotation(fld *gen.Field) (*pbfield, error) {
	annot, ok := fld.Annotations[AutoFieldAnnotation]
	if !ok {
		return nil, fmt.Errorf("field %q is missing %q annotation", fld.Name, AutoFieldAnnotation)
	}
	var pbf pbfield
	if err := mapstructure.Decode(annot, &pbf); err != nil {
		return nil, fmt.Errorf("decoding %q annotation for field %q: %w", AutoFieldAnnotation, fld.Name, err)
	}
	if pbf.Number == 0 {
		return nil, fmt.Errorf("field %q has invalid number in %q annotation", fld.Name, AutoFieldAnnotation)
	}
	return &pbf, nil
}
