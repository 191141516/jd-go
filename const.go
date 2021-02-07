package jd

import (
	"github.com/cliod/jd-go/jd_method"
)

// 商品接口
const (
	GoodsJingfenQuery            = jd_method.GoodsJingfenQuery            // 京粉精选商品查询接口
	GoodsQuery                   = jd_method.GoodsQuery                   // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQuery = jd_method.GoodsPromotiongoodsinfoQuery // 根据skuid查询商品信息接口
	CategoryGoodsGet             = jd_method.CategoryGoodsGet             // 商品类目查询接口
	GoodsGigfieldQuery           = jd_method.GoodsGigfieldQuery           // 商品详情查询接口【申请】
)
const (
	GoodsJingfenQueryResponse            = "jd_union_open_goods_jingfen_query_response"            // 精选商品查询
	GoodsQueryResponse                   = "jd_union_open_goods_query_response"                    // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQueryResponse = "jd_union_open_goods_promotiongoodsinfo_query_response" // 根据skuid查询商品信息接口
	CategoryGoodsGetResponse             = "jd_union_open_category_goods_get_response"             // 商品类目查询接口
	GoodsGigfieldQueryResponse           = "jd_union_open_goods_bigfield_query_response"           // 商品详情查询接口【申请】
)
const (
	GoodsJingfenQueryResponce            = "jd_union_open_goods_jingfen_query_responce"            // 精选商品查询
	GoodsQueryResponce                   = "jd_union_open_goods_query_responce"                    // 关键词商品查询接口【申请】
	GoodsPromotiongoodsinfoQueryResponce = "jd_union_open_goods_promotiongoodsinfo_query_responce" // 根据skuid查询商品信息接口
	CategoryGoodsGetResponce             = "jd_union_open_category_goods_get_responce"             // 商品类目查询接口
	GoodsGigfieldQueryResponce           = "jd_union_open_goods_bigfield_query_responce"           // 商品详情查询接口【申请】
)

// 优惠价接口
const (
	CouponQuery = jd_method.CouponQuery // 优惠券领取情况查询接口【申请】
)
const (
	CouponQueryResponse = "jd_union_open_coupon_query_response" // 优惠券领取情况查询接口【申请】
)
const (
	CouponQueryResponce = "jd_union_open_coupon_query_responce" // 优惠券领取情况查询接口【申请】
)

// 推广接口
const (
	PromotionCommonGet       = jd_method.PromotionCommonGet       // 网站/APP获取推广链接接口
	PromotionBysubunionidGet = jd_method.PromotionBysubunionidGet // 社交媒体获取推广链接接口【申请】
	PromotionByunionidGet    = jd_method.PromotionByunionidGet    // 通过小程序获取推广链接【申请】
)
const (
	PromotionCommonGetResponse       = "jd_union_open_promotion_common_get_response"       // 网站/APP获取推广链接接口
	PromotionBysubunionidGetResponse = "jd_union_open_promotion_bysubunionid_get_response" // 社交媒体获取推广链接接口【申请】
	PromotionByunionidGetResponse    = "jd_union_open_promotion_byunionid_get_response"    // 工具商获取推广链接接口【申请】
	PromotionAppletGetResponse       = "jd_union_open_promotion_applet_get_response"       // 通过小程序获取推广链接【申请】
)
const (
	PromotionCommonGetResponce       = "jd_union_open_promotion_common_get_responce"       // 网站/APP获取推广链接接口
	PromotionBysubunionidGetResponce = "jd_union_open_promotion_bysubunionid_get_responce" // 社交媒体获取推广链接接口【申请】
	PromotionByunionidGetResponce    = "jd_union_open_promotion_byunionid_get_responce"    // 工具商获取推广链接接口【申请】
	PromotionAppletGetResponce       = "jd_union_open_promotion_applet_get_responce"       // 通过小程序获取推广链接【申请】
)

// 订单接口
const (
	OrderQuery      = jd_method.OrderQuery      // 订单查询接口
	OrderBonusQuery = jd_method.OrderBonusQuery // 奖励订单查询接口【申请】
	OrderRowQuery   = jd_method.OrderRowQuery   // 订单行查询接口
)
const (
	OrderQueryResponse      = "jd_union_open_order_query_response"       // 订单查询接口
	OrderBonusQueryResponse = "jd_union_open_order_bonus_query_response" // 奖励订单查询接口【申请】
	OrderRowQueryResponse   = "jd_union_open_order_row_query_response"   // 订单行查询接口
)
const (
	OrderQueryResponce      = "jd_union_open_order_query_responce"       // 订单查询接口
	OrderBonusQueryResponce = "jd_union_open_order_bonus_query_responce" // 奖励订单查询接口【申请】
	OrderRowQueryResponce   = "jd_union_open_order_row_query_responce"   // 订单行查询接口
)

// 推广位接口
const (
	PositionQuery  = jd_method.PositionQuery  // 查询推广位【申请】
	PositionCreate = jd_method.PositionCreate // 创建推广位【申请】
	UserPidGet     = jd_method.UserPidGet     // 获取PID【申请】
)
const (
	PositionQueryResponse  = "jd_union_open_position_query_response"  // 查询推广位【申请】
	PositionCreateResponse = "jd_union_open_position_create_response" // 创建推广位【申请】
	UserPidGetResponse     = "jd_union_open_user_pid_get_response"    // 获取PID【申请】
)
const (
	PositionQueryResponce  = "jd_union_open_position_query_responce"  // 查询推广位【申请】
	PositionCreateResponce = "jd_union_open_position_create_responce" // 创建推广位【申请】
	UserPidGetResponce     = "jd_union_open_user_pid_get_responce"    // 获取PID【申请】
)

// 营销接口
const (
	CouponGiftStop           = jd_method.CouponGiftStop           // 礼金停止
	CouponGiftGet            = jd_method.CouponGiftGet            // 礼金创建
	StatisticGiftCouponQuery = jd_method.StatisticGiftCouponQuery // 礼金效果数据
)
const (
	CouponGiftStopResponse           = "jd_union_open_coupon_gift_stop_response"            // 礼金停止
	CouponGiftGetResponse            = "jd_union_open_coupon_gift_get_response"             // 礼金创建
	StatisticGiftCouponQueryResponse = "jd_union_open_statistics_giftcoupon_query_response" // 礼金效果数据
)
const (
	CouponGiftStopResponce           = "jd_union_open_coupon_gift_stop_responce"            // 礼金停止
	CouponGiftGetResponce            = "jd_union_open_coupon_gift_get_responce"             // 礼金创建
	StatisticGiftCouponQueryResponce = "jd_union_open_statistics_giftcoupon_query_responce" // 礼金效果数据
)

// 活动接口
const (
	ActivityQuery            = jd_method.ActivityQuery            // 活动查询接口
	StatisticsRedpacketQuery = jd_method.StatisticsRedpacketQuery // 京享红包效果数据
)
const (
	ActivityQueryResponse            = "jd_union_open_activity_query_response"             // 活动查询接口
	StatisticsRedpacketQueryResponse = "jd_union_open_statistics_redpacket_query_response" // 京享红包效果数据
)
const (
	ActivityQueryResponce            = "jd_union_open_activity_query_responce"             // 活动查询接口
	StatisticsRedpacketQueryResponce = "jd_union_open_statistics_redpacket_query_responce" // 京享红包效果数据
)

const (
	UnionRootEndpoint = "https://router.jd.com/api"     // 京东联盟API路由: 参数: param_json
	JosRootEndpoint   = "https://api.jd.com/routerjson" // 宙斯API路由 🤣: 参数: 360buy_param_json
	BaseUrl           = UnionRootEndpoint
)

// 参数验证/校验等级
type VCode int

const (
	Non VCode = iota // 不校验😅
	NotEmpty
)
