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

func Registry_updateapideploymentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		deploymentVal, ok := args["deployment"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: deployment"), nil
		}
		deployment, ok := deploymentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: deployment"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["updateMask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updateMask=%v", val))
		}
		if val, ok := args["allowMissing"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("allowMissing=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ApiDeployment
		
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
		url := fmt.Sprintf("%s/v1/projects/%s/locations/%s/apis/%s/deployments/%s%s", cfg.BaseURL, project, location, api, deployment, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.ApiDeployment
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

func CreateRegistry_updateapideploymentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_projects_project_locations_location_apis_api_deployments_deployment",
		mcp.WithDescription("UpdateApiDeployment can be used to modify a specified deployment."),
		mcp.WithString("project", mcp.Required(), mcp.Description("The project id.")),
		mcp.WithString("location", mcp.Required(), mcp.Description("The location id.")),
		mcp.WithString("api", mcp.Required(), mcp.Description("The api id.")),
		mcp.WithString("deployment", mcp.Required(), mcp.Description("The deployment id.")),
		mcp.WithString("updateMask", mcp.Description("The list of fields to be updated. If omitted, all fields are updated that are set in the request message (fields set to default values are ignored). If a \"*\" is specified, all fields are updated, including fields that are unspecified/default in the request.")),
		mcp.WithBoolean("allowMissing", mcp.Description("If set to true, and the deployment is not found, a new deployment will be created. In this situation, `update_mask` is ignored.")),
		mcp.WithString("revisionUpdateTime", mcp.Description("Input parameter: Output only. Last update timestamp: when the represented revision was last modified.")),
		mcp.WithString("accessGuidance", mcp.Description("Input parameter: Text briefly describing how to access the endpoint. Changes to this value will not affect the revision.")),
		mcp.WithObject("annotations", mcp.Description("Input parameter: Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.")),
		mcp.WithString("apiSpecRevision", mcp.Description("Input parameter: The full resource name (including revision id) of the spec of the API being served by the deployment. Changes to this value will update the revision. Format: apis/{api}/deployments/{deployment}")),
		mcp.WithString("endpointUri", mcp.Description("Input parameter: The address where the deployment is serving. Changes to this value will update the revision.")),
		mcp.WithString("externalChannelUri", mcp.Description("Input parameter: The address of the external channel of the API (e.g. the Developer Portal). Changes to this value will not affect the revision.")),
		mcp.WithString("intendedAudience", mcp.Description("Input parameter: Text briefly identifying the intended audience of the API. Changes to this value will not affect the revision.")),
		mcp.WithString("name", mcp.Description("Input parameter: Resource name.")),
		mcp.WithString("revisionCreateTime", mcp.Description("Input parameter: Output only. Revision creation timestamp; when the represented revision was created.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Creation timestamp; when the deployment resource was created.")),
		mcp.WithString("description", mcp.Description("Input parameter: A detailed description.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: Human-meaningful name.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with \"registry.googleapis.com/\" and cannot be changed.")),
		mcp.WithString("revisionId", mcp.Description("Input parameter: Output only. Immutable. The revision ID of the deployment. A new revision is committed whenever the deployment contents are changed. The format is an 8-character hexadecimal string.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Registry_updateapideploymentHandler(cfg),
	}
}
