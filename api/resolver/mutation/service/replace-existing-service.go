package service

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceMutation) ReplaceExistingService(
	ctx context.Context,
	input gqlgen.ReplaceExistingServiceInput,
	language *string,
) (*gqlgen.ReplaceExistingServicePayload, error) {

	// Validate input.Services to ensure it's not empty or contains invalid IDs

	if len(input.Services) == 0 {
		return nil, gqlerrors.NewValidationError("Services can not be empty", "EmptyData")
	}

	for _, serviceID := range input.Services {
		_, err := r.Prisma.UpdateService(prisma.ServiceUpdateParams{
			Where: prisma.ServiceWhereUniqueInput{
				ID: &serviceID,
			},
			Data: prisma.ServiceUpdateInput{
				Category: &prisma.ServiceCategoryUpdateOneInput{
					Connect: &prisma.ServiceCategoryWhereUniqueInput{
						ID: &input.CategoryID,
					},
				},
				SubCategory: &prisma.ServiceSubCategoryUpdateOneWithoutServicesInput{
					Connect: &prisma.ServiceSubCategoryWhereUniqueInput{
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

	subCategory, err := r.Prisma.ServiceSubCategory(prisma.ServiceSubCategoryWhereUniqueInput{
		ID: &input.SubCategoryID,
	}).Exec(ctx)

	if err != nil {
		// Log the error for debugging
		return nil, err
	}

	// Log the successful operation
	return &gqlgen.ReplaceExistingServicePayload{
		ServiceSubCategory: subCategory,
	}, nil
}
