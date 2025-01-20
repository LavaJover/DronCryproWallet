module github.com/LavaJover/DronCryptoWallet/api-gateway

go 1.23.2

require (
	github.com/LavaJover/DronCryptoWallet/auth-service v0.0.0
	google.golang.org/grpc v1.69.4
)

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)

replace github.com/LavaJover/DronCryptoWallet/auth-service => ../auth-service
