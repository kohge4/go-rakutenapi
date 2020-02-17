package rakuten

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	client *Client
	ctx    context.Context
)

func init() {
	client = NewTestClient()
	ctx = context.Background()
}

func TestBooksTotalSearch(t *testing.T) {
	params := &BooksTotalSearchParams{
		Keyword: "伝説の",
	}
	_, _, err := client.Books.TotalSearch(ctx, params)
	assert.Nil(t, err)
}

func TestBooksBookSearch(t *testing.T) {
	params := &BookSearchParams{
		Title: "ドメイン駆動設計",
	}
	_, _, err := client.Books.BookSearch(ctx, params)
	assert.Nil(t, err)
}

func TestBooksCDSearch(t *testing.T) {
	params := &BooksCDSearchParams{
		ArtistName: "cero",
	}
	_, _, err := client.Books.CDSearch(ctx, params)
	assert.Nil(t, err)
}

func TestBooksDvdBluraySearch(t *testing.T) {
	params := &BooksDvdBluraySearchParams{
		ArtistName: "吉岡里穂",
		Hits:       5,
	}
	resp, _, err := client.Books.DvdBluraySearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestBooksForeignBookSearch(t *testing.T) {
	params := &BooksForeignBookSearchParams{
		Title: "ninja",
		Hits:  5,
	}
	resp, _, err := client.Books.ForeignBookSearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestBooksGameSearch(t *testing.T) {
	params := &BooksGameSearchParams{
		Hardware: "ps4",
		Hits:     5,
	}
	resp, _, err := client.Books.GameSearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestBooksGenreSearch(t *testing.T) {
	params := &BooksGenreSearchParams{
		BooksGenreID: "006513001",
	}
	_, _, err := client.Books.GenreSearch(ctx, params)
	assert.Nil(t, err)
}

func TestBooksMagazineSearch(t *testing.T) {
	params := &BooksMagazineSearchParams{
		Title: "ジャンプ",
		Hits:  5,
	}
	resp, _, err := client.Books.MagazineSearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestBooksSoftwareSearch(t *testing.T) {
	params := &BooksSoftwareSearchParams{
		OS:   "Mac",
		Hits: 10,
	}
	resp, _, err := client.Books.SoftwareSearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}
