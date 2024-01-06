package branch

import (
	"context"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/prisma"
)

func (r *Branch) Slots(ctx context.Context, obj *prisma.Branch) ([]*gqlgen.BranchImageSlot, error) {

	slots, err := r.Prisma.BranchImageSlots(&prisma.BranchImageSlotsParams{
		Where: &prisma.BranchImageSlotWhereInput{
			Branch: &prisma.BranchWhereInput{
				ID: &obj.ID,
			},
		},
	}).Exec(ctx)

	var nodes []*gqlgen.BranchImageSlot

	if err != nil {
		return nil, err
	}
	//branchImages, err := r.Images(ctx, obj)
	//
	//if err != nil {
	//	return nil, err
	//}

	for _, slot := range slots {
		//imageID := strings.TrimPrefix(img.URL, file.BasePath)

		clone := gqlgen.BranchImageSlot{
			SlotNumber: int(slot.SlotNumber),
			ImageURL:   slot.ImageUrl,
			UpdatedAt:  slot.UpdatedAt,
			CreatedAt:  slot.CreatedAt,
			ID:         slot.ID,
		}
		nodes = append(nodes, &clone)

	}

	return nodes, err

}
