package cyclonedx

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Plugin initializes the steampipe plugin with the dynamic table map.
func Plugin(ctx context.Context) *plugin.Plugin {
	// List of table names and their corresponding table creation functions
	tables := map[string]func(ctx context.Context, tableName string) *plugin.Table{
		"cyclonedx_bom_info_file": tableCyclonedxBomInfo,
		"cyclonedx_metadata":      tableCyclonedxMetadata,
		"cyclonedx_tools":         tableCyclonedxTools,
		"cyclonedx_authors":       tableCyclonedxAuthors,
		"cyclonedx_lifecycles":    tableCyclonedxLifecycles,
		"cyclonedx_components":    tableCyclonedxComponents,
		"cyclonedx_dependencies":  tableCyclonedxDependencies,
	}

	// Dynamically populate the TableMap
	tableMap := make(map[string]*plugin.Table)
	for tableName, tableFunc := range tables {
		tableMap[tableName] = tableFunc(ctx, tableName)
	}

	return &plugin.Plugin{
		Name: "steampipe-plugin-cyclonedx",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: tableMap,
	}
}
