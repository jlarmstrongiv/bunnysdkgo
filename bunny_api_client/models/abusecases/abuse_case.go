package abusecases

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AbuseCase struct {
    // The ActualUrl property
    actualUrl *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The DateCreated property
    dateCreated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DateUpdated property
    dateUpdated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Deadline property
    deadline *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Id property
    id *int64
    // The Message property
    message *string
    // The Path property
    path *string
    // The PullZoneId property
    pullZoneId *int64
    // The PullZoneName property
    pullZoneName *string
    // The Status property
    status *float64
    // The Urls property
    urls []Urlable
}
// NewAbuseCase instantiates a new AbuseCase and sets the default values.
func NewAbuseCase()(*AbuseCase) {
    m := &AbuseCase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAbuseCaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAbuseCaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAbuseCase(), nil
}
// GetActualUrl gets the ActualUrl property value. The ActualUrl property
// returns a *string when successful
func (m *AbuseCase) GetActualUrl()(*string) {
    return m.actualUrl
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AbuseCase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDateCreated gets the DateCreated property value. The DateCreated property
// returns a *Time when successful
func (m *AbuseCase) GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateCreated
}
// GetDateUpdated gets the DateUpdated property value. The DateUpdated property
// returns a *Time when successful
func (m *AbuseCase) GetDateUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateUpdated
}
// GetDeadline gets the Deadline property value. The Deadline property
// returns a *Time when successful
func (m *AbuseCase) GetDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deadline
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AbuseCase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ActualUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActualUrl(val)
        }
        return nil
    }
    res["DateCreated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateCreated(val)
        }
        return nil
    }
    res["DateUpdated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateUpdated(val)
        }
        return nil
    }
    res["Deadline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadline(val)
        }
        return nil
    }
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["Message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["Path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["PullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneId(val)
        }
        return nil
    }
    res["PullZoneName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneName(val)
        }
        return nil
    }
    res["Status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["Urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Urlable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Urlable)
                }
            }
            m.SetUrls(res)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *AbuseCase) GetId()(*int64) {
    return m.id
}
// GetMessage gets the Message property value. The Message property
// returns a *string when successful
func (m *AbuseCase) GetMessage()(*string) {
    return m.message
}
// GetPath gets the Path property value. The Path property
// returns a *string when successful
func (m *AbuseCase) GetPath()(*string) {
    return m.path
}
// GetPullZoneId gets the PullZoneId property value. The PullZoneId property
// returns a *int64 when successful
func (m *AbuseCase) GetPullZoneId()(*int64) {
    return m.pullZoneId
}
// GetPullZoneName gets the PullZoneName property value. The PullZoneName property
// returns a *string when successful
func (m *AbuseCase) GetPullZoneName()(*string) {
    return m.pullZoneName
}
// GetStatus gets the Status property value. The Status property
// returns a *float64 when successful
func (m *AbuseCase) GetStatus()(*float64) {
    return m.status
}
// GetUrls gets the Urls property value. The Urls property
// returns a []Urlable when successful
func (m *AbuseCase) GetUrls()([]Urlable) {
    return m.urls
}
// Serialize serializes information the current object
func (m *AbuseCase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ActualUrl", m.GetActualUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("DateCreated", m.GetDateCreated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("DateUpdated", m.GetDateUpdated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("Deadline", m.GetDeadline())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("Id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("PullZoneId", m.GetPullZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PullZoneName", m.GetPullZoneName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetUrls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUrls()))
        for i, v := range m.GetUrls() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("Urls", cast)
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
// SetActualUrl sets the ActualUrl property value. The ActualUrl property
func (m *AbuseCase) SetActualUrl(value *string)() {
    m.actualUrl = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AbuseCase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDateCreated sets the DateCreated property value. The DateCreated property
func (m *AbuseCase) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateCreated = value
}
// SetDateUpdated sets the DateUpdated property value. The DateUpdated property
func (m *AbuseCase) SetDateUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateUpdated = value
}
// SetDeadline sets the Deadline property value. The Deadline property
func (m *AbuseCase) SetDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deadline = value
}
// SetId sets the Id property value. The Id property
func (m *AbuseCase) SetId(value *int64)() {
    m.id = value
}
// SetMessage sets the Message property value. The Message property
func (m *AbuseCase) SetMessage(value *string)() {
    m.message = value
}
// SetPath sets the Path property value. The Path property
func (m *AbuseCase) SetPath(value *string)() {
    m.path = value
}
// SetPullZoneId sets the PullZoneId property value. The PullZoneId property
func (m *AbuseCase) SetPullZoneId(value *int64)() {
    m.pullZoneId = value
}
// SetPullZoneName sets the PullZoneName property value. The PullZoneName property
func (m *AbuseCase) SetPullZoneName(value *string)() {
    m.pullZoneName = value
}
// SetStatus sets the Status property value. The Status property
func (m *AbuseCase) SetStatus(value *float64)() {
    m.status = value
}
// SetUrls sets the Urls property value. The Urls property
func (m *AbuseCase) SetUrls(value []Urlable)() {
    m.urls = value
}
type AbuseCaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActualUrl()(*string)
    GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateUpdated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeadline()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int64)
    GetMessage()(*string)
    GetPath()(*string)
    GetPullZoneId()(*int64)
    GetPullZoneName()(*string)
    GetStatus()(*float64)
    GetUrls()([]Urlable)
    SetActualUrl(value *string)()
    SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeadline(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int64)()
    SetMessage(value *string)()
    SetPath(value *string)()
    SetPullZoneId(value *int64)()
    SetPullZoneName(value *string)()
    SetStatus(value *float64)()
    SetUrls(value []Urlable)()
}
