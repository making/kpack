package v1alpha1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1alpha1 "knative.dev/pkg/apis/duck/v1alpha1"
)

const (
	BuilderNotFound = "BuilderNotFound"
	BuilderNotReady = "BuilderNotReady"
)

func (im *Image) BuilderNotFound() duckv1alpha1.Conditions {
	return duckv1alpha1.Conditions{
		{
			Type:               duckv1alpha1.ConditionReady,
			Status:             corev1.ConditionFalse,
			Reason:             BuilderNotFound,
			Message:            fmt.Sprintf("Unable to find builder %s.", im.Spec.Builder.Name),
			LastTransitionTime: apis.VolatileTime{Inner: metav1.Now()},
		},
	}
}
