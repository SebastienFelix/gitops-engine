package syncwaves

import (
	"strconv"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/argoproj/gitops-engine/pkg/sync/common"
	helmhook "github.com/argoproj/gitops-engine/pkg/sync/hook/helm"
)

func Wave(obj *unstructured.Unstructured) int {
	text, ok := obj.GetAnnotations()[common.AnnotationSyncWave]
	if ok {
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
	}
	return helmhook.Weight(obj)
}

func WaveOrdering(obj *unstructured.Unstructured) string {
	text, ok := obj.GetAnnotations()[common.AnnotationSyncWaveOrder]
	if ok {
		if text == "BTree" {
			return text
		}
		return "Normal"
	}
	return "Normal"
}
