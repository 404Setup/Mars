package helpful

import (
	schemas2 "Mars/database/schemas"
	"Mars/server/schemas"
)

func CreateBuildResponse(project schemas2.Project, version schemas2.Version, build schemas2.Build) *schemas.ProjectVersionBuildSchema {
	changes := build.UnmarshalChanges()

	downloads := build.UnmarshalDownloads()

	return &schemas.ProjectVersionBuildSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   project.ID,
			ProjectName: project.Name,
		},
		Channel:   string(build.Channel),
		Version:   version.Name,
		Build:     build.Number,
		Time:      build.Time,
		Promoted:  build.Promoted,
		Changes:   changes,
		Downloads: downloads,
	}
}
