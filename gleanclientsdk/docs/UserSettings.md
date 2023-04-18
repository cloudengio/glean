# UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**[]Setting**](Setting.md) | List of user setting objects. In case of multiple keys, the last one is used. | [optional] 

## Methods

### NewUserSettings

`func NewUserSettings() *UserSettings`

NewUserSettings instantiates a new UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsWithDefaults

`func NewUserSettingsWithDefaults() *UserSettings`

NewUserSettingsWithDefaults instantiates a new UserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *UserSettings) GetSettings() []Setting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserSettings) GetSettingsOk() (*[]Setting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserSettings) SetSettings(v []Setting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UserSettings) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


