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

type GCPObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GCPParameters struct {

	// Silence monitors for expected GCE instance shutdowns.
	// +kubebuilder:validation:Optional
	Automute *bool `json:"automute,omitempty" tf:"automute,omitempty"`

	// Your email found in your JSON service account key.
	// +kubebuilder:validation:Required
	ClientEmail *string `json:"clientEmail" tf:"client_email,omitempty"`

	// Your ID found in your JSON service account key.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.
	// +kubebuilder:validation:Optional
	HostFilters *string `json:"hostFilters,omitempty" tf:"host_filters,omitempty"`

	// Your private key ID found in your JSON service account key.
	// +kubebuilder:validation:Required
	PrivateKeyID *string `json:"privateKeyId" tf:"private_key_id,omitempty"`

	// Your private key name found in your JSON service account key.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// Your Google Cloud project ID found in your JSON service account key.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

// GCPSpec defines the desired state of GCP
type GCPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GCPParameters `json:"forProvider"`
}

// GCPStatus defines the observed state of GCP.
type GCPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GCPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GCP is the Schema for the GCPs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type GCP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GCPSpec   `json:"spec"`
	Status            GCPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPList contains a list of GCPs
type GCPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCP `json:"items"`
}

// Repository type metadata.
var (
	GCP_Kind             = "GCP"
	GCP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GCP_Kind}.String()
	GCP_KindAPIVersion   = GCP_Kind + "." + CRDGroupVersion.String()
	GCP_GroupVersionKind = CRDGroupVersion.WithKind(GCP_Kind)
)

func init() {
	SchemeBuilder.Register(&GCP{}, &GCPList{})
}