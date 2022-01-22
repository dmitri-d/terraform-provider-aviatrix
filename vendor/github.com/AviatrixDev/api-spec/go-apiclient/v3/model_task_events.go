/*
api spec

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
	"time"
)

// TaskEvents struct for TaskEvents
type TaskEvents struct {
	Time *time.Time `json:"time,omitempty"`
	Event *string `json:"event,omitempty"`
}

// NewTaskEvents instantiates a new TaskEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskEvents() *TaskEvents {
	this := TaskEvents{}
	return &this
}

// NewTaskEventsWithDefaults instantiates a new TaskEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskEventsWithDefaults() *TaskEvents {
	this := TaskEvents{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *TaskEvents) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvents) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *TaskEvents) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *TaskEvents) SetTime(v time.Time) {
	o.Time = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TaskEvents) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvents) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TaskEvents) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *TaskEvents) SetEvent(v string) {
	o.Event = &v
}

func (o TaskEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}

type NullableTaskEvents struct {
	value *TaskEvents
	isSet bool
}

func (v NullableTaskEvents) Get() *TaskEvents {
	return v.value
}

func (v *NullableTaskEvents) Set(val *TaskEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskEvents(val *TaskEvents) *NullableTaskEvents {
	return &NullableTaskEvents{value: val, isSet: true}
}

func (v NullableTaskEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

