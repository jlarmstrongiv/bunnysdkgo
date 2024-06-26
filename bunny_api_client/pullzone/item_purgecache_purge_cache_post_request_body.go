package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemPurgecachePurgeCachePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CacheTag property
    cacheTag *string
}
// NewItemPurgecachePurgeCachePostRequestBody instantiates a new ItemPurgecachePurgeCachePostRequestBody and sets the default values.
func NewItemPurgecachePurgeCachePostRequestBody()(*ItemPurgecachePurgeCachePostRequestBody) {
    m := &ItemPurgecachePurgeCachePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPurgecachePurgeCachePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPurgecachePurgeCachePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPurgecachePurgeCachePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPurgecachePurgeCachePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCacheTag gets the CacheTag property value. The CacheTag property
// returns a *string when successful
func (m *ItemPurgecachePurgeCachePostRequestBody) GetCacheTag()(*string) {
    return m.cacheTag
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPurgecachePurgeCachePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CacheTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheTag(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemPurgecachePurgeCachePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("CacheTag", m.GetCacheTag())
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
func (m *ItemPurgecachePurgeCachePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCacheTag sets the CacheTag property value. The CacheTag property
func (m *ItemPurgecachePurgeCachePostRequestBody) SetCacheTag(value *string)() {
    m.cacheTag = value
}
type ItemPurgecachePurgeCachePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCacheTag()(*string)
    SetCacheTag(value *string)()
}
