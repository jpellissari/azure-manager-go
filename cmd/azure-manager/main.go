package main

import (
	"fmt"
	"log"

	azureapi "github.com/jpellissari/azure-manager-go/pkg/azure-api"
	"github.com/jpellissari/azure-manager-go/pkg/cli"
)

type ColumnsTracked string

const (
	Blocked            ColumnsTracked = "Blocked"
	Backlog            ColumnsTracked = "Backlog"
	In_Progress        ColumnsTracked = "In Progress"
	Dev_Testing        ColumnsTracked = "Dev Testing"
	Acceptance_Backlog ColumnsTracked = "Acceptance Backlog"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatal("Whoops, something went wrong parsings args")
	}

	if len(opts.Args) == 0 {
		defaultMenu()
	}
}

func defaultMenu() {
	groupedWorkItems := make(map[string][]azureapi.WorkItem)
	workItems := azureapi.LoadActiveWorkItems()
	for _, wi := range workItems {
		if wi.Column == string(Backlog) {
			groupedWorkItems[string(Backlog)] = append(groupedWorkItems[string(Backlog)], wi)
		} else if wi.Column == string(In_Progress) {
			groupedWorkItems[string(In_Progress)] = append(groupedWorkItems[string(In_Progress)], wi)
		} else if wi.Column == string(Dev_Testing) {
			groupedWorkItems[string(Dev_Testing)] = append(groupedWorkItems[string(Dev_Testing)], wi)
		} else if wi.Column == string(Acceptance_Backlog) {
			groupedWorkItems[string(Acceptance_Backlog)] = append(groupedWorkItems[string(Acceptance_Backlog)], wi)
		} else if wi.Column == string(Blocked) {
			groupedWorkItems[string(Blocked)] = append(groupedWorkItems[string(Blocked)], wi)
		} else {
			groupedWorkItems["Untracked"] = append(groupedWorkItems["Untracked"], wi)
		}
	}

	fmt.Printf("Welcome to this amazing tool :)\n")
	fmt.Printf("There's %d workitems assigned to you right now:\n", len(workItems))
	if len(groupedWorkItems[string(Backlog)]) > 0 {
		fmt.Printf("- Backlog:\n")
		for _, wi := range groupedWorkItems[string(Backlog)] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
	if len(groupedWorkItems[string(Blocked)]) > 0 {
		fmt.Printf("- Blocked\n")
		for _, wi := range groupedWorkItems[string(Blocked)] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
	if len(groupedWorkItems[string(In_Progress)]) > 0 {
		fmt.Printf("- In Progress\n")
		for _, wi := range groupedWorkItems[string(In_Progress)] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
	if len(groupedWorkItems[string(Dev_Testing)]) > 0 {
		fmt.Printf("- Dev Testing\n")
		for _, wi := range groupedWorkItems[string(Dev_Testing)] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
	if len(groupedWorkItems[string(Acceptance_Backlog)]) > 0 {
		fmt.Printf("- Acceptance Backlog\n")
		for _, wi := range groupedWorkItems[string(Acceptance_Backlog)] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
	if len(groupedWorkItems["Untracked"]) > 0 {
		fmt.Printf("- Untracked\n")
		for _, wi := range groupedWorkItems["Untracked"] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
}
