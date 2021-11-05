package engine

import "testing"
import "github.com/stretchr/testify/assert"

func TestEngine_getPattern(t *testing.T) {
	var ret string
	ret, _ = getPattern("/hello/world")
	assert.Equal(t, "/hello/world", ret)

}
