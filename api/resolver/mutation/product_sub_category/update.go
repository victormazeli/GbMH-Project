package product_sub_category

import (
	"context"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductSubCategoryMutation) UpdateProductSubCategory(ctx context.Context, input gqlgen.UpdateProductSubCategoryInput) (*gqlgen.UpdateProductSubCategoryPayload, error) {
	user, err := sessctx.User(ctx)

	if err != nil {
		return nil, err
	}

	if user.Type != prisma.UserTypeAdministrator  && user.Type != prisma.UserTypeManager{
		return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	}

	subCg, err := r.Prisma.UpdateProductSubCategory(prisma.ProductSubCategoryUpdateParams{
		Where: prisma.ProductSubCategoryWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.ProductSubCategoryUpdateInput{
			Name: input.Patch.Name,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.UpdateProductSubCategoryPayload{
		SubCategory: subCg,
	}, nil
}
