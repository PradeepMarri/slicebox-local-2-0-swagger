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

func Get_images_id_pngHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["framenumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("framenumber=%v", val))
		}
		if val, ok := args["windowmin"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("windowmin=%v", val))
		}
		if val, ok := args["windowmax"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("windowmax=%v", val))
		}
		if val, ok := args["imageheight"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("imageheight=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/images/%s/png%s", cfg.BaseURL, id, queryString)
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

func CreateGet_images_id_pngTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_images_id_png",
		mcp.WithDescription("get a PNG image representation of the image corresponding to the supplied ID"),
		mcp.WithNumber("id", mcp.Required(), mcp.Description("ID of image")),
		mcp.WithNumber("framenumber", mcp.Description("frame/slice to show")),
		mcp.WithNumber("windowmin", mcp.Description("intensity window minimum value. If not specified or set to zero, windowing will be selected from relevant DICOM attributes")),
		mcp.WithNumber("windowmax", mcp.Description("intensity window maximum value. If not specified or set to zero, windowing will be selected from relevant DICOM attributes")),
		mcp.WithNumber("imageheight", mcp.Description("height of PNG image. If not specified or set to zero, the image height will equal that of the data")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_images_id_pngHandler(cfg),
	}
}
