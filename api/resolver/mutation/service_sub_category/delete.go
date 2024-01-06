package service_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategoryMutation) DeleteServiceSubCategory(ctx context.Context, input gqlgen.DeleteServiceSubCategoryInput) (*gqlgen.DeleteServiceSubCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator && user.Type != prisma.UserTypeManager {
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	subcg, err := r.Prisma.DeleteServiceSubCategory(prisma.ServiceSubCategoryWhereUniqueInput{
		ID: &input.ID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.DeleteServiceSubCategoryPayload{
		SubCategory: subcg,
	}, nil
}
