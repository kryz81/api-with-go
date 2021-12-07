package handlers

import (
	"api/services"
	"api/types"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)

func GetRecipes(context *gin.Context) {
	recipes, err := services.ReadRecipes()
	if err != nil {
		context.JSON(500, gin.H{"msg": err.Error()})
	}
	context.JSON(200, recipes)
}

func GetRecipe(context *gin.Context) {
	recipes, err := services.ReadRecipes()
	if err != nil {
		context.JSON(500, gin.H{"msg": err.Error()})
	}

	id := context.Params.ByName("id")
	index := -1
	for i, recipe := range recipes {
		if id == recipe.ID {
			index = i
			break
		}
	}
	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{"msg": "recipe not found"})
		return
	}

	context.JSON(200, recipes[index])
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
	recipes, err := services.ReadRecipes()

	if err != nil {
		context.JSON(500, gin.H{"msg": err.Error()})
	}

	recipes = append(recipes, recipe)
	services.SaveRecipes(recipes)

	context.JSON(201, recipe)
}

func UpdateRecipe(context *gin.Context) {
	var recipe types.Recipe
	// bind body json to the recipe var
	if err := context.ShouldBindJSON(&recipe); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	id := context.Params.ByName("id")
	recipes, _ := services.ReadRecipes()
	index := -1
	for i, recipe := range recipes {
		if id == recipe.ID {
			index = i
			break
		}
	}
	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{"msg": "recipe not found"})
		return
	}

	recipes[index] = types.Recipe{
		ID:          id,
		Name:        recipe.Name,
		Tags:        recipe.Tags,
		Ingredients: recipe.Ingredients,
		PublishedAt: recipes[index].PublishedAt,
	}

	services.SaveRecipes(recipes)

	context.JSON(http.StatusOK, recipes[index])
}

func DeleteRecipe(context *gin.Context) {
	recipes, err := services.ReadRecipes()
	if err != nil {
		context.JSON(500, gin.H{"msg": err.Error()})
	}

	id := context.Params.ByName("id")
	index := -1
	for i, recipe := range recipes {
		if id == recipe.ID {
			index = i
			break
		}
	}
	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{"msg": "recipe not found"})
		return
	}

	newRecipes := make([]types.Recipe, len(recipes)-1)
	place := 0
	for _, recipe := range recipes {
		if recipe.ID != id {
			newRecipes[place] = recipe
			place++
		}
	}
	services.SaveRecipes(newRecipes)

	context.Status(200)
}
