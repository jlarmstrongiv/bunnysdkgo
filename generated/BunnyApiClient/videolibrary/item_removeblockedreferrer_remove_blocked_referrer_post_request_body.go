package videolibrary

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hostname that will be removed as a blocked referer
    hostname *string
}
// NewItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody instantiates a new ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody and sets the default values.
func NewItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody()(*ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) {
    m := &ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the Hostname property value. The hostname that will be removed as a blocked referer
// returns a *string when successful
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) GetHostname()(*string) {
    return m.hostname
}
// Serialize serializes information the current object
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Hostname", m.GetHostname())
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
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostname sets the Hostname property value. The hostname that will be removed as a blocked referer
func (m *ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBody) SetHostname(value *string)() {
    m.hostname = value
}
type ItemRemoveblockedreferrerRemoveBlockedReferrerPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHostname()(*string)
    SetHostname(value *string)()
}
