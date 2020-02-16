package test

import (
	"testing"

	"github.com/kohge4/go-rakutenapi/rakuten"

	"github.com/stretchr/testify/assert"
)

func TestIchibaItemSearch(t *testing.T) {
	params := &rakuten.IchibaItemSearchParams{
		Keyword: "オードリー",
		Hits:    5,
		GenreID: 204209,
	}
	resp, _, err := client.Ichiba.Search(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestIchibaGenreSearch(t *testing.T) {
	params := &rakuten.IchibaGenreSearchParams{
		GenreID: 204209,
	}
	_, _, err := client.Ichiba.GenreSearch(ctx, params)
	assert.Nil(t, err)
}

func TestIchibaKakakunavi(t *testing.T) {
	params := &rakuten.IchibaProductSearchParams{
		Keyword: "ハービーハンコック",
		Hits:    5,
	}
	resp, _, err := client.Ichiba.ProductSearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestIchibaRanking(t *testing.T) {
	params := &rakuten.IchibaRankingParams{
		GenreID: 204209,
	}
	_, _, err := client.Ichiba.Ranking(ctx, params)
	assert.Nil(t, err)
}

func TestIchibaTagSearch(t *testing.T) {
	params := &rakuten.IchibaTagParams{
		TagID: 1000317,
	}
	_, _, err := client.Ichiba.TagSearch(ctx, params)
	assert.Nil(t, err)
}
