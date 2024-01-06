package branch

import (
	"context"

	"github.com/steebchen/keskin-api/api/permissions"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Branch) SMTPPassword(ctx context.Context, obj *prisma.Branch) (*string, error) {
	if err := permissions.CanAccessBranch(ctx, obj.ID, r.Prisma, allowedTypes); err != nil {
		return nil, nil
	}
	return obj.SmtpPassword, nil
}
