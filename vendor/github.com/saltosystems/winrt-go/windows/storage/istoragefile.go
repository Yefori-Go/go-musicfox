// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package storage

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const GUIDIStorageFile string = "fa3f6186-4214-428c-a64c-14c9ac7315ea"
const SignatureIStorageFile string = "{fa3f6186-4214-428c-a64c-14c9ac7315ea}"

type IStorageFile struct {
	ole.IInspectable
}

type IStorageFileVtbl struct {
	ole.IInspectableVtbl

	GetFileType                       uintptr
	GetContentType                    uintptr
	OpenAsync                         uintptr
	OpenTransactedWriteAsync          uintptr
	CopyOverloadDefaultNameAndOptions uintptr
	CopyOverloadDefaultOptions        uintptr
	CopyOverload                      uintptr
	CopyAndReplaceAsync               uintptr
	MoveOverloadDefaultNameAndOptions uintptr
	MoveOverloadDefaultOptions        uintptr
	MoveOverload                      uintptr
	MoveAndReplaceAsync               uintptr
}

func (v *IStorageFile) VTable() *IStorageFileVtbl {
	return (*IStorageFileVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IStorageFile) GetFileType() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetFileType,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *IStorageFile) GetContentType() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetContentType,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *IStorageFile) OpenAsync(accessMode FileAccessMode) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().OpenAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(accessMode),           // in FileAccessMode
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) OpenTransactedWriteAsync() (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().OpenTransactedWriteAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) CopyOverloadDefaultNameAndOptions(destinationFolder *IStorageFolder) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().CopyOverloadDefaultNameAndOptions,
		uintptr(unsafe.Pointer(v)),                 // this
		uintptr(unsafe.Pointer(destinationFolder)), // in IStorageFolder
		uintptr(unsafe.Pointer(&out)),              // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) CopyOverloadDefaultOptions(destinationFolder *IStorageFolder, desiredNewName string) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	desiredNewNameHStr, err := ole.NewHString(desiredNewName)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().CopyOverloadDefaultOptions,
		uintptr(unsafe.Pointer(v)),                 // this
		uintptr(unsafe.Pointer(destinationFolder)), // in IStorageFolder
		uintptr(desiredNewNameHStr),                // in string
		uintptr(unsafe.Pointer(&out)),              // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) CopyOverload(destinationFolder *IStorageFolder, desiredNewName string, option NameCollisionOption) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	desiredNewNameHStr, err := ole.NewHString(desiredNewName)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().CopyOverload,
		uintptr(unsafe.Pointer(v)),                 // this
		uintptr(unsafe.Pointer(destinationFolder)), // in IStorageFolder
		uintptr(desiredNewNameHStr),                // in string
		uintptr(option),                            // in NameCollisionOption
		uintptr(unsafe.Pointer(&out)),              // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) CopyAndReplaceAsync(fileToReplace *IStorageFile) (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	hr, _, _ := syscall.SyscallN(
		v.VTable().CopyAndReplaceAsync,
		uintptr(unsafe.Pointer(v)),             // this
		uintptr(unsafe.Pointer(fileToReplace)), // in IStorageFile
		uintptr(unsafe.Pointer(&out)),          // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) MoveOverloadDefaultNameAndOptions(destinationFolder *IStorageFolder) (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	hr, _, _ := syscall.SyscallN(
		v.VTable().MoveOverloadDefaultNameAndOptions,
		uintptr(unsafe.Pointer(v)),                 // this
		uintptr(unsafe.Pointer(destinationFolder)), // in IStorageFolder
		uintptr(unsafe.Pointer(&out)),              // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) MoveOverloadDefaultOptions(destinationFolder *IStorageFolder, desiredNewName string) (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	desiredNewNameHStr, err := ole.NewHString(desiredNewName)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().MoveOverloadDefaultOptions,
		uintptr(unsafe.Pointer(v)),                 // this
		uintptr(unsafe.Pointer(destinationFolder)), // in IStorageFolder
		uintptr(desiredNewNameHStr),                // in string
		uintptr(unsafe.Pointer(&out)),              // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) MoveOverload(destinationFolder *IStorageFolder, desiredNewName string, option NameCollisionOption) (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	desiredNewNameHStr, err := ole.NewHString(desiredNewName)
	if err != nil {
		return nil, err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().MoveOverload,
		uintptr(unsafe.Pointer(v)),                 // this
		uintptr(unsafe.Pointer(destinationFolder)), // in IStorageFolder
		uintptr(desiredNewNameHStr),                // in string
		uintptr(option),                            // in NameCollisionOption
		uintptr(unsafe.Pointer(&out)),              // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IStorageFile) MoveAndReplaceAsync(fileToReplace *IStorageFile) (*foundation.IAsyncAction, error) {
	var out *foundation.IAsyncAction
	hr, _, _ := syscall.SyscallN(
		v.VTable().MoveAndReplaceAsync,
		uintptr(unsafe.Pointer(v)),             // this
		uintptr(unsafe.Pointer(fileToReplace)), // in IStorageFile
		uintptr(unsafe.Pointer(&out)),          // out foundation.IAsyncAction
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
