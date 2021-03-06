package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	config := New()

	assert.Equal(t, defaultLogLevel, config.LogLevel)
	assert.Equal(t, defaultServiceName, config.ServiceName)
	assert.Equal(t, defaultServiceGRPCPort, config.ServiceGRPCPort)
	assert.Equal(t, defaultServiceHTTPPort, config.ServiceHTTPPort)
	assert.Equal(t, defaultServerTimeout, config.ServerTimeout)
	assert.Equal(t, defaultArangoEndpoints, config.ArangoEndpoints)
	assert.Equal(t, defaultArangoUser, config.ArangoUser)
	assert.Empty(t, config.ArangoPassword)
	assert.Equal(t, defaultArangoDatabase, config.ArangoDatabase)
	assert.Equal(t, defaultArangoCollection, config.ArangoCollection)
	assert.Equal(t, defaultJaegerAgentAddr, config.JaegerAgentAddr)
	assert.Equal(t, defaultJaegerLogSpans, config.JaegerLogSpans)
	assert.Empty(t, config.CAChainFile)
	assert.Empty(t, config.ServerCertFile)
	assert.Empty(t, config.ServerKeyFile)
}
