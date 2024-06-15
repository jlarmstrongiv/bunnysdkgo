package edgerule

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Trigger struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The trigger parameter 1. The value depends on the type of trigger.
    parameter1 *string
    // The list of pattern matches that will trigger the edge rule
    patternMatches []string
    // The PatternMatchingType property
    patternMatchingType *float64
    // The Type property
    typeEscaped *float64
}
// NewTrigger instantiates a new Trigger and sets the default values.
func NewTrigger()(*Trigger) {
    m := &Trigger{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrigger(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Trigger) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Trigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Parameter1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameter1(val)
        }
        return nil
    }
    res["PatternMatches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPatternMatches(res)
        }
        return nil
    }
    res["PatternMatchingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatternMatchingType(val)
        }
        return nil
    }
    res["Type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetParameter1 gets the Parameter1 property value. The trigger parameter 1. The value depends on the type of trigger.
// returns a *string when successful
func (m *Trigger) GetParameter1()(*string) {
    return m.parameter1
}
// GetPatternMatches gets the PatternMatches property value. The list of pattern matches that will trigger the edge rule
// returns a []string when successful
func (m *Trigger) GetPatternMatches()([]string) {
    return m.patternMatches
}
// GetPatternMatchingType gets the PatternMatchingType property value. The PatternMatchingType property
// returns a *float64 when successful
func (m *Trigger) GetPatternMatchingType()(*float64) {
    return m.patternMatchingType
}
// GetTypeEscaped gets the Type property value. The Type property
// returns a *float64 when successful
func (m *Trigger) GetTypeEscaped()(*float64) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Trigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Parameter1", m.GetParameter1())
        if err != nil {
            return err
        }
    }
    if m.GetPatternMatches() != nil {
        err := writer.WriteCollectionOfStringValues("PatternMatches", m.GetPatternMatches())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("PatternMatchingType", m.GetPatternMatchingType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Type", m.GetTypeEscaped())
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
func (m *Trigger) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameter1 sets the Parameter1 property value. The trigger parameter 1. The value depends on the type of trigger.
func (m *Trigger) SetParameter1(value *string)() {
    m.parameter1 = value
}
// SetPatternMatches sets the PatternMatches property value. The list of pattern matches that will trigger the edge rule
func (m *Trigger) SetPatternMatches(value []string)() {
    m.patternMatches = value
}
// SetPatternMatchingType sets the PatternMatchingType property value. The PatternMatchingType property
func (m *Trigger) SetPatternMatchingType(value *float64)() {
    m.patternMatchingType = value
}
// SetTypeEscaped sets the Type property value. The Type property
func (m *Trigger) SetTypeEscaped(value *float64)() {
    m.typeEscaped = value
}
type Triggerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameter1()(*string)
    GetPatternMatches()([]string)
    GetPatternMatchingType()(*float64)
    GetTypeEscaped()(*float64)
    SetParameter1(value *string)()
    SetPatternMatches(value []string)()
    SetPatternMatchingType(value *float64)()
    SetTypeEscaped(value *float64)()
}
