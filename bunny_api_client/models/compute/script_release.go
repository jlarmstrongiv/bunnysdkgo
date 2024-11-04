package compute

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ScriptRelease struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Author property
    author *string
    // The AuthorEmail property
    authorEmail *string
    // The Code property
    code *string
    // The CommitSha property
    commitSha *string
    // The DatePublished property
    datePublished *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The DateReleased property
    dateReleased *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Deleted property
    deleted *bool
    // The Id property
    id *int64
    // The Note property
    note *string
    // The Status property
    status *float64
    // The Uuid property
    uuid *string
}
// NewScriptRelease instantiates a new ScriptRelease and sets the default values.
func NewScriptRelease()(*ScriptRelease) {
    m := &ScriptRelease{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScriptReleaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScriptReleaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScriptRelease(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ScriptRelease) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the Author property value. The Author property
// returns a *string when successful
func (m *ScriptRelease) GetAuthor()(*string) {
    return m.author
}
// GetAuthorEmail gets the AuthorEmail property value. The AuthorEmail property
// returns a *string when successful
func (m *ScriptRelease) GetAuthorEmail()(*string) {
    return m.authorEmail
}
// GetCode gets the Code property value. The Code property
// returns a *string when successful
func (m *ScriptRelease) GetCode()(*string) {
    return m.code
}
// GetCommitSha gets the CommitSha property value. The CommitSha property
// returns a *string when successful
func (m *ScriptRelease) GetCommitSha()(*string) {
    return m.commitSha
}
// GetDatePublished gets the DatePublished property value. The DatePublished property
// returns a *Time when successful
func (m *ScriptRelease) GetDatePublished()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.datePublished
}
// GetDateReleased gets the DateReleased property value. The DateReleased property
// returns a *Time when successful
func (m *ScriptRelease) GetDateReleased()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateReleased
}
// GetDeleted gets the Deleted property value. The Deleted property
// returns a *bool when successful
func (m *ScriptRelease) GetDeleted()(*bool) {
    return m.deleted
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ScriptRelease) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val)
        }
        return nil
    }
    res["AuthorEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorEmail(val)
        }
        return nil
    }
    res["Code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["CommitSha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitSha(val)
        }
        return nil
    }
    res["DatePublished"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatePublished(val)
        }
        return nil
    }
    res["DateReleased"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateReleased(val)
        }
        return nil
    }
    res["Deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val)
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
    res["Note"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNote(val)
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
    res["Uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *ScriptRelease) GetId()(*int64) {
    return m.id
}
// GetNote gets the Note property value. The Note property
// returns a *string when successful
func (m *ScriptRelease) GetNote()(*string) {
    return m.note
}
// GetStatus gets the Status property value. The Status property
// returns a *float64 when successful
func (m *ScriptRelease) GetStatus()(*float64) {
    return m.status
}
// GetUuid gets the Uuid property value. The Uuid property
// returns a *string when successful
func (m *ScriptRelease) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *ScriptRelease) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("AuthorEmail", m.GetAuthorEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CommitSha", m.GetCommitSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("DatePublished", m.GetDatePublished())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("DateReleased", m.GetDateReleased())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Deleted", m.GetDeleted())
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
        err := writer.WriteStringValue("Note", m.GetNote())
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
    {
        err := writer.WriteStringValue("Uuid", m.GetUuid())
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
func (m *ScriptRelease) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the Author property value. The Author property
func (m *ScriptRelease) SetAuthor(value *string)() {
    m.author = value
}
// SetAuthorEmail sets the AuthorEmail property value. The AuthorEmail property
func (m *ScriptRelease) SetAuthorEmail(value *string)() {
    m.authorEmail = value
}
// SetCode sets the Code property value. The Code property
func (m *ScriptRelease) SetCode(value *string)() {
    m.code = value
}
// SetCommitSha sets the CommitSha property value. The CommitSha property
func (m *ScriptRelease) SetCommitSha(value *string)() {
    m.commitSha = value
}
// SetDatePublished sets the DatePublished property value. The DatePublished property
func (m *ScriptRelease) SetDatePublished(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.datePublished = value
}
// SetDateReleased sets the DateReleased property value. The DateReleased property
func (m *ScriptRelease) SetDateReleased(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateReleased = value
}
// SetDeleted sets the Deleted property value. The Deleted property
func (m *ScriptRelease) SetDeleted(value *bool)() {
    m.deleted = value
}
// SetId sets the Id property value. The Id property
func (m *ScriptRelease) SetId(value *int64)() {
    m.id = value
}
// SetNote sets the Note property value. The Note property
func (m *ScriptRelease) SetNote(value *string)() {
    m.note = value
}
// SetStatus sets the Status property value. The Status property
func (m *ScriptRelease) SetStatus(value *float64)() {
    m.status = value
}
// SetUuid sets the Uuid property value. The Uuid property
func (m *ScriptRelease) SetUuid(value *string)() {
    m.uuid = value
}
type ScriptReleaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(*string)
    GetAuthorEmail()(*string)
    GetCode()(*string)
    GetCommitSha()(*string)
    GetDatePublished()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateReleased()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeleted()(*bool)
    GetId()(*int64)
    GetNote()(*string)
    GetStatus()(*float64)
    GetUuid()(*string)
    SetAuthor(value *string)()
    SetAuthorEmail(value *string)()
    SetCode(value *string)()
    SetCommitSha(value *string)()
    SetDatePublished(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateReleased(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeleted(value *bool)()
    SetId(value *int64)()
    SetNote(value *string)()
    SetStatus(value *float64)()
    SetUuid(value *string)()
}
