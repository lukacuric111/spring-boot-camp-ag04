package services

import (
	"context"
	domainmodels "elProfessor/internal/api/controllers/models"
)

type MemberHandler interface {
	InsertMember(memberDto domainmodels.MemberDto) error
	UpdateMemberSkills(ctx context.Context, memberSkillsUpdate domainmodels.MemberSkillsUpdateDto, memberId string) error
}