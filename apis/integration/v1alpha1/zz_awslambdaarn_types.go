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

type AwsLambdaArnObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AwsLambdaArnParameters struct {

	// Your AWS Account ID without dashes. If your account is a GovCloud or China account, specify the `access_key_id` here.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// The ARN of the Datadog forwarder Lambda.
	// +kubebuilder:validation:Required
	LambdaArn *string `json:"lambdaArn" tf:"lambda_arn,omitempty"`
}

// AwsLambdaArnSpec defines the desired state of AwsLambdaArn
type AwsLambdaArnSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsLambdaArnParameters `json:"forProvider"`
}

// AwsLambdaArnStatus defines the observed state of AwsLambdaArn.
type AwsLambdaArnStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsLambdaArnObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AwsLambdaArn is the Schema for the AwsLambdaArns API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadogjet}
type AwsLambdaArn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLambdaArnSpec   `json:"spec"`
	Status            AwsLambdaArnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsLambdaArnList contains a list of AwsLambdaArns
type AwsLambdaArnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwsLambdaArn `json:"items"`
}

// Repository type metadata.
var (
	AwsLambdaArn_Kind             = "AwsLambdaArn"
	AwsLambdaArn_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AwsLambdaArn_Kind}.String()
	AwsLambdaArn_KindAPIVersion   = AwsLambdaArn_Kind + "." + CRDGroupVersion.String()
	AwsLambdaArn_GroupVersionKind = CRDGroupVersion.WithKind(AwsLambdaArn_Kind)
)

func init() {
	SchemeBuilder.Register(&AwsLambdaArn{}, &AwsLambdaArnList{})
}
