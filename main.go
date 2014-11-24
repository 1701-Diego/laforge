package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry-incubator/receptor"
)

type Experiment struct {
	Name        string
	Description string
	Experiment  func(receptor.Client, string, string) error
}

var experiments = []Experiment{
	{
		Name:        "away-team",
		Description: "Deploy riker as an LRP (no Docker, not for Diego-Edge)",
		Experiment:  AwayTeam,
	},
	{
		Name:        "docker-away-team",
		Description: "Deploy the away-team Dockerimage as an LRP",
		Experiment:  DockerAwayTeam,
	},
	{
		Name:        "modulate-frequencies",
		Description: "Perform a one-off task",
		Experiment:  ModulateFrequencies,
	},
}

func main() {
	if len(os.Args) != 3 {
		PrintUsageAndExit()
	}

	experimentName := os.Args[1]
	domain := os.Args[2]

	receptorAddr := os.Getenv("RECEPTOR")
	if receptorAddr == "" {
		fmt.Println("No RECEPTOR set")
		PrintUsageAndExit()
	}

	client := receptor.NewClient(receptorAddr)
	routeRoot := strings.Split(receptorAddr, "receptor.")[1]

	for _, experiment := range experiments {
		if experiment.Name == experimentName {
			fmt.Printf("Running experiment: %s\n\n", experiment.Name)
			err := experiment.Experiment(client, domain, routeRoot)
			if err != nil {
				fmt.Printf("Failed!\n%s\n", err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}
	}

	fmt.Printf("Unknown experiment:%s\n", experimentName)
	PrintUsageAndExit()
}

func PrintUsageAndExit() {
	experimentDocs := []string{}

	for _, experiment := range experiments {
		experimentDocs = append(experimentDocs, fmt.Sprintf("  %s: %s", experiment.Name, experiment.Description))
	}

	fmt.Println(fmt.Sprintf(`Usage:
laforge EXPERIMENT DOMAIN

Set the receptor address with the RECEPTOR environment:
    export RECEPTOR=http://username:password@receptor.ketchup.cf-app.com

The address for a local Diego Edge box can be set via: 
    export RECEPTOR=http://receptor.192.168.11.11.xip.io

Available EXPERIMENTs:

%s
`, strings.Join(experimentDocs, "\n")))
	os.Exit(1)
}
