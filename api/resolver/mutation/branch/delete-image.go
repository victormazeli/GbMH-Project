package branch

import (
	"context"
	"github.com/steebchen/keskin-api/api/permissions"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/lib/file"
	"github.com/steebchen/keskin-api/prisma"
	"strings"
)

func (r *BranchMutation) DeleteBranchImage(
	ctx context.Context,
	input gqlgen.DeleteBranchImageInput,
) (*gqlgen.DeleteBranchPayload, error) {
	company, err := r.Prisma.Branch(prisma.BranchWhereUniqueInput{
		ID: &input.BranchID,
	}).Company().Exec(ctx)

	if err != nil {
		return nil, err
	}

	if err := permissions.CanAccessCompany(ctx, company.ID, r.Prisma, allowedTypes); err != nil {
		return nil, err
	}

	// get the slot by slotID
	slot, err := r.Prisma.BranchImageSlot(prisma.BranchImageSlotWhereUniqueInput{
		ID: &input.SlotID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// get branch
	branch, err := r.Prisma.Branch(prisma.BranchWhereUniqueInput{
		ID: &input.BranchID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// get imageId
	imageID := strings.TrimPrefix(slot.ImageUrl, file.BasePath)

	imageList := branch.Images

	var newImageList []string

	// remove image
	for _, img := range imageList {
		if img != imageID {
			newImageList = append(newImageList, img)
		}
	}

	_, err = r.Prisma.DeleteBranchImageSlot(prisma.BranchImageSlotWhereUniqueInput{
		ID: &slot.ID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	updateBranch, err := r.Prisma.UpdateBranch(prisma.BranchUpdateParams{
		Where: prisma.BranchWhereUniqueInput{
			ID: &input.BranchID,
		},
		Data: prisma.BranchUpdateInput{
			Images: &prisma.BranchUpdateimagesInput{
				Set: newImageList,
			},
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &gqlgen.DeleteBranchPayload{
		Branch: updateBranch,
	}, nil
}
