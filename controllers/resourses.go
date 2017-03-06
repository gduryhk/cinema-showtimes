package controllers

import (
	"github.com/gduryhk/cinema-showtimes/models"
)

type (
	// For Get - /showtimes
	ShowTimesResource struct {
		Data []models.ShowTime `json:"data"`
	}
	// For Post/Put - /showtimes
	ShowTimeResource struct {
		Data models.ShowTime `json:"data"`
	}
)
