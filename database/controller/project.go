package controller

import (
	"Mars/database/schemas"
	"Mars/database/shared"
	"Mars/shared/utils"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func newProject(projectName string) *gorm.DB {
	project := schemas.Project{
		ID:           strings.ToLower(projectName),
		Name:         projectName,
		FriendlyName: projectName,
	}
	return shared.DB.Create(&project)
}

func CreateProject(projectName string) (bool, error) {
	if !utils.IsValidProjectName(projectName) {
		return false, fmt.Errorf("invalid project name: failed string mandatory validation")
	}

	var project schemas.Project

	if err := shared.DB.Where("id = ?", strings.ToLower(projectName)).First(&project).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = newProject(projectName).Error; err != nil {
				return false, fmt.Errorf("failed to create project %s: %w", projectName, err)
			}
			return true, nil
		}
		return false, err
	}
	return false, fmt.Errorf("project %s already exists", projectName)
}

// FindProjectByID finds a Project by its id
func FindProjectByID(id string) (schemas.Project, error) {
	if !utils.IsValidProjectName(id) {
		return schemas.Project{}, fmt.Errorf("invalid project id: failed string mandatory validation")
	}
	var project schemas.Project
	result := shared.DB.Where("id = ?", id).First(&project)
	return project, result.Error
}
