/*
DataInfra Pinot Operator (C) 2023 - 2024 DataInfra.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PinotSchemaSpec defines the desired state of PinotSchema
type PinotSchemaSpec struct {
	SchemaJson string `json:"schema.json,omitempty"`
}

// PinotSchemaStatus defines the observed state of PinotSchema
type PinotSchemaStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PinotSchema is the Schema for the pinotschemas API
type PinotSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinotSchemaSpec   `json:"spec,omitempty"`
	Status PinotSchemaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PinotSchemaList contains a list of PinotSchema
type PinotSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinotSchema `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PinotSchema{}, &PinotSchemaList{})
}
