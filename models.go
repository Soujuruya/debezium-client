package debezium_client

type GetConnectorsResponse struct {
	Connectors []struct {
		Status struct {
			Name      string `json:"name"`
			Connector struct {
				State    string `json:"state"`
				WorkerId string `json:"worker_id"`
			} `json:"connector"`
			Tasks []struct {
				Id       int    `json:"id"`
				State    string `json:"state"`
				WorkerId string `json:"worker_id"`
			} `json:"tasks"`
			Type string `json:"type"`
		} `json:"status"`
	} `json:"connector"`
}

type CreateConnectorRequest struct {
	Name   string                `json:"name"`
	Config CreateConnectorConfig `json:"config"`
}

type CreateConnectorConfig struct {
	ConnectorClass       string            `json:"connector.class"`
	TasksMax             string            `json:"tasks.max"`
	DatabaseHostname     string            `json:"database.hostname"`
	DatabasePort         string            `json:"database.port"`
	DatabaseUser         string            `json:"database.user"`
	DatabasePassword     string            `json:"database.password"`
	DatabaseDbname       string            `json:"database.dbname"`
	DatabaseServerName   string            `json:"database.server.name"`
	AdditionalParametres map[string]string `json:"-,omitempty"`
}

type CreateConnectorErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type CreateConnectorResponse struct {
	Name   string                `json:"name"`
	Config CreateConnectorConfig `json:"config"`
	Tasks  []any                 `json:"tasks"`
	Type   string                `json:"type"`
}

type DeleteConnectorErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type GetConnectorByNameErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

type ConnectorInfoResponse struct {
	Name   string                `json:"name"`
	Config CreateConnectorConfig `json:"config"`
	Tasks  []ConnectorTask       `json:"tasks"`
	Type   string                `json:"type"`
}

type ConnectorTask struct {
	Connector string `json:"connector"`
	Task      int    `json:"task"`
}
