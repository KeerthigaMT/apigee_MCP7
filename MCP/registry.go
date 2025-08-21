package main

import (
	"github.com/registry-api/mcp-server/config"
	"github.com/registry-api/mcp-server/models"
	tools_registry "github.com/registry-api/mcp-server/tools/registry"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_registry.CreateRegistry_deleteapispecTool(cfg),
		tools_registry.CreateRegistry_getapispecTool(cfg),
		tools_registry.CreateRegistry_updateapispecTool(cfg),
		tools_registry.CreateRegistry_getartifactcontentsTool(cfg),
		tools_registry.CreateRegistry_deleteapiversionTool(cfg),
		tools_registry.CreateRegistry_getapiversionTool(cfg),
		tools_registry.CreateRegistry_updateapiversionTool(cfg),
		tools_registry.CreateRegistry_createapispecTool(cfg),
		tools_registry.CreateRegistry_listapispecsTool(cfg),
		tools_registry.CreateRegistry_rollbackapispecTool(cfg),
		tools_registry.CreateRegistry_listartifactsTool(cfg),
		tools_registry.CreateRegistry_createartifactTool(cfg),
		tools_registry.CreateRegistry_deleteapispecrevisionTool(cfg),
		tools_registry.CreateRegistry_deleteapideploymentTool(cfg),
		tools_registry.CreateRegistry_getapideploymentTool(cfg),
		tools_registry.CreateRegistry_updateapideploymentTool(cfg),
		tools_registry.CreateRegistry_listapiversionsTool(cfg),
		tools_registry.CreateRegistry_createapiversionTool(cfg),
		tools_registry.CreateRegistry_deleteapiTool(cfg),
		tools_registry.CreateRegistry_getapiTool(cfg),
		tools_registry.CreateRegistry_updateapiTool(cfg),
		tools_registry.CreateRegistry_listapideploymentrevisionsTool(cfg),
		tools_registry.CreateRegistry_tagapideploymentrevisionTool(cfg),
		tools_registry.CreateRegistry_listapispecrevisionsTool(cfg),
		tools_registry.CreateRegistry_deleteartifactTool(cfg),
		tools_registry.CreateRegistry_getartifactTool(cfg),
		tools_registry.CreateRegistry_replaceartifactTool(cfg),
		tools_registry.CreateRegistry_deleteapideploymentrevisionTool(cfg),
		tools_registry.CreateRegistry_rollbackapideploymentTool(cfg),
		tools_registry.CreateRegistry_listapisTool(cfg),
		tools_registry.CreateRegistry_createapiTool(cfg),
		tools_registry.CreateRegistry_tagapispecrevisionTool(cfg),
		tools_registry.CreateRegistry_getapispeccontentsTool(cfg),
		tools_registry.CreateRegistry_listapideploymentsTool(cfg),
		tools_registry.CreateRegistry_createapideploymentTool(cfg),
	}
}
