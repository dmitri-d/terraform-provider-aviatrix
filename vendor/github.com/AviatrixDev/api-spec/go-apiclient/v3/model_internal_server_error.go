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

// InternalServerError struct for InternalServerError
type InternalServerError struct {
	Description string `json:"description"`
}

// NewInternalServerError instantiates a new InternalServerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerError(description string) *InternalServerError {
	this := InternalServerError{}
	this.Description = description
	return &this
}

// NewInternalServerErrorWithDefaults instantiates a new InternalServerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerErrorWithDefaults() *InternalServerError {
	this := InternalServerError{}
	return &this
}

// GetDescription returns the Description field value
func (o *InternalServerError) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InternalServerError) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InternalServerError) SetDescription(v string) {
	o.Description = v
}

func (o InternalServerError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInternalServerError struct {
	value *InternalServerError
	isSet bool
}

func (v NullableInternalServerError) Get() *InternalServerError {
	return v.value
}

func (v *NullableInternalServerError) Set(val *InternalServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerError(val *InternalServerError) *NullableInternalServerError {
	return &NullableInternalServerError{value: val, isSet: true}
}

func (v NullableInternalServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


