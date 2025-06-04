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

type autoPBField struct{ pbfield }

func (autoPBField) Name() string { return AutoFieldAnnotation }

// AutoField annotates a field for automatic protobuf numbering. Options may be
// provided to override the protobuf type or type name. The field number is
// assigned during code generation.
func AutoField(options ...FieldOption) schema.Annotation {
	f := autoPBField{}
	for _, apply := range options {
		apply(&f.pbfield)
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
	return &pbf, nil
}

func extractAutoEdgeAnnotation(edge *gen.Edge) (*pbfield, error) {
	annot, ok := edge.Annotations[AutoFieldAnnotation]
	if !ok {
		return nil, fmt.Errorf("edge %q is missing %q annotation", edge.Name, AutoFieldAnnotation)
	}
	var pbf pbfield
	if err := mapstructure.Decode(annot, &pbf); err != nil {
		return nil, fmt.Errorf("decoding %q annotation for edge %q: %w", AutoFieldAnnotation, edge.Name, err)
	}
	return &pbf, nil
}
