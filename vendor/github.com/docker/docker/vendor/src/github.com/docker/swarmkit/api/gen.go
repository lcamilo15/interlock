package api

<<<<<<< HEAD
//go:generate protoc -I.:../protobuf:../vendor:../vendor/github.com/gogo/protobuf --gogoswarm_out=plugins=grpc+deepcopy+raftproxy+authenticatedwrapper,import_path=github.com/docker/swarmkit/api,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,Mtimestamp/timestamp.proto=github.com/docker/swarmkit/api/timestamp,Mduration/duration.proto=github.com/docker/swarmkit/api/duration,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,Mplugin/plugin.proto=github.com/docker/swarmkit/protobuf/plugin:. types.proto specs.proto objects.proto control.proto dispatcher.proto ca.proto snapshot.proto raft.proto
=======
//go:generate protoc -I.:../protobuf:../vendor:../vendor/github.com/gogo/protobuf --gogoswarm_out=plugins=grpc+deepcopy+raftproxy+authenticatedwrapper,import_path=github.com/docker/swarmkit/api,Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,Mtimestamp/timestamp.proto=github.com/docker/swarmkit/api/timestamp,Mduration/duration.proto=github.com/docker/swarmkit/api/duration,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,Mplugin/plugin.proto=github.com/docker/swarmkit/protobuf/plugin:. types.proto specs.proto objects.proto control.proto dispatcher.proto ca.proto snapshot.proto raft.proto health.proto
>>>>>>> 12a5469... start on swarm services; move to glade
