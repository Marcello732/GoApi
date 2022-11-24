package models

import (
	"gorm.io/gorm"
)

type Stickers struct {
	RecipeID uint
	Sticker string
}

type Ingredients struct {
	RecipeID uint
	Ingredient string
	Amount string
}

type Preparations struct {
	RecipeID uint
	Step string
}

type Images struct {
	RecipeID uint
	Image string
}

type Recipe struct {
	//gorm.Model
	ID              uint     
	Name            string  
	Description     string  
	PreparationTime int     
	DifficultyLevel int     
	Rating          float64 
	Calories 		float64
	Proteins 		float64 
	Carbs    		float64 
	Fats     		float64
	AuthorID 		uint
	Author			User 
	Stickers 		[]Stickers
	IngredientsAndAmounts []Ingredients
	PreparationSteps 	[]Preparations
	RecipeImages []Images
}

//create a recipe
func CreateRecipe(db *gorm.DB, Recipe *Recipe) (err error) {
	err = db.Create(Recipe).Error
	if err != nil {
		return err
	}
	return nil
}

//get recipes
func GetRecipes(db *gorm.DB, Recipe *[]Recipe) (err error) {
	err = db.Find(Recipe).Error
	if err != nil {
		return err
	}
	return nil
}

//get recipe by id
func GetRecipe(db *gorm.DB, Recipe *Recipe, id int) (err error) {
	err = db.Where("ID = ?", id).First(Recipe).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAuthor(db *gorm.DB, id int, User *User) (err error) {
	err = db.Where("ID = ?", id).Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetStickers(db *gorm.DB, id int, Stickers *[]Stickers) (err error) {
	err = db.Where("RecipeID = ?", id).Find(Stickers).Error
	if err != nil {
		return err
	}
	return nil
}

func GetIngredientsAndAmounts(db *gorm.DB, id int, IngredientsAndAmounts *[]Ingredients) (err error) {
	err = db.Where("RecipeID = ?", id).Find(IngredientsAndAmounts).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPreparationSteps(db *gorm.DB, id int, PreparationSteps *[]Preparations) (err error) {
	err = db.Where("RecipeID = ?", id).Find(PreparationSteps).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRecipeImages(db *gorm.DB, id int, RecipeImages *[]Images) (err error) {
	err = db.Where("RecipeID = ?", id).Find(RecipeImages).Error
	if err != nil {
		return err
	}
	return nil
}

//update recipe
func UpdateRecipe(db *gorm.DB, Recipe *Recipe) (err error) {
	db.Save(Recipe)
	return nil
}

//delete recipe
func DeleteRecipe(db *gorm.DB, Recipe *Recipe, id int) (err error) {
	db.Where("ID = ?", id).Delete(Recipe)
	return nil
}