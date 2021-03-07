package service

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/KennyChenFight/gin-starter/pkg/util"

	"github.com/KennyChenFight/gin-starter/pkg/dao"
	"github.com/gin-gonic/gin"
)

func (s *BaseService) CreateMember(c *gin.Context) {
	var request struct {
		dao.Member
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		sendErrorResponse(c, s, util.NewBusinessError(util.InvalidParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}
	request.ID = uuid.NewV4().String()
	if digest, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost); err != nil {
		sendErrorResponse(c, s, util.NewBusinessError(util.Unknown, http.StatusBadRequest, "internal server error", err))
		return
	} else {
		request.PasswordDigest = string(digest)
	}

	member := dao.Member{ID: request.ID, Name: request.Name, Email: request.Email, PasswordDigest: request.PasswordDigest}
	memberID, err := s.memberDAO.Create(member)
	if err != nil {
		sendErrorResponse(c, s, err)
		return
	}
	sendSuccessResponse(c, http.StatusCreated, gin.H{"memberID": memberID})
}

func (s *BaseService) GetMember(c *gin.Context) {
	memberID := c.Param("id")
	if memberID == "" {
		sendErrorResponse(c, s, util.NewBusinessError(util.InvalidParse, http.StatusBadRequest, "invalid memberID", nil))
		return
	}

	member, err := s.memberDAO.Get(memberID)
	if err != nil {
		sendErrorResponse(c, s, err)
		return
	}
	sendSuccessResponse(c, http.StatusOK, member)
}

func (s *BaseService) UpdateMember(c *gin.Context) {
	var requestMemberID struct {
		ID string `uri:"id" binding:"uuid4"`
	}
	if err := c.ShouldBindUri(&requestMemberID); err != nil {
		sendErrorResponse(c, s, util.NewBusinessError(util.InvalidParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}

	memberID := requestMemberID.ID
	var requestMember struct {
		dao.Member
	}
	if err := c.ShouldBindJSON(&requestMember); err != nil {
		sendErrorResponse(c, s, util.NewBusinessError(util.InvalidParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}

	member := dao.Member{ID: memberID, Email: requestMember.Email, Name: requestMember.Name}
	err := s.memberDAO.Update(member)
	if err != nil {
		sendErrorResponse(c, s, err)
		return
	}
	sendSuccessResponse(c, http.StatusNoContent, nil)
}

func (s *BaseService) DeleteMember(c *gin.Context) {
	var requestMemberID struct {
		ID string `uri:"id" binding:"uuid4"`
	}
	if err := c.ShouldBindUri(&requestMemberID); err != nil {
		sendErrorResponse(c, s, util.NewBusinessError(util.InvalidParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}

	memberID := requestMemberID.ID
	err := s.memberDAO.Delete(memberID)
	if err != nil {
		sendErrorResponse(c, s, err)
		return
	}
	sendSuccessResponse(c, http.StatusNoContent, nil)
}
