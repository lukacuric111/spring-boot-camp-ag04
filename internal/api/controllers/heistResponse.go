package controllers

import (
	"context"
	"elProfessor/internal/api/controllers/models"
)

// MemberResponse implements member related functions
type HeistResponse interface {
	InsertHeist(heistDto models.HeistDto) error
	UpdateHeistSkills(ctx context.Context, heistSkills models.HeistSkillsDto, heistId string) error
	AddHeistMembers(members []string, id string) (string,error)
	StartHeist(id string) (string,error)
	GetHeistById(ctx context.Context, id string) (models.HeistDto, bool, error)
}