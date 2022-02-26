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

type ArchiveOrderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ArchiveOrderParameters struct {

	// The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If `archive_ids` is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	// +kubebuilder:validation:Optional
	ArchiveIds []*string `json:"archiveIds,omitempty" tf:"archive_ids,omitempty"`
}

// ArchiveOrderSpec defines the desired state of ArchiveOrder
type ArchiveOrderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ArchiveOrderParameters `json:"forProvider"`
}

// ArchiveOrderStatus defines the observed state of ArchiveOrder.
type ArchiveOrderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ArchiveOrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveOrder is the Schema for the ArchiveOrders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type ArchiveOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArchiveOrderSpec   `json:"spec"`
	Status            ArchiveOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveOrderList contains a list of ArchiveOrders
type ArchiveOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArchiveOrder `json:"items"`
}

// Repository type metadata.
var (
	ArchiveOrder_Kind             = "ArchiveOrder"
	ArchiveOrder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ArchiveOrder_Kind}.String()
	ArchiveOrder_KindAPIVersion   = ArchiveOrder_Kind + "." + CRDGroupVersion.String()
	ArchiveOrder_GroupVersionKind = CRDGroupVersion.WithKind(ArchiveOrder_Kind)
)

func init() {
	SchemeBuilder.Register(&ArchiveOrder{}, &ArchiveOrderList{})
}
