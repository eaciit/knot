package evening

import "github.com/eaciit/knot/knot.v1"

type Evening struct {
}

func (h *Evening) Index(r *knot.WebContext) interface{} {
	message := "Accessing /evening/index. There is message from /morning/index: " +
		knot.GetSharedObject().Get("name", "").(string)
	return message
}
