package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemSetforcesslSetForceSSLPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Set to true to force SSL on the given pull zone hostname
    forceSSL *bool
    // The hostname that will be updated
    hostname *string
}
// NewItemSetforcesslSetForceSSLPostRequestBody instantiates a new ItemSetforcesslSetForceSSLPostRequestBody and sets the default values.
func NewItemSetforcesslSetForceSSLPostRequestBody()(*ItemSetforcesslSetForceSSLPostRequestBody) {
    m := &ItemSetforcesslSetForceSSLPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSetforcesslSetForceSSLPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSetforcesslSetForceSSLPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSetforcesslSetForceSSLPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSetforcesslSetForceSSLPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSetforcesslSetForceSSLPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ForceSSL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceSSL(val)
        }
        return nil
    }
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
// GetForceSSL gets the ForceSSL property value. Set to true to force SSL on the given pull zone hostname
// returns a *bool when successful
func (m *ItemSetforcesslSetForceSSLPostRequestBody) GetForceSSL()(*bool) {
    return m.forceSSL
}
// GetHostname gets the Hostname property value. The hostname that will be updated
// returns a *string when successful
func (m *ItemSetforcesslSetForceSSLPostRequestBody) GetHostname()(*string) {
    return m.hostname
}
// Serialize serializes information the current object
func (m *ItemSetforcesslSetForceSSLPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("ForceSSL", m.GetForceSSL())
        if err != nil {
            return err
        }
    }
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
func (m *ItemSetforcesslSetForceSSLPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetForceSSL sets the ForceSSL property value. Set to true to force SSL on the given pull zone hostname
func (m *ItemSetforcesslSetForceSSLPostRequestBody) SetForceSSL(value *bool)() {
    m.forceSSL = value
}
// SetHostname sets the Hostname property value. The hostname that will be updated
func (m *ItemSetforcesslSetForceSSLPostRequestBody) SetHostname(value *string)() {
    m.hostname = value
}
type ItemSetforcesslSetForceSSLPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetForceSSL()(*bool)
    GetHostname()(*string)
    SetForceSSL(value *bool)()
    SetHostname(value *string)()
}
