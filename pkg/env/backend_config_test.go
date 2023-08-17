package env

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/onsi/gomega"
)

func Test_GetBackendConfig(t *testing.T) {
	g := NewGomegaWithT(t)
	envs := map[string]string{
		// optional
		"PUBLISHER_REQUEST_TIMEOUT": "10s",
	}

	for k, v := range envs {
		t.Setenv(k, v)
	}
	backendConfig := GetBackendConfig()
	// Ensure optional variables can be set
	g.Expect(backendConfig.PublisherConfig.RequestTimeout).To(Equal(envs["PUBLISHER_REQUEST_TIMEOUT"]))
}

func Test_ToECENVDefaultSubscriptionConfig(t *testing.T) {
	// given
	givenConfig := DefaultSubscriptionConfig{
		MaxInFlightMessages:   55,
		DispatcherRetryPeriod: 56,
		DispatcherMaxRetries:  57,
	}

	// when
	result := givenConfig.ToECENVDefaultSubscriptionConfig()

	// then
	require.Equal(t, givenConfig.DispatcherMaxRetries, result.DispatcherMaxRetries)
	require.Equal(t, givenConfig.DispatcherRetryPeriod, result.DispatcherRetryPeriod)
	require.Equal(t, givenConfig.MaxInFlightMessages, result.MaxInFlightMessages)
}
