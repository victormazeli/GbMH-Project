package service_category

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategory) Services(ctx context.Context, obj *prisma.ServiceCategory) ([]*prisma.Service, error) {
	subCg, err := r.Prisma.Services(&prisma.ServicesParams{
		Where: &prisma.ServiceWhereInput{
			Category: &prisma.ServiceCategoryWhereInput{
				ID: &obj.ID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, nil
	}

	var nodes []*prisma.Service
	for _, cg := range subCg {
		clone := cg
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
