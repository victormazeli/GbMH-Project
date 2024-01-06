package service

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Service) SubCategory(ctx context.Context, obj *prisma.Service) (*prisma.ServiceSubCategory, error) {
	if obj.Deleted {
		return nil, nil
	}

	subcategory, err := r.Prisma.Service(prisma.ServiceWhereUniqueInput{
		ID: &obj.ID,
	}).SubCategory().Exec(ctx)

	if err != nil {
		return nil, nil
	}

	return subcategory, nil

}
