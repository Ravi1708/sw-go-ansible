package main

import (
	"context"
	"log"

	"github.com/Ravi1708/sw-go-ansible/pkg/adhoc"
	"github.com/Ravi1708/sw-go-ansible/pkg/options"
)

func main() {

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  "127.0.0.1,",
		ModuleName: "debug",
		Args: `msg="
{{ arg1 }}
{{ arg2 }}
{{ arg3 }}
"`,
		ExtraVars: map[string]interface{}{
			"arg1": map[string]interface{}{"subargument": "subargument_value"},
			"arg2": "arg2_value",
			"arg3": "arg3_value",
		},
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		//StdoutCallback:    "oneline",
	}

	log.Println("Command: ", adhoc)

	err := adhoc.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
