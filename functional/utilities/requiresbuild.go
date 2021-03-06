package testutilities

import (
	"testing"

	"github.com/Microsoft/hcsshim/osversion"
)

func RequiresBuild(t *testing.T, b uint16) {
	if osversion.Get().Build < b {
		t.Skipf("Requires build %d+", b)
	}
}
