package aufs

import (
	"context"
	"testing"

	"github.com/containerd/containerd/snapshots"
	"github.com/containerd/containerd/snapshots/testsuite"
	"github.com/containerd/containerd/testutil"
)

func TestAufs(t *testing.T) {
	testutil.RequiresRoot(t)
	testsuite.SnapshotterSuite(t, "Aufs", func(ctx context.Context, root string) (snapshots.Snapshotter, func() error, error) {
		s, err := New(root)
		if err != nil {
			return nil, nil, err
		}
		return s, s.Close, nil
	})
}
