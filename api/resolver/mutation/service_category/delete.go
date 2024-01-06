package service_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategoryMutation) DeleteServiceCategory(ctx context.Context, input gqlgen.DeleteServiceCategoryInput) (*gqlgen.DeleteServiceCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator && user.Type != prisma.UserTypeManager {
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	cg, err := r.Prisma.DeleteServiceCategory(prisma.ServiceCategoryWhereUniqueInput{
		ID: &input.ID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.DeleteServiceCategoryPayload{
		Category: cg,
	}, nil
}
