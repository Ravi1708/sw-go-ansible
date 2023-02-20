package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/Ravi1708/sw-go-ansible/pkg/execute"
	"github.com/Ravi1708/sw-go-ansible/pkg/options"
	"github.com/Ravi1708/sw-go-ansible/pkg/playbook"
	"github.com/Ravi1708/sw-go-ansible/pkg/stdoutcallback/results"
)

func main() {

	var timeout int
	flag.IntVar(&timeout, "timeout", 15, "Timeout in seconds")
	flag.Parse()

	fmt.Printf("Timeout: %d seconds\n", timeout)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithTransformers(
				results.Prepend("Go-ansible example"),
			),
		),
		StdoutCallback: "json",
	}

	err := playbook.Run(ctx)
	if err != nil {
		panic(err)
	}
}
