package repository

import (
	"context"

	"GO_APP/internal/model/entity"
)

//go:generate mockgen -destination=../../testfile/repository/category.go -package=mock_repository . Category

//Category defines the methods available for Category repository
type User interface {
	InsertUserData(ctx context.Context) (*entity.User, error)
	InsertAddrData(ctx context.Context) (*entity.UserAddress, error)
	// ConstructCategory() error
	// ParseBitfield(ctx context.Context, bitString string) *entity.BitInformation
	// ConstructCategoryImage() (err error)
	// FetchCategoryFromDB(context.Context) ([]*entity.Category, error)
	// GetAllCategory(context.Context) map[int]*entity.Category
	// SetAllCategory(ctx context.Context, allCategory map[int]*entity.Category)
	// GetCategoryMapByIDSortedByPopular(context.Context) map[int]*entity.Category
	// SetCategoryMapByIDSortedByPopular(ctx context.Context, allCategory map[int]*entity.Category)
	// SetAllCategoryImageGray(context.Context, map[int]*entity.CategoryImage)
	// SetAllCategoryImageBanner(context.Context, map[int]*entity.CategoryImage)
	// GetAllCategoryByIdentifier(context.Context) map[string]*entity.Category
	// SetAllCategoryByIdentifier(ctx context.Context, allCategoryByIdentifier map[string]*entity.Category)
	// GetCategoryMapByIdentifierSortedByPopular(context.Context) map[string]*entity.Category
	// SetCategoryMapByIdentifierSortedByPopular(ctx context.Context, allCategory map[string]*entity.Category)
	// GetAllSortedCategory(context.Context) []*entity.Category
	// SetAllSortedCategory(ctx context.Context, allSortedCategory []*entity.Category)
	// GetAllSortedByPopularCategory(context.Context) []*entity.Category
	// SetAllSortedByPopularCategory(ctx context.Context, allSortedByPopularCategory []*entity.Category)
	// GetAllCategoryByLevel(context.Context) *entity.CategoryByLevel
	// SetAllCategoryByLevel(ctx context.Context, allCategoryByLevel *entity.CategoryByLevel)
	// GetCategoryBannerByID(ctx context.Context, ID int) ([]*entity.CategoryBanner, error)
	// GetCategoryImageByIDs(ctx context.Context, IDs []int, Type []int) ([]*entity.CategoryImage, error)
	// GetImageHexByImageID(ctx context.Context, imageID int) (string, error)
	// GetCategoryParentAndSibling(ctx context.Context, ID int) ([]*entity.Category, error)
	// GetRedirectionURLByID(ctx context.Context, ID int) ([]*entity.RedirectionURL, error)
	// GetTrendingCategories(context.Context) ([]*entity.Category, error)
	// CategoryByLevel(context.Context) *entity.CategoryByLevel
	// GetIconImageURLGray(ctx context.Context, categoryID int) (imageGray *entity.CategoryImage, err error)
	// GetIconBanner(ctx context.Context, categoryID int) (imageBanner *entity.CategoryImage, err error)
	// GetImageStaticHeader(ctx context.Context, categoryID int) (headerImage *entity.CategoryImage)
	// GetCategoryRedirectionMap() (err error)
	// GetRedirectionMap(ctx context.Context) map[string]string
	// SetRedirectionMap(ctx context.Context, redirectionMap map[string]string)
	// GetBitfieldDefinition(ctx context.Context) (bitDefinition map[string]int64)
}