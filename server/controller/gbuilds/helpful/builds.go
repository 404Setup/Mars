package helpful

import (
	schemas2 "Mars/database/schemas"
	"Mars/server/schemas"
)

func DatabaseBuildsToServerBuilds(builds []schemas2.Build) []schemas.BuildsSchema {
	build := make([]schemas.BuildsSchema, len(builds))
	if builds == nil || len(builds) == 0 {
		return build
	}
	for i, b := range builds {
		downloads := b.UnmarshalDownloads()
		changes := b.UnmarshalChanges()
		build[i] = schemas.BuildsSchema{
			Build:     b.Number,
			Time:      b.Time,
			Channel:   string(b.Channel),
			Promoted:  b.Promoted,
			Changes:   changes,
			Downloads: downloads,
		}
	}
	return build
}

func CreateBuildsResponse(project schemas2.Project, version schemas2.Version, builds []schemas2.Build) *schemas.ProjectVersionsBuildsSchema {
	buildResponses := make([]schemas.BuildsSchema, len(builds))
	for i, build := range builds {
		downloads := build.UnmarshalDownloads()
		changes := build.UnmarshalChanges()

		buildResponses[i] = schemas.BuildsSchema{
			Build:     build.Number,
			Time:      build.Time,
			Channel:   string(build.Channel),
			Promoted:  build.Promoted,
			Changes:   changes,
			Downloads: downloads,
		}
	}

	return &schemas.ProjectVersionsBuildsSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   project.ID,
			ProjectName: project.Name,
		},
		Version: version.Name,
		Builds:  buildResponses,
	}
}
