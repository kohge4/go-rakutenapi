package quickstart

import (
	"context"
	"fmt"

	"github.com/kohge4/go-rakutenapi/rakuten"
)

const (
	applicationID     = "ApplicationID"
	applicationSecret = "ApplicationSecret"
	affiliateID       = "AffiliateID"
)

func QuickStart() {
	ctx := context.Background()
	tp := rakuten.Transport{}

	// Rakuten API client
	client := rakuten.NewClient(tp.Client(), applicationID, affiliateID)

	// QueryParameter for Search argument
	sOptions := &rakuten.IchibaItemSearchParams{
		Keyword: "cero",
		Hits:    6,
	}

	// Search Items from Rakuten Ichiba API
	ichiba, _, err := client.Ichiba.Search(ctx, sOptions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ichiba)
}
