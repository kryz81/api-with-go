package handlers

import (
	"api/types"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"time"
)

const RecipesFile = "../data/recipes.json"

func readRecipes() ([]types.Recipe, error) {
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

func saveRecipes(recipes []types.Recipe) {
	// convert to json
	file, _ := json.MarshalIndent(recipes, "", " ")
	// write json to file
	_ = ioutil.WriteFile(RecipesFile, file, 0644)
}

func GetRecipes(context *gin.Context) {
	recipes, err := readRecipes()
	if err != nil {
		context.JSON(500, gin.H{"msg": err.Error()})
	}
	context.JSON(200, recipes)
}

func NewRecipe(context *gin.Context) {
	var recipe types.Recipe
	// bind body json to the recipe var
	if err := context.ShouldBindJSON(&recipe); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// add custom fields
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()

	// add the new recipe to the list and save all recipes to file
	recipes, err := readRecipes()

	if err != nil {
		context.JSON(500, gin.H{"msg": err.Error()})
	}

	recipes = append(recipes, recipe)
	saveRecipes(recipes)

	context.JSON(201, recipe)
}
