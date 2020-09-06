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

			// get config
			config, ok := ic.Config.(*Config)
			if !ok {
				return nil, errors.New("invalid aufs configuration")
			}

			// use default ic.Root as root path if config doesn't have a valid root path
			root := ic.Root
			if len(config.RootPath) != 0 {
				root = config.RootPath
			}
			ic.Meta.Exports["root"] = ic.Root

			snapshotter, err := aufs.New(root)
			if err != nil {
				return nil, errors.Wrap(plugin.ErrSkipPlugin, err.Error())
			}
			return snapshotter, nil
		},
	})
}
