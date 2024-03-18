package storage

import (
	"github.com/spf13/cobra"
	"github.com/ydb-platform/ydbops/internal/cobra_util"
	"github.com/ydb-platform/ydbops/pkg/options"
	"github.com/ydb-platform/ydbops/pkg/rolling"
	"github.com/ydb-platform/ydbops/pkg/rolling/restarters"
)

func NewStorageK8sCmd() *cobra.Command {
	restartOpts := options.RestartOptionsInstance
	rootOpts := options.RootOptionsInstance
	restarter := restarters.NewStorageK8sRestarter(options.Logger)

	cmd := cobra_util.SetDefaultsOn(&cobra.Command{
		Use:   "k8s",
		Short: "Restarts a specified subset of YDB storage Pods in a Kubernetes cluster",
		Long: `ydbops restart storage k8s:
  Restarts a specified subset of YDB storage Pods in a Kubernetes cluster.
  Not specifying any filters will restart all storage Pods.`,
		Run: func(cmd *cobra.Command, args []string) {
			rolling.ExecuteRolling(*restartOpts, *rootOpts, options.Logger, restarter)
		},
	}, restarter.Opts)

	restarter.Opts.DefineFlags(cmd.Flags())
	return cmd
}

func init() {
}
