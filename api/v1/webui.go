package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WebUi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WebUiSpec `json:"spec"`
}

type WebUiSpec struct {
	Contents string `json:"contents"`
	Image    string `json:"image"`
	Replicas int    `json:"replicas"`
}

type WebUiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []WebUi `json:"items"`
}
