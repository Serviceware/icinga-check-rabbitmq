package internal

import (
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
	"strings"
)

type Check string

const (
	ALARMS                       Check = "alarms"
	LOCAL_ALARMS                 Check = "local-alarms"
	CERTIFICATE_EXPIRATION       Check = "certificate-expiration"
	PORT_LISTENER                Check = "port-listener"
	PROTOCOL_LISTENER            Check = "protocol-listener"
	VIRTUAL_HOSTS                Check = "virtual-hosts"
	NODE_IS_MIRROR_SYNC_CRITICAL Check = "node-is-mirror-sync-critical"
	NODE_IS_QUORUM_CRITICAL      Check = "node-is-quorum-critical"
)

type CheckHealthOpts struct {
	Alarms                   Void                 `command:"alarms" description:""`
	LocalAlarms              Void                 `command:"local-alarms" description:""`
	CertificateExpiration    CertExpirationOpts   `command:"certificate-expiration" description:""`
	PortListener             PortListenerOpts     `command:"port-listener" description:""`
	ProtocolLister           ProtocolListenerOpts `command:"protocol-listener" description:""`
	VirtualHosts             Void                 `command:"virtual-hosts" description:""`
	NodeIsMirrorSyncCritical Void                 `command:"node-is-mirror-sync-critical" description:""`
	NodeIsQuorumCritical     Void                 `command:"node-is-quorum-critical" description:""`
}

type CertExpirationOpts struct {
	Within uint                `long:"within" description:"Option for certificate-expiration"`
	Unit   rabbithole.TimeUnit `long:"unit" description:"Option for certificate-expiration"`
}

type PortListenerOpts struct {
	Port uint `long:"port" dscription:"Option for port-listener"`
}

type ProtocolListenerOpts struct {
	Protocol rabbithole.Protocol `long:""`
}

func CheckHealth(client *rabbithole.Client, check Check, opts *CheckHealthOpts) int {

	var health *rabbithole.Health
	var err error
	switch check {
	case ALARMS:
		health, err = client.HealthCheckAlarms()
	case LOCAL_ALARMS:
		health, err = client.HealthCheckLocalAlarms()
	case "certificate-expiration":
		health, err = client.HealthCheckCertificateExpiration(opts.CertificateExpiration.Within, opts.CertificateExpiration.Unit)
	case "port-listener":
		health, err = client.HealthCheckPortListenerListener(opts.PortListener.Port)
	case "protocol-listener":
		health, err = client.HealthCheckProtocolListener(opts.ProtocolLister.Protocol)
	case "virtual-hosts":
		health, err = client.HealthCheckVirtualHosts()
	case "node-is-mirror-sync-critical":
		health, err = client.HealthCheckNodeIsMirrorSyncCritical()
	case "node-is-quorum-critical":
		health, err = client.HealthCheckNodeIsQuorumCritical()
	}

	if health == nil {
		println(err.Error())
		return UNKNOWN
	}

	if health.Status != "ok" {
		println("status =", health.Status)
		println("reason =", health.Reason)
		println("missing = ", health.Missing)
		println("ports", strings.Join(health.Ports, ","))
		println("protocols", strings.Join(health.Protocols, ","))
		return WARNING
	} else {
		println("status=ok")
		return OK
	}
}
