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

// checks if the CompanyStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyStatus{}

// CompanyStatus struct for CompanyStatus
type CompanyStatus struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 50;
	Name string `json:"name"`
	DefaultFlag NullableBool `json:"defaultFlag,omitempty"`
	InactiveFlag NullableBool `json:"inactiveFlag,omitempty"`
	NotifyFlag NullableBool `json:"notifyFlag,omitempty"`
	DisallowSavingFlag NullableBool `json:"disallowSavingFlag,omitempty"`
	//  Max length: 500;
	NotificationMessage *string `json:"notificationMessage,omitempty"`
	CustomNoteFlag NullableBool `json:"customNoteFlag,omitempty"`
	CancelOpenTracksFlag NullableBool `json:"cancelOpenTracksFlag,omitempty"`
	Track *TrackReference `json:"track,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _CompanyStatus CompanyStatus

// NewCompanyStatus instantiates a new CompanyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyStatus(name string) *CompanyStatus {
	this := CompanyStatus{}
	this.Name = name
	return &this
}

// NewCompanyStatusWithDefaults instantiates a new CompanyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyStatusWithDefaults() *CompanyStatus {
	this := CompanyStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyStatus) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyStatus) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyStatus) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CompanyStatus) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CompanyStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompanyStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompanyStatus) SetName(v string) {
	o.Name = v
}

// GetDefaultFlag returns the DefaultFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyStatus) GetDefaultFlag() bool {
	if o == nil || IsNil(o.DefaultFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultFlag.Get()
}

// GetDefaultFlagOk returns a tuple with the DefaultFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyStatus) GetDefaultFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFlag.Get(), o.DefaultFlag.IsSet()
}

// HasDefaultFlag returns a boolean if a field has been set.
func (o *CompanyStatus) HasDefaultFlag() bool {
	if o != nil && o.DefaultFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultFlag gets a reference to the given NullableBool and assigns it to the DefaultFlag field.
func (o *CompanyStatus) SetDefaultFlag(v bool) {
	o.DefaultFlag.Set(&v)
}
// SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil
func (o *CompanyStatus) SetDefaultFlagNil() {
	o.DefaultFlag.Set(nil)
}

// UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
func (o *CompanyStatus) UnsetDefaultFlag() {
	o.DefaultFlag.Unset()
}

// GetInactiveFlag returns the InactiveFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyStatus) GetInactiveFlag() bool {
	if o == nil || IsNil(o.InactiveFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.InactiveFlag.Get()
}

// GetInactiveFlagOk returns a tuple with the InactiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyStatus) GetInactiveFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InactiveFlag.Get(), o.InactiveFlag.IsSet()
}

// HasInactiveFlag returns a boolean if a field has been set.
func (o *CompanyStatus) HasInactiveFlag() bool {
	if o != nil && o.InactiveFlag.IsSet() {
		return true
	}

	return false
}

// SetInactiveFlag gets a reference to the given NullableBool and assigns it to the InactiveFlag field.
func (o *CompanyStatus) SetInactiveFlag(v bool) {
	o.InactiveFlag.Set(&v)
}
// SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil
func (o *CompanyStatus) SetInactiveFlagNil() {
	o.InactiveFlag.Set(nil)
}

// UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
func (o *CompanyStatus) UnsetInactiveFlag() {
	o.InactiveFlag.Unset()
}

// GetNotifyFlag returns the NotifyFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyStatus) GetNotifyFlag() bool {
	if o == nil || IsNil(o.NotifyFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.NotifyFlag.Get()
}

// GetNotifyFlagOk returns a tuple with the NotifyFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyStatus) GetNotifyFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyFlag.Get(), o.NotifyFlag.IsSet()
}

// HasNotifyFlag returns a boolean if a field has been set.
func (o *CompanyStatus) HasNotifyFlag() bool {
	if o != nil && o.NotifyFlag.IsSet() {
		return true
	}

	return false
}

// SetNotifyFlag gets a reference to the given NullableBool and assigns it to the NotifyFlag field.
func (o *CompanyStatus) SetNotifyFlag(v bool) {
	o.NotifyFlag.Set(&v)
}
// SetNotifyFlagNil sets the value for NotifyFlag to be an explicit nil
func (o *CompanyStatus) SetNotifyFlagNil() {
	o.NotifyFlag.Set(nil)
}

// UnsetNotifyFlag ensures that no value is present for NotifyFlag, not even an explicit nil
func (o *CompanyStatus) UnsetNotifyFlag() {
	o.NotifyFlag.Unset()
}

// GetDisallowSavingFlag returns the DisallowSavingFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyStatus) GetDisallowSavingFlag() bool {
	if o == nil || IsNil(o.DisallowSavingFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DisallowSavingFlag.Get()
}

// GetDisallowSavingFlagOk returns a tuple with the DisallowSavingFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyStatus) GetDisallowSavingFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisallowSavingFlag.Get(), o.DisallowSavingFlag.IsSet()
}

// HasDisallowSavingFlag returns a boolean if a field has been set.
func (o *CompanyStatus) HasDisallowSavingFlag() bool {
	if o != nil && o.DisallowSavingFlag.IsSet() {
		return true
	}

	return false
}

// SetDisallowSavingFlag gets a reference to the given NullableBool and assigns it to the DisallowSavingFlag field.
func (o *CompanyStatus) SetDisallowSavingFlag(v bool) {
	o.DisallowSavingFlag.Set(&v)
}
// SetDisallowSavingFlagNil sets the value for DisallowSavingFlag to be an explicit nil
func (o *CompanyStatus) SetDisallowSavingFlagNil() {
	o.DisallowSavingFlag.Set(nil)
}

// UnsetDisallowSavingFlag ensures that no value is present for DisallowSavingFlag, not even an explicit nil
func (o *CompanyStatus) UnsetDisallowSavingFlag() {
	o.DisallowSavingFlag.Unset()
}

// GetNotificationMessage returns the NotificationMessage field value if set, zero value otherwise.
func (o *CompanyStatus) GetNotificationMessage() string {
	if o == nil || IsNil(o.NotificationMessage) {
		var ret string
		return ret
	}
	return *o.NotificationMessage
}

// GetNotificationMessageOk returns a tuple with the NotificationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyStatus) GetNotificationMessageOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationMessage) {
		return nil, false
	}
	return o.NotificationMessage, true
}

// HasNotificationMessage returns a boolean if a field has been set.
func (o *CompanyStatus) HasNotificationMessage() bool {
	if o != nil && !IsNil(o.NotificationMessage) {
		return true
	}

	return false
}

// SetNotificationMessage gets a reference to the given string and assigns it to the NotificationMessage field.
func (o *CompanyStatus) SetNotificationMessage(v string) {
	o.NotificationMessage = &v
}

// GetCustomNoteFlag returns the CustomNoteFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyStatus) GetCustomNoteFlag() bool {
	if o == nil || IsNil(o.CustomNoteFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.CustomNoteFlag.Get()
}

// GetCustomNoteFlagOk returns a tuple with the CustomNoteFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyStatus) GetCustomNoteFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomNoteFlag.Get(), o.CustomNoteFlag.IsSet()
}

// HasCustomNoteFlag returns a boolean if a field has been set.
func (o *CompanyStatus) HasCustomNoteFlag() bool {
	if o != nil && o.CustomNoteFlag.IsSet() {
		return true
	}

	return false
}

// SetCustomNoteFlag gets a reference to the given NullableBool and assigns it to the CustomNoteFlag field.
func (o *CompanyStatus) SetCustomNoteFlag(v bool) {
	o.CustomNoteFlag.Set(&v)
}
// SetCustomNoteFlagNil sets the value for CustomNoteFlag to be an explicit nil
func (o *CompanyStatus) SetCustomNoteFlagNil() {
	o.CustomNoteFlag.Set(nil)
}

// UnsetCustomNoteFlag ensures that no value is present for CustomNoteFlag, not even an explicit nil
func (o *CompanyStatus) UnsetCustomNoteFlag() {
	o.CustomNoteFlag.Unset()
}

// GetCancelOpenTracksFlag returns the CancelOpenTracksFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyStatus) GetCancelOpenTracksFlag() bool {
	if o == nil || IsNil(o.CancelOpenTracksFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.CancelOpenTracksFlag.Get()
}

// GetCancelOpenTracksFlagOk returns a tuple with the CancelOpenTracksFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyStatus) GetCancelOpenTracksFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelOpenTracksFlag.Get(), o.CancelOpenTracksFlag.IsSet()
}

// HasCancelOpenTracksFlag returns a boolean if a field has been set.
func (o *CompanyStatus) HasCancelOpenTracksFlag() bool {
	if o != nil && o.CancelOpenTracksFlag.IsSet() {
		return true
	}

	return false
}

// SetCancelOpenTracksFlag gets a reference to the given NullableBool and assigns it to the CancelOpenTracksFlag field.
func (o *CompanyStatus) SetCancelOpenTracksFlag(v bool) {
	o.CancelOpenTracksFlag.Set(&v)
}
// SetCancelOpenTracksFlagNil sets the value for CancelOpenTracksFlag to be an explicit nil
func (o *CompanyStatus) SetCancelOpenTracksFlagNil() {
	o.CancelOpenTracksFlag.Set(nil)
}

// UnsetCancelOpenTracksFlag ensures that no value is present for CancelOpenTracksFlag, not even an explicit nil
func (o *CompanyStatus) UnsetCancelOpenTracksFlag() {
	o.CancelOpenTracksFlag.Unset()
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *CompanyStatus) GetTrack() TrackReference {
	if o == nil || IsNil(o.Track) {
		var ret TrackReference
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyStatus) GetTrackOk() (*TrackReference, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *CompanyStatus) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given TrackReference and assigns it to the Track field.
func (o *CompanyStatus) SetTrack(v TrackReference) {
	o.Track = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *CompanyStatus) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyStatus) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *CompanyStatus) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *CompanyStatus) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o CompanyStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.DefaultFlag.IsSet() {
		toSerialize["defaultFlag"] = o.DefaultFlag.Get()
	}
	if o.InactiveFlag.IsSet() {
		toSerialize["inactiveFlag"] = o.InactiveFlag.Get()
	}
	if o.NotifyFlag.IsSet() {
		toSerialize["notifyFlag"] = o.NotifyFlag.Get()
	}
	if o.DisallowSavingFlag.IsSet() {
		toSerialize["disallowSavingFlag"] = o.DisallowSavingFlag.Get()
	}
	if !IsNil(o.NotificationMessage) {
		toSerialize["notificationMessage"] = o.NotificationMessage
	}
	if o.CustomNoteFlag.IsSet() {
		toSerialize["customNoteFlag"] = o.CustomNoteFlag.Get()
	}
	if o.CancelOpenTracksFlag.IsSet() {
		toSerialize["cancelOpenTracksFlag"] = o.CancelOpenTracksFlag.Get()
	}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *CompanyStatus) UnmarshalJSON(data []byte) (err error) {
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

	varCompanyStatus := _CompanyStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyStatus)

	if err != nil {
		return err
	}

	*o = CompanyStatus(varCompanyStatus)

	return err
}

type NullableCompanyStatus struct {
	value *CompanyStatus
	isSet bool
}

func (v NullableCompanyStatus) Get() *CompanyStatus {
	return v.value
}

func (v *NullableCompanyStatus) Set(val *CompanyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyStatus(val *CompanyStatus) *NullableCompanyStatus {
	return &NullableCompanyStatus{value: val, isSet: true}
}

func (v NullableCompanyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


