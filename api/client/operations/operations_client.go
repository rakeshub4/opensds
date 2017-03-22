package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateShare create a share in specified backend resource
*/
func (a *Client) CreateShare(params *CreateShareParams) (*CreateShareOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateShareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateShare",
		Method:             "POST",
		PathPattern:        "/api/v1/shares/{resourceType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateShareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateShareOK), nil

}

/*
CreateVolume create a volume in specified backend resource
*/
func (a *Client) CreateVolume(params *CreateVolumeParams) (*CreateVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateVolume",
		Method:             "POST",
		PathPattern:        "/api/v1/volumes/{resourceType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateVolumeOK), nil

}

/*
DeleteShare delete specified share in specified backend resource
*/
func (a *Client) DeleteShare(params *DeleteShareParams) (*DeleteShareOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteShareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteShare",
		Method:             "DELETE",
		PathPattern:        "/api/v1/shares/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteShareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteShareOK), nil

}

/*
DeleteVolume delete specified volume in specified backend resource
*/
func (a *Client) DeleteVolume(params *DeleteVolumeParams) (*DeleteVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteVolume",
		Method:             "DELETE",
		PathPattern:        "/api/v1/volumes/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteVolumeOK), nil

}

/*
GetShare get specified share in specified backend resource
*/
func (a *Client) GetShare(params *GetShareParams) (*GetShareOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetShare",
		Method:             "GET",
		PathPattern:        "/api/v1/shares/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetShareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetShareOK), nil

}

/*
GetVersionv1 show API version
*/
func (a *Client) GetVersionv1(params *GetVersionv1Params) (*GetVersionv1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionv1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVersionv1",
		Method:             "GET",
		PathPattern:        "/api/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVersionv1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersionv1OK), nil

}

/*
GetVolume get specified volume in specified backend resource
*/
func (a *Client) GetVolume(params *GetVolumeParams) (*GetVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVolume",
		Method:             "GET",
		PathPattern:        "/api/v1/volumes/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVolumeOK), nil

}

/*
ListShareResources list share backend resource types
*/
func (a *Client) ListShareResources(params *ListShareResourcesParams) (*ListShareResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListShareResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListShareResources",
		Method:             "GET",
		PathPattern:        "/api/v1/shares",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListShareResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListShareResourcesOK), nil

}

/*
ListShares list shares in specified backend resource
*/
func (a *Client) ListShares(params *ListSharesParams) (*ListSharesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSharesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListShares",
		Method:             "GET",
		PathPattern:        "/api/v1/shares/{resourceType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSharesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSharesOK), nil

}

/*
ListVersions get available API versions
*/
func (a *Client) ListVersions(params *ListVersionsParams) (*ListVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVersions",
		Method:             "GET",
		PathPattern:        "/api/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVersionsOK), nil

}

/*
ListVolumeResources list volume backend resource types
*/
func (a *Client) ListVolumeResources(params *ListVolumeResourcesParams) (*ListVolumeResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVolumeResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVolumeResources",
		Method:             "GET",
		PathPattern:        "/api/v1/volumes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVolumeResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVolumeResourcesOK), nil

}

/*
ListVolumes list volumes in specified backend resource
*/
func (a *Client) ListVolumes(params *ListVolumesParams) (*ListVolumesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVolumesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVolumes",
		Method:             "GET",
		PathPattern:        "/api/v1/volumes/{resourceType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVolumesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVolumesOK), nil

}

/*
OperateVolume opreate specified volume in specified backend resource
*/
func (a *Client) OperateVolume(params *OperateVolumeParams) (*OperateVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOperateVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OperateVolume",
		Method:             "POST",
		PathPattern:        "/api/v1/volumes/action/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OperateVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OperateVolumeOK), nil

}

/*
UpdateShare update specified share in specified backend resource
*/
func (a *Client) UpdateShare(params *UpdateShareParams) (*UpdateShareOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateShareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateShare",
		Method:             "PUT",
		PathPattern:        "/api/v1/shares/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateShareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateShareOK), nil

}

/*
UpdateVolume update specified volume in specified backend resource
*/
func (a *Client) UpdateVolume(params *UpdateVolumeParams) (*UpdateVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVolumeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateVolume",
		Method:             "PUT",
		PathPattern:        "/api/v1/volumes/{resourceType}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVolumeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateVolumeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}