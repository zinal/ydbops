package restarters

import "github.com/spf13/pflag"

type baremetalOpts struct {
	sshArgs                 []string
}

func (o *baremetalOpts) DefineFlags(fs *pflag.FlagSet) {
	fs.StringSliceVarP(&o.sshArgs, "ssh-args", "", nil,
		`This argument will be used when ssh-ing to the nodes. It may be used to override 
the ssh command itself, ssh username or any additional arguments.
E.g.:
	--ssh-args=pssh,-A,-J,<some jump host>,--yc-profile,<YC profile name>`)
}

func (o *baremetalOpts) Validate() error {
	return nil
}
