/*
api spec

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
)

// SubnetVpc struct for SubnetVpc
type SubnetVpc struct {
	Id *string `json:"id,omitempty"`
	Self *string `json:"self,omitempty"`
}

// NewSubnetVpc instantiates a new SubnetVpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetVpc() *SubnetVpc {
	this := SubnetVpc{}
	return &this
}

// NewSubnetVpcWithDefaults instantiates a new SubnetVpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetVpcWithDefaults() *SubnetVpc {
	this := SubnetVpc{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubnetVpc) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetVpc) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubnetVpc) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubnetVpc) SetId(v string) {
	o.Id = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SubnetVpc) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetVpc) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SubnetVpc) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *SubnetVpc) SetSelf(v string) {
	o.Self = &v
}

func (o SubnetVpc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableSubnetVpc struct {
	value *SubnetVpc
	isSet bool
}

func (v NullableSubnetVpc) Get() *SubnetVpc {
	return v.value
}

func (v *NullableSubnetVpc) Set(val *SubnetVpc) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetVpc(val *SubnetVpc) *NullableSubnetVpc {
	return &NullableSubnetVpc{value: val, isSet: true}
}

func (v NullableSubnetVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


