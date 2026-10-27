package recipe

import (
	"fmt"

	"github.com/lurkshark/fireweed/packages/assets/model"
)

func BuiltinRecipe(name string) (model.Recipe, error) {
	return model.Recipe{}, fmt.Errorf("recipe %q is not implemented", name)
}
