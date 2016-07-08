package main

import (
	"fmt"
	"os"
	"path/filepath"
<<<<<<< HEAD
=======
	"sync"
>>>>>>> 12a5469... start on swarm services; move to glade
	"testing"

	"github.com/docker/docker/cliconfig"
	"github.com/docker/docker/pkg/reexec"
<<<<<<< HEAD
=======
	"github.com/docker/engine-api/types/swarm"
>>>>>>> 12a5469... start on swarm services; move to glade
	"github.com/go-check/check"
)

func Test(t *testing.T) {
	reexec.Init() // This is required for external graphdriver tests

	if !isLocalDaemon {
		fmt.Println("INFO: Testing against a remote daemon")
	} else {
		fmt.Println("INFO: Testing against a local daemon")
	}

	check.TestingT(t)
}

func init() {
	check.Suite(&DockerSuite{})
}

type DockerSuite struct {
}

func (s *DockerSuite) TearDownTest(c *check.C) {
	unpauseAllContainers()
	deleteAllContainers()
	deleteAllImages()
	deleteAllVolumes()
	deleteAllNetworks()
}

func init() {
	check.Suite(&DockerRegistrySuite{
		ds: &DockerSuite{},
	})
}

type DockerRegistrySuite struct {
	ds  *DockerSuite
	reg *testRegistryV2
	d   *Daemon
}

func (s *DockerRegistrySuite) SetUpTest(c *check.C) {
	testRequires(c, DaemonIsLinux, RegistryHosting)
	s.reg = setupRegistry(c, false, "", "")
	s.d = NewDaemon(c)
}

func (s *DockerRegistrySuite) TearDownTest(c *check.C) {
	if s.reg != nil {
		s.reg.Close()
	}
	if s.d != nil {
		s.d.Stop()
	}
	s.ds.TearDownTest(c)
}

func init() {
	check.Suite(&DockerSchema1RegistrySuite{
		ds: &DockerSuite{},
	})
}

type DockerSchema1RegistrySuite struct {
	ds  *DockerSuite
	reg *testRegistryV2
	d   *Daemon
}

func (s *DockerSchema1RegistrySuite) SetUpTest(c *check.C) {
<<<<<<< HEAD
	testRequires(c, DaemonIsLinux, RegistryHosting)
=======
	testRequires(c, DaemonIsLinux, RegistryHosting, NotArm64)
>>>>>>> 12a5469... start on swarm services; move to glade
	s.reg = setupRegistry(c, true, "", "")
	s.d = NewDaemon(c)
}

func (s *DockerSchema1RegistrySuite) TearDownTest(c *check.C) {
	if s.reg != nil {
		s.reg.Close()
	}
	if s.d != nil {
		s.d.Stop()
	}
	s.ds.TearDownTest(c)
}

func init() {
	check.Suite(&DockerRegistryAuthHtpasswdSuite{
		ds: &DockerSuite{},
	})
}

type DockerRegistryAuthHtpasswdSuite struct {
	ds  *DockerSuite
	reg *testRegistryV2
	d   *Daemon
}

func (s *DockerRegistryAuthHtpasswdSuite) SetUpTest(c *check.C) {
	testRequires(c, DaemonIsLinux, RegistryHosting)
	s.reg = setupRegistry(c, false, "htpasswd", "")
	s.d = NewDaemon(c)
}

func (s *DockerRegistryAuthHtpasswdSuite) TearDownTest(c *check.C) {
	if s.reg != nil {
		out, err := s.d.Cmd("logout", privateRegistryURL)
		c.Assert(err, check.IsNil, check.Commentf(out))
		s.reg.Close()
	}
	if s.d != nil {
		s.d.Stop()
	}
	s.ds.TearDownTest(c)
}

func init() {
	check.Suite(&DockerRegistryAuthTokenSuite{
		ds: &DockerSuite{},
	})
}

type DockerRegistryAuthTokenSuite struct {
	ds  *DockerSuite
	reg *testRegistryV2
	d   *Daemon
}

func (s *DockerRegistryAuthTokenSuite) SetUpTest(c *check.C) {
	testRequires(c, DaemonIsLinux, RegistryHosting)
	s.d = NewDaemon(c)
}

func (s *DockerRegistryAuthTokenSuite) TearDownTest(c *check.C) {
	if s.reg != nil {
		out, err := s.d.Cmd("logout", privateRegistryURL)
		c.Assert(err, check.IsNil, check.Commentf(out))
		s.reg.Close()
	}
	if s.d != nil {
		s.d.Stop()
	}
	s.ds.TearDownTest(c)
}

func (s *DockerRegistryAuthTokenSuite) setupRegistryWithTokenService(c *check.C, tokenURL string) {
	if s == nil {
		c.Fatal("registry suite isn't initialized")
	}
	s.reg = setupRegistry(c, false, "token", tokenURL)
}

func init() {
	check.Suite(&DockerDaemonSuite{
		ds: &DockerSuite{},
	})
}

type DockerDaemonSuite struct {
	ds *DockerSuite
	d  *Daemon
}

func (s *DockerDaemonSuite) SetUpTest(c *check.C) {
	testRequires(c, DaemonIsLinux)
	s.d = NewDaemon(c)
}

func (s *DockerDaemonSuite) TearDownTest(c *check.C) {
	testRequires(c, DaemonIsLinux)
	if s.d != nil {
		s.d.Stop()
	}
	s.ds.TearDownTest(c)
}

const defaultSwarmPort = 2477

func init() {
	check.Suite(&DockerSwarmSuite{
		ds: &DockerSuite{},
	})
}

type DockerSwarmSuite struct {
<<<<<<< HEAD
	ds        *DockerSuite
	daemons   []*SwarmDaemon
	portIndex int
=======
	ds          *DockerSuite
	daemons     []*SwarmDaemon
	daemonsLock sync.Mutex // protect access to daemons
	portIndex   int
>>>>>>> 12a5469... start on swarm services; move to glade
}

func (s *DockerSwarmSuite) SetUpTest(c *check.C) {
	testRequires(c, DaemonIsLinux)
}

func (s *DockerSwarmSuite) AddDaemon(c *check.C, joinSwarm, manager bool) *SwarmDaemon {
	d := &SwarmDaemon{
		Daemon: NewDaemon(c),
		port:   defaultSwarmPort + s.portIndex,
	}
	d.listenAddr = fmt.Sprintf("0.0.0.0:%d", d.port)
	err := d.StartWithBusybox("--iptables=false") // avoid networking conflicts
	c.Assert(err, check.IsNil)

	if joinSwarm == true {
		if len(s.daemons) > 0 {
<<<<<<< HEAD
			c.Assert(d.Join(s.daemons[0].listenAddr, "", "", manager), check.IsNil)
		} else {
			aa := make(map[string]bool)
			aa["worker"] = true
			aa["manager"] = true
			c.Assert(d.Init(aa, ""), check.IsNil)
=======
			c.Assert(d.Join(swarm.JoinRequest{
				RemoteAddrs: []string{s.daemons[0].listenAddr},
				Manager:     manager}), check.IsNil)
		} else {
			c.Assert(d.Init(swarm.InitRequest{
				Spec: swarm.Spec{
					AcceptancePolicy: autoAcceptPolicy,
				},
			}), check.IsNil)
>>>>>>> 12a5469... start on swarm services; move to glade
		}
	}

	s.portIndex++
<<<<<<< HEAD
	s.daemons = append(s.daemons, d)
=======
	s.daemonsLock.Lock()
	s.daemons = append(s.daemons, d)
	s.daemonsLock.Unlock()
>>>>>>> 12a5469... start on swarm services; move to glade

	return d
}

func (s *DockerSwarmSuite) TearDownTest(c *check.C) {
	testRequires(c, DaemonIsLinux)
<<<<<<< HEAD
=======
	s.daemonsLock.Lock()
>>>>>>> 12a5469... start on swarm services; move to glade
	for _, d := range s.daemons {
		d.Stop()
	}
	s.daemons = nil
<<<<<<< HEAD
	s.portIndex = 0

=======
	s.daemonsLock.Unlock()

	s.portIndex = 0
>>>>>>> 12a5469... start on swarm services; move to glade
	s.ds.TearDownTest(c)
}

func init() {
	check.Suite(&DockerTrustSuite{
		ds: &DockerSuite{},
	})
}

type DockerTrustSuite struct {
	ds  *DockerSuite
	reg *testRegistryV2
	not *testNotary
}

func (s *DockerTrustSuite) SetUpTest(c *check.C) {
	testRequires(c, RegistryHosting, NotaryServerHosting)
	s.reg = setupRegistry(c, false, "", "")
	s.not = setupNotary(c)
}

func (s *DockerTrustSuite) TearDownTest(c *check.C) {
	if s.reg != nil {
		s.reg.Close()
	}
	if s.not != nil {
		s.not.Close()
	}

	// Remove trusted keys and metadata after test
	os.RemoveAll(filepath.Join(cliconfig.ConfigDir(), "trust"))
	s.ds.TearDownTest(c)
}
