package checks

import (
	"github.com/Serviceware/rabbit-hole/v2"
	"strconv"
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
	Port uint `long:"port" description:"Option for port-listener"`
}

type ProtocolListenerOpts struct {
	Protocol rabbithole.Protocol `long:""`
}

func CheckHealth(client *rabbithole.Client, check Check, opts *CheckHealthOpts) int {

	switch check {
	case ALARMS:
		return handleResourceAlarm(client.HealthCheckAlarms)
	case LOCAL_ALARMS:
		return handleResourceAlarm(client.HealthCheckLocalAlarms)
	case CERTIFICATE_EXPIRATION:
		check := func() (rabbithole.HealthCheckStatus, error) {
			return client.HealthCheckCertificateExpiration(opts.CertificateExpiration.Within, opts.CertificateExpiration.Unit)
		}
		return handleGeneralCheck(check)
	case PORT_LISTENER:
		check := func() (rabbithole.PortListenerCheckStatus, error) {
			return client.HealthCheckPortListener(opts.PortListener.Port)
		}
		return handlePortListenerCheck(check)
	case PROTOCOL_LISTENER:
		check := func() (rabbithole.ProtocolListenerCheckStatus, error) {
			return client.HealthCheckProtocolListener(opts.ProtocolLister.Protocol)
		}
		return handleProtocolListenerCheck(check)
	case VIRTUAL_HOSTS:
		return handleGeneralCheck(client.HealthCheckVirtualHosts)
	case NODE_IS_MIRROR_SYNC_CRITICAL:
		return handleGeneralCheck(client.HealthCheckNodeIsMirrorSyncCritical)
	case NODE_IS_QUORUM_CRITICAL:
		return handleGeneralCheck(client.HealthCheckNodeIsQuorumCritical)
	}

	return UNKNOWN
}

func handleResourceAlarm(f func() (rabbithole.ResourceAlarmCheckStatus, error)) int {
	status, err := f()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	if status.Ok() {
		return OK
	} else {
		println(status.Status + " - " + status.Reason)
		for _, alarm := range status.Alarms {
			println(alarm.Node + " - " + alarm.Resource)
		}
		return WARNING
	}
}

func handleGeneralCheck(check func() (rabbithole.HealthCheckStatus, error)) int {
	status, err := check()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	if status.Ok() {
		return OK
	} else {
		println(status.Status + " - " + status.Reason)
		return WARNING
	}
}

func handlePortListenerCheck(check func() (rabbithole.PortListenerCheckStatus, error)) int {
	status, err := check()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	if status.Ok() {
		return OK
	} else {
		println(status.Status + " - " + status.Reason)
		println("port:", status.Port)
		println("missing:", status.Missing)
		println("ports:", strings.Join(toStringArray(status.Ports), ","))
		return WARNING
	}
}

func toStringArray(intList []uint) (res []string) {
	for _, u := range intList {
		res = append(res, strconv.Itoa(int(u)))
	}

	return res
}

func handleProtocolListenerCheck(check func() (rabbithole.ProtocolListenerCheckStatus, error)) int {
	status, err := check()

	if err != nil {
		println(err.Error())
		return UNKNOWN
	}

	if status.Ok() {
		return OK
	} else {
		println(status.Status + " - " + status.Reason)
		println("protocol:", status.Protocols)
		println("missing:", status.Missing)
		println("protocols:", strings.Join(status.Protocols, ","))
		return WARNING
	}
}