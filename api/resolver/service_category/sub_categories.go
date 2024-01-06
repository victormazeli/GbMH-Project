package service_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategory) SubCategories(ctx context.Context, obj *prisma.ServiceCategory) ([]*prisma.ServiceSubCategory, error) {
	subCg, err := r.Prisma.ServiceSubCategories(&prisma.ServiceSubCategoriesParams{
		Where: &prisma.ServiceSubCategoryWhereInput{
			Category: &prisma.ServiceCategoryWhereInput{
				ID: &obj.ID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, nil
	}

	nodes := []*prisma.ServiceSubCategory{}
	for _, cg := range subCg {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
