/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivateLocationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateLocationParameters struct {

	// Description of the private location.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A list of tags to associate with your synthetics private location.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PrivateLocationSpec defines the desired state of PrivateLocation
type PrivateLocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateLocationParameters `json:"forProvider"`
}

// PrivateLocationStatus defines the observed state of PrivateLocation.
type PrivateLocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateLocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLocation is the Schema for the PrivateLocations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type PrivateLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateLocationSpec   `json:"spec"`
	Status            PrivateLocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateLocationList contains a list of PrivateLocations
type PrivateLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateLocation `json:"items"`
}

// Repository type metadata.
var (
	PrivateLocation_Kind             = "PrivateLocation"
	PrivateLocation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateLocation_Kind}.String()
	PrivateLocation_KindAPIVersion   = PrivateLocation_Kind + "." + CRDGroupVersion.String()
	PrivateLocation_GroupVersionKind = CRDGroupVersion.WithKind(PrivateLocation_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateLocation{}, &PrivateLocationList{})
}
