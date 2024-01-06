package product

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductMutation) SetProductActive(
	ctx context.Context,
	input gqlgen.SetProductActiveInput,
) (*gqlgen.SetProductActivePayload, error) {
	active := false

	if input.Active == false {
		updateProduct, err := r.Prisma.UpdateProduct(prisma.ProductUpdateParams{
			Where: prisma.ProductWhereUniqueInput{
				ID: &input.ProductID,
			},
			Data: prisma.ProductUpdateInput{
				Active: &active,
			},
		}).Exec(ctx)

		if err != nil {
			return nil, err
		}
		return &gqlgen.SetProductActivePayload{
			Product: updateProduct,
		}, nil
	} else {
		active = true
		updateProduct, err := r.Prisma.UpdateProduct(prisma.ProductUpdateParams{
			Where: prisma.ProductWhereUniqueInput{
				ID: &input.ProductID,
			},
			Data: prisma.ProductUpdateInput{
				Active: &active,
			},
		}).Exec(ctx)

		if err != nil {
			return nil, err
		}
		return &gqlgen.SetProductActivePayload{
			Product: updateProduct,
		}, nil
	}
}
