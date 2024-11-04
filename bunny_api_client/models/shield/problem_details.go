package shield

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProblemDetails struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The detail property
    detail *string
    // The instance property
    instance *string
    // The status property
    status *int32
    // The title property
    title *string
    // The type property
    typeEscaped *string
}
// NewProblemDetails instantiates a new ProblemDetails and sets the default values.
func NewProblemDetails()(*ProblemDetails) {
    m := &ProblemDetails{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProblemDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProblemDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProblemDetails(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *ProblemDetails) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProblemDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetail gets the detail property value. The detail property
// returns a *string when successful
func (m *ProblemDetails) GetDetail()(*string) {
    return m.detail
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProblemDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val)
        }
        return nil
    }
    res["instance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstance(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetInstance gets the instance property value. The instance property
// returns a *string when successful
func (m *ProblemDetails) GetInstance()(*string) {
    return m.instance
}
// GetStatus gets the status property value. The status property
// returns a *int32 when successful
func (m *ProblemDetails) GetStatus()(*int32) {
    return m.status
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *ProblemDetails) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *ProblemDetails) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ProblemDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instance", m.GetInstance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("status", m.GetStatus())
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
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *ProblemDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetail sets the detail property value. The detail property
func (m *ProblemDetails) SetDetail(value *string)() {
    m.detail = value
}
// SetInstance sets the instance property value. The instance property
func (m *ProblemDetails) SetInstance(value *string)() {
    m.instance = value
}
// SetStatus sets the status property value. The status property
func (m *ProblemDetails) SetStatus(value *int32)() {
    m.status = value
}
// SetTitle sets the title property value. The title property
func (m *ProblemDetails) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *ProblemDetails) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type ProblemDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetail()(*string)
    GetInstance()(*string)
    GetStatus()(*int32)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    SetDetail(value *string)()
    SetInstance(value *string)()
    SetStatus(value *int32)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
}
