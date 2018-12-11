// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMSpec) DeepCopyInto(out *PMMSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMSpec.
func (in *PMMSpec) DeepCopy() *PMMSpec {
	if in == nil {
		return nil
	}
	out := new(PMMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCBackupSpec) DeepCopyInto(out *PXCBackupSpec) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCBackupSpec.
func (in *PXCBackupSpec) DeepCopy() *PXCBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PXCBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCBackupStatus) DeepCopyInto(out *PXCBackupStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCBackupStatus.
func (in *PXCBackupStatus) DeepCopy() *PXCBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PXCBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PXCBackupVolume) DeepCopyInto(out *PXCBackupVolume) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PXCBackupVolume.
func (in *PXCBackupVolume) DeepCopy() *PXCBackupVolume {
	if in == nil {
		return nil
	}
	out := new(PXCBackupVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBBackup) DeepCopyInto(out *PerconaXtraDBBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBBackup.
func (in *PerconaXtraDBBackup) DeepCopy() *PerconaXtraDBBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBBackupList) DeepCopyInto(out *PerconaXtraDBBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtraDBBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBBackupList.
func (in *PerconaXtraDBBackupList) DeepCopy() *PerconaXtraDBBackupList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBCluster) DeepCopyInto(out *PerconaXtraDBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBCluster.
func (in *PerconaXtraDBCluster) DeepCopy() *PerconaXtraDBCluster {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterList) DeepCopyInto(out *PerconaXtraDBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtraDBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterList.
func (in *PerconaXtraDBClusterList) DeepCopy() *PerconaXtraDBClusterList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtraDBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterSpec) DeepCopyInto(out *PerconaXtraDBClusterSpec) {
	*out = *in
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(Platform)
		**out = **in
	}
	if in.PXC != nil {
		in, out := &in.PXC, &out.PXC
		*out = new(PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ProxySQL != nil {
		in, out := &in.ProxySQL, &out.ProxySQL
		*out = new(PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PMM != nil {
		in, out := &in.PMM, &out.PMM
		*out = new(PMMSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterSpec.
func (in *PerconaXtraDBClusterSpec) DeepCopy() *PerconaXtraDBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtraDBClusterStatus) DeepCopyInto(out *PerconaXtraDBClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtraDBClusterStatus.
func (in *PerconaXtraDBClusterStatus) DeepCopy() *PerconaXtraDBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaXtraDBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResources) DeepCopyInto(out *PodResources) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourcesList)
		**out = **in
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourcesList)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResources.
func (in *PodResources) DeepCopy() *PodResources {
	if in == nil {
		return nil
	}
	out := new(PodResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpec) DeepCopyInto(out *PodSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(PodResources)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeSpec != nil {
		in, out := &in.VolumeSpec, &out.VolumeSpec
		*out = new(PodVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpec.
func (in *PodSpec) DeepCopy() *PodSpec {
	if in == nil {
		return nil
	}
	out := new(PodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeSpec) DeepCopyInto(out *PodVolumeSpec) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeSpec.
func (in *PodVolumeSpec) DeepCopy() *PodVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(PodVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesList) DeepCopyInto(out *ResourcesList) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesList.
func (in *ResourcesList) DeepCopy() *ResourcesList {
	if in == nil {
		return nil
	}
	out := new(ResourcesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerVersion) DeepCopyInto(out *ServerVersion) {
	*out = *in
	out.Info = in.Info
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerVersion.
func (in *ServerVersion) DeepCopy() *ServerVersion {
	if in == nil {
		return nil
	}
	out := new(ServerVersion)
	in.DeepCopyInto(out)
	return out
}
