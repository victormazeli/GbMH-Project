package service_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategoryMutation) UpdateServiceCategory(ctx context.Context, input gqlgen.UpdateServiceCategoryInput) (*gqlgen.UpdateServiceCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator && user.Type != prisma.UserTypeManager{
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	cg, err := r.Prisma.UpdateServiceCategory(prisma.ServiceCategoryUpdateParams{
		Where: prisma.ServiceCategoryWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.ServiceCategoryUpdateInput{
			Name: input.Patch.Name,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.UpdateServiceCategoryPayload{
		Category: cg,
	}, nil
}
