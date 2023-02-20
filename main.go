package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/Ravi1708/sw-go-ansible/pkg/execute"
	"github.com/Ravi1708/sw-go-ansible/pkg/options"
	"github.com/Ravi1708/sw-go-ansible/pkg/playbook"
	"github.com/Ravi1708/sw-go-ansible/pkg/stdoutcallback/results"
)

func main() {

	// var res *results.AnsiblePlaybookJSONResults

	// buff := new(bytes.Buffer)

	// ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
	// 	User: "ravi",
	// }

	// ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
	// 	// Inventory: "/etc/ansible/hosts",
	// 	ExtraVars: map[string]interface{}{
	// 		"target": "633ebda0916487c773e7cc86",
	// 		"ansible_host" : "local",

	// 	},
	// }

	// ansiblePlaybookPrivilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{
	// 	Become: true,
	// }

	// playbook := &playbook.AnsiblePlaybookCmd{
	// 	Playbooks:                  []string{"site.yml"},
	// 	ConnectionOptions:          ansiblePlaybookConnectionOptions,
	// 	PrivilegeEscalationOptions: ansiblePlaybookPrivilegeEscalationOptions,
	// 	Options:                    ansiblePlaybookOptions,
	// 	Exec: execute.NewDefaultExecute(
	// 		execute.WithEnvVar("ANSIBLE_FORCE_COLOR", "true"),
	// 		execute.WithTransformers(
	// 			results.Prepend("Stackwatch Execution"),
	// 		),
	// 	),
	// }

	// err := playbook.Run(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }

	// results.ParseJSONResultsStream(io.Reader(buff))
	
	// res, err = results.ParseJSONResultsStream(io.Reader(buff))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(res.String())

	// node := map[string]interface{}{
	// 	"ID" : "633ebda0916487c773e7cc86",
	// 	"Hostname": "23.82.14.238",
	// 	"Ipaddress": "23.82.14.238",
	// 	"Port" : 22,
	// 	"AuthType" : "password",
	// 	"SudoUser" : true,
	// 	"SshUser": "ubuntu",
	// 	"SshPass": "fw7YED",
	// 	"Password" : "fw7YED",
	// 	"PrivilageEscalation" : "sudo",
	// 	"becomePassword": "fw7YED",
	// }
	// playbookFile := "site.yml"

	//  data, err := json.Marshal(node)
	// if err != nil {
	//     fmt.Println(err)
	// }

	// argumets := map[string]interface{}{
	// 	"ansible_host" : "23.82.14.238",
	// }

	// _ = playbook.PlaybookExecute(data, playbookFile, argumets)

	// println(err)

	err := PlaybookExecute("site.yml", map[string]interface{}{ "ansible_host" : "23.82.14.238"}) 
	if err != nil {
		panic(err)
	}

}

func PlaybookExecute(playbookFile string, arguments map[string]interface{}) error {

	// var res *results.AnsiblePlaybookJSONResults
	buff := new(bytes.Buffer)

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "inventory.txt",
		ExtraVars: map[string]interface{}{
			"ansible_host" : "23.82.14.238",
			"ansible_user" : "ubuntu",
			"ansible_password" : "fw7YED",
			"ansible_become" : true,
			"ansible_become_method" : "sudo",
			"ansible_become_user" : "root",
			"ansible_become_pass" : "fw7YED",
		},
	}

	ansiblePlaybookPrivilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{
		Become: true,
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:                  []string{playbookFile},
		PrivilegeEscalationOptions: ansiblePlaybookPrivilegeEscalationOptions,
		Options:                    ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithEnvVar("ANSIBLE_FORCE_COLOR", "true"),
			execute.WithTransformers(
				results.Prepend("Stackwatch Execution"),
			),
			execute.WithWrite(io.Writer(buff)),
		),
		StdoutCallback:    "default",
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}

	// var reader io.Reader
	// reader = bufio.NewReader(strings.NewReader( buff.String()))

	wbuff := bytes.Buffer{}
	writer := io.Writer(&wbuff)


	err = results.DefaultStdoutCallbackResults(context.TODO(),io.Reader(buff), writer )

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Output:", writer)

	return nil
}