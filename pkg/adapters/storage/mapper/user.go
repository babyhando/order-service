package mapper

import (
	"order-service/internal/user/domain"
	"order-service/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func UserDomain2Storage(userDomain domain.User) *types.User {
	return &types.User{
		Model: gorm.Model{
			ID:        uint(userDomain.ID),
			CreatedAt: userDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(userDomain.DeletedAt)),
		},
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		Phone:     string(userDomain.Phone),
	}
}
