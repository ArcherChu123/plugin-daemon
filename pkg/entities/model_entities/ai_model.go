package model_entities

import "github.com/ArcherChu123/plugin-daemon/pkg/entities/plugin_entities"

type GetModelSchemasResponse struct {
	ModelSchema *plugin_entities.ModelDeclaration `json:"model_schema" validate:"omitempty"`
}
