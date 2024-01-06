package product_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductCategoryMutation) UpdateProductCategory(ctx context.Context, input gqlgen.UpdateProductCategoryInput) (*gqlgen.UpdateProductCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator && user.Type != prisma.UserTypeManager {
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	cg, err := r.Prisma.UpdateProductCategory(prisma.ProductCategoryUpdateParams{
		Where: prisma.ProductCategoryWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.ProductCategoryUpdateInput{
			Name: input.Patch.Name,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.UpdateProductCategoryPayload{
		Category: cg,
	}, nil
}
