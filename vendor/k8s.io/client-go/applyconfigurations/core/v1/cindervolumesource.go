/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CinderVolumeSourceApplyConfiguration represents an declarative configuration of the CinderVolumeSource type for use
// with apply.
type CinderVolumeSourceApplyConfiguration struct {
	VolumeID  *string                                 `json:"volumeID,omitempty"`
	FSType    *string                                 `json:"fsType,omitempty"`
	ReadOnly  *bool                                   `json:"readOnly,omitempty"`
	SecretRef *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// CinderVolumeSourceApplyConfiguration constructs an declarative configuration of the CinderVolumeSource type for use with
// apply.
func CinderVolumeSource() *CinderVolumeSourceApplyConfiguration {
	return &CinderVolumeSourceApplyConfiguration{}
}

// WithVolumeID sets the VolumeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeID field is set to the value of the last call.
func (b *CinderVolumeSourceApplyConfiguration) WithVolumeID(value string) *CinderVolumeSourceApplyConfiguration {
	b.VolumeID = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *CinderVolumeSourceApplyConfiguration) WithFSType(value string) *CinderVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *CinderVolumeSourceApplyConfiguration) WithReadOnly(value bool) *CinderVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *CinderVolumeSourceApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *CinderVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}
