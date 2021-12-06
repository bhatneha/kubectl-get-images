package cmd

import (
	"path/filepath"

	"github.com/bhatneha/kubectl-get-images/images"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

var (
	c = &images.Config{}
)

func registerFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&c.KubeConfig, "kubeconfig", "k", filepath.Join(homedir.HomeDir(), ".kube", "config"), "absolute path to the kubeconfig file")
	cmd.Flags().BoolVar(&c.All, "all", false, "all the images in a cluster")
	cmd.Flags().StringVarP(&c.NameSpace, "namespace", "n", "default", "namespace in a cluster")
}
