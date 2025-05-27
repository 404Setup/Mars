package schemas

type FamilyVersionsSchema struct {
	ProjectRootSchema
	VersionGroup string   `json:"version_group"`
	Versions     []string `json:"versions"`
}

type FamilyVersionsBuildsSchema struct {
	FamilyVersionsSchema
	Builds []BuildsSchema `json:"builds"`
}
