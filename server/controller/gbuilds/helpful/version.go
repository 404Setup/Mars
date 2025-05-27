package helpful

import (
	schemas2 "Mars/database/schemas"
	"Mars/server/schemas"
)

func CreateVersionResponse(project schemas2.Project, version schemas2.Version, builds []schemas2.Build) *schemas.ProjectVersionsSchema {
	buildNumbers := make([]int, len(builds))
	for i, build := range builds {
		buildNumbers[i] = build.Number
	}
	return &schemas.ProjectVersionsSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   project.ID,
			ProjectName: project.Name,
		},
		Version: version.Name,
		Builds:  buildNumbers,
	}
}
