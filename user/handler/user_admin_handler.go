package handler

import (
	"ecommerce/model"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (e *UserHandler) viewAllUser(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}

	user, err := e.userUsecase.ViewAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, user)
}

func (e *UserHandler) userUpdateAdmin(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	var user = model.User{}
	err = c.Bind(&user)
	if err != nil {
		fmt.Printf("[Userhandler.updateUser] error bind data %v \n", err)
		utils.HandleError(c, http.StatusBadRequest, "Opps Sever has be wrong")
		return
	}

	_, err = e.userUsecase.ViewById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	if user.Password != "" || user.ID != 0 {
		utils.HandleError(c, http.StatusForbidden, "no password updates")
		return
	}

	upUser, err := e.userUsecase.UpdateUser(id, &user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
	return
}
	utils.HandleSucces(c, upUser)
}

func (e *UserHandler) getUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}

	_, err = e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := e.userUsecase.ViewById(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSucces(c, user)
}

func (e *UserHandler) blashEmailVertifiled(c *gin.Context) {
	user, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var newsletter = model.NewsletterInput{}
	err = c.Bind(&newsletter)
	if err != nil {
		fmt.Printf("[userHandler.blashEmaiVertifiled] Error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Ooops server somting wrong")
		return
	}
	if newsletter.Article == "" || newsletter.Subjec == "" {
		utils.HandleError(c, http.StatusBadRequest, "column cannot be empty")
		return
	}
	users, err := e.userUsecase.ViewAll()
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	counter := 0
	for i := 0; i < len(*users); i++ {
		if  (*users)[i].Status == "vertifiled" {
			utils.NewsLetterEmail((*users)[i].Email, newsletter.Subjec, "asset/template/email_message_newsletter.html", newsletter.Article)
			counter++
		}
	}
	news := model.Newletter{
		DateSend: time.Now().Format("2006-01-02"),
		Subject: newsletter.Subjec,
		Article: newsletter.Article,
		UserID: user.ID,
		MessageSend: counter,
	}
	_, err = e.userUsecase.InsertNewsletter(&news)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, "Succes send email newsletter")
}

func (e *UserHandler) checkMitraWaiting(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	mitra, err := e.userUsecase.ViewAllMitra()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var waitMitra []model.Mitra
	for i := 0; i < len(*mitra); i++ {
		if (*mitra)[i].Status == "waiting for approval" {
			waitMitra = append(waitMitra, (*mitra)[i])
		}
	}
	if waitMitra == nil {
		utils.HandleError(c, http.StatusNotFound, "no partner submission data awaiting approval")
		return
	}
	utils.HandleSucces(c, waitMitra)
}

func (e *UserHandler) approveMitra(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	status := "approved"
	user, mitra, err := e.userUsecase.RequestMitraProcess(id, status)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	user.Role = "mitra"
	upUser, err := e.userUsecase.UpdateUser(int(user.ID), user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WelcomeEmail("Ecommerce - Permintaan Upgrate Mitra Telah Disetujui","", upUser.Email, "", "asset/template/email_approve_mitra.html")
	utils.HandleSucces(c, mitra)
}

func (e *UserHandler) rejectMitra(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "id has be number")
		return
	}
	status := "rejected"
	user, mitra, err := e.userUsecase.RequestMitraProcess(id, status)
	if err != nil {
		utils.HandleError(c, http.StatusForbidden, err.Error())
		return
	}
	utils.WelcomeEmail("Ecommerce - Permintaan Upgrate Mitra Telah Ditolak","", user.Email, "", "asset/template/email_reject_mitra.html")
	utils.HandleSucces(c, mitra)
}

func (e *UserHandler) historyNewsletter(c *gin.Context) {
	_, err := e.userUsecase.AdminOnly(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	news, err := e.userUsecase.ViewAllNewsletter()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(c, news)
}


