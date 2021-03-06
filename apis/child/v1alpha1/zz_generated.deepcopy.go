//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyObservation) DeepCopyInto(out *APIKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyObservation.
func (in *APIKeyObservation) DeepCopy() *APIKeyObservation {
	if in == nil {
		return nil
	}
	out := new(APIKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyParameters) DeepCopyInto(out *APIKeyParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyParameters.
func (in *APIKeyParameters) DeepCopy() *APIKeyParameters {
	if in == nil {
		return nil
	}
	out := new(APIKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationKeyObservation) DeepCopyInto(out *ApplicationKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationKeyObservation.
func (in *ApplicationKeyObservation) DeepCopy() *ApplicationKeyObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationKeyParameters) DeepCopyInto(out *ApplicationKeyParameters) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationKeyParameters.
func (in *ApplicationKeyParameters) DeepCopy() *ApplicationKeyParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Organization) DeepCopyInto(out *Organization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Organization.
func (in *Organization) DeepCopy() *Organization {
	if in == nil {
		return nil
	}
	out := new(Organization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Organization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationList) DeepCopyInto(out *OrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Organization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationList.
func (in *OrganizationList) DeepCopy() *OrganizationList {
	if in == nil {
		return nil
	}
	out := new(OrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationObservation) DeepCopyInto(out *OrganizationObservation) {
	*out = *in
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = make([]APIKeyObservation, len(*in))
		copy(*out, *in)
	}
	if in.ApplicationKey != nil {
		in, out := &in.ApplicationKey, &out.ApplicationKey
		*out = make([]ApplicationKeyObservation, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PublicID != nil {
		in, out := &in.PublicID, &out.PublicID
		*out = new(string)
		**out = **in
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make([]SettingsObservation, len(*in))
		copy(*out, *in)
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = make([]UserObservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationObservation.
func (in *OrganizationObservation) DeepCopy() *OrganizationObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationParameters) DeepCopyInto(out *OrganizationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationParameters.
func (in *OrganizationParameters) DeepCopy() *OrganizationParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpec) DeepCopyInto(out *OrganizationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpec.
func (in *OrganizationSpec) DeepCopy() *OrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationStatus) DeepCopyInto(out *OrganizationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationStatus.
func (in *OrganizationStatus) DeepCopy() *OrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLAutocreateUsersDomainsObservation) DeepCopyInto(out *SAMLAutocreateUsersDomainsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLAutocreateUsersDomainsObservation.
func (in *SAMLAutocreateUsersDomainsObservation) DeepCopy() *SAMLAutocreateUsersDomainsObservation {
	if in == nil {
		return nil
	}
	out := new(SAMLAutocreateUsersDomainsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLAutocreateUsersDomainsParameters) DeepCopyInto(out *SAMLAutocreateUsersDomainsParameters) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLAutocreateUsersDomainsParameters.
func (in *SAMLAutocreateUsersDomainsParameters) DeepCopy() *SAMLAutocreateUsersDomainsParameters {
	if in == nil {
		return nil
	}
	out := new(SAMLAutocreateUsersDomainsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLIdpInitiatedLoginObservation) DeepCopyInto(out *SAMLIdpInitiatedLoginObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLIdpInitiatedLoginObservation.
func (in *SAMLIdpInitiatedLoginObservation) DeepCopy() *SAMLIdpInitiatedLoginObservation {
	if in == nil {
		return nil
	}
	out := new(SAMLIdpInitiatedLoginObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLIdpInitiatedLoginParameters) DeepCopyInto(out *SAMLIdpInitiatedLoginParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLIdpInitiatedLoginParameters.
func (in *SAMLIdpInitiatedLoginParameters) DeepCopy() *SAMLIdpInitiatedLoginParameters {
	if in == nil {
		return nil
	}
	out := new(SAMLIdpInitiatedLoginParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLObservation) DeepCopyInto(out *SAMLObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLObservation.
func (in *SAMLObservation) DeepCopy() *SAMLObservation {
	if in == nil {
		return nil
	}
	out := new(SAMLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLParameters) DeepCopyInto(out *SAMLParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLParameters.
func (in *SAMLParameters) DeepCopy() *SAMLParameters {
	if in == nil {
		return nil
	}
	out := new(SAMLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLStrictModeObservation) DeepCopyInto(out *SAMLStrictModeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLStrictModeObservation.
func (in *SAMLStrictModeObservation) DeepCopy() *SAMLStrictModeObservation {
	if in == nil {
		return nil
	}
	out := new(SAMLStrictModeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLStrictModeParameters) DeepCopyInto(out *SAMLStrictModeParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLStrictModeParameters.
func (in *SAMLStrictModeParameters) DeepCopy() *SAMLStrictModeParameters {
	if in == nil {
		return nil
	}
	out := new(SAMLStrictModeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsObservation) DeepCopyInto(out *SettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsObservation.
func (in *SettingsObservation) DeepCopy() *SettingsObservation {
	if in == nil {
		return nil
	}
	out := new(SettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingsParameters) DeepCopyInto(out *SettingsParameters) {
	*out = *in
	if in.PrivateWidgetShare != nil {
		in, out := &in.PrivateWidgetShare, &out.PrivateWidgetShare
		*out = new(bool)
		**out = **in
	}
	if in.SAML != nil {
		in, out := &in.SAML, &out.SAML
		*out = make([]SAMLParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SAMLAutocreateAccessRole != nil {
		in, out := &in.SAMLAutocreateAccessRole, &out.SAMLAutocreateAccessRole
		*out = new(string)
		**out = **in
	}
	if in.SAMLAutocreateUsersDomains != nil {
		in, out := &in.SAMLAutocreateUsersDomains, &out.SAMLAutocreateUsersDomains
		*out = make([]SAMLAutocreateUsersDomainsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SAMLCanBeEnabled != nil {
		in, out := &in.SAMLCanBeEnabled, &out.SAMLCanBeEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SAMLIdpEndpoint != nil {
		in, out := &in.SAMLIdpEndpoint, &out.SAMLIdpEndpoint
		*out = new(string)
		**out = **in
	}
	if in.SAMLIdpInitiatedLogin != nil {
		in, out := &in.SAMLIdpInitiatedLogin, &out.SAMLIdpInitiatedLogin
		*out = make([]SAMLIdpInitiatedLoginParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SAMLIdpMetadataUploaded != nil {
		in, out := &in.SAMLIdpMetadataUploaded, &out.SAMLIdpMetadataUploaded
		*out = new(bool)
		**out = **in
	}
	if in.SAMLLoginURL != nil {
		in, out := &in.SAMLLoginURL, &out.SAMLLoginURL
		*out = new(string)
		**out = **in
	}
	if in.SAMLStrictMode != nil {
		in, out := &in.SAMLStrictMode, &out.SAMLStrictMode
		*out = make([]SAMLStrictModeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsParameters.
func (in *SettingsParameters) DeepCopy() *SettingsParameters {
	if in == nil {
		return nil
	}
	out := new(SettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserObservation) DeepCopyInto(out *UserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserObservation.
func (in *UserObservation) DeepCopy() *UserObservation {
	if in == nil {
		return nil
	}
	out := new(UserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserParameters) DeepCopyInto(out *UserParameters) {
	*out = *in
	if in.AccessRole != nil {
		in, out := &in.AccessRole, &out.AccessRole
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserParameters.
func (in *UserParameters) DeepCopy() *UserParameters {
	if in == nil {
		return nil
	}
	out := new(UserParameters)
	in.DeepCopyInto(out)
	return out
}
