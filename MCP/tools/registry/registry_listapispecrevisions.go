package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/registry-api/mcp-server/config"
	"github.com/registry-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Registry_listapispecrevisionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		projectVal, ok := args["project"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: project"), nil
		}
		project, ok := projectVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: project"), nil
		}
		locationVal, ok := args["location"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: location"), nil
		}
		location, ok := locationVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: location"), nil
		}
		apiVal, ok := args["api"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: api"), nil
		}
		api, ok := apiVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: api"), nil
		}
		versionVal, ok := args["version"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: version"), nil
		}
		version, ok := versionVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: version"), nil
		}
		specVal, ok := args["spec"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: spec"), nil
		}
		spec, ok := specVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: spec"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/projects/%s/locations/%s/apis/%s/versions/%s/specs/%s:listRevisions%s", cfg.BaseURL, project, location, api, version, spec, queryString)
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
		var result models.ListApiSpecRevisionsResponse
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

func CreateRegistry_listapispecrevisionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_projects_project_locations_location_apis_api_versions_version_specs_spec:listRevisions",
		mcp.WithDescription("ListApiSpecRevisions lists all revisions of a spec.
 Revisions are returned in descending order of revision creation time."),
		mcp.WithString("project", mcp.Required(), mcp.Description("The project id.")),
		mcp.WithString("location", mcp.Required(), mcp.Description("The location id.")),
		mcp.WithString("api", mcp.Required(), mcp.Description("The api id.")),
		mcp.WithString("version", mcp.Required(), mcp.Description("The version id.")),
		mcp.WithString("spec", mcp.Required(), mcp.Description("The spec id.")),
		mcp.WithNumber("pageSize", mcp.Description("The maximum number of revisions to return per page.")),
		mcp.WithString("pageToken", mcp.Description("The page token, received from a previous ListApiSpecRevisions call. Provide this to retrieve the subsequent page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Registry_listapispecrevisionsHandler(cfg),
	}
}
