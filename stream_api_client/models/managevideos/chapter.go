package managevideos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Chapter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The end time of the chapter in seconds
    end *int32
    // The start time of the chapter in seconds
    start *int32
    // The title of the chapter
    title *string
}
// NewChapter instantiates a new Chapter and sets the default values.
func NewChapter()(*Chapter) {
    m := &Chapter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChapterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChapterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChapter(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Chapter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnd gets the end property value. The end time of the chapter in seconds
// returns a *int32 when successful
func (m *Chapter) GetEnd()(*int32) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Chapter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetStart gets the start property value. The start time of the chapter in seconds
// returns a *int32 when successful
func (m *Chapter) GetStart()(*int32) {
    return m.start
}
// GetTitle gets the title property value. The title of the chapter
// returns a *string when successful
func (m *Chapter) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *Chapter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *Chapter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnd sets the end property value. The end time of the chapter in seconds
func (m *Chapter) SetEnd(value *int32)() {
    m.end = value
}
// SetStart sets the start property value. The start time of the chapter in seconds
func (m *Chapter) SetStart(value *int32)() {
    m.start = value
}
// SetTitle sets the title property value. The title of the chapter
func (m *Chapter) SetTitle(value *string)() {
    m.title = value
}
type Chapterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnd()(*int32)
    GetStart()(*int32)
    GetTitle()(*string)
    SetEnd(value *int32)()
    SetStart(value *int32)()
    SetTitle(value *string)()
}
