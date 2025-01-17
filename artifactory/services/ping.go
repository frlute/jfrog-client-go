package services

import (
	"net/http"

	"github.com/frlute/jfrog-client-go/auth"
	"github.com/frlute/jfrog-client-go/http/jfroghttpclient"
	"github.com/frlute/jfrog-client-go/utils"
	"github.com/frlute/jfrog-client-go/utils/errorutils"
	"github.com/frlute/jfrog-client-go/utils/log"
)

type PingService struct {
	client     *jfroghttpclient.JfrogHttpClient
	artDetails *auth.ServiceDetails
}

func NewPingService(artDetails auth.ServiceDetails, client *jfroghttpclient.JfrogHttpClient) *PingService {
	return &PingService{artDetails: &artDetails, client: client}
}

func (ps *PingService) GetArtifactoryDetails() auth.ServiceDetails {
	return *ps.artDetails
}

func (ps *PingService) GetJfrogHttpClient() *jfroghttpclient.JfrogHttpClient {
	return ps.client
}

func (ps *PingService) IsDryRun() bool {
	return false
}

func (ps *PingService) Ping() ([]byte, error) {
	url, err := utils.BuildUrl(ps.GetArtifactoryDetails().GetUrl(), "api/system/ping", nil)
	if err != nil {
		return nil, err
	}
	httpClientDetails := ps.GetArtifactoryDetails().CreateHttpClientDetails()
	resp, body, _, err := ps.client.SendGet(url, true, &httpClientDetails)
	if err != nil {
		return nil, err
	}
	if err = errorutils.CheckResponseStatusWithBody(resp, body, http.StatusOK); err != nil {
		return body, err
	}
	log.Debug("Artifactory response:", resp.Status)
	return body, nil
}
