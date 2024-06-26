package getstatistics

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Statistics struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The AverageOriginResponseTime property
    averageOriginResponseTime *int32
    // The BandwidthCachedChart property
    bandwidthCachedChart Statistics_BandwidthCachedChartable
    // The BandwidthUsedChart property
    bandwidthUsedChart Statistics_BandwidthUsedChartable
    // The CacheHitRate property
    cacheHitRate *float64
    // The CacheHitRateChart property
    cacheHitRateChart Statistics_CacheHitRateChartable
    // The Error3xxChart property
    error3xxChart Statistics_Error3xxChartable
    // The Error4xxChart property
    error4xxChart Statistics_Error4xxChartable
    // The Error5xxChart property
    error5xxChart Statistics_Error5xxChartable
    // The GeoTrafficDistribution property
    geoTrafficDistribution Statistics_GeoTrafficDistributionable
    // The OriginResponseTimeChart property
    originResponseTimeChart Statistics_OriginResponseTimeChartable
    // The OriginShieldBandwidthUsedChart property
    originShieldBandwidthUsedChart Statistics_OriginShieldBandwidthUsedChartable
    // The OriginShieldInternalBandwidthUsedChart property
    originShieldInternalBandwidthUsedChart Statistics_OriginShieldInternalBandwidthUsedChartable
    // The OriginTrafficChart property
    originTrafficChart Statistics_OriginTrafficChartable
    // The PullRequestsPulledChart property
    pullRequestsPulledChart Statistics_PullRequestsPulledChartable
    // The RequestsServedChart property
    requestsServedChart Statistics_RequestsServedChartable
    // The TotalBandwidthUsed property
    totalBandwidthUsed *int64
    // The TotalOriginTraffic property
    totalOriginTraffic *int64
    // The TotalRequestsServed property
    totalRequestsServed *int64
    // The UserBalanceHistoryChart property
    userBalanceHistoryChart Statistics_UserBalanceHistoryChartable
}
// NewStatistics instantiates a new Statistics and sets the default values.
func NewStatistics()(*Statistics) {
    m := &Statistics{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatistics(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Statistics) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAverageOriginResponseTime gets the AverageOriginResponseTime property value. The AverageOriginResponseTime property
// returns a *int32 when successful
func (m *Statistics) GetAverageOriginResponseTime()(*int32) {
    return m.averageOriginResponseTime
}
// GetBandwidthCachedChart gets the BandwidthCachedChart property value. The BandwidthCachedChart property
// returns a Statistics_BandwidthCachedChartable when successful
func (m *Statistics) GetBandwidthCachedChart()(Statistics_BandwidthCachedChartable) {
    return m.bandwidthCachedChart
}
// GetBandwidthUsedChart gets the BandwidthUsedChart property value. The BandwidthUsedChart property
// returns a Statistics_BandwidthUsedChartable when successful
func (m *Statistics) GetBandwidthUsedChart()(Statistics_BandwidthUsedChartable) {
    return m.bandwidthUsedChart
}
// GetCacheHitRate gets the CacheHitRate property value. The CacheHitRate property
// returns a *float64 when successful
func (m *Statistics) GetCacheHitRate()(*float64) {
    return m.cacheHitRate
}
// GetCacheHitRateChart gets the CacheHitRateChart property value. The CacheHitRateChart property
// returns a Statistics_CacheHitRateChartable when successful
func (m *Statistics) GetCacheHitRateChart()(Statistics_CacheHitRateChartable) {
    return m.cacheHitRateChart
}
// GetError3xxChart gets the Error3xxChart property value. The Error3xxChart property
// returns a Statistics_Error3xxChartable when successful
func (m *Statistics) GetError3xxChart()(Statistics_Error3xxChartable) {
    return m.error3xxChart
}
// GetError4xxChart gets the Error4xxChart property value. The Error4xxChart property
// returns a Statistics_Error4xxChartable when successful
func (m *Statistics) GetError4xxChart()(Statistics_Error4xxChartable) {
    return m.error4xxChart
}
// GetError5xxChart gets the Error5xxChart property value. The Error5xxChart property
// returns a Statistics_Error5xxChartable when successful
func (m *Statistics) GetError5xxChart()(Statistics_Error5xxChartable) {
    return m.error5xxChart
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Statistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AverageOriginResponseTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOriginResponseTime(val)
        }
        return nil
    }
    res["BandwidthCachedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_BandwidthCachedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthCachedChart(val.(Statistics_BandwidthCachedChartable))
        }
        return nil
    }
    res["BandwidthUsedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_BandwidthUsedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthUsedChart(val.(Statistics_BandwidthUsedChartable))
        }
        return nil
    }
    res["CacheHitRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheHitRate(val)
        }
        return nil
    }
    res["CacheHitRateChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_CacheHitRateChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheHitRateChart(val.(Statistics_CacheHitRateChartable))
        }
        return nil
    }
    res["Error3xxChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_Error3xxChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError3xxChart(val.(Statistics_Error3xxChartable))
        }
        return nil
    }
    res["Error4xxChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_Error4xxChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError4xxChart(val.(Statistics_Error4xxChartable))
        }
        return nil
    }
    res["Error5xxChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_Error5xxChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError5xxChart(val.(Statistics_Error5xxChartable))
        }
        return nil
    }
    res["GeoTrafficDistribution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_GeoTrafficDistributionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeoTrafficDistribution(val.(Statistics_GeoTrafficDistributionable))
        }
        return nil
    }
    res["OriginResponseTimeChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_OriginResponseTimeChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginResponseTimeChart(val.(Statistics_OriginResponseTimeChartable))
        }
        return nil
    }
    res["OriginShieldBandwidthUsedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_OriginShieldBandwidthUsedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldBandwidthUsedChart(val.(Statistics_OriginShieldBandwidthUsedChartable))
        }
        return nil
    }
    res["OriginShieldInternalBandwidthUsedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_OriginShieldInternalBandwidthUsedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginShieldInternalBandwidthUsedChart(val.(Statistics_OriginShieldInternalBandwidthUsedChartable))
        }
        return nil
    }
    res["OriginTrafficChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_OriginTrafficChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginTrafficChart(val.(Statistics_OriginTrafficChartable))
        }
        return nil
    }
    res["PullRequestsPulledChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_PullRequestsPulledChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullRequestsPulledChart(val.(Statistics_PullRequestsPulledChartable))
        }
        return nil
    }
    res["RequestsServedChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_RequestsServedChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestsServedChart(val.(Statistics_RequestsServedChartable))
        }
        return nil
    }
    res["TotalBandwidthUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBandwidthUsed(val)
        }
        return nil
    }
    res["TotalOriginTraffic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalOriginTraffic(val)
        }
        return nil
    }
    res["TotalRequestsServed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRequestsServed(val)
        }
        return nil
    }
    res["UserBalanceHistoryChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStatistics_UserBalanceHistoryChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserBalanceHistoryChart(val.(Statistics_UserBalanceHistoryChartable))
        }
        return nil
    }
    return res
}
// GetGeoTrafficDistribution gets the GeoTrafficDistribution property value. The GeoTrafficDistribution property
// returns a Statistics_GeoTrafficDistributionable when successful
func (m *Statistics) GetGeoTrafficDistribution()(Statistics_GeoTrafficDistributionable) {
    return m.geoTrafficDistribution
}
// GetOriginResponseTimeChart gets the OriginResponseTimeChart property value. The OriginResponseTimeChart property
// returns a Statistics_OriginResponseTimeChartable when successful
func (m *Statistics) GetOriginResponseTimeChart()(Statistics_OriginResponseTimeChartable) {
    return m.originResponseTimeChart
}
// GetOriginShieldBandwidthUsedChart gets the OriginShieldBandwidthUsedChart property value. The OriginShieldBandwidthUsedChart property
// returns a Statistics_OriginShieldBandwidthUsedChartable when successful
func (m *Statistics) GetOriginShieldBandwidthUsedChart()(Statistics_OriginShieldBandwidthUsedChartable) {
    return m.originShieldBandwidthUsedChart
}
// GetOriginShieldInternalBandwidthUsedChart gets the OriginShieldInternalBandwidthUsedChart property value. The OriginShieldInternalBandwidthUsedChart property
// returns a Statistics_OriginShieldInternalBandwidthUsedChartable when successful
func (m *Statistics) GetOriginShieldInternalBandwidthUsedChart()(Statistics_OriginShieldInternalBandwidthUsedChartable) {
    return m.originShieldInternalBandwidthUsedChart
}
// GetOriginTrafficChart gets the OriginTrafficChart property value. The OriginTrafficChart property
// returns a Statistics_OriginTrafficChartable when successful
func (m *Statistics) GetOriginTrafficChart()(Statistics_OriginTrafficChartable) {
    return m.originTrafficChart
}
// GetPullRequestsPulledChart gets the PullRequestsPulledChart property value. The PullRequestsPulledChart property
// returns a Statistics_PullRequestsPulledChartable when successful
func (m *Statistics) GetPullRequestsPulledChart()(Statistics_PullRequestsPulledChartable) {
    return m.pullRequestsPulledChart
}
// GetRequestsServedChart gets the RequestsServedChart property value. The RequestsServedChart property
// returns a Statistics_RequestsServedChartable when successful
func (m *Statistics) GetRequestsServedChart()(Statistics_RequestsServedChartable) {
    return m.requestsServedChart
}
// GetTotalBandwidthUsed gets the TotalBandwidthUsed property value. The TotalBandwidthUsed property
// returns a *int64 when successful
func (m *Statistics) GetTotalBandwidthUsed()(*int64) {
    return m.totalBandwidthUsed
}
// GetTotalOriginTraffic gets the TotalOriginTraffic property value. The TotalOriginTraffic property
// returns a *int64 when successful
func (m *Statistics) GetTotalOriginTraffic()(*int64) {
    return m.totalOriginTraffic
}
// GetTotalRequestsServed gets the TotalRequestsServed property value. The TotalRequestsServed property
// returns a *int64 when successful
func (m *Statistics) GetTotalRequestsServed()(*int64) {
    return m.totalRequestsServed
}
// GetUserBalanceHistoryChart gets the UserBalanceHistoryChart property value. The UserBalanceHistoryChart property
// returns a Statistics_UserBalanceHistoryChartable when successful
func (m *Statistics) GetUserBalanceHistoryChart()(Statistics_UserBalanceHistoryChartable) {
    return m.userBalanceHistoryChart
}
// Serialize serializes information the current object
func (m *Statistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("AverageOriginResponseTime", m.GetAverageOriginResponseTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("BandwidthCachedChart", m.GetBandwidthCachedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("BandwidthUsedChart", m.GetBandwidthUsedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("CacheHitRate", m.GetCacheHitRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("CacheHitRateChart", m.GetCacheHitRateChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("Error3xxChart", m.GetError3xxChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("Error4xxChart", m.GetError4xxChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("Error5xxChart", m.GetError5xxChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("GeoTrafficDistribution", m.GetGeoTrafficDistribution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("OriginResponseTimeChart", m.GetOriginResponseTimeChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("OriginShieldBandwidthUsedChart", m.GetOriginShieldBandwidthUsedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("OriginShieldInternalBandwidthUsedChart", m.GetOriginShieldInternalBandwidthUsedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("OriginTrafficChart", m.GetOriginTrafficChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("PullRequestsPulledChart", m.GetPullRequestsPulledChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("RequestsServedChart", m.GetRequestsServedChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalBandwidthUsed", m.GetTotalBandwidthUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalOriginTraffic", m.GetTotalOriginTraffic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("TotalRequestsServed", m.GetTotalRequestsServed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("UserBalanceHistoryChart", m.GetUserBalanceHistoryChart())
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
func (m *Statistics) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAverageOriginResponseTime sets the AverageOriginResponseTime property value. The AverageOriginResponseTime property
func (m *Statistics) SetAverageOriginResponseTime(value *int32)() {
    m.averageOriginResponseTime = value
}
// SetBandwidthCachedChart sets the BandwidthCachedChart property value. The BandwidthCachedChart property
func (m *Statistics) SetBandwidthCachedChart(value Statistics_BandwidthCachedChartable)() {
    m.bandwidthCachedChart = value
}
// SetBandwidthUsedChart sets the BandwidthUsedChart property value. The BandwidthUsedChart property
func (m *Statistics) SetBandwidthUsedChart(value Statistics_BandwidthUsedChartable)() {
    m.bandwidthUsedChart = value
}
// SetCacheHitRate sets the CacheHitRate property value. The CacheHitRate property
func (m *Statistics) SetCacheHitRate(value *float64)() {
    m.cacheHitRate = value
}
// SetCacheHitRateChart sets the CacheHitRateChart property value. The CacheHitRateChart property
func (m *Statistics) SetCacheHitRateChart(value Statistics_CacheHitRateChartable)() {
    m.cacheHitRateChart = value
}
// SetError3xxChart sets the Error3xxChart property value. The Error3xxChart property
func (m *Statistics) SetError3xxChart(value Statistics_Error3xxChartable)() {
    m.error3xxChart = value
}
// SetError4xxChart sets the Error4xxChart property value. The Error4xxChart property
func (m *Statistics) SetError4xxChart(value Statistics_Error4xxChartable)() {
    m.error4xxChart = value
}
// SetError5xxChart sets the Error5xxChart property value. The Error5xxChart property
func (m *Statistics) SetError5xxChart(value Statistics_Error5xxChartable)() {
    m.error5xxChart = value
}
// SetGeoTrafficDistribution sets the GeoTrafficDistribution property value. The GeoTrafficDistribution property
func (m *Statistics) SetGeoTrafficDistribution(value Statistics_GeoTrafficDistributionable)() {
    m.geoTrafficDistribution = value
}
// SetOriginResponseTimeChart sets the OriginResponseTimeChart property value. The OriginResponseTimeChart property
func (m *Statistics) SetOriginResponseTimeChart(value Statistics_OriginResponseTimeChartable)() {
    m.originResponseTimeChart = value
}
// SetOriginShieldBandwidthUsedChart sets the OriginShieldBandwidthUsedChart property value. The OriginShieldBandwidthUsedChart property
func (m *Statistics) SetOriginShieldBandwidthUsedChart(value Statistics_OriginShieldBandwidthUsedChartable)() {
    m.originShieldBandwidthUsedChart = value
}
// SetOriginShieldInternalBandwidthUsedChart sets the OriginShieldInternalBandwidthUsedChart property value. The OriginShieldInternalBandwidthUsedChart property
func (m *Statistics) SetOriginShieldInternalBandwidthUsedChart(value Statistics_OriginShieldInternalBandwidthUsedChartable)() {
    m.originShieldInternalBandwidthUsedChart = value
}
// SetOriginTrafficChart sets the OriginTrafficChart property value. The OriginTrafficChart property
func (m *Statistics) SetOriginTrafficChart(value Statistics_OriginTrafficChartable)() {
    m.originTrafficChart = value
}
// SetPullRequestsPulledChart sets the PullRequestsPulledChart property value. The PullRequestsPulledChart property
func (m *Statistics) SetPullRequestsPulledChart(value Statistics_PullRequestsPulledChartable)() {
    m.pullRequestsPulledChart = value
}
// SetRequestsServedChart sets the RequestsServedChart property value. The RequestsServedChart property
func (m *Statistics) SetRequestsServedChart(value Statistics_RequestsServedChartable)() {
    m.requestsServedChart = value
}
// SetTotalBandwidthUsed sets the TotalBandwidthUsed property value. The TotalBandwidthUsed property
func (m *Statistics) SetTotalBandwidthUsed(value *int64)() {
    m.totalBandwidthUsed = value
}
// SetTotalOriginTraffic sets the TotalOriginTraffic property value. The TotalOriginTraffic property
func (m *Statistics) SetTotalOriginTraffic(value *int64)() {
    m.totalOriginTraffic = value
}
// SetTotalRequestsServed sets the TotalRequestsServed property value. The TotalRequestsServed property
func (m *Statistics) SetTotalRequestsServed(value *int64)() {
    m.totalRequestsServed = value
}
// SetUserBalanceHistoryChart sets the UserBalanceHistoryChart property value. The UserBalanceHistoryChart property
func (m *Statistics) SetUserBalanceHistoryChart(value Statistics_UserBalanceHistoryChartable)() {
    m.userBalanceHistoryChart = value
}
type Statisticsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageOriginResponseTime()(*int32)
    GetBandwidthCachedChart()(Statistics_BandwidthCachedChartable)
    GetBandwidthUsedChart()(Statistics_BandwidthUsedChartable)
    GetCacheHitRate()(*float64)
    GetCacheHitRateChart()(Statistics_CacheHitRateChartable)
    GetError3xxChart()(Statistics_Error3xxChartable)
    GetError4xxChart()(Statistics_Error4xxChartable)
    GetError5xxChart()(Statistics_Error5xxChartable)
    GetGeoTrafficDistribution()(Statistics_GeoTrafficDistributionable)
    GetOriginResponseTimeChart()(Statistics_OriginResponseTimeChartable)
    GetOriginShieldBandwidthUsedChart()(Statistics_OriginShieldBandwidthUsedChartable)
    GetOriginShieldInternalBandwidthUsedChart()(Statistics_OriginShieldInternalBandwidthUsedChartable)
    GetOriginTrafficChart()(Statistics_OriginTrafficChartable)
    GetPullRequestsPulledChart()(Statistics_PullRequestsPulledChartable)
    GetRequestsServedChart()(Statistics_RequestsServedChartable)
    GetTotalBandwidthUsed()(*int64)
    GetTotalOriginTraffic()(*int64)
    GetTotalRequestsServed()(*int64)
    GetUserBalanceHistoryChart()(Statistics_UserBalanceHistoryChartable)
    SetAverageOriginResponseTime(value *int32)()
    SetBandwidthCachedChart(value Statistics_BandwidthCachedChartable)()
    SetBandwidthUsedChart(value Statistics_BandwidthUsedChartable)()
    SetCacheHitRate(value *float64)()
    SetCacheHitRateChart(value Statistics_CacheHitRateChartable)()
    SetError3xxChart(value Statistics_Error3xxChartable)()
    SetError4xxChart(value Statistics_Error4xxChartable)()
    SetError5xxChart(value Statistics_Error5xxChartable)()
    SetGeoTrafficDistribution(value Statistics_GeoTrafficDistributionable)()
    SetOriginResponseTimeChart(value Statistics_OriginResponseTimeChartable)()
    SetOriginShieldBandwidthUsedChart(value Statistics_OriginShieldBandwidthUsedChartable)()
    SetOriginShieldInternalBandwidthUsedChart(value Statistics_OriginShieldInternalBandwidthUsedChartable)()
    SetOriginTrafficChart(value Statistics_OriginTrafficChartable)()
    SetPullRequestsPulledChart(value Statistics_PullRequestsPulledChartable)()
    SetRequestsServedChart(value Statistics_RequestsServedChartable)()
    SetTotalBandwidthUsed(value *int64)()
    SetTotalOriginTraffic(value *int64)()
    SetTotalRequestsServed(value *int64)()
    SetUserBalanceHistoryChart(value Statistics_UserBalanceHistoryChartable)()
}
