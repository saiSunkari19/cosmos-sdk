package types

import (
	"fmt"

	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

// ConnectionPaths define all the connection paths for a given client state.
type ConnectionPaths struct {
	ClientID string   `json:"client_id" yaml:"client_id"`
	Paths    []string `json:"paths" yaml:"paths"`
}

// NewConnectionPaths creates a ConnectionPaths instance.
func NewConnectionPaths(id string, paths []string) ConnectionPaths {
	return ConnectionPaths{
		ClientID: id,
		Paths:    paths,
	}
}

// GenesisState defines the ibc connection submodule's genesis state.
type GenesisState struct {
	Connections           []ConnectionEnd   `json:"connections" yaml:"connections"`
	ClientConnectionPaths []ConnectionPaths `json:"client_connection_paths" yaml:"client_connection_paths"`
}

// NewGenesisState creates a GenesisState instance.
func NewGenesisState(
	connections []ConnectionEnd, connPaths []ConnectionPaths,
) GenesisState {
	return GenesisState{
		Connections:           connections,
		ClientConnectionPaths: connPaths,
	}
}

// DefaultGenesisState returns the ibc connection submodule's default genesis state.
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Connections:           []ConnectionEnd{},
		ClientConnectionPaths: []ConnectionPaths{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for i, conn := range gs.Connections {
		if err := conn.ValidateBasic(); err != nil {
			return fmt.Errorf("invalid connection %d: %w", i, err)
		}
	}

	for i, conPaths := range gs.ClientConnectionPaths {
		if err := host.DefaultClientIdentifierValidator(conPaths.ClientID); err != nil {
			return fmt.Errorf("invalid client connection path %d: %w", i, err)
		}
		for _, path := range conPaths.Paths {
			if err := host.DefaultPathValidator(path); err != nil {
				return fmt.Errorf("invalid client connection path %d: %w", i, err)
			}
		}
	}

	return nil
}
