package rakuten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKoboGenre(t *testing.T) {
	params := &KoboGenreParams{
		KoboGenreID: 101912005001,
	}
	_, _, err := client.Kobo.GenreSearch(ctx, params)
	assert.Nil(t, err)
}

func TestKoboEbooks(t *testing.T) {
	params := &KoboEbooksParams{
		Keyword: "Go言語",
	}
	_, _, err := client.Kobo.EbooksSearch(ctx, params)
	assert.Nil(t, err)
}
