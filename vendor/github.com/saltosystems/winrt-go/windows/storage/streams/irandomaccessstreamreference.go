// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package streams

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const GUIDIRandomAccessStreamReference string = "33ee3134-1dd6-4e3a-8067-d1c162e8642b"
const SignatureIRandomAccessStreamReference string = "{33ee3134-1dd6-4e3a-8067-d1c162e8642b}"

type IRandomAccessStreamReference struct {
	ole.IInspectable
}

type IRandomAccessStreamReferenceVtbl struct {
	ole.IInspectableVtbl

	OpenReadAsync uintptr
}

func (v *IRandomAccessStreamReference) VTable() *IRandomAccessStreamReferenceVtbl {
	return (*IRandomAccessStreamReferenceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IRandomAccessStreamReference) OpenReadAsync() (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().OpenReadAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
