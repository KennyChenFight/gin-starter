package dao

import (
	"fmt"
	"time"

	"github.com/KennyChenFight/golib/pglib"

	"github.com/KennyChenFight/gin-starter/pkg/business"
	"github.com/KennyChenFight/golib/loglib"
)

func NewPGMemberDAO(logger *loglib.Logger, client *pglib.GOPGClient) *PGMemberDAO {
	return &PGMemberDAO{logger: logger, client: client}
}

type PGMemberDAO struct {
	logger *loglib.Logger
	client *pglib.GOPGClient
}

func (p *PGMemberDAO) Create(member Member) (string, *business.Error) {
	_, err := p.client.Model(&member).Insert()
	if err != nil {
		fmt.Println(member, err)
		return "", pgErrorHandle(p.logger, err)
	}
	return member.ID, nil
}

func (p *PGMemberDAO) Get(memberID string) (Member, *business.Error) {
	var member Member
	member.ID = memberID
	if err := p.client.Model(&member).WherePK().Where("deleted_at is null").Select(); err != nil {
		return member, pgErrorHandle(p.logger, err)
	}
	return member, nil
}

func (p *PGMemberDAO) Update(member Member) *business.Error {
	targetMember := Member{ID: member.ID}
	err := p.client.Model(&targetMember).WherePK().Where("deleted_at is null").Select()
	if err != nil {
		return pgErrorHandle(p.logger, err)
	}

	targetMember.Email = member.Email
	targetMember.Name = member.Name
	targetMember.UpdatedAt = time.Now()
	_, err = p.client.Model(&targetMember).WherePK().Update()
	if err != nil {
		return pgErrorHandle(p.logger, err)
	}
	return nil
}

func (p *PGMemberDAO) Delete(memberID string) *business.Error {
	var member Member
	member.ID = memberID
	_, err := p.client.Model(&member).WherePK().Delete()
	if err != nil {
		return pgErrorHandle(p.logger, err)
	}
	return nil
}
