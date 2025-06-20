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

// checks if the BundleResultsCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleResultsCollection{}

// BundleResultsCollection struct for BundleResultsCollection
type BundleResultsCollection struct {
	Results []BundleResult `json:"results,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewBundleResultsCollection instantiates a new BundleResultsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleResultsCollection() *BundleResultsCollection {
	this := BundleResultsCollection{}
	return &this
}

// NewBundleResultsCollectionWithDefaults instantiates a new BundleResultsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleResultsCollectionWithDefaults() *BundleResultsCollection {
	this := BundleResultsCollection{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BundleResultsCollection) GetResults() []BundleResult {
	if o == nil || IsNil(o.Results) {
		var ret []BundleResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleResultsCollection) GetResultsOk() ([]BundleResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BundleResultsCollection) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BundleResult and assigns it to the Results field.
func (o *BundleResultsCollection) SetResults(v []BundleResult) {
	o.Results = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BundleResultsCollection) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleResultsCollection) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BundleResultsCollection) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *BundleResultsCollection) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o BundleResultsCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleResultsCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableBundleResultsCollection struct {
	value *BundleResultsCollection
	isSet bool
}

func (v NullableBundleResultsCollection) Get() *BundleResultsCollection {
	return v.value
}

func (v *NullableBundleResultsCollection) Set(val *BundleResultsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleResultsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleResultsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleResultsCollection(val *BundleResultsCollection) *NullableBundleResultsCollection {
	return &NullableBundleResultsCollection{value: val, isSet: true}
}

func (v NullableBundleResultsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleResultsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


