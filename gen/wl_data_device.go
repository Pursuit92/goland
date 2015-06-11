package gen

// There is one wl_data_device per seat which can be obtained
// from the global wl_data_device_manager singleton.
// A wl_data_device provides access to inter-client data transfer
// mechanisms such as copy-and-paste and drag-and-drop.
type WlDataDevice interface {
	// The data_offer event introduces a new wl_data_offer object,
	// which will subsequently be used in either the
	// data_device.enter event (for drag-and-drop) or the
	// data_device.selection event (for selections).  Immediately
	// following the data_device_data_offer event, the new data_offer
	// object will send out data_offer.offer events to describe the
	// mime types it offers.
	DataOffer(Id WlNewId)
	// This event is sent when an active drag-and-drop pointer enters
	// a surface owned by the client.  The position of the pointer at
	// enter time is provided by the x and y arguments, in surface
	// local coordinates.
	Enter(Serial WlUint, Surface WlObject, X WlFixed, Y WlFixed, Id WlObject)
	// This event is sent when the drag-and-drop pointer leaves the
	// surface and the session ends.  The client must destroy the
	// wl_data_offer introduced at enter time at this point.
	Leave()
	// This event is sent when the drag-and-drop pointer moves within
	// the currently focused surface. The new position of the pointer
	// is provided by the x and y arguments, in surface local
	// coordinates.
	Motion(Time WlUint, X WlFixed, Y WlFixed)
	// The event is sent when a drag-and-drop operation is ended
	// because the implicit grab is removed.
	Drop()
	// The selection event is sent out to notify the client of a new
	// wl_data_offer for the selection for this device.  The
	// data_device.data_offer and the data_offer.offer events are
	// sent out immediately before this event to introduce the data
	// offer object.  The selection event is sent to a client
	// immediately before receiving keyboard focus and when a new
	// selection is set while the client has keyboard focus.  The
	// data_offer is valid until a new data_offer or NULL is received
	// or until the client loses keyboard focus.  The client must
	// destroy the previous selection data_offer, if any, upon receiving
	// this event.
	Selection(Id WlObject)
	// This request asks the compositor to start a drag-and-drop
	// operation on behalf of the client.
	// The source argument is the data source that provides the data
	// for the eventual data transfer. If source is NULL, enter, leave
	// and motion events are sent only to the client that initiated the
	// drag and the client is expected to handle the data passing
	// internally.
	// The origin surface is the surface where the drag originates and
	// the client must have an active implicit grab that matches the
	// serial.
	// The icon surface is an optional (can be NULL) surface that
	// provides an icon to be moved around with the cursor.  Initially,
	// the top-left corner of the icon surface is placed at the cursor
	// hotspot, but subsequent wl_surface.attach request can move the
	// relative position. Attach requests must be confirmed with
	// wl_surface.commit as usual. The icon surface is given the role of
	// a drag-and-drop icon. If the icon surface already has another role,
	// it raises a protocol error.
	// The current and pending input regions of the icon wl_surface are
	// cleared, and wl_surface.set_input_region is ignored until the
	// wl_surface is no longer used as the icon surface. When the use
	// as an icon ends, the current and pending input regions become
	// undefined, and the wl_surface is unmapped.
	StartDrag(Source WlObject, Origin WlObject, Icon WlObject, Serial WlUint)
	// This request asks the compositor to set the selection
	// to the data from the source on behalf of the client.
	// To unset the selection, set the source to NULL.
	SetSelection(Source WlObject, Serial WlUint)
	// This request destroys the data device.
	Release()
}
type WlDataDeviceError uint32

const (
	WlDataDeviceRole WlDataDeviceError = 0
)
