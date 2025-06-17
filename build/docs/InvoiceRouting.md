# InvoiceRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ReviewedFlag** | Pointer to **bool** |  | [optional] 
**DateReviewedUTC** | Pointer to **time.Time** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceRouting

`func NewInvoiceRouting() *InvoiceRouting`

NewInvoiceRouting instantiates a new InvoiceRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRoutingWithDefaults

`func NewInvoiceRoutingWithDefaults() *InvoiceRouting`

NewInvoiceRoutingWithDefaults instantiates a new InvoiceRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceRouting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceRouting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceRouting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceRouting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoice

`func (o *InvoiceRouting) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceRouting) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceRouting) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceRouting) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *InvoiceRouting) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *InvoiceRouting) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *InvoiceRouting) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *InvoiceRouting) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *InvoiceRouting) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *InvoiceRouting) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetMember

`func (o *InvoiceRouting) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *InvoiceRouting) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *InvoiceRouting) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *InvoiceRouting) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetReviewedFlag

`func (o *InvoiceRouting) GetReviewedFlag() bool`

GetReviewedFlag returns the ReviewedFlag field if non-nil, zero value otherwise.

### GetReviewedFlagOk

`func (o *InvoiceRouting) GetReviewedFlagOk() (*bool, bool)`

GetReviewedFlagOk returns a tuple with the ReviewedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedFlag

`func (o *InvoiceRouting) SetReviewedFlag(v bool)`

SetReviewedFlag sets ReviewedFlag field to given value.

### HasReviewedFlag

`func (o *InvoiceRouting) HasReviewedFlag() bool`

HasReviewedFlag returns a boolean if a field has been set.

### GetDateReviewedUTC

`func (o *InvoiceRouting) GetDateReviewedUTC() time.Time`

GetDateReviewedUTC returns the DateReviewedUTC field if non-nil, zero value otherwise.

### GetDateReviewedUTCOk

`func (o *InvoiceRouting) GetDateReviewedUTCOk() (*time.Time, bool)`

GetDateReviewedUTCOk returns a tuple with the DateReviewedUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReviewedUTC

`func (o *InvoiceRouting) SetDateReviewedUTC(v time.Time)`

SetDateReviewedUTC sets DateReviewedUTC field to given value.

### HasDateReviewedUTC

`func (o *InvoiceRouting) HasDateReviewedUTC() bool`

HasDateReviewedUTC returns a boolean if a field has been set.

### GetInfo

`func (o *InvoiceRouting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceRouting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceRouting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceRouting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


