package service_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategoryQuery) ServiceCategory(ctx context.Context, id string) (*prisma.ServiceCategory, error) {
	return r.Prisma.ServiceCategory(prisma.ServiceCategoryWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)
}
