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

// checks if the ProcurementSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcurementSetting{}

// ProcurementSetting struct for ProcurementSetting
type ProcurementSetting struct {
	Id *int32 `json:"id,omitempty"`
	StartingPurchaseOrderNum int32 `json:"startingPurchaseOrderNum"`
	//  Max length: 5;
	PurchaseOrderPrefix *string `json:"purchaseOrderPrefix,omitempty"`
	//  Max length: 5;
	PurchaseOrderSuffix *string `json:"purchaseOrderSuffix,omitempty"`
	PrefixSuffixType NullableString `json:"prefixSuffixType,omitempty"`
	DisableCostUpdatesFlag NullableBool `json:"disableCostUpdatesFlag,omitempty"`
	DisableNegativeInventoryFlag NullableBool `json:"disableNegativeInventoryFlag,omitempty"`
	CostingMethod NullableString `json:"costingMethod"`
	AutoClosePurchaseOrderFlag NullableBool `json:"autoClosePurchaseOrderFlag,omitempty"`
	AutoClosePurchaseOrderItemFlag NullableBool `json:"autoClosePurchaseOrderItemFlag,omitempty"`
	AutoApprovePurchaseOrderFlag NullableBool `json:"autoApprovePurchaseOrderFlag,omitempty"`
	TaxPurchaseOrderFlag NullableBool `json:"taxPurchaseOrderFlag,omitempty"`
	TaxFreightFlag NullableBool `json:"taxFreightFlag,omitempty"`
	UseVendorTaxCodeFlag NullableBool `json:"useVendorTaxCodeFlag,omitempty"`
	NumDecimalPlaces NullableInt32 `json:"numDecimalPlaces,omitempty"`
	DisableAutoPickFlag NullableBool `json:"disableAutoPickFlag,omitempty"`
	DefaultProductTaxableFlag NullableBool `json:"defaultProductTaxableFlag,omitempty"`
	//  Max length: 50;
	EoriNumber *string `json:"eoriNumber,omitempty"`
	Info *map[string]string `json:"_info,omitempty"`
	NotificationForChangesInShippingInfoFlag NullableBool `json:"notificationForChangesInShippingInfoFlag,omitempty"`
	//  Max length: 250;
	ShippingInfoNotificationEmail *string `json:"shippingInfoNotificationEmail,omitempty"`
}

type _ProcurementSetting ProcurementSetting

// NewProcurementSetting instantiates a new ProcurementSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcurementSetting(startingPurchaseOrderNum int32, costingMethod NullableString) *ProcurementSetting {
	this := ProcurementSetting{}
	this.StartingPurchaseOrderNum = startingPurchaseOrderNum
	this.CostingMethod = costingMethod
	return &this
}

// NewProcurementSettingWithDefaults instantiates a new ProcurementSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcurementSettingWithDefaults() *ProcurementSetting {
	this := ProcurementSetting{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcurementSetting) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcurementSetting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProcurementSetting) SetId(v int32) {
	o.Id = &v
}

// GetStartingPurchaseOrderNum returns the StartingPurchaseOrderNum field value
func (o *ProcurementSetting) GetStartingPurchaseOrderNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartingPurchaseOrderNum
}

// GetStartingPurchaseOrderNumOk returns a tuple with the StartingPurchaseOrderNum field value
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetStartingPurchaseOrderNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartingPurchaseOrderNum, true
}

// SetStartingPurchaseOrderNum sets field value
func (o *ProcurementSetting) SetStartingPurchaseOrderNum(v int32) {
	o.StartingPurchaseOrderNum = v
}

// GetPurchaseOrderPrefix returns the PurchaseOrderPrefix field value if set, zero value otherwise.
func (o *ProcurementSetting) GetPurchaseOrderPrefix() string {
	if o == nil || IsNil(o.PurchaseOrderPrefix) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderPrefix
}

// GetPurchaseOrderPrefixOk returns a tuple with the PurchaseOrderPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetPurchaseOrderPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderPrefix) {
		return nil, false
	}
	return o.PurchaseOrderPrefix, true
}

// HasPurchaseOrderPrefix returns a boolean if a field has been set.
func (o *ProcurementSetting) HasPurchaseOrderPrefix() bool {
	if o != nil && !IsNil(o.PurchaseOrderPrefix) {
		return true
	}

	return false
}

// SetPurchaseOrderPrefix gets a reference to the given string and assigns it to the PurchaseOrderPrefix field.
func (o *ProcurementSetting) SetPurchaseOrderPrefix(v string) {
	o.PurchaseOrderPrefix = &v
}

// GetPurchaseOrderSuffix returns the PurchaseOrderSuffix field value if set, zero value otherwise.
func (o *ProcurementSetting) GetPurchaseOrderSuffix() string {
	if o == nil || IsNil(o.PurchaseOrderSuffix) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderSuffix
}

// GetPurchaseOrderSuffixOk returns a tuple with the PurchaseOrderSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetPurchaseOrderSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderSuffix) {
		return nil, false
	}
	return o.PurchaseOrderSuffix, true
}

// HasPurchaseOrderSuffix returns a boolean if a field has been set.
func (o *ProcurementSetting) HasPurchaseOrderSuffix() bool {
	if o != nil && !IsNil(o.PurchaseOrderSuffix) {
		return true
	}

	return false
}

// SetPurchaseOrderSuffix gets a reference to the given string and assigns it to the PurchaseOrderSuffix field.
func (o *ProcurementSetting) SetPurchaseOrderSuffix(v string) {
	o.PurchaseOrderSuffix = &v
}

// GetPrefixSuffixType returns the PrefixSuffixType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetPrefixSuffixType() string {
	if o == nil || IsNil(o.PrefixSuffixType.Get()) {
		var ret string
		return ret
	}
	return *o.PrefixSuffixType.Get()
}

// GetPrefixSuffixTypeOk returns a tuple with the PrefixSuffixType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetPrefixSuffixTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrefixSuffixType.Get(), o.PrefixSuffixType.IsSet()
}

// HasPrefixSuffixType returns a boolean if a field has been set.
func (o *ProcurementSetting) HasPrefixSuffixType() bool {
	if o != nil && o.PrefixSuffixType.IsSet() {
		return true
	}

	return false
}

// SetPrefixSuffixType gets a reference to the given NullableString and assigns it to the PrefixSuffixType field.
func (o *ProcurementSetting) SetPrefixSuffixType(v string) {
	o.PrefixSuffixType.Set(&v)
}
// SetPrefixSuffixTypeNil sets the value for PrefixSuffixType to be an explicit nil
func (o *ProcurementSetting) SetPrefixSuffixTypeNil() {
	o.PrefixSuffixType.Set(nil)
}

// UnsetPrefixSuffixType ensures that no value is present for PrefixSuffixType, not even an explicit nil
func (o *ProcurementSetting) UnsetPrefixSuffixType() {
	o.PrefixSuffixType.Unset()
}

// GetDisableCostUpdatesFlag returns the DisableCostUpdatesFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetDisableCostUpdatesFlag() bool {
	if o == nil || IsNil(o.DisableCostUpdatesFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DisableCostUpdatesFlag.Get()
}

// GetDisableCostUpdatesFlagOk returns a tuple with the DisableCostUpdatesFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetDisableCostUpdatesFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisableCostUpdatesFlag.Get(), o.DisableCostUpdatesFlag.IsSet()
}

// HasDisableCostUpdatesFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasDisableCostUpdatesFlag() bool {
	if o != nil && o.DisableCostUpdatesFlag.IsSet() {
		return true
	}

	return false
}

// SetDisableCostUpdatesFlag gets a reference to the given NullableBool and assigns it to the DisableCostUpdatesFlag field.
func (o *ProcurementSetting) SetDisableCostUpdatesFlag(v bool) {
	o.DisableCostUpdatesFlag.Set(&v)
}
// SetDisableCostUpdatesFlagNil sets the value for DisableCostUpdatesFlag to be an explicit nil
func (o *ProcurementSetting) SetDisableCostUpdatesFlagNil() {
	o.DisableCostUpdatesFlag.Set(nil)
}

// UnsetDisableCostUpdatesFlag ensures that no value is present for DisableCostUpdatesFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetDisableCostUpdatesFlag() {
	o.DisableCostUpdatesFlag.Unset()
}

// GetDisableNegativeInventoryFlag returns the DisableNegativeInventoryFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetDisableNegativeInventoryFlag() bool {
	if o == nil || IsNil(o.DisableNegativeInventoryFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DisableNegativeInventoryFlag.Get()
}

// GetDisableNegativeInventoryFlagOk returns a tuple with the DisableNegativeInventoryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetDisableNegativeInventoryFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisableNegativeInventoryFlag.Get(), o.DisableNegativeInventoryFlag.IsSet()
}

// HasDisableNegativeInventoryFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasDisableNegativeInventoryFlag() bool {
	if o != nil && o.DisableNegativeInventoryFlag.IsSet() {
		return true
	}

	return false
}

// SetDisableNegativeInventoryFlag gets a reference to the given NullableBool and assigns it to the DisableNegativeInventoryFlag field.
func (o *ProcurementSetting) SetDisableNegativeInventoryFlag(v bool) {
	o.DisableNegativeInventoryFlag.Set(&v)
}
// SetDisableNegativeInventoryFlagNil sets the value for DisableNegativeInventoryFlag to be an explicit nil
func (o *ProcurementSetting) SetDisableNegativeInventoryFlagNil() {
	o.DisableNegativeInventoryFlag.Set(nil)
}

// UnsetDisableNegativeInventoryFlag ensures that no value is present for DisableNegativeInventoryFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetDisableNegativeInventoryFlag() {
	o.DisableNegativeInventoryFlag.Unset()
}

// GetCostingMethod returns the CostingMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProcurementSetting) GetCostingMethod() string {
	if o == nil || o.CostingMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.CostingMethod.Get()
}

// GetCostingMethodOk returns a tuple with the CostingMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetCostingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostingMethod.Get(), o.CostingMethod.IsSet()
}

// SetCostingMethod sets field value
func (o *ProcurementSetting) SetCostingMethod(v string) {
	o.CostingMethod.Set(&v)
}

// GetAutoClosePurchaseOrderFlag returns the AutoClosePurchaseOrderFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetAutoClosePurchaseOrderFlag() bool {
	if o == nil || IsNil(o.AutoClosePurchaseOrderFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoClosePurchaseOrderFlag.Get()
}

// GetAutoClosePurchaseOrderFlagOk returns a tuple with the AutoClosePurchaseOrderFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetAutoClosePurchaseOrderFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoClosePurchaseOrderFlag.Get(), o.AutoClosePurchaseOrderFlag.IsSet()
}

// HasAutoClosePurchaseOrderFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasAutoClosePurchaseOrderFlag() bool {
	if o != nil && o.AutoClosePurchaseOrderFlag.IsSet() {
		return true
	}

	return false
}

// SetAutoClosePurchaseOrderFlag gets a reference to the given NullableBool and assigns it to the AutoClosePurchaseOrderFlag field.
func (o *ProcurementSetting) SetAutoClosePurchaseOrderFlag(v bool) {
	o.AutoClosePurchaseOrderFlag.Set(&v)
}
// SetAutoClosePurchaseOrderFlagNil sets the value for AutoClosePurchaseOrderFlag to be an explicit nil
func (o *ProcurementSetting) SetAutoClosePurchaseOrderFlagNil() {
	o.AutoClosePurchaseOrderFlag.Set(nil)
}

// UnsetAutoClosePurchaseOrderFlag ensures that no value is present for AutoClosePurchaseOrderFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetAutoClosePurchaseOrderFlag() {
	o.AutoClosePurchaseOrderFlag.Unset()
}

// GetAutoClosePurchaseOrderItemFlag returns the AutoClosePurchaseOrderItemFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetAutoClosePurchaseOrderItemFlag() bool {
	if o == nil || IsNil(o.AutoClosePurchaseOrderItemFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoClosePurchaseOrderItemFlag.Get()
}

// GetAutoClosePurchaseOrderItemFlagOk returns a tuple with the AutoClosePurchaseOrderItemFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetAutoClosePurchaseOrderItemFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoClosePurchaseOrderItemFlag.Get(), o.AutoClosePurchaseOrderItemFlag.IsSet()
}

// HasAutoClosePurchaseOrderItemFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasAutoClosePurchaseOrderItemFlag() bool {
	if o != nil && o.AutoClosePurchaseOrderItemFlag.IsSet() {
		return true
	}

	return false
}

// SetAutoClosePurchaseOrderItemFlag gets a reference to the given NullableBool and assigns it to the AutoClosePurchaseOrderItemFlag field.
func (o *ProcurementSetting) SetAutoClosePurchaseOrderItemFlag(v bool) {
	o.AutoClosePurchaseOrderItemFlag.Set(&v)
}
// SetAutoClosePurchaseOrderItemFlagNil sets the value for AutoClosePurchaseOrderItemFlag to be an explicit nil
func (o *ProcurementSetting) SetAutoClosePurchaseOrderItemFlagNil() {
	o.AutoClosePurchaseOrderItemFlag.Set(nil)
}

// UnsetAutoClosePurchaseOrderItemFlag ensures that no value is present for AutoClosePurchaseOrderItemFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetAutoClosePurchaseOrderItemFlag() {
	o.AutoClosePurchaseOrderItemFlag.Unset()
}

// GetAutoApprovePurchaseOrderFlag returns the AutoApprovePurchaseOrderFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetAutoApprovePurchaseOrderFlag() bool {
	if o == nil || IsNil(o.AutoApprovePurchaseOrderFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoApprovePurchaseOrderFlag.Get()
}

// GetAutoApprovePurchaseOrderFlagOk returns a tuple with the AutoApprovePurchaseOrderFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetAutoApprovePurchaseOrderFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoApprovePurchaseOrderFlag.Get(), o.AutoApprovePurchaseOrderFlag.IsSet()
}

// HasAutoApprovePurchaseOrderFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasAutoApprovePurchaseOrderFlag() bool {
	if o != nil && o.AutoApprovePurchaseOrderFlag.IsSet() {
		return true
	}

	return false
}

// SetAutoApprovePurchaseOrderFlag gets a reference to the given NullableBool and assigns it to the AutoApprovePurchaseOrderFlag field.
func (o *ProcurementSetting) SetAutoApprovePurchaseOrderFlag(v bool) {
	o.AutoApprovePurchaseOrderFlag.Set(&v)
}
// SetAutoApprovePurchaseOrderFlagNil sets the value for AutoApprovePurchaseOrderFlag to be an explicit nil
func (o *ProcurementSetting) SetAutoApprovePurchaseOrderFlagNil() {
	o.AutoApprovePurchaseOrderFlag.Set(nil)
}

// UnsetAutoApprovePurchaseOrderFlag ensures that no value is present for AutoApprovePurchaseOrderFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetAutoApprovePurchaseOrderFlag() {
	o.AutoApprovePurchaseOrderFlag.Unset()
}

// GetTaxPurchaseOrderFlag returns the TaxPurchaseOrderFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetTaxPurchaseOrderFlag() bool {
	if o == nil || IsNil(o.TaxPurchaseOrderFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.TaxPurchaseOrderFlag.Get()
}

// GetTaxPurchaseOrderFlagOk returns a tuple with the TaxPurchaseOrderFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetTaxPurchaseOrderFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxPurchaseOrderFlag.Get(), o.TaxPurchaseOrderFlag.IsSet()
}

// HasTaxPurchaseOrderFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasTaxPurchaseOrderFlag() bool {
	if o != nil && o.TaxPurchaseOrderFlag.IsSet() {
		return true
	}

	return false
}

// SetTaxPurchaseOrderFlag gets a reference to the given NullableBool and assigns it to the TaxPurchaseOrderFlag field.
func (o *ProcurementSetting) SetTaxPurchaseOrderFlag(v bool) {
	o.TaxPurchaseOrderFlag.Set(&v)
}
// SetTaxPurchaseOrderFlagNil sets the value for TaxPurchaseOrderFlag to be an explicit nil
func (o *ProcurementSetting) SetTaxPurchaseOrderFlagNil() {
	o.TaxPurchaseOrderFlag.Set(nil)
}

// UnsetTaxPurchaseOrderFlag ensures that no value is present for TaxPurchaseOrderFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetTaxPurchaseOrderFlag() {
	o.TaxPurchaseOrderFlag.Unset()
}

// GetTaxFreightFlag returns the TaxFreightFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetTaxFreightFlag() bool {
	if o == nil || IsNil(o.TaxFreightFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.TaxFreightFlag.Get()
}

// GetTaxFreightFlagOk returns a tuple with the TaxFreightFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetTaxFreightFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxFreightFlag.Get(), o.TaxFreightFlag.IsSet()
}

// HasTaxFreightFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasTaxFreightFlag() bool {
	if o != nil && o.TaxFreightFlag.IsSet() {
		return true
	}

	return false
}

// SetTaxFreightFlag gets a reference to the given NullableBool and assigns it to the TaxFreightFlag field.
func (o *ProcurementSetting) SetTaxFreightFlag(v bool) {
	o.TaxFreightFlag.Set(&v)
}
// SetTaxFreightFlagNil sets the value for TaxFreightFlag to be an explicit nil
func (o *ProcurementSetting) SetTaxFreightFlagNil() {
	o.TaxFreightFlag.Set(nil)
}

// UnsetTaxFreightFlag ensures that no value is present for TaxFreightFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetTaxFreightFlag() {
	o.TaxFreightFlag.Unset()
}

// GetUseVendorTaxCodeFlag returns the UseVendorTaxCodeFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetUseVendorTaxCodeFlag() bool {
	if o == nil || IsNil(o.UseVendorTaxCodeFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.UseVendorTaxCodeFlag.Get()
}

// GetUseVendorTaxCodeFlagOk returns a tuple with the UseVendorTaxCodeFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetUseVendorTaxCodeFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseVendorTaxCodeFlag.Get(), o.UseVendorTaxCodeFlag.IsSet()
}

// HasUseVendorTaxCodeFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasUseVendorTaxCodeFlag() bool {
	if o != nil && o.UseVendorTaxCodeFlag.IsSet() {
		return true
	}

	return false
}

// SetUseVendorTaxCodeFlag gets a reference to the given NullableBool and assigns it to the UseVendorTaxCodeFlag field.
func (o *ProcurementSetting) SetUseVendorTaxCodeFlag(v bool) {
	o.UseVendorTaxCodeFlag.Set(&v)
}
// SetUseVendorTaxCodeFlagNil sets the value for UseVendorTaxCodeFlag to be an explicit nil
func (o *ProcurementSetting) SetUseVendorTaxCodeFlagNil() {
	o.UseVendorTaxCodeFlag.Set(nil)
}

// UnsetUseVendorTaxCodeFlag ensures that no value is present for UseVendorTaxCodeFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetUseVendorTaxCodeFlag() {
	o.UseVendorTaxCodeFlag.Unset()
}

// GetNumDecimalPlaces returns the NumDecimalPlaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetNumDecimalPlaces() int32 {
	if o == nil || IsNil(o.NumDecimalPlaces.Get()) {
		var ret int32
		return ret
	}
	return *o.NumDecimalPlaces.Get()
}

// GetNumDecimalPlacesOk returns a tuple with the NumDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetNumDecimalPlacesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumDecimalPlaces.Get(), o.NumDecimalPlaces.IsSet()
}

// HasNumDecimalPlaces returns a boolean if a field has been set.
func (o *ProcurementSetting) HasNumDecimalPlaces() bool {
	if o != nil && o.NumDecimalPlaces.IsSet() {
		return true
	}

	return false
}

// SetNumDecimalPlaces gets a reference to the given NullableInt32 and assigns it to the NumDecimalPlaces field.
func (o *ProcurementSetting) SetNumDecimalPlaces(v int32) {
	o.NumDecimalPlaces.Set(&v)
}
// SetNumDecimalPlacesNil sets the value for NumDecimalPlaces to be an explicit nil
func (o *ProcurementSetting) SetNumDecimalPlacesNil() {
	o.NumDecimalPlaces.Set(nil)
}

// UnsetNumDecimalPlaces ensures that no value is present for NumDecimalPlaces, not even an explicit nil
func (o *ProcurementSetting) UnsetNumDecimalPlaces() {
	o.NumDecimalPlaces.Unset()
}

// GetDisableAutoPickFlag returns the DisableAutoPickFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetDisableAutoPickFlag() bool {
	if o == nil || IsNil(o.DisableAutoPickFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DisableAutoPickFlag.Get()
}

// GetDisableAutoPickFlagOk returns a tuple with the DisableAutoPickFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetDisableAutoPickFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisableAutoPickFlag.Get(), o.DisableAutoPickFlag.IsSet()
}

// HasDisableAutoPickFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasDisableAutoPickFlag() bool {
	if o != nil && o.DisableAutoPickFlag.IsSet() {
		return true
	}

	return false
}

// SetDisableAutoPickFlag gets a reference to the given NullableBool and assigns it to the DisableAutoPickFlag field.
func (o *ProcurementSetting) SetDisableAutoPickFlag(v bool) {
	o.DisableAutoPickFlag.Set(&v)
}
// SetDisableAutoPickFlagNil sets the value for DisableAutoPickFlag to be an explicit nil
func (o *ProcurementSetting) SetDisableAutoPickFlagNil() {
	o.DisableAutoPickFlag.Set(nil)
}

// UnsetDisableAutoPickFlag ensures that no value is present for DisableAutoPickFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetDisableAutoPickFlag() {
	o.DisableAutoPickFlag.Unset()
}

// GetDefaultProductTaxableFlag returns the DefaultProductTaxableFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetDefaultProductTaxableFlag() bool {
	if o == nil || IsNil(o.DefaultProductTaxableFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.DefaultProductTaxableFlag.Get()
}

// GetDefaultProductTaxableFlagOk returns a tuple with the DefaultProductTaxableFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetDefaultProductTaxableFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultProductTaxableFlag.Get(), o.DefaultProductTaxableFlag.IsSet()
}

// HasDefaultProductTaxableFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasDefaultProductTaxableFlag() bool {
	if o != nil && o.DefaultProductTaxableFlag.IsSet() {
		return true
	}

	return false
}

// SetDefaultProductTaxableFlag gets a reference to the given NullableBool and assigns it to the DefaultProductTaxableFlag field.
func (o *ProcurementSetting) SetDefaultProductTaxableFlag(v bool) {
	o.DefaultProductTaxableFlag.Set(&v)
}
// SetDefaultProductTaxableFlagNil sets the value for DefaultProductTaxableFlag to be an explicit nil
func (o *ProcurementSetting) SetDefaultProductTaxableFlagNil() {
	o.DefaultProductTaxableFlag.Set(nil)
}

// UnsetDefaultProductTaxableFlag ensures that no value is present for DefaultProductTaxableFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetDefaultProductTaxableFlag() {
	o.DefaultProductTaxableFlag.Unset()
}

// GetEoriNumber returns the EoriNumber field value if set, zero value otherwise.
func (o *ProcurementSetting) GetEoriNumber() string {
	if o == nil || IsNil(o.EoriNumber) {
		var ret string
		return ret
	}
	return *o.EoriNumber
}

// GetEoriNumberOk returns a tuple with the EoriNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetEoriNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EoriNumber) {
		return nil, false
	}
	return o.EoriNumber, true
}

// HasEoriNumber returns a boolean if a field has been set.
func (o *ProcurementSetting) HasEoriNumber() bool {
	if o != nil && !IsNil(o.EoriNumber) {
		return true
	}

	return false
}

// SetEoriNumber gets a reference to the given string and assigns it to the EoriNumber field.
func (o *ProcurementSetting) SetEoriNumber(v string) {
	o.EoriNumber = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ProcurementSetting) GetInfo() map[string]string {
	if o == nil || IsNil(o.Info) {
		var ret map[string]string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ProcurementSetting) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given map[string]string and assigns it to the Info field.
func (o *ProcurementSetting) SetInfo(v map[string]string) {
	o.Info = &v
}

// GetNotificationForChangesInShippingInfoFlag returns the NotificationForChangesInShippingInfoFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProcurementSetting) GetNotificationForChangesInShippingInfoFlag() bool {
	if o == nil || IsNil(o.NotificationForChangesInShippingInfoFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationForChangesInShippingInfoFlag.Get()
}

// GetNotificationForChangesInShippingInfoFlagOk returns a tuple with the NotificationForChangesInShippingInfoFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProcurementSetting) GetNotificationForChangesInShippingInfoFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationForChangesInShippingInfoFlag.Get(), o.NotificationForChangesInShippingInfoFlag.IsSet()
}

// HasNotificationForChangesInShippingInfoFlag returns a boolean if a field has been set.
func (o *ProcurementSetting) HasNotificationForChangesInShippingInfoFlag() bool {
	if o != nil && o.NotificationForChangesInShippingInfoFlag.IsSet() {
		return true
	}

	return false
}

// SetNotificationForChangesInShippingInfoFlag gets a reference to the given NullableBool and assigns it to the NotificationForChangesInShippingInfoFlag field.
func (o *ProcurementSetting) SetNotificationForChangesInShippingInfoFlag(v bool) {
	o.NotificationForChangesInShippingInfoFlag.Set(&v)
}
// SetNotificationForChangesInShippingInfoFlagNil sets the value for NotificationForChangesInShippingInfoFlag to be an explicit nil
func (o *ProcurementSetting) SetNotificationForChangesInShippingInfoFlagNil() {
	o.NotificationForChangesInShippingInfoFlag.Set(nil)
}

// UnsetNotificationForChangesInShippingInfoFlag ensures that no value is present for NotificationForChangesInShippingInfoFlag, not even an explicit nil
func (o *ProcurementSetting) UnsetNotificationForChangesInShippingInfoFlag() {
	o.NotificationForChangesInShippingInfoFlag.Unset()
}

// GetShippingInfoNotificationEmail returns the ShippingInfoNotificationEmail field value if set, zero value otherwise.
func (o *ProcurementSetting) GetShippingInfoNotificationEmail() string {
	if o == nil || IsNil(o.ShippingInfoNotificationEmail) {
		var ret string
		return ret
	}
	return *o.ShippingInfoNotificationEmail
}

// GetShippingInfoNotificationEmailOk returns a tuple with the ShippingInfoNotificationEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcurementSetting) GetShippingInfoNotificationEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingInfoNotificationEmail) {
		return nil, false
	}
	return o.ShippingInfoNotificationEmail, true
}

// HasShippingInfoNotificationEmail returns a boolean if a field has been set.
func (o *ProcurementSetting) HasShippingInfoNotificationEmail() bool {
	if o != nil && !IsNil(o.ShippingInfoNotificationEmail) {
		return true
	}

	return false
}

// SetShippingInfoNotificationEmail gets a reference to the given string and assigns it to the ShippingInfoNotificationEmail field.
func (o *ProcurementSetting) SetShippingInfoNotificationEmail(v string) {
	o.ShippingInfoNotificationEmail = &v
}

func (o ProcurementSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcurementSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["startingPurchaseOrderNum"] = o.StartingPurchaseOrderNum
	if !IsNil(o.PurchaseOrderPrefix) {
		toSerialize["purchaseOrderPrefix"] = o.PurchaseOrderPrefix
	}
	if !IsNil(o.PurchaseOrderSuffix) {
		toSerialize["purchaseOrderSuffix"] = o.PurchaseOrderSuffix
	}
	if o.PrefixSuffixType.IsSet() {
		toSerialize["prefixSuffixType"] = o.PrefixSuffixType.Get()
	}
	if o.DisableCostUpdatesFlag.IsSet() {
		toSerialize["disableCostUpdatesFlag"] = o.DisableCostUpdatesFlag.Get()
	}
	if o.DisableNegativeInventoryFlag.IsSet() {
		toSerialize["disableNegativeInventoryFlag"] = o.DisableNegativeInventoryFlag.Get()
	}
	toSerialize["costingMethod"] = o.CostingMethod.Get()
	if o.AutoClosePurchaseOrderFlag.IsSet() {
		toSerialize["autoClosePurchaseOrderFlag"] = o.AutoClosePurchaseOrderFlag.Get()
	}
	if o.AutoClosePurchaseOrderItemFlag.IsSet() {
		toSerialize["autoClosePurchaseOrderItemFlag"] = o.AutoClosePurchaseOrderItemFlag.Get()
	}
	if o.AutoApprovePurchaseOrderFlag.IsSet() {
		toSerialize["autoApprovePurchaseOrderFlag"] = o.AutoApprovePurchaseOrderFlag.Get()
	}
	if o.TaxPurchaseOrderFlag.IsSet() {
		toSerialize["taxPurchaseOrderFlag"] = o.TaxPurchaseOrderFlag.Get()
	}
	if o.TaxFreightFlag.IsSet() {
		toSerialize["taxFreightFlag"] = o.TaxFreightFlag.Get()
	}
	if o.UseVendorTaxCodeFlag.IsSet() {
		toSerialize["useVendorTaxCodeFlag"] = o.UseVendorTaxCodeFlag.Get()
	}
	if o.NumDecimalPlaces.IsSet() {
		toSerialize["numDecimalPlaces"] = o.NumDecimalPlaces.Get()
	}
	if o.DisableAutoPickFlag.IsSet() {
		toSerialize["disableAutoPickFlag"] = o.DisableAutoPickFlag.Get()
	}
	if o.DefaultProductTaxableFlag.IsSet() {
		toSerialize["defaultProductTaxableFlag"] = o.DefaultProductTaxableFlag.Get()
	}
	if !IsNil(o.EoriNumber) {
		toSerialize["eoriNumber"] = o.EoriNumber
	}
	if !IsNil(o.Info) {
		toSerialize["_info"] = o.Info
	}
	if o.NotificationForChangesInShippingInfoFlag.IsSet() {
		toSerialize["notificationForChangesInShippingInfoFlag"] = o.NotificationForChangesInShippingInfoFlag.Get()
	}
	if !IsNil(o.ShippingInfoNotificationEmail) {
		toSerialize["shippingInfoNotificationEmail"] = o.ShippingInfoNotificationEmail
	}
	return toSerialize, nil
}

func (o *ProcurementSetting) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startingPurchaseOrderNum",
		"costingMethod",
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

	varProcurementSetting := _ProcurementSetting{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProcurementSetting)

	if err != nil {
		return err
	}

	*o = ProcurementSetting(varProcurementSetting)

	return err
}

type NullableProcurementSetting struct {
	value *ProcurementSetting
	isSet bool
}

func (v NullableProcurementSetting) Get() *ProcurementSetting {
	return v.value
}

func (v *NullableProcurementSetting) Set(val *ProcurementSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableProcurementSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableProcurementSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcurementSetting(val *ProcurementSetting) *NullableProcurementSetting {
	return &NullableProcurementSetting{value: val, isSet: true}
}

func (v NullableProcurementSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcurementSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


