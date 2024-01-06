package service_sub_category

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategory) Services(ctx context.Context, obj *prisma.ServiceSubCategory) ([]*prisma.Service, error) {

	services, err := r.Prisma.Services(&prisma.ServicesParams{
		Where: &prisma.ServiceWhereInput{
			SubCategory: &prisma.ServiceSubCategoryWhereInput{
				ID: &obj.ID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, nil
	}

	var nodes []*prisma.Service
	for _, service := range services {
		clone := service
		nodes = append(nodes, &clone)
	}

	return nodes, nil
}
