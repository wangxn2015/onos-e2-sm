module github.com/wangxn2015/onos-e2-sm/servicemodels/e2smkpmv2

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/onosproject/onos-api/go v0.7.110
	github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_kpm_v2_go v0.8.7
	github.com/wangxn2015/onos-lib-go v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.27.1
	gotest.tools v2.2.0+incompatible
)

replace github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_kpm_v2_go => /home/baicells/go_project/modified-onos-module/onos-e2-sm/servicemodels/e2sm_kpm_v2_go

replace github.com/wangxn2015/onos-lib-go => /home/baicells/go_project/modified-onos-module/onos-lib-go
