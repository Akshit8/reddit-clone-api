package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const randomKey = "randomkey:q3434234"
const randomValue = "randomvalue:234343204923-049-30492-03"

func TestRedisImpl(t *testing.T) {
	err := client.Set(context.Background(), randomKey, randomValue, time.Duration(1 * time.Hour))
	require.NoError(t, err)

	value, err := client.GetString(context.Background(), randomKey)
	require.NoError(t, err)
	require.Equal(t, value, randomValue)
}
