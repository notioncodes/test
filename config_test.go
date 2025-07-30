package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoads(t *testing.T) {
	assert.Contains(t, TestConfig.NotionAPIKey, "ntn_")
}
