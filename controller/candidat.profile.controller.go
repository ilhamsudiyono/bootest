package controller

import (
	"github.com/labstack/echo"

	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var candidatProfileService service.CandidatProfileService = service.CandidatProfileServiceImpl{}

func SetCandidatProfile(c *echo.Group) {
	c.POST("/candidate", createCandidatProfile)
}

func createCandidatProfile(c echo.Context) (e error) {
	defer catchError(&e)
	data := new(model.CandidateProfiles)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	var err = candidatProfileService.CreateCandidat(data)
	if err == nil {
		return res(c, data)
	}
	return resErr(c, err)
}
