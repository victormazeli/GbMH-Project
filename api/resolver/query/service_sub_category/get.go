package service_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategoryQuery) ServiceSubCategory(ctx context.Context, id string) (*prisma.ServiceSubCategory, error) {
	subCg, err := r.Prisma.ServiceSubCategory(prisma.ServiceSubCategoryWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return subCg, nil
}
