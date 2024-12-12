package v1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in WebUi) DeepCopyInto(out *WebUi) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = WebUiSpec{
		Replicas: in.Spec.Replicas,
		Image:    in.Spec.Image,
		Contents: in.Spec.Contents,
	}
}

func (in WebUi) DeepCopyObject() runtime.Object {
	out := WebUi{}
	in.DeepCopyInto(&out)

	return &out
}

func (in WebUiList) DeepCopyObject() runtime.Object {
	out := WebUiList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]WebUi, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
