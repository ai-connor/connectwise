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

// checks if the OrderStatusNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusNotification{}

// OrderStatusNotification struct for OrderStatusNotification
type OrderStatusNotification struct {
	Id *int32 `json:"id,omitempty"`
	NotifyWho NotificationRecipientReference `json:"notifyWho"`
	Status *OrderStatusReference `json:"status,omitempty"`
	Member *MemberReference `json:"member,omitempty"`
	// Order Status Notification sendEmail must be entered if the notify type is \"Email Address\". Max length: 50;
	Email *string `json:"email,omitempty"`
	WorkflowStep NullableInt32 `json:"workflowStep,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
}

type _OrderStatusNotification OrderStatusNotification

// NewOrderStatusNotification instantiates a new OrderStatusNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusNotification(notifyWho NotificationRecipientReference) *OrderStatusNotification {
	this := OrderStatusNotification{}
	this.NotifyWho = notifyWho
	return &this
}

// NewOrderStatusNotificationWithDefaults instantiates a new OrderStatusNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusNotificationWithDefaults() *OrderStatusNotification {
	this := OrderStatusNotification{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderStatusNotification) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusNotification) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderStatusNotification) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrderStatusNotification) SetId(v int32) {
	o.Id = &v
}

// GetNotifyWho returns the NotifyWho field value
func (o *OrderStatusNotification) GetNotifyWho() NotificationRecipientReference {
	if o == nil {
		var ret NotificationRecipientReference
		return ret
	}

	return o.NotifyWho
}

// GetNotifyWhoOk returns a tuple with the NotifyWho field value
// and a boolean to check if the value has been set.
func (o *OrderStatusNotification) GetNotifyWhoOk() (*NotificationRecipientReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyWho, true
}

// SetNotifyWho sets field value
func (o *OrderStatusNotification) SetNotifyWho(v NotificationRecipientReference) {
	o.NotifyWho = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderStatusNotification) GetStatus() OrderStatusReference {
	if o == nil || IsNil(o.Status) {
		var ret OrderStatusReference
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusNotification) GetStatusOk() (*OrderStatusReference, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderStatusNotification) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrderStatusReference and assigns it to the Status field.
func (o *OrderStatusNotification) SetStatus(v OrderStatusReference) {
	o.Status = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *OrderStatusNotification) GetMember() MemberReference {
	if o == nil || IsNil(o.Member) {
		var ret MemberReference
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusNotification) GetMemberOk() (*MemberReference, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *OrderStatusNotification) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given MemberReference and assigns it to the Member field.
func (o *OrderStatusNotification) SetMember(v MemberReference) {
	o.Member = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderStatusNotification) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusNotification) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderStatusNotification) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderStatusNotification) SetEmail(v string) {
	o.Email = &v
}

// GetWorkflowStep returns the WorkflowStep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderStatusNotification) GetWorkflowStep() int32 {
	if o == nil || IsNil(o.WorkflowStep.Get()) {
		var ret int32
		return ret
	}
	return *o.WorkflowStep.Get()
}

// GetWorkflowStepOk returns a tuple with the WorkflowStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderStatusNotification) GetWorkflowStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowStep.Get(), o.WorkflowStep.IsSet()
}

// HasWorkflowStep returns a boolean if a field has been set.
func (o *OrderStatusNotification) HasWorkflowStep() bool {
	if o != nil && o.WorkflowStep.IsSet() {
		return true
	}

	return false
}

// SetWorkflowStep gets a reference to the given NullableInt32 and assigns it to the WorkflowStep field.
func (o *OrderStatusNotification) SetWorkflowStep(v int32) {
	o.WorkflowStep.Set(&v)
}
// SetWorkflowStepNil sets the value for WorkflowStep to be an explicit nil
func (o *OrderStatusNotification) SetWorkflowStepNil() {
	o.WorkflowStep.Set(nil)
}

// UnsetWorkflowStep ensures that no value is present for WorkflowStep, not even an explicit nil
func (o *OrderStatusNotification) UnsetWorkflowStep() {
	o.WorkflowStep.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *OrderStatusNotification) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusNotification) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *OrderStatusNotification) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *OrderStatusNotification) SetInfo(v map[string]string) {
	o.Info = &v
}

func (o OrderStatusNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["notifyWho"] = o.NotifyWho
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if o.WorkflowStep.IsSet() {
		toSerialize["workflowStep"] = o.WorkflowStep.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	return toSerialize, nil
}

func (o *OrderStatusNotification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifyWho",
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

	varOrderStatusNotification := _OrderStatusNotification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderStatusNotification)

	if err != nil {
		return err
	}

	*o = OrderStatusNotification(varOrderStatusNotification)

	return err
}

type NullableOrderStatusNotification struct {
	value *OrderStatusNotification
	isSet bool
}

func (v NullableOrderStatusNotification) Get() *OrderStatusNotification {
	return v.value
}

func (v *NullableOrderStatusNotification) Set(val *OrderStatusNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusNotification(val *OrderStatusNotification) *NullableOrderStatusNotification {
	return &NullableOrderStatusNotification{value: val, isSet: true}
}

func (v NullableOrderStatusNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


