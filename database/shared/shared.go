package shared

import (
	"errors"

	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	ProjectAndVersionNotFoundError = errors.New("project or version not found")
)
