package product_sub_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductSubCategoryQuery) ProductSubCategories(ctx context.Context) ([]*prisma.ProductSubCategory, error) {
	companyId := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	cgs, err := r.Prisma.ProductSubCategories(&prisma.ProductSubCategoriesParams{
		Where: &prisma.ProductSubCategoryWhereInput{
			Company: &prisma.CompanyWhereInput{
				ID: &companyId,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var nodes []*prisma.ProductSubCategory

	for _, cg := range cgs {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
