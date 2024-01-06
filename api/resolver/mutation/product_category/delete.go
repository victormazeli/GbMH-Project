package product_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductCategoryMutation) DeleteProductCategory(ctx context.Context, input gqlgen.DeleteProductCategoryInput) (*gqlgen.DeleteProductCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator && user.Type != prisma.UserTypeManager {
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	cg, err := r.Prisma.DeleteProductCategory(prisma.ProductCategoryWhereUniqueInput{
		ID: &input.ID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.DeleteProductCategoryPayload{
		Category: cg,
	}, nil
}
