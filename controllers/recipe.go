package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Marcello732/goApi/database"
	"github.com/Marcello732/goApi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecipeRepo struct {
	Db *gorm.DB
}

func NewRecipeRepo() *RecipeRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Recipe{}, &models.User{}, &models.Stickers{}, &models.Ingredients{}, &models.Preparations{}, &models.Images{})
	return &RecipeRepo{Db: db}
}

//create recipe
func (repository *RecipeRepo) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	c.BindJSON(&recipe)
	err := models.CreateRecipe(repository.Db, &recipe)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

//get recipes
func (repository *RecipeRepo) GetRecipes(c *gin.Context) {
	var recipe []models.Recipe
	err := models.GetRecipes(repository.Db, &recipe)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

//get recipe by id
func (repository *RecipeRepo) GetRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var recipe models.Recipe
	var stickers []models.Stickers
	var ingredientsAndAmounts []models.Ingredients
	var preparationSteps []models.Preparations
	var recipeImages []models.Images
	var author models.User
	err := models.GetRecipe(repository.Db, &recipe, id)
	err2 := models.GetStickers(repository.Db, id, &stickers)
	err3 := models.GetIngredientsAndAmounts(repository.Db, id, &ingredientsAndAmounts)
	err4 := models.GetPreparationSteps(repository.Db, id, &preparationSteps)
	err5 := models.GetRecipeImages(repository.Db, id, &recipeImages)
	err6 := models.GetAuthor(repository.Db, int(recipe.AuthorID), &author)
	recipe.Stickers = stickers
	recipe.IngredientsAndAmounts = ingredientsAndAmounts
	recipe.PreparationSteps = preparationSteps
	recipe.RecipeImages = recipeImages
	recipe.Author = author
	if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

// update recipe
func (repository *RecipeRepo) UpdateRecipe(c *gin.Context) {
	var recipe models.Recipe
	//var stickers []models.Stickers
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetRecipe(repository.Db, &recipe, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&recipe)
	err = models.UpdateRecipe(repository.Db, &recipe)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

// delete recipe
func (repository *RecipeRepo) DeleteRecipe(c *gin.Context) {
	var recipe models.Recipe
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteRecipe(repository.Db, &recipe, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}