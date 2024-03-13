package handler

import (
	"github.com/Schierke/personal-site/view/resume"
	"github.com/labstack/echo/v4"
)

type resumeHandler struct {
}

func NewResumeHandler() *resumeHandler {
	handler := &resumeHandler{}

	return handler
}

func (h resumeHandler) RegisterRoutes(router *echo.Echo) {
	router.GET("/resume", h.ResumeShow)
}

func (h resumeHandler) ResumeShow(c echo.Context) error {
	return render(c, resume.ShowResume())
}
