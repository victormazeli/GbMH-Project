package product_sub_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ProductSubCategory struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ProductSubCategory {
	return &ProductSubCategory{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
