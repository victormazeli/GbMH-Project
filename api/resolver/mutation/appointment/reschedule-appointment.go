package appointment

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/prisma"
	"time"
)

func (r *Appointment) RescheduleAppointment(
	ctx context.Context,
	input gqlgen.RescheduleAppointmentInput,
	language *string,
) (*gqlgen.RescheduleAppointmentPayload, error) {

	//1. check if appointment exist
	employee, err := r.Prisma.Appointment(prisma.AppointmentWhereUniqueInput{
		ID: &input.ID,
	}).Employee().Exec(ctx)

	//if errors.Is(err, prisma.ErrNoResult) {
	//
	//	//return nil, errors.New("could not find an appointment")
	//}else {
	//	return nil, err
	//}

	if err != nil {
		return nil, gqlerrors.NewFormatNodeError(err, input.ID)

	}

	employeeID := employee.ID

	// 2. check if customer wants to use the same employee (employeeID)
	if input.EmployeeID != nil {
		staff, err := r.Prisma.User(prisma.UserWhereUniqueInput{
			ID: input.EmployeeID,
		}).Exec(ctx)

		if err != nil {
			return nil, gqlerrors.NewFormatNodeError(err, *input.EmployeeID)
		}

		if staff.Type != prisma.UserTypeManager && staff.Type != prisma.UserTypeEmployee {
			return nil, gqlerrors.NewValidationError("staff type is invalid", "InvalidUserType")

		}

		if staff.Deleted {
			return nil, gqlerrors.NewPermissionError("Staff is not available")
		}
		employeeID = *input.EmployeeID
	}

	// 3. update with new time
	var duration time.Duration
	var services []*prisma.Service

	serviceLinks, err := r.Prisma.AppointmentServiceLinks(&prisma.AppointmentServiceLinksParams{
		Where: &prisma.AppointmentServiceLinkWhereInput{
			Appointment: &prisma.AppointmentWhereInput{
				ID: &input.ID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	for _, serviceLink := range serviceLinks {
		service, err := r.Prisma.AppointmentServiceLink(prisma.AppointmentServiceLinkWhereUniqueInput{
			ID: &serviceLink.ID,
		}).Service().Exec(ctx)

		if err != nil {
			return nil, err
		}

		clone := service
		services = append(services, clone)
	}

	for _, service := range services {
		duration += time.Duration(int(service.Duration)) * time.Minute
	}

	start := *prisma.TimeString(input.NewDate)
	end := *prisma.TimeString(input.NewDate.Add(duration))

	updateAppointment, err := r.Prisma.UpdateAppointment(prisma.AppointmentUpdateParams{
		Where: prisma.AppointmentWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.AppointmentUpdateInput{
			Start: &start,
			End:   &end,
			Employee: &prisma.UserUpdateOneRequiredInput{
				Connect: &prisma.UserWhereUniqueInput{
					ID: &employeeID,
				},
			},
		},
	}).Exec(ctx)

	return &gqlgen.RescheduleAppointmentPayload{
		Appointment: updateAppointment,
	}, nil
}
