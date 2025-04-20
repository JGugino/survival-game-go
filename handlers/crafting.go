package handlers

import (
	"errors"

	"github.com/JGugino/survival-game-go/utils"
)

type CraftingRecipe struct {
	Id          string
	RecipeItems []RecipeItem
}

type RecipeItem struct {
	Id       utils.ItemId
	Quantity int
}

type Crafting struct {
	CraftingRecipes map[string]CraftingRecipe
}

func (c *Crafting) InitCraftingRecipes() {
	c.CraftingRecipes["pickaxe"] = CraftingRecipe{
		Id: "pickaxe",
		RecipeItems: []RecipeItem{
			RecipeItem{Id: utils.I_ROCK, Quantity: 3},
			RecipeItem{Id: utils.I_WOOD, Quantity: 2},
		},
	}
}

func (c *Crafting) CanCraftItem(recipeId string) (canCaft bool, missingItems []RecipeItem, err error) {
	recipe, ok := c.CraftingRecipes[recipeId]

	if !ok {
		return false, []RecipeItem{}, errors.New("no-recipe")
	}

	missingItems = make([]RecipeItem, 0)

	for _, i := range recipe.RecipeItems {
		_, stack, err := utils.GetItemStackByItemId(i.Id)
		if err != nil {
			missingItems = append(missingItems, i)
			continue
		}

		if stack.StackSize < i.Quantity {
			missingItems = append(missingItems, i)
			continue
		}
	}

	if len(missingItems) > 0 {
		return false, missingItems, errors.New("missing-items")
	}

	return true, missingItems, nil
}
