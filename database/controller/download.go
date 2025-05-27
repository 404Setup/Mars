package controller

import (
	schemas2 "Mars/database/schemas"
	"Mars/database/shared"
)

func CreateDownload(project, version string, build int, app schemas2.Build) error {
	/*, err := FindBuildByProjectAndVersionAndNumber(project, version, build)
	if err != nil {
		return err
	}*/
	return shared.DB.Where("project = ? AND version = ? AND number = ?", project, version, build).Updates(app).Error
}
