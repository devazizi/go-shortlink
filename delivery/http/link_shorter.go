package http

import (
	"crypto/sha512"
	"encoding/base64"
	"goshorter/adapter"
	"goshorter/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterNewRoute(db adapter.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := dto.RequestCreateNewShort{}
		reqErr := c.Bind(&request)
		if reqErr != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "bad request",
			})
		}

		hashSha512 := sha512.New()
		hashSha512.Write([]byte(request.Url))
		hashCodeOfUrl := hashSha512.Sum(nil)

		base64OfUrlHash := base64.StdEncoding.EncodeToString(hashCodeOfUrl)
		link := adapter.Link{
			Short:    base64OfUrlHash[0:10],
			FullLink: request.Url,
		}

		db.Conn.Create(&link)

		return c.JSON(http.StatusCreated, map[string]any{
			"id":        link.ID,
			"shortUrl":  c.Request().Host + "/r/" + link.Short,
			"fullUrl":   link.FullLink,
			"createdAt": link.CreatedAt,
		})
	}
}

func RedirectToRealLink(db adapter.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		link := adapter.Link{}
		db.Conn.Where("short = ?", id).First(&link)
		link.ClickCount += 1
		db.Conn.Save(&link)

		return c.Redirect(http.StatusOK, link.FullLink)
	}
}
