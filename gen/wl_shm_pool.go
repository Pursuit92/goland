package gen

// The wl_shm_pool object encapsulates a piece of memory shared
// between the compositor and client.  Through the wl_shm_pool
// object, the client can allocate shared memory wl_buffer objects.
// All objects created through the same pool share the same
// underlying mapped memory. Reusing the mapped memory avoids the
// setup/teardown overhead and is useful when interactively resizing
// a surface or for many small buffers.
type WlShmPool interface {
	// Create a wl_buffer object from the pool.
	// The buffer is created offset bytes into the pool and has
	// width and height as specified.  The stride arguments specifies
	// the number of bytes from beginning of one row to the beginning
	// of the next.  The format is the pixel format of the buffer and
	// must be one of those advertised through the wl_shm.format event.
	// A buffer will keep a reference to the pool it was created from
	// so it is valid to destroy the pool immediately after creating
	// a buffer from it.
	CreateBuffer(Id WlNewId, Offset WlInt, Width WlInt, Height WlInt, Stride WlInt, Format WlUint)
	// Destroy the shared memory pool.
	// The mmapped memory will be released when all
	// buffers that have been created from this pool
	// are gone.
	Destroy()
	// This request will cause the server to remap the backing memory
	// for the pool from the file descriptor passed when the pool was
	// created, but using the new size.  This request can only be
	// used to make the pool bigger.
	Resize(Size WlInt)
}
