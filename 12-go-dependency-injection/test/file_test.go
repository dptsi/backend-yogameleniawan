package test

import (
	"testing"

	"github.com/dptsi/backend-yogameleniawan/12-go-dependency-injection/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databaes")
	assert.NotNil(t, connection)

	cleanup()
}
