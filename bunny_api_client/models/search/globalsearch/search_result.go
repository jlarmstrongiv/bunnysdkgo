package globalsearch

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SearchResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of the search result item linked object
    id *int64
    // The name of the object found
    name *string
    // The type of the search result item. Possible values: cdn, storage, dns, script, stream
    typeEscaped *GlobalSearchType
}
// NewSearchResult instantiates a new SearchResult and sets the default values.
func NewSearchResult()(*SearchResult) {
    m := &SearchResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSearchResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SearchResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SearchResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["Name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["Type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGlobalSearchType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*GlobalSearchType))
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The ID of the search result item linked object
// returns a *int64 when successful
func (m *SearchResult) GetId()(*int64) {
    return m.id
}
// GetName gets the Name property value. The name of the object found
// returns a *string when successful
func (m *SearchResult) GetName()(*string) {
    return m.name
}
// GetTypeEscaped gets the Type property value. The type of the search result item. Possible values: cdn, storage, dns, script, stream
// returns a *GlobalSearchType when successful
func (m *SearchResult) GetTypeEscaped()(*GlobalSearchType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SearchResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("Id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("Type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the Id property value. The ID of the search result item linked object
func (m *SearchResult) SetId(value *int64)() {
    m.id = value
}
// SetName sets the Name property value. The name of the object found
func (m *SearchResult) SetName(value *string)() {
    m.name = value
}
// SetTypeEscaped sets the Type property value. The type of the search result item. Possible values: cdn, storage, dns, script, stream
func (m *SearchResult) SetTypeEscaped(value *GlobalSearchType)() {
    m.typeEscaped = value
}
type SearchResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int64)
    GetName()(*string)
    GetTypeEscaped()(*GlobalSearchType)
    SetId(value *int64)()
    SetName(value *string)()
    SetTypeEscaped(value *GlobalSearchType)()
}
