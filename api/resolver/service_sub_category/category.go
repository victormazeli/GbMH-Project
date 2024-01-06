package service_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategory) Category(ctx context.Context, obj *prisma.ServiceSubCategory) (*prisma.ServiceCategory, error) {
	return r.Prisma.ServiceSubCategory(prisma.ServiceSubCategoryWhereUniqueInput{
		ID: &obj.ID,
	}).Category().Exec(ctx)
}
