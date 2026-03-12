//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllStaffMembersReadModel interface {
	FindAllStaffMembers(ctx context.Context) ([]StaffMember, error)
}

type FindAllStaffMembersHandler struct {
	readModel FindAllStaffMembersReadModel
}

func NewFindAllStaffMembersHandler(
	readModel FindAllStaffMembersReadModel,
) FindAllStaffMembersHandler {
	return FindAllStaffMembersHandler{readModel: readModel}
}

func (h FindAllStaffMembersHandler) Handle(ctx context.Context) ([]StaffMember, error) {
	ctx, span := telemetree.AddSpan(ctx, "veterinarystaff.app.query.find_all_staff_members.handle")
	defer span.End()

	exams, err := h.readModel.FindAllStaffMembers(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
