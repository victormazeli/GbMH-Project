package product_category

import (
	"github.com/google/wire"
	"github.com/steebchen/keskin-api/prisma"
)

type ProductCategory struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ProductCategory {
	return &ProductCategory{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
