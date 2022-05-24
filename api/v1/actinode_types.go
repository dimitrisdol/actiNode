/*
Copyright 2022.

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

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ActiNodeSpec defines the desired state of ActiNode
type ActiNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ActiNode. Edit actinode_types.go to remove/update
	Cpus    int    `json:"cpus,omitempty"`
	Pods    string `json:"pods,omitempty"`
	Mapping string `json:"mapping,omitempty"`
}

// ActiNodeStatus defines the observed state of ActiNode
type ActiNodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ActiNode is the Schema for the actinodes API
type ActiNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActiNodeSpec   `json:"spec,omitempty"`
	Status ActiNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ActiNodeList contains a list of ActiNode
type ActiNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ActiNode{}, &ActiNodeList{})
}
