package postgres
// main
// postgres

import (
	"context"
	// "errors"
	// "fmt"
    // "net/http"
	// "github.com/lib/pq"
	// "github.com/opentracing/o pentracing-go"
	"GO_APP/internal/model/entity"
)


// Insert User data
func (postgres *Postgres) InsertUserData(ctx context.Context) (*entity.User, error) {
	c := entity.User{}

	err := postgres.queries.insertUserData.SelectContext(ctx, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Insert User Address data
func (postgres *Postgres) InsertAddrData(ctx context.Context) (*entity.UserAddress, error) {
	c := entity.UserAddress{}

	err := postgres.queries.insertAddrData.SelectContext(ctx, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}



// //GetAllCategory get map of id with Category object from Memcache
// func (postgres *Postgres) InsertUserData(ctx context.Context) (*entity.User, error) {
// 	c := entity.User{}

// 	err := postgres.queries.insertUserData.SelectContext(ctx, &c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &c, nil
// }


// //FetchCategoryFromDB fetch the entire category tree from DB
// func (postgres *Postgres) FetchCategoryFromDB(ctx context.Context) ([]*entity.Category, error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "FetchCategoryFromDB")
// 		defer span.Finish()
// 	}

// 	c := []*entity.Category{}

// 	err := postgres.queries.getCategory.SelectContext(ctx, &c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// parse bit information
// 	for _, v := range c {
// 		bitInformation := postgres.ParseBitfield(ctx, v.BitInformation.String)
// 		if bitInformation.IsCrawlable {
// 			v.IsCrawlable = 1
// 		}
// 	}

// 	return c, nil
// }

// //ConstructCategory Construct Category object from DB
// func (postgres *Postgres) ConstructCategory() error {
// 	var ctx = context.Background()

// 	// wrap everything into a span context
// 	span := opentracing.GlobalTracer().StartSpan("ConstructCategory")
// 	defer span.Finish()
// 	ctx = opentracing.ContextWithSpan(ctx, span)

// 	tmpArrCategory := []*entity.Category{}
// 	tmpSortedCategory := []*entity.Category{}
// 	tmpCategoryByID := make(map[int]*entity.Category)
// 	tmpCategoryByIdentifier := make(map[string]*entity.Category)

// 	err := postgres.queries.getCategory.SelectContext(ctx, &tmpArrCategory)
// 	if err != nil {
// 		return err
// 	}

// 	// TODO: refactor: this logic and the Set process should not be present in the repository
// 	// it's causing getting fresh data rom DB quite difficult
// 	// copying allCategory won't work since it's a slice of pointers
// 	// current workaround is a partial duplicate of this function (FetchCategoryFromDB) to return raw data
// 	// after refactoring FetchCategoryFromDB() should be deleted and use this instead or vice-versa
// 	for _, v := range tmpArrCategory {
// 		// parse bit information using existing loop
// 		bitInformation := postgres.ParseBitfield(ctx, v.BitInformation.String)
// 		if bitInformation.IsCrawlable {
// 			v.IsCrawlable = 1
// 		}

// 		tmpCategoryByID[v.ID] = v
// 		tmpCategoryByIdentifier[v.Identifier] = v
// 	}

// 	for i := range tmpArrCategory {
// 		if (*tmpArrCategory[i]).Tree == 1 {
// 			tmpSortedCategory = append(tmpSortedCategory, tmpArrCategory[i])
// 		}
// 		if (*tmpArrCategory[i]).Parent != 0 && tmpCategoryByID[(*tmpArrCategory[i]).Parent] != nil && tmpCategoryByID[(*tmpArrCategory[i]).Parent].Tree < (*tmpArrCategory[i]).Tree {
// 			tmpCategoryByID[(*tmpArrCategory[i]).Parent].Children = append(tmpCategoryByID[(*tmpArrCategory[i]).Parent].Children, tmpArrCategory[i])
// 		}
// 	}

// 	postgres.SetAllCategory(ctx, tmpCategoryByID)
// 	postgres.SetAllCategoryByIdentifier(ctx, tmpCategoryByIdentifier)
// 	postgres.SetAllSortedCategory(ctx, tmpSortedCategory)
// 	postgres.SetAllCategoryByLevel(ctx, postgres.CategoryByLevel(ctx))

// 	return nil
// }

// //ConstructCategoryImage Construct Category Image object from DB
// func (postgres *Postgres) ConstructCategoryImage() (err error) {
// 	var ctx = context.Background()

// 	// wrap everything into a span context
// 	span := opentracing.GlobalTracer().StartSpan("ConstructCategoryImage")
// 	defer span.Finish()
// 	ctx = opentracing.ContextWithSpan(ctx, span)

// 	tmpCategoryImage := &entity.CategoryImageByType{}

// 	tmpCategoryImage, err = postgres.GetAllCategoryImageActiveByType(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	postgres.SetAllCategoryImageGray(ctx, tmpCategoryImage.CategoryImageGray)
// 	postgres.SetAllCategoryImageBanner(ctx, tmpCategoryImage.CategoryImageBanner)
// 	postgres.SetAllCategoryImageStaticHeader(ctx, tmpCategoryImage.CategoryImageStaticHeader)

// 	return
// }

// //GetAllCategoryImageActiveByType return map of category id in category image
// func (postgres *Postgres) GetAllCategoryImageActiveByType(ctx context.Context) (categoryImage *entity.CategoryImageByType, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetAllCategoryImageActiveByType")
// 		defer span.Finish()
// 	}

// 	tmpImage := []*entity.CategoryImage{}
// 	categoryImage = &entity.CategoryImageByType{}
// 	categoryImageGray := make(map[int]*entity.CategoryImage)
// 	categoryImageBanner := make(map[int]*entity.CategoryImage)
// 	categoryImageStaticHeader := make(map[int]*entity.CategoryImage)

// 	err = postgres.queries.getAllCategoryImageActive.SelectContext(ctx, &tmpImage)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, v := range tmpImage {
// 		if _, ok := postgres.allCategory[v.CategoryID]; ok {
// 			if v.Type == 1 { // used in detail for imageStaticHeader type Header
// 				categoryImageStaticHeader[v.CategoryID] = v
// 			} else if v.Type == 5 { // icon image url gray used for tree
// 				categoryImageGray[v.CategoryID] = v
// 			} else if v.Type == 6 { // icon banner used for tree
// 				categoryImageBanner[v.CategoryID] = v
// 			}
// 		}
// 	}
// 	categoryImage.CategoryImageGray = categoryImageGray
// 	categoryImage.CategoryImageBanner = categoryImageBanner
// 	categoryImage.CategoryImageStaticHeader = categoryImageStaticHeader

// 	return
// }

// //GetIconImageURLGray return map of category id in category image
// func (postgres *Postgres) GetIconImageURLGray(ctx context.Context, categoryID int) (imageGray *entity.CategoryImage, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetIconImageURLGray")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	imageGrays := postgres.allCategoryImageGray
// 	if _, ok := imageGrays[categoryID]; ok {
// 		imageGray = imageGrays[categoryID]
// 	} else {
// 		err = errors.New("not found")
// 	}

// 	return
// }

// //GetIconBanner return map of category id in category image
// func (postgres *Postgres) GetIconBanner(ctx context.Context, categoryID int) (imageBanner *entity.CategoryImage, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetIconBanner")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	imageBanners := postgres.allCategoryImageBanner
// 	if _, ok := imageBanners[categoryID]; ok {
// 		imageBanner = imageBanners[categoryID]
// 	} else {
// 		err = errors.New("not found")
// 	}

// 	return
// }

// //GetImageStaticHeader return map of category id in category image
// func (postgres *Postgres) GetImageStaticHeader(ctx context.Context, categoryID int) (headerImage *entity.CategoryImage) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetImageStaticHeader")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	imageStaticHeader := postgres.allCategoryImageStaticHeader
// 	if _, ok := imageStaticHeader[categoryID]; ok {
// 		headerImage = imageStaticHeader[categoryID]
// 	}

// 	return
// }

// //GetAllCategory get map of id with Category object from Memcache
// func (postgres *Postgres) GetAllCategory(ctx context.Context) map[int]*entity.Category {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetAllCategory")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.allCategory
// }

// //SetAllCategory set map of id with Category object from Memcache
// func (postgres *Postgres) SetAllCategory(ctx context.Context, allCategory map[int]*entity.Category) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllCategory")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allCategory = allCategory
// 	return
// }

// //GetCategoryMapByIDSortedByPopular get map of id with Category object from Memcache sorted by popularity
// func (postgres *Postgres) GetCategoryMapByIDSortedByPopular(ctx context.Context) map[int]*entity.Category {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetCategoryMapByIDSortedByPopular")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.categoryMapByIDSortedByPopular
// }

// //SetCategoryMapByIDSortedByPopular set map of id with Category object from Memcache sorted by popularity
// func (postgres *Postgres) SetCategoryMapByIDSortedByPopular(ctx context.Context, sortedCategory map[int]*entity.Category) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetCategoryMapByIDSortedByPopular")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.categoryMapByIDSortedByPopular = sortedCategory
// 	return
// }

// //GetAllCategoryByIdentifier get map of identifier with Category object from Memcache
// func (postgres *Postgres) GetAllCategoryByIdentifier(ctx context.Context) map[string]*entity.Category {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetAllCategoryByIdentifier")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.allCategoryByIdentifier
// }

// //SetAllCategoryByIdentifier set map of identifier with Category object from Memcache
// func (postgres *Postgres) SetAllCategoryByIdentifier(ctx context.Context, allCategoryByIdentifier map[string]*entity.Category) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllCategoryByIdentifier")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allCategoryByIdentifier = allCategoryByIdentifier
// 	return
// }

// //GetCategoryMapByIdentifierSortedByPopular get map of identifier with Category object from Memcache
// func (postgres *Postgres) GetCategoryMapByIdentifierSortedByPopular(ctx context.Context) map[string]*entity.Category {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetCategoryMapByIdentifierSortedByPopular")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.categoryMapByIdentifierSortedByPopular
// }

// //SetCategoryMapByIdentifierSortedByPopular set map of identifier with Category object from Memcache
// func (postgres *Postgres) SetCategoryMapByIdentifierSortedByPopular(ctx context.Context, sortedCategory map[string]*entity.Category) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetCategoryMapByIdentifierSortedByPopular")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.categoryMapByIdentifierSortedByPopular = sortedCategory
// 	return
// }

// //SetAllCategoryImageGray set map of identifier with Category object from Memcache
// func (postgres *Postgres) SetAllCategoryImageGray(ctx context.Context, allCategoryImageGray map[int]*entity.CategoryImage) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllCategoryImageGray")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allCategoryImageGray = allCategoryImageGray
// 	return
// }

// //SetAllCategoryImageStaticHeader set map of identifier with Category object from Memcache
// func (postgres *Postgres) SetAllCategoryImageStaticHeader(ctx context.Context, allCategoryImageStaticHeader map[int]*entity.CategoryImage) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllCategoryImageStaticHeader")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allCategoryImageStaticHeader = allCategoryImageStaticHeader
// 	return
// }

// //SetAllCategoryImageBanner set map of identifier with Category object from Memcache
// func (postgres *Postgres) SetAllCategoryImageBanner(ctx context.Context, allCategoryImageGray map[int]*entity.CategoryImage) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllCategoryImageBanner")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allCategoryImageBanner = allCategoryImageGray
// 	return
// }

// //GetAllSortedCategory get sorted list Category object from Memcache
// func (postgres *Postgres) GetAllSortedCategory(ctx context.Context) []*entity.Category {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetAllSortedCategory")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.allSortedCategory
// }

// //SetAllSortedCategory set sorted list Category object from Memcache
// func (postgres *Postgres) SetAllSortedCategory(ctx context.Context, allSortedCategory []*entity.Category) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllSortedCategory")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allSortedCategory = allSortedCategory
// 	return
// }

// // GetAllSortedByPopularCategory get sorted list Category object from Memcache  with additional sort by popular categories from bigquery
// // ALWAYS check for nil first, the source is an external calls
// func (postgres *Postgres) GetAllSortedByPopularCategory(ctx context.Context) []*entity.Category {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetAllSortedByPopularCategory")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.allSortedByPopularCategory
// }

// // SetAllSortedByPopularCategory set sorted list Category object from Memcache with additional sort by popular categories from bigquery
// func (postgres *Postgres) SetAllSortedByPopularCategory(ctx context.Context, allSortedByPopularCategory []*entity.Category) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllSortedByPopularCategory")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allSortedByPopularCategory = allSortedByPopularCategory
// 	return
// }

// //GetAllCategoryByLevel get sorted list by level Category object from Memcache
// func (postgres *Postgres) GetAllCategoryByLevel(ctx context.Context) *entity.CategoryByLevel {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetAllCategoryByLevel")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()

// 	return postgres.allCategoryByLevel
// }

// //SetAllCategoryByLevel set sorted list by level Category object from Memcache
// func (postgres *Postgres) SetAllCategoryByLevel(ctx context.Context, allCategoryByLevel *entity.CategoryByLevel) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetAllCategoryByLevel")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()

// 	postgres.allCategoryByLevel = allCategoryByLevel
// 	return
// }

// //GetCategoryBannerByID get category banner from DB
// func (postgres *Postgres) GetCategoryBannerByID(ctx context.Context, ID int) (result []*entity.CategoryBanner, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetCategoryBannerByID")
// 		defer span.Finish()
// 	}

// 	err = postgres.queries.getCategoryBanner.SelectContext(ctx, &result, ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// //GetCategoryImageByIDs get Category Image from DB
// func (postgres *Postgres) GetCategoryImageByIDs(ctx context.Context, ID []int, Type []int) (result []*entity.CategoryImage, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetCategoryImageByIDs")
// 		defer span.Finish()
// 	}

// 	err = postgres.queries.getCategoryImageByIDsAndTypes.SelectContext(ctx, &result, pq.Array(ID), pq.Array(Type))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// //GetImageHexByImageID returns HEX Color from ImageID
// func (postgres *Postgres) GetImageHexByImageID(ctx context.Context, imageID int) (argsHex string, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetImageHexByImageID")
// 		defer span.Finish()
// 	}

// 	err = postgres.queries.getImageHexColor.GetContext(ctx, &argsHex, imageID)
// 	return
// }

// //GetCategoryParentAndSibling get parent and sibling from DB
// func (postgres *Postgres) GetCategoryParentAndSibling(ctx context.Context, ID int) (result []*entity.Category, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetCategoryParentAndSibling")
// 		defer span.Finish()
// 	}

// 	err = postgres.queries.getChildByParentCategoryID.SelectContext(ctx, &result, ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// //CategoryByLevel Construct Category By level from Memcache
// func (postgres *Postgres) CategoryByLevel(ctx context.Context) *entity.CategoryByLevel {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "CategoryByLevel")
// 		defer span.Finish()
// 	}

// 	var (
// 		result              *entity.CategoryByLevel
// 		temp1, temp2, temp3 []*entity.Category
// 	)
// 	for _, e := range postgres.GetAllCategory(ctx) {
// 		if e.Tree == 1 {
// 			temp1 = append(temp1, e)
// 		} else if e.Tree == 2 {
// 			temp2 = append(temp2, e)
// 		} else {
// 			temp3 = append(temp3, e)
// 		}
// 	}

// 	result = &entity.CategoryByLevel{
// 		Level1: temp1,
// 		Level2: temp2,
// 		Level3: temp3,
// 	}

// 	return result
// }

// //GetRedirectionURLByID get the list of possible redirection combination for the specified category ID
// func (postgres *Postgres) GetRedirectionURLByID(ctx context.Context, ID int) (result []*entity.RedirectionURL, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetRedirectionURLByID")
// 		defer span.Finish()
// 	}

// 	err = postgres.queries.getRedirectionURLByID.SelectContext(ctx, &result, ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// //GetTrendingCategories get the list of trending categories from database
// func (postgres *Postgres) GetTrendingCategories(ctx context.Context) (trendingCategory []*entity.Category, err error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetTrendingCategories")
// 		defer span.Finish()
// 	}

// 	err = postgres.queries.getTrendingCategories.SelectContext(ctx, &trendingCategory)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return
// }

// // GetCategoryRedirectionMap returns map of old and new identifier
// func (postgres *Postgres) GetCategoryRedirectionMap() (err error) {
// 	var ctx = context.Background()

// 	// wrap everything into a span context
// 	span := opentracing.GlobalTracer().StartSpan("GetCategoryRedirectionMap")
// 	defer span.Finish()
// 	ctx = opentracing.ContextWithSpan(ctx, span)

// 	var dest []entity.RedirectionMapTable
// 	var redirectionMap = make(map[string]string)

// 	err = postgres.queries.getRedirectionMap.SelectContext(ctx, &dest)
// 	if err != nil {
// 		err = fmt.Errorf("[GetCategoryRedirectionMap] %v", err.Error())
// 		return err
// 	}
// 	for _, val := range dest {
// 		redirectionMap[val.OldIdentifier] = val.NewIdentifier
// 	}
// 	postgres.SetRedirectionMap(ctx, redirectionMap)
// 	return nil
// }

// //GetRedirectionMap return RedirectionMap
// func (postgres *Postgres) GetRedirectionMap(ctx context.Context) map[string]string {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetRedirectionMap")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.RLock()
// 	defer postgres.mtxDBMap.RUnlock()
// 	return postgres.RedirectionMap
// }

// //SetRedirectionMap set Redirection Map
// func (postgres *Postgres) SetRedirectionMap(ctx context.Context, redirectionMap map[string]string) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "SetRedirectionMap")
// 		defer span.Finish()
// 	}

// 	postgres.mtxDBMap.Lock()
// 	defer postgres.mtxDBMap.Unlock()
// 	postgres.RedirectionMap = redirectionMap
// 	return
// }

// // GetCategoryTaxData return CategoryTax data from DB specified by category ID
// func (postgres *Postgres) GetCategoryTaxData(ctx context.Context, categoryID int) ([]*entity.CategoryTax, error) {
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span, ctx = opentracing.StartSpanFromContext(ctx, "GetCategoryTaxData")
// 		defer span.Finish()
// 	}

// 	var result []*entity.CategoryTax

// 	err := postgres.queries.getCategoryTaxData.SelectContext(ctx, &result, categoryID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

