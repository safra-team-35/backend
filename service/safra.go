package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/safra-team-35/backend/domain"
	"github.com/safra-team-35/backend/domain/contract"
	"github.com/safra-team-35/backend/domain/entity"
)

type SafraTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type safraService struct {
	svc    *Service
	client *http.Client
}

//newSafraService return a new instance of the service
func newSafraService(svc *Service) contract.SafraService {
	client := svc.httpClient
	return &safraService{
		svc:    svc,
		client: client,
	}
}

func (s *safraService) GetMorningCalls() (data entity.SafraMorningCallsResponse, err resterrors.RestErr) {

	accessToken, err := s.getLoginToken()
	if err != nil {
		return data, err
	}

	statusCode, response, err := s.call(accessToken, http.MethodGet, domain.SafraMorningCallsURL, nil)
	if err != nil {
		return data, err
	}

	if statusCode != http.StatusOK {
		return data, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}

	jsonErr := json.Unmarshal(response, &data)
	if jsonErr != nil {
		logger.Error("Error to get token: ", err)
		return data, resterrors.NewInternalServerError("Service unavailable, contact the support")
	}

	return data, nil
}

//getLoginToken gets the login token for the first data api
func (s *safraService) getLoginToken() (token string, err resterrors.RestErr) {

	url := domain.SafraTokenURL
	payload := strings.NewReader("grant_type=client_credentials&scope=urn:opc:resource:consumer::all")

	req, reqErr := http.NewRequest(http.MethodPost, url, payload)
	if reqErr != nil {
		logger.Error("Error to create request to get token: ", reqErr)
		return token, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}

	authString := domain.SafraClientID + ":" + domain.SafraSecrect

	base64Token := base64.StdEncoding.EncodeToString([]byte(authString))

	req.Header.Add("authorization", "Basic "+base64Token)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("postman-token", "280d6ac2-0e1c-d7ed-fc20-85de145f3d1c")

	resp, reqErr := s.client.Do(req)
	if reqErr != nil {
		logger.Error("Error to get token: ", reqErr)
		return token, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("Error to get token: ", err)
		return token, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}

	loginResponse := SafraTokenResponse{}
	json.NewDecoder(resp.Body).Decode(&loginResponse)

	token = loginResponse.AccessToken

	return token, nil
}

func (s *safraService) call(token, method, url string, request interface{}) (statusCode int, response []byte, restErr resterrors.RestErr) {

	var data []byte
	var err error
	if request != nil {
		data, err = json.Marshal(request)
		if err != nil {
			logger.Error("Error to get token: ", err)
			return 0, nil, resterrors.NewInternalServerError("Service unavailable, contact the support")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		logger.Error("Error to get token: ", err)
		return 0, nil, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := s.svc.httpClient.Do(req)
	if err != nil {
		logger.Error("Error to get token: ", err)
		return 0, nil, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}

	defer res.Body.Close()

	response, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error("Error to get token: ", err)
		return 0, response, resterrors.NewInternalServerError("Error to get morning calls, please try again")
	}

	return res.StatusCode, response, nil
}
