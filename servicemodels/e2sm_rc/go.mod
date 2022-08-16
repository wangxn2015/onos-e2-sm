module github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/onosproject/onos-api/go v0.9.11
	github.com/wangxn2015/onos-lib-go v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.0
	gotest.tools v2.2.0+incompatible
)

//replace github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc => /home/baicells/go_project/modified-onos-module/onos-e2-sm/servicemodels/e2sm_rc

replace github.com/wangxn2015/onos-lib-go => /home/baicells/go_project/modified-onos-module/onos-lib-go
