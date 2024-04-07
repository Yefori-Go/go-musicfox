// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package core

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
)

const SignatureMseSourceBufferList string = "rc(Windows.Media.Core.MseSourceBufferList;{95fae8e7-a8e7-4ebf-8927-145e940ba511})"

type MseSourceBufferList struct {
	ole.IUnknown
}

func (impl *MseSourceBufferList) AddSourceBufferAdded(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseSourceBufferList))
	defer itf.Release()
	v := (*iMseSourceBufferList)(unsafe.Pointer(itf))
	return v.AddSourceBufferAdded(handler)
}

func (impl *MseSourceBufferList) RemoveSourceBufferAdded(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseSourceBufferList))
	defer itf.Release()
	v := (*iMseSourceBufferList)(unsafe.Pointer(itf))
	return v.RemoveSourceBufferAdded(token)
}

func (impl *MseSourceBufferList) AddSourceBufferRemoved(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseSourceBufferList))
	defer itf.Release()
	v := (*iMseSourceBufferList)(unsafe.Pointer(itf))
	return v.AddSourceBufferRemoved(handler)
}

func (impl *MseSourceBufferList) RemoveSourceBufferRemoved(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseSourceBufferList))
	defer itf.Release()
	v := (*iMseSourceBufferList)(unsafe.Pointer(itf))
	return v.RemoveSourceBufferRemoved(token)
}

func (impl *MseSourceBufferList) GetBuffers() (*collections.IVectorView, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMseSourceBufferList))
	defer itf.Release()
	v := (*iMseSourceBufferList)(unsafe.Pointer(itf))
	return v.GetBuffers()
}

const GUIDiMseSourceBufferList string = "95fae8e7-a8e7-4ebf-8927-145e940ba511"
const SignatureiMseSourceBufferList string = "{95fae8e7-a8e7-4ebf-8927-145e940ba511}"

type iMseSourceBufferList struct {
	ole.IInspectable
}

type iMseSourceBufferListVtbl struct {
	ole.IInspectableVtbl

	AddSourceBufferAdded      uintptr
	RemoveSourceBufferAdded   uintptr
	AddSourceBufferRemoved    uintptr
	RemoveSourceBufferRemoved uintptr
	GetBuffers                uintptr
}

func (v *iMseSourceBufferList) VTable() *iMseSourceBufferListVtbl {
	return (*iMseSourceBufferListVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMseSourceBufferList) AddSourceBufferAdded(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddSourceBufferAdded,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseSourceBufferList) RemoveSourceBufferAdded(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveSourceBufferAdded,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseSourceBufferList) AddSourceBufferRemoved(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddSourceBufferRemoved,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMseSourceBufferList) RemoveSourceBufferRemoved(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveSourceBufferRemoved,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMseSourceBufferList) GetBuffers() (*collections.IVectorView, error) {
	var out *collections.IVectorView
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetBuffers,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVectorView
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
