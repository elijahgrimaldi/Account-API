package handler

import (
	"log"
	"net/http"

	"github.com/elijahgrimaldi/Account-API/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// bindData is a helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	// Bind incoming JSON to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					Field: err.Field(),
					Value: err.Value().(string),
					Tag:   err.Tag(),
					Param: err.Param(),
				})
			}

			err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

			c.JSON(http.StatusBadRequest, gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
		} else {
			// Handle other types of binding errors, e.g., invalid JSON format
			fallBack := apperrors.NewBadRequest("Bad Request")

			c.JSON(http.StatusBadRequest, gin.H{"error": fallBack})
		}
		return false
	}

	return true
}
