package branch

import (
	"context"
	"github.com/steebchen/keskin-api/api/permissions"
	"github.com/steebchen/keskin-api/gqlgen"
	"github.com/steebchen/keskin-api/gqlgen/gqlerrors"
	"github.com/steebchen/keskin-api/lib/file"
	"github.com/steebchen/keskin-api/lib/picture"
	"github.com/steebchen/keskin-api/prisma"
	"strings"
)

func (r *BranchMutation) AddImageSlot(
	ctx context.Context,
	input gqlgen.BranchImageSlotInput,
) (*gqlgen.UpdateBranchPayload, error) {
	if err := permissions.CanAccessBranch(ctx, input.BranchID, r.Prisma, allowedTypes); err != nil {
		return nil, err
	}

	// check if current process is new or old

	if input.New == true {
		// upload image and add it to branch and image slot record.
		id, err := file.MaybeUpload(&input.NewImage, true)
		if err != nil {
			return nil, gqlerrors.NewValidationError("error uploading picture", "ErrorUploadingPicture")
		}

		// get previous images list and add new image to list
		branch, err := r.Prisma.Branch(prisma.BranchWhereUniqueInput{
			ID: &input.BranchID,
		}).Exec(ctx)

		prevImageList := branch.Images

		prevImageList = append(prevImageList, *id)

		// create image structure
		image := picture.FromID(id)

		//find slot and update with new image
		_, err = r.Prisma.CreateBranchImageSlot(prisma.BranchImageSlotCreateInput{
			Branch: &prisma.BranchCreateOneWithoutSlotsInput{
				Connect: &prisma.BranchWhereUniqueInput{
					ID: &input.BranchID,
				},
			},
			SlotNumber: int32(*input.SlotNumber),
			ImageUrl:   image.URL,
		}).Exec(ctx)

		if err != nil {
			return nil, gqlerrors.NewValidationError("error adding slot", "ErrorAddingSlot")
		}

		updatedBranch, err := r.Prisma.UpdateBranch(prisma.BranchUpdateParams{
			Where: prisma.BranchWhereUniqueInput{
				ID: &input.BranchID,
			},
			Data: prisma.BranchUpdateInput{
				Images: &prisma.BranchUpdateimagesInput{
					Set: prevImageList,
				},
			},
		}).Exec(ctx)

		return &gqlgen.UpdateBranchPayload{
			Branch: updatedBranch,
		}, nil

	} else {
		// if old get slotID and replace the image url

		// upload image and add it to branch and image slot record.
		id, err := file.MaybeUpload(&input.NewImage, true)
		if err != nil {
			return nil, gqlerrors.NewValidationError("error uploading picture", "ErrorUploadingPicture")
		}

		// get the slot
		slot, err := r.Prisma.BranchImageSlot(prisma.BranchImageSlotWhereUniqueInput{
			ID: input.SlotID,
		}).Exec(ctx)

		if err != nil {
			return nil, gqlerrors.NewValidationError("error getting slot", "ErrorGettingSLot")
		}

		// get imageId
		imageID := strings.TrimPrefix(slot.ImageUrl, file.BasePath)

		// get previous images list
		branch, err := r.Prisma.Branch(prisma.BranchWhereUniqueInput{
			ID: &input.BranchID,
		}).Exec(ctx)

		if err != nil {
			return nil, gqlerrors.NewValidationError("error getting branch", "ErrorGettingBranch")
		}

		imageList := branch.Images

		var newImageList []string

		// remove old image
		for _, img := range imageList {
			if img != imageID {
				newImageList = append(newImageList, img)
			}
		}

		// add new image to list
		newImageList = append(newImageList, *id)

		newSlotImageUrl := picture.FromID(id)

		// update slot
		_, err = r.Prisma.UpdateBranchImageSlot(prisma.BranchImageSlotUpdateParams{
			Where: prisma.BranchImageSlotWhereUniqueInput{
				ID: input.SlotID,
			},
			Data: prisma.BranchImageSlotUpdateInput{
				ImageUrl: &newSlotImageUrl.URL,
			},
		}).Exec(ctx)

		if err != nil {
			return nil, gqlerrors.NewValidationError("error updating slot", "ErrorUpdatingSlot")
		}

		// update branch with new images

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
			return nil, gqlerrors.NewValidationError("error updating branch", "ErrorUpdatingBranch")
		}

		return &gqlgen.UpdateBranchPayload{
			Branch: updateBranch,
		}, nil

	}

}
