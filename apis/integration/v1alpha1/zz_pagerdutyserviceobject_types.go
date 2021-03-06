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

type PagerdutyServiceObjectObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PagerdutyServiceObjectParameters struct {

	// Your Service name associated service key in PagerDuty. Note: Since the Datadog API never returns service keys, it is impossible to detect [drifts](https://www.hashicorp.com/blog/detecting-and-managing-drift-with-terraform). The best way to solve a drift is to manually mark the Service Object resource with [terraform taint](https://www.terraform.io/docs/commands/taint.html) to have it destroyed and recreated.
	// +kubebuilder:validation:Required
	ServiceKeySecretRef v1.SecretKeySelector `json:"serviceKeySecretRef" tf:"-"`

	// Your Service name in PagerDuty.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// PagerdutyServiceObjectSpec defines the desired state of PagerdutyServiceObject
type PagerdutyServiceObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PagerdutyServiceObjectParameters `json:"forProvider"`
}

// PagerdutyServiceObjectStatus defines the observed state of PagerdutyServiceObject.
type PagerdutyServiceObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PagerdutyServiceObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PagerdutyServiceObject is the Schema for the PagerdutyServiceObjects API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type PagerdutyServiceObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PagerdutyServiceObjectSpec   `json:"spec"`
	Status            PagerdutyServiceObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PagerdutyServiceObjectList contains a list of PagerdutyServiceObjects
type PagerdutyServiceObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PagerdutyServiceObject `json:"items"`
}

// Repository type metadata.
var (
	PagerdutyServiceObject_Kind             = "PagerdutyServiceObject"
	PagerdutyServiceObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PagerdutyServiceObject_Kind}.String()
	PagerdutyServiceObject_KindAPIVersion   = PagerdutyServiceObject_Kind + "." + CRDGroupVersion.String()
	PagerdutyServiceObject_GroupVersionKind = CRDGroupVersion.WithKind(PagerdutyServiceObject_Kind)
)

func init() {
	SchemeBuilder.Register(&PagerdutyServiceObject{}, &PagerdutyServiceObjectList{})
}
