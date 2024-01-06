package product_sub_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ProductSubCategoryQuery struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ProductSubCategoryQuery {
	return &ProductSubCategoryQuery{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
