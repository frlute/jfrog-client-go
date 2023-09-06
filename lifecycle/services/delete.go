package services

import (
	"net/http"
	"path"
	"strconv"

	"github.com/frlute/jfrog-client-go/utils"
	"github.com/frlute/jfrog-client-go/utils/errorutils"
)

func (rbs *ReleaseBundlesService) DeleteReleaseBundle(rbDetails ReleaseBundleDetails, params ReleaseBundleQueryParams) error {
	queryParams := getProjectQueryParam(params.ProjectKey)
	queryParams[async] = strconv.FormatBool(params.Async)
	restApi := path.Join(releaseBundleBaseApi, records, rbDetails.ReleaseBundleName, rbDetails.ReleaseBundleVersion)
	requestFullUrl, err := utils.BuildUrl(rbs.GetLifecycleDetails().GetUrl(), restApi, queryParams)
	if err != nil {
		return err
	}
	httpClientsDetails := rbs.GetLifecycleDetails().CreateHttpClientDetails()
	resp, body, err := rbs.client.SendDelete(requestFullUrl, nil, &httpClientsDetails)
	if err != nil {
		return err
	}
	return errorutils.CheckResponseStatusWithBody(resp, body, http.StatusNoContent)
}
