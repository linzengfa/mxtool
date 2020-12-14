// Copyright (c) 2020.
// ALL Rights reserved.
// @Description WXMapObj.go
// @Author moxiao
// @Date 2020/12/14 20:41

package mxmap

// 位置坐标
type Location struct {
	Lat float64 // 纬度
	Lng float64 // 经度
}

// 位置描述
type FormattedAddresses struct {
	Recommend string //经过腾讯地图优化过的描述方式，更具人性化特点
	Rough     string // 大致位置，可用于对位置的粗略描述
}

// 地址部件，address不满足需求时可自行拼接
type AddressComponent struct {
	Nation       string // 国家
	Province     string // 省
	City         string // 市
	District     string // 区，可能为空字串
	Street       string // 街道，可能为空字串
	StreetNumber string // 门牌，可能为空字串
}

// 行政区划信息
type AdInfo struct {
	NationCode string   `json:"nation_code"` // 国家代码
	AdCode     string   `json:"adcode"`      //行政区划代码
	CityCode   string   `json:"city_code"`   //城市
	Name       string   // 行政区划名称
	Location   Location //行政区划中心点坐标
	Nation     string   // 国家
	Province   string   // 省 / 直辖市
	City       string   // 市 / 地级区 及同级行政区划
	district   string   // 区 / 县级市 及同级行政区划
}

// 知名区域，如商圈或人们普遍认为有较高知名度的区域
type FamousArea struct {
	Id       string   // 地点唯一标识
	Title    string   // 名称/标题
	Location Location //行政区划中心点坐标
	Distance float64  `json:"_distance"` // 此参考位置到输入坐标的直线距离
	DirDesc  string   `json:"_dir_desc"` //此参考位置到输入坐标的方位关系，如：北、南、内
}

// 乡镇街道
type Town struct {
	Id       string   // 地点唯一标识
	Title    string   // 名称/标题
	Location Location //行政区划中心点坐标
	Distance float64  `json:"_distance"` // 此参考位置到输入坐标的直线距离
	DirDesc  string   `json:"_dir_desc"` //此参考位置到输入坐标的方位关系，如：北、南、内
}

// 坐标相对位置参考
type AddressReference struct {
	FamousArea   FamousArea `famous_area` // 知名区域，如商圈或人们普遍认为有较高知名度的区域
	Town         Town       // 知名区域，如商圈或人们普遍认为有较高知名度的区域
	LandmarkL1   FamousArea `json:"landmark_l_1"`  // 一级地标，可识别性较强、规模较大的地点、小区等
	LandmarkL2   FamousArea `json:"landmark_l2"`   // 二级地标，较一级地标更为精确，规模更小
	Street       FamousArea `json:"street"`        // 街道
	StreetNumber FamousArea `json:"street_number"` // 门牌
	Crossroad    FamousArea `json:"crossroad"`     // 交叉路口
	Water        FamousArea `json:"water"`         // 水系
}

// POI行政区划信息
type PosAdInfo struct {
	AdCode   string `json:"adcode"` // 行政区划代码
	Province string // 省 / 直辖市
	City     string // 市 / 地级区 及同级行政区划
	district string // 区 / 县级市 及同级行政区划
}

type Poi struct {
	Id       string    // 地点唯一标识
	Title    string    // 名称/标题
	Address  string    // 地址
	Category string    // POI分类
	Location Location  //行政区划中心点坐标
	AdInfo   PosAdInfo `json:"ad_info"`   //行政区划信息
	Distance float64   `json:"_distance"` // 此参考位置到输入坐标的直线距离
}

// 逆地址解析结果
type Result struct {
	Address            string             //地址描述
	FormattedAddresses FormattedAddresses `json:"formatted_addresses"` //位置描述
	AddressComponent   AddressComponent   `json:"address_component"`   //地址部件
	AdInfo             AdInfo             `json:"ad_info"`             //行政区划信息
	AddressReference   AddressReference   `json:"address_reference"`   //坐标相对位置参考
	PoiCount           int64              `json:"poi_count"`           //查询的周边poi的总数
	Pois               []Poi              `json:"pois"`                //POI数组，对象中每个子项为一个POI对象
}

// 逆地址解析应答
type GeoCoderResponse struct {
	Status    int64  `json:"status"`     //状态码 0为正常 310请求参数信息有误 311Key格式错误 306请求有护持信息请检查字符串 110请求来源未被授权
	Message   string `json:"message"`    //状态说明
	RequestId string `json:"request_id"` //本次请求的唯一标识
	Result    Result `json:"result"`     //逆地址解析结果
}
