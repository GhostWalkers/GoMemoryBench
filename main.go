package main

import (
	"encoding/binary"
	"github.com/ghetzel/shmtool/shm"
	"math"
	"unsafe"
)

func WriteFloat64Bench() {
	segment, err := shm.Create(1024)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(0.42))
	segment.Write(b)
	/*if segment, err := shm.Create(1024 * 1024 * 28); err == nil {
		// Mark the segment for destruction when the program exits
		//
		// NOTE: The memory segment will only be destroyed when all processes that have attached
		//       to it have detached, which can happen explicitly (here via segment.Detach(ptr)),
		//       or implicitly when the process exits.
		//
		// NOTE: Memory is not overwritten / zeroed out when destroyed.  If you have sensitive data in
		//       this memory segment, you must overwrite it yourself before detaching.
		//
		defer segment.Destroy()

		// Call the Attach() function on the created segment to get the memory address
		// where data can be read or written.  You must do this even if you don't use the address
		// directly as this is what makes the shared memory segment "part of" this process's memory
		// space, and thus allowing you to read from/write to it.
		if segmentAddress, err := segment.Attach(); err == nil {
			defer segment.Detach(segmentAddress)

			// Write the contents of standard input to the shared memory area.
			if _, err := io.Copy(segment, os.Stdin); err != nil {
				panic(err.Error())
			}

			// Do something, maybe tell another process to start and read from this segment (which
			// is communicated by giving the other process the address in segmentAddress).
			//

			// Read the contents of the shared memory area, which may (or may not) have been modified by
			// another program.
			if _, err := io.Copy(os.Stdout, segment); err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	*/
}

func WriteInt64Bench() {
	buffer := make([]byte, 1024)
	ptr := unsafe.Pointer(&buffer[0])

	*(*int64)(unsafe.Pointer(uintptr(ptr))) = 0xDEAD_BEEF
}

func WriteRawBench() {
	buffer := make([]byte, 1024)
	ptr := unsafe.Pointer(&buffer[0])

	*(*string)(unsafe.Pointer(uintptr(ptr))) = "ABCDEFGH"
}
