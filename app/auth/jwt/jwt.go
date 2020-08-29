package jwt

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hollson/gooz/repo"
	"github.com/sirupsen/logrus"
)

const (
	VERPROFIX = "jwt_" // version前缀
)

// todo 配置
var (
	JWTSECRET  = "7KCQC8YW3AMF7A98OIUV2Q6MEKE3GTGSYB3RFWI89KS071URZ2MJDPZKTJIUVZUY"
	TokenTerm  = time.Hour * 24
	VersionCtl = false // 开启版本控制
)

type Claims struct {
	jwt.StandardClaims // 标准Claim
	Version int64      // 服务端版本号
}

// 生成JWT令牌：
// uuid ：用户唯一标识，
func Generate(uuid string) (string, error) {
	var token *jwt.Token

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  uuid,                             // 签发对象
			ExpiresAt: time.Now().Add(TokenTerm).Unix(), // 过期时间

			// 为了精简化token，可省略其他缺省项
			// Id:        util.EncodeMD5(fmt.Sprintf("%d+%s",time.Now().UnixNano(),uuid)), // JWT标识
			// Subject:   "gooz_service",                            // 应用主题
			// Issuer:    "gooz.org",                                // 签发机构
			// IssuedAt:  jwt.TimeFunc().Unix(),                         // 签发时间
			// NotBefore: time.Now().Unix(),                             // 生效时间
		},
		Version: time.Now().Unix(),
	}
	if VersionCtl {

	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(JWTSECRET))
	return tokenStr, err
}

// 将Token字符串解析成Claim对象
func Resolve(token string) (*Claims, error) {
	token = strings.TrimPrefix(token, "Bearer ")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSECRET), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token数据错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token尚未生效")
			} else {
				return nil, errors.New("token验证失败")
			}
		}
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			if VersionCtl {
				v, err := repo.GetValue(claims.Audience)
				if err != nil {
					VersionCtl = false
					logrus.Errorln(err)
					return nil, err
				}

				ver := strconv.Itoa(int(claims.Version))
				if v != ver {
					return nil, errors.New("token已失效")
				}
			}
			return claims, nil
		}
	}
	return nil, err
}

// Token版本升级
// 用于账号退出、强制下线、安全检测等场景，即在服务端升级token版本号，
// 使得token强制失效。该功能需要在服务端开启Redis或Memory缓存服务，
// 用于存储版本号，并在token验证服务中添加对版本号的验证。
func upgrade() {

}

// 验证本版号的一致性
func verifyVersion(uuid string) (bool, error) {
	if VersionCtl {
		v, err := repo.GetValue(uuid)
		if err != nil {
			VersionCtl = false
			logrus.Errorln(err)
			return false, err
		}
		return v == uuid, nil
	}
	return false, errors.New("version ctl is closed")
}

// 滑动延时
// func (j *JWT) RefreshToken(tokenString string) (string, error) {
// 	jwt.TimeFunc = func() time.Time {
// 		return time.Unix(0, 0)
// 	}
// 	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return j.SigningKey, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
// 		jwt.TimeFunc = time.Now
// 		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
// 		return j.CreateToken(*claims)
// 	}
// 	return "", TokenInvalid
// }
