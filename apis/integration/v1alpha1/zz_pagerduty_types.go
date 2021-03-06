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

type PagerdutyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PagerdutyParameters struct {

	// Your PagerDuty API token.
	// +kubebuilder:validation:Optional
	APITokenSecretRef *v1.SecretKeySelector `json:"apiTokenSecretRef,omitempty" tf:"-"`

	// Array of your schedule URLs.
	// +kubebuilder:validation:Optional
	Schedules []*string `json:"schedules,omitempty" tf:"schedules,omitempty"`

	// Your PagerDuty account’s personalized subdomain name.
	// +kubebuilder:validation:Required
	Subdomain *string `json:"subdomain" tf:"subdomain,omitempty"`
}

// PagerdutySpec defines the desired state of Pagerduty
type PagerdutySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PagerdutyParameters `json:"forProvider"`
}

// PagerdutyStatus defines the observed state of Pagerduty.
type PagerdutyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PagerdutyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pagerduty is the Schema for the Pagerdutys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type Pagerduty struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PagerdutySpec   `json:"spec"`
	Status            PagerdutyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PagerdutyList contains a list of Pagerdutys
type PagerdutyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pagerduty `json:"items"`
}

// Repository type metadata.
var (
	Pagerduty_Kind             = "Pagerduty"
	Pagerduty_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pagerduty_Kind}.String()
	Pagerduty_KindAPIVersion   = Pagerduty_Kind + "." + CRDGroupVersion.String()
	Pagerduty_GroupVersionKind = CRDGroupVersion.WithKind(Pagerduty_Kind)
)

func init() {
	SchemeBuilder.Register(&Pagerduty{}, &PagerdutyList{})
}
