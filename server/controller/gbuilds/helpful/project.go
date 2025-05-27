package helpful

import (
	"sort"

	schemas2 "Mars/database/schemas"
	"Mars/server/schemas"
)

func CreateProjectResponse(project schemas2.Project, families []schemas2.VersionFamily, versions []schemas2.Version) *schemas.ProjectSchema {
	sort.Slice(families, func(i, j int) bool {
		return families[i].Name < families[j].Name
	})

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Name < versions[j].Name
	})

	familyNames := make([]string, len(families))
	for i, family := range families {
		familyNames[i] = family.Name
	}

	versionNames := make([]string, len(versions))
	for i, version := range versions {
		versionNames[i] = version.Name
	}
	return &schemas.ProjectSchema{
		ProjectRootSchema: schemas.ProjectRootSchema{
			ProjectId:   project.ID,
			ProjectName: project.Name,
		},
		VersionGroups: familyNames,
		Versions:      versionNames,
	}
}
