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

func Delete_seriestypes_rules_ruleid_attributes_attributeidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ruleIdVal, ok := args["ruleId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ruleId"), nil
		}
		ruleId, ok := ruleIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ruleId"), nil
		}
		attributeIdVal, ok := args["attributeId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: attributeId"), nil
		}
		attributeId, ok := attributeIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: attributeId"), nil
		}
		url := fmt.Sprintf("%s/seriestypes/rules/%s/attributes/%s", cfg.BaseURL, ruleId, attributeId)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDelete_seriestypes_rules_ruleid_attributes_attributeidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_seriestypes_rules_ruleId_attributes_attributeId",
		mcp.WithDescription("remove the series type rule attribute corresponding to the supplied series type and attribute IDs"),
		mcp.WithNumber("ruleId", mcp.Required(), mcp.Description("id of series type rule for which to remove an attribute")),
		mcp.WithNumber("attributeId", mcp.Required(), mcp.Description("id of attribute to remove")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_seriestypes_rules_ruleid_attributes_attributeidHandler(cfg),
	}
}
