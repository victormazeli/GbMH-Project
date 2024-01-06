package product_category

import "github.com/steebchen/keskin-api/prisma"

type ProductCategoryMutation struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *ProductCategoryMutation {
	return &ProductCategoryMutation{
		Prisma: client,
	}
}
