// Package generators contains all your models generators; functions used to generate mocked data and help you testing
package generators

import "github.com/go-viper/mapstructure/v2"

type testingI interface {
	Fatal(args ...any)
}

func decode(t testingI, args map[string]any, holder any) {
	if err := mapstructure.Decode(args, &holder); err != nil {
		t.Fatal(err)
	}
}
