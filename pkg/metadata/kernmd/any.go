package kernmd

import (
	"github.com/weaveworks/ignite/pkg/metadata"
)

func LoadKernel(id string) (metadata.Metadata, error) {
	md, err := NewKernel(id, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := md.Load(); err != nil {
		return nil, err
	}

	return md, nil
}

func LoadAllKernel() ([]metadata.Metadata, error) {
	return metadata.LoadAllMetadata((&Kernel{}).TypePath(), LoadKernel)
}

func ToKernel(md metadata.Metadata) *Kernel {
	return md.(*Kernel) // This type assert is internal, we don't need to validate it
}

func ToKernelAll(any []metadata.Metadata) []*Kernel {
	var mds []*Kernel

	for _, md := range any {
		mds = append(mds, ToKernel(md))
	}

	return mds
}
