package main

import (
	"fmt"

	"github.com/cloudfoundry-incubator/receptor"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

func AwayTeam(client receptor.Client, domain string, routeRoot string) error {
	processGuid := fmt.Sprintf("away-team-%s", domain)
	route := fmt.Sprintf("away-team-%s.%s", domain, routeRoot)
	err := client.CreateDesiredLRP(receptor.DesiredLRPCreateRequest{
		ProcessGuid: processGuid,
		Domain:      domain,
		Instances:   1,
		Stack:       "lucid64",
		//no rootfs here, instead we download riker and crusher:
		Setup: &models.SerialAction{
			Actions: []models.Action{
				&models.DownloadAction{
					From:     "http://onsi-public.s3.amazonaws.com/riker.tar.gz",
					To:       "/tmp",
					CacheKey: "riker",
				},
				&models.DownloadAction{
					From:     "http://onsi-public.s3.amazonaws.com/crusher.tar.gz",
					To:       "/tmp",
					CacheKey: "crusher",
				},
			},
		},
		Action: &models.RunAction{
			Path:      "/tmp/riker",
			LogSource: "RIKER",
		},
		Monitor: &models.RunAction{
			Path:      "/tmp/crusher",
			Args:      []string{"--port-check=8080"},
			LogSource: "CRUSHER",
		},
		DiskMB:    128,
		MemoryMB:  64,
		Ports:     []uint32{8080},
		Routes:    []string{route},
		LogGuid:   processGuid,
		LogSource: "AWAY-TEAM",
	})
	if err != nil {
		return err
	}

	fmt.Println("The away team is deployed.")
	fmt.Printf("To make contact:\n  http://%s/\n", route)
	fmt.Printf("To stream logs:\n  picard %s\n", processGuid)
	fmt.Printf("To see what's running:\n  troy %s\n", domain)
	fmt.Printf("To delete the LRP:\n  worf destroy-lrp %s\n", processGuid)

	return nil
}
