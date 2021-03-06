package api

import (
	"github.com/codersgarage/golang-restful-boilerplate/app"
	"github.com/codersgarage/golang-restful-boilerplate/errors"
	"github.com/codersgarage/golang-restful-boilerplate/models"
	"github.com/codersgarage/golang-restful-boilerplate/repos"
	"net/http"
)

type MonkeyRepo interface {
	SaveMonkey(s *app.Scope) (*models.Monkey, error)
	GetMonkey(s *app.Scope) (*models.Monkey, error)
	UpdateMonkey(s *app.Scope) (*models.Monkey, error)
	ListMonkey(s *app.Scope) ([]models.Monkey, error)
	DeleteMonkey(s *app.Scope) error
}

type MonkeyRoutes struct {
	Repo MonkeyRepo
}

func NewMonkeyRoutes() *MonkeyRoutes {
	return &MonkeyRoutes{
		Repo: repos.NewMonkeyRepo(),
	}
}

func (hr *MonkeyRoutes) saveMonkey(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Monkey, err := hr.Repo.SaveMonkey(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusCreated
	resp.Data = Monkey
	resp.Title = "Monkey has been created"
	resp.ServerJSON(w)
}

func (hr *MonkeyRoutes) getMonkey(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Monkey, err := hr.Repo.GetMonkey(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Monkey
	resp.ServerJSON(w)
}

func (hr *MonkeyRoutes) updateMonkey(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Monkey, err := hr.Repo.UpdateMonkey(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Monkey
	resp.ServerJSON(w)
}

func (hr *MonkeyRoutes) listMonkey(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Monkeys, err := hr.Repo.ListMonkey(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Monkeys
	resp.ServerJSON(w)
}

func (hr *MonkeyRoutes) deleteMonkey(w http.ResponseWriter, r *http.Request) {
	// TODO
}
