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

type CaseObservation struct {
}

type CaseParameters struct {

	// Notification targets for each rule case.
	// +kubebuilder:validation:Required
	Notifications []*string `json:"notifications" tf:"notifications,omitempty"`

	// Status of the rule case to match. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	// +kubebuilder:validation:Required
	Status *string `json:"status" tf:"status,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// The type of filtering action. Allowed enum values: require, suppress Valid values are `require`, `suppress`.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Query for selecting logs to apply the filtering action.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`
}

type MonitoringDefaultRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitoringDefaultRuleParameters struct {

	// Cases of the rule, this is used to update notifications.
	// +kubebuilder:validation:Optional
	Case []CaseParameters `json:"case,omitempty" tf:"case,omitempty"`

	// Enable the rule.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Additional queries to filter matched events before they are processed.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`
}

// MonitoringDefaultRuleSpec defines the desired state of MonitoringDefaultRule
type MonitoringDefaultRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitoringDefaultRuleParameters `json:"forProvider"`
}

// MonitoringDefaultRuleStatus defines the observed state of MonitoringDefaultRule.
type MonitoringDefaultRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitoringDefaultRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringDefaultRule is the Schema for the MonitoringDefaultRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type MonitoringDefaultRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitoringDefaultRuleSpec   `json:"spec"`
	Status            MonitoringDefaultRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringDefaultRuleList contains a list of MonitoringDefaultRules
type MonitoringDefaultRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringDefaultRule `json:"items"`
}

// Repository type metadata.
var (
	MonitoringDefaultRule_Kind             = "MonitoringDefaultRule"
	MonitoringDefaultRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitoringDefaultRule_Kind}.String()
	MonitoringDefaultRule_KindAPIVersion   = MonitoringDefaultRule_Kind + "." + CRDGroupVersion.String()
	MonitoringDefaultRule_GroupVersionKind = CRDGroupVersion.WithKind(MonitoringDefaultRule_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitoringDefaultRule{}, &MonitoringDefaultRuleList{})
}
