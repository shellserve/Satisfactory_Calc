// Package recipes holds loader logic and models used in collecting recipe information from the satisfactory master schema
// model.go: model structs
// loader.go: loader logic
//
// TODO
// UnlockedBy is currently not used.
//
//	This can either be a string or array of strings
//	we need a normalization function to handle this.
package recipes

type ItemAmount struct {
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
}

type RecipeEntry struct {
	ClassName                string       `json:"className"`
	Name                     string       `json:"name"`
	UnlockedBy               any          `json:"-"`
	Duration                 float64      `json:"duration"`
	Ingredients              []ItemAmount `json:"ingredients"`
	Products                 []ItemAmount `json:"products"`
	ProducedIn               []string     `json:"producedIn"`
	InCraftBench             bool         `json:"inCraftBench"`
	InWorkshop               bool         `json:"inWorkshop"`
	InBuildGun               bool         `json:"inBuildGun"`
	InCustomizer             bool         `json:"inCustomizer"`
	ManualCraftingMultiplier float64      `json:"manualCraftingMultiplier"`
	Alternate                bool         `json:"alternate"`
	MinPower                 *float64     `json:"minPower"`
	MaxPower                 *float64     `json:"maxPower"`
	Seasons                  []string     `json:"seasons"`
	Stable                   bool         `json:"stable"`
	Experimental             bool         `json:"experimental"`
}

type Recipes map[string][]RecipeEntry
