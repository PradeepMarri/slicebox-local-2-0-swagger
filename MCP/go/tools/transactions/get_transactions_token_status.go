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

func Get_transactions_token_statusHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		tokenVal, ok := args["token"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: token"), nil
		}
		token, ok := tokenVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: token"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["transactionid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("transactionid=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/transactions/%s/status%s", cfg.BaseURL, token, queryString)
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

func CreateGet_transactions_token_statusTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_transactions_token_status",
		mcp.WithDescription("get the status of the remote incoming transaction with the supplied transaction ID"),
		mcp.WithString("token", mcp.Required(), mcp.Description("authentication token identifying the current box-to-box connection")),
		mcp.WithNumber("transactionid", mcp.Required(), mcp.Description("the ID of the client's outgoing transaction")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_transactions_token_statusHandler(cfg),
	}
}
