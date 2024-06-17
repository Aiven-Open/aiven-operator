//go:build !ignore_autogenerated

// Copyright (c) 2024 Aiven, Helsinki, Finland. https://aiven.io/

// Code generated by controller-gen. DO NOT EDIT.

package rsysloguserconfig

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RsyslogUserConfig) DeepCopyInto(out *RsyslogUserConfig) {
	*out = *in
	if in.Ca != nil {
		in, out := &in.Ca, &out.Ca
		*out = new(string)
		**out = **in
	}
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Logline != nil {
		in, out := &in.Logline, &out.Logline
		*out = new(string)
		**out = **in
	}
	if in.MaxMessageSize != nil {
		in, out := &in.MaxMessageSize, &out.MaxMessageSize
		*out = new(int)
		**out = **in
	}
	if in.Sd != nil {
		in, out := &in.Sd, &out.Sd
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RsyslogUserConfig.
func (in *RsyslogUserConfig) DeepCopy() *RsyslogUserConfig {
	if in == nil {
		return nil
	}
	out := new(RsyslogUserConfig)
	in.DeepCopyInto(out)
	return out
}
