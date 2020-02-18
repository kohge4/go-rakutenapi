# go-rakutenapi

- https://webservice.rakuten.co.jp/document/
- 楽天ウェブサービスAPIのSDK
- 参考: https://github.com/google/go-github

## クイックスタート

```go
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
		Keyword: "マイケルジャクソン",
		Hits:    6,
	}

	// Search Items from Rakuten Ichiba API
	ichiba, _, err := client.Ichiba.Search(ctx, sOptions)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ichiba)
}

```
### 認証周り等の詳しい使い方は, example/auth_app のコードが参考になるかと思います

## 各エンドポイントとの対応

### 楽天市場系API

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|----------------|
|GET /IchibaItem/Search/20170706 | IchibaItemSearchParams | Ichiba.Search| アプリID認可方式 |
|GET /IchibaGenre/Search/20140222 | IchibaGenreSearchParams | Ichiba.GenreSearch| アプリID認可方式 |
|GET /IchibaTag/Search/20140222 | IchibatagSearchParams | Ichiba.TagSearch| アプリID認可方式 |
|GET /IchibaItem/Ranking/20170628 | IchibaRankingParams | Ichiba.Ranking| アプリID認可方式 |
|GET /Product/Search/20170426 | IchibaPriductSearchParams | Ichiba.ProductSearch| アプリID認可方式 |

### 楽天ブックス系API 

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|----------------|
|GET /BooksTotal/Search/20170404 | BooksTotalSearchParams | Books.TotalSearch| アプリID認可方式 |
|GET /BooksBook/Search/20170404 | BooksBookSearchParams | Books.BookSearch| アプリID認可方式 |
|GET /BooksCD/Search/20170404 | BooksCDSearchParams | Books.CDSearch| アプリID認可方式 |
|GET /BooksDVD/Search/20170404 | BooksDvdBluraySearchParams | Books.DvdBluraySearch| アプリID認可方式 |
|GET /BooksForeignBook/Search/20170404 | BooksForeignBookSearchParams | Books.ForeignBookSearch| アプリID認可方式 |
|GET /BooksMagazine/Search/201704044 | BooksMagazineSearchParams | Books.MagazineSearch| アプリID認可方式 |
|GET /BooksGame/Search/20170404 | BooksGameSearchParams | Books.GameSearch| アプリID認可方式 |
|GET /BooksSoftware/Search/20170404 | BooksSoftwareSearchParams | Books.SoftwareSearch| アプリID認可方式 |
|GET /BooksGenre/Search/20121128 | BooksGenreSearchParams | Books.GenreSearch| アプリID認可方式 |

### 楽天トラベル系API 

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|----------------|
|GET /Travel/SimpleHotelSearch/20170426 | TravelSimpleHotelSearchParams | Travel.SimpleHotelSearch| アプリID認可方式 |
|GET /Travel/HotelDetailSearch/20170426 | TravelHotelDetailSearchParams | Travel.HotelDetailSearch| アプリID認可方式 |
|GET /Travel/VacantHotelSearch/20170426 | VacantHotelSearchParams | Travel.VacantHotelSearch| アプリID認可方式 |
|GET /Travel/GetAreaClass/20131024 | TravelGetAreaParams | Travel.GetArea| アプリID認可方式 |
|GET /Travel/KeywordHotelSearch/20170426 | TravelKeywordHotelSearchParams | Travel.KeywordHotelSearch| アプリID認可方式 |
|GET /Travel/GetHotelChainList/20131024 | TravelHotelChainParams | Travel.HotelChainSearch| アプリID認可方式 |
|GET /Travel/HotelRanking/20170426 | TravelHotelRankingParams | Travel.HotelRanking| アプリID認可方式 |

### 楽天ブックマーク系API 

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|---------------|
|GET /FavoriteBookmark/List/20170426 | FavoGetParams | Favorite.Get| OAuth2.0 |
|GET /FavoriteBookmark/Delete/20120627 | FavoAddParams | Favorite.Add| 0Auth2.0 |
|GET /FavoriteBookmark/List/20170426 | FavoDeleteParams | Favorite.Delete| OAuth2.0 |

### 楽天レシピ系API

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|----------------|
|GET /Recipe/CategoryList/20170426 | RecipeCategoryParams | Recipe.Category| アプリID認可方式 |
|GET /Recipe/CategoryRanking/20170426 | RecipeRankingParams | Recipe.Ranking| アプリID認可方式 |

### 楽天Kobo系API

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|----------------|
|GET /Kobo/EbookSearch/20170426 | KoboEbooksParams | Kobo.EbooksSearch| アプリID認可方式 |
|GET /Kobo/GenreSearch/20131010 | KoboGenreSearchParams | Kobo.GenreSearch| アプリID認可方式 |

### 楽天Gora系API

|Endpoint | Parameter Struct | Method | Other |
|---------|------------------|--------|----------------|
|GET /Gora/GoraGolfCourseSearch/20170623 | GolfCourseParams | Gora.GolfCourseSearch| アプリID認可方式 |
|GET /Gora/GoraGolfCourseDetail/20170623 | GolfCourseDetailParams | Gora.GolfCourseDetailSearch| アプリID認可方式 |
|GET /Gora/GoraPlanSearch/20170623 | GolfPlanParams | Gora.GolfPlanSearch| アプリID認可方式 |



