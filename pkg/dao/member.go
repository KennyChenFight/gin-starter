package dao

import (
	"time"

	"github.com/KennyChenFight/gin-starter/pkg/util"
)

type Member struct {
	ID             string    `json:"id"`
	Email          string    `json:"email" binding:"required,email"`
	PasswordDigest string    `json:"-"`
	Name           string    `json:"name" binding:"required"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `pg:",soft_delete" json:"deletedAt"`
}

type MemberDAO interface {
	Create(member Member) (string, *util.BusinessError)
	Get(memberID string) (Member, *util.BusinessError)
	Update(member Member) *util.BusinessError
	Delete(memberID string) *util.BusinessError
}
