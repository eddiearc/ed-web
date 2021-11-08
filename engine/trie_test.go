package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_formatUrlPath(t *testing.T) {
	r, err := formatUrlPath("/hello/world")
	assert.Equal(t, nil, err)
	assert.Equal(t, "/hello/world", r)

	r, err = formatUrlPath("/hello/world/ed-web")
	assert.Equal(t, nil, err)
	assert.Equal(t, "/hello/world/ed-web", r)

	_, err = formatUrlPath("///hello///world//ed-web")
	assert.Equal(t, ErrPatternFormatError, err)

	_, err = formatUrlPath("///hello//world//ed-web")
	assert.Equal(t, ErrPatternFormatError, err)

	_, err = formatUrlPath("/hello//world/ed-web")
	assert.Equal(t, ErrPatternFormatError, err)

	_, err = formatUrlPath("/hello/world/ed-web?id=666")
	assert.Equal(t, ErrPatternFormatError, err)
}

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
