package gaurun

import (
	"github.com/mercari/gaurun/gcm"

	"go.uber.org/zap"
)

var (
	// Toml configuration for Gaurun
	ConfGaurun ConfToml
	// push notification Queue
	QueueNotification chan RequestGaurunNotification
	// Stat for Gaurun
	StatGaurun StatApp
	// http client for APNs (token based push)
	APNSClient APNsClient
	// http client for APNs (certificate based push)
	CertAPNSClients map[NotificationOption]APNsClient = map[NotificationOption]APNsClient{}
	// http client for GCM/FCM
	GCMClient *gcm.Client
	// access and error logger
	LogAccess *zap.Logger
	LogError  *zap.Logger
	// sequence ID for numbering push
	SeqID uint64
)
