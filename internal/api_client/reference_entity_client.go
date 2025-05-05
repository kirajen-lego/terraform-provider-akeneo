package api_client

import (
  goakeneo "github.com/ezifyio/go-akeneo"
  "path"
)

const (
  ReferenceEntityBasePath = "/api/rest/v1/reference-entities"
)

// ReferenceEntityService is an interface for interfacing with the ReferenceEntity
type ReferenceEntityService interface {
  ListWithPagination(options any) ([]ReferenceEntity, goakeneo.Links, error)
  GetReferenceEntity(code string, options any) (*ReferenceEntity, error)
  GetReferenceEntityAttributes(code string, attributeCode string, options any) ([]ReferenceEntityAttributes, goakeneo.Links, error)
  GetReferenceEntityAttributesOption(code string, attributeCode string, optionCode string, options any) ([]ReferenceEntityAttributesOption, Links, error)
}

// ReferenceEntityOp handles communication with the ReferenceEntity related methods of the Akeneo API.
type ReferenceEntityOp struct {
  client *goakeneo.Client
}

// ListWithPagination lists ReferenceEntitys with pagination
func (c *ReferenceEntityOp) ListWithPagination(options any) ([]ReferenceEntity, goakeneo.Links, error) {
  ReferenceEntityResponse := new(ReferenceEntitysResponse)
  if err := c.client.GET(
    ReferenceEntityBasePath,
    options,
    nil,
    ReferenceEntityResponse,
  ); err != nil {
    return nil, goakeneo.Links{}, err
  }
  return ReferenceEntityResponse.Embedded.Items, ReferenceEntityResponse.Links, nil
}

// GetReferenceEntity gets an ReferenceEntity by code
func (c *ReferenceEntityOp) GetReferenceEntity(code string, options any) (*ReferenceEntity, error) {
  sourcePath := path.Join(ReferenceEntityBasePath, code)
  ReferenceEntity := new(ReferenceEntity)
  if err := c.client.GET(
    sourcePath,
    options,
    nil,
    ReferenceEntity,
  ); err != nil {
    return nil, err
  }
  return ReferenceEntity, nil
}

// GetReferenceEntityOptions gets an ReferenceEntity's options by code
func (c *ReferenceEntityOp) GetReferenceEntityOptions(code string, options any) ([]ReferenceEntityOption, goakeneo.Links, error) {
  sourcePath := path.Join(ReferenceEntityBasePath, code, "options")
  ReferenceEntityOptionsResponse := new(ReferenceEntityOptionsResponse)
  if err := c.client.GET(
    sourcePath,
    options,
    nil,
    ReferenceEntityOptionsResponse,
  ); err != nil {
    return nil, goakeneo.Links{}, err
  }
  return ReferenceEntityOptionsResponse.Embedded.Items, ReferenceEntityOptionsResponse.Links, nil
}

// ReferenceEntitysResponse is the struct for a akeneo ReferenceEntitys response
type ReferenceEntitysResponse struct {
  Links       goakeneo.Links       `json:"_links" mapstructure:"_links"`
  CurrentPage int                  `json:"current_page" mapstructure:"current_page"`
  Embedded    ReferenceEntityItems `json:"_embedded" mapstructure:"_embedded"`
}

type ReferenceEntityItems struct {
  Items []ReferenceEntity `json:"items" mapstructure:"items"`
}

// ReferenceEntityOptionsResponse is the struct for a akeneo ReferenceEntity options response
type ReferenceEntityOptionsResponse struct {
  Links       goakeneo.Links             `json:"_links" mapstructure:"_links"`
  CurrentPage int                        `json:"current_page" mapstructure:"current_page"`
  Embedded    ReferenceEntityOptionItems `json:"_embedded" mapstructure:"_embedded"`
}

type ReferenceEntityOptionItems struct {
  Items []ReferenceEntityOption `json:"items" mapstructure:"items"`
}
