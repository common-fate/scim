package patch

import (
	"github.com/common-fate/scim/optional"
	"github.com/common-fate/scim/schema"
)

var (
	patchSchema = schema.Schema{
		ID:          "test:PatchEntity",
		Name:        optional.NewString("Patch"),
		Description: optional.NewString("Patch Entity"),
		Attributes: []schema.CoreAttribute{
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name: "attr1",
			})),
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name:        "multiValued",
				MultiValued: true,
			})),
			schema.ComplexCoreAttribute(schema.ComplexParams{
				Name: "complex",
				SubAttributes: []schema.SimpleParams{
					schema.SimpleStringParams(schema.StringParams{
						Name: "attr1",
					}),
					schema.SimpleStringParams(schema.StringParams{
						Name: "attr2",
					}),
				},
			}),
			schema.ComplexCoreAttribute(schema.ComplexParams{
				Name:        "complexMultiValued",
				MultiValued: true,
				SubAttributes: []schema.SimpleParams{
					schema.SimpleStringParams(schema.StringParams{
						Name: "attr1",
					}),
				},
			}),
		},
	}
	patchSchemaExtension = schema.Schema{
		ID:          "test:PatchExtension",
		Name:        optional.NewString("Patch Extension"),
		Description: optional.NewString("Patch Entity Extension"),
		Attributes: []schema.CoreAttribute{
			schema.SimpleCoreAttribute(schema.SimpleStringParams(schema.StringParams{
				Name: "attr1",
			})),
		},
	}
)
