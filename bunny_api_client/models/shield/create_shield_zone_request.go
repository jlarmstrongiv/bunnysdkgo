package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateShieldZoneRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The pullZoneId property
    pullZoneId *int64
    // The shieldZone property
    shieldZone ShieldZoneRequestable
}
// NewCreateShieldZoneRequest instantiates a new CreateShieldZoneRequest and sets the default values.
func NewCreateShieldZoneRequest()(*CreateShieldZoneRequest) {
    m := &CreateShieldZoneRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateShieldZoneRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateShieldZoneRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateShieldZoneRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateShieldZoneRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateShieldZoneRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneId(val)
        }
        return nil
    }
    res["shieldZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShieldZoneRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldZone(val.(ShieldZoneRequestable))
        }
        return nil
    }
    return res
}
// GetPullZoneId gets the pullZoneId property value. The pullZoneId property
// returns a *int64 when successful
func (m *CreateShieldZoneRequest) GetPullZoneId()(*int64) {
    return m.pullZoneId
}
// GetShieldZone gets the shieldZone property value. The shieldZone property
// returns a ShieldZoneRequestable when successful
func (m *CreateShieldZoneRequest) GetShieldZone()(ShieldZoneRequestable) {
    return m.shieldZone
}
// Serialize serializes information the current object
func (m *CreateShieldZoneRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("pullZoneId", m.GetPullZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("shieldZone", m.GetShieldZone())
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
func (m *CreateShieldZoneRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPullZoneId sets the pullZoneId property value. The pullZoneId property
func (m *CreateShieldZoneRequest) SetPullZoneId(value *int64)() {
    m.pullZoneId = value
}
// SetShieldZone sets the shieldZone property value. The shieldZone property
func (m *CreateShieldZoneRequest) SetShieldZone(value ShieldZoneRequestable)() {
    m.shieldZone = value
}
type CreateShieldZoneRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPullZoneId()(*int64)
    GetShieldZone()(ShieldZoneRequestable)
    SetPullZoneId(value *int64)()
    SetShieldZone(value ShieldZoneRequestable)()
}
