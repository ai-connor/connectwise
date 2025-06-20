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

// checks if the PageValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageValues{}

// PageValues struct for PageValues
type PageValues struct {
	Page NullableInt32 `json:"page,omitempty"`
	PageSize NullableInt32 `json:"pageSize,omitempty"`
	PageId NullableInt32 `json:"pageId,omitempty"`
}

// NewPageValues instantiates a new PageValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageValues() *PageValues {
	this := PageValues{}
	return &this
}

// NewPageValuesWithDefaults instantiates a new PageValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageValuesWithDefaults() *PageValues {
	this := PageValues{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PageValues) GetPage() int32 {
	if o == nil || IsNil(o.Page.Get()) {
		var ret int32
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PageValues) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *PageValues) HasPage() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableInt32 and assigns it to the Page field.
func (o *PageValues) SetPage(v int32) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *PageValues) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *PageValues) UnsetPage() {
	o.Page.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PageValues) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize.Get()) {
		var ret int32
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PageValues) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *PageValues) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt32 and assigns it to the PageSize field.
func (o *PageValues) SetPageSize(v int32) {
	o.PageSize.Set(&v)
}
// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *PageValues) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *PageValues) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetPageId returns the PageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PageValues) GetPageId() int32 {
	if o == nil || IsNil(o.PageId.Get()) {
		var ret int32
		return ret
	}
	return *o.PageId.Get()
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PageValues) GetPageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageId.Get(), o.PageId.IsSet()
}

// HasPageId returns a boolean if a field has been set.
func (o *PageValues) HasPageId() bool {
	if o != nil && o.PageId.IsSet() {
		return true
	}

	return false
}

// SetPageId gets a reference to the given NullableInt32 and assigns it to the PageId field.
func (o *PageValues) SetPageId(v int32) {
	o.PageId.Set(&v)
}
// SetPageIdNil sets the value for PageId to be an explicit nil
func (o *PageValues) SetPageIdNil() {
	o.PageId.Set(nil)
}

// UnsetPageId ensures that no value is present for PageId, not even an explicit nil
func (o *PageValues) UnsetPageId() {
	o.PageId.Unset()
}

func (o PageValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	if o.PageId.IsSet() {
		toSerialize["pageId"] = o.PageId.Get()
	}
	return toSerialize, nil
}

type NullablePageValues struct {
	value *PageValues
	isSet bool
}

func (v NullablePageValues) Get() *PageValues {
	return v.value
}

func (v *NullablePageValues) Set(val *PageValues) {
	v.value = val
	v.isSet = true
}

func (v NullablePageValues) IsSet() bool {
	return v.isSet
}

func (v *NullablePageValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageValues(val *PageValues) *NullablePageValues {
	return &NullablePageValues{value: val, isSet: true}
}

func (v NullablePageValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


