/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GcsSpec defines the desired state of Gcs
type GcsSpec struct {
	ProjectID  string `json:"projectID"`
	BucketName string `json:"bucketName"`
}

// GcsStatus defines the observed state of Gcs
type GcsStatus struct {
	ProjectID  string `json:"projectID"`
	BucketName string `json:"bucketName"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Gcs is the Schema for the gcs API
type Gcs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GcsSpec   `json:"spec,omitempty"`
	Status GcsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GcsList contains a list of Gcs
type GcsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gcs `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Gcs{}, &GcsList{})
}
