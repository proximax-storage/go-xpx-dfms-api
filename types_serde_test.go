package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalUnmarshalJSONVerifyResult(t *testing.T) {
	in := RandVerifyResult(t)

	data, err := in.MarshalJSON()
	require.Nil(t, err, err)

	out := VerifyResult{}
	err = out.UnmarshalJSON(data)
	require.Nil(t, err, err)

	assert.Equal(t, in, out)
}
