package db

import (
	"context"
	"github.com/laeni/pri-dns/types"
)

type Store interface {
	// FindForwardByHostAndName 查询客户端对应的转发配置，当 host 为 “” 时表示查询全局配置.
	FindForwardByHostAndName(ctx context.Context, host, name string) []Forward

	// FindDomainByHostAndName 查询 qname 的解析记录。如果 host 不为空，则查询host下的解析，如果为空则只查询全局解析
	FindDomainByHostAndName(ctx context.Context, host, qname string) []Domain
}

type RecordFilter interface {
	NameVal() string
	TypeVal() string
}

// Domain 解析记录表.
type Domain struct {
	ID int64
	// 客户端地址（生效范围）。<br />如果全局生效，则该字段为空。
	Host string
	// 主机记录。由于可能存在泛域名，所以为了便于使用索引，存储时将采用反转格式，如：example.com
	Name string
	// 记录值
	Value string
	// TTL
	Ttl int32
	// 状态。<br />ENABLE-启用
	Status string
	// 记录类型。带"NO_"前缀的表示用于禁用全局解析。<br />A / NO_ALL / NO_A
	Type string
	// 优先级。值越小优先级越高。
	Priority int32
	// 创建时间
	CreateTime types.LocalTime
	// 修改时间
	UpdateTime types.LocalTime
}

func (d Domain) NameVal() string {
	return d.Name
}
func (d Domain) TypeVal() string {
	return d.Type
}

// Forward 转发配置.
type Forward struct {
	ID int64
	// 客户端地址（生效范围）。<br />如果全局生效，则该字段为空。
	Host string
	// 需要转发解析的域名
	Name string
	// 转发目标DNS服务器
	DnsSvr []string
	// 解析记录，用于导出使用
	History []string
	// Allow-正常转发 Deny-否定全局解析
	Type string
	// 状态。<br />ENABLE-启用
	Status string
	// 创建时间
	CreateTime types.LocalTime
	// 修改时间
	UpdateTime types.LocalTime
}

func (f Forward) NameVal() string {
	return f.Name
}
func (f Forward) TypeVal() string {
	return f.Type
}
