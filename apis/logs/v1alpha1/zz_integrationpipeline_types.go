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

type IntegrationPipelineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IntegrationPipelineParameters struct {

	// Boolean value to enable your pipeline.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

// IntegrationPipelineSpec defines the desired state of IntegrationPipeline
type IntegrationPipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationPipelineParameters `json:"forProvider"`
}

// IntegrationPipelineStatus defines the observed state of IntegrationPipeline.
type IntegrationPipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationPipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationPipeline is the Schema for the IntegrationPipelines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type IntegrationPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationPipelineSpec   `json:"spec"`
	Status            IntegrationPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationPipelineList contains a list of IntegrationPipelines
type IntegrationPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationPipeline `json:"items"`
}

// Repository type metadata.
var (
	IntegrationPipeline_Kind             = "IntegrationPipeline"
	IntegrationPipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationPipeline_Kind}.String()
	IntegrationPipeline_KindAPIVersion   = IntegrationPipeline_Kind + "." + CRDGroupVersion.String()
	IntegrationPipeline_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationPipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationPipeline{}, &IntegrationPipelineList{})
}
