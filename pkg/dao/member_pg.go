package dao

import (
	"fmt"
	"time"

	"github.com/KennyChenFight/gin-starter/pkg/util"

	"github.com/go-pg/pg/v10"
)

func NewPGMemberDAO(client *pg.DB) *PGMemberDAO {
	return &PGMemberDAO{client: client}
}

type PGMemberDAO struct {
	client *pg.DB
}

func (p *PGMemberDAO) Create(member Member) (string, *util.BusinessError) {
	_, err := p.client.Model(&member).Insert()
	if err != nil {
		fmt.Println(member, err)
		return "", pgErrorHandle(err)
	}
	return member.ID, nil
}

func (p *PGMemberDAO) Get(memberID string) (Member, *util.BusinessError) {
	var member Member
	member.ID = memberID
	if err := p.client.Model(&member).WherePK().Where("deleted_at is null").Select(); err != nil {
		return member, pgErrorHandle(err)
	}
	return member, nil
}

func (p *PGMemberDAO) Update(member Member) *util.BusinessError {
	targetMember := Member{ID: member.ID}
	err := p.client.Model(&targetMember).WherePK().Where("deleted_at is null").Select()
	if err != nil {
		return pgErrorHandle(err)
	}

	targetMember.Email = member.Email
	targetMember.Name = member.Name
	targetMember.UpdatedAt = time.Now()
	_, err = p.client.Model(&targetMember).WherePK().Update()
	if err != nil {
		return pgErrorHandle(err)
	}
	return nil
}

func (p *PGMemberDAO) Delete(memberID string) *util.BusinessError {
	var member Member
	member.ID = memberID
	_, err := p.client.Model(&member).WherePK().Delete()
	if err != nil {
		return pgErrorHandle(err)
	}
	return nil
}
