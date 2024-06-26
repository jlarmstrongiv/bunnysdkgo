package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemOptimizerStatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Average compression chart of the responses
    averageCompressionChart ItemOptimizerStatisticsGetResponse_AverageCompressionChartable
    // The average compression ratio of CDN responses
    averageCompressionRatio *float64
    // The average processing time of each request
    averageProcessingTime *float64
    // The AverageProcessingTimeChart property
    averageProcessingTimeChart ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable
    // The constructed chart of optimized requests
    requestsOptimizedChart ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable
    // The total number of optimized requests
    totalRequestsOptimized *float64
    // The total requests saved
    totalTrafficSaved *float64
    // The constructed chart of saved traffic
    trafficSavedChart ItemOptimizerStatisticsGetResponse_TrafficSavedChartable
}
// NewItemOptimizerStatisticsGetResponse instantiates a new ItemOptimizerStatisticsGetResponse and sets the default values.
func NewItemOptimizerStatisticsGetResponse()(*ItemOptimizerStatisticsGetResponse) {
    m := &ItemOptimizerStatisticsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOptimizerStatisticsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOptimizerStatisticsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOptimizerStatisticsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemOptimizerStatisticsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAverageCompressionChart gets the AverageCompressionChart property value. Average compression chart of the responses
// returns a ItemOptimizerStatisticsGetResponse_AverageCompressionChartable when successful
func (m *ItemOptimizerStatisticsGetResponse) GetAverageCompressionChart()(ItemOptimizerStatisticsGetResponse_AverageCompressionChartable) {
    return m.averageCompressionChart
}
// GetAverageCompressionRatio gets the AverageCompressionRatio property value. The average compression ratio of CDN responses
// returns a *float64 when successful
func (m *ItemOptimizerStatisticsGetResponse) GetAverageCompressionRatio()(*float64) {
    return m.averageCompressionRatio
}
// GetAverageProcessingTime gets the AverageProcessingTime property value. The average processing time of each request
// returns a *float64 when successful
func (m *ItemOptimizerStatisticsGetResponse) GetAverageProcessingTime()(*float64) {
    return m.averageProcessingTime
}
// GetAverageProcessingTimeChart gets the AverageProcessingTimeChart property value. The AverageProcessingTimeChart property
// returns a ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable when successful
func (m *ItemOptimizerStatisticsGetResponse) GetAverageProcessingTimeChart()(ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable) {
    return m.averageProcessingTimeChart
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemOptimizerStatisticsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AverageCompressionChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemOptimizerStatisticsGetResponse_AverageCompressionChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageCompressionChart(val.(ItemOptimizerStatisticsGetResponse_AverageCompressionChartable))
        }
        return nil
    }
    res["AverageCompressionRatio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageCompressionRatio(val)
        }
        return nil
    }
    res["AverageProcessingTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageProcessingTime(val)
        }
        return nil
    }
    res["AverageProcessingTimeChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageProcessingTimeChart(val.(ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable))
        }
        return nil
    }
    res["RequestsOptimizedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemOptimizerStatisticsGetResponse_RequestsOptimizedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestsOptimizedChart(val.(ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable))
        }
        return nil
    }
    res["TotalRequestsOptimized"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRequestsOptimized(val)
        }
        return nil
    }
    res["TotalTrafficSaved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTrafficSaved(val)
        }
        return nil
    }
    res["TrafficSavedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemOptimizerStatisticsGetResponse_TrafficSavedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficSavedChart(val.(ItemOptimizerStatisticsGetResponse_TrafficSavedChartable))
        }
        return nil
    }
    return res
}
// GetRequestsOptimizedChart gets the RequestsOptimizedChart property value. The constructed chart of optimized requests
// returns a ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable when successful
func (m *ItemOptimizerStatisticsGetResponse) GetRequestsOptimizedChart()(ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable) {
    return m.requestsOptimizedChart
}
// GetTotalRequestsOptimized gets the TotalRequestsOptimized property value. The total number of optimized requests
// returns a *float64 when successful
func (m *ItemOptimizerStatisticsGetResponse) GetTotalRequestsOptimized()(*float64) {
    return m.totalRequestsOptimized
}
// GetTotalTrafficSaved gets the TotalTrafficSaved property value. The total requests saved
// returns a *float64 when successful
func (m *ItemOptimizerStatisticsGetResponse) GetTotalTrafficSaved()(*float64) {
    return m.totalTrafficSaved
}
// GetTrafficSavedChart gets the TrafficSavedChart property value. The constructed chart of saved traffic
// returns a ItemOptimizerStatisticsGetResponse_TrafficSavedChartable when successful
func (m *ItemOptimizerStatisticsGetResponse) GetTrafficSavedChart()(ItemOptimizerStatisticsGetResponse_TrafficSavedChartable) {
    return m.trafficSavedChart
}
// Serialize serializes information the current object
func (m *ItemOptimizerStatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("AverageCompressionChart", m.GetAverageCompressionChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("AverageCompressionRatio", m.GetAverageCompressionRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("AverageProcessingTime", m.GetAverageProcessingTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("AverageProcessingTimeChart", m.GetAverageProcessingTimeChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("RequestsOptimizedChart", m.GetRequestsOptimizedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("TotalRequestsOptimized", m.GetTotalRequestsOptimized())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("TotalTrafficSaved", m.GetTotalTrafficSaved())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("TrafficSavedChart", m.GetTrafficSavedChart())
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
func (m *ItemOptimizerStatisticsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAverageCompressionChart sets the AverageCompressionChart property value. Average compression chart of the responses
func (m *ItemOptimizerStatisticsGetResponse) SetAverageCompressionChart(value ItemOptimizerStatisticsGetResponse_AverageCompressionChartable)() {
    m.averageCompressionChart = value
}
// SetAverageCompressionRatio sets the AverageCompressionRatio property value. The average compression ratio of CDN responses
func (m *ItemOptimizerStatisticsGetResponse) SetAverageCompressionRatio(value *float64)() {
    m.averageCompressionRatio = value
}
// SetAverageProcessingTime sets the AverageProcessingTime property value. The average processing time of each request
func (m *ItemOptimizerStatisticsGetResponse) SetAverageProcessingTime(value *float64)() {
    m.averageProcessingTime = value
}
// SetAverageProcessingTimeChart sets the AverageProcessingTimeChart property value. The AverageProcessingTimeChart property
func (m *ItemOptimizerStatisticsGetResponse) SetAverageProcessingTimeChart(value ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable)() {
    m.averageProcessingTimeChart = value
}
// SetRequestsOptimizedChart sets the RequestsOptimizedChart property value. The constructed chart of optimized requests
func (m *ItemOptimizerStatisticsGetResponse) SetRequestsOptimizedChart(value ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable)() {
    m.requestsOptimizedChart = value
}
// SetTotalRequestsOptimized sets the TotalRequestsOptimized property value. The total number of optimized requests
func (m *ItemOptimizerStatisticsGetResponse) SetTotalRequestsOptimized(value *float64)() {
    m.totalRequestsOptimized = value
}
// SetTotalTrafficSaved sets the TotalTrafficSaved property value. The total requests saved
func (m *ItemOptimizerStatisticsGetResponse) SetTotalTrafficSaved(value *float64)() {
    m.totalTrafficSaved = value
}
// SetTrafficSavedChart sets the TrafficSavedChart property value. The constructed chart of saved traffic
func (m *ItemOptimizerStatisticsGetResponse) SetTrafficSavedChart(value ItemOptimizerStatisticsGetResponse_TrafficSavedChartable)() {
    m.trafficSavedChart = value
}
type ItemOptimizerStatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageCompressionChart()(ItemOptimizerStatisticsGetResponse_AverageCompressionChartable)
    GetAverageCompressionRatio()(*float64)
    GetAverageProcessingTime()(*float64)
    GetAverageProcessingTimeChart()(ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable)
    GetRequestsOptimizedChart()(ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable)
    GetTotalRequestsOptimized()(*float64)
    GetTotalTrafficSaved()(*float64)
    GetTrafficSavedChart()(ItemOptimizerStatisticsGetResponse_TrafficSavedChartable)
    SetAverageCompressionChart(value ItemOptimizerStatisticsGetResponse_AverageCompressionChartable)()
    SetAverageCompressionRatio(value *float64)()
    SetAverageProcessingTime(value *float64)()
    SetAverageProcessingTimeChart(value ItemOptimizerStatisticsGetResponse_AverageProcessingTimeChartable)()
    SetRequestsOptimizedChart(value ItemOptimizerStatisticsGetResponse_RequestsOptimizedChartable)()
    SetTotalRequestsOptimized(value *float64)()
    SetTotalTrafficSaved(value *float64)()
    SetTrafficSavedChart(value ItemOptimizerStatisticsGetResponse_TrafficSavedChartable)()
}
