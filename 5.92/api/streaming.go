package api // import "github.com/akss1/vksdk/5.92/api"

// StreamingGetServerURLResponse struct
type StreamingGetServerURLResponse struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
}

// StreamingGetServerURL allows to receive data for the connection to Streaming API.
//
// https://vk.com/dev/streaming.getServerUrl
func (vk *VK) StreamingGetServerURL(params map[string]string) (response StreamingGetServerURLResponse, vkErr Error) {
	vk.RequestUnmarshal("streaming.getServerUrl", params, &response, &vkErr)
	return
}

// StreamingGetSettingsResponse struct
type StreamingGetSettingsResponse struct {
	MonthlyLimit string `json:"monthly_limit"`
}

// StreamingGetSettings allows to receive monthly tier for Streaming API.
//
// https://vk.com/dev/streaming.getSettings
func (vk *VK) StreamingGetSettings(params map[string]string) (response StreamingGetSettingsResponse, vkErr Error) {
	vk.RequestUnmarshal("streaming.getSettings", params, &response, &vkErr)
	return
}

// StreamingGetStatsResponse struct
type StreamingGetStatsResponse []struct {
	EventType string `json:"event_type"`
	Stats     []struct {
		Timestamp int `json:"timestamp"`
		Value     int `json:"value"`
	} `json:"stats"`
}

// StreamingGetStats allows to receive statistics for prepared and received events in Streaming API.
//
// https://vk.com/dev/streaming.getStats
func (vk *VK) StreamingGetStats(params map[string]string) (response StreamingGetStatsResponse, vkErr Error) {
	vk.RequestUnmarshal("streaming.getStats", params, &response, &vkErr)
	return
}

// StreamingGetStemResponse struct
type StreamingGetStemResponse struct {
	Stem string `json:"stem"`
}

// StreamingGetStem allows to receive the stem of the word.
//
// https://vk.com/dev/streaming.getStem
func (vk *VK) StreamingGetStem(params map[string]string) (response StreamingGetStemResponse, vkErr Error) {
	vk.RequestUnmarshal("streaming.getStem", params, &response, &vkErr)
	return
}

// StreamingSetSettings allows to set monthly tier for Streaming API.
//
// https://vk.com/dev/streaming.setSettings
func (vk *VK) StreamingSetSettings(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("streaming.setSettings", params, &response, &vkErr)

	return
}
