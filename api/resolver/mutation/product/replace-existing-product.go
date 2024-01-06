package product

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductMutation) ReplaceExistingProduct(
	ctx context.Context,
	input gqlgen.ReplaceExistingProductInput,
	language *string,
) (*gqlgen.ReplaceExistingProductsPayload, error) {

	// Validate input.Services to ensure it's not empty or contains invalid IDs

	if len(input.Products) == 0 {
		return nil, gqlerrors.NewValidationError("Products can not be empty", "EmptyData")
	}

	for _, productID := range input.Products {
		_, err := r.Prisma.UpdateProduct(prisma.ProductUpdateParams{
			Where: prisma.ProductWhereUniqueInput{
				ID: &productID,
			},
			Data: prisma.ProductUpdateInput{
				Category: &prisma.ProductCategoryUpdateOneInput{
					Connect: &prisma.ProductCategoryWhereUniqueInput{
						ID: &input.CategoryID,
					},
				},
				SubCategory: &prisma.ProductSubCategoryUpdateOneWithoutProductsInput{
					Connect: &prisma.ProductSubCategoryWhereUniqueInput{
						ID: &input.SubCategoryID,
					},
				},
			},
		}).Exec(ctx)

		if err != nil {
			// Log the error for debugging
			return nil, err
		}
	}

	subCategory, err := r.Prisma.ProductSubCategory(prisma.ProductSubCategoryWhereUniqueInput{
		ID: &input.SubCategoryID,
	}).Exec(ctx)

	if err != nil {
		// Log the error for debugging
		return nil, err
	}

	// Log the successful operation
	return &gqlgen.ReplaceExistingProductsPayload{
		ProductSubCategory: subCategory,
	}, nil
}
