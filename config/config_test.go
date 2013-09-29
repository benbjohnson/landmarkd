package config

import (
    "bytes"
    "testing"
    "github.com/stretchr/testify/assert"
)

const testConfig = `
port=9000
pid-path = "/home/pid"
`

func TestDecode(t *testing.T) {
    config := NewConfig()
    err := config.Decode(bytes.NewBufferString(testConfig))
    assert.Nil(t, err)
    assert.Equal(t, config.Port, 9000)
    assert.Equal(t, config.PidPath, "/home/pid")
}
