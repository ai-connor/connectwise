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

// checks if the UserDefinedFieldInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDefinedFieldInfo{}

// UserDefinedFieldInfo struct for UserDefinedFieldInfo
type UserDefinedFieldInfo struct {
	// ID of the custom user defined field
	Id *int32 `json:"id,omitempty"`
	// Id of the Pod where the custom field will be placed
	PodId NullableInt32 `json:"podId,omitempty"`
	// Field caption
	Caption *string `json:"caption,omitempty"`
	// Must be between 1 and 500.  This defines the order in which the custom fields will appear
	SequenceNumber NullableInt32 `json:"sequenceNumber,omitempty"`
	// Help text to accompany the custom field
	HelpText *string `json:"helpText,omitempty"`
	FieldTypeIdentifier NullableString `json:"fieldTypeIdentifier,omitempty"`
	// Only valid for Number or percent
	NumberDecimals NullableInt32 `json:"numberDecimals,omitempty"`
	EntryTypeIdentifier NullableString `json:"entryTypeIdentifier,omitempty"`
	RequiredFlag NullableBool `json:"requiredFlag,omitempty"`
	DisplayOnScreenFlag NullableBool `json:"displayOnScreenFlag,omitempty"`
	ReadOnlyFlag NullableBool `json:"readOnlyFlag,omitempty"`
	// Denotes that this custom field is included on a list view
	ListViewFlag NullableBool `json:"listViewFlag,omitempty"`
	// Only available with Button Field Type. Required when entryTypeIdentifier is button
	ButtonUrl *string `json:"buttonUrl,omitempty"`
	Options []UserDefinedFieldOption `json:"options,omitempty"`
	// List of business unit ids using custom field
	BusinessUnitIds []int32 `json:"businessUnitIds,omitempty"`
	// List of locations ids using custom field
	LocationIds []int32 `json:"locationIds,omitempty"`
	// Date in UTC the custom field was created
	DateCreated *string `json:"dateCreated,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewUserDefinedFieldInfo instantiates a new UserDefinedFieldInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefinedFieldInfo() *UserDefinedFieldInfo {
	this := UserDefinedFieldInfo{}
	return &this
}

// NewUserDefinedFieldInfoWithDefaults instantiates a new UserDefinedFieldInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefinedFieldInfoWithDefaults() *UserDefinedFieldInfo {
	this := UserDefinedFieldInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserDefinedFieldInfo) SetId(v int32) {
	o.Id = &v
}

// GetPodId returns the PodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetPodId() int32 {
	if o == nil || IsNil(o.PodId.Get()) {
		var ret int32
		return ret
	}
	return *o.PodId.Get()
}

// GetPodIdOk returns a tuple with the PodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetPodIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PodId.Get(), o.PodId.IsSet()
}

// HasPodId returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasPodId() bool {
	if o != nil && o.PodId.IsSet() {
		return true
	}

	return false
}

// SetPodId gets a reference to the given NullableInt32 and assigns it to the PodId field.
func (o *UserDefinedFieldInfo) SetPodId(v int32) {
	o.PodId.Set(&v)
}
// SetPodIdNil sets the value for PodId to be an explicit nil
func (o *UserDefinedFieldInfo) SetPodIdNil() {
	o.PodId.Set(nil)
}

// UnsetPodId ensures that no value is present for PodId, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetPodId() {
	o.PodId.Unset()
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *UserDefinedFieldInfo) SetCaption(v string) {
	o.Caption = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetSequenceNumber() int32 {
	if o == nil || IsNil(o.SequenceNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.SequenceNumber.Get()
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetSequenceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SequenceNumber.Get(), o.SequenceNumber.IsSet()
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasSequenceNumber() bool {
	if o != nil && o.SequenceNumber.IsSet() {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given NullableInt32 and assigns it to the SequenceNumber field.
func (o *UserDefinedFieldInfo) SetSequenceNumber(v int32) {
	o.SequenceNumber.Set(&v)
}
// SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil
func (o *UserDefinedFieldInfo) SetSequenceNumberNil() {
	o.SequenceNumber.Set(nil)
}

// UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetSequenceNumber() {
	o.SequenceNumber.Unset()
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetHelpText() string {
	if o == nil || IsNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetHelpTextOk() (*string, bool) {
	if o == nil || IsNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasHelpText() bool {
	if o != nil && !IsNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *UserDefinedFieldInfo) SetHelpText(v string) {
	o.HelpText = &v
}

// GetFieldTypeIdentifier returns the FieldTypeIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetFieldTypeIdentifier() string {
	if o == nil || IsNil(o.FieldTypeIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.FieldTypeIdentifier.Get()
}

// GetFieldTypeIdentifierOk returns a tuple with the FieldTypeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetFieldTypeIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldTypeIdentifier.Get(), o.FieldTypeIdentifier.IsSet()
}

// HasFieldTypeIdentifier returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasFieldTypeIdentifier() bool {
	if o != nil && o.FieldTypeIdentifier.IsSet() {
		return true
	}

	return false
}

// SetFieldTypeIdentifier gets a reference to the given NullableString and assigns it to the FieldTypeIdentifier field.
func (o *UserDefinedFieldInfo) SetFieldTypeIdentifier(v string) {
	o.FieldTypeIdentifier.Set(&v)
}
// SetFieldTypeIdentifierNil sets the value for FieldTypeIdentifier to be an explicit nil
func (o *UserDefinedFieldInfo) SetFieldTypeIdentifierNil() {
	o.FieldTypeIdentifier.Set(nil)
}

// UnsetFieldTypeIdentifier ensures that no value is present for FieldTypeIdentifier, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetFieldTypeIdentifier() {
	o.FieldTypeIdentifier.Unset()
}

// GetNumberDecimals returns the NumberDecimals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetNumberDecimals() int32 {
	if o == nil || IsNil(o.NumberDecimals.Get()) {
		var ret int32
		return ret
	}
	return *o.NumberDecimals.Get()
}

// GetNumberDecimalsOk returns a tuple with the NumberDecimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetNumberDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberDecimals.Get(), o.NumberDecimals.IsSet()
}

// HasNumberDecimals returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasNumberDecimals() bool {
	if o != nil && o.NumberDecimals.IsSet() {
		return true
	}

	return false
}

// SetNumberDecimals gets a reference to the given NullableInt32 and assigns it to the NumberDecimals field.
func (o *UserDefinedFieldInfo) SetNumberDecimals(v int32) {
	o.NumberDecimals.Set(&v)
}
// SetNumberDecimalsNil sets the value for NumberDecimals to be an explicit nil
func (o *UserDefinedFieldInfo) SetNumberDecimalsNil() {
	o.NumberDecimals.Set(nil)
}

// UnsetNumberDecimals ensures that no value is present for NumberDecimals, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetNumberDecimals() {
	o.NumberDecimals.Unset()
}

// GetEntryTypeIdentifier returns the EntryTypeIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetEntryTypeIdentifier() string {
	if o == nil || IsNil(o.EntryTypeIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.EntryTypeIdentifier.Get()
}

// GetEntryTypeIdentifierOk returns a tuple with the EntryTypeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetEntryTypeIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryTypeIdentifier.Get(), o.EntryTypeIdentifier.IsSet()
}

// HasEntryTypeIdentifier returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasEntryTypeIdentifier() bool {
	if o != nil && o.EntryTypeIdentifier.IsSet() {
		return true
	}

	return false
}

// SetEntryTypeIdentifier gets a reference to the given NullableString and assigns it to the EntryTypeIdentifier field.
func (o *UserDefinedFieldInfo) SetEntryTypeIdentifier(v string) {
	o.EntryTypeIdentifier.Set(&v)
}
// SetEntryTypeIdentifierNil sets the value for EntryTypeIdentifier to be an explicit nil
func (o *UserDefinedFieldInfo) SetEntryTypeIdentifierNil() {
	o.EntryTypeIdentifier.Set(nil)
}

// UnsetEntryTypeIdentifier ensures that no value is present for EntryTypeIdentifier, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetEntryTypeIdentifier() {
	o.EntryTypeIdentifier.Unset()
}

// GetRequiredFlag returns the RequiredFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetRequiredFlag() bool {
	if o == nil || IsNil(o.RequiredFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.RequiredFlag.Get()
}

// GetRequiredFlagOk returns a tuple with the RequiredFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetRequiredFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredFlag.Get(), o.RequiredFlag.IsSet()
}

// HasRequiredFlag returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasRequiredFlag() bool {
	if o != nil && o.RequiredFlag.IsSet() {
		return true
	}

	return false
}

// SetRequiredFlag gets a reference to the given NullableBool and assigns it to the RequiredFlag field.
func (o *UserDefinedFieldInfo) SetRequiredFlag(v bool) {
	o.RequiredFlag.Set(&v)
}
// SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil
func (o *UserDefinedFieldInfo) SetRequiredFlagNil() {
	o.RequiredFlag.Set(nil)
}

// UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetRequiredFlag() {
	o.RequiredFlag.Unset()
}

// GetDisplayOnScreenFlag returns the DisplayOnScreenFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetDisplayOnScreenFlag() bool {
	if o == nil || IsNil(o.DisplayOnScreenFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DisplayOnScreenFlag.Get()
}

// GetDisplayOnScreenFlagOk returns a tuple with the DisplayOnScreenFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetDisplayOnScreenFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayOnScreenFlag.Get(), o.DisplayOnScreenFlag.IsSet()
}

// HasDisplayOnScreenFlag returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasDisplayOnScreenFlag() bool {
	if o != nil && o.DisplayOnScreenFlag.IsSet() {
		return true
	}

	return false
}

// SetDisplayOnScreenFlag gets a reference to the given NullableBool and assigns it to the DisplayOnScreenFlag field.
func (o *UserDefinedFieldInfo) SetDisplayOnScreenFlag(v bool) {
	o.DisplayOnScreenFlag.Set(&v)
}
// SetDisplayOnScreenFlagNil sets the value for DisplayOnScreenFlag to be an explicit nil
func (o *UserDefinedFieldInfo) SetDisplayOnScreenFlagNil() {
	o.DisplayOnScreenFlag.Set(nil)
}

// UnsetDisplayOnScreenFlag ensures that no value is present for DisplayOnScreenFlag, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetDisplayOnScreenFlag() {
	o.DisplayOnScreenFlag.Unset()
}

// GetReadOnlyFlag returns the ReadOnlyFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetReadOnlyFlag() bool {
	if o == nil || IsNil(o.ReadOnlyFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ReadOnlyFlag.Get()
}

// GetReadOnlyFlagOk returns a tuple with the ReadOnlyFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetReadOnlyFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadOnlyFlag.Get(), o.ReadOnlyFlag.IsSet()
}

// HasReadOnlyFlag returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasReadOnlyFlag() bool {
	if o != nil && o.ReadOnlyFlag.IsSet() {
		return true
	}

	return false
}

// SetReadOnlyFlag gets a reference to the given NullableBool and assigns it to the ReadOnlyFlag field.
func (o *UserDefinedFieldInfo) SetReadOnlyFlag(v bool) {
	o.ReadOnlyFlag.Set(&v)
}
// SetReadOnlyFlagNil sets the value for ReadOnlyFlag to be an explicit nil
func (o *UserDefinedFieldInfo) SetReadOnlyFlagNil() {
	o.ReadOnlyFlag.Set(nil)
}

// UnsetReadOnlyFlag ensures that no value is present for ReadOnlyFlag, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetReadOnlyFlag() {
	o.ReadOnlyFlag.Unset()
}

// GetListViewFlag returns the ListViewFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDefinedFieldInfo) GetListViewFlag() bool {
	if o == nil || IsNil(o.ListViewFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ListViewFlag.Get()
}

// GetListViewFlagOk returns a tuple with the ListViewFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDefinedFieldInfo) GetListViewFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ListViewFlag.Get(), o.ListViewFlag.IsSet()
}

// HasListViewFlag returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasListViewFlag() bool {
	if o != nil && o.ListViewFlag.IsSet() {
		return true
	}

	return false
}

// SetListViewFlag gets a reference to the given NullableBool and assigns it to the ListViewFlag field.
func (o *UserDefinedFieldInfo) SetListViewFlag(v bool) {
	o.ListViewFlag.Set(&v)
}
// SetListViewFlagNil sets the value for ListViewFlag to be an explicit nil
func (o *UserDefinedFieldInfo) SetListViewFlagNil() {
	o.ListViewFlag.Set(nil)
}

// UnsetListViewFlag ensures that no value is present for ListViewFlag, not even an explicit nil
func (o *UserDefinedFieldInfo) UnsetListViewFlag() {
	o.ListViewFlag.Unset()
}

// GetButtonUrl returns the ButtonUrl field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetButtonUrl() string {
	if o == nil || IsNil(o.ButtonUrl) {
		var ret string
		return ret
	}
	return *o.ButtonUrl
}

// GetButtonUrlOk returns a tuple with the ButtonUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetButtonUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ButtonUrl) {
		return nil, false
	}
	return o.ButtonUrl, true
}

// HasButtonUrl returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasButtonUrl() bool {
	if o != nil && !IsNil(o.ButtonUrl) {
		return true
	}

	return false
}

// SetButtonUrl gets a reference to the given string and assigns it to the ButtonUrl field.
func (o *UserDefinedFieldInfo) SetButtonUrl(v string) {
	o.ButtonUrl = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetOptions() []UserDefinedFieldOption {
	if o == nil || IsNil(o.Options) {
		var ret []UserDefinedFieldOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetOptionsOk() ([]UserDefinedFieldOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []UserDefinedFieldOption and assigns it to the Options field.
func (o *UserDefinedFieldInfo) SetOptions(v []UserDefinedFieldOption) {
	o.Options = v
}

// GetBusinessUnitIds returns the BusinessUnitIds field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetBusinessUnitIds() []int32 {
	if o == nil || IsNil(o.BusinessUnitIds) {
		var ret []int32
		return ret
	}
	return o.BusinessUnitIds
}

// GetBusinessUnitIdsOk returns a tuple with the BusinessUnitIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetBusinessUnitIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.BusinessUnitIds) {
		return nil, false
	}
	return o.BusinessUnitIds, true
}

// HasBusinessUnitIds returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasBusinessUnitIds() bool {
	if o != nil && !IsNil(o.BusinessUnitIds) {
		return true
	}

	return false
}

// SetBusinessUnitIds gets a reference to the given []int32 and assigns it to the BusinessUnitIds field.
func (o *UserDefinedFieldInfo) SetBusinessUnitIds(v []int32) {
	o.BusinessUnitIds = v
}

// GetLocationIds returns the LocationIds field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetLocationIds() []int32 {
	if o == nil || IsNil(o.LocationIds) {
		var ret []int32
		return ret
	}
	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetLocationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.LocationIds) {
		return nil, false
	}
	return o.LocationIds, true
}

// HasLocationIds returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasLocationIds() bool {
	if o != nil && !IsNil(o.LocationIds) {
		return true
	}

	return false
}

// SetLocationIds gets a reference to the given []int32 and assigns it to the LocationIds field.
func (o *UserDefinedFieldInfo) SetLocationIds(v []int32) {
	o.LocationIds = v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetDateCreated() string {
	if o == nil || IsNil(o.DateCreated) {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetDateCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *UserDefinedFieldInfo) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *UserDefinedFieldInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefinedFieldInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *UserDefinedFieldInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *UserDefinedFieldInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o UserDefinedFieldInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDefinedFieldInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.PodId.IsSet() {
		toSerialize["podId"] = o.PodId.Get()
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if o.SequenceNumber.IsSet() {
		toSerialize["sequenceNumber"] = o.SequenceNumber.Get()
	}
	if !IsNil(o.HelpText) {
		toSerialize["helpText"] = o.HelpText
	}
	if o.FieldTypeIdentifier.IsSet() {
		toSerialize["fieldTypeIdentifier"] = o.FieldTypeIdentifier.Get()
	}
	if o.NumberDecimals.IsSet() {
		toSerialize["numberDecimals"] = o.NumberDecimals.Get()
	}
	if o.EntryTypeIdentifier.IsSet() {
		toSerialize["entryTypeIdentifier"] = o.EntryTypeIdentifier.Get()
	}
	if o.RequiredFlag.IsSet() {
		toSerialize["requiredFlag"] = o.RequiredFlag.Get()
	}
	if o.DisplayOnScreenFlag.IsSet() {
		toSerialize["displayOnScreenFlag"] = o.DisplayOnScreenFlag.Get()
	}
	if o.ReadOnlyFlag.IsSet() {
		toSerialize["readOnlyFlag"] = o.ReadOnlyFlag.Get()
	}
	if o.ListViewFlag.IsSet() {
		toSerialize["listViewFlag"] = o.ListViewFlag.Get()
	}
	if !IsNil(o.ButtonUrl) {
		toSerialize["buttonUrl"] = o.ButtonUrl
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.BusinessUnitIds) {
		toSerialize["businessUnitIds"] = o.BusinessUnitIds
	}
	if !IsNil(o.LocationIds) {
		toSerialize["locationIds"] = o.LocationIds
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableUserDefinedFieldInfo struct {
	value *UserDefinedFieldInfo
	isSet bool
}

func (v NullableUserDefinedFieldInfo) Get() *UserDefinedFieldInfo {
	return v.value
}

func (v *NullableUserDefinedFieldInfo) Set(val *UserDefinedFieldInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefinedFieldInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefinedFieldInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefinedFieldInfo(val *UserDefinedFieldInfo) *NullableUserDefinedFieldInfo {
	return &NullableUserDefinedFieldInfo{value: val, isSet: true}
}

func (v NullableUserDefinedFieldInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefinedFieldInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


