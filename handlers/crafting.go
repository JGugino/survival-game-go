package handlers

import (
	"errors"

	"github.com/JGugino/survival-game-go/utils"
)

type CraftingRecipe struct {
	Id          string
	RecipeItems []RecipeItem
	OutputItem  utils.ItemId
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
		OutputItem: utils.I_PICKAXE,
	}
}

func (c *Crafting) CanCraftItem(recipeId string) (canCaft bool, foundStacks []*utils.ItemStack, missingItems []RecipeItem, err error) {
	recipe, ok := c.CraftingRecipes[recipeId]

	if !ok {
		return false, []*utils.ItemStack{}, []RecipeItem{}, errors.New("no-recipe")
	}

	missingItems = make([]RecipeItem, 0)
	foundStacks = make([]*utils.ItemStack, 0)

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

		foundStacks = append(foundStacks, stack)
	}

	if len(missingItems) > 0 {
		return false, foundStacks, missingItems, errors.New("missing-items")
	}

	return true, foundStacks, missingItems, nil
}
