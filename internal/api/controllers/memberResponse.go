package controllers

import (
	"context"
	"elProfessor/internal/api/controllers/models"
)

// MemberResponse implements member related functions
type MemberResponse interface {
	InsertMember(memberDto models.MemberDto) (error)
	UpdateMemberSkills(ctx context.Context, memberSkillsUpdate models.MemberSkillsUpdateDto, memberId string) error
}