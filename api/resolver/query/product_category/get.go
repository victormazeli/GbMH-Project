package product_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductCategoryQuery) ProductCategory(ctx context.Context, id string) (*prisma.ProductCategory, error) {
	return r.Prisma.ProductCategory(prisma.ProductCategoryWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)
}
