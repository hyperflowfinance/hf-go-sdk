package codegen

import (
	"encoding/json"
	"fmt"
	"io"
)

type Schema struct {
	QueryType        *RootType `json:"queryType"`
	MutationType     *RootType `json:"mutationType"`
	SubscriptionType *RootType `json:"subscriptionType"`
	Types            []Type    `json:"types"`
}

type RootType struct {
	Name string `json:"name"`
}

type Type struct {
	Kind          string      `json:"kind"`
	Name          string      `json:"name"`
	Fields        []Field     `json:"fields"`
	InputFields   []Input     `json:"inputFields"`
	EnumValues    []EnumValue `json:"enumValues"`
	PossibleTypes []TypeRef   `json:"possibleTypes"`
}

type Field struct {
	Name string  `json:"name"`
	Type TypeRef `json:"type"`
	Args []Input `json:"args"`
}

type Input struct {
	Name         string  `json:"name"`
	Type         TypeRef `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

type EnumValue struct {
	Name string `json:"name"`
}

type TypeRef struct {
	Kind   string   `json:"kind"`
	Name   string   `json:"name"`
	OfType *TypeRef `json:"ofType"`
}

type envelope struct {
	Data struct {
		Schema Schema `json:"__schema"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func ParseSchema(reader io.Reader) (*Schema, error) {
	data, err := io.ReadAll(io.LimitReader(reader, (16<<20)+1))
	if err != nil {
		return nil, fmt.Errorf("read introspection response: %w", err)
	}
	if len(data) > 16<<20 {
		return nil, fmt.Errorf("introspection response exceeds 16 MiB")
	}
	var response envelope
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("decode introspection response: %w", err)
	}
	if len(response.Errors) != 0 {
		return nil, fmt.Errorf("schema introspection failed: %s", response.Errors[0].Message)
	}
	if response.Data.Schema.QueryType == nil || response.Data.Schema.QueryType.Name == "" {
		return nil, fmt.Errorf("introspection response has no query root")
	}
	return &response.Data.Schema, nil
}
