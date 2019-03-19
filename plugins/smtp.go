package plugins

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type SMTPWriter struct {
	Username           string   `json:"username"`
	Password           string   `json:"password"`
	Host               string   `json:"host"`
	Subject            string   `json:"subject"`
	FromAddress        string   `json:"fromAddress"`
	RecipientAddresses []string `json:"sendTos"`
	Level              int      `json:"level"`
}
