package evening

import "github.com/eaciit/knot/knot.v1"

type Evening struct {
}

func (h *Evening) Index(r *knot.WebContext) interface{} {
	sharedMessage := knot.GetSharedObject().Get("name")

	if sharedMessage != nil {
		message := sharedMessage.(string)
		// or `knot.GetSharedObject().GetDefaultValue("name", "").(string)`

		return "There is message from /morning/index: " + message
	} else {
		return "Accessing /evening/index."
	}
}
