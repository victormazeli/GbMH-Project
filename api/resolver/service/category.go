package service

import (
	"context"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Service) Category(ctx context.Context, obj *prisma.Service) (*prisma.ServiceCategory, error) {
	if obj.Deleted {
		return nil, nil
	}

	category, err := r.Prisma.Service(prisma.ServiceWhereUniqueInput{
		ID: &obj.ID,
	}).Category().Exec(ctx)

	if err != nil {
		return nil, nil
	}

	return category, nil

}
