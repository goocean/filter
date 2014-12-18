package filter

import (
	"net/url"
)

const COLON = ':'

// Filter Params from Values.
// If b is true, the param will be deleted.
func ParamsFromValues(v url.Values, b bool) (p url.Values) {
	if len(v) == 0 {
		return
	}
	p = make(url.Values)
	for k, _ := range v {
		if k[0] == COLON {
			p.Add(k[1:], v.Get(k))
			if b {
				v.Del(k)
			}
		}
	}
	return
}
