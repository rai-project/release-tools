package cmd

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `manifest.template`,
		FileModTime: time.Unix(1503356312, 0),
		Content:     string("---\napplications:\n- path: .\n  memory: {{ .MemorySize }}\n  instances: 1\n  domain: mybluemix.net\n  name: {{ .Package }}\n  host: {{ .HostName }}\n  disk_quota: {{ .DistQuota }}\n  buildpack: https://github.com/cloudfoundry/go-buildpack.git\n  env:\n    GOVERSION: go1.7\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `version_gen.go.template`,
		FileModTime: time.Unix(1503356312, 0),
		Content:     string("package {{ .Package }}\n\nconst (\n\tVersionInfo   = \"{{ .VersionInfo }}\"\n\tRepository    = \"{{ .Repository }}\"\n\tBuildTimeInfo = \"{{ .BuildTimeInfo }}\"\n)\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1503356312, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // manifest.template
			file3, // version_gen.go.template

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`_fixtures`, &embedded.EmbeddedBox{
		Name: `_fixtures`,
		Time: time.Unix(1503356312, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"manifest.template":       file2,
			"version_gen.go.template": file3,
		},
	})
}
