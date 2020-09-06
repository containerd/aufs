package plugin

import (
	"github.com/containerd/aufs"
	"github.com/containerd/containerd/platforms"
	"github.com/containerd/containerd/plugin"
	"github.com/pkg/errors"
)


// Config represents configuration for the zfs plugin
type Config struct {
	// Root directory for the plugin
	RootPath string `toml:"root_path"`
}


func init() {
	plugin.Register(&plugin.Registration{
		Type: plugin.SnapshotPlugin,
		ID:   "aufs",
		Config: &Config{},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			ic.Meta.Platforms = append(ic.Meta.Platforms, platforms.DefaultSpec())
			ic.Meta.Exports["root"] = ic.Root

			snapshotter, err := aufs.New(ic.Root)
			if err != nil {
				return nil, errors.Wrap(plugin.ErrSkipPlugin, err.Error())
			}
			return snapshotter, nil
		},
	})
}
