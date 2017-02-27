package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// ListVolumesURL generates an URL for the list volumes operation
type ListVolumesURL struct {
	ResourceType string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListVolumesURL) WithBasePath(bp string) *ListVolumesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ListVolumesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ListVolumesURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/api/v1/volumes/{resourceType}"

	resourceType := o.ResourceType
	if resourceType != "" {
		_path = strings.Replace(_path, "{resourceType}", resourceType, -1)
	} else {
		return nil, errors.New("ResourceType is required on ListVolumesURL")
	}
	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ListVolumesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ListVolumesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ListVolumesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ListVolumesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ListVolumesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ListVolumesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
