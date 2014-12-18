package filter

import (
	"net/url"
	"testing"

	"github.com/bmizerany/assert"
)

func TestParamsFromValuesDontDel(t *testing.T) {
	v := url.Values{}
	v.Set(":name", "Ava")
	v.Add("friend", "Jess")
	v.Set("name", "Zoe")
	p := ParamsFromValues(v, false)

	assert.Equal(t, "name=Ava", p.Encode())
	assert.Equal(t, v.Get(":name"), p.Get("name"))
}

func TestParamsFromValues(t *testing.T) {
	v := url.Values{}
	v.Set(":name", "Ava")
	v.Add("friend", "Jess")
	v.Set("name", "Zoe")
	p := ParamsFromValues(v, true)

	assert.Equal(t, "friend=Jess&name=Zoe", v.Encode())
	assert.Equal(t, "name=Ava", p.Encode())
	assert.Equal(t, "Ava", p.Get("name"))
}
