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

// Vpc struct for Vpc
type Vpc struct {
	CloudProvider *CloudProviders `json:"cloud_provider,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Cidr *string `json:"cidr,omitempty"`
	Id *string `json:"id,omitempty"`
	Region *AwsRegions `json:"region,omitempty"`
	Self *string `json:"self,omitempty"`
}

// NewVpc instantiates a new Vpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpc() *Vpc {
	this := Vpc{}
	return &this
}

// NewVpcWithDefaults instantiates a new Vpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpcWithDefaults() *Vpc {
	this := Vpc{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *Vpc) GetCloudProvider() CloudProviders {
	if o == nil || o.CloudProvider == nil {
		var ret CloudProviders
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetCloudProviderOk() (*CloudProviders, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *Vpc) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given CloudProviders and assigns it to the CloudProvider field.
func (o *Vpc) SetCloudProvider(v CloudProviders) {
	o.CloudProvider = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *Vpc) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *Vpc) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *Vpc) SetAccountName(v string) {
	o.AccountName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vpc) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vpc) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vpc) SetName(v string) {
	o.Name = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *Vpc) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *Vpc) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *Vpc) SetCidr(v string) {
	o.Cidr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Vpc) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Vpc) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Vpc) SetId(v string) {
	o.Id = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Vpc) GetRegion() AwsRegions {
	if o == nil || o.Region == nil {
		var ret AwsRegions
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetRegionOk() (*AwsRegions, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Vpc) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given AwsRegions and assigns it to the Region field.
func (o *Vpc) SetRegion(v AwsRegions) {
	o.Region = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Vpc) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpc) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Vpc) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *Vpc) SetSelf(v string) {
	o.Self = &v
}

func (o Vpc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if o.AccountName != nil {
		toSerialize["account_name"] = o.AccountName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableVpc struct {
	value *Vpc
	isSet bool
}

func (v NullableVpc) Get() *Vpc {
	return v.value
}

func (v *NullableVpc) Set(val *Vpc) {
	v.value = val
	v.isSet = true
}

func (v NullableVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpc(val *Vpc) *NullableVpc {
	return &NullableVpc{value: val, isSet: true}
}

func (v NullableVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


