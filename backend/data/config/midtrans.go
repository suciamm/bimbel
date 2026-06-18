package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client

func InitMidtrans() {
	// Server Key Sandbox
	midtrans.ServerKey = "Mid-server-TIb0fqEmsEQfWt9L8n_SyZh8"
	midtrans.Environment = midtrans.Sandbox

	SnapClient.New(midtrans.ServerKey, midtrans.Sandbox)
}
