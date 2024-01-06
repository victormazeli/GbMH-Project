package product_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ProductCategoryQuery struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ProductCategoryQuery {
	return &ProductCategoryQuery{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
