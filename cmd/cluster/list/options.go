package list

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/ydb-platform/ydbops/pkg/cmdutil"
)

type Options struct {
}

const (
	DefaultMaintenanceDurationSeconds = 3600
)

func (o *Options) DefineFlags(fs *pflag.FlagSet) {
}

func (o *Options) Validate() error {
	return nil
}

func (o *Options) Run(f cmdutil.Factory) error {
	nodes, err := f.GetCMSClient().Nodes()
	if err != nil {
		return err
	}

	fmt.Printf("--- begin nodes ---\n")
	for _, node := range nodes {
		fmt.Printf("%d %s:%d %s", node.GetNodeId(), node.GetHost(), node.GetPort(), node.GetType())
	}
	fmt.Printf("--- end nodes ---\n")

	return nil
}
