package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"RecurroControl/models"
)

func generateUniqueKey() string {
	const charset = "ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz123456789"

	seed := int64(time.Now().UnixNano())
	src := rand.NewSource(seed)
	r := rand.New(src)

	key := make([]byte, 36)
	for i := 0; i < 36; i++ {
		key[i] = charset[r.Intn(len(charset))]
	}

	return string(key)
}

func (h *Handler) createLicenseKeys(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var inputKeys models.InputCreate
	if err := c.BindJSON(&inputKeys); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	if len(inputKeys.Notes) > 256 {
		newErrorResponse(c, http.StatusBadRequest, err, "max sim 256")
		return
	}

	if inputKeys.CountKeys > 5 && user.Role != models.Admin {
		newErrorResponse(c, http.StatusBadRequest, err, "limit keys")
		return
	}

	var licenseKeys []models.LicenseKeys

	for x := 0; x < inputKeys.CountKeys; x++ {
		var tempLicenseKey models.LicenseKeys

		tempLicenseKey.Creator = user.Login
		tempLicenseKey.Holder = inputKeys.Holder
		tempLicenseKey.TTLCheat = inputKeys.TTLCheat
		tempLicenseKey.Cheat = inputKeys.Cheat
		tempLicenseKey.LicenseKeys = generateUniqueKey()
		tempLicenseKey.Notes = &inputKeys.Notes

		licenseKeys = append(licenseKeys, tempLicenseKey)
	}

	err = h.services.LicenseKeys.CreateLicenseKeys(licenseKeys)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"keys": licenseKeys})
}

func (h *Handler) getLicenseKeys(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		return
	}

	pageStr, ok := c.GetQuery("page")
	if !ok {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, "invalid page format")
		return
	}

	query, ok := c.GetQuery("query")
	if !ok {
		query = ""
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	offset := (page - 1) * 100

	usersI, err := h.services.Users.GetHierarchyUsers(user.Login)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	keys, err := h.services.LicenseKeys.GetLicenseKeys(user.Login, user.Role, usersI, 100, offset, query)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"keys": keys})
}

func (h *Handler) getCustomLicenseKeys(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	date, ok := c.GetQuery("date")
	if !ok {
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err, "set date query")
			return
		}
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	t, err := time.Parse("2006-01-02", date)

	res, err := h.services.LicenseKeys.GetCustomLicenseKeys(user.Login, t)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"res": res})

}

type inputAction struct {
	Id int `json:"id" binding:"required"`
}

func (h *Handler) licenseKeyResetHWID(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	err = h.services.LicenseKeys.ResetHWID(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) licenseKeyBan(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	if user.Role == models.Salesman {
		newErrorResponse(c, http.StatusBadRequest, err, ErrAccessDenied)
		return
	}

	err = h.services.LicenseKeys.Ban(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}

type InputData struct {
	Date string `json:"date" binding:"required"`
}

func (h *Handler) licenseKeysBanDate(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input InputData
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	if user.Role == models.Salesman || user.Role == models.Reseller {
		newErrorResponse(c, http.StatusBadRequest, err, ErrAccessDenied)
		return
	}

	t, err := time.Parse("2006-01-02", input.Date)
	fmt.Println(t)
	err = h.services.LicenseKeys.BanOfDate(user.Login, t)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) licenseKeyUnban(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	if user.Role == models.Salesman {
		newErrorResponse(c, http.StatusBadRequest, err, ErrAccessDenied)
		return
	}

	err = h.services.LicenseKeys.Unban(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) licenseKeyDelete(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputAction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err, ErrNotFields)
		return
	}

	user, err := h.services.Users.GetUser(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	if user.Role != models.Admin {
		newErrorResponse(c, http.StatusBadRequest, err, ErrAccessDenied)
		return
	}

	err = h.services.LicenseKeys.Delete(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err, ErrServerError)
		return
	}

	c.Status(http.StatusOK)
}
