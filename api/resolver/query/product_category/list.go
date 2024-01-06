package product_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductCategoryQuery) ProductCategories(ctx context.Context) ([]*prisma.ProductCategory, error) {

	companyId := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)
	cgs, err := r.Prisma.ProductCategories(&prisma.ProductCategoriesParams{
		Where: &prisma.ProductCategoryWhereInput{
			Company: &prisma.CompanyWhereInput{
				ID: &companyId,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var nodes []*prisma.ProductCategory

	for _, cg := range cgs {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
