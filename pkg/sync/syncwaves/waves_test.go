package syncwaves

import (
	"testing"

	"github.com/stretchr/testify/assert"

	testingutils "github.com/argoproj/gitops-engine/pkg/utils/testing"
)

func TestWave(t *testing.T) {
	assert.Equal(t, 0, Wave(testingutils.NewPod()))
	assert.Equal(t, 1, Wave(testingutils.Annotate(testingutils.NewPod(), "argocd.argoproj.io/sync-wave", "1")))
	assert.Equal(t, 1, Wave(testingutils.Annotate(testingutils.NewPod(), "helm.sh/hook-weight", "1")))
}

func TestWaveOrdering(t *testing.T) {
	assert.Equal(t, "Normal", WaveOrdering(testingutils.NewPod()))
	assert.Equal(t, "Normal", Wave(testingutils.Annotate(testingutils.NewPod(), "argocd.argoproj.io/sync-wave-order", "Normal")))
	assert.Equal(t, "BTree", Wave(testingutils.Annotate(testingutils.NewPod(), "argocd.argoproj.io/sync-wave-order", "BTree")))
}
