// +build !ignore_autogenerated

//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoDeploymentReplication) DeepCopyInto(out *ArangoDeploymentReplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoDeploymentReplication.
func (in *ArangoDeploymentReplication) DeepCopy() *ArangoDeploymentReplication {
	if in == nil {
		return nil
	}
	out := new(ArangoDeploymentReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoDeploymentReplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArangoDeploymentReplicationList) DeepCopyInto(out *ArangoDeploymentReplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ArangoDeploymentReplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArangoDeploymentReplicationList.
func (in *ArangoDeploymentReplicationList) DeepCopy() *ArangoDeploymentReplicationList {
	if in == nil {
		return nil
	}
	out := new(ArangoDeploymentReplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArangoDeploymentReplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentReplicationSpec) DeepCopyInto(out *DeploymentReplicationSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentReplicationSpec.
func (in *DeploymentReplicationSpec) DeepCopy() *DeploymentReplicationSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentReplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentReplicationStatus) DeepCopyInto(out *DeploymentReplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(ConditionList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentReplicationStatus.
func (in *DeploymentReplicationStatus) DeepCopy() *DeploymentReplicationStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentReplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointAuthenticationSpec) DeepCopyInto(out *EndpointAuthenticationSpec) {
	*out = *in
	if in.KeyfileSecretName != nil {
		in, out := &in.KeyfileSecretName, &out.KeyfileSecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.UserSecretName != nil {
		in, out := &in.UserSecretName, &out.UserSecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointAuthenticationSpec.
func (in *EndpointAuthenticationSpec) DeepCopy() *EndpointAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
	if in.DeploymentName != nil {
		in, out := &in.DeploymentName, &out.DeploymentName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.TLS.DeepCopyInto(&out.TLS)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointTLSSpec) DeepCopyInto(out *EndpointTLSSpec) {
	*out = *in
	if in.CASecretName != nil {
		in, out := &in.CASecretName, &out.CASecretName
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointTLSSpec.
func (in *EndpointTLSSpec) DeepCopy() *EndpointTLSSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointTLSSpec)
	in.DeepCopyInto(out)
	return out
}
