package test

import (
	"go-rakutenapi/rakuten"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoraCourseSearch(t *testing.T) {
	params := &rakuten.GolfCourseParams{
		Keyword: "グリーン",
		Hits:    10,
	}
	resp, _, err := client.Gora.GolfCourseSearch(ctx, params)
	assert.Nil(t, err)
	assert.Equal(t, params.Hits, resp.Hits)
}

func TestGoraPlanSearch(t *testing.T) {
	params := &rakuten.GolfPlanParams{
		GolfCourseID: 90076,
		PlayDate:     "2020-04-09",
	}
	_, _, err := client.Gora.GolfPlanSearch(ctx, params)
	assert.Nil(t, err)
}

func TestGoraCourseDetailSearch(t *testing.T) {
	params := &rakuten.GolfCourseDetailParams{
		GolfCourseID: 90076,
	}
	_, _, err := client.Gora.GolfCourseDetailSearch(ctx, params)
	assert.Nil(t, err)
}
