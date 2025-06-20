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

// checks if the PurchaseOrderStatusEmailTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseOrderStatusEmailTemplate{}

// PurchaseOrderStatusEmailTemplate struct for PurchaseOrderStatusEmailTemplate
type PurchaseOrderStatusEmailTemplate struct {
	Id *int32 `json:"id,omitempty"`
	Status *PurchaseOrderStatusReference `json:"status,omitempty"`
	UseSenderFlag NullableBool `json:"useSenderFlag,omitempty"`
	//  Max length: 100;
	FirstName *string `json:"firstName,omitempty"`
	//  Max length: 100;
	LastName *string `json:"lastName,omitempty"`
	//  Max length: 100;
	EmailAddress *string `json:"emailAddress,omitempty"`
	//  Max length: 200;
	Subject string `json:"subject"`
	Body *string `json:"body,omitempty"`
	CopySenderFlag NullableBool `json:"copySenderFlag,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _PurchaseOrderStatusEmailTemplate PurchaseOrderStatusEmailTemplate

// NewPurchaseOrderStatusEmailTemplate instantiates a new PurchaseOrderStatusEmailTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrderStatusEmailTemplate(subject string) *PurchaseOrderStatusEmailTemplate {
	this := PurchaseOrderStatusEmailTemplate{}
	this.Subject = subject
	return &this
}

// NewPurchaseOrderStatusEmailTemplateWithDefaults instantiates a new PurchaseOrderStatusEmailTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderStatusEmailTemplateWithDefaults() *PurchaseOrderStatusEmailTemplate {
	this := PurchaseOrderStatusEmailTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PurchaseOrderStatusEmailTemplate) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetStatus() PurchaseOrderStatusReference {
	if o == nil || IsNil(o.Status) {
		var ret PurchaseOrderStatusReference
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetStatusOk() (*PurchaseOrderStatusReference, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PurchaseOrderStatusReference and assigns it to the Status field.
func (o *PurchaseOrderStatusEmailTemplate) SetStatus(v PurchaseOrderStatusReference) {
	o.Status = &v
}

// GetUseSenderFlag returns the UseSenderFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderStatusEmailTemplate) GetUseSenderFlag() bool {
	if o == nil || IsNil(o.UseSenderFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.UseSenderFlag.Get()
}

// GetUseSenderFlagOk returns a tuple with the UseSenderFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderStatusEmailTemplate) GetUseSenderFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseSenderFlag.Get(), o.UseSenderFlag.IsSet()
}

// HasUseSenderFlag returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasUseSenderFlag() bool {
	if o != nil && o.UseSenderFlag.IsSet() {
		return true
	}

	return false
}

// SetUseSenderFlag gets a reference to the given NullableBool and assigns it to the UseSenderFlag field.
func (o *PurchaseOrderStatusEmailTemplate) SetUseSenderFlag(v bool) {
	o.UseSenderFlag.Set(&v)
}
// SetUseSenderFlagNil sets the value for UseSenderFlag to be an explicit nil
func (o *PurchaseOrderStatusEmailTemplate) SetUseSenderFlagNil() {
	o.UseSenderFlag.Set(nil)
}

// UnsetUseSenderFlag ensures that no value is present for UseSenderFlag, not even an explicit nil
func (o *PurchaseOrderStatusEmailTemplate) UnsetUseSenderFlag() {
	o.UseSenderFlag.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PurchaseOrderStatusEmailTemplate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PurchaseOrderStatusEmailTemplate) SetLastName(v string) {
	o.LastName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *PurchaseOrderStatusEmailTemplate) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetSubject returns the Subject field value
func (o *PurchaseOrderStatusEmailTemplate) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *PurchaseOrderStatusEmailTemplate) SetSubject(v string) {
	o.Subject = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *PurchaseOrderStatusEmailTemplate) SetBody(v string) {
	o.Body = &v
}

// GetCopySenderFlag returns the CopySenderFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderStatusEmailTemplate) GetCopySenderFlag() bool {
	if o == nil || IsNil(o.CopySenderFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.CopySenderFlag.Get()
}

// GetCopySenderFlagOk returns a tuple with the CopySenderFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderStatusEmailTemplate) GetCopySenderFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopySenderFlag.Get(), o.CopySenderFlag.IsSet()
}

// HasCopySenderFlag returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasCopySenderFlag() bool {
	if o != nil && o.CopySenderFlag.IsSet() {
		return true
	}

	return false
}

// SetCopySenderFlag gets a reference to the given NullableBool and assigns it to the CopySenderFlag field.
func (o *PurchaseOrderStatusEmailTemplate) SetCopySenderFlag(v bool) {
	o.CopySenderFlag.Set(&v)
}
// SetCopySenderFlagNil sets the value for CopySenderFlag to be an explicit nil
func (o *PurchaseOrderStatusEmailTemplate) SetCopySenderFlagNil() {
	o.CopySenderFlag.Set(nil)
}

// UnsetCopySenderFlag ensures that no value is present for CopySenderFlag, not even an explicit nil
func (o *PurchaseOrderStatusEmailTemplate) UnsetCopySenderFlag() {
	o.CopySenderFlag.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *PurchaseOrderStatusEmailTemplate) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderStatusEmailTemplate) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *PurchaseOrderStatusEmailTemplate) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *PurchaseOrderStatusEmailTemplate) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o PurchaseOrderStatusEmailTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurchaseOrderStatusEmailTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.UseSenderFlag.IsSet() {
		toSerialize["useSenderFlag"] = o.UseSenderFlag.Get()
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if o.CopySenderFlag.IsSet() {
		toSerialize["copySenderFlag"] = o.CopySenderFlag.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *PurchaseOrderStatusEmailTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subject",
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

	varPurchaseOrderStatusEmailTemplate := _PurchaseOrderStatusEmailTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPurchaseOrderStatusEmailTemplate)

	if err != nil {
		return err
	}

	*o = PurchaseOrderStatusEmailTemplate(varPurchaseOrderStatusEmailTemplate)

	return err
}

type NullablePurchaseOrderStatusEmailTemplate struct {
	value *PurchaseOrderStatusEmailTemplate
	isSet bool
}

func (v NullablePurchaseOrderStatusEmailTemplate) Get() *PurchaseOrderStatusEmailTemplate {
	return v.value
}

func (v *NullablePurchaseOrderStatusEmailTemplate) Set(val *PurchaseOrderStatusEmailTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderStatusEmailTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderStatusEmailTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderStatusEmailTemplate(val *PurchaseOrderStatusEmailTemplate) *NullablePurchaseOrderStatusEmailTemplate {
	return &NullablePurchaseOrderStatusEmailTemplate{value: val, isSet: true}
}

func (v NullablePurchaseOrderStatusEmailTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderStatusEmailTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


