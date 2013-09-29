package config

import (
    "github.com/BurntSushi/toml"
    "io"
    "os"
)

const (
    DefaultPort       = 9500
    DefaultPidPath    = "/var/run/landmarkd.pid"
)

// The configuration for running landmarkd.
type Config struct {
    Port       int    `toml:"port"`
    PidPath    string `toml:"pid-path"`
}

// Creates a new configuration object.
func NewConfig() *Config {
    return &Config{
        Port:       DefaultPort,
        PidPath:    DefaultPidPath,
    }
}

// Decodes the configuration from a TOML reader.
func (c *Config) Decode(r io.Reader) error {
    if _, err := toml.DecodeReader(r, &c); err != nil {
        return err
    }
    return nil
}

// Decodes the configuration from a TOML file path.
func (c *Config) DecodeFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    if err = c.Decode(f); err != nil {
        return err
    }
    return nil
}
