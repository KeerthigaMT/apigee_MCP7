package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/registry-api/mcp-server/config"
	"github.com/registry-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Registry_createapiversionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["apiVersionId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiVersionId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ApiVersion
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/projects/%s/locations/%s/apis/%s/versions%s", cfg.BaseURL, project, location, api, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.ApiVersion
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

func CreateRegistry_createapiversionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_projects_project_locations_location_apis_api_versions",
		mcp.WithDescription("CreateApiVersion creates a specified version."),
		mcp.WithString("project", mcp.Required(), mcp.Description("The project id.")),
		mcp.WithString("location", mcp.Required(), mcp.Description("The location id.")),
		mcp.WithString("api", mcp.Required(), mcp.Description("The api id.")),
		mcp.WithString("apiVersionId", mcp.Description("Required. The ID to use for the version, which will become the final component of the version's resource name. This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/. Following AIP-162, IDs must not have the form of a UUID.")),
		mcp.WithObject("annotations", mcp.Description("Input parameter: Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Creation timestamp.")),
		mcp.WithString("description", mcp.Description("Input parameter: A detailed description.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: Human-meaningful name.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with \"apigeeregistry.googleapis.com/\" and cannot be changed.")),
		mcp.WithString("name", mcp.Description("Input parameter: Resource name.")),
		mcp.WithString("state", mcp.Description("Input parameter: A user-definable description of the lifecycle phase of this API version. Format: free-form, but we expect single words that describe API maturity, e.g. \"CONCEPT\", \"DESIGN\", \"DEVELOPMENT\", \"STAGING\", \"PRODUCTION\", \"DEPRECATED\", \"RETIRED\".")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. Last update timestamp.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Registry_createapiversionHandler(cfg),
	}
}
