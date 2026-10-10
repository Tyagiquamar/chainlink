package cmd

import (
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestNewHTTPClientEnforcesTLS12Minimum(t *testing.T) {
	t.Parallel()

	client := newHTTPClient(logger.Test(t), false)
	transport, ok := client.Transport.(*http.Transport)
	require.True(t, ok)
	require.NotNil(t, transport.TLSClientConfig)
	require.Equal(t, uint16(tls.VersionTLS12), transport.TLSClientConfig.MinVersion)
}
