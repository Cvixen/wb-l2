package main

import (
	"dev01/cur_time"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTime(t *testing.T) {
	time, err := cur_time.GetTime()
	require.NoError(t, err)
	require.NotEmpty(t, time)
}
