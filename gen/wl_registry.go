package gen

// The global registry object.  The server has a number of global
// objects that are available to all clients.  These objects
// typically represent an actual object in the server (for example,
// an input device) or they are singleton objects that provide
// extension functionality.
// When a client creates a registry object, the registry object
// will emit a global event for each global currently in the
// registry.  Globals come and go as a result of device or
// monitor hotplugs, reconfiguration or other events, and the
// registry will send out global and global_remove events to
// keep the client up to date with the changes.  To mark the end
// of the initial burst of events, the client can use the
// wl_display.sync request immediately after calling
// wl_display.get_registry.
// A client can bind to a global object by using the bind
// request.  This creates a client-side handle that lets the object
// emit events to the client and lets the client invoke requests on
// the object.
type WlRegistry interface {
	// Notify the client of global objects.
	// The event notifies the client that a global object with
	// the given name is now available, and it implements the
	// given version of the given interface.
	Global(Name WlUint, WlInterface WlString, Version WlUint)
	// Notify the client of removed global objects.
	// This event notifies the client that the global identified
	// by name is no longer available.  If the client bound to
	// the global using the bind request, the client should now
	// destroy that object.
	// The object remains valid and requests to the object will be
	// ignored until the client destroys it, to avoid races between
	// the global going away and a client sending a request to it.
	GlobalRemove(Name WlUint)
	// Binds a new, client-created object to the server using the
	// specified name as the identifier.
	Bind(Name WlUint, Id WlNewId)
}
