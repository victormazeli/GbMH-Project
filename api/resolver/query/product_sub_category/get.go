package product_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductSubCategoryQuery) ProductSubCategory(ctx context.Context, id string) (*prisma.ProductSubCategory, error) {
	subCg, err := r.Prisma.ProductSubCategory(prisma.ProductSubCategoryWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return subCg, nil
}
