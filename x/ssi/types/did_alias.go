package types

import (
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

type DidAliasConfig struct {
	PidxNodeConfigDir string
	DidAliasDir      string
}

func GetDidAliasConfig(cmd *cobra.Command) (*DidAliasConfig, error) {
	configDir, err := cmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return nil, err
	}

	didAliasDir := filepath.Join(configDir, "generated-ssi-docs")

	return &DidAliasConfig{
		PidxNodeConfigDir: configDir,
		DidAliasDir:      didAliasDir,
	}, nil
}
