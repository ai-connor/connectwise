/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"encoding/json"
)

// checks if the BundleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleRequest{}

// BundleRequest struct for BundleRequest
type BundleRequest struct {
	SequenceNumber *int32 `json:"sequenceNumber,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Version *string `json:"version,omitempty"`
	ApiRequest *ApiRequest `json:"apiRequest,omitempty"`
}

// NewBundleRequest instantiates a new BundleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleRequest() *BundleRequest {
	this := BundleRequest{}
	return &this
}

// NewBundleRequestWithDefaults instantiates a new BundleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleRequestWithDefaults() *BundleRequest {
	this := BundleRequest{}
	return &this
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *BundleRequest) GetSequenceNumber() int32 {
	if o == nil || IsNil(o.SequenceNumber) {
		var ret int32
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetSequenceNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SequenceNumber) {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *BundleRequest) HasSequenceNumber() bool {
	if o != nil && !IsNil(o.SequenceNumber) {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int32 and assigns it to the SequenceNumber field.
func (o *BundleRequest) SetSequenceNumber(v int32) {
	o.SequenceNumber = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BundleRequest) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BundleRequest) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BundleRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BundleRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BundleRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BundleRequest) SetVersion(v string) {
	o.Version = &v
}

// GetApiRequest returns the ApiRequest field value if set, zero value otherwise.
func (o *BundleRequest) GetApiRequest() ApiRequest {
	if o == nil || IsNil(o.ApiRequest) {
		var ret ApiRequest
		return ret
	}
	return *o.ApiRequest
}

// GetApiRequestOk returns a tuple with the ApiRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleRequest) GetApiRequestOk() (*ApiRequest, bool) {
	if o == nil || IsNil(o.ApiRequest) {
		return nil, false
	}
	return o.ApiRequest, true
}

// HasApiRequest returns a boolean if a field has been set.
func (o *BundleRequest) HasApiRequest() bool {
	if o != nil && !IsNil(o.ApiRequest) {
		return true
	}

	return false
}

// SetApiRequest gets a reference to the given ApiRequest and assigns it to the ApiRequest field.
func (o *BundleRequest) SetApiRequest(v ApiRequest) {
	o.ApiRequest = &v
}

func (o BundleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SequenceNumber) {
		toSerialize["sequenceNumber"] = o.SequenceNumber
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ApiRequest) {
		toSerialize["apiRequest"] = o.ApiRequest
	}
	return toSerialize, nil
}

type NullableBundleRequest struct {
	value *BundleRequest
	isSet bool
}

func (v NullableBundleRequest) Get() *BundleRequest {
	return v.value
}

func (v *NullableBundleRequest) Set(val *BundleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleRequest(val *BundleRequest) *NullableBundleRequest {
	return &NullableBundleRequest{value: val, isSet: true}
}

func (v NullableBundleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


