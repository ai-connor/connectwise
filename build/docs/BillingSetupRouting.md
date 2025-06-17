# BillingSetupRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | **NullableInt32** |  | 
**InvoiceRule** | **NullableString** |  | 
**RoutingRule** | **NullableString** |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingSetupRouting

`func NewBillingSetupRouting(sequenceNumber NullableInt32, invoiceRule NullableString, routingRule NullableString, ) *BillingSetupRouting`

NewBillingSetupRouting instantiates a new BillingSetupRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSetupRoutingWithDefaults

`func NewBillingSetupRoutingWithDefaults() *BillingSetupRouting`

NewBillingSetupRoutingWithDefaults instantiates a new BillingSetupRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingSetupRouting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSetupRouting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSetupRouting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingSetupRouting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *BillingSetupRouting) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *BillingSetupRouting) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *BillingSetupRouting) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### SetSequenceNumberNil

`func (o *BillingSetupRouting) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *BillingSetupRouting) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetInvoiceRule

`func (o *BillingSetupRouting) GetInvoiceRule() string`

GetInvoiceRule returns the InvoiceRule field if non-nil, zero value otherwise.

### GetInvoiceRuleOk

`func (o *BillingSetupRouting) GetInvoiceRuleOk() (*string, bool)`

GetInvoiceRuleOk returns a tuple with the InvoiceRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRule

`func (o *BillingSetupRouting) SetInvoiceRule(v string)`

SetInvoiceRule sets InvoiceRule field to given value.


### SetInvoiceRuleNil

`func (o *BillingSetupRouting) SetInvoiceRuleNil(b bool)`

 SetInvoiceRuleNil sets the value for InvoiceRule to be an explicit nil

### UnsetInvoiceRule
`func (o *BillingSetupRouting) UnsetInvoiceRule()`

UnsetInvoiceRule ensures that no value is present for InvoiceRule, not even an explicit nil
### GetRoutingRule

`func (o *BillingSetupRouting) GetRoutingRule() string`

GetRoutingRule returns the RoutingRule field if non-nil, zero value otherwise.

### GetRoutingRuleOk

`func (o *BillingSetupRouting) GetRoutingRuleOk() (*string, bool)`

GetRoutingRuleOk returns a tuple with the RoutingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRule

`func (o *BillingSetupRouting) SetRoutingRule(v string)`

SetRoutingRule sets RoutingRule field to given value.


### SetRoutingRuleNil

`func (o *BillingSetupRouting) SetRoutingRuleNil(b bool)`

 SetRoutingRuleNil sets the value for RoutingRule to be an explicit nil

### UnsetRoutingRule
`func (o *BillingSetupRouting) UnsetRoutingRule()`

UnsetRoutingRule ensures that no value is present for RoutingRule, not even an explicit nil
### GetMember

`func (o *BillingSetupRouting) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *BillingSetupRouting) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *BillingSetupRouting) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *BillingSetupRouting) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *BillingSetupRouting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingSetupRouting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingSetupRouting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingSetupRouting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


