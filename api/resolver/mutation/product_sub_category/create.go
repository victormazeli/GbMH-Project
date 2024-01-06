package product_sub_category

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductSubCategoryMutation) CreateProductSubCategory(ctx context.Context, input gqlgen.CreateProductSubCategoryInput) (*gqlgen.CreateProductSubCategoryPayload, error) {
	companyID := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	cg, err := r.Prisma.ProductCategory(prisma.ProductCategoryWhereUniqueInput{
		ID: &input.Data.ProductCategoryID,
	}).Exec(ctx)

	if err != nil {
		return nil, gqlerrors.NewNotFoundError(input.Data.ProductCategoryID)
	}

	subCg, err := r.Prisma.CreateProductSubCategory(prisma.ProductSubCategoryCreateInput{
		Name: *input.Data.Name,
		Category: &prisma.ProductCategoryCreateOneWithoutSubCategoriesInput{
			Connect: &prisma.ProductCategoryWhereUniqueInput{
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

	return &gqlgen.CreateProductSubCategoryPayload{
		SubCategory: subCg,
	}, nil
}
