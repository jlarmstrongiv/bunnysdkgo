package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WafMappedEnum struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The isPremium property
    isPremium *bool
    // The name property
    name *string
    // The value property
    value *int32
}
// NewWafMappedEnum instantiates a new WafMappedEnum and sets the default values.
func NewWafMappedEnum()(*WafMappedEnum) {
    m := &WafMappedEnum{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWafMappedEnumFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWafMappedEnumFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWafMappedEnum(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WafMappedEnum) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WafMappedEnum) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isPremium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPremium(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetIsPremium gets the isPremium property value. The isPremium property
// returns a *bool when successful
func (m *WafMappedEnum) GetIsPremium()(*bool) {
    return m.isPremium
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *WafMappedEnum) GetName()(*string) {
    return m.name
}
// GetValue gets the value property value. The value property
// returns a *int32 when successful
func (m *WafMappedEnum) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *WafMappedEnum) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isPremium", m.GetIsPremium())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *WafMappedEnum) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsPremium sets the isPremium property value. The isPremium property
func (m *WafMappedEnum) SetIsPremium(value *bool)() {
    m.isPremium = value
}
// SetName sets the name property value. The name property
func (m *WafMappedEnum) SetName(value *string)() {
    m.name = value
}
// SetValue sets the value property value. The value property
func (m *WafMappedEnum) SetValue(value *int32)() {
    m.value = value
}
type WafMappedEnumable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsPremium()(*bool)
    GetName()(*string)
    GetValue()(*int32)
    SetIsPremium(value *bool)()
    SetName(value *string)()
    SetValue(value *int32)()
}
