package models

import (
	"html/template"
	"time"
)

var Shortcuts *template.Template
var Lists []string
var Original []string
var Dir string

type Session struct {
	UserID       string
	LastActivity time.Time
}

var DBSessions = map[string]Session{}
