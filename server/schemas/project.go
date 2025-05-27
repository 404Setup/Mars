package schemas

type ProjectRootSchema struct {
	ProjectId   string `json:"project_id"`
	ProjectName string `json:"project_name"`
}

type NewProjectSchema struct {
	ProjectId string `json:"project_id"`
}

type ProjectsSchema struct {
	Projects []string `json:"projects"`
}

type ProjectSchema struct {
	ProjectRootSchema
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}
