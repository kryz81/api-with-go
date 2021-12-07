package services

import (
	"api/types"
	"encoding/json"
	"io/ioutil"
)

const RecipesFile = "data/recipes.json"

func ReadRecipes() ([]types.Recipe, error) {
	recipes := make([]types.Recipe, 0)
	// read the content of the file
	file, err := ioutil.ReadFile(RecipesFile)
	if err != nil {
		return recipes, err
	}
	// parse json and save result to the recipes
	err = json.Unmarshal(file, &recipes)
	if err != nil {
		return recipes, err
	}
	return recipes, nil
}

func SaveRecipes(recipes []types.Recipe) {
	// convert to json
	file, _ := json.MarshalIndent(recipes, "", " ")
	// write json to file
	_ = ioutil.WriteFile(RecipesFile, file, 0644)
}
