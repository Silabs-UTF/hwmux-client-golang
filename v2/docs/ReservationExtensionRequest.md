# ReservationExtensionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionDurationS** | **int32** |  | 
**ExtendExisting** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewReservationExtensionRequest

`func NewReservationExtensionRequest(extensionDurationS int32, ) *ReservationExtensionRequest`

NewReservationExtensionRequest instantiates a new ReservationExtensionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationExtensionRequestWithDefaults

`func NewReservationExtensionRequestWithDefaults() *ReservationExtensionRequest`

NewReservationExtensionRequestWithDefaults instantiates a new ReservationExtensionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionDurationS

`func (o *ReservationExtensionRequest) GetExtensionDurationS() int32`

GetExtensionDurationS returns the ExtensionDurationS field if non-nil, zero value otherwise.

### GetExtensionDurationSOk

`func (o *ReservationExtensionRequest) GetExtensionDurationSOk() (*int32, bool)`

GetExtensionDurationSOk returns a tuple with the ExtensionDurationS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionDurationS

`func (o *ReservationExtensionRequest) SetExtensionDurationS(v int32)`

SetExtensionDurationS sets ExtensionDurationS field to given value.


### GetExtendExisting

`func (o *ReservationExtensionRequest) GetExtendExisting() bool`

GetExtendExisting returns the ExtendExisting field if non-nil, zero value otherwise.

### GetExtendExistingOk

`func (o *ReservationExtensionRequest) GetExtendExistingOk() (*bool, bool)`

GetExtendExistingOk returns a tuple with the ExtendExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendExisting

`func (o *ReservationExtensionRequest) SetExtendExisting(v bool)`

SetExtendExisting sets ExtendExisting field to given value.

### HasExtendExisting

`func (o *ReservationExtensionRequest) HasExtendExisting() bool`

HasExtendExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


