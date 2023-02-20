package main

import (
	"context"

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

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"input.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
