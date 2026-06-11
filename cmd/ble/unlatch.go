package cmd

import (
	"runtime"
	"time"

	c "github.com/nuki-io/nuki-cli/cmd"
	"github.com/nuki-io/nuki-cli/pkg/blecommands"
	"github.com/nuki-io/nuki-cli/pkg/bleflows"
	"github.com/nuki-io/nuki-cli/pkg/nukible"
	"github.com/spf13/cobra"
)

var unlatchCmd = &cobra.Command{
	Use:     "unlatch",
	Short:   "Unlatch a device via Bluetooth",
	PreRunE: mustDeviceId,
	Run: func(cmd *cobra.Command, args []string) {
		ble, err := nukible.NewNukiBle()
		if err != nil {
			c.Logger.Error("Failed to enable bluetooth device", "error", err.Error())
			return
		}
		if runtime.GOOS == "linux" {
			err = ble.ScanForDevice(deviceId, 10*time.Second)
			if err != nil {
				c.Logger.Error("Failed to scan", "error", err.Error())
				return
			}
		}
		flow, err := bleflows.NewAuthenticatedFlow(ble, deviceId)
		if flow == nil {
			c.Logger.Error("Failed to create BLE flow", "error", err.Error())
			return
		}
		defer flow.DisconnectDevice()
		err = flow.PerformLockOperation(blecommands.Unlatch)

		if err != nil {
			c.Logger.Error("Failed to perform unlatch operation", "error", err.Error())
			return
		}
	},
}

func init() {
	bleCmd.AddCommand(unlatchCmd)
}
