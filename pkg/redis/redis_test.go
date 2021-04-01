package redis

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const randomKey = "forgot_password:50168da5"
const randomValue = "13"

func TestRedisImpl(t *testing.T) {
	err := client.Set(randomKey, randomValue, time.Duration(1*time.Hour))
	require.NoError(t, err)

	value, err := client.GetString(randomKey)
	require.NoError(t, err)
	require.Equal(t, value, randomValue)

	err = client.Delete(randomKey)
	require.NoError(t, err)

	_, err = client.GetString(randomKey)
	require.Error(t, err)
}
