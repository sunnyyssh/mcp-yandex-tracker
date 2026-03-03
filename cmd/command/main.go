package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/mcpserver"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

func main() {
	trackerClient := tracker.New(tracker.Config{
		Token:      os.Getenv("TRACKER_TOKEN"),
		OrgID:      os.Getenv("TRACKER_ORG_ID"),
		CloudOrgID: os.Getenv("TRACKER_CLOUD_ORG_ID"),
	})

	mcpServer := mcpserver.NewServer(trackerClient)

	if err := mcpServer.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
