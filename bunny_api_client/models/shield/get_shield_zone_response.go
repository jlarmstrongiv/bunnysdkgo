package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GetShieldZoneResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The data property
    data ShieldZoneResponseable
}
// NewGetShieldZoneResponse instantiates a new GetShieldZoneResponse and sets the default values.
func NewGetShieldZoneResponse()(*GetShieldZoneResponse) {
    m := &GetShieldZoneResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetShieldZoneResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetShieldZoneResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetShieldZoneResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetShieldZoneResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. The data property
// returns a ShieldZoneResponseable when successful
func (m *GetShieldZoneResponse) GetData()(ShieldZoneResponseable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetShieldZoneResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShieldZoneResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(ShieldZoneResponseable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GetShieldZoneResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
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
func (m *GetShieldZoneResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *GetShieldZoneResponse) SetData(value ShieldZoneResponseable)() {
    m.data = value
}
type GetShieldZoneResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetData()(ShieldZoneResponseable)
    SetData(value ShieldZoneResponseable)()
}
