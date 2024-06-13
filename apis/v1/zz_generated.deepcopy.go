//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
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
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectBucket) DeepCopyInto(out *ObjectBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectBucket.
func (in *ObjectBucket) DeepCopy() *ObjectBucket {
	if in == nil {
		return nil
	}
	out := new(ObjectBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectBucketParameters) DeepCopyInto(out *ObjectBucketParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectBucketParameters.
func (in *ObjectBucketParameters) DeepCopy() *ObjectBucketParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectBucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectBucketSpec) DeepCopyInto(out *ObjectBucketSpec) {
	*out = *in
	out.Parameters = in.Parameters
	out.WriteConnectionSecretToRef = in.WriteConnectionSecretToRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectBucketSpec.
func (in *ObjectBucketSpec) DeepCopy() *ObjectBucketSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectBucketStatus) DeepCopyInto(out *ObjectBucketStatus) {
	*out = *in
	if in.AccessUserConditions != nil {
		in, out := &in.AccessUserConditions, &out.AccessUserConditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BucketConditions != nil {
		in, out := &in.BucketConditions, &out.BucketConditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectBucketStatus.
func (in *ObjectBucketStatus) DeepCopy() *ObjectBucketStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XObjectBucket) DeepCopyInto(out *XObjectBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XObjectBucket.
func (in *XObjectBucket) DeepCopy() *XObjectBucket {
	if in == nil {
		return nil
	}
	out := new(XObjectBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XObjectBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XObjectBucketSpec) DeepCopyInto(out *XObjectBucketSpec) {
	*out = *in
	out.Parameters = in.Parameters
	out.WriteConnectionSecretToRef = in.WriteConnectionSecretToRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XObjectBucketSpec.
func (in *XObjectBucketSpec) DeepCopy() *XObjectBucketSpec {
	if in == nil {
		return nil
	}
	out := new(XObjectBucketSpec)
	in.DeepCopyInto(out)
	return out
}
