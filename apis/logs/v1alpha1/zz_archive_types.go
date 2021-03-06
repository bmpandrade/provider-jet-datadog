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

type ArchiveObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ArchiveParameters struct {

	// Definition of an azure archive.
	// +kubebuilder:validation:Optional
	AzureArchive []AzureArchiveParameters `json:"azureArchive,omitempty" tf:"azure_archive,omitempty"`

	// Definition of a GCS archive.
	// +kubebuilder:validation:Optional
	GcsArchive []GcsArchiveParameters `json:"gcsArchive,omitempty" tf:"gcs_archive,omitempty"`

	// To store the tags in the archive, set the value `true`. If it is set to `false`, the tags will be dropped when the logs are sent to the archive.
	// +kubebuilder:validation:Optional
	IncludeTags *bool `json:"includeTags,omitempty" tf:"include_tags,omitempty"`

	// The archive query/filter. Logs matching this query are included in the archive.
	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// An array of tags to add to rehydrated logs from an archive.
	// +kubebuilder:validation:Optional
	RehydrationTags []*string `json:"rehydrationTags,omitempty" tf:"rehydration_tags,omitempty"`

	// Definition of an s3 archive.
	// +kubebuilder:validation:Optional
	S3Archive []S3ArchiveParameters `json:"s3Archive,omitempty" tf:"s3_archive,omitempty"`
}

type AzureArchiveObservation struct {
}

type AzureArchiveParameters struct {

	// Your client id.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The container where the archive will be stored.
	// +kubebuilder:validation:Required
	Container *string `json:"container" tf:"container,omitempty"`

	// The path where the archive will be stored.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The associated storage account.
	// +kubebuilder:validation:Required
	StorageAccount *string `json:"storageAccount" tf:"storage_account,omitempty"`

	// Your tenant id.
	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type GcsArchiveObservation struct {
}

type GcsArchiveParameters struct {

	// Name of your GCS bucket.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Your client email.
	// +kubebuilder:validation:Required
	ClientEmail *string `json:"clientEmail" tf:"client_email,omitempty"`

	// Path where the archive will be stored.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Your project id.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

type S3ArchiveObservation struct {
}

type S3ArchiveParameters struct {

	// Your AWS account id.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// Name of your s3 bucket.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Path where the archive will be stored.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Your AWS role name
	// +kubebuilder:validation:Required
	RoleName *string `json:"roleName" tf:"role_name,omitempty"`
}

// ArchiveSpec defines the desired state of Archive
type ArchiveSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ArchiveParameters `json:"forProvider"`
}

// ArchiveStatus defines the observed state of Archive.
type ArchiveStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ArchiveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Archive is the Schema for the Archives API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type Archive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArchiveSpec   `json:"spec"`
	Status            ArchiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveList contains a list of Archives
type ArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Archive `json:"items"`
}

// Repository type metadata.
var (
	Archive_Kind             = "Archive"
	Archive_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Archive_Kind}.String()
	Archive_KindAPIVersion   = Archive_Kind + "." + CRDGroupVersion.String()
	Archive_GroupVersionKind = CRDGroupVersion.WithKind(Archive_Kind)
)

func init() {
	SchemeBuilder.Register(&Archive{}, &ArchiveList{})
}
