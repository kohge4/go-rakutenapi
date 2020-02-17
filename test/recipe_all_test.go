package test

import (
	"testing"

	"github.com/kohge4/go-rakutenapi/rakuten"

	"github.com/stretchr/testify/assert"
)

func TestRecipeRanking(t *testing.T) {
	params := &rakuten.RecipeRankingParams{
		CategoryID: "46",
	}
	_, _, err := client.Recipe.Ranking(ctx, params)
	assert.Nil(t, err)
}

func TestRecipecategory(t *testing.T) {
	params := &rakuten.RecipeCategoryParams{
		CategoryType: "small",
		// Category Type は small, medium, large のうちどれか選択
	}
	_, _, err := client.Recipe.CategoryList(ctx, params)
	assert.Nil(t, err)
}
