package gostSDK

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dpnetca/gostSDK/internal/agents"
	"github.com/dpnetca/gostSDK/internal/client"
	"github.com/dpnetca/gostSDK/internal/contracts"
	"github.com/dpnetca/gostSDK/internal/factions"
	"github.com/dpnetca/gostSDK/internal/fleet"
	"github.com/dpnetca/gostSDK/models"
	"github.com/dpnetca/gostSDK/internal/system"
)

type SpaceTrader struct {
	Client    *client.Client
	Agents    *agents.Client
	Fleet     *fleet.Client
	System    *system.Client
	Contracts *contracts.Client
	Factions  *factions.Client
}

func NewSpaceTrader(token string) *SpaceTrader {
	client := client.New(token)
	agent := agents.New(client)
	fleet := fleet.New(client)
	system := system.New(client)
	contacts := contracts.New(client)
	factions := factions.New(client)

	return &SpaceTrader{
		Client:    client,
		Agents:    agent,
		Fleet:     fleet,
		System:    system,
		Contracts: contacts,
		Factions:  factions,
	}
}

type Status struct {
	Status      string      `json:"status"`
	Version     string      `json:"version"`
	ResetDate   string      `json:"resetDate"`
	Description string      `json:"description"`
	Stats       StatusStats `json:"stats"`
}

type StatusStats struct {
	Agents    int `json:"agents"`
	Ships     int `json:"ships"`
	Systems   int `json:"systems"`
	Waypoints int `json:"waypoints"`
}

// Return the status of the game server. This also includes a few global elements, such as announcements, server reset dates and leaderboards.
func (s SpaceTrader) GetStatus() (*Status, error) {

	req, _ := http.NewRequest("GET", s.Client.Base_url, nil)
	res, err := s.Client.Send(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var status Status
	if err = json.Unmarshal(data, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

type registerNewAgentRequest struct {
	Faction string `json:"faction"`
	Symbol  string `json:"symbol"`
	Email   string `json:"email"`
}

type registerNewAgentResponse struct {
	Data NewAgent `json:"data"`
}

type NewAgent struct {
	Agent models.Agent `json:"agent"`
	// Contract ... `json:"contract"`
	// Faction ... `json:"faction"`
	Ship  models.Ship `json:"ship"`
	Token string      `json:"token"`
}

func (s SpaceTrader) RegisterNewAgent(faction, symbol, email string) (*NewAgent, error) {

	registrationData := registerNewAgentRequest{
		Faction: faction,
		Symbol:  symbol,
		Email:   email,
	}
	body, err := json.Marshal(registrationData)
	if err != nil {
		return nil, err
	}
	url := s.Client.Base_url + "/register"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	res, err := s.Client.Send(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%s", data)
	}
	var newAgentReponse registerNewAgentResponse
	if err = json.Unmarshal(data, &newAgentReponse); err != nil {
		return nil, err
	}
	return &newAgentReponse.Data, nil
}
