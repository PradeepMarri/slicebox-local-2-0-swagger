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

func Get_anonymization_keysHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["orderby"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderby=%v", val))
		}
		if val, ok := args["orderascending"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderascending=%v", val))
		}
		if val, ok := args["filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/anonymization/keys%s", cfg.BaseURL, queryString)
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
		var result []AnonymizationKey
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

func CreateGet_anonymization_keysTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_anonymization_keys",
		mcp.WithDescription("get a list of anonymization keys, each specifying how vital DICOM attributes have been anonymized for a particular image"),
		mcp.WithNumber("startindex", mcp.Description("start index of returned slice of anonymization keys")),
		mcp.WithNumber("count", mcp.Description("size of returned slice of anonymization keys")),
		mcp.WithString("orderby", mcp.Description("property to order results by")),
		mcp.WithBoolean("orderascending", mcp.Description("order result ascendingly if true, descendingly otherwise")),
		mcp.WithString("filter", mcp.Description("filter the results by matching substrings of properties against this value")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_anonymization_keysHandler(cfg),
	}
}
