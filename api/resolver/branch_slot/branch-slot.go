package branch_slot

import (
	"github.com/google/wire"

	"github.com/steebchen/keskin-api/prisma"
)

var allowedTypes = []prisma.UserType{
	prisma.UserTypeManager,
}

type BranchImageSlot struct {
	Prisma *prisma.Client
}

func New(client *prisma.Client) *BranchImageSlot {
	return &BranchImageSlot{
		Prisma: client,
	}
}

var ProviderSet = wire.NewSet(
	New,
)
