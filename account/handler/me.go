package handler

import (
	"net/http"

	"github.com/elijahgrimaldi/Account-API/model"
	"github.com/elijahgrimaldi/Account-API/model/apperrors"
	"github.com/gin-gonic/gin"
)

// Me handler calls services for getting
// a user's details
func (h *Handler) Me(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	uid := user.(*model.User).UID

	u, err := h.UserService.Get(c, uid)

	if err != nil {
		e := apperrors.NewNotFound("user", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
