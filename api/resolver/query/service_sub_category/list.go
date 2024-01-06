package service_sub_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategoryQuery) ServiceSubCategories(ctx context.Context) ([]*prisma.ServiceSubCategory, error) {

	companyId := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	cgs, err := r.Prisma.ServiceSubCategories(&prisma.ServiceSubCategoriesParams{
		Where: &prisma.ServiceSubCategoryWhereInput{
			Company: &prisma.CompanyWhereInput{
				ID: &companyId,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var nodes []*prisma.ServiceSubCategory

	for _, cg := range cgs {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
