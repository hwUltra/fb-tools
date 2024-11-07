package mqttx

type Conf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username,optional"`
	Password string `json:"password,optional"`
	ClientId string `json:"clientId,optional"`
	//Qos      byte   `json:"qos"`
	//Tls      bool   `json:"tls"`
	//Action   string `json:"action"`
	//Topic    string `json:"topic"`
	//CaCert   string `json:"caCert"`
}

type MqtMsg struct {
	Topic     string
	Duplicate bool
	Qos       byte
	Retained  bool
	MessageID uint16
	Payload   []byte
	Ack       func()
}
