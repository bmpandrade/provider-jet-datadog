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

type AwsTagFilterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AwsTagFilterParameters struct {

	// Your AWS Account ID without dashes. If your account is a GovCloud or China account, specify the `access_key_id` here.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// The namespace associated with the tag filter entry. Valid values are `elb`, `application_elb`, `sqs`, `rds`, `custom`, `network_elb`, `lambda`.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// The tag filter string.
	// +kubebuilder:validation:Required
	TagFilterStr *string `json:"tagFilterStr" tf:"tag_filter_str,omitempty"`
}

// AwsTagFilterSpec defines the desired state of AwsTagFilter
type AwsTagFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsTagFilterParameters `json:"forProvider"`
}

// AwsTagFilterStatus defines the observed state of AwsTagFilter.
type AwsTagFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsTagFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AwsTagFilter is the Schema for the AwsTagFilters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type AwsTagFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsTagFilterSpec   `json:"spec"`
	Status            AwsTagFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsTagFilterList contains a list of AwsTagFilters
type AwsTagFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwsTagFilter `json:"items"`
}

// Repository type metadata.
var (
	AwsTagFilter_Kind             = "AwsTagFilter"
	AwsTagFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AwsTagFilter_Kind}.String()
	AwsTagFilter_KindAPIVersion   = AwsTagFilter_Kind + "." + CRDGroupVersion.String()
	AwsTagFilter_GroupVersionKind = CRDGroupVersion.WithKind(AwsTagFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&AwsTagFilter{}, &AwsTagFilterList{})
}
