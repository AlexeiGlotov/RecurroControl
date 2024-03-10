package handler

import (
	"math/rand"
	"net/http"
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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.GetUserStruct(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var licenseKeys []models.LicenseKeys

	for x := 0; x < inputKeys.CountKeys; x++ {
		var tempLicenseKey models.LicenseKeys

		tempLicenseKey.Creator = user.Login
		tempLicenseKey.Holder = inputKeys.Holder
		tempLicenseKey.TTLCheat = inputKeys.TTLCheat
		tempLicenseKey.LicenseKeys = generateUniqueKey()

		licenseKeys = append(licenseKeys, tempLicenseKey)
	}

	err = h.services.LicenseKeys.CreateLicenseKeys(licenseKeys)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"keys": licenseKeys})
}
