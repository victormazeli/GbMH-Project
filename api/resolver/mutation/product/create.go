package product

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"log"

	"github.com/steebchen/keskin-api/api/permissions"
	"github.com/steebchen/keskin-api/api/resolver/product_service_attribute"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/i18n"
	"github.com/steebchen/keskin-api/lib/file"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ProductMutation) CreateProduct(
	ctx context.Context,
	input gqlgen.CreateProductInput,
	language *string,
) (*gqlgen.CreateProductPayload, error) {
	if err := permissions.CanAccessBranch(ctx, input.Branch, r.Prisma, allowedTypes); err != nil {
		return nil, err
	}

	companyID := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	log.Print("companyID >>", companyID)

	imageID, err := file.MaybeUpload(input.Data.Image, true)

	if err != nil {
		return nil, err
	}

	product, err := r.Prisma.CreateProduct(prisma.ProductCreateInput{
		Name:     *i18n.CreateLocalizedString(ctx, &input.Data.Name),
		Desc:     *i18n.CreateLocalizedString(ctx, input.Data.Desc),
		Price:    *prisma.Price(&input.Data.Price),
		Capacity: input.Data.Capacity,
		SubCategory: &prisma.ProductSubCategoryCreateOneWithoutProductsInput{
			Connect: &prisma.ProductSubCategoryWhereUniqueInput{
				ID: &input.Data.Subcategory,
			},
		},
		Category: &prisma.ProductCategoryCreateOneInput{
			Connect: &prisma.ProductCategoryWhereUniqueInput{
				ID: &input.Data.Category,
			},
		},
		Image: imageID,
		Company: &prisma.CompanyCreateOneInput{
			Connect: &prisma.CompanyWhereUniqueInput{
				ID: &companyID,
			},
		},

		Branch: prisma.BranchCreateOneInput{
			Connect: &prisma.BranchWhereUniqueInput{
				ID: &input.Branch,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	err = product_service_attribute.UpsertAttributes(r.Prisma, ctx, &product.ID, nil, input.Data.Attributes)

	if err != nil {
		return nil, err
	}

	return &gqlgen.CreateProductPayload{
		Product: product,
	}, nil
}
