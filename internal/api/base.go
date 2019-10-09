package api

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/xiaomeng79/istio-micro/cinit"
	"github.com/xiaomeng79/istio-micro/internal/utils"

	"github.com/asaskevich/govalidator"
	jsoniter "github.com/json-iterator/go"
	"github.com/xiaomeng79/go-utils/crypto"
	"github.com/xiaomeng79/go-utils/time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ReqParam struct {
	AppID     int        `json:"-"`                                                    // AppID
	AppKey    string     `json:"app_key" valid:"required~appKey必须存在"`                  // 密钥ID
	AppSecret string     `json:"app_secret" valid:"-"`                                 // 密钥
	RequestID string     `json:"request_id" valid:"required~required必须存在"`             // 32位的唯一请求标识，用于问题排查和防止重复提交
	Timestamp string     `json:"timestamp" valid:"required~毫秒时间戳必须存在"`                 // 毫秒时间戳
	Custom    string     `json:"custom" valid:"-"`                                     // 第三方自定义内容
	Nonce     string     `json:"nonce" valid:"required~随机数必须存在,length(8|8)~随机数必须满足8位"` // 8 位随机数
	Language  string     `json:"language" valid:"in(cn|en)"`                           // 多语言支持
	Sign      string     `json:"sign" valid:"required~签名必须存在"`                         // 签名
	SignType  string     `json:"sign_type" valid:"required~签名类型必须存在"`                  // 签名类型：MD5 SHA_1 SHA_256 SHA_512
	Encode    bool       `json:"encode" valid:"-"`                                     // 响应数据data是否进行base64编码，默认true
	Data      string     `json:"data" valid:"-"`                                       // 请求的数据
	Remark    string     `json:"remark"`
	Code      string     `json:"code"`
	Message   string     `json:"message"`
	Page      utils.Page `json:"page"`    // 分页
	IsPage    bool       `json:"is_page"` // 是否分页
}

/**
解析参
*/

func (r *ReqParam) Decode(s string) error {
	return json.Unmarshal([]byte(s), r)
}

/**
验证参数
*/

func (r *ReqParam) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	if err != nil {
		return err
	}
	err = r.ValidateAppSecret()
	return err
}

/**
验证appSecret
*/

func (r *ReqParam) ValidateAppSecret() error {
	if r.AppKey != cinit.Config.Service.AppKey {
		return errors.New("appKey不存在")
	}
	r.AppSecret = cinit.Config.Service.AppSecret
	return nil
}

/**
**解析data参数
**input: v point
**output:  error
 */
func (r *ReqParam) DataDecode(v interface{}) error {
	if r.Data == "" {
		return nil
	}
	// 判断参数是否base64编码
	var decoded []byte
	var err error
	if r.Encode { // 解码
		decoded, err = base64.StdEncoding.DecodeString(r.Data)
		if err != nil {
			return errors.New("(base64解析失败):" + err.Error())
		}
	} else {
		decoded = []byte(r.Data)
	}
	err = json.Unmarshal(decoded, v)
	if err != nil {
		return errors.New("(json解析失败):" + err.Error())
	}

	return nil
}

/**
**编码data参数
**input: v interface
**output:  error
 */
func (r *ReqParam) DataEncode(v interface{}) error {
	if v == "" {
		r.Data = ""
		return nil
	}
	var encoded []byte
	var err error
	// json marshal
	encoded, err = json.Marshal(v)
	if err != nil {
		return errors.New("json编码失败")
	}

	// 判断参数是否base64编码
	if r.Encode { // 解码
		r.Data = base64.StdEncoding.EncodeToString(encoded)
	} else {
		r.Data = string(encoded)
	}

	return nil
}

// 生成签名 string
func (r *ReqParam) CreateSign() (string, error) {
	_signType := strings.ToUpper(strings.Trim(r.SignType, " "))
	var originSign string
	// 组合签名字符串
	if r.Data == "" {
		originSign = r.AppKey + r.AppSecret + r.Nonce + r.Timestamp
	} else {
		originSign = r.AppKey + r.AppSecret + r.Data + r.Nonce + r.Timestamp
	}
	var _sign string
	var err error
	switch _signType {
	case "MD5":
		_sign = crypto.MD5(originSign)
	case "SHA_1":
		_sign = crypto.SHA1(originSign)
	case "SHA_256":
		_sign = crypto.SHA256(originSign)
	case "SHA_512":
		_sign = crypto.SHA512(originSign)
	default:
		err = errors.New("签名类型不存在")
	}
	if err != nil {
		return "", err
	}
	return _sign, nil
}

/**
根据签名类型比较签名 true:相同 false：不同
*/
func (r *ReqParam) CompareSign() (bool, error) {
	_sign, err := r.CreateSign()
	if err != nil {
		return false, err
	}
	if strings.EqualFold(r.Sign, _sign) {
		return true, nil
	}
	return false, errors.New("签名不匹配")
}

/**
生成返回数据
*/
func (r *ReqParam) genData(code ErrorNo, errmsg string, v interface{}) (interface{}, error) {
	var err error
	err = r.DataEncode(v)
	if err != nil {
		return nil, err
	}
	r.Timestamp = time.GenMicTime()
	r.Sign, err = r.CreateSign()
	if err != nil {
		return nil, err
	}
	_v := map[string]interface{}{
		"code":       code.String(),
		"message":    ReturnMsg[code] + errmsg,
		"appKey":     r.AppKey,
		"request_id": r.RequestID,
		"timestamp":  r.Timestamp,
		"custom":     r.Custom,
		"sign":       r.Sign,
		"sign_type":  r.SignType,
		"encode":     r.Encode,
		"data":       r.Data,
	}
	if r.IsPage {
		_v["page"] = r.Page
	}
	return _v, err
}

// 返回数据
/*
0.生成data数据和page数据
1.encode data数据
2.生成时间错和签名
3.生成错误吗和错误信息
4.返回数据
*/
func (r *ReqParam) R(v interface{}) (interface{}, error) {
	return r.genData(Success, "", v)
}
