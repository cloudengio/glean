/*
Glean Client API

# Introduction These are the public APIs to enable implementing a custom client interface to the Glean system.  # Usage guidelines This API is evolving fast. Glean will provide advance notice of any planned backwards incompatible changes along with a 6-month sunset period for anything that requires developers to adopt the new versions. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleanclientsdk

import (
	"encoding/json"
	"time"
)

// checks if the DocumentMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentMetadata{}

// DocumentMetadata struct for DocumentMetadata
type DocumentMetadata struct {
	Datasource *string `json:"datasource,omitempty"`
	// The datasource instance from which the document was extracted.
	DatasourceInstance *string `json:"datasourceInstance,omitempty"`
	// The type of the result. Interpretation is specific to each datasource. (e.g. for Jira issues, this is the issue type such as Bug or Feature Request).
	ObjectType *string `json:"objectType,omitempty"`
	// The name of the container (higher level parent, not direct parent) of the result. Interpretation is specific to each datasource (e.g. Channels for Slack, Project for Jira). cf. parentId
	Container *string `json:"container,omitempty"`
	// The id of the direct parent of the result. Interpretation is specific to each datasource (e.g. parent issue for Jira). cf. container
	ParentId *string `json:"parentId,omitempty"`
	MimeType *string `json:"mimeType,omitempty"`
	// The index-wide unique identifier.
	DocumentId *string `json:"documentId,omitempty"`
	// Hash of documentId.
	DocumentIdHash *string `json:"documentIdHash,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	Author *Person `json:"author,omitempty"`
	Owner *Person `json:"owner,omitempty"`
	Visibility *DocumentVisibility `json:"visibility,omitempty"`
	// A list of components this result is associated with. Interpretation is specific to each datasource. (e.g. for Jira issues, these are [components](https://confluence.atlassian.com/jirasoftwarecloud/organizing-work-with-components-764478279.html).)
	Components []string `json:"components,omitempty"`
	// The status or disposition of the result. Interpretation is specific to each datasource. (e.g. for Jira issues, this is the issue status such as Done, In Progress or Will Not Fix).
	Status *string `json:"status,omitempty"`
	// The status category of the result. Meant to be more general than status. Interpretation is specific to each datasource.
	StatusCategory *string `json:"statusCategory,omitempty"`
	// A list of stars associated with this result.  \"Pin\" is an older name.
	Pins []PinDocument `json:"pins,omitempty"`
	// The document priority. Interpretation is datasource specific.
	Priority *string `json:"priority,omitempty"`
	AssignedTo *Person `json:"assignedTo,omitempty"`
	UpdatedBy *Person `json:"updatedBy,omitempty"`
	// A list of tags for the document. Interpretation is datasource specific.
	Labels []string `json:"labels,omitempty"`
	// A list of collections that the document belongs to.
	Collections []Collection `json:"collections,omitempty"`
	// The user-visible datasource specific id (e.g. Salesforce case number for example, GitHub PR number).
	DatasourceId *string `json:"datasourceId,omitempty"`
	Interactions *DocumentInteractions `json:"interactions,omitempty"`
	Verification *Verification `json:"verification,omitempty"`
	ViewerInfo *ViewerInfo `json:"viewerInfo,omitempty"`
	Permissions *ObjectPermissions `json:"permissions,omitempty"`
	VisitCount *CountInfo `json:"visitCount,omitempty"`
	// A list of shortcuts of which destination url is for the document.
	Shortcuts []Shortcut `json:"shortcuts,omitempty"`
	// For file datasources like onedrive/github etc this has the path to the file
	Path *string `json:"path,omitempty"`
	// Custom fields specific to individual datasources
	CustomData *map[string]CustomDataValue `json:"customData,omitempty"`
	// The document's document_category(.proto).
	DocumentCategory *string `json:"documentCategory,omitempty"`
	ContactPerson *Person `json:"contactPerson,omitempty"`
	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
	IndexStatus *IndexStatus `json:"indexStatus,omitempty"`
}

// NewDocumentMetadata instantiates a new DocumentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentMetadata() *DocumentMetadata {
	this := DocumentMetadata{}
	return &this
}

// NewDocumentMetadataWithDefaults instantiates a new DocumentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentMetadataWithDefaults() *DocumentMetadata {
	this := DocumentMetadata{}
	return &this
}

// GetDatasource returns the Datasource field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDatasource() string {
	if o == nil || IsNil(o.Datasource) {
		var ret string
		return ret
	}
	return *o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDatasourceOk() (*string, bool) {
	if o == nil || IsNil(o.Datasource) {
		return nil, false
	}
	return o.Datasource, true
}

// HasDatasource returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDatasource() bool {
	if o != nil && !IsNil(o.Datasource) {
		return true
	}

	return false
}

// SetDatasource gets a reference to the given string and assigns it to the Datasource field.
func (o *DocumentMetadata) SetDatasource(v string) {
	o.Datasource = &v
}

// GetDatasourceInstance returns the DatasourceInstance field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDatasourceInstance() string {
	if o == nil || IsNil(o.DatasourceInstance) {
		var ret string
		return ret
	}
	return *o.DatasourceInstance
}

// GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDatasourceInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceInstance) {
		return nil, false
	}
	return o.DatasourceInstance, true
}

// HasDatasourceInstance returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDatasourceInstance() bool {
	if o != nil && !IsNil(o.DatasourceInstance) {
		return true
	}

	return false
}

// SetDatasourceInstance gets a reference to the given string and assigns it to the DatasourceInstance field.
func (o *DocumentMetadata) SetDatasourceInstance(v string) {
	o.DatasourceInstance = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *DocumentMetadata) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *DocumentMetadata) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *DocumentMetadata) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *DocumentMetadata) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *DocumentMetadata) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *DocumentMetadata) SetContainer(v string) {
	o.Container = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *DocumentMetadata) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *DocumentMetadata) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *DocumentMetadata) SetParentId(v string) {
	o.ParentId = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *DocumentMetadata) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *DocumentMetadata) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *DocumentMetadata) SetMimeType(v string) {
	o.MimeType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDocumentId() string {
	if o == nil || IsNil(o.DocumentId) {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *DocumentMetadata) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentIdHash returns the DocumentIdHash field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDocumentIdHash() string {
	if o == nil || IsNil(o.DocumentIdHash) {
		var ret string
		return ret
	}
	return *o.DocumentIdHash
}

// GetDocumentIdHashOk returns a tuple with the DocumentIdHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDocumentIdHashOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentIdHash) {
		return nil, false
	}
	return o.DocumentIdHash, true
}

// HasDocumentIdHash returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDocumentIdHash() bool {
	if o != nil && !IsNil(o.DocumentIdHash) {
		return true
	}

	return false
}

// SetDocumentIdHash gets a reference to the given string and assigns it to the DocumentIdHash field.
func (o *DocumentMetadata) SetDocumentIdHash(v string) {
	o.DocumentIdHash = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *DocumentMetadata) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *DocumentMetadata) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *DocumentMetadata) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *DocumentMetadata) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *DocumentMetadata) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *DocumentMetadata) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *DocumentMetadata) GetAuthor() Person {
	if o == nil || IsNil(o.Author) {
		var ret Person
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetAuthorOk() (*Person, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *DocumentMetadata) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given Person and assigns it to the Author field.
func (o *DocumentMetadata) SetAuthor(v Person) {
	o.Author = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DocumentMetadata) GetOwner() Person {
	if o == nil || IsNil(o.Owner) {
		var ret Person
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetOwnerOk() (*Person, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DocumentMetadata) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Person and assigns it to the Owner field.
func (o *DocumentMetadata) SetOwner(v Person) {
	o.Owner = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *DocumentMetadata) GetVisibility() DocumentVisibility {
	if o == nil || IsNil(o.Visibility) {
		var ret DocumentVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetVisibilityOk() (*DocumentVisibility, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *DocumentMetadata) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given DocumentVisibility and assigns it to the Visibility field.
func (o *DocumentMetadata) SetVisibility(v DocumentVisibility) {
	o.Visibility = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *DocumentMetadata) GetComponents() []string {
	if o == nil || IsNil(o.Components) {
		var ret []string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetComponentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *DocumentMetadata) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *DocumentMetadata) SetComponents(v []string) {
	o.Components = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DocumentMetadata) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DocumentMetadata) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DocumentMetadata) SetStatus(v string) {
	o.Status = &v
}

// GetStatusCategory returns the StatusCategory field value if set, zero value otherwise.
func (o *DocumentMetadata) GetStatusCategory() string {
	if o == nil || IsNil(o.StatusCategory) {
		var ret string
		return ret
	}
	return *o.StatusCategory
}

// GetStatusCategoryOk returns a tuple with the StatusCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetStatusCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCategory) {
		return nil, false
	}
	return o.StatusCategory, true
}

// HasStatusCategory returns a boolean if a field has been set.
func (o *DocumentMetadata) HasStatusCategory() bool {
	if o != nil && !IsNil(o.StatusCategory) {
		return true
	}

	return false
}

// SetStatusCategory gets a reference to the given string and assigns it to the StatusCategory field.
func (o *DocumentMetadata) SetStatusCategory(v string) {
	o.StatusCategory = &v
}

// GetPins returns the Pins field value if set, zero value otherwise.
func (o *DocumentMetadata) GetPins() []PinDocument {
	if o == nil || IsNil(o.Pins) {
		var ret []PinDocument
		return ret
	}
	return o.Pins
}

// GetPinsOk returns a tuple with the Pins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetPinsOk() ([]PinDocument, bool) {
	if o == nil || IsNil(o.Pins) {
		return nil, false
	}
	return o.Pins, true
}

// HasPins returns a boolean if a field has been set.
func (o *DocumentMetadata) HasPins() bool {
	if o != nil && !IsNil(o.Pins) {
		return true
	}

	return false
}

// SetPins gets a reference to the given []PinDocument and assigns it to the Pins field.
func (o *DocumentMetadata) SetPins(v []PinDocument) {
	o.Pins = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DocumentMetadata) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DocumentMetadata) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *DocumentMetadata) SetPriority(v string) {
	o.Priority = &v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *DocumentMetadata) GetAssignedTo() Person {
	if o == nil || IsNil(o.AssignedTo) {
		var ret Person
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetAssignedToOk() (*Person, bool) {
	if o == nil || IsNil(o.AssignedTo) {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *DocumentMetadata) HasAssignedTo() bool {
	if o != nil && !IsNil(o.AssignedTo) {
		return true
	}

	return false
}

// SetAssignedTo gets a reference to the given Person and assigns it to the AssignedTo field.
func (o *DocumentMetadata) SetAssignedTo(v Person) {
	o.AssignedTo = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DocumentMetadata) GetUpdatedBy() Person {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret Person
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetUpdatedByOk() (*Person, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DocumentMetadata) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given Person and assigns it to the UpdatedBy field.
func (o *DocumentMetadata) SetUpdatedBy(v Person) {
	o.UpdatedBy = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DocumentMetadata) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DocumentMetadata) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *DocumentMetadata) SetLabels(v []string) {
	o.Labels = v
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *DocumentMetadata) GetCollections() []Collection {
	if o == nil || IsNil(o.Collections) {
		var ret []Collection
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetCollectionsOk() ([]Collection, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *DocumentMetadata) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []Collection and assigns it to the Collections field.
func (o *DocumentMetadata) SetCollections(v []Collection) {
	o.Collections = v
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDatasourceId() string {
	if o == nil || IsNil(o.DatasourceId) {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDatasourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasourceId) {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDatasourceId() bool {
	if o != nil && !IsNil(o.DatasourceId) {
		return true
	}

	return false
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *DocumentMetadata) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

// GetInteractions returns the Interactions field value if set, zero value otherwise.
func (o *DocumentMetadata) GetInteractions() DocumentInteractions {
	if o == nil || IsNil(o.Interactions) {
		var ret DocumentInteractions
		return ret
	}
	return *o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetInteractionsOk() (*DocumentInteractions, bool) {
	if o == nil || IsNil(o.Interactions) {
		return nil, false
	}
	return o.Interactions, true
}

// HasInteractions returns a boolean if a field has been set.
func (o *DocumentMetadata) HasInteractions() bool {
	if o != nil && !IsNil(o.Interactions) {
		return true
	}

	return false
}

// SetInteractions gets a reference to the given DocumentInteractions and assigns it to the Interactions field.
func (o *DocumentMetadata) SetInteractions(v DocumentInteractions) {
	o.Interactions = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *DocumentMetadata) GetVerification() Verification {
	if o == nil || IsNil(o.Verification) {
		var ret Verification
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetVerificationOk() (*Verification, bool) {
	if o == nil || IsNil(o.Verification) {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *DocumentMetadata) HasVerification() bool {
	if o != nil && !IsNil(o.Verification) {
		return true
	}

	return false
}

// SetVerification gets a reference to the given Verification and assigns it to the Verification field.
func (o *DocumentMetadata) SetVerification(v Verification) {
	o.Verification = &v
}

// GetViewerInfo returns the ViewerInfo field value if set, zero value otherwise.
func (o *DocumentMetadata) GetViewerInfo() ViewerInfo {
	if o == nil || IsNil(o.ViewerInfo) {
		var ret ViewerInfo
		return ret
	}
	return *o.ViewerInfo
}

// GetViewerInfoOk returns a tuple with the ViewerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetViewerInfoOk() (*ViewerInfo, bool) {
	if o == nil || IsNil(o.ViewerInfo) {
		return nil, false
	}
	return o.ViewerInfo, true
}

// HasViewerInfo returns a boolean if a field has been set.
func (o *DocumentMetadata) HasViewerInfo() bool {
	if o != nil && !IsNil(o.ViewerInfo) {
		return true
	}

	return false
}

// SetViewerInfo gets a reference to the given ViewerInfo and assigns it to the ViewerInfo field.
func (o *DocumentMetadata) SetViewerInfo(v ViewerInfo) {
	o.ViewerInfo = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *DocumentMetadata) GetPermissions() ObjectPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret ObjectPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetPermissionsOk() (*ObjectPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *DocumentMetadata) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ObjectPermissions and assigns it to the Permissions field.
func (o *DocumentMetadata) SetPermissions(v ObjectPermissions) {
	o.Permissions = &v
}

// GetVisitCount returns the VisitCount field value if set, zero value otherwise.
func (o *DocumentMetadata) GetVisitCount() CountInfo {
	if o == nil || IsNil(o.VisitCount) {
		var ret CountInfo
		return ret
	}
	return *o.VisitCount
}

// GetVisitCountOk returns a tuple with the VisitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetVisitCountOk() (*CountInfo, bool) {
	if o == nil || IsNil(o.VisitCount) {
		return nil, false
	}
	return o.VisitCount, true
}

// HasVisitCount returns a boolean if a field has been set.
func (o *DocumentMetadata) HasVisitCount() bool {
	if o != nil && !IsNil(o.VisitCount) {
		return true
	}

	return false
}

// SetVisitCount gets a reference to the given CountInfo and assigns it to the VisitCount field.
func (o *DocumentMetadata) SetVisitCount(v CountInfo) {
	o.VisitCount = &v
}

// GetShortcuts returns the Shortcuts field value if set, zero value otherwise.
func (o *DocumentMetadata) GetShortcuts() []Shortcut {
	if o == nil || IsNil(o.Shortcuts) {
		var ret []Shortcut
		return ret
	}
	return o.Shortcuts
}

// GetShortcutsOk returns a tuple with the Shortcuts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetShortcutsOk() ([]Shortcut, bool) {
	if o == nil || IsNil(o.Shortcuts) {
		return nil, false
	}
	return o.Shortcuts, true
}

// HasShortcuts returns a boolean if a field has been set.
func (o *DocumentMetadata) HasShortcuts() bool {
	if o != nil && !IsNil(o.Shortcuts) {
		return true
	}

	return false
}

// SetShortcuts gets a reference to the given []Shortcut and assigns it to the Shortcuts field.
func (o *DocumentMetadata) SetShortcuts(v []Shortcut) {
	o.Shortcuts = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DocumentMetadata) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DocumentMetadata) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DocumentMetadata) SetPath(v string) {
	o.Path = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *DocumentMetadata) GetCustomData() map[string]CustomDataValue {
	if o == nil || IsNil(o.CustomData) {
		var ret map[string]CustomDataValue
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetCustomDataOk() (*map[string]CustomDataValue, bool) {
	if o == nil || IsNil(o.CustomData) {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *DocumentMetadata) HasCustomData() bool {
	if o != nil && !IsNil(o.CustomData) {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given map[string]CustomDataValue and assigns it to the CustomData field.
func (o *DocumentMetadata) SetCustomData(v map[string]CustomDataValue) {
	o.CustomData = &v
}

// GetDocumentCategory returns the DocumentCategory field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDocumentCategory() string {
	if o == nil || IsNil(o.DocumentCategory) {
		var ret string
		return ret
	}
	return *o.DocumentCategory
}

// GetDocumentCategoryOk returns a tuple with the DocumentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDocumentCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentCategory) {
		return nil, false
	}
	return o.DocumentCategory, true
}

// HasDocumentCategory returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDocumentCategory() bool {
	if o != nil && !IsNil(o.DocumentCategory) {
		return true
	}

	return false
}

// SetDocumentCategory gets a reference to the given string and assigns it to the DocumentCategory field.
func (o *DocumentMetadata) SetDocumentCategory(v string) {
	o.DocumentCategory = &v
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise.
func (o *DocumentMetadata) GetContactPerson() Person {
	if o == nil || IsNil(o.ContactPerson) {
		var ret Person
		return ret
	}
	return *o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetContactPersonOk() (*Person, bool) {
	if o == nil || IsNil(o.ContactPerson) {
		return nil, false
	}
	return o.ContactPerson, true
}

// HasContactPerson returns a boolean if a field has been set.
func (o *DocumentMetadata) HasContactPerson() bool {
	if o != nil && !IsNil(o.ContactPerson) {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given Person and assigns it to the ContactPerson field.
func (o *DocumentMetadata) SetContactPerson(v Person) {
	o.ContactPerson = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *DocumentMetadata) GetThumbnail() Thumbnail {
	if o == nil || IsNil(o.Thumbnail) {
		var ret Thumbnail
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetThumbnailOk() (*Thumbnail, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *DocumentMetadata) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given Thumbnail and assigns it to the Thumbnail field.
func (o *DocumentMetadata) SetThumbnail(v Thumbnail) {
	o.Thumbnail = &v
}

// GetIndexStatus returns the IndexStatus field value if set, zero value otherwise.
func (o *DocumentMetadata) GetIndexStatus() IndexStatus {
	if o == nil || IsNil(o.IndexStatus) {
		var ret IndexStatus
		return ret
	}
	return *o.IndexStatus
}

// GetIndexStatusOk returns a tuple with the IndexStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetIndexStatusOk() (*IndexStatus, bool) {
	if o == nil || IsNil(o.IndexStatus) {
		return nil, false
	}
	return o.IndexStatus, true
}

// HasIndexStatus returns a boolean if a field has been set.
func (o *DocumentMetadata) HasIndexStatus() bool {
	if o != nil && !IsNil(o.IndexStatus) {
		return true
	}

	return false
}

// SetIndexStatus gets a reference to the given IndexStatus and assigns it to the IndexStatus field.
func (o *DocumentMetadata) SetIndexStatus(v IndexStatus) {
	o.IndexStatus = &v
}

func (o DocumentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Datasource) {
		toSerialize["datasource"] = o.Datasource
	}
	if !IsNil(o.DatasourceInstance) {
		toSerialize["datasourceInstance"] = o.DatasourceInstance
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	if !IsNil(o.DocumentId) {
		toSerialize["documentId"] = o.DocumentId
	}
	if !IsNil(o.DocumentIdHash) {
		toSerialize["documentIdHash"] = o.DocumentIdHash
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusCategory) {
		toSerialize["statusCategory"] = o.StatusCategory
	}
	if !IsNil(o.Pins) {
		toSerialize["pins"] = o.Pins
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.AssignedTo) {
		toSerialize["assignedTo"] = o.AssignedTo
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Collections) {
		toSerialize["collections"] = o.Collections
	}
	if !IsNil(o.DatasourceId) {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if !IsNil(o.Interactions) {
		toSerialize["interactions"] = o.Interactions
	}
	if !IsNil(o.Verification) {
		toSerialize["verification"] = o.Verification
	}
	if !IsNil(o.ViewerInfo) {
		toSerialize["viewerInfo"] = o.ViewerInfo
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.VisitCount) {
		toSerialize["visitCount"] = o.VisitCount
	}
	if !IsNil(o.Shortcuts) {
		toSerialize["shortcuts"] = o.Shortcuts
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.CustomData) {
		toSerialize["customData"] = o.CustomData
	}
	if !IsNil(o.DocumentCategory) {
		toSerialize["documentCategory"] = o.DocumentCategory
	}
	if !IsNil(o.ContactPerson) {
		toSerialize["contactPerson"] = o.ContactPerson
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.IndexStatus) {
		toSerialize["indexStatus"] = o.IndexStatus
	}
	return toSerialize, nil
}

type NullableDocumentMetadata struct {
	value *DocumentMetadata
	isSet bool
}

func (v NullableDocumentMetadata) Get() *DocumentMetadata {
	return v.value
}

func (v *NullableDocumentMetadata) Set(val *DocumentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentMetadata(val *DocumentMetadata) *NullableDocumentMetadata {
	return &NullableDocumentMetadata{value: val, isSet: true}
}

func (v NullableDocumentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


