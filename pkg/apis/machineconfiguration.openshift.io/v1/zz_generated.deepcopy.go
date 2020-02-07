// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfig) DeepCopyInto(out *ContainerRuntimeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfig.
func (in *ContainerRuntimeConfig) DeepCopy() *ContainerRuntimeConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerRuntimeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigCondition) DeepCopyInto(out *ContainerRuntimeConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigCondition.
func (in *ContainerRuntimeConfigCondition) DeepCopy() *ContainerRuntimeConfigCondition {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigList) DeepCopyInto(out *ContainerRuntimeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerRuntimeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigList.
func (in *ContainerRuntimeConfigList) DeepCopy() *ContainerRuntimeConfigList {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerRuntimeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigSpec) DeepCopyInto(out *ContainerRuntimeConfigSpec) {
	*out = *in
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerRuntimeConfig != nil {
		in, out := &in.ContainerRuntimeConfig, &out.ContainerRuntimeConfig
		*out = new(ContainerRuntimeConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigSpec.
func (in *ContainerRuntimeConfigSpec) DeepCopy() *ContainerRuntimeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfigStatus) DeepCopyInto(out *ContainerRuntimeConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ContainerRuntimeConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfigStatus.
func (in *ContainerRuntimeConfigStatus) DeepCopy() *ContainerRuntimeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfiguration) DeepCopyInto(out *ContainerRuntimeConfiguration) {
	*out = *in
	out.LogSizeMax = in.LogSizeMax.DeepCopy()
	out.OverlaySize = in.OverlaySize.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfiguration.
func (in *ContainerRuntimeConfiguration) DeepCopy() *ContainerRuntimeConfiguration {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfig) DeepCopyInto(out *ControllerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfig.
func (in *ControllerConfig) DeepCopy() *ControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigList) DeepCopyInto(out *ControllerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigList.
func (in *ControllerConfigList) DeepCopy() *ControllerConfigList {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigSpec) DeepCopyInto(out *ControllerConfigSpec) {
	*out = *in
	if in.KubeAPIServerServingCAData != nil {
		in, out := &in.KubeAPIServerServingCAData, &out.KubeAPIServerServingCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EtcdCAData != nil {
		in, out := &in.EtcdCAData, &out.EtcdCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EtcdMetricCAData != nil {
		in, out := &in.EtcdMetricCAData, &out.EtcdMetricCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RootCAData != nil {
		in, out := &in.RootCAData, &out.RootCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CloudProviderCAData != nil {
		in, out := &in.CloudProviderCAData, &out.CloudProviderCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalTrustBundle != nil {
		in, out := &in.AdditionalTrustBundle, &out.AdditionalTrustBundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(configv1.ProxyStatus)
		**out = **in
	}
	if in.Infra != nil {
		in, out := &in.Infra, &out.Infra
		*out = new(configv1.Infrastructure)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigSpec.
func (in *ControllerConfigSpec) DeepCopy() *ControllerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigStatus) DeepCopyInto(out *ControllerConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ControllerConfigStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigStatus.
func (in *ControllerConfigStatus) DeepCopy() *ControllerConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigStatusCondition) DeepCopyInto(out *ControllerConfigStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigStatusCondition.
func (in *ControllerConfigStatusCondition) DeepCopy() *ControllerConfigStatusCondition {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfig) DeepCopyInto(out *KubeletConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfig.
func (in *KubeletConfig) DeepCopy() *KubeletConfig {
	if in == nil {
		return nil
	}
	out := new(KubeletConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigCondition) DeepCopyInto(out *KubeletConfigCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigCondition.
func (in *KubeletConfigCondition) DeepCopy() *KubeletConfigCondition {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigList) DeepCopyInto(out *KubeletConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeletConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigList.
func (in *KubeletConfigList) DeepCopy() *KubeletConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigSpec) DeepCopyInto(out *KubeletConfigSpec) {
	*out = *in
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeletConfig != nil {
		in, out := &in.KubeletConfig, &out.KubeletConfig
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigSpec.
func (in *KubeletConfigSpec) DeepCopy() *KubeletConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfigStatus) DeepCopyInto(out *KubeletConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KubeletConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfigStatus.
func (in *KubeletConfigStatus) DeepCopy() *KubeletConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeletConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigList) DeepCopyInto(out *MachineConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigList.
func (in *MachineConfigList) DeepCopy() *MachineConfigList {
	if in == nil {
		return nil
	}
	out := new(MachineConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPool) DeepCopyInto(out *MachineConfigPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPool.
func (in *MachineConfigPool) DeepCopy() *MachineConfigPool {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfigPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolCondition) DeepCopyInto(out *MachineConfigPoolCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolCondition.
func (in *MachineConfigPoolCondition) DeepCopy() *MachineConfigPoolCondition {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolList) DeepCopyInto(out *MachineConfigPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfigPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolList.
func (in *MachineConfigPoolList) DeepCopy() *MachineConfigPoolList {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineConfigPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolSpec) DeepCopyInto(out *MachineConfigPoolSpec) {
	*out = *in
	if in.MachineConfigSelector != nil {
		in, out := &in.MachineConfigSelector, &out.MachineConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolSpec.
func (in *MachineConfigPoolSpec) DeepCopy() *MachineConfigPoolSpec {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolStatus) DeepCopyInto(out *MachineConfigPoolStatus) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MachineConfigPoolCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolStatus.
func (in *MachineConfigPoolStatus) DeepCopy() *MachineConfigPoolStatus {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineConfigPoolStatusConfiguration) DeepCopyInto(out *MachineConfigPoolStatusConfiguration) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineConfigPoolStatusConfiguration.
func (in *MachineConfigPoolStatusConfiguration) DeepCopy() *MachineConfigPoolStatusConfiguration {
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolStatusConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenderedMachineConfig) DeepCopyInto(out *RenderedMachineConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenderedMachineConfig.
func (in *RenderedMachineConfig) DeepCopy() *RenderedMachineConfig {
	if in == nil {
		return nil
	}
	out := new(RenderedMachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RenderedMachineConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenderedMachineConfigList) DeepCopyInto(out *RenderedMachineConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RenderedMachineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenderedMachineConfigList.
func (in *RenderedMachineConfigList) DeepCopy() *RenderedMachineConfigList {
	if in == nil {
		return nil
	}
	out := new(RenderedMachineConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RenderedMachineConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
