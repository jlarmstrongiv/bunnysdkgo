package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemAddBlockedIpPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The IP that will be blocked from accessing the pull zone
    blockedIp *string
}
// NewItemAddBlockedIpPostRequestBody instantiates a new ItemAddBlockedIpPostRequestBody and sets the default values.
func NewItemAddBlockedIpPostRequestBody()(*ItemAddBlockedIpPostRequestBody) {
    m := &ItemAddBlockedIpPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAddBlockedIpPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAddBlockedIpPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAddBlockedIpPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAddBlockedIpPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlockedIp gets the BlockedIp property value. The IP that will be blocked from accessing the pull zone
// returns a *string when successful
func (m *ItemAddBlockedIpPostRequestBody) GetBlockedIp()(*string) {
    return m.blockedIp
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAddBlockedIpPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["BlockedIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockedIp(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemAddBlockedIpPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("BlockedIp", m.GetBlockedIp())
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
func (m *ItemAddBlockedIpPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlockedIp sets the BlockedIp property value. The IP that will be blocked from accessing the pull zone
func (m *ItemAddBlockedIpPostRequestBody) SetBlockedIp(value *string)() {
    m.blockedIp = value
}
type ItemAddBlockedIpPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlockedIp()(*string)
    SetBlockedIp(value *string)()
}
