package services

import (
	"context"
	domainmodels "elProfessor/internal/api/controllers/models"
	"github.com/gin-gonic/gin"
)

type HeistHandler interface {
	InsertHeist(heistDto domainmodels.HeistDto) (string, error)
	UpdateHeistSkills(ctx context.Context, heistSkills domainmodels.HeistSkillsDto, heistId string) error
	AddHeistMembers(members []string, id string) (string,error, []string)
	StartHeist(id string) (string,error)
	GetHeistById(ctx context.Context, id string) (domainmodels.HeistDto, bool, error)
	GetHeistMembersByHeistId(ctx context.Context, id string) ([]domainmodels.MemberDto, bool, error)
	GetHeistSkillsByHeistId(ctx *gin.Context, id string) (domainmodels.HeistSkillsDto, error)
	GetHeistStatusByHeistId(ctx *gin.Context, id string) (string, error)
	EndHeist(id string) (string, error)
	GetHeistOutcomeByHeistId(ctx *gin.Context, id string) (string, bool, error)
}