package transcodingmessage

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TranscodingMessage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The issueCode property
    issueCode *float64
    // The level property
    level *float64
    // The message property
    message *string
    // The timeStamp property
    timeStamp *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The value property
    value *string
}
// NewTranscodingMessage instantiates a new TranscodingMessage and sets the default values.
func NewTranscodingMessage()(*TranscodingMessage) {
    m := &TranscodingMessage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTranscodingMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTranscodingMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTranscodingMessage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TranscodingMessage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TranscodingMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["issueCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueCode(val)
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["timeStamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeStamp(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetIssueCode gets the issueCode property value. The issueCode property
// returns a *float64 when successful
func (m *TranscodingMessage) GetIssueCode()(*float64) {
    return m.issueCode
}
// GetLevel gets the level property value. The level property
// returns a *float64 when successful
func (m *TranscodingMessage) GetLevel()(*float64) {
    return m.level
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *TranscodingMessage) GetMessage()(*string) {
    return m.message
}
// GetTimeStamp gets the timeStamp property value. The timeStamp property
// returns a *Time when successful
func (m *TranscodingMessage) GetTimeStamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timeStamp
}
// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *TranscodingMessage) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *TranscodingMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("issueCode", m.GetIssueCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("timeStamp", m.GetTimeStamp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *TranscodingMessage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIssueCode sets the issueCode property value. The issueCode property
func (m *TranscodingMessage) SetIssueCode(value *float64)() {
    m.issueCode = value
}
// SetLevel sets the level property value. The level property
func (m *TranscodingMessage) SetLevel(value *float64)() {
    m.level = value
}
// SetMessage sets the message property value. The message property
func (m *TranscodingMessage) SetMessage(value *string)() {
    m.message = value
}
// SetTimeStamp sets the timeStamp property value. The timeStamp property
func (m *TranscodingMessage) SetTimeStamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeStamp = value
}
// SetValue sets the value property value. The value property
func (m *TranscodingMessage) SetValue(value *string)() {
    m.value = value
}
type TranscodingMessageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssueCode()(*float64)
    GetLevel()(*float64)
    GetMessage()(*string)
    GetTimeStamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetValue()(*string)
    SetIssueCode(value *float64)()
    SetLevel(value *float64)()
    SetMessage(value *string)()
    SetTimeStamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetValue(value *string)()
}
