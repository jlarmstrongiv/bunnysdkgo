package library

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28 "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/models/managevideos"
)

type ItemVideosGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The currentPage property
    currentPage *int32
    // The items property
    items []i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable
    // The itemsPerPage property
    itemsPerPage *int32
    // The totalItems property
    totalItems *int32
}
// NewItemVideosGetResponse instantiates a new ItemVideosGetResponse and sets the default values.
func NewItemVideosGetResponse()(*ItemVideosGetResponse) {
    m := &ItemVideosGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemVideosGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemVideosGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemVideosGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemVideosGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentPage gets the currentPage property value. The currentPage property
// returns a *int32 when successful
func (m *ItemVideosGetResponse) GetCurrentPage()(*int32) {
    return m.currentPage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemVideosGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPage(val)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.CreateVideoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    res["itemsPerPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemsPerPage(val)
        }
        return nil
    }
    res["totalItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetItems gets the items property value. The items property
// returns a []Videoable when successful
func (m *ItemVideosGetResponse) GetItems()([]i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable) {
    return m.items
}
// GetItemsPerPage gets the itemsPerPage property value. The itemsPerPage property
// returns a *int32 when successful
func (m *ItemVideosGetResponse) GetItemsPerPage()(*int32) {
    return m.itemsPerPage
}
// GetTotalItems gets the totalItems property value. The totalItems property
// returns a *int32 when successful
func (m *ItemVideosGetResponse) GetTotalItems()(*int32) {
    return m.totalItems
}
// Serialize serializes information the current object
func (m *ItemVideosGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("currentPage", m.GetCurrentPage())
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
        err := writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("itemsPerPage", m.GetItemsPerPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalItems", m.GetTotalItems())
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
func (m *ItemVideosGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentPage sets the currentPage property value. The currentPage property
func (m *ItemVideosGetResponse) SetCurrentPage(value *int32)() {
    m.currentPage = value
}
// SetItems sets the items property value. The items property
func (m *ItemVideosGetResponse) SetItems(value []i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable)() {
    m.items = value
}
// SetItemsPerPage sets the itemsPerPage property value. The itemsPerPage property
func (m *ItemVideosGetResponse) SetItemsPerPage(value *int32)() {
    m.itemsPerPage = value
}
// SetTotalItems sets the totalItems property value. The totalItems property
func (m *ItemVideosGetResponse) SetTotalItems(value *int32)() {
    m.totalItems = value
}
type ItemVideosGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPage()(*int32)
    GetItems()([]i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable)
    GetItemsPerPage()(*int32)
    GetTotalItems()(*int32)
    SetCurrentPage(value *int32)()
    SetItems(value []i804737ef689cf33cc1996aaa50bf60df828fa9ff20878b7ccb1696459a4f4e28.Videoable)()
    SetItemsPerPage(value *int32)()
    SetTotalItems(value *int32)()
}
