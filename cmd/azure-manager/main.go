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
	Testing            ColumnsTracked = "Dev Testing"
	Acceptance_Backlog ColumnsTracked = "Acceptance Backlog"
	Untracked_Column   ColumnsTracked = "Untracked Column"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatal("Whoops, something went wrong parsings args")
	}

	if len(opts.Args) == 0 {
		defaultMenu()
	}

	fmt.Printf("%v", opts.Args)
}

func defaultMenu() {
	groupedWorkItems := make(map[string][]azureapi.WorkItem)
	workItems := azureapi.LoadActiveWorkItems()
	for _, wi := range workItems {
		if wi.Column == string(Backlog) {
			groupedWorkItems[string(Backlog)] = append(groupedWorkItems[string(Backlog)], wi)
		} else if wi.Column == string(In_Progress) {
			groupedWorkItems[string(In_Progress)] = append(groupedWorkItems[string(In_Progress)], wi)
		} else if wi.Column == string(Testing) {
			groupedWorkItems[string(Testing)] = append(groupedWorkItems[string(Testing)], wi)
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
	printGroupedWI(Backlog, groupedWorkItems)
	printGroupedWI(Blocked, groupedWorkItems)
	printGroupedWI(In_Progress, groupedWorkItems)
	printGroupedWI(Testing, groupedWorkItems)
	printGroupedWI(Acceptance_Backlog, groupedWorkItems)
	printGroupedWI(Untracked_Column, groupedWorkItems)
}

func printGroupedWI(column ColumnsTracked, groupedWorkItems map[string][]azureapi.WorkItem) {
	if len(groupedWorkItems[string(column)]) > 0 {
		fmt.Printf("- ")
		fmt.Printf("%s\n", string(column))
		for _, wi := range groupedWorkItems[string(column)] {
			fmt.Printf("    - %d: %s\n", wi.Id, wi.Title)
		}
	}
}
