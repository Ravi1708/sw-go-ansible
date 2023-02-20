package main

import (
	"context"

	"github.com/Ravi1708/sw-go-ansible/pkg/execute"
	"github.com/Ravi1708/sw-go-ansible/pkg/execute/measure"
	"github.com/Ravi1708/sw-go-ansible/pkg/options"
	"github.com/Ravi1708/sw-go-ansible/pkg/playbook"
)

func main() {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	executorTimeMeasurement := measure.NewExecutorTimeMeasurement(
		execute.NewDefaultExecute(
			execute.WithEnvVar("ANSIBLE_FORCE_COLOR", "true"),
		),
		measure.WithShowDuration(),
	)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              executorTimeMeasurement,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
