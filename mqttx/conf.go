package mqttx

type Conf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	ClientId string `json:"clientId"`
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
