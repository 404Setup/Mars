//go:build amd64

package sonic

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"

	"github.com/3JoB/ulib/timings"
	"github.com/bytedance/sonic"

	schemas2 "Mars/server/schemas"
	"Mars/shared/configure"
	"Mars/shared/schemas"
)

func PreTouch() {
	if configure.Get().Json.Runtime != "sonic" || !configure.Get().Json.Sonic.JITPretouch {
		return
	}
	timer := time.Now()
	errors := 0
	var b = []reflect.Type{
		// b
		reflect.TypeOf(schemas.Change{}),
		reflect.TypeOf([]schemas.Change{}),
		reflect.TypeOf(schemas.NewDownloadSchema{}),
		reflect.TypeOf(schemas.IndexVersion{}),
		reflect.TypeOf(schemas.ApplicationVersionsSchema{}),
		reflect.TypeOf(map[string]schemas.ApplicationVersionsSchema{}),

		// c
		reflect.TypeOf(schemas2.ProjectRootSchema{}),
		reflect.TypeOf(schemas2.ResultSchema{}),
		reflect.TypeOf(schemas2.ErrorSchema{}),
		reflect.TypeOf(schemas2.NewBuildSchema{}),
		reflect.TypeOf(schemas2.NewVersionGroupSchema{}),
		reflect.TypeOf(schemas2.NewVersionSchema{}),
		reflect.TypeOf(schemas2.NewProjectSchema{}),
		reflect.TypeOf(schemas2.NewBuildResult{}),
		reflect.TypeOf(schemas2.BuildsSchema{}),
		reflect.TypeOf([]schemas2.BuildsSchema{}),
		reflect.TypeOf(schemas2.FamilyVersionsSchema{}),
		reflect.TypeOf(schemas2.FamilyVersionsBuildsSchema{}),
		reflect.TypeOf(schemas2.ProjectVersionBuildSchema{}),
		reflect.TypeOf(schemas2.ProjectVersionsBuildsSchema{}),
		reflect.TypeOf(schemas2.ProjectVersionsSchema{}),
		reflect.TypeOf(schemas2.ProjectSchema{}),
		reflect.TypeOf(schemas2.ProjectsSchema{}),
	}
	timings.ParallelForEach(b, timings.AttrSlice(b), func(v int, b reflect.Type) {
		if err := sonic.Pretouch(b); err != nil {
			errors++
		}
	})
	slog.Info(fmt.Sprintf("JITPretouch completed, compilation failed %v times, took %v ms.", errors, time.Since(timer).Milliseconds()))
}
