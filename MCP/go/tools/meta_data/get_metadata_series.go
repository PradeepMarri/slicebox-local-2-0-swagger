package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/slicebox-api/mcp-server/config"
	"github.com/slicebox-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_metadata_seriesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["startindex"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startindex=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["studyid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("studyid=%v", val))
		}
		if val, ok := args["sources"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sources=%v", val))
		}
		if val, ok := args["seriestypes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seriestypes=%v", val))
		}
		if val, ok := args["seriestags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seriestags=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/metadata/series%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result []Series
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGet_metadata_seriesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_metadata_series",
		mcp.WithDescription("Returns a list of metadata on the series level of the DICOM hierarchy"),
		mcp.WithNumber("startindex", mcp.Description("start index of returned slice of series")),
		mcp.WithNumber("count", mcp.Description("size of returned slice of series")),
		mcp.WithNumber("studyid", mcp.Required(), mcp.Description("reference to study to list series for")),
		mcp.WithString("sources", mcp.Description("filter the results by matching on one or more series sources. Examples of sources are user, box, directory or scp. The list of sources to filter results by must have the form TYPE1:ID1,TYPE2:ID2,...,TYPEN:IDN. For instance, the argument sources=box:1,user:5 shows results either sent from (slice)box with id 1 or uploaded by user with id 5.")),
		mcp.WithString("seriestypes", mcp.Description("filter the results by matching on one or more series types. The supplied list of series types must be a comma separated list of series type ids. For instance, the argument seriestypes=3,7,22 shows series assigned to either of the series types with ids 3, 7 and 22.")),
		mcp.WithString("seriestags", mcp.Description("filter the results by matching on one or more series tags. The supplied list of series tags must be a comma separated list of series tag ids. For instance, the argument seriestags=6,2,11 shows series with either of the series tags with ids 6, 2 and 11.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_metadata_seriesHandler(cfg),
	}
}
