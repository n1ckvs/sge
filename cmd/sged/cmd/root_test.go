package cmd_test

import (
	"testing"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/sge-network/sge/app"
	"github.com/sge-network/sge/cmd/sged/cmd"
	"github.com/stretchr/testify/require"
)

func TestRootCmdConfig(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"config",          // Test the config cmd
		"keyring-backend", // key
		"test",            // value
	})

	require.NoError(t, svrcmd.Execute(rootCmd, app.DefaultNodeHome))
}
