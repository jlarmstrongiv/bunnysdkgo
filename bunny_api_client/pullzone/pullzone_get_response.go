package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/pullzone"
)

type PullzoneGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CurrentPage property
    currentPage *int32
    // The HasMoreItems property
    hasMoreItems *bool
    // The Items property
    items []i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable
    // The TotalItems property
    totalItems *int32
}
// NewPullzoneGetResponse instantiates a new PullzoneGetResponse and sets the default values.
func NewPullzoneGetResponse()(*PullzoneGetResponse) {
    m := &PullzoneGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullzoneGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullzoneGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullzoneGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullzoneGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentPage gets the CurrentPage property value. The CurrentPage property
// returns a *int32 when successful
func (m *PullzoneGetResponse) GetCurrentPage()(*int32) {
    return m.currentPage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullzoneGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CurrentPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPage(val)
        }
        return nil
    }
    res["HasMoreItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMoreItems(val)
        }
        return nil
    }
    res["Items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.CreatePullZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["TotalItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalItems(val)
        }
        return nil
    }
    return res
}
// GetHasMoreItems gets the HasMoreItems property value. The HasMoreItems property
// returns a *bool when successful
func (m *PullzoneGetResponse) GetHasMoreItems()(*bool) {
    return m.hasMoreItems
}
// GetItems gets the Items property value. The Items property
// returns a []PullZoneable when successful
func (m *PullzoneGetResponse) GetItems()([]i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable) {
    return m.items
}
// GetTotalItems gets the TotalItems property value. The TotalItems property
// returns a *int32 when successful
func (m *PullzoneGetResponse) GetTotalItems()(*int32) {
    return m.totalItems
}
// Serialize serializes information the current object
func (m *PullzoneGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("CurrentPage", m.GetCurrentPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("HasMoreItems", m.GetHasMoreItems())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("Items", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("TotalItems", m.GetTotalItems())
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
func (m *PullzoneGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentPage sets the CurrentPage property value. The CurrentPage property
func (m *PullzoneGetResponse) SetCurrentPage(value *int32)() {
    m.currentPage = value
}
// SetHasMoreItems sets the HasMoreItems property value. The HasMoreItems property
func (m *PullzoneGetResponse) SetHasMoreItems(value *bool)() {
    m.hasMoreItems = value
}
// SetItems sets the Items property value. The Items property
func (m *PullzoneGetResponse) SetItems(value []i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)() {
    m.items = value
}
// SetTotalItems sets the TotalItems property value. The TotalItems property
func (m *PullzoneGetResponse) SetTotalItems(value *int32)() {
    m.totalItems = value
}
type PullzoneGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPage()(*int32)
    GetHasMoreItems()(*bool)
    GetItems()([]i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)
    GetTotalItems()(*int32)
    SetCurrentPage(value *int32)()
    SetHasMoreItems(value *bool)()
    SetItems(value []i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)()
    SetTotalItems(value *int32)()
}
