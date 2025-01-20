module github.com/LavaJover/DronCryptoWallet/api-gateway

go 1.23.2

require (
	github.com/LavaJover/DronCryptoWallet/auth-service v0.0.0
	github.com/LavaJover/DronCryptoWallet/wallet-service v0.0.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	google.golang.org/grpc v1.69.4
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/LavaJover/DronCryptoWallet/auth-service => ../auth-service

replace github.com/LavaJover/DronCryptoWallet/wallet-service => ../wallet-service
