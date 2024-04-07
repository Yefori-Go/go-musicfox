// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package enumeration

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureDeviceInformation string = "rc(Windows.Devices.Enumeration.DeviceInformation;{aba0fb95-4398-489d-8e44-e6130927011f})"

type DeviceInformation struct {
	ole.IUnknown
}

const GUIDiDeviceInformation string = "aba0fb95-4398-489d-8e44-e6130927011f"
const SignatureiDeviceInformation string = "{aba0fb95-4398-489d-8e44-e6130927011f}"

type iDeviceInformation struct {
	ole.IInspectable
}

type iDeviceInformationVtbl struct {
	ole.IInspectableVtbl

	GetId                  uintptr
	GetName                uintptr
	GetIsEnabled           uintptr
	GetIsDefault           uintptr
	GetEnclosureLocation   uintptr
	GetProperties          uintptr
	Update                 uintptr
	GetThumbnailAsync      uintptr
	GetGlyphThumbnailAsync uintptr
}

func (v *iDeviceInformation) VTable() *iDeviceInformationVtbl {
	return (*iDeviceInformationVtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiDeviceInformation2 string = "f156a638-7997-48d9-a10c-269d46533f48"
const SignatureiDeviceInformation2 string = "{f156a638-7997-48d9-a10c-269d46533f48}"

type iDeviceInformation2 struct {
	ole.IInspectable
}

type iDeviceInformation2Vtbl struct {
	ole.IInspectableVtbl

	GetKind    uintptr
	GetPairing uintptr
}

func (v *iDeviceInformation2) VTable() *iDeviceInformation2Vtbl {
	return (*iDeviceInformation2Vtbl)(unsafe.Pointer(v.RawVTable))
}
