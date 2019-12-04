package api // import "github.com/akss1/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/akss1/vksdk/5.92/object"
)

// UtilsCheckLinkResponse struct
type UtilsCheckLinkResponse object.UtilsLinkChecked

// UtilsCheckLink checks whether a link is blocked in VK.
//
// https://vk.com/dev/utils.checkLink
func (vk *VK) UtilsCheckLink(params map[string]string) (response UtilsCheckLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("utils.checkLink", params, &response, &vkErr)
	return
}

// UtilsDeleteFromLastShortened deletes shortened link from user's list.
//
// https://vk.com/dev/utils.deleteFromLastShortened
func (vk *VK) UtilsDeleteFromLastShortened(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("utils.deleteFromLastShortened", params, &response, &vkErr)
	return
}

// UtilsGetLastShortenedLinksResponse struct
type UtilsGetLastShortenedLinksResponse struct {
	Count int                             `json:"count"`
	Items []object.UtilsLastShortenedLink `json:"items"`
}

// UtilsGetLastShortenedLinks returns a list of user's shortened links.
//
// https://vk.com/dev/utils.getLastShortenedLinks
func (vk *VK) UtilsGetLastShortenedLinks(params map[string]string) (response UtilsGetLastShortenedLinksResponse, vkErr Error) {
	vk.RequestUnmarshal("utils.getLastShortenedLinks", params, &response, &vkErr)
	return
}

// UtilsGetLinkStatsResponse struct
type UtilsGetLinkStatsResponse object.UtilsLinkStats

// UtilsGetLinkStats returns stats data for shortened link.
//
// extended=0
//
// https://vk.com/dev/utils.getLinkStats
func (vk *VK) UtilsGetLinkStats(params map[string]string) (response UtilsGetLinkStatsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("utils.getLinkStats", params, &response, &vkErr)
	return
}

// UtilsGetLinkStatsExtendedResponse struct
type UtilsGetLinkStatsExtendedResponse object.UtilsLinkStatsExtended

// UtilsGetLinkStatsExtended returns stats data for shortened link.
//
// extended=1
//
// https://vk.com/dev/utils.getLinkStats
func (vk *VK) UtilsGetLinkStatsExtended(params map[string]string) (response UtilsGetLinkStatsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("utils.getLinkStats", params, &response, &vkErr)
	return
}

// UtilsGetServerTime returns the current time of the VK server.
//
// https://vk.com/dev/utils.getServerTime
func (vk *VK) UtilsGetServerTime(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("utils.getServerTime", params, &response, &vkErr)
	return
}

// UtilsGetShortLinkResponse struct
type UtilsGetShortLinkResponse object.UtilsShortLink

// UtilsGetShortLink allows to receive a link shortened via vk.cc.
//
// https://vk.com/dev/utils.getShortLink
func (vk *VK) UtilsGetShortLink(params map[string]string) (response UtilsGetShortLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("utils.getShortLink", params, &response, &vkErr)
	return
}

// UtilsResolveScreenNameResponse struct
type UtilsResolveScreenNameResponse object.UtilsDomainResolved

// UtilsResolveScreenName detects a type of object (e.g., user, community, application) and its ID by screen name.
//
// https://vk.com/dev/utils.resolveScreenName
func (vk *VK) UtilsResolveScreenName(params map[string]string) (response UtilsResolveScreenNameResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("utils.resolveScreenName", params)
	// Если короткое имя screen_name не занято, то будет возвращён пустой объект.
	if vkErr.Code != 0 || string(rawResponse) == "[]" {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "utils.resolveScreenName", params)
	}

	return
}
