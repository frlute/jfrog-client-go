package services

import (
	"encoding/json"
	"net/http"

	"github.com/frlute/jfrog-client-go/artifactory/services/utils"
	"github.com/frlute/jfrog-client-go/auth"
	"github.com/frlute/jfrog-client-go/http/jfroghttpclient"
	"github.com/frlute/jfrog-client-go/utils/errorutils"
	"github.com/frlute/jfrog-client-go/utils/log"
)

type CreateReplicationService struct {
	client     *jfroghttpclient.JfrogHttpClient
	ArtDetails auth.ServiceDetails
}

func NewCreateReplicationService(client *jfroghttpclient.JfrogHttpClient) *CreateReplicationService {
	return &CreateReplicationService{client: client}
}

func (rs *CreateReplicationService) GetJfrogHttpClient() *jfroghttpclient.JfrogHttpClient {
	return rs.client
}

func (rs *CreateReplicationService) performRequest(params *utils.UpdateReplicationBody) error {
	content, err := json.Marshal(params)
	if err != nil {
		return errorutils.CheckError(err)
	}
	httpClientsDetails := rs.ArtDetails.CreateHttpClientDetails()
	utils.SetContentType("application/vnd.org.jfrog.artifactory.replications.ReplicationConfigRequest+json", &httpClientsDetails.Headers)
	url := rs.ArtDetails.GetUrl() + "api/replications/" + params.RepoKey
	log.Info("Creating replication...")
	resp, body, err := rs.client.SendPut(url, content, &httpClientsDetails)
	if err != nil {
		return err
	}
	if err = errorutils.CheckResponseStatusWithBody(resp, body, http.StatusOK, http.StatusCreated); err != nil {
		return err
	}
	log.Debug("Artifactory response:", resp.Status)
	log.Info("Done creating replication.")
	return nil
}

func (rs *CreateReplicationService) CreateReplication(params CreateReplicationParams) error {
	return rs.performRequest(utils.CreateUpdateReplicationBody(params.ReplicationParams))
}

func NewCreateReplicationParams() CreateReplicationParams {
	return CreateReplicationParams{}
}

type CreateReplicationParams struct {
	utils.ReplicationParams
}
