package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type KafkaTopicSpec struct {
	Name string `json:"name"`
}

type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KafkaTopicSpec `json:"spec"`
}

type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KafkaTopic `json:"items"`
}

func (in *KafkaTopic) DeepCopyInto(out *KafkaTopic) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

func (in *KafkaTopic) DeepCopy() *KafkaTopic {
	if in == nil {
		return nil
	}
	out := new(KafkaTopic)
	in.DeepCopyInto(out)
	return out
}

func (in *KafkaTopic) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

func (in *KafkaTopicList) DeepCopyInto(out *KafkaTopicList) {
	*out = *in
	in.ListMeta.DeepCopyInto(&out.ListMeta)
}

func (in *KafkaTopicList) DeepCopy() *KafkaTopicList {
	if in == nil {
		return nil
	}
	out := new(KafkaTopicList)
	in.DeepCopyInto(out)
	return out
}

func (in *KafkaTopicList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}
