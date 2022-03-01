module micro_one

go 1.16

require (
	cloud.google.com/go/bigquery v1.4.0 // indirect
	github.com/Azure/azure-sdk-for-go v32.4.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.9.0 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.1.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.2.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.1.0 // indirect
	github.com/OpenDNS/vegadns2client v0.0.0-20180418235048-a3fa4a771d87 // indirect
	github.com/akamai/AkamaiOPEN-edgegrid-golang v1.1.0 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.976 // indirect
	github.com/aws/aws-sdk-go v1.37.27 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cenkalti/backoff/v4 v4.1.0 // indirect
	github.com/cloudflare/cloudflare-go v0.14.0 // indirect
	github.com/cpu/goacmedns v0.1.1 // indirect
	github.com/dnsimple/dnsimple-go v0.63.0 // indirect
	github.com/ef-ds/deque v1.0.4 // indirect
	github.com/evanphx/json-patch/v5 v5.5.0 // indirect
	github.com/exoscale/egoscale v0.46.0 // indirect
	github.com/fsouza/go-dockerclient v1.7.3 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.4 // indirect
	github.com/golang/mock v1.4.1 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/google/pprof v0.0.0-20200229191704-1ebb73c60ed3 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gophercloud/gophercloud v0.16.0 // indirect
	github.com/gophercloud/utils v0.0.0-20210216074907-f6de111f2eae // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/iij/doapi v0.0.0-20190504054126-0bbf12d6d7df // indirect
	github.com/infobloxopen/infoblox-go-client v1.1.1 // indirect
	github.com/labbsr0x/bindman-dns-webhook v1.0.2 // indirect
	github.com/linode/linodego v0.25.3 // indirect
	github.com/liquidweb/liquidweb-go v1.6.3 // indirect
	github.com/m3o/m3o-go/client v0.0.0-20210421144725-8bfd7992ada3 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/namedotcom/go v0.0.0-20180403034216-08470befbe04 // indirect
	github.com/nrdcg/auroradns v1.0.1 // indirect
	github.com/nrdcg/desec v0.5.0 // indirect
	github.com/nrdcg/dnspod-go v0.4.0 // indirect
	github.com/nrdcg/goinwx v0.8.1 // indirect
	github.com/nrdcg/namesilo v0.2.1 // indirect
	github.com/nrdcg/porkbun v0.1.1 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oracle/oci-go-sdk v24.3.0+incompatible // indirect
	github.com/ovh/go-ovh v1.1.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pquerna/otp v1.3.0 // indirect
	github.com/rainycape/memcache v0.0.0-20150622160815-1031fa0ce2f2 // indirect
	github.com/sacloud/libsacloud v1.36.2 // indirect
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.7.0.20210127161313-bd30bebeac4f // indirect
	github.com/transip/gotransip/v6 v6.2.0 // indirect
	github.com/urfave/cli v1.22.4 // indirect
	github.com/vinyldns/go-vinyldns v0.0.0-20200917153823-148a5f6b8f14 // indirect
	github.com/vultr/govultr/v2 v2.0.0 // indirect
	go.opencensus.io v0.22.3 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/api v0.20.0 // indirect
	google.golang.org/grpc v1.27.1 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
	gopkg.in/ns1/ns1-go.v2 v2.4.4 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
