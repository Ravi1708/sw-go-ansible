package main

import (
	"context"
	"fmt"

	"github.com/Ravi1708/sw-go-ansible/pkg/adhoc"
	"github.com/Ravi1708/sw-go-ansible/pkg/options"
)

func main() {

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  " 127.0.0.1,",
		ModuleName: "command",
		Args:       "ping 127.0.0.1 -c 2",
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}

	fmt.Println("Command: ", adhoc.String())

	err := adhoc.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
