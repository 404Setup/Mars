package controller

import (
	"Mars/database/schemas"
	"Mars/database/shared"
)

func FindAllProjects() []schemas.Project {
	var projects []schemas.Project
	shared.DB.Find(&projects)
	return projects
}
