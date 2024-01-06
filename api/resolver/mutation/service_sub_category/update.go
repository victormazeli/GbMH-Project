package service_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategoryMutation) UpdateServiceSubCategory(ctx context.Context, input gqlgen.UpdateServiceSubCategoryInput) (*gqlgen.UpdateServiceSubCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator && user.Type != prisma.UserTypeManager{
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	subCg, err := r.Prisma.UpdateServiceSubCategory(prisma.ServiceSubCategoryUpdateParams{
		Where: prisma.ServiceSubCategoryWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.ServiceSubCategoryUpdateInput{
			Name: input.Patch.Name,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.UpdateServiceSubCategoryPayload{
		SubCategory: subCg,
	}, nil
}
