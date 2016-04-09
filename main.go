package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"k8s.io/kubernetes/pkg/api"
	//"k8s.io/kubernetes/pkg/fields"
	//"k8s.io/kubernetes/pkg/labels"

	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned"
)

func checkKubePods(c *cli.Context) {

	clientConfig := restclient.Config{}
	clientConfig.Host = "127.0.0.1:8080"
	clientU, err := unversioned.New(&clientConfig)
	if err != nil {
		fmt.Println("new unverioned err!")
	}
	pods, err := clientU.Pods("").List(api.ListOptions{})
	if err != nil {
		fmt.Printf("list pods err! err := %v\n", err)
	}
	for _, pod := range pods.Items {
		for _, cond := range pod.Status.Conditions {
			fmt.Printf("pod.Name := %v pod.Type := %v, pod.Status := %v ,pod.Status.Phase := %v\n", pod.Name, cond.Type, cond.Status, pod.Status.Phase)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "check_kube_nodes"
	app.HelpName = app.Name
	app.Usage = "Nagios check to verify Kubernetes resources status"
	app.Version = "1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "api-endpoint",
			Value: "",
			Usage: "Kubernetes API Endpoint",
		},
		cli.StringFlag{
			Name:  "username",
			Value: "",
			Usage: "Kubernetes API Username",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "",
			Usage: "Kubernetes API Password",
		},
		cli.BoolFlag{
			Name:  "skip-tls-verify",
			Usage: "Skip TLS certificate verification",
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "pod",
			Aliases: []string{"p"},
			Usage:   "check pod status",
			Action: func(c *cli.Context) {
				checkKubePods(c)
			},
		},
	}

	app.Run(os.Args)
}
