package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getParts(t *testing.T) {
	r1 := getParts("/hello/world")

	r2 := getParts("/hello/world/ed-web")

	r3 := getParts("//hello/world//ed-web")

	r4 := getParts("//hello///////world//ed-web")

	r5 := getParts("///hello/////world///ed-web")

	assert.Equal(t, []string{"hello", "world"}, r1)
	assert.Equal(t, []string{"hello", "world", "ed-web"}, r2)
	assert.Equal(t, []string{"hello", "world", "ed-web"}, r3)
	assert.Equal(t, []string{"hello", "world", "ed-web"}, r4)
	assert.Equal(t, []string{"hello", "world", "ed-web"}, r5)
}
