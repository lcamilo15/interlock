package swarm

import "time"

// Swarm represents a swarm.
type Swarm struct {
	ID string
	Meta
	Spec Spec
}

// Spec represents the spec of a swarm.
type Spec struct {
	Annotations

	AcceptancePolicy AcceptancePolicy    `json:",omitempty"`
	Orchestration    OrchestrationConfig `json:",omitempty"`
	Raft             RaftConfig          `json:",omitempty"`
	Dispatcher       DispatcherConfig    `json:",omitempty"`
	CAConfig         CAConfig            `json:",omitempty"`
<<<<<<< HEAD
=======

	// DefaultLogDriver sets the log driver to use at task creation time if
	// unspecified by a task.
	//
	// Updating this value will only have an affect on new tasks. Old tasks
	// will continue use their previously configured log driver until
	// recreated.
	DefaultLogDriver *Driver `json:",omitempty"`
>>>>>>> 12a5469... start on swarm services; move to glade
}

// AcceptancePolicy represents the list of policies.
type AcceptancePolicy struct {
	Policies []Policy `json:",omitempty"`
}

// Policy represents a role, autoaccept and secret.
type Policy struct {
	Role       NodeRole
	Autoaccept bool
	Secret     *string `json:",omitempty"`
}

<<<<<<< HEAD
// OrchestrationConfig represents ochestration configuration.
=======
// OrchestrationConfig represents orchestration configuration.
>>>>>>> 12a5469... start on swarm services; move to glade
type OrchestrationConfig struct {
	TaskHistoryRetentionLimit int64 `json:",omitempty"`
}

// RaftConfig represents raft configuration.
type RaftConfig struct {
	SnapshotInterval           uint64 `json:",omitempty"`
	KeepOldSnapshots           uint64 `json:",omitempty"`
	LogEntriesForSlowFollowers uint64 `json:",omitempty"`
	HeartbeatTick              uint32 `json:",omitempty"`
	ElectionTick               uint32 `json:",omitempty"`
}

// DispatcherConfig represents dispatcher configuration.
type DispatcherConfig struct {
	HeartbeatPeriod uint64 `json:",omitempty"`
}

// CAConfig represents CA configuration.
type CAConfig struct {
	NodeCertExpiry time.Duration `json:",omitempty"`
<<<<<<< HEAD
=======
	ExternalCAs    []*ExternalCA `json:",omitempty"`
}

// ExternalCAProtocol represents type of external CA.
type ExternalCAProtocol string

// ExternalCAProtocolCFSSL CFSSL
const ExternalCAProtocolCFSSL ExternalCAProtocol = "cfssl"

// ExternalCA defines external CA to be used by the cluster.
type ExternalCA struct {
	Protocol ExternalCAProtocol
	URL      string
	Options  map[string]string `json:",omitempty"`
>>>>>>> 12a5469... start on swarm services; move to glade
}

// InitRequest is the request used to init a swarm.
type InitRequest struct {
	ListenAddr      string
	ForceNewCluster bool
	Spec            Spec
}

// JoinRequest is the request used to join a swarm.
type JoinRequest struct {
	ListenAddr  string
	RemoteAddrs []string
	Secret      string // accept by secret
	CACertHash  string
	Manager     bool
}

// LocalNodeState represents the state of the local node.
type LocalNodeState string

const (
	// LocalNodeStateInactive INACTIVE
	LocalNodeStateInactive LocalNodeState = "inactive"
	// LocalNodeStatePending PENDING
	LocalNodeStatePending LocalNodeState = "pending"
	// LocalNodeStateActive ACTIVE
	LocalNodeStateActive LocalNodeState = "active"
	// LocalNodeStateError ERROR
	LocalNodeStateError LocalNodeState = "error"
)

// Info represents generic information about swarm.
type Info struct {
	NodeID string

	LocalNodeState   LocalNodeState
	ControlAvailable bool
	Error            string

	RemoteManagers []Peer
	Nodes          int
	Managers       int
	CACertHash     string
}

// Peer represents a peer.
type Peer struct {
	NodeID string
	Addr   string
}
