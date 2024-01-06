package employee

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *EmployeeMutation) EditEmployeeAvailability(
	ctx context.Context,
	input gqlgen.EditEmployeeAvailabilityInput,
) (*gqlgen.EditEmployeeAvailabilityPayload, error) {

	// find the employee by employeeID
	checkEmployee, err := r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &input.EmployeeID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	if checkEmployee.Deleted {
		return nil, gqlerrors.NewPermissionError("Staff is not available")
	}

	// find the working time for the employee and update the status

	workingHour, err := r.Prisma.WorkingHourses(&prisma.WorkingHoursesParams{
		Where: &prisma.WorkingHoursWhereInput{
			ID: &input.EmployeeID,
			Or: []prisma.WorkingHoursWhereInput{
				{
					StartForenoon: prisma.TimeString(input.StartDate),
					EndForenoon:   prisma.TimeString(input.EndDate),
				},
				{
					StartAfternoon: prisma.TimeString(input.StartDate),
					EndAfternoon:   prisma.TimeString(input.EndDate),
				},
			},
		},
	}).Exec(ctx)

	for _, workHourItem := range workingHour {
		_, err := r.Prisma.UpdateWorkingHours(prisma.WorkingHoursUpdateParams{
			Where: prisma.WorkingHoursWhereUniqueInput{
				ID: &workHourItem.ID,
			},
			Data: prisma.WorkingHoursUpdateInput{
				Status: (*prisma.AvailabilityStatus)(&input.Status),
			},
		}).Exec(ctx)

		if err != nil {
			return nil, err
		}
	}

	employee, err := r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &input.EmployeeID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.EditEmployeeAvailabilityPayload{
		Employee: &prisma.Employee{User: employee},
	}, nil

}
