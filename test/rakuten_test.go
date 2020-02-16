package test

import (
	"github.com/kohge4/go-rakutenapi/rakuten"
)

const (
	applicationID     = "1058291654304154113"
	applicationSecret = "1faf7b39ed12271e735d7f67b5bf9a8aa5758afb"
	affiliateID       = "1a11cd3a.d53d02bd.1a11cd3b.54ce7249"
)

func NewTestClient() *rakuten.Client {
	tp := rakuten.Transport{}
	//client := rakuten.NewClient(tp.Client(), os.Getenv("applicationID"), os.Getenv("affiliateID"))
	client := rakuten.NewClient(tp.Client(), applicationID, affiliateID)
	return client
}
