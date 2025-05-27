package schemas

import (
	"time"

	schemas2 "Mars/database/schemas"
	"Mars/shared/schemas"
)

type NewBuildSchema struct {
	Project  string           `json:"project"`
	Version  string           `json:"version"`
	Family   string           `json:"family"`
	Channel  schemas2.Channel `json:"channel"`
	Promoted bool             `json:"promoted"`
	Changes  []schemas.Change `json:"changes,omitempty"`
}

type NewBuildResult struct {
	Project string `json:"project"`
	Version string `json:"version"`
	Family  string `json:"family"`
	BuildID int    `json:"build_id"`
}

type BuildsSchema struct {
	Build     int                                          `json:"build"`
	Time      time.Time                                    `json:"time"`
	Channel   string                                       `json:"channel"`
	Promoted  bool                                         `json:"promoted"`
	Changes   []schemas.Change                             `json:"changes"`
	Downloads map[string]schemas.ApplicationVersionsSchema `json:"downloads"`
}

type ProjectVersionBuildSchema struct {
	ProjectRootSchema
	Channel   string                                       `json:"channel"`
	Version   string                                       `json:"version"`
	Build     int                                          `json:"build"`
	Promoted  bool                                         `json:"promoted"`
	Time      time.Time                                    `json:"time"`
	Changes   []schemas.Change                             `json:"changes"`
	Downloads map[string]schemas.ApplicationVersionsSchema `json:"downloads"`
}
