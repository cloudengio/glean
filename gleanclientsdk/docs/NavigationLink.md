# NavigationLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title to show for the navigation link | 
**Target** | **string** | Destination for the link | 
**External** | Pointer to **bool** | Hint indicating the link points to an external site | [optional] 
**ClientSideChecks** | Pointer to [**[]ClientSideCheck**](ClientSideCheck.md) | Client side checks to be performed before render | [optional] 

## Methods

### NewNavigationLink

`func NewNavigationLink(title string, target string, ) *NavigationLink`

NewNavigationLink instantiates a new NavigationLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationLinkWithDefaults

`func NewNavigationLinkWithDefaults() *NavigationLink`

NewNavigationLinkWithDefaults instantiates a new NavigationLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NavigationLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NavigationLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NavigationLink) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTarget

`func (o *NavigationLink) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NavigationLink) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NavigationLink) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetExternal

`func (o *NavigationLink) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *NavigationLink) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *NavigationLink) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *NavigationLink) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetClientSideChecks

`func (o *NavigationLink) GetClientSideChecks() []ClientSideCheck`

GetClientSideChecks returns the ClientSideChecks field if non-nil, zero value otherwise.

### GetClientSideChecksOk

`func (o *NavigationLink) GetClientSideChecksOk() (*[]ClientSideCheck, bool)`

GetClientSideChecksOk returns a tuple with the ClientSideChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideChecks

`func (o *NavigationLink) SetClientSideChecks(v []ClientSideCheck)`

SetClientSideChecks sets ClientSideChecks field to given value.

### HasClientSideChecks

`func (o *NavigationLink) HasClientSideChecks() bool`

HasClientSideChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


