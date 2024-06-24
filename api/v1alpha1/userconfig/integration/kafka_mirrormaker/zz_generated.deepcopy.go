//go:build !ignore_autogenerated

// Copyright (c) 2024 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package kafkamirrormakeruserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaMirrormaker) DeepCopyInto(out *KafkaMirrormaker) {
	*out = *in
	if in.ConsumerAutoOffsetReset != nil {
		in, out := &in.ConsumerAutoOffsetReset, &out.ConsumerAutoOffsetReset
		*out = new(string)
		**out = **in
	}
	if in.ConsumerFetchMinBytes != nil {
		in, out := &in.ConsumerFetchMinBytes, &out.ConsumerFetchMinBytes
		*out = new(int)
		**out = **in
	}
	if in.ConsumerMaxPollRecords != nil {
		in, out := &in.ConsumerMaxPollRecords, &out.ConsumerMaxPollRecords
		*out = new(int)
		**out = **in
	}
	if in.ProducerBatchSize != nil {
		in, out := &in.ProducerBatchSize, &out.ProducerBatchSize
		*out = new(int)
		**out = **in
	}
	if in.ProducerBufferMemory != nil {
		in, out := &in.ProducerBufferMemory, &out.ProducerBufferMemory
		*out = new(int)
		**out = **in
	}
	if in.ProducerCompressionType != nil {
		in, out := &in.ProducerCompressionType, &out.ProducerCompressionType
		*out = new(string)
		**out = **in
	}
	if in.ProducerLingerMs != nil {
		in, out := &in.ProducerLingerMs, &out.ProducerLingerMs
		*out = new(int)
		**out = **in
	}
	if in.ProducerMaxRequestSize != nil {
		in, out := &in.ProducerMaxRequestSize, &out.ProducerMaxRequestSize
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaMirrormaker.
func (in *KafkaMirrormaker) DeepCopy() *KafkaMirrormaker {
	if in == nil {
		return nil
	}
	out := new(KafkaMirrormaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaMirrormakerUserConfig) DeepCopyInto(out *KafkaMirrormakerUserConfig) {
	*out = *in
	if in.ClusterAlias != nil {
		in, out := &in.ClusterAlias, &out.ClusterAlias
		*out = new(string)
		**out = **in
	}
	if in.KafkaMirrormaker != nil {
		in, out := &in.KafkaMirrormaker, &out.KafkaMirrormaker
		*out = new(KafkaMirrormaker)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaMirrormakerUserConfig.
func (in *KafkaMirrormakerUserConfig) DeepCopy() *KafkaMirrormakerUserConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaMirrormakerUserConfig)
	in.DeepCopyInto(out)
	return out
}
