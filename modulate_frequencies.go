package main

import (
	"fmt"

	"github.com/cloudfoundry-incubator/receptor"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

func ModulateFrequencies(client receptor.Client, domain string, _ string) error {
	taskGuid := fmt.Sprintf("modulate-frequencies-%s", domain)
	err := client.CreateTask(receptor.TaskCreateRequest{
		TaskGuid:   taskGuid,
		Domain:     domain,
		RootFSPath: "docker:///onsi/away-team",
		Stack:      "lucid64",
		Action: &models.RunAction{
			Path: "sh",
			Args: []string{"-c", `echo "Operating..."; printf "Current Temporal Displacement\n` + "`date`" + `\n" > /tmp/result`},
		},
		DiskMB:     128,
		MemoryMB:   64,
		LogGuid:    taskGuid,
		LogSource:  "TRICHORDER",
		ResultFile: "/tmp/result",
	})
	if err != nil {
		return err
	}

	fmt.Println("Modulating Frequencies")
	fmt.Printf("To stream logs:\n  picard %s\n", taskGuid)
	fmt.Printf("To view current status:\n  troy %s\n", domain)
	fmt.Printf("To delete the task:\n  worf delete-task %s\n", taskGuid)

	return nil
}
