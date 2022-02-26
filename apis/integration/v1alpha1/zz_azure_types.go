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

type AzureObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AzureParameters struct {

	// Silence monitors for expected Azure VM shutdowns.
	// +kubebuilder:validation:Optional
	Automute *bool `json:"automute,omitempty" tf:"automute,omitempty"`

	// Your Azure web application ID.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// (Required for Initial Creation) Your Azure web application secret key.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// String of host tag(s) (in the form `key:value,key:value`) defines a filter that Datadog will use when collecting metrics from Azure. Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. e.x. `env:production,deploymentgroup:red`
	// +kubebuilder:validation:Optional
	HostFilters *string `json:"hostFilters,omitempty" tf:"host_filters,omitempty"`

	// Your Azure Active Directory ID.
	// +kubebuilder:validation:Required
	TenantName *string `json:"tenantName" tf:"tenant_name,omitempty"`
}

// AzureSpec defines the desired state of Azure
type AzureSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureParameters `json:"forProvider"`
}

// AzureStatus defines the observed state of Azure.
type AzureStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Azure is the Schema for the Azures API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type Azure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureSpec   `json:"spec"`
	Status            AzureStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureList contains a list of Azures
type AzureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Azure `json:"items"`
}

// Repository type metadata.
var (
	Azure_Kind             = "Azure"
	Azure_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Azure_Kind}.String()
	Azure_KindAPIVersion   = Azure_Kind + "." + CRDGroupVersion.String()
	Azure_GroupVersionKind = CRDGroupVersion.WithKind(Azure_Kind)
)

func init() {
	SchemeBuilder.Register(&Azure{}, &AzureList{})
}
