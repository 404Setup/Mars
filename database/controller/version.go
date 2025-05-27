package controller

import (
	"time"

	"gorm.io/gorm"

	"Mars/database/schemas"
	"Mars/database/shared"
)

func newVersion(project, name, group string) *gorm.DB {
	now := time.Now().UTC()
	version := schemas.Version{
		Project: project,
		Group:   group,
		Name:    name,
		Time:    &now,
	}
	return shared.DB.Create(&version)
}

func CreateVersion(project, name, group string) (bool, string) {
	var version schemas.Version
	if err := shared.DB.Where("project = ? AND name = ? AND `group` = ?", project, name, group).First(&version).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = newVersion(project, name, group).Error; err != nil {
				return false, err.Error()
			}
			return true, "done"
		}
		return false, err.Error()
	}
	return false, "The same content already exists"
}

// FindVersionByProjectAndName finds a Version by project and name
func FindVersionByProjectAndName(project, name string) (schemas.Version, error) {
	var version schemas.Version
	result := shared.DB.Where("project = ? AND name = ?", project, name).First(&version)
	return version, result.Error
}
