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

// checks if the DocumentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentInfo{}

// DocumentInfo struct for DocumentInfo
type DocumentInfo struct {
	Id *int32 `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	ServerFileName *string `json:"serverFileName,omitempty"`
	Owner *string `json:"owner,omitempty"`
	LinkFlag NullableBool `json:"linkFlag,omitempty"`
	ImageFlag NullableBool `json:"imageFlag,omitempty"`
	PublicFlag NullableBool `json:"publicFlag,omitempty"`
	HtmlTemplateFlag NullableBool `json:"htmlTemplateFlag,omitempty"`
	ReadOnlyFlag NullableBool `json:"readOnlyFlag,omitempty"`
	Size NullableInt32 `json:"size,omitempty"`
	UrlFlag NullableBool `json:"urlFlag,omitempty"`
	CreatedOnDate *string `json:"createdOnDate,omitempty"`
	DocumentType *DocumentTypeReference `json:"documentType,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewDocumentInfo instantiates a new DocumentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentInfo() *DocumentInfo {
	this := DocumentInfo{}
	return &this
}

// NewDocumentInfoWithDefaults instantiates a new DocumentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentInfoWithDefaults() *DocumentInfo {
	this := DocumentInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentInfo) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DocumentInfo) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DocumentInfo) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DocumentInfo) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DocumentInfo) SetTitle(v string) {
	o.Title = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DocumentInfo) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DocumentInfo) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DocumentInfo) SetFileName(v string) {
	o.FileName = &v
}

// GetServerFileName returns the ServerFileName field value if set, zero value otherwise.
func (o *DocumentInfo) GetServerFileName() string {
	if o == nil || IsNil(o.ServerFileName) {
		var ret string
		return ret
	}
	return *o.ServerFileName
}

// GetServerFileNameOk returns a tuple with the ServerFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetServerFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFileName) {
		return nil, false
	}
	return o.ServerFileName, true
}

// HasServerFileName returns a boolean if a field has been set.
func (o *DocumentInfo) HasServerFileName() bool {
	if o != nil && !IsNil(o.ServerFileName) {
		return true
	}

	return false
}

// SetServerFileName gets a reference to the given string and assigns it to the ServerFileName field.
func (o *DocumentInfo) SetServerFileName(v string) {
	o.ServerFileName = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DocumentInfo) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DocumentInfo) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *DocumentInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetLinkFlag returns the LinkFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetLinkFlag() bool {
	if o == nil || IsNil(o.LinkFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.LinkFlag.Get()
}

// GetLinkFlagOk returns a tuple with the LinkFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetLinkFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkFlag.Get(), o.LinkFlag.IsSet()
}

// HasLinkFlag returns a boolean if a field has been set.
func (o *DocumentInfo) HasLinkFlag() bool {
	if o != nil && o.LinkFlag.IsSet() {
		return true
	}

	return false
}

// SetLinkFlag gets a reference to the given NullableBool and assigns it to the LinkFlag field.
func (o *DocumentInfo) SetLinkFlag(v bool) {
	o.LinkFlag.Set(&v)
}
// SetLinkFlagNil sets the value for LinkFlag to be an explicit nil
func (o *DocumentInfo) SetLinkFlagNil() {
	o.LinkFlag.Set(nil)
}

// UnsetLinkFlag ensures that no value is present for LinkFlag, not even an explicit nil
func (o *DocumentInfo) UnsetLinkFlag() {
	o.LinkFlag.Unset()
}

// GetImageFlag returns the ImageFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetImageFlag() bool {
	if o == nil || IsNil(o.ImageFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ImageFlag.Get()
}

// GetImageFlagOk returns a tuple with the ImageFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetImageFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageFlag.Get(), o.ImageFlag.IsSet()
}

// HasImageFlag returns a boolean if a field has been set.
func (o *DocumentInfo) HasImageFlag() bool {
	if o != nil && o.ImageFlag.IsSet() {
		return true
	}

	return false
}

// SetImageFlag gets a reference to the given NullableBool and assigns it to the ImageFlag field.
func (o *DocumentInfo) SetImageFlag(v bool) {
	o.ImageFlag.Set(&v)
}
// SetImageFlagNil sets the value for ImageFlag to be an explicit nil
func (o *DocumentInfo) SetImageFlagNil() {
	o.ImageFlag.Set(nil)
}

// UnsetImageFlag ensures that no value is present for ImageFlag, not even an explicit nil
func (o *DocumentInfo) UnsetImageFlag() {
	o.ImageFlag.Unset()
}

// GetPublicFlag returns the PublicFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetPublicFlag() bool {
	if o == nil || IsNil(o.PublicFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.PublicFlag.Get()
}

// GetPublicFlagOk returns a tuple with the PublicFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetPublicFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicFlag.Get(), o.PublicFlag.IsSet()
}

// HasPublicFlag returns a boolean if a field has been set.
func (o *DocumentInfo) HasPublicFlag() bool {
	if o != nil && o.PublicFlag.IsSet() {
		return true
	}

	return false
}

// SetPublicFlag gets a reference to the given NullableBool and assigns it to the PublicFlag field.
func (o *DocumentInfo) SetPublicFlag(v bool) {
	o.PublicFlag.Set(&v)
}
// SetPublicFlagNil sets the value for PublicFlag to be an explicit nil
func (o *DocumentInfo) SetPublicFlagNil() {
	o.PublicFlag.Set(nil)
}

// UnsetPublicFlag ensures that no value is present for PublicFlag, not even an explicit nil
func (o *DocumentInfo) UnsetPublicFlag() {
	o.PublicFlag.Unset()
}

// GetHtmlTemplateFlag returns the HtmlTemplateFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetHtmlTemplateFlag() bool {
	if o == nil || IsNil(o.HtmlTemplateFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.HtmlTemplateFlag.Get()
}

// GetHtmlTemplateFlagOk returns a tuple with the HtmlTemplateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetHtmlTemplateFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HtmlTemplateFlag.Get(), o.HtmlTemplateFlag.IsSet()
}

// HasHtmlTemplateFlag returns a boolean if a field has been set.
func (o *DocumentInfo) HasHtmlTemplateFlag() bool {
	if o != nil && o.HtmlTemplateFlag.IsSet() {
		return true
	}

	return false
}

// SetHtmlTemplateFlag gets a reference to the given NullableBool and assigns it to the HtmlTemplateFlag field.
func (o *DocumentInfo) SetHtmlTemplateFlag(v bool) {
	o.HtmlTemplateFlag.Set(&v)
}
// SetHtmlTemplateFlagNil sets the value for HtmlTemplateFlag to be an explicit nil
func (o *DocumentInfo) SetHtmlTemplateFlagNil() {
	o.HtmlTemplateFlag.Set(nil)
}

// UnsetHtmlTemplateFlag ensures that no value is present for HtmlTemplateFlag, not even an explicit nil
func (o *DocumentInfo) UnsetHtmlTemplateFlag() {
	o.HtmlTemplateFlag.Unset()
}

// GetReadOnlyFlag returns the ReadOnlyFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetReadOnlyFlag() bool {
	if o == nil || IsNil(o.ReadOnlyFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.ReadOnlyFlag.Get()
}

// GetReadOnlyFlagOk returns a tuple with the ReadOnlyFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetReadOnlyFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadOnlyFlag.Get(), o.ReadOnlyFlag.IsSet()
}

// HasReadOnlyFlag returns a boolean if a field has been set.
func (o *DocumentInfo) HasReadOnlyFlag() bool {
	if o != nil && o.ReadOnlyFlag.IsSet() {
		return true
	}

	return false
}

// SetReadOnlyFlag gets a reference to the given NullableBool and assigns it to the ReadOnlyFlag field.
func (o *DocumentInfo) SetReadOnlyFlag(v bool) {
	o.ReadOnlyFlag.Set(&v)
}
// SetReadOnlyFlagNil sets the value for ReadOnlyFlag to be an explicit nil
func (o *DocumentInfo) SetReadOnlyFlagNil() {
	o.ReadOnlyFlag.Set(nil)
}

// UnsetReadOnlyFlag ensures that no value is present for ReadOnlyFlag, not even an explicit nil
func (o *DocumentInfo) UnsetReadOnlyFlag() {
	o.ReadOnlyFlag.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetSize() int32 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int32
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *DocumentInfo) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt32 and assigns it to the Size field.
func (o *DocumentInfo) SetSize(v int32) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *DocumentInfo) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *DocumentInfo) UnsetSize() {
	o.Size.Unset()
}

// GetUrlFlag returns the UrlFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentInfo) GetUrlFlag() bool {
	if o == nil || IsNil(o.UrlFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.UrlFlag.Get()
}

// GetUrlFlagOk returns a tuple with the UrlFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentInfo) GetUrlFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlFlag.Get(), o.UrlFlag.IsSet()
}

// HasUrlFlag returns a boolean if a field has been set.
func (o *DocumentInfo) HasUrlFlag() bool {
	if o != nil && o.UrlFlag.IsSet() {
		return true
	}

	return false
}

// SetUrlFlag gets a reference to the given NullableBool and assigns it to the UrlFlag field.
func (o *DocumentInfo) SetUrlFlag(v bool) {
	o.UrlFlag.Set(&v)
}
// SetUrlFlagNil sets the value for UrlFlag to be an explicit nil
func (o *DocumentInfo) SetUrlFlagNil() {
	o.UrlFlag.Set(nil)
}

// UnsetUrlFlag ensures that no value is present for UrlFlag, not even an explicit nil
func (o *DocumentInfo) UnsetUrlFlag() {
	o.UrlFlag.Unset()
}

// GetCreatedOnDate returns the CreatedOnDate field value if set, zero value otherwise.
func (o *DocumentInfo) GetCreatedOnDate() string {
	if o == nil || IsNil(o.CreatedOnDate) {
		var ret string
		return ret
	}
	return *o.CreatedOnDate
}

// GetCreatedOnDateOk returns a tuple with the CreatedOnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetCreatedOnDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedOnDate) {
		return nil, false
	}
	return o.CreatedOnDate, true
}

// HasCreatedOnDate returns a boolean if a field has been set.
func (o *DocumentInfo) HasCreatedOnDate() bool {
	if o != nil && !IsNil(o.CreatedOnDate) {
		return true
	}

	return false
}

// SetCreatedOnDate gets a reference to the given string and assigns it to the CreatedOnDate field.
func (o *DocumentInfo) SetCreatedOnDate(v string) {
	o.CreatedOnDate = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *DocumentInfo) GetDocumentType() DocumentTypeReference {
	if o == nil || IsNil(o.DocumentType) {
		var ret DocumentTypeReference
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetDocumentTypeOk() (*DocumentTypeReference, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *DocumentInfo) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given DocumentTypeReference and assigns it to the DocumentType field.
func (o *DocumentInfo) SetDocumentType(v DocumentTypeReference) {
	o.DocumentType = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *DocumentInfo) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *DocumentInfo) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *DocumentInfo) SetGuid(v string) {
	o.Guid = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *DocumentInfo) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *DocumentInfo) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *DocumentInfo) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o DocumentInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.ServerFileName) {
		toSerialize["serverFileName"] = o.ServerFileName
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if o.LinkFlag.IsSet() {
		toSerialize["linkFlag"] = o.LinkFlag.Get()
	}
	if o.ImageFlag.IsSet() {
		toSerialize["imageFlag"] = o.ImageFlag.Get()
	}
	if o.PublicFlag.IsSet() {
		toSerialize["publicFlag"] = o.PublicFlag.Get()
	}
	if o.HtmlTemplateFlag.IsSet() {
		toSerialize["htmlTemplateFlag"] = o.HtmlTemplateFlag.Get()
	}
	if o.ReadOnlyFlag.IsSet() {
		toSerialize["readOnlyFlag"] = o.ReadOnlyFlag.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.UrlFlag.IsSet() {
		toSerialize["urlFlag"] = o.UrlFlag.Get()
	}
	if !IsNil(o.CreatedOnDate) {
		toSerialize["createdOnDate"] = o.CreatedOnDate
	}
	if !IsNil(o.DocumentType) {
		toSerialize["documentType"] = o.DocumentType
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableDocumentInfo struct {
	value *DocumentInfo
	isSet bool
}

func (v NullableDocumentInfo) Get() *DocumentInfo {
	return v.value
}

func (v *NullableDocumentInfo) Set(val *DocumentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentInfo(val *DocumentInfo) *NullableDocumentInfo {
	return &NullableDocumentInfo{value: val, isSet: true}
}

func (v NullableDocumentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


