package validators

import (
	"testing"

	"github.com/Nearrivers/dnd-grid-server/api/presenter"
	"github.com/go-playground/validator/v10"
)

func TestValidateLevel(t *testing.T) {
	cases := []struct {
		description string
		level       presenter.Level
		fieldName   string
	}{
		{
			description: "Missing name",
			level: presenter.Level{
				ID:          0,
				Name:        "",
				ImagePath:   "./path",
				GridWidth:   25,
				GridColor:   "#fff",
				GridSpacing: 3,
			},
			fieldName: "Name",
		},
		{
			description: "Missing image path",
			level: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "",
				GridWidth:   25,
				GridColor:   "#fff",
				GridSpacing: 3,
			},
			fieldName: "ImagePath",
		},
		{
			description: "Missing grid width",
			level: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "./path",
				GridColor:   "#fff",
				GridWidth:   0,
				GridSpacing: 3,
			},
			fieldName: "GridWidth",
		},
		{
			description: "Missing grid color",
			level: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "./path",
				GridWidth:   25,
				GridSpacing: 3,
				GridColor:   "",
			},
			fieldName: "GridColor",
		},
		{
			description: "Missing grid spacing",
			level: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "./path",
				GridWidth:   25,
				GridColor:   "#fff",
				GridSpacing: 0,
			},
			fieldName: "GridSpacing",
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			err := ValidateLevel(tt.level)

			if err == nil {
				t.Fatal("didn't get an error but should have")
			}

			if _, ok := err.(validator.FieldError); !ok {
				t.Fatalf("error isn't of type validator.FieldError, got %v", err)
			}
		})
	}
}
