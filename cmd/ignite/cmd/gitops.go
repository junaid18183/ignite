package cmd

import (
	"io"

	"github.com/lithammer/dedent"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/weaveworks/ignite/pkg/errutils"
	"github.com/weaveworks/ignite/pkg/gitops"
)

type gitOpsFlags struct {
	branch string
}

// NewCmdGitOps runs the GitOps functionality of Ignite
func NewCmdGitOps(out io.Writer) *cobra.Command {
	f := &gitOpsFlags{}
	cmd := &cobra.Command{
		Use:   "gitops <repo-url>",
		Short: "Run the GitOps feature of Ignite",
		Long: dedent.Dedent(`
			Run Ignite in GitOps mode watching the given repository. The repository needs
			to be publicly cloneable. Ignite will watch for changes in the master branch
			by default, overridable with the branch flag (-b, --branch). If any new/changed
			VM specification files are found in the repo (in JSON/YAML format), their
			configuration will automatically be declaratively applied.

			To quit GitOps mode, use (Ctrl + C).
		`),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			errutils.Check(gitops.RunLoop(args[0], f.branch))
		},
	}

	addGitOpsFlags(cmd.Flags(), f)
	return cmd
}

func addGitOpsFlags(fs *pflag.FlagSet, f *gitOpsFlags) {
	fs.StringVarP(&f.branch, "branch", "b", "master", "What branch to sync")
	// TODO: Add repo subdirectories, ssh key opts etc.
}
