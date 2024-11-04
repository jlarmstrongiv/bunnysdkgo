package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WafMappedEnumList struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enumName property
    enumName *string
    // The enumValues property
    enumValues []WafMappedEnumable
}
// NewWafMappedEnumList instantiates a new WafMappedEnumList and sets the default values.
func NewWafMappedEnumList()(*WafMappedEnumList) {
    m := &WafMappedEnumList{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWafMappedEnumListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWafMappedEnumListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWafMappedEnumList(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WafMappedEnumList) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnumName gets the enumName property value. The enumName property
// returns a *string when successful
func (m *WafMappedEnumList) GetEnumName()(*string) {
    return m.enumName
}
// GetEnumValues gets the enumValues property value. The enumValues property
// returns a []WafMappedEnumable when successful
func (m *WafMappedEnumList) GetEnumValues()([]WafMappedEnumable) {
    return m.enumValues
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WafMappedEnumList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enumName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnumName(val)
        }
        return nil
    }
    res["enumValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWafMappedEnumFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WafMappedEnumable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WafMappedEnumable)
                }
            }
            m.SetEnumValues(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WafMappedEnumList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("enumName", m.GetEnumName())
        if err != nil {
            return err
        }
    }
    if m.GetEnumValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnumValues()))
        for i, v := range m.GetEnumValues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("enumValues", cast)
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
func (m *WafMappedEnumList) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnumName sets the enumName property value. The enumName property
func (m *WafMappedEnumList) SetEnumName(value *string)() {
    m.enumName = value
}
// SetEnumValues sets the enumValues property value. The enumValues property
func (m *WafMappedEnumList) SetEnumValues(value []WafMappedEnumable)() {
    m.enumValues = value
}
type WafMappedEnumListable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnumName()(*string)
    GetEnumValues()([]WafMappedEnumable)
    SetEnumName(value *string)()
    SetEnumValues(value []WafMappedEnumable)()
}
