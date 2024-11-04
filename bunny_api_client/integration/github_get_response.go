package integration

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/integrations"
)

type GithubGetResponse struct {
    // The Accounts property
    accounts []i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewGithubGetResponse instantiates a new GithubGetResponse and sets the default values.
func NewGithubGetResponse()(*GithubGetResponse) {
    m := &GithubGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGithubGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGithubGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGithubGetResponse(), nil
}
// GetAccounts gets the Accounts property value. The Accounts property
// returns a []Accountable when successful
func (m *GithubGetResponse) GetAccounts()([]i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable) {
    return m.accounts
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GithubGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GithubGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Accounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable)
                }
            }
            m.SetAccounts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GithubGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccounts()))
        for i, v := range m.GetAccounts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("Accounts", cast)
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
// SetAccounts sets the Accounts property value. The Accounts property
func (m *GithubGetResponse) SetAccounts(value []i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable)() {
    m.accounts = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GithubGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type GithubGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccounts()([]i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable)
    SetAccounts(value []i992bc58bc9656a17fd1c36e36d09357969d1fdf7d4127abf3227a2e812004195.Accountable)()
}
