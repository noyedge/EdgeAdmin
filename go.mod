module github.com/noyedge/EdgeAdmin

go 1.21

replace github.com/noyedge/EdgeCommon => ../EdgeCommon

require (
	github.com/cespare/xxhash/v2 v2.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/iwind/TeaGo v0.0.0-20240429060313-31a7bc8e9cc9
	github.com/iwind/gosock v0.0.0-20211103081026-ee4652210ca4
	github.com/miekg/dns v1.1.43
	github.com/noyedge/EdgeCommon v0.0.0-00010101000000-000000000000
	github.com/shirou/gopsutil/v3 v3.22.5
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/xlzd/gotp v0.1.0
	golang.org/x/crypto v0.22.0
	golang.org/x/sys v0.19.0
	google.golang.org/grpc v1.63.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/shabbyrobe/xmlwriter v0.0.0-20200208144257-9fca06d00ffa // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
