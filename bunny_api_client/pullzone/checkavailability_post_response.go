package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CheckavailabilityPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Available property
    available *bool
}
// NewCheckavailabilityPostResponse instantiates a new CheckavailabilityPostResponse and sets the default values.
func NewCheckavailabilityPostResponse()(*CheckavailabilityPostResponse) {
    m := &CheckavailabilityPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckavailabilityPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckavailabilityPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckavailabilityPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckavailabilityPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailable gets the Available property value. The Available property
// returns a *bool when successful
func (m *CheckavailabilityPostResponse) GetAvailable()(*bool) {
    return m.available
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckavailabilityPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Available"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailable(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CheckavailabilityPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("Available", m.GetAvailable())
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
func (m *CheckavailabilityPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailable sets the Available property value. The Available property
func (m *CheckavailabilityPostResponse) SetAvailable(value *bool)() {
    m.available = value
}
type CheckavailabilityPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailable()(*bool)
    SetAvailable(value *bool)()
}
