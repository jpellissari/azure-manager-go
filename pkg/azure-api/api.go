package azureapi

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/workitemtracking"
)

type connection struct {
	con *azuredevops.Connection
	ctx *context.Context
}

type WorkItem struct {
	Id     int
	Title  string
	Column string
}

func start() connection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Whoops, something went wrong parsing the env file")
	}

	organizationUrl := os.Getenv("AZURE_ORGANIZATION_URL")
	personalAccessToken := os.Getenv("AZURE_PERSONAL_TOKEN")

	// Create a connection to your organization
	azureConnection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

	ctx := context.Background()

	return connection{
		con: azureConnection,
		ctx: &ctx,
	}
}

func getActiveWorkItemsIds(ctx context.Context, client workitemtracking.Client) []int {
	queryString := `
    SELECT 
        [System.Id]
    FROM workitems
    WHERE
       [Assigned to] = @Me
       AND [System.State] <> 'Closed'
    `
	wiql := workitemtracking.Wiql{
		Query: &queryString,
	}
	queryByWiqlArgs := workitemtracking.QueryByWiqlArgs{
		Wiql: &wiql,
	}
	workItems, err := client.QueryByWiql(ctx, queryByWiqlArgs)
	if err != nil {
		log.Fatal(err)
	}

	ids := []int{}
	for _, wi := range *workItems.WorkItems {
		ids = append(ids, *wi.Id)
	}

	return ids
}

func LoadActiveWorkItems() []WorkItem {
	con := start()

	client, err := workitemtracking.NewClient(*con.ctx, con.con)
	if err != nil {
		log.Fatal(err)
	}

	workItemsIDs := getActiveWorkItemsIds(*con.ctx, client)

	workItemsResponse, err := client.GetWorkItems(*con.ctx, workitemtracking.GetWorkItemsArgs{Ids: &workItemsIDs})
	if err != nil {
		log.Fatal(err)
	}

	workItems := []WorkItem{}
	for _, wi := range *workItemsResponse {
		workItem := WorkItem{
			Id:     *wi.Id,
			Title:  (*wi.Fields)["System.Title"].(string),
			Column: (*wi.Fields)["System.BoardColumn"].(string),
		}
		workItems = append(workItems, workItem)
	}

	return workItems
}
