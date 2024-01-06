package service_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceCategoryMutation) CreateServiceCategory(ctx context.Context, input gqlgen.CreateServiceCategoryInput) (*gqlgen.CreateServiceCategoryPayload, error) {
	// user, err := sessctx.User(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// if user.Type != prisma.UserTypeAdministrator {
	// 	return nil, gqlerrors.NewPermissionError(gqlerrors.InvalidUserType)
	// }

	companyID := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	cg, err := r.Prisma.CreateServiceCategory(prisma.ServiceCategoryCreateInput{
		Name: *input.Data.Name,
		Company: &prisma.CompanyCreateOneInput{
			Connect: &prisma.CompanyWhereUniqueInput{
				ID: &companyID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.CreateServiceCategoryPayload{
		Category: cg,
	}, nil
}
