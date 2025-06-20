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

// checks if the GLCaption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GLCaption{}

// GLCaption struct for GLCaption
type GLCaption struct {
	Id *int32 `json:"id,omitempty"`
	//  Max length: 255;
	Segment1 *string `json:"segment1,omitempty"`
	//  Max length: 255;
	Segment2 *string `json:"segment2,omitempty"`
	//  Max length: 255;
	Segment3 *string `json:"segment3,omitempty"`
	//  Max length: 255;
	Segment4 *string `json:"segment4,omitempty"`
	//  Max length: 255;
	Segment5 *string `json:"segment5,omitempty"`
	//  Max length: 255;
	Segment6 *string `json:"segment6,omitempty"`
	//  Max length: 255;
	Segment7 *string `json:"segment7,omitempty"`
	//  Max length: 255;
	Segment8 *string `json:"segment8,omitempty"`
	//  Max length: 255;
	Segment9 *string `json:"segment9,omitempty"`
	//  Max length: 255;
	Segment10 *string `json:"segment10,omitempty"`
	Segment1type NullableString `json:"segment1type,omitempty"`
	Segment2type NullableString `json:"segment2type,omitempty"`
	Segment3type NullableString `json:"segment3type,omitempty"`
	Segment4type NullableString `json:"segment4type,omitempty"`
	Segment5type NullableString `json:"segment5type,omitempty"`
	Segment6type NullableString `json:"segment6type,omitempty"`
	Segment7type NullableString `json:"segment7type,omitempty"`
	Segment8type NullableString `json:"segment8type,omitempty"`
	Segment9type NullableString `json:"segment9type,omitempty"`
	Segment10type NullableString `json:"segment10type,omitempty"`
	//  Max length: 255;
	Cogs1 *string `json:"cogs1,omitempty"`
	//  Max length: 255;
	Cogs2 *string `json:"cogs2,omitempty"`
	//  Max length: 255;
	Cogs3 *string `json:"cogs3,omitempty"`
	//  Max length: 255;
	Cogs4 *string `json:"cogs4,omitempty"`
	//  Max length: 255;
	Cogs5 *string `json:"cogs5,omitempty"`
	//  Max length: 255;
	Cogs6 *string `json:"cogs6,omitempty"`
	//  Max length: 255;
	Cogs7 *string `json:"cogs7,omitempty"`
	//  Max length: 255;
	Cogs8 *string `json:"cogs8,omitempty"`
	//  Max length: 255;
	Cogs9 *string `json:"cogs9,omitempty"`
	//  Max length: 255;
	Cogs10 *string `json:"cogs10,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

// NewGLCaption instantiates a new GLCaption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGLCaption() *GLCaption {
	this := GLCaption{}
	return &this
}

// NewGLCaptionWithDefaults instantiates a new GLCaption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGLCaptionWithDefaults() *GLCaption {
	this := GLCaption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GLCaption) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GLCaption) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GLCaption) SetId(v int32) {
	o.Id = &v
}

// GetSegment1 returns the Segment1 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment1() string {
	if o == nil || IsNil(o.Segment1) {
		var ret string
		return ret
	}
	return *o.Segment1
}

// GetSegment1Ok returns a tuple with the Segment1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment1Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment1) {
		return nil, false
	}
	return o.Segment1, true
}

// HasSegment1 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment1() bool {
	if o != nil && !IsNil(o.Segment1) {
		return true
	}

	return false
}

// SetSegment1 gets a reference to the given string and assigns it to the Segment1 field.
func (o *GLCaption) SetSegment1(v string) {
	o.Segment1 = &v
}

// GetSegment2 returns the Segment2 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment2() string {
	if o == nil || IsNil(o.Segment2) {
		var ret string
		return ret
	}
	return *o.Segment2
}

// GetSegment2Ok returns a tuple with the Segment2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment2Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment2) {
		return nil, false
	}
	return o.Segment2, true
}

// HasSegment2 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment2() bool {
	if o != nil && !IsNil(o.Segment2) {
		return true
	}

	return false
}

// SetSegment2 gets a reference to the given string and assigns it to the Segment2 field.
func (o *GLCaption) SetSegment2(v string) {
	o.Segment2 = &v
}

// GetSegment3 returns the Segment3 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment3() string {
	if o == nil || IsNil(o.Segment3) {
		var ret string
		return ret
	}
	return *o.Segment3
}

// GetSegment3Ok returns a tuple with the Segment3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment3Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment3) {
		return nil, false
	}
	return o.Segment3, true
}

// HasSegment3 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment3() bool {
	if o != nil && !IsNil(o.Segment3) {
		return true
	}

	return false
}

// SetSegment3 gets a reference to the given string and assigns it to the Segment3 field.
func (o *GLCaption) SetSegment3(v string) {
	o.Segment3 = &v
}

// GetSegment4 returns the Segment4 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment4() string {
	if o == nil || IsNil(o.Segment4) {
		var ret string
		return ret
	}
	return *o.Segment4
}

// GetSegment4Ok returns a tuple with the Segment4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment4Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment4) {
		return nil, false
	}
	return o.Segment4, true
}

// HasSegment4 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment4() bool {
	if o != nil && !IsNil(o.Segment4) {
		return true
	}

	return false
}

// SetSegment4 gets a reference to the given string and assigns it to the Segment4 field.
func (o *GLCaption) SetSegment4(v string) {
	o.Segment4 = &v
}

// GetSegment5 returns the Segment5 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment5() string {
	if o == nil || IsNil(o.Segment5) {
		var ret string
		return ret
	}
	return *o.Segment5
}

// GetSegment5Ok returns a tuple with the Segment5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment5Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment5) {
		return nil, false
	}
	return o.Segment5, true
}

// HasSegment5 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment5() bool {
	if o != nil && !IsNil(o.Segment5) {
		return true
	}

	return false
}

// SetSegment5 gets a reference to the given string and assigns it to the Segment5 field.
func (o *GLCaption) SetSegment5(v string) {
	o.Segment5 = &v
}

// GetSegment6 returns the Segment6 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment6() string {
	if o == nil || IsNil(o.Segment6) {
		var ret string
		return ret
	}
	return *o.Segment6
}

// GetSegment6Ok returns a tuple with the Segment6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment6Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment6) {
		return nil, false
	}
	return o.Segment6, true
}

// HasSegment6 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment6() bool {
	if o != nil && !IsNil(o.Segment6) {
		return true
	}

	return false
}

// SetSegment6 gets a reference to the given string and assigns it to the Segment6 field.
func (o *GLCaption) SetSegment6(v string) {
	o.Segment6 = &v
}

// GetSegment7 returns the Segment7 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment7() string {
	if o == nil || IsNil(o.Segment7) {
		var ret string
		return ret
	}
	return *o.Segment7
}

// GetSegment7Ok returns a tuple with the Segment7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment7Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment7) {
		return nil, false
	}
	return o.Segment7, true
}

// HasSegment7 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment7() bool {
	if o != nil && !IsNil(o.Segment7) {
		return true
	}

	return false
}

// SetSegment7 gets a reference to the given string and assigns it to the Segment7 field.
func (o *GLCaption) SetSegment7(v string) {
	o.Segment7 = &v
}

// GetSegment8 returns the Segment8 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment8() string {
	if o == nil || IsNil(o.Segment8) {
		var ret string
		return ret
	}
	return *o.Segment8
}

// GetSegment8Ok returns a tuple with the Segment8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment8Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment8) {
		return nil, false
	}
	return o.Segment8, true
}

// HasSegment8 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment8() bool {
	if o != nil && !IsNil(o.Segment8) {
		return true
	}

	return false
}

// SetSegment8 gets a reference to the given string and assigns it to the Segment8 field.
func (o *GLCaption) SetSegment8(v string) {
	o.Segment8 = &v
}

// GetSegment9 returns the Segment9 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment9() string {
	if o == nil || IsNil(o.Segment9) {
		var ret string
		return ret
	}
	return *o.Segment9
}

// GetSegment9Ok returns a tuple with the Segment9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment9Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment9) {
		return nil, false
	}
	return o.Segment9, true
}

// HasSegment9 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment9() bool {
	if o != nil && !IsNil(o.Segment9) {
		return true
	}

	return false
}

// SetSegment9 gets a reference to the given string and assigns it to the Segment9 field.
func (o *GLCaption) SetSegment9(v string) {
	o.Segment9 = &v
}

// GetSegment10 returns the Segment10 field value if set, zero value otherwise.
func (o *GLCaption) GetSegment10() string {
	if o == nil || IsNil(o.Segment10) {
		var ret string
		return ret
	}
	return *o.Segment10
}

// GetSegment10Ok returns a tuple with the Segment10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetSegment10Ok() (*string, bool) {
	if o == nil || IsNil(o.Segment10) {
		return nil, false
	}
	return o.Segment10, true
}

// HasSegment10 returns a boolean if a field has been set.
func (o *GLCaption) HasSegment10() bool {
	if o != nil && !IsNil(o.Segment10) {
		return true
	}

	return false
}

// SetSegment10 gets a reference to the given string and assigns it to the Segment10 field.
func (o *GLCaption) SetSegment10(v string) {
	o.Segment10 = &v
}

// GetSegment1type returns the Segment1type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment1type() string {
	if o == nil || IsNil(o.Segment1type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment1type.Get()
}

// GetSegment1typeOk returns a tuple with the Segment1type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment1typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment1type.Get(), o.Segment1type.IsSet()
}

// HasSegment1type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment1type() bool {
	if o != nil && o.Segment1type.IsSet() {
		return true
	}

	return false
}

// SetSegment1type gets a reference to the given NullableString and assigns it to the Segment1type field.
func (o *GLCaption) SetSegment1type(v string) {
	o.Segment1type.Set(&v)
}
// SetSegment1typeNil sets the value for Segment1type to be an explicit nil
func (o *GLCaption) SetSegment1typeNil() {
	o.Segment1type.Set(nil)
}

// UnsetSegment1type ensures that no value is present for Segment1type, not even an explicit nil
func (o *GLCaption) UnsetSegment1type() {
	o.Segment1type.Unset()
}

// GetSegment2type returns the Segment2type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment2type() string {
	if o == nil || IsNil(o.Segment2type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment2type.Get()
}

// GetSegment2typeOk returns a tuple with the Segment2type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment2typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment2type.Get(), o.Segment2type.IsSet()
}

// HasSegment2type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment2type() bool {
	if o != nil && o.Segment2type.IsSet() {
		return true
	}

	return false
}

// SetSegment2type gets a reference to the given NullableString and assigns it to the Segment2type field.
func (o *GLCaption) SetSegment2type(v string) {
	o.Segment2type.Set(&v)
}
// SetSegment2typeNil sets the value for Segment2type to be an explicit nil
func (o *GLCaption) SetSegment2typeNil() {
	o.Segment2type.Set(nil)
}

// UnsetSegment2type ensures that no value is present for Segment2type, not even an explicit nil
func (o *GLCaption) UnsetSegment2type() {
	o.Segment2type.Unset()
}

// GetSegment3type returns the Segment3type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment3type() string {
	if o == nil || IsNil(o.Segment3type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment3type.Get()
}

// GetSegment3typeOk returns a tuple with the Segment3type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment3typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment3type.Get(), o.Segment3type.IsSet()
}

// HasSegment3type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment3type() bool {
	if o != nil && o.Segment3type.IsSet() {
		return true
	}

	return false
}

// SetSegment3type gets a reference to the given NullableString and assigns it to the Segment3type field.
func (o *GLCaption) SetSegment3type(v string) {
	o.Segment3type.Set(&v)
}
// SetSegment3typeNil sets the value for Segment3type to be an explicit nil
func (o *GLCaption) SetSegment3typeNil() {
	o.Segment3type.Set(nil)
}

// UnsetSegment3type ensures that no value is present for Segment3type, not even an explicit nil
func (o *GLCaption) UnsetSegment3type() {
	o.Segment3type.Unset()
}

// GetSegment4type returns the Segment4type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment4type() string {
	if o == nil || IsNil(o.Segment4type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment4type.Get()
}

// GetSegment4typeOk returns a tuple with the Segment4type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment4typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment4type.Get(), o.Segment4type.IsSet()
}

// HasSegment4type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment4type() bool {
	if o != nil && o.Segment4type.IsSet() {
		return true
	}

	return false
}

// SetSegment4type gets a reference to the given NullableString and assigns it to the Segment4type field.
func (o *GLCaption) SetSegment4type(v string) {
	o.Segment4type.Set(&v)
}
// SetSegment4typeNil sets the value for Segment4type to be an explicit nil
func (o *GLCaption) SetSegment4typeNil() {
	o.Segment4type.Set(nil)
}

// UnsetSegment4type ensures that no value is present for Segment4type, not even an explicit nil
func (o *GLCaption) UnsetSegment4type() {
	o.Segment4type.Unset()
}

// GetSegment5type returns the Segment5type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment5type() string {
	if o == nil || IsNil(o.Segment5type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment5type.Get()
}

// GetSegment5typeOk returns a tuple with the Segment5type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment5typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment5type.Get(), o.Segment5type.IsSet()
}

// HasSegment5type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment5type() bool {
	if o != nil && o.Segment5type.IsSet() {
		return true
	}

	return false
}

// SetSegment5type gets a reference to the given NullableString and assigns it to the Segment5type field.
func (o *GLCaption) SetSegment5type(v string) {
	o.Segment5type.Set(&v)
}
// SetSegment5typeNil sets the value for Segment5type to be an explicit nil
func (o *GLCaption) SetSegment5typeNil() {
	o.Segment5type.Set(nil)
}

// UnsetSegment5type ensures that no value is present for Segment5type, not even an explicit nil
func (o *GLCaption) UnsetSegment5type() {
	o.Segment5type.Unset()
}

// GetSegment6type returns the Segment6type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment6type() string {
	if o == nil || IsNil(o.Segment6type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment6type.Get()
}

// GetSegment6typeOk returns a tuple with the Segment6type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment6typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment6type.Get(), o.Segment6type.IsSet()
}

// HasSegment6type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment6type() bool {
	if o != nil && o.Segment6type.IsSet() {
		return true
	}

	return false
}

// SetSegment6type gets a reference to the given NullableString and assigns it to the Segment6type field.
func (o *GLCaption) SetSegment6type(v string) {
	o.Segment6type.Set(&v)
}
// SetSegment6typeNil sets the value for Segment6type to be an explicit nil
func (o *GLCaption) SetSegment6typeNil() {
	o.Segment6type.Set(nil)
}

// UnsetSegment6type ensures that no value is present for Segment6type, not even an explicit nil
func (o *GLCaption) UnsetSegment6type() {
	o.Segment6type.Unset()
}

// GetSegment7type returns the Segment7type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment7type() string {
	if o == nil || IsNil(o.Segment7type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment7type.Get()
}

// GetSegment7typeOk returns a tuple with the Segment7type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment7typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment7type.Get(), o.Segment7type.IsSet()
}

// HasSegment7type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment7type() bool {
	if o != nil && o.Segment7type.IsSet() {
		return true
	}

	return false
}

// SetSegment7type gets a reference to the given NullableString and assigns it to the Segment7type field.
func (o *GLCaption) SetSegment7type(v string) {
	o.Segment7type.Set(&v)
}
// SetSegment7typeNil sets the value for Segment7type to be an explicit nil
func (o *GLCaption) SetSegment7typeNil() {
	o.Segment7type.Set(nil)
}

// UnsetSegment7type ensures that no value is present for Segment7type, not even an explicit nil
func (o *GLCaption) UnsetSegment7type() {
	o.Segment7type.Unset()
}

// GetSegment8type returns the Segment8type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment8type() string {
	if o == nil || IsNil(o.Segment8type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment8type.Get()
}

// GetSegment8typeOk returns a tuple with the Segment8type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment8typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment8type.Get(), o.Segment8type.IsSet()
}

// HasSegment8type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment8type() bool {
	if o != nil && o.Segment8type.IsSet() {
		return true
	}

	return false
}

// SetSegment8type gets a reference to the given NullableString and assigns it to the Segment8type field.
func (o *GLCaption) SetSegment8type(v string) {
	o.Segment8type.Set(&v)
}
// SetSegment8typeNil sets the value for Segment8type to be an explicit nil
func (o *GLCaption) SetSegment8typeNil() {
	o.Segment8type.Set(nil)
}

// UnsetSegment8type ensures that no value is present for Segment8type, not even an explicit nil
func (o *GLCaption) UnsetSegment8type() {
	o.Segment8type.Unset()
}

// GetSegment9type returns the Segment9type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment9type() string {
	if o == nil || IsNil(o.Segment9type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment9type.Get()
}

// GetSegment9typeOk returns a tuple with the Segment9type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment9typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment9type.Get(), o.Segment9type.IsSet()
}

// HasSegment9type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment9type() bool {
	if o != nil && o.Segment9type.IsSet() {
		return true
	}

	return false
}

// SetSegment9type gets a reference to the given NullableString and assigns it to the Segment9type field.
func (o *GLCaption) SetSegment9type(v string) {
	o.Segment9type.Set(&v)
}
// SetSegment9typeNil sets the value for Segment9type to be an explicit nil
func (o *GLCaption) SetSegment9typeNil() {
	o.Segment9type.Set(nil)
}

// UnsetSegment9type ensures that no value is present for Segment9type, not even an explicit nil
func (o *GLCaption) UnsetSegment9type() {
	o.Segment9type.Unset()
}

// GetSegment10type returns the Segment10type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GLCaption) GetSegment10type() string {
	if o == nil || IsNil(o.Segment10type.Get()) {
		var ret string
		return ret
	}
	return *o.Segment10type.Get()
}

// GetSegment10typeOk returns a tuple with the Segment10type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GLCaption) GetSegment10typeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segment10type.Get(), o.Segment10type.IsSet()
}

// HasSegment10type returns a boolean if a field has been set.
func (o *GLCaption) HasSegment10type() bool {
	if o != nil && o.Segment10type.IsSet() {
		return true
	}

	return false
}

// SetSegment10type gets a reference to the given NullableString and assigns it to the Segment10type field.
func (o *GLCaption) SetSegment10type(v string) {
	o.Segment10type.Set(&v)
}
// SetSegment10typeNil sets the value for Segment10type to be an explicit nil
func (o *GLCaption) SetSegment10typeNil() {
	o.Segment10type.Set(nil)
}

// UnsetSegment10type ensures that no value is present for Segment10type, not even an explicit nil
func (o *GLCaption) UnsetSegment10type() {
	o.Segment10type.Unset()
}

// GetCogs1 returns the Cogs1 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs1() string {
	if o == nil || IsNil(o.Cogs1) {
		var ret string
		return ret
	}
	return *o.Cogs1
}

// GetCogs1Ok returns a tuple with the Cogs1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs1Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs1) {
		return nil, false
	}
	return o.Cogs1, true
}

// HasCogs1 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs1() bool {
	if o != nil && !IsNil(o.Cogs1) {
		return true
	}

	return false
}

// SetCogs1 gets a reference to the given string and assigns it to the Cogs1 field.
func (o *GLCaption) SetCogs1(v string) {
	o.Cogs1 = &v
}

// GetCogs2 returns the Cogs2 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs2() string {
	if o == nil || IsNil(o.Cogs2) {
		var ret string
		return ret
	}
	return *o.Cogs2
}

// GetCogs2Ok returns a tuple with the Cogs2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs2Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs2) {
		return nil, false
	}
	return o.Cogs2, true
}

// HasCogs2 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs2() bool {
	if o != nil && !IsNil(o.Cogs2) {
		return true
	}

	return false
}

// SetCogs2 gets a reference to the given string and assigns it to the Cogs2 field.
func (o *GLCaption) SetCogs2(v string) {
	o.Cogs2 = &v
}

// GetCogs3 returns the Cogs3 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs3() string {
	if o == nil || IsNil(o.Cogs3) {
		var ret string
		return ret
	}
	return *o.Cogs3
}

// GetCogs3Ok returns a tuple with the Cogs3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs3Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs3) {
		return nil, false
	}
	return o.Cogs3, true
}

// HasCogs3 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs3() bool {
	if o != nil && !IsNil(o.Cogs3) {
		return true
	}

	return false
}

// SetCogs3 gets a reference to the given string and assigns it to the Cogs3 field.
func (o *GLCaption) SetCogs3(v string) {
	o.Cogs3 = &v
}

// GetCogs4 returns the Cogs4 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs4() string {
	if o == nil || IsNil(o.Cogs4) {
		var ret string
		return ret
	}
	return *o.Cogs4
}

// GetCogs4Ok returns a tuple with the Cogs4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs4Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs4) {
		return nil, false
	}
	return o.Cogs4, true
}

// HasCogs4 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs4() bool {
	if o != nil && !IsNil(o.Cogs4) {
		return true
	}

	return false
}

// SetCogs4 gets a reference to the given string and assigns it to the Cogs4 field.
func (o *GLCaption) SetCogs4(v string) {
	o.Cogs4 = &v
}

// GetCogs5 returns the Cogs5 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs5() string {
	if o == nil || IsNil(o.Cogs5) {
		var ret string
		return ret
	}
	return *o.Cogs5
}

// GetCogs5Ok returns a tuple with the Cogs5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs5Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs5) {
		return nil, false
	}
	return o.Cogs5, true
}

// HasCogs5 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs5() bool {
	if o != nil && !IsNil(o.Cogs5) {
		return true
	}

	return false
}

// SetCogs5 gets a reference to the given string and assigns it to the Cogs5 field.
func (o *GLCaption) SetCogs5(v string) {
	o.Cogs5 = &v
}

// GetCogs6 returns the Cogs6 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs6() string {
	if o == nil || IsNil(o.Cogs6) {
		var ret string
		return ret
	}
	return *o.Cogs6
}

// GetCogs6Ok returns a tuple with the Cogs6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs6Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs6) {
		return nil, false
	}
	return o.Cogs6, true
}

// HasCogs6 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs6() bool {
	if o != nil && !IsNil(o.Cogs6) {
		return true
	}

	return false
}

// SetCogs6 gets a reference to the given string and assigns it to the Cogs6 field.
func (o *GLCaption) SetCogs6(v string) {
	o.Cogs6 = &v
}

// GetCogs7 returns the Cogs7 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs7() string {
	if o == nil || IsNil(o.Cogs7) {
		var ret string
		return ret
	}
	return *o.Cogs7
}

// GetCogs7Ok returns a tuple with the Cogs7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs7Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs7) {
		return nil, false
	}
	return o.Cogs7, true
}

// HasCogs7 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs7() bool {
	if o != nil && !IsNil(o.Cogs7) {
		return true
	}

	return false
}

// SetCogs7 gets a reference to the given string and assigns it to the Cogs7 field.
func (o *GLCaption) SetCogs7(v string) {
	o.Cogs7 = &v
}

// GetCogs8 returns the Cogs8 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs8() string {
	if o == nil || IsNil(o.Cogs8) {
		var ret string
		return ret
	}
	return *o.Cogs8
}

// GetCogs8Ok returns a tuple with the Cogs8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs8Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs8) {
		return nil, false
	}
	return o.Cogs8, true
}

// HasCogs8 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs8() bool {
	if o != nil && !IsNil(o.Cogs8) {
		return true
	}

	return false
}

// SetCogs8 gets a reference to the given string and assigns it to the Cogs8 field.
func (o *GLCaption) SetCogs8(v string) {
	o.Cogs8 = &v
}

// GetCogs9 returns the Cogs9 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs9() string {
	if o == nil || IsNil(o.Cogs9) {
		var ret string
		return ret
	}
	return *o.Cogs9
}

// GetCogs9Ok returns a tuple with the Cogs9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs9Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs9) {
		return nil, false
	}
	return o.Cogs9, true
}

// HasCogs9 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs9() bool {
	if o != nil && !IsNil(o.Cogs9) {
		return true
	}

	return false
}

// SetCogs9 gets a reference to the given string and assigns it to the Cogs9 field.
func (o *GLCaption) SetCogs9(v string) {
	o.Cogs9 = &v
}

// GetCogs10 returns the Cogs10 field value if set, zero value otherwise.
func (o *GLCaption) GetCogs10() string {
	if o == nil || IsNil(o.Cogs10) {
		var ret string
		return ret
	}
	return *o.Cogs10
}

// GetCogs10Ok returns a tuple with the Cogs10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetCogs10Ok() (*string, bool) {
	if o == nil || IsNil(o.Cogs10) {
		return nil, false
	}
	return o.Cogs10, true
}

// HasCogs10 returns a boolean if a field has been set.
func (o *GLCaption) HasCogs10() bool {
	if o != nil && !IsNil(o.Cogs10) {
		return true
	}

	return false
}

// SetCogs10 gets a reference to the given string and assigns it to the Cogs10 field.
func (o *GLCaption) SetCogs10(v string) {
	o.Cogs10 = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *GLCaption) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GLCaption) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *GLCaption) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *GLCaption) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o GLCaption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GLCaption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Segment1) {
		toSerialize["segment1"] = o.Segment1
	}
	if !IsNil(o.Segment2) {
		toSerialize["segment2"] = o.Segment2
	}
	if !IsNil(o.Segment3) {
		toSerialize["segment3"] = o.Segment3
	}
	if !IsNil(o.Segment4) {
		toSerialize["segment4"] = o.Segment4
	}
	if !IsNil(o.Segment5) {
		toSerialize["segment5"] = o.Segment5
	}
	if !IsNil(o.Segment6) {
		toSerialize["segment6"] = o.Segment6
	}
	if !IsNil(o.Segment7) {
		toSerialize["segment7"] = o.Segment7
	}
	if !IsNil(o.Segment8) {
		toSerialize["segment8"] = o.Segment8
	}
	if !IsNil(o.Segment9) {
		toSerialize["segment9"] = o.Segment9
	}
	if !IsNil(o.Segment10) {
		toSerialize["segment10"] = o.Segment10
	}
	if o.Segment1type.IsSet() {
		toSerialize["segment1type"] = o.Segment1type.Get()
	}
	if o.Segment2type.IsSet() {
		toSerialize["segment2type"] = o.Segment2type.Get()
	}
	if o.Segment3type.IsSet() {
		toSerialize["segment3type"] = o.Segment3type.Get()
	}
	if o.Segment4type.IsSet() {
		toSerialize["segment4type"] = o.Segment4type.Get()
	}
	if o.Segment5type.IsSet() {
		toSerialize["segment5type"] = o.Segment5type.Get()
	}
	if o.Segment6type.IsSet() {
		toSerialize["segment6type"] = o.Segment6type.Get()
	}
	if o.Segment7type.IsSet() {
		toSerialize["segment7type"] = o.Segment7type.Get()
	}
	if o.Segment8type.IsSet() {
		toSerialize["segment8type"] = o.Segment8type.Get()
	}
	if o.Segment9type.IsSet() {
		toSerialize["segment9type"] = o.Segment9type.Get()
	}
	if o.Segment10type.IsSet() {
		toSerialize["segment10type"] = o.Segment10type.Get()
	}
	if !IsNil(o.Cogs1) {
		toSerialize["cogs1"] = o.Cogs1
	}
	if !IsNil(o.Cogs2) {
		toSerialize["cogs2"] = o.Cogs2
	}
	if !IsNil(o.Cogs3) {
		toSerialize["cogs3"] = o.Cogs3
	}
	if !IsNil(o.Cogs4) {
		toSerialize["cogs4"] = o.Cogs4
	}
	if !IsNil(o.Cogs5) {
		toSerialize["cogs5"] = o.Cogs5
	}
	if !IsNil(o.Cogs6) {
		toSerialize["cogs6"] = o.Cogs6
	}
	if !IsNil(o.Cogs7) {
		toSerialize["cogs7"] = o.Cogs7
	}
	if !IsNil(o.Cogs8) {
		toSerialize["cogs8"] = o.Cogs8
	}
	if !IsNil(o.Cogs9) {
		toSerialize["cogs9"] = o.Cogs9
	}
	if !IsNil(o.Cogs10) {
		toSerialize["cogs10"] = o.Cogs10
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

type NullableGLCaption struct {
	value *GLCaption
	isSet bool
}

func (v NullableGLCaption) Get() *GLCaption {
	return v.value
}

func (v *NullableGLCaption) Set(val *GLCaption) {
	v.value = val
	v.isSet = true
}

func (v NullableGLCaption) IsSet() bool {
	return v.isSet
}

func (v *NullableGLCaption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGLCaption(val *GLCaption) *NullableGLCaption {
	return &NullableGLCaption{value: val, isSet: true}
}

func (v NullableGLCaption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGLCaption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


