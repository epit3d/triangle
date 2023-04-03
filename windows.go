//go:build windows

package triangle

import "unsafe"
import "C"

// The allocations here are using C.malloc and C.free over trimalloc and
// trifree, as C.malloc works the same but gives a better error report in cases
// if we go out of memory.

func trimalloc(size C.ulonglong) unsafe.Pointer {
	return C.malloc(size)
}

func intSliceToCArr(slice []int32) (*C.int, C.int) {
	sz := len(slice)
	ptr := (*C.int)(trimalloc(C.sizeof_int * (C.ulonglong)(sz)))
	cArr := (*[1 << 30]C.int)(unsafe.Pointer(ptr))[:sz:sz]
	for i := range slice {
		cArr[i] = C.int(slice[i])
	}
	return ptr, C.int(len(slice))
}

func intSlice2DToCArr(slice [][2]int32) (*C.int, C.int) {
	sz := 2 * len(slice)
	ptr := (*C.int)(trimalloc(C.sizeof_int * (C.ulonglong)(sz)))
	cArr := (*[1 << 30]C.int)(unsafe.Pointer(ptr))[:sz:sz]
	for i := range slice {
		j := 2 * i
		cArr[j] = C.int(slice[i][0])
		cArr[j+1] = C.int(slice[i][1])
	}
	return ptr, C.int(len(slice))
}

func intSlice3DToCArr(slice [][3]int32) (*C.int, C.int) {
	sz := 3 * len(slice)
	ptr := (*C.int)(trimalloc(C.sizeof_int * (C.ulonglong)(sz)))
	cArr := (*[1 << 30]C.int)(unsafe.Pointer(ptr))[:sz:sz]
	for i := range slice {
		j := 3 * i
		cArr[j] = C.int(slice[i][0])
		cArr[j+1] = C.int(slice[i][1])
		cArr[j+2] = C.int(slice[i][2])
	}
	return ptr, C.int(len(slice))
}

func flt64SliceToCArr(slice []float64) (*C.double, C.int) {
	sz := len(slice)
	ptr := (*C.double)(trimalloc(C.sizeof_double * (C.ulonglong)(sz)))
	cArr := (*[1 << 30]C.double)(unsafe.Pointer(ptr))[:sz:sz]
	for i := range slice {
		cArr[i] = C.double(slice[i])
	}
	return ptr, C.int(len(slice))
}

func flt64Slice2DToCArr(slice [][2]float64) (*C.double, C.int) {
	sz := 2 * len(slice)
	ptr := (*C.double)(trimalloc(C.sizeof_double * (C.ulonglong)(sz)))
	cArr := (*[1 << 30]C.double)(unsafe.Pointer(ptr))[:sz:sz]
	for i := range slice {
		j := i * 2
		cArr[j] = C.double(slice[i][0])
		cArr[j+1] = C.double(slice[i][1])
	}
	return ptr, C.int(len(slice))
}
