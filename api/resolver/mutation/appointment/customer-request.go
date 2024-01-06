package appointment

import (
	"context"
	"github.com/steebchen/keskin-api/api/resolver/mutation/email_template"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/i18n"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Appointment) CustomerRequestAppointment(
	ctx context.Context,
	input gqlgen.CustomerRequestAppointmentInput,
	language *string,
) (*gqlgen.CustomerRequestAppointmentPayload, error) {
	user, err := sessctx.User(ctx)
	if err != nil {
		return nil, err
	}
	respMessage := "appointment booked successfully"
	// Verify branch existence
	_, err = r.Prisma.Branch(prisma.BranchWhereUniqueInput{ID: &input.Branch}).Exec(ctx)
	if err != nil {
		return nil, gqlerrors.NewFormatNodeError(err, input.Branch)
	}

	if input.Employee != nil {
		// Create appointments for each employee
		staff, err := r.validateAndRetrieveStaff(ctx, input.Employee)
		if err != nil {
			return nil, err
		}

		appointment, err := r.createAppointment(ctx, user.ID, *input.Employee, input)
		if err != nil {
			return nil, err
		}

		if input.Data.BeforeImage != nil {
			beforeImagePayload := gqlgen.UpdateAppointmentImageInput{
				Appointment: appointment.ID,
				Image:       *input.Data.BeforeImage,
			}
			_, err := r.UpdateBeforeImage(ctx, beforeImagePayload, nil)

			if err != nil {
				return nil, err
			}
		}

		go r.sendBookingNotification(ctx, staff, appointment, input.Branch)

		return &gqlgen.CustomerRequestAppointmentPayload{
			Message: &respMessage,
		}, nil

	}

	// fetch all employees for that branch
	employees, err := r.Prisma.Branch(prisma.BranchWhereUniqueInput{
		ID: &input.Branch,
	}).Employees(nil).Exec(ctx)

	// Create appointments for each employee
	for _, employee := range employees {
		staff, err := r.validateAndRetrieveStaff(ctx, &employee.ID)
		if err != nil {
			return nil, err
		}

		appointment, err := r.createAppointment(ctx, user.ID, employee.ID, input)

		if err != nil {
			return nil, err
		}

		if input.Data.BeforeImage != nil {
			beforeImagePayload := gqlgen.UpdateAppointmentImageInput{
				Appointment: appointment.ID,
				Image:       *input.Data.BeforeImage,
			}
			_, err := r.UpdateBeforeImage(ctx, beforeImagePayload, nil)

			if err != nil {
				return nil, err
			}
		}

		go r.sendBookingNotification(ctx, staff, appointment, input.Branch)
	}

	return &gqlgen.CustomerRequestAppointmentPayload{
		Message: &respMessage,
	}, nil
}

func (r *Appointment) validateAndRetrieveStaff(ctx context.Context, employeeID *string) (*prisma.User, error) {
	staff, err := r.Prisma.User(prisma.UserWhereUniqueInput{ID: employeeID}).Exec(ctx)
	if err != nil {
		return nil, gqlerrors.NewFormatNodeError(err, *employeeID)
	}

	if staff.Type != prisma.UserTypeManager && staff.Type != prisma.UserTypeEmployee {
		return nil, gqlerrors.NewValidationError("staff type is invalid", "InvalidUserType")
	}

	if staff.Deleted {
		return nil, gqlerrors.NewPermissionError("Staff is not available")
	}

	return staff, nil
}

//func (r *Appointment) checkIfAppointmentBooked(ctx context.Context, employeeID *string, startTime time.Time, endTime time.Time) (*bool, error) {
//
//	appointment, err := r.Prisma.Appointments(&prisma.AppointmentsParams{
//		Where: &prisma.AppointmentWhereInput{
//			Start:
//		}
//	}).Exec(ctx)
//
//}

func (r *Appointment) createAppointment(
	ctx context.Context,
	customerID string,
	employeeID string,
	input gqlgen.CustomerRequestAppointmentInput,
) (*prisma.Appointment, error) {
	return CreateAppointment(CreateAppointmentInput{
		Client:          r.Prisma,
		Context:         ctx,
		EmployeeID:      &employeeID,
		CustomerID:      customerID,
		Branch:          input.Branch,
		Desc:            *input.Data.Desc,
		Start:           input.Data.Start,
		ProductRequests: input.Data.Products,
		ServiceRequests: input.Data.Services,
		DefaultStatus:   prisma.AppointmentStatusRequested,
	})
}

func (r *Appointment) sendBookingNotification(
	ctx context.Context,
	staff *prisma.User,
	appointment *prisma.Appointment,
	branchID string,
) {
	appointmentDate := i18n.FormatDate(ctx, appointment.Start)
	appointmentTime := i18n.FormatTime(ctx, appointment.Start)

	notificationContext := context.Background()
	notificationContext = sessctx.SetLanguage(notificationContext, staff.Language)

	go email_template.SendEmailTemplate(
		notificationContext,
		r.Prisma,
		"appointmentBooked",
		branchID,
		staff.Email,
		staff.Gender,
		staff.LastName,
		staff.FirstName,
		&appointmentDate,
		&appointmentTime,
		nil,
		nil,
		&staff.FirstName,
		&staff.LastName,
	)

}
