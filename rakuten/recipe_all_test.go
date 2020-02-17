package rakuten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecipeRanking(t *testing.T) {
	params := &RecipeRankingParams{
		CategoryID: "46",
	}
	_, _, err := client.Recipe.Ranking(ctx, params)
	assert.Nil(t, err)
}

func TestRecipecategory(t *testing.T) {
	params := &RecipeCategoryParams{
		CategoryType: "small",
		// Category Type は small, medium, large のうちどれか選択
	}
	_, _, err := client.Recipe.CategoryList(ctx, params)
	assert.Nil(t, err)
}
