package test

import (
	"testing"

	"github.com/kohge4/go-rakutenapi/rakuten"

	"github.com/stretchr/testify/assert"
)

func TestTravelGetAreaClass(t *testing.T) {
	params := &rakuten.TravelGetAreaParams{}
	_, _, err := client.Travel.GetArea(ctx, params)
	assert.Nil(t, err)
}

func TestTravelHotelChain(t *testing.T) {
	params := &rakuten.TravelHotelChainParams{}
	_, _, err := client.Travel.HotelChain(ctx, params)
	assert.Nil(t, err)
}

func TestTravelHotelRanking(t *testing.T) {
	params := &rakuten.TravelHotelRankingParams{
		Genre:   "onsen",
		Carrier: 0,
	}
	_, _, err := client.Travel.HotelRanking(ctx, params)
	assert.Nil(t, err)
}

func TestTravelKeywordHotel(t *testing.T) {
	params := &rakuten.TravelKeywordHotelSearchParams{
		Keyword: "北海道",
		Hits:    5,
	}
	_, _, err := client.Travel.KeywordHotelSearch(ctx, params)
	assert.Nil(t, err)
}

func TestTravelSimpleSearch(t *testing.T) {
	params := &rakuten.TravelSimpleHotelSearchParams{
		LargeClassCode:  "japan",
		MiddleClassCode: "akita",
		SmallClassCode:  "tazawa",
	}
	_, _, err := client.Travel.SimpleHotelSearch(ctx, params)
	assert.Nil(t, err)
}

func TestTraveHotelDetailSearch(t *testing.T) {
	params := &rakuten.TravelHotelDetailSearchParams{
		HotelNo: 4624,
	}
	_, _, err := client.Travel.HotelDetailSearch(ctx, params)
	assert.Nil(t, err)
}

func TestTravelVacantHotelSearch(t *testing.T) {
	params := &rakuten.TravelVacantHotelSearchParams{
		LargeClassCode:  "japan",
		MiddleClassCode: "akita",
		SmallClassCode:  "tazawa",
		CheckDate:       "2020-03-15",
	}
	_, _, err := client.Travel.VacantHotelSearch(ctx, params)
	assert.Nil(t, err)
}
