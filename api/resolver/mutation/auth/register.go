package auth

import (
	"context"
	"fmt"

	"github.com/steebchen/keskin-api/api/resolver/mutation/email_template"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/auth"
	"github.com/steebchen/keskin-api/lib/sessctx"
	"github.com/steebchen/keskin-api/lib/users"
	"github.com/steebchen/keskin-api/lib/validator"
	"github.com/steebchen/keskin-api/prisma"
)

// Registers a user of type customer
func (a *Auth) Register(ctx context.Context, input gqlgen.RegisterInput) (*gqlgen.RegisterPayload, error) {
	companyID := sessctx.CompanyWithFallback(ctx, a.Prisma, input.Company)

	// a User of type Customer requires a company relation
	if companyID == "" {
		return nil, gqlerrors.NewInternalError("company id is required", "CompanyHeaderRequired")
	}

	err := validator.Phone(input.PhoneNumber)
	if err != nil {
		return nil, gqlerrors.NewValidationError(err.Error(), "InvalidPhoneNumber")
	}

	err = validator.Email(input.Email)
	if err != nil {
		return nil, gqlerrors.NewValidationError(err.Error(), "InvalidEmail")
	}

	emailInUse, err := users.EmailInUse(ctx, a.Prisma, input.Email, &companyID, nil, nil)

	if err != nil {
		return nil, err
	}

	if emailInUse {
		return nil, gqlerrors.NewValidationError("Email already used for another account", "DuplicateEmail")
	}

	activateToken, err := GenerateActivateToken(a.Prisma, ctx)
	fmt.Printf("TOKEN: %v", activateToken)

	if err != nil {
		return nil, err
	}

	birthdate := ""

	if input.Birthday != nil {
		birthdate = (*input.Birthday)[5:10]
	}

	_, err = a.Prisma.CreateUser(prisma.UserCreateInput{
		Email:         input.Email,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Birthday:      input.Birthday,
		Birthdate:     &birthdate,
		Gender:        &input.Gender,
		PhoneNumber:   &input.PhoneNumber,
		PasswordHash:  auth.HashPassword(input.Password),
		Type:          prisma.UserTypeCustomer,
		ActivateToken: &activateToken,
		Activated:     prisma.Bool(false),
		Company: &prisma.CompanyCreateOneWithoutUsersInput{
			Connect: &prisma.CompanyWhereUniqueInput{
				ID: &companyID,
			},
		},
	}).Exec(ctx)

	if err != nil {
		fmt.Printf("error %+v", err)
		return nil, err
	}

	branches, err := a.Prisma.Branches(&prisma.BranchesParams{
		Where: &prisma.BranchWhereInput{
			Company: &prisma.CompanyWhereInput{
				ID: &companyID,
			},
		},
	}).Exec(ctx)

	if err == nil && len(branches) > 0 {
		_, err := email_template.SendEmailTemplate(context.Background(), a.Prisma, "register", branches[0].ID, input.Email, input.Gender, input.LastName, input.FirstName, nil, nil, nil, &activateToken, nil, nil)
		if err != nil {
			return nil, err
		}
	}

	return &gqlgen.RegisterPayload{
		Status: "OK",
	}, nil
}
