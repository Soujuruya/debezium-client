package debezium_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	getConnectorsStatuses = "/connectors?expand=status"
	postCreateConnector   = "/connectors"
	deleteConnector       = "/connectors/"
	getConnector          = "/connectors/"
)

// Получение статуса коннекторов
func (c *Client) GetConnectorsStatuses(ctx context.Context) (GetConnectorsResponse, error) {
	var connectorResponse GetConnectorsResponse
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+getConnectorsStatuses, nil)

	if err != nil {
		return GetConnectorsResponse{}, fmt.Errorf("GetConnectorsStatuses.NewRequestWithContext: %w", err)
	}
	resp, err := c.cc.Do(req)

	if err != nil {
		return GetConnectorsResponse{}, fmt.Errorf("GetConnectorsStatuses.Client.Do: %w", err)
	}
	defer resp.Body.Close()
	//ВСЕ BODY ВСЕГДА ЗАКРЫВАЕ!!!

	if err := json.NewDecoder(resp.Body).Decode(&connectorResponse); err != nil {
		return GetConnectorsResponse{}, fmt.Errorf("GetConnectorsStatuses.UnmarshalJSON: %w", err)
	}

	return GetConnectorsResponse{}, nil
}

// Создание коннектора
func (c *Client) PostCreateConnectors(ctx context.Context, data CreateConnectorRequest) (bool, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("json.Marshal: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+postCreateConnector, bytes.NewBuffer(d))

	if err != nil {
		return false, fmt.Errorf("PostCreateConnectors.NewRequestWithContext: %w", err)
	}
	resp, err := c.cc.Do(req)

	if err != nil {
		return false, fmt.Errorf("PostCreateConnectors.Client.Do: %w", err)
	}
	defer resp.Body.Close()
	//ВСЕ BODY ВСЕГДА ЗАКРЫВАЕ!!!
	if resp.StatusCode != http.StatusCreated {
		var errorResponse CreateConnectorErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return false, fmt.Errorf("PostCreateConnectors.UnmarshalJSON: %w", err)
		}
		return false, fmt.Errorf("PostCreateConnectors.Client.Do: status code %s", errorResponse.Message)
	}

	var connectorResponse CreateConnectorResponse
	if err := json.NewDecoder(resp.Body).Decode(&connectorResponse); err != nil {
		return false, fmt.Errorf("PostCreateConnectors.UnmarshalJSON: %w", err)
	}

	return true, nil
}

// Метод удаления коннектора по имени
func (c *Client) DeleteConnector(ctx context.Context, connectorName string) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+deleteConnector+connectorName, nil)

	if err != nil {
		return fmt.Errorf("DeleteConnector.NewRequestWithContext: %w", err)
	}

	resp, err := c.cc.Do(req)

	if err != nil {
		return fmt.Errorf("DeleteConnector.Client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		var errorResponse DeleteConnectorErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("DeleteConnector.UnmarshalJSON: %w", err)
		}
		return fmt.Errorf("DeleteConnector.Client.Do: status code %s", errorResponse.Message)
	}

	return nil
}

// Метод получения данных о коннекторе по имени
func (c *Client) GetConnectorByName(ctx context.Context, connectorName string) (ConnectorInfoResponse, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+getConnector+connectorName, nil)

	if err != nil {
		return ConnectorInfoResponse{}, fmt.Errorf("GetConnectorByName.NewRequestWithContext: %w", err)
	}

	resp, err := c.cc.Do(req)

	if err != nil {
		return ConnectorInfoResponse{}, fmt.Errorf("GetConnectorByName.Client.Do: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse GetConnectorByNameErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return ConnectorInfoResponse{}, fmt.Errorf("GetConnectorByName.UnmarshalJSON: %w", err)
		}
		return ConnectorInfoResponse{}, fmt.Errorf("GetConnectorByName.Client.Do: status code %s", errorResponse.Message)
	}

	var connectorResponse ConnectorInfoResponse

	if err := json.NewDecoder(resp.Body).Decode(&connectorResponse); err != nil {
		return ConnectorInfoResponse{}, fmt.Errorf("GetConnectorByName.UnmarchalJSON: %w", err)
	}

	return connectorResponse, nil
}
