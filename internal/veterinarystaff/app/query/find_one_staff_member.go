//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneStaffMember struct {
	StaffMemberID string
}

type FindOneStaffMemberReadModel interface {
	FindOneStaffMember(ctx context.Context, staffmemberId uuid.UUID) (StaffMember, error)
}

type FindOneStaffMemberHandler struct {
	readModel FindOneStaffMemberReadModel
}

func NewFindOneStaffMemberHandler(
	readModel FindOneStaffMemberReadModel,
) FindOneStaffMemberHandler {
	return FindOneStaffMemberHandler{readModel: readModel}
}

func (h FindOneStaffMemberHandler) Handle(ctx context.Context, qry FindOneStaffMember) (StaffMember, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.query.find_one_staff_member.handle")
	defer span.End()

	staffmember, err := h.readModel.FindOneStaffMember(ctx, uuid.MustParse(qry.StaffMemberID))
	if err != nil {
		return StaffMember{}, err
	}

	return staffmember, nil
}
