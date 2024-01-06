package appointment

import (
	"context"
	"fmt"
	"github.com/steebchen/keskin-api/api/permissions"
	"github.com/steebchen/keskin-api/api/resolver/mutation/email_template"
	"github.com/steebchen/keskin-api/api/resolver/mutation/notification"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/i18n"
	"github.com/steebchen/keskin-api/lib/hours"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/prisma"
)

var statusApproved = prisma.AppointmentStatusApproved

func (r *Appointment) StaffApproveAppointment(
	ctx context.Context,
	input gqlgen.StaffApproveAppointmentInput,
	language *string,
) (*gqlgen.StaffApproveAppointmentPayload, error) {
	branch, err := r.Prisma.Appointment(prisma.AppointmentWhereUniqueInput{
		ID: &input.ID,
	}).Branch().Exec(ctx)

	if err != nil {
		return nil, err
	}

	t := []prisma.UserType{prisma.UserTypeEmployee, prisma.UserTypeManager}
	if err := permissions.CanAccessBranch(ctx, branch.ID, r.Prisma, t); err != nil {
		return nil, err
	}

	if input.Patch == nil {
		input.Patch = &gqlgen.StaffApproveAppointmentPatch{}
	}

	// check if appointment is already approved, if approved return error

	checkAppointment, err := r.Prisma.Appointment(prisma.AppointmentWhereUniqueInput{
		ID: &input.ID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	if checkAppointment.Status == statusApproved {
		return nil, gqlerrors.NewValidationError("appointment already approved", "InvalidApprovalAction")
	}
	appointment, err := r.Prisma.UpdateAppointment(prisma.AppointmentUpdateParams{
		Where: prisma.AppointmentWhereUniqueInput{
			ID: &input.ID,
		},
		Data: prisma.AppointmentUpdateInput{
			Status: &statusApproved,
			Desc:   i18n.UpdateLocalizedString(ctx, input.Patch.Desc),
			Note:   input.Patch.Note,
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// update status working hours of employee
	err = r.FindEmployeeHoursAndUpdate(appointment, ctx)

	if err != nil {
		return nil, err
	}

	customer, err := r.Prisma.Appointment(prisma.AppointmentWhereUniqueInput{
		ID: &input.ID,
	}).Customer().Exec(ctx)

	if err != nil {
		return nil, err
	}

	if customer.Deleted {
		return nil, gqlerrors.NewPermissionError("Customer is deleted")
	}

	appointmentDate := i18n.FormatDate(ctx, appointment.Start)
	appointmentTime := i18n.FormatTime(ctx, appointment.Start)

	notificationContext := context.Background()
	notificationContext = sessctx.SetLanguage(notificationContext, customer.Language)

	go notification.Send(
		r.Prisma,
		r.MessagingClient,
		notificationContext,
		customer.ID,
		i18n.Language(notificationContext)["APPOINTMENT_APPROVED_TITLE"],
		fmt.Sprintf(
			i18n.Language(notificationContext)["APPOINTMENT_APPROVED_TEXT"],
			appointmentDate,
			appointmentTime,
		),
	)

	go email_template.SendEmailTemplate(
		notificationContext,
		r.Prisma,
		"appointmentConfirmed",
		branch.ID,
		customer.Email,
		customer.Gender,
		customer.LastName,
		customer.FirstName,
		&appointmentDate,
		&appointmentTime,
		nil,
		nil,
		nil,
		nil,
	)

	return &gqlgen.StaffApproveAppointmentPayload{
		Appointment: appointment,
	}, nil
}

func (r *Appointment) FindEmployeeHoursAndUpdate(appointment *prisma.Appointment, ctx context.Context) error {

	employee, err := r.Prisma.Appointment(prisma.AppointmentWhereUniqueInput{
		ID: &appointment.ID,
	}).Employee().Exec(ctx)

	if err != nil {
		return err
	}

	workingHours, err := r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &employee.ID,
	}).WorkingHours(nil).Exec(ctx)

	if err != nil {
		return err
	}

	startToday := hours.TodaySetTime(&appointment.Start)
	endToday := hours.TodaySetTime(&appointment.End)
	var newStatus = prisma.AvailabilityStatusBooked

	for _, workingHoursItem := range workingHours {
		isWithinWorkingHours := false

		forenoonStart := hours.TodaySetTime(workingHoursItem.StartForenoon)
		forenoonEnd := hours.TodaySetTime(workingHoursItem.EndForenoon)

		if (forenoonStart.Before(startToday) ||
			forenoonStart.Equal(startToday)) &&
			(forenoonEnd.After(endToday) ||
				forenoonEnd.Equal(endToday)) {
			isWithinWorkingHours = true
		}

		if !isWithinWorkingHours && workingHoursItem.Break {
			afternoonStart := hours.TodaySetTime(workingHoursItem.StartAfternoon)
			afternoonEnd := hours.TodaySetTime(workingHoursItem.EndAfternoon)

			if (afternoonStart.Before(startToday) ||
				afternoonStart.Equal(startToday)) &&
				(afternoonEnd.After(endToday) ||
					afternoonEnd.Equal(endToday)) {
				isWithinWorkingHours = true
			}
		}

		if isWithinWorkingHours {
			_, err := r.Prisma.UpdateWorkingHours(prisma.WorkingHoursUpdateParams{
				Where: prisma.WorkingHoursWhereUniqueInput{
					ID: &workingHoursItem.ID,
				},
				Data: prisma.WorkingHoursUpdateInput{
					Status: &newStatus,
				},
			}).Exec(ctx)

			if err != nil {
				return err
			}
		}
	}
	return nil

}
