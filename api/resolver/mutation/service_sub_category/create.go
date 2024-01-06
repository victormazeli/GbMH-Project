package service_sub_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceSubCategoryMutation) CreateServiceSubCategory(ctx context.Context, input gqlgen.CreateServiceSubCategoryInput) (*gqlgen.CreateServiceSubCategoryPayload, error) {

	companyID := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	cg, err := r.Prisma.ServiceCategory(prisma.ServiceCategoryWhereUniqueInput{
		ID: &input.Data.CategoryID,
	}).Exec(ctx)

	if err != nil {
		return nil, gqlerrors.NewNotFoundError(input.Data.CategoryID)
	}

	subCg, err := r.Prisma.CreateServiceSubCategory(prisma.ServiceSubCategoryCreateInput{
		Name: *input.Data.Name,
		Category: &prisma.ServiceCategoryCreateOneWithoutSubCategoriesInput{
			Connect: &prisma.ServiceCategoryWhereUniqueInput{
				ID: &cg.ID,
			},
		},
		Company: &prisma.CompanyCreateOneInput{
			Connect: &prisma.CompanyWhereUniqueInput{
				ID: &companyID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.CreateServiceSubCategoryPayload{
		SubCategory: subCg,
	}, nil
}
