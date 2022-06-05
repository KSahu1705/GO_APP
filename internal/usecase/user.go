package usecase

import (
	"context"

	"GO_APP/internal/model/entity"
)

//UserCase is an interface
type UserCase interface {
	PutUserData(ctx context.Context) (*entity.User, error)
	// GetCategoryList(ctx context.Context, params *entity.CategoryParams) ([]*entity.CategoryItem, error)
	// GetCategoryTree(ctx context.Context, params *entity.CategoryParams) ([]*entity.CategoryItem, error)
	// GetRelatedCategory(ctx context.Context, categoryID int, childCategory []*entity.Category, parentOfCategoryID int, deviceType string) ([]entity.RelatedCategory, error)
	// CollectCategoryListDefault(ctx context.Context, id int, tree int, safeSearch bool, categoryIDExists bool) (temp1, temp2, temp3 []*entity.CategoryItem)
	// CollectCategoryListBreadcrumb(ctx context.Context, id int, tree int, safeSearch bool, categoryIDExists bool) (temp1, temp2, temp3 []*entity.CategoryItem)
	// CollectCategoryList(ctx context.Context, id int, safeSearch bool, level string, categoryIDExists bool) ([]*entity.CategoryItem, error)
	// AssignCategoryResponse(ctx context.Context, e *entity.Category) *entity.CategoryItem
	// AssignCategoryResponseWithChild(ctx context.Context, e *entity.Category, params entity.ParamsAssignCategoryResponseWithChild) *entity.CategoryItem
	// GetApplinksByID(ctx context.Context, e *entity.Category) (applinks string)
	// GetCategoryIDFromIdentifier(ctx context.Context, identifier string) (CategoryID int)
	// GetCategoryTreeLite(ctx context.Context, sender string, show string) ([]*entity.CategoryLite, error)
	// OverwriteDiscoURL(ctx context.Context, categoryTree []*entity.CategoryItem, inputDiscoItems string, device string) (categories []*entity.CategoryItem)
	// ParseParams(ctx context.Context, params *entity.DetailCategoryParams) (err error)
	// ValidateIdentifierFromRedirectionMap(ctx context.Context, params *entity.DetailCategoryParams) (err error)
	// ValidateIDFromCategory(ctx context.Context, params *entity.DetailCategoryParams) (err error)
	// GetBreadcrumb(ctx context.Context, id int) (breadcrumb []entity.Breadcrumb, err error)
	// RefreshPopularCategories() error
	// GetModularData(ctx context.Context, params *entity.DetailCategoryParams, child []*pb.DetailCategoryChild, tree int) ([]*pb.Components, error)
	// GetBitfieldDefinition(ctx context.Context, source string) (bitfieldDefinition map[string]int64, err error)
	// BuildTabsComponent(params *entity.DetailCategoryParams, children []*pb.DetailCategoryChild, dataComponent []*pb.Components) (tabComponent *pb.Components, err error)
}