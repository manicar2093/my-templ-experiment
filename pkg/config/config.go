// Package config contains a struct with all your API configs
package config

import (
    "templ-demo/core"
    "templ-demo/core/connections"
)

type Config struct {
    core.Config
    connections.DatabaseConnectionConfig
}
