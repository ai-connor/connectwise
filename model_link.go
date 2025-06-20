/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Link type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Link{}

// Link struct for Link
type Link struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	TableReferenceId NullableInt32 `json:"tableReferenceId,omitempty"`
	//  Max length: 1000;
	Url *string `json:"url,omitempty"`
	ScreenLink NullableString `json:"screenLink,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _Link Link

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink(name string) *Link {
	this := Link{}
	this.Name = name
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Link) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Link) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Link) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Link) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Link) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Link) SetName(v string) {
	o.Name = v
}

// GetTableReferenceId returns the TableReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Link) GetTableReferenceId() int32 {
	if o == nil || IsNil(o.TableReferenceId.Get()) {
		var ret int32
		return ret
	}
	return *o.TableReferenceId.Get()
}

// GetTableReferenceIdOk returns a tuple with the TableReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Link) GetTableReferenceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TableReferenceId.Get(), o.TableReferenceId.IsSet()
}

// HasTableReferenceId returns a boolean if a field has been set.
func (o *Link) HasTableReferenceId() bool {
	if o != nil && o.TableReferenceId.IsSet() {
		return true
	}

	return false
}

// SetTableReferenceId gets a reference to the given NullableInt32 and assigns it to the TableReferenceId field.
func (o *Link) SetTableReferenceId(v int32) {
	o.TableReferenceId.Set(&v)
}
// SetTableReferenceIdNil sets the value for TableReferenceId to be an explicit nil
func (o *Link) SetTableReferenceIdNil() {
	o.TableReferenceId.Set(nil)
}

// UnsetTableReferenceId ensures that no value is present for TableReferenceId, not even an explicit nil
func (o *Link) UnsetTableReferenceId() {
	o.TableReferenceId.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Link) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Link) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Link) SetUrl(v string) {
	o.Url = &v
}

// GetScreenLink returns the ScreenLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Link) GetScreenLink() string {
	if o == nil || IsNil(o.ScreenLink.Get()) {
		var ret string
		return ret
	}
	return *o.ScreenLink.Get()
}

// GetScreenLinkOk returns a tuple with the ScreenLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Link) GetScreenLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScreenLink.Get(), o.ScreenLink.IsSet()
}

// HasScreenLink returns a boolean if a field has been set.
func (o *Link) HasScreenLink() bool {
	if o != nil && o.ScreenLink.IsSet() {
		return true
	}

	return false
}

// SetScreenLink gets a reference to the given NullableString and assigns it to the ScreenLink field.
func (o *Link) SetScreenLink(v string) {
	o.ScreenLink.Set(&v)
}
// SetScreenLinkNil sets the value for ScreenLink to be an explicit nil
func (o *Link) SetScreenLinkNil() {
	o.ScreenLink.Set(nil)
}

// UnsetScreenLink ensures that no value is present for ScreenLink, not even an explicit nil
func (o *Link) UnsetScreenLink() {
	o.ScreenLink.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *Link) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *Link) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *Link) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o Link) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Link) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.TableReferenceId.IsSet() {
		toSerialize["tableReferenceId"] = o.TableReferenceId.Get()
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if o.ScreenLink.IsSet() {
		toSerialize["screenLink"] = o.ScreenLink.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *Link) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLink := _Link{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLink)

	if err != nil {
		return err
	}

	*o = Link(varLink)

	return err
}

type NullableLink struct {
	value *Link
	isSet bool
}

func (v NullableLink) Get() *Link {
	return v.value
}

func (v *NullableLink) Set(val *Link) {
	v.value = val
	v.isSet = true
}

func (v NullableLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLink(val *Link) *NullableLink {
	return &NullableLink{value: val, isSet: true}
}

func (v NullableLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


