// +build !ignore_autogenerated

/*
Copyright 2021 The idx Authors.

http://git.inspur.com/middleware/idx-component

### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataTopicAndThread) DeepCopyInto(out *DataTopicAndThread) {
	*out = *in
	if in.ThreadNum != nil {
		in, out := &in.ThreadNum, &out.ThreadNum
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataTopicAndThread.
func (in *DataTopicAndThread) DeepCopy() *DataTopicAndThread {
	if in == nil {
		return nil
	}
	out := new(DataTopicAndThread)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSync) DeepCopyInto(out *FileSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSync.
func (in *FileSync) DeepCopy() *FileSync {
	if in == nil {
		return nil
	}
	out := new(FileSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSyncList) DeepCopyInto(out *FileSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSyncList.
func (in *FileSyncList) DeepCopy() *FileSyncList {
	if in == nil {
		return nil
	}
	out := new(FileSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSyncSpec) DeepCopyInto(out *FileSyncSpec) {
	*out = *in
	out.Global = in.Global
	if in.NATS != nil {
		in, out := &in.NATS, &out.NATS
		*out = new(NATSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Monit != nil {
		in, out := &in.Monit, &out.Monit
		*out = new(MonitSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reader != nil {
		in, out := &in.Reader, &out.Reader
		*out = new(ReaderSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Writer != nil {
		in, out := &in.Writer, &out.Writer
		*out = new(WriterSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSyncSpec.
func (in *FileSyncSpec) DeepCopy() *FileSyncSpec {
	if in == nil {
		return nil
	}
	out := new(FileSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSyncStatus) DeepCopyInto(out *FileSyncStatus) {
	*out = *in
	if in.ReaderStatus != nil {
		in, out := &in.ReaderStatus, &out.ReaderStatus
		*out = new(string)
		**out = **in
	}
	if in.WriterStatus != nil {
		in, out := &in.WriterStatus, &out.WriterStatus
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSyncStatus.
func (in *FileSyncStatus) DeepCopy() *FileSyncStatus {
	if in == nil {
		return nil
	}
	out := new(FileSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalSpec) DeepCopyInto(out *GlobalSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalSpec.
func (in *GlobalSpec) DeepCopy() *GlobalSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSpec) DeepCopyInto(out *LogSpec) {
	*out = *in
	if in.LogPath != nil {
		in, out := &in.LogPath, &out.LogPath
		*out = new(string)
		**out = **in
	}
	if in.LogPathPVC != nil {
		in, out := &in.LogPathPVC, &out.LogPathPVC
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSpec.
func (in *LogSpec) DeepCopy() *LogSpec {
	if in == nil {
		return nil
	}
	out := new(LogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitSpec) DeepCopyInto(out *MonitSpec) {
	*out = *in
	if in.EndMsgTopic != nil {
		in, out := &in.EndMsgTopic, &out.EndMsgTopic
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitSpec.
func (in *MonitSpec) DeepCopy() *MonitSpec {
	if in == nil {
		return nil
	}
	out := new(MonitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NATSSpec) DeepCopyInto(out *NATSSpec) {
	*out = *in
	if in.CredsSecret != nil {
		in, out := &in.CredsSecret, &out.CredsSecret
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NATSSpec.
func (in *NATSSpec) DeepCopy() *NATSSpec {
	if in == nil {
		return nil
	}
	out := new(NATSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReaderSpec) DeepCopyInto(out *ReaderSpec) {
	*out = *in
	in.SyncMetadata.DeepCopyInto(&out.SyncMetadata)
	in.LogSpec.DeepCopyInto(&out.LogSpec)
	in.DataTopicAndThread.DeepCopyInto(&out.DataTopicAndThread)
	if in.SrcPathPVC != nil {
		in, out := &in.SrcPathPVC, &out.SrcPathPVC
		*out = new(string)
		**out = **in
	}
	if in.ShadowPath != nil {
		in, out := &in.ShadowPath, &out.ShadowPath
		*out = new(string)
		**out = **in
	}
	if in.ShadowPathPVC != nil {
		in, out := &in.ShadowPathPVC, &out.ShadowPathPVC
		*out = new(string)
		**out = **in
	}
	if in.DirShadowPrefix != nil {
		in, out := &in.DirShadowPrefix, &out.DirShadowPrefix
		*out = new(string)
		**out = **in
	}
	if in.HandleModeAfterRead != nil {
		in, out := &in.HandleModeAfterRead, &out.HandleModeAfterRead
		*out = new(string)
		**out = **in
	}
	if in.TrashPath != nil {
		in, out := &in.TrashPath, &out.TrashPath
		*out = new(string)
		**out = **in
	}
	if in.TrashPathPVC != nil {
		in, out := &in.TrashPathPVC, &out.TrashPathPVC
		*out = new(string)
		**out = **in
	}
	if in.IncScanMode != nil {
		in, out := &in.IncScanMode, &out.IncScanMode
		*out = new(int)
		**out = **in
	}
	if in.IncScanInterval != nil {
		in, out := &in.IncScanInterval, &out.IncScanInterval
		*out = new(int)
		**out = **in
	}
	if in.IncQnotifyMode != nil {
		in, out := &in.IncQnotifyMode, &out.IncQnotifyMode
		*out = new(int)
		**out = **in
	}
	if in.DirLevelUseThread != nil {
		in, out := &in.DirLevelUseThread, &out.DirLevelUseThread
		*out = new(int)
		**out = **in
	}
	if in.IncSkipDays != nil {
		in, out := &in.IncSkipDays, &out.IncSkipDays
		*out = new(int)
		**out = **in
	}
	if in.IncludedFiletype != nil {
		in, out := &in.IncludedFiletype, &out.IncludedFiletype
		*out = new(string)
		**out = **in
	}
	if in.ExcludedFiletype != nil {
		in, out := &in.ExcludedFiletype, &out.ExcludedFiletype
		*out = new(string)
		**out = **in
	}
	if in.FilepathRegex != nil {
		in, out := &in.FilepathRegex, &out.FilepathRegex
		*out = new(string)
		**out = **in
	}
	if in.ReaderLabel != nil {
		in, out := &in.ReaderLabel, &out.ReaderLabel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReaderSpec.
func (in *ReaderSpec) DeepCopy() *ReaderSpec {
	if in == nil {
		return nil
	}
	out := new(ReaderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncMetadata) DeepCopyInto(out *SyncMetadata) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncMetadata.
func (in *SyncMetadata) DeepCopy() *SyncMetadata {
	if in == nil {
		return nil
	}
	out := new(SyncMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WriterSpec) DeepCopyInto(out *WriterSpec) {
	*out = *in
	in.SyncMetadata.DeepCopyInto(&out.SyncMetadata)
	in.LogSpec.DeepCopyInto(&out.LogSpec)
	in.DataTopicAndThread.DeepCopyInto(&out.DataTopicAndThread)
	if in.RcvPathPVC != nil {
		in, out := &in.RcvPathPVC, &out.RcvPathPVC
		*out = new(string)
		**out = **in
	}
	if in.TrashPath != nil {
		in, out := &in.TrashPath, &out.TrashPath
		*out = new(string)
		**out = **in
	}
	if in.TrashPathPVC != nil {
		in, out := &in.TrashPathPVC, &out.TrashPathPVC
		*out = new(string)
		**out = **in
	}
	if in.BisyncShadowPath != nil {
		in, out := &in.BisyncShadowPath, &out.BisyncShadowPath
		*out = new(string)
		**out = **in
	}
	if in.BisyncMode != nil {
		in, out := &in.BisyncMode, &out.BisyncMode
		*out = new(int)
		**out = **in
	}
	if in.BigfileSliceCount != nil {
		in, out := &in.BigfileSliceCount, &out.BigfileSliceCount
		*out = new(int)
		**out = **in
	}
	if in.WriterLabel != nil {
		in, out := &in.WriterLabel, &out.WriterLabel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WriterSpec.
func (in *WriterSpec) DeepCopy() *WriterSpec {
	if in == nil {
		return nil
	}
	out := new(WriterSpec)
	in.DeepCopyInto(out)
	return out
}