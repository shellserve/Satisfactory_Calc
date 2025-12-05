package recipes

import (
	"encoding/json"
	"fmt"
)

func LoadRecipeFromJSON(jsonBytes []byte) (Recipes, error) {
	var r Recipes
	if err := json.Unmarshal(jsonBytes, &r); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return r, nil
}
