package unet

import (
	"github.com/xh4n3/ucloud-sdk-go/ucloud"
)

type AllocateEIPParams struct {
	ucloud.CommonRequest

	OperatorName string
	Region       string
	Bandwidth    int
	ChargeType   string
	Quantity     int
}

type EIPAddr struct {
	OperatorName string
	IP           string
}

type EIPSet struct {
	EIPId         string
	CurBandwidth  float32
	Weight        int
	BandwidthType int
	Bandwidth     int
	Status        string
	ChargeType    string
	CreateTime    int
	ExpireTime    int
	Name          string
	Tag           string
	Remark        string
	EIPAddr       *[]EIPAddr
	Resource      *ucloud.Resource
}

type AllocateEIPResponse struct {
	ucloud.CommonResponse

	EIPIds *[]string
	EIPSet *[]EIPSet
}

func (u *UNet) AllocateEIP(params *AllocateEIPParams) (*AllocateEIPResponse, error) {
	response := &AllocateEIPResponse{}
	err := u.DoRequest("AllocateEIP", params, response)

	return response, err
}

type DescribeEIPParams struct {
	ucloud.CommonRequest

	Region string
	EIPIds []string
	OffSet int
	Limit  int
}

type DescribeEIPResponse struct {
	ucloud.CommonResponse

	TotalCount     int
	TotalBandwidth int
	EIPSet         *[]EIPSet
}

func (u *UNet) DescribeEIP(params *DescribeEIPParams) (*DescribeEIPResponse, error) {
	response := &DescribeEIPResponse{}
	err := u.DoRequest("DescribeEIP", params, response)

	return response, err
}

type UpdateEIPAttributeParams struct {
	ucloud.CommonRequest

	Region string
	EIPId  string
	Name   string
	Tag    string
	Remark string
}

type UpdateEIPAttributeResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) UpdateEIPAttribute(params *UpdateEIPAttributeParams) (*UpdateEIPAttributeResponse, error) {
	response := &UpdateEIPAttributeResponse{}
	err := u.DoRequest("UpdateEIPAttribute", params, response)

	return response, err
}

type ReleaseEIPParams struct {
	ucloud.CommonRequest

	Region string
	EIPId  string
}

type ReleaseEIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ReleaseEIP(params *ReleaseEIPParams) (*ReleaseEIPResponse, error) {
	response := &ReleaseEIPResponse{}
	err := u.DoRequest("ReleaseEIP", params, response)

	return response, err
}

type BindEIPParams struct {
	ucloud.CommonRequest

	Region       string
	EIPId        string
	ResourceType string
	ResourceId   string
}

type BindEIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) BindEIP(params *BindEIPParams) (*BindEIPResponse, error) {
	response := &BindEIPResponse{}
	err := u.DoRequest("BindEIP", params, response)

	return response, err
}

type UnBindEIPParams struct {
	ucloud.CommonRequest

	Region       string
	EIPId        string
	ResourceType string
	ResourceId   string
}

type UnBindEIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) UnBindEIP(params *UnBindEIPParams) (*UnBindEIPResponse, error) {
	response := &UnBindEIPResponse{}
	err := u.DoRequest("UnBindEIP", params, response)

	return response, err
}

type ModifyEIPBandwidthParams struct {
	ucloud.CommonRequest

	Region    string
	EIPId     string
	Bandwidth int
}

type ModifyEIPBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ModifyEIPBandwidth(params *ModifyEIPBandwidthParams) (*ModifyEIPBandwidthResponse, error) {
	response := &ModifyEIPBandwidthResponse{}
	err := u.DoRequest("ModifyEIPBandwidth", params, response)

	return response, err
}

type ModifyEIPWeightParams struct {
	ucloud.CommonRequest

	Region string
	EIPId  string
	Weight int
}

type ModifyEIPWeightResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ModifyEIPWeight(params *ModifyEIPWeightParams) (*ModifyEIPWeightResponse, error) {
	response := &ModifyEIPWeightResponse{}
	err := u.DoRequest("ModifyEIPWeight", params, response)

	return response, err
}

type GetEIPPriceParams struct {
	ucloud.CommonRequest

	OperatorName string
	Bandwidth    int
	ChargeType   string
}

type PriceSet struct {
	ChargeType    string
	Price         float32
	PurchaseValue int
}
type GetEIPPriceResponse struct {
	ucloud.CommonResponse

	PriceSet *[]PriceSet
}

func (u *UNet) GetEIPPrice(params *GetEIPPriceParams) (*GetEIPPriceResponse, error) {
	response := &GetEIPPriceResponse{}
	err := u.DoRequest("GetEIPPrice", params, response)

	return response, err
}

type AllocateVIPParams struct {
	ucloud.CommonRequest

	Region string
	Count  int
}

type AllocateVIPResponse struct {
	ucloud.CommonResponse

	DataSet []string
}

func (u *UNet) AllocateVIP(params *AllocateVIPParams) (*AllocateVIPResponse, error) {
	response := &AllocateVIPResponse{}
	err := u.DoRequest("AllocateVIP", params, response)

	return response, err
}

type DescribeVIPParams struct {
	ucloud.CommonRequest

	Region string
}

type DescribeVIPResponse struct {
	ucloud.CommonResponse

	DataSet []string
}

func (u *UNet) DescribeVIP(params *DescribeVIPParams) (*DescribeVIPResponse, error) {
	response := &DescribeVIPResponse{}
	err := u.DoRequest("DescribeVIP", params, response)

	return response, err
}

type ReleaseVIPParams struct {
	ucloud.CommonRequest

	Region string
	VIP    string
}

type ReleaseVIPResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ReleaseVIP(params *ReleaseVIPParams) (*ReleaseVIPResponse, error) {
	response := &ReleaseVIPResponse{}
	err := u.DoRequest("ReleaseVIP", params, response)

	return response, err
}

type CreateSecurityGroupParams struct {
	ucloud.CommonRequest

	Region      string
	GroupName   string
	Description string
	Rule        []string
}

type CreateSecurityGroupResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) CreateSecurityGroup(params *CreateSecurityGroupParams) (*CreateSecurityGroupResponse, error) {
	response := &CreateSecurityGroupResponse{}
	err := u.DoRequest("CreateSecurityGroup", params, response)

	return response, err
}

type DescribeSecurityGroupParams struct {
	ucloud.CommonRequest

	Region       string
	ResourceType string
	ResourceId   int
	GroupId      string
}

type Rule struct {
	SrcIP        string
	Priority     int
	ProtocolType string
	DstPort      string
	RuleAction   string
}

type FirewallDataSet struct {
	GroupId    int
	GroupName  string
	CreateTime int
	Type       int
	Rule       *[]Rule
}

type DescribeSecurityGroupResponse struct {
	ucloud.CommonResponse

	DataSet []FirewallDataSet
}

func (u *UNet) DescribeSecurityGroup(params *DescribeSecurityGroupParams) (*DescribeSecurityGroupResponse, error) {
	response := &DescribeSecurityGroupResponse{}
	err := u.DoRequest("DescribeSecurityGroup", params, response)

	return response, err
}

type DescribeSecurityGroupResourceParams struct {
	ucloud.CommonRequest

	Region  string
	GroupId int
}

type DescribeSecurityGroupResourceResponse struct {
	ucloud.CommonResponse

	DataSet []string
}

func (u *UNet) DescribeSecurityGroupResource(params *DescribeSecurityGroupParams) (*DescribeSecurityGroupResourceResponse, error) {
	response := &DescribeSecurityGroupResourceResponse{}
	err := u.DoRequest("DescribeSecurityGroupResource", params, response)

	return response, err
}

type UpdateSecurityGroupParams struct {
	ucloud.CommonRequest

	Region  string
	GroupId int
	Rule    []string
}

type UpdateSecurityGroupResponse struct {
	ucloud.CommonResponse

	DataSet []string
}

func (u *UNet) UpdateSecurityGroup(params *UpdateSecurityGroupParams) (*UpdateSecurityGroupResponse, error) {
	response := &UpdateSecurityGroupResponse{}
	err := u.DoRequest("UpdateSecurityGroup", params, response)

	return response, err
}

type GrantSecurityGroupParams struct {
	ucloud.CommonRequest

	Region       string
	GroupId      int
	ResourceType string
	ResourceId   string
}

type GrantSecurityGroupResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) GrantSecurityGroup(params *GrantSecurityGroupParams) (*GrantSecurityGroupResponse, error) {
	response := &GrantSecurityGroupResponse{}
	err := u.DoRequest("GrantSecurityGroup", params, response)

	return response, err
}

type DeleteSecurityGroupParams struct {
	ucloud.CommonRequest

	Region  string
	GroupId int
}

type DeleteSecurityGroupResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) DeleteSecurityGroup(params *DeleteSecurityGroupParams) (*DeleteSecurityGroupResponse, error) {
	response := &DeleteSecurityGroupResponse{}
	err := u.DoRequest("DeleteSecurityGroup", params, response)

	return response, err
}

type CreateBandwidthPackageParams struct {
	ucloud.CommonRequest

	Region     string
	EIPId      string
	Bandwidth  int
	EnableTime int
	TimeRange  int
}

type CreateBandwidthPackageResponse struct {
	ucloud.CommonResponse

	BandwidthPackageId string
}

func (u *UNet) CreateBandwidthPackage(params *CreateBandwidthPackageParams) (*CreateBandwidthPackageResponse, error) {
	response := &CreateBandwidthPackageResponse{}
	err := u.DoRequest("CreateBandwidthPackage", params, response)

	return response, err
}

type DescribeBandwidthPackageParams struct {
	ucloud.CommonRequest

	Region string
	Limit  int
	OffSet int
}

type BandwidthPackageDataSet struct {
	BandwidthPackageId string
	EnableTime         int
	DisableTime        int
	CreateTime         int
	Bandwidth          int
	EIPId              string
	EIPAddr            *[]EIPAddr
}

type DescribeBandwidthPackageResponse struct {
	ucloud.CommonResponse

	TotalCount int
	DataSets   *[]BandwidthPackageDataSet
}

func (u *UNet) DescribeBandwidthPackage(params *DescribeBandwidthPackageParams) (*DescribeBandwidthPackageResponse, error) {
	response := &DescribeBandwidthPackageResponse{}
	err := u.DoRequest("DescribeBandwidthPackage", params, response)

	return response, err
}

type DeleteBandwidthPackageParams struct {
	ucloud.CommonRequest

	Region             string
	BandwidthPackageId string
}

type DeleteBandwidthPackageResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) DeleteBandwidthPackage(params *DeleteBandwidthPackageParams) (*DeleteBandwidthPackageResponse, error) {
	response := &DeleteBandwidthPackageResponse{}
	err := u.DoRequest("DeleteBandwidthPackage", params, response)

	return response, err
}

type DescribeShareBandwidthParams struct {
	ucloud.CommonRequest

	Region string
}

type DescribeShareBandwidthResponse struct {
	ucloud.CommonResponse

	TotalCount int
	DataSet    *[]UnetShareBandwidthSet
}

type UnetShareBandwidthSet struct {
	Name             string
	ShareBandwidth   int
	ShareBandwidthId string
	ChargeType       string
	CreateTime       int64
	ExpireTime       int64
	EIPSet           *[]EIPSet
}

func (u *UNet) DescribeShareBandwidth(params *DescribeShareBandwidthParams) (*DescribeShareBandwidthResponse, error) {
	response := &DescribeShareBandwidthResponse{}
	err := u.DoRequest("DescribeShareBandwidth", params, response)

	return response, err
}

type ResizeShareBandwidthParams struct {
	ucloud.CommonRequest

	Region           string
	ShareBandwidth   int
	ShareBandwidthId string
}

type ResizeShareBandwidthResponse struct {
	ucloud.CommonResponse
}

func (u *UNet) ResizeShareBandwidth(params *ResizeShareBandwidthParams) (*ResizeShareBandwidthResponse, error) {
	response := &ResizeShareBandwidthResponse{}
	err := u.DoRequest("ResizeShareBandwidth", params, response)

	return response, err
}

type DescribeBandwidthUsageParams struct {
	ucloud.CommonRequest

	Region string
	OffSet int
	Limit  int
}

type DescribeBandwidthUsageResponse struct {
	ucloud.CommonResponse

	TotalCount int
	EIPSet     *[]UnetBandwidthUsageEIPSet
}

type UnetBandwidthUsageEIPSet struct {
	EIPId        string
	CurBandwidth float32
}

func (u *UNet) DescribeBandwidthUsage(params *DescribeBandwidthUsageParams) (*DescribeBandwidthUsageResponse, error) {
	response := &DescribeBandwidthUsageResponse{}
	err := u.DoRequest("DescribeBandwidthUsage", params, response)

	return response, err
}
