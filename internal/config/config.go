package config

import (
	"log"
	"os"
	"sync"

	"github.com/iostrovok/godotenv"

	vip "github.com/iostrovok/coinbaseapi/internal/config/viper"
)

const (
	// https://docs.cdp.coinbase.com/exchange/docs/sandbox
	BaseUrl = "https://api.coinbase.com"

	SandboxRestApiUrl               = "https://api-public.sandbox.exchange.coinbase.com"
	SandboxWebsocketFeedUrl         = "wss://ws-feed-public.sandbox.exchange.coinbase.com"
	SandboxWebsocketDirectFeedUrl   = "wss://ws-direct.sandbox.exchange.coinbase.com"
	SandboxFixApiOrderEntry42Url    = "tcp+ssl://fix-public.sandbox.exchange.coinbase.com:4198"
	SandboxFixApiOrderEntry50Sp2Url = "tcp+ssl://fix-ord.sandbox.exchange.coinbase.com:6121"
	SandboxFixApiMarketData50Sp2Url = "tcp+ssl://fix-md.sandbox.exchange.coinbase.com:6121"

	RestApiUrl               = "https://api.exchange.coinbase.com"
	WebsocketFeedUrl         = "wss://ws-feed.exchange.coinbase.com"
	WebsocketDirectFeedUrl   = "wss://ws-direct.exchange.coinbase.com"
	FixApiOrderEntry42Url    = "tcp+ssl://fix.exchange.coinbase.com:4198"
	FixApiOrderEntry50Sp2Url = "tcp+ssl://fix-ord.exchange.coinbase.com:6121"
	FixApiMarketData50Sp2Url = "tcp+ssl://fix-md.exchange.coinbase.com:6121"
)

type Config struct {
	sync.RWMutex
	// for logs
	viper *vip.Viper

	// common info data
	DebugMode bool // DEBUG_MODE

	KeyName      string // aka "organizations/{org_id}/apiKeys/{key_id}"
	KeySecret    string // aka "-----BEGIN EC PRIVATE KEY-----\nYOUR PRIVATE KEY\n-----END EC PRIVATE KEY-----\n"
	CoinbaseHost string // default "https://api.coinbase.com"
}

var CG *Config

func New() *Config {
	return &Config{}
}

func Reload() *Config {
	CG = New()
	CG.Load()

	return CG
}

func CFG() *Config {
	if CG != nil {
		return CG
	}
	return Reload()
}

func ChDir(testHome string) {
	if testHome == "" {
		return
	}

	err := os.Chdir(testHome)
	if err != nil {
		panic(err)
	}
}

func (c *Config) Load() {
	c.Lock()
	defer c.Unlock()

	//sandboxSSLCertificateB, err := SSLCertificate.ReadFile("certificates/sandbox_ssl_certificate.pem")
	//if err != nil {
	//	return panic(err)
	//}
	//
	//sandboxSSLCertificate, err := tls.X509KeyPair(CertFilePath, KeyFilePath)
	//if err != nil {
	//	log.Fatalf("Error loading certificate and key file: %v", err)
	//}

	viper := vip.New()
	c.viper = viper

	ChDir(viper.GetString("TEST_SOURCE_PATH"))
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
	}

	c.DebugMode = viper.GetBool("DEBUG")

	c.KeyName = viper.GetString("KEY_NAME")
	if c.KeyName == "" {
		log.Print("KEY_NAME not found")
	}

	c.KeySecret = viper.GetString("KEY_SECRET")
	if c.KeySecret == "" {
		log.Print("KEY_SECRET not found")
	}

	c.CoinbaseHost = viper.GetString("COINBASE_HOST")
	if c.CoinbaseHost == "" {
		c.CoinbaseHost = BaseUrl
	}

}
