package aufs

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/containerd/containerd/reaper"
	"github.com/containerd/containerd/snapshot"
	"github.com/containerd/containerd/snapshot/testsuite"
	"github.com/containerd/containerd/testutil"
)

func init() {
	// start a reaper to the tests because aufs execs modprobe
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGCHLD)
	go func() {
		for range s {
			reaper.Reap()
		}
	}()
}

// NOTE: aufs tests should not run on travis.ci beacuse their nodes do not support aufs
func TestAufs(t *testing.T) {
	testutil.RequiresRoot(t)
	testsuite.SnapshotterSuite(t, "Aufs", func(ctx context.Context, root string) (snapshot.Snapshotter, func() error, error) {
		s, err := New(root)
		if err != nil {
			return nil, nil, err
		}
		return s, nil, nil
	})
}
