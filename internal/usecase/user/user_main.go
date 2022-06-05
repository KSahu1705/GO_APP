package category

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	// "sort"

	"GO_APP/internal/model/entity"
)

// RefreshPopularCategories fetches popular categories from google bigquery via redis
func (i *Implementation) PutUserData(ctx context.Context) (*entity.User, error) {
	fmt.Println("line22222222@@@@@@@")
	userData, err := i.UserRepository.InsertUserData(ctx)
	if err != nil {
		fmt.Println("[RefreshPopularCategories][RedisGET] ", err)
		return userData, err
	}

	return userData, nil
}