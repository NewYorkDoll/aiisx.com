package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"aiisx.com/src/database/graphql/graph/model"
	"aiisx.com/src/version"
)

// Version is the resolver for the version field.
func (r *queryResolver) Version(ctx context.Context) (*model.VersionInfo, error) {
	var links []*model.Link
	link := &model.Link{
		Name: "aiisx.com",
		URL:  "//aiisx.com"}
	build, ok := debug.ReadBuildInfo()
	links = append(links, link)
	versionINfo := &model.VersionInfo{
		Name:      "",
		Version:   version.Version,
		Commit:    version.CommitHash,
		Date:      version.CompileDate,
		Command:   filepath.Base(os.Args[0]),
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Links:     links,
	}
	if ok {
		versionINfo.Name = build.Main.Path
	}

	return versionINfo, nil
}
