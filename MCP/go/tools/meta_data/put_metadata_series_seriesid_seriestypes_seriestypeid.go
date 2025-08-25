package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/slicebox-api/mcp-server/config"
	"github.com/slicebox-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_metadata_series_seriesid_seriestypes_seriestypeidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		seriesIdVal, ok := args["seriesId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: seriesId"), nil
		}
		seriesId, ok := seriesIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: seriesId"), nil
		}
		seriesTypeIdVal, ok := args["seriesTypeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: seriesTypeId"), nil
		}
		seriesTypeId, ok := seriesTypeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: seriesTypeId"), nil
		}
		url := fmt.Sprintf("%s/metadata/series/%s/seriestypes/%s", cfg.BaseURL, seriesId, seriesTypeId)
		req, err := http.NewRequest("PUT", url, nil)
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
		var result map[string]interface{}
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

func CreatePut_metadata_series_seriesid_seriestypes_seriestypeidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_metadata_series_seriesId_seriestypes_seriesTypeId",
		mcp.WithDescription("Add the series type with the supplied series type ID to the series with the supplied series ID"),
		mcp.WithNumber("seriesId", mcp.Required(), mcp.Description("ID of series")),
		mcp.WithNumber("seriesTypeId", mcp.Required(), mcp.Description("ID of series type to add")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_metadata_series_seriesid_seriestypes_seriestypeidHandler(cfg),
	}
}
