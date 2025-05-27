package schemas

type NewVersionSchema struct {
	Version string `json:"version"`
	Group   string `json:"group"`
	Project string `json:"project"`
}

type NewVersionGroupSchema struct {
	Name    string `json:"name"`
	Project string `json:"project"`
}

type ProjectVersionsSchema struct {
	ProjectRootSchema
	Version string `json:"version"`
	Builds  []int  `json:"builds"`
}

type ProjectVersionsBuildsSchema struct {
	ProjectRootSchema
	Version string         `json:"version"`
	Builds  []BuildsSchema `json:"builds"`
}
