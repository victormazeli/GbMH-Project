package service

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *ServiceMutation) SetServiceActive(
	ctx context.Context,
	input gqlgen.SetServiceActiveInput,
) (*gqlgen.SetServiceActivePayload, error) {
	active := false

	if input.Active == false {
		updateService, err := r.Prisma.UpdateService(prisma.ServiceUpdateParams{
			Where: prisma.ServiceWhereUniqueInput{
				ID: &input.ServiceID,
			},
			Data: prisma.ServiceUpdateInput{
				Active: &active,
			},
		}).Exec(ctx)

		if err != nil {
			return nil, err
		}
		return &gqlgen.SetServiceActivePayload{
			Service: updateService,
		}, nil
	} else {
		active = true
		updateService, err := r.Prisma.UpdateService(prisma.ServiceUpdateParams{
			Where: prisma.ServiceWhereUniqueInput{
				ID: &input.ServiceID,
			},
			Data: prisma.ServiceUpdateInput{
				Active: &active,
			},
		}).Exec(ctx)

		if err != nil {
			return nil, err
		}
		return &gqlgen.SetServiceActivePayload{
			Service: updateService,
		}, nil
	}
}
