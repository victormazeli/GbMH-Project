package service

import (
	"context"
	"github.com/steebchen/keskin-api/lib/sessctx"

	"github.com/steebchen/keskin-api/api/permissions"
	"github.com/steebchen/keskin-api/api/resolver/product_service_attribute"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/i18n"
	"github.com/steebchen/keskin-api/lib/file"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceMutation) CreateService(
	ctx context.Context,
	input gqlgen.CreateServiceInput,
	language *string,
) (*gqlgen.CreateServicePayload, error) {
	companyID := sessctx.CompanyWithFallback(ctx, r.Prisma, nil)

	if err := permissions.CanAccessBranch(ctx, input.Branch, r.Prisma, allowedTypes); err != nil {
		return nil, err
	}

	imageID, err := file.MaybeUpload(input.Data.Image, true)

	if err != nil {
		return nil, err
	}

	service, err := r.Prisma.CreateService(prisma.ServiceCreateInput{
		Name:         *i18n.CreateLocalizedString(ctx, &input.Data.Name),
		Desc:         *i18n.CreateLocalizedString(ctx, input.Data.Desc),
		Price:        *prisma.Price(&input.Data.Price),
		Duration:     int32(input.Data.Duration),
		GenderTarget: &input.Data.GenderTarget,
		Image:        imageID,
		SubCategory: &prisma.ServiceSubCategoryCreateOneWithoutServicesInput{
			Connect: &prisma.ServiceSubCategoryWhereUniqueInput{
				ID: &input.Data.Subcategory,
			},
		},
		Category: &prisma.ServiceCategoryCreateOneInput{
			Connect: &prisma.ServiceCategoryWhereUniqueInput{
				ID: &input.Data.Category,
			},
		},

		Branch: prisma.BranchCreateOneInput{
			Connect: &prisma.BranchWhereUniqueInput{
				ID: &input.Branch,
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

	err = product_service_attribute.UpsertAttributes(r.Prisma, ctx, nil, &service.ID, input.Data.Attributes)

	if err != nil {
		return nil, err
	}

	return &gqlgen.CreateServicePayload{
		Service: service,
	}, nil
}
