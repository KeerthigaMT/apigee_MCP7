package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Api represents the Api schema from the OpenAPI specification
type Api struct {
	Availability string `json:"availability,omitempty"` // A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g. "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
	Recommendeddeployment string `json:"recommendedDeployment,omitempty"` // The recommended deployment of the API. Format: apis/{api}/deployments/{deployment}
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
	Description string `json:"description,omitempty"` // A detailed description.
	Displayname string `json:"displayName,omitempty"` // Human-meaningful name.
	Name string `json:"name,omitempty"` // Resource name.
	Recommendedversion string `json:"recommendedVersion,omitempty"` // The recommended version of the API. Format: apis/{api}/versions/{version}
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
}

// ListApiDeploymentsResponse represents the ListApiDeploymentsResponse schema from the OpenAPI specification
type ListApiDeploymentsResponse struct {
	Apideployments []ApiDeployment `json:"apiDeployments,omitempty"` // The deployments from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Hash string `json:"hash,omitempty"` // Output only. A SHA-256 hash of the artifact's contents. If the artifact is gzipped, this is the hash of the uncompressed artifact.
	Mimetype string `json:"mimeType,omitempty"` // A content type specifier for the artifact. Content type specifiers are Media Types (https://en.wikipedia.org/wiki/Media_type) with a possible "schema" parameter that specifies a schema for the stored information. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	Name string `json:"name,omitempty"` // Resource name.
	Sizebytes int `json:"sizeBytes,omitempty"` // Output only. The size of the artifact in bytes. If the artifact is gzipped, this is the size of the uncompressed artifact.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
	Contents string `json:"contents,omitempty"` // Input only. The contents of the artifact. Provided by API callers when artifacts are created or replaced. To access the contents of an artifact, use GetArtifactContents.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
}

// ListApiVersionsResponse represents the ListApiVersionsResponse schema from the OpenAPI specification
type ListApiVersionsResponse struct {
	Apiversions []ApiVersion `json:"apiVersions,omitempty"` // The versions from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// ListApiSpecRevisionsResponse represents the ListApiSpecRevisionsResponse schema from the OpenAPI specification
type ListApiSpecRevisionsResponse struct {
	Apispecs []ApiSpec `json:"apiSpecs,omitempty"` // The revisions of the spec.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// ListApiSpecsResponse represents the ListApiSpecsResponse schema from the OpenAPI specification
type ListApiSpecsResponse struct {
	Apispecs []ApiSpec `json:"apiSpecs,omitempty"` // The specs from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// RollbackApiDeploymentRequest represents the RollbackApiDeploymentRequest schema from the OpenAPI specification
type RollbackApiDeploymentRequest struct {
	Name string `json:"name"` // Required. The deployment being rolled back.
	Revisionid string `json:"revisionId"` // Required. The revision ID to roll back to. It must be a revision of the same deployment. Example: c7cfa2a8
}

// ApiSpec represents the ApiSpec schema from the OpenAPI specification
type ApiSpec struct {
	Name string `json:"name,omitempty"` // Resource name.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Revisionupdatetime string `json:"revisionUpdateTime,omitempty"` // Output only. Last update timestamp: when the represented revision was last modified.
	Sourceuri string `json:"sourceUri,omitempty"` // The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
	Description string `json:"description,omitempty"` // A detailed description.
	Hash string `json:"hash,omitempty"` // Output only. A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.
	Revisionid string `json:"revisionId,omitempty"` // Output only. Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.
	Sizebytes int `json:"sizeBytes,omitempty"` // Output only. The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.
	Filename string `json:"filename,omitempty"` // A possibly-hierarchical name used to refer to the spec from other specs.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
	Revisioncreatetime string `json:"revisionCreateTime,omitempty"` // Output only. Revision creation timestamp; when the represented revision was created.
	Contents string `json:"contents,omitempty"` // Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp; when the spec resource was created.
	Mimetype string `json:"mimeType,omitempty"` // A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include "application/vnd.apigee.proto", "application/vnd.apigee.openapi", and "application/vnd.apigee.graphql", with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
}

// ListApisResponse represents the ListApisResponse schema from the OpenAPI specification
type ListApisResponse struct {
	Apis []Api `json:"apis,omitempty"` // The APIs from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Details []GoogleProtobufAny `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
}

// TagApiSpecRevisionRequest represents the TagApiSpecRevisionRequest schema from the OpenAPI specification
type TagApiSpecRevisionRequest struct {
	Name string `json:"name"` // Required. The name of the spec to be tagged, including the revision ID.
	Tag string `json:"tag"` // Required. The tag to apply. The tag should be at most 40 characters, and match `[a-z][a-z0-9-]{3,39}`.
}

// GoogleProtobufAny represents the GoogleProtobufAny schema from the OpenAPI specification
type GoogleProtobufAny struct {
	TypeField string `json:"@type,omitempty"` // The type of the serialized message.
}

// ListArtifactsResponse represents the ListArtifactsResponse schema from the OpenAPI specification
type ListArtifactsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	Artifacts []Artifact `json:"artifacts,omitempty"` // The artifacts from the specified publisher.
}

// TagApiDeploymentRevisionRequest represents the TagApiDeploymentRevisionRequest schema from the OpenAPI specification
type TagApiDeploymentRevisionRequest struct {
	Name string `json:"name"` // Required. The name of the deployment to be tagged, including the revision ID.
	Tag string `json:"tag"` // Required. The tag to apply. The tag should be at most 40 characters, and match `[a-z][a-z0-9-]{3,39}`.
}

// ApiDeployment represents the ApiDeployment schema from the OpenAPI specification
type ApiDeployment struct {
	Revisioncreatetime string `json:"revisionCreateTime,omitempty"` // Output only. Revision creation timestamp; when the represented revision was created.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp; when the deployment resource was created.
	Description string `json:"description,omitempty"` // A detailed description.
	Displayname string `json:"displayName,omitempty"` // Human-meaningful name.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "registry.googleapis.com/" and cannot be changed.
	Revisionid string `json:"revisionId,omitempty"` // Output only. Immutable. The revision ID of the deployment. A new revision is committed whenever the deployment contents are changed. The format is an 8-character hexadecimal string.
	Revisionupdatetime string `json:"revisionUpdateTime,omitempty"` // Output only. Last update timestamp: when the represented revision was last modified.
	Accessguidance string `json:"accessGuidance,omitempty"` // Text briefly describing how to access the endpoint. Changes to this value will not affect the revision.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Apispecrevision string `json:"apiSpecRevision,omitempty"` // The full resource name (including revision id) of the spec of the API being served by the deployment. Changes to this value will update the revision. Format: apis/{api}/deployments/{deployment}
	Endpointuri string `json:"endpointUri,omitempty"` // The address where the deployment is serving. Changes to this value will update the revision.
	Externalchanneluri string `json:"externalChannelUri,omitempty"` // The address of the external channel of the API (e.g. the Developer Portal). Changes to this value will not affect the revision.
	Intendedaudience string `json:"intendedAudience,omitempty"` // Text briefly identifying the intended audience of the API. Changes to this value will not affect the revision.
	Name string `json:"name,omitempty"` // Resource name.
}

// ListApiDeploymentRevisionsResponse represents the ListApiDeploymentRevisionsResponse schema from the OpenAPI specification
type ListApiDeploymentRevisionsResponse struct {
	Apideployments []ApiDeployment `json:"apiDeployments,omitempty"` // The revisions of the deployment.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// RollbackApiSpecRequest represents the RollbackApiSpecRequest schema from the OpenAPI specification
type RollbackApiSpecRequest struct {
	Name string `json:"name"` // Required. The spec being rolled back.
	Revisionid string `json:"revisionId"` // Required. The revision ID to roll back to. It must be a revision of the same spec. Example: c7cfa2a8
}

// ApiVersion represents the ApiVersion schema from the OpenAPI specification
type ApiVersion struct {
	Name string `json:"name,omitempty"` // Resource name.
	State string `json:"state,omitempty"` // A user-definable description of the lifecycle phase of this API version. Format: free-form, but we expect single words that describe API maturity, e.g. "CONCEPT", "DESIGN", "DEVELOPMENT", "STAGING", "PRODUCTION", "DEPRECATED", "RETIRED".
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
	Description string `json:"description,omitempty"` // A detailed description.
	Displayname string `json:"displayName,omitempty"` // Human-meaningful name.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
}
