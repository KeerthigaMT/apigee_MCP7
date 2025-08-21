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

func Registry_createartifactHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["artifactId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("artifactId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Artifact
		
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
		url := fmt.Sprintf("%s/v1/projects/%s/locations/%s/artifacts%s", cfg.BaseURL, project, location, queryString)
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
		var result models.Artifact
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

func CreateRegistry_createartifactTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_projects_project_locations_location_artifacts",
		mcp.WithDescription("CreateArtifact creates a specified artifact."),
		mcp.WithString("project", mcp.Required(), mcp.Description("The project id.")),
		mcp.WithString("location", mcp.Required(), mcp.Description("The location id.")),
		mcp.WithString("artifactId", mcp.Description("Required. The ID to use for the artifact, which will become the final component of the artifact's resource name. This value should be 4-63 characters, and valid characters are /[a-z][0-9]-/. Following AIP-162, IDs must not have the form of a UUID.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. Last update timestamp.")),
		mcp.WithString("contents", mcp.Description("Input parameter: Input only. The contents of the artifact. Provided by API callers when artifacts are created or replaced. To access the contents of an artifact, use GetArtifactContents.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Creation timestamp.")),
		mcp.WithString("hash", mcp.Description("Input parameter: Output only. A SHA-256 hash of the artifact's contents. If the artifact is gzipped, this is the hash of the uncompressed artifact.")),
		mcp.WithString("mimeType", mcp.Description("Input parameter: A content type specifier for the artifact. Content type specifiers are Media Types (https://en.wikipedia.org/wiki/Media_type) with a possible \"schema\" parameter that specifies a schema for the stored information. Content types can specify compression. Currently only GZip compression is supported (indicated with \"+gzip\").")),
		mcp.WithString("name", mcp.Description("Input parameter: Resource name.")),
		mcp.WithNumber("sizeBytes", mcp.Description("Input parameter: Output only. The size of the artifact in bytes. If the artifact is gzipped, this is the size of the uncompressed artifact.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Registry_createartifactHandler(cfg),
	}
}
