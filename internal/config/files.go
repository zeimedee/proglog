package config

import (
	"os"
	"path/filepath"
)

var (
	CAFile               = ConfigFile("ca.pem")
	ServerCertFile       = ConfigFile("server.pem")
	ServerKeyFile        = ConfigFile("server-key.pem")
	ClientCertFile       = ConfigFile("client.pem")
	ClientKeyFile        = ConfigFile("client-key.pem")
	RootClientCertFile   = ConfigFile("root-client.pem")
	RootClientKeyFile    = ConfigFile("root-client-key.pem")
	NoBodyClientCertFile = ConfigFile("nobody-client.pem")
	NoBodyClientKeyFile  = ConfigFile("nobody-client-key.pem")
	ACLModelFile         = ConfigFile("model.conf")
	ACLPolicyFile        = ConfigFile("policy.csv")
)

func ConfigFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homedir, "/Desktop/goProjects/proglog/certs", filename)
}
