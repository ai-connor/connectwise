# ClosedInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**BillingStatusReference**](BillingStatusReference.md) |  | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewClosedInvoice

`func NewClosedInvoice() *ClosedInvoice`

NewClosedInvoice instantiates a new ClosedInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClosedInvoiceWithDefaults

`func NewClosedInvoiceWithDefaults() *ClosedInvoice`

NewClosedInvoiceWithDefaults instantiates a new ClosedInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClosedInvoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClosedInvoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClosedInvoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClosedInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ClosedInvoice) GetStatus() BillingStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClosedInvoice) GetStatusOk() (*BillingStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClosedInvoice) SetStatus(v BillingStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClosedInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInternalNotes

`func (o *ClosedInvoice) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *ClosedInvoice) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *ClosedInvoice) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *ClosedInvoice) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetInfo

`func (o *ClosedInvoice) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ClosedInvoice) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ClosedInvoice) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ClosedInvoice) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


