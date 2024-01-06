package product_sub_category

import (
	"github.com/steebchen/keskin-api/prisma"
)

type ProductSubCategoryMutation struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ProductSubCategoryMutation {
	return &ProductSubCategoryMutation{
		Prisma: client,
	}
}
