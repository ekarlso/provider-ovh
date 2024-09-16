//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalInformationInitParameters) DeepCopyInto(out *ExternalInformationInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalInformationInitParameters.
func (in *ExternalInformationInitParameters) DeepCopy() *ExternalInformationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalInformationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalInformationObservation) DeepCopyInto(out *ExternalInformationObservation) {
	*out = *in
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]IpsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalInformationObservation.
func (in *ExternalInformationObservation) DeepCopy() *ExternalInformationObservation {
	if in == nil {
		return nil
	}
	out := new(ExternalInformationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalInformationParameters) DeepCopyInto(out *ExternalInformationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalInformationParameters.
func (in *ExternalInformationParameters) DeepCopy() *ExternalInformationParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalInformationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfacesInitParameters) DeepCopyInto(out *InterfacesInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfacesInitParameters.
func (in *InterfacesInitParameters) DeepCopy() *InterfacesInitParameters {
	if in == nil {
		return nil
	}
	out := new(InterfacesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfacesObservation) DeepCopyInto(out *InterfacesObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfacesObservation.
func (in *InterfacesObservation) DeepCopy() *InterfacesObservation {
	if in == nil {
		return nil
	}
	out := new(InterfacesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfacesParameters) DeepCopyInto(out *InterfacesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfacesParameters.
func (in *InterfacesParameters) DeepCopy() *InterfacesParameters {
	if in == nil {
		return nil
	}
	out := new(InterfacesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsInitParameters) DeepCopyInto(out *IpsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsInitParameters.
func (in *IpsInitParameters) DeepCopy() *IpsInitParameters {
	if in == nil {
		return nil
	}
	out := new(IpsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsObservation) DeepCopyInto(out *IpsObservation) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsObservation.
func (in *IpsObservation) DeepCopy() *IpsObservation {
	if in == nil {
		return nil
	}
	out := new(IpsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IpsParameters) DeepCopyInto(out *IpsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IpsParameters.
func (in *IpsParameters) DeepCopy() *IpsParameters {
	if in == nil {
		return nil
	}
	out := new(IpsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGateway) DeepCopyInto(out *ProjectGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGateway.
func (in *ProjectGateway) DeepCopy() *ProjectGateway {
	if in == nil {
		return nil
	}
	out := new(ProjectGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGatewayInitParameters) DeepCopyInto(out *ProjectGatewayInitParameters) {
	*out = *in
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.NetworkIDRef != nil {
		in, out := &in.NetworkIDRef, &out.NetworkIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkIDSelector != nil {
		in, out := &in.NetworkIDSelector, &out.NetworkIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGatewayInitParameters.
func (in *ProjectGatewayInitParameters) DeepCopy() *ProjectGatewayInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectGatewayInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGatewayList) DeepCopyInto(out *ProjectGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGatewayList.
func (in *ProjectGatewayList) DeepCopy() *ProjectGatewayList {
	if in == nil {
		return nil
	}
	out := new(ProjectGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGatewayObservation) DeepCopyInto(out *ProjectGatewayObservation) {
	*out = *in
	if in.ExternalInformation != nil {
		in, out := &in.ExternalInformation, &out.ExternalInformation
		*out = make([]ExternalInformationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]InterfacesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGatewayObservation.
func (in *ProjectGatewayObservation) DeepCopy() *ProjectGatewayObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectGatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGatewayParameters) DeepCopyInto(out *ProjectGatewayParameters) {
	*out = *in
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.NetworkIDRef != nil {
		in, out := &in.NetworkIDRef, &out.NetworkIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkIDSelector != nil {
		in, out := &in.NetworkIDSelector, &out.NetworkIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGatewayParameters.
func (in *ProjectGatewayParameters) DeepCopy() *ProjectGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGatewaySpec) DeepCopyInto(out *ProjectGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGatewaySpec.
func (in *ProjectGatewaySpec) DeepCopy() *ProjectGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(ProjectGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectGatewayStatus) DeepCopyInto(out *ProjectGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectGatewayStatus.
func (in *ProjectGatewayStatus) DeepCopy() *ProjectGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectGatewayStatus)
	in.DeepCopyInto(out)
	return out
}