package service_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategoryQuery) ServiceCategories(ctx context.Context) ([]*prisma.ServiceCategory, error) {

	companyId := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	cgs, err := r.Prisma.ServiceCategories(&prisma.ServiceCategoriesParams{
		Where: &prisma.ServiceCategoryWhereInput{
			Company: &prisma.CompanyWhereInput{
				ID: &companyId,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	var nodes []*prisma.ServiceCategory

	for _, cg := range cgs {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
