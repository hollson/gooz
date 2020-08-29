[TOC]

# 橱窗DIY-Api文档　v2.2




# 1. 我的DIY


## 1.1. 橱窗列表
**说明：** 主态的列表包含素材或商品，根据`IsGoods`字段决定详情接口

**请求:**

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/diy/list/my`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:200701150318519815
resTags:119
pageNum:1    //分页编号(从1开始)
pageSize:10  //分页尺寸
otherAccount:
searchType:0
```
_字段参考说明：_
```go
account       //账号id
resTags       //资源标签，全部0、 上装110、下装111、鞋112、表情109、舞蹈127、表情泡129
pageNum      //分页编号(从1开始)
pageSize     //分页尺寸
```

Demo:

```html
POST /api/v2/diy/list/res?appKey=100000&timestamp=666&nonce=1234&sign=ABC HTTP/1.1
Host: x.x.x.x:19722
Content-Type: application/x-www-form-urlencoded
Authorization: Bearer xxxxxx

account=200701150318519815&resTags=119&pageIndex=0&pageNum=20&otherAccount=200701150318519815
```



**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "list": [
            {
                "id": 74867,      // 素材或商品ID
                "name": "0b0ca59d0c",
                "priceType": 0,    
                "price": 0,
                "thumbName": "T1DabTB4WT1RCvBVdK",
                "onsale": false,    // 是否在售
                "sales": 0,        // 销量
                "isGoods": false,   // 是否为商品，根据该字段决定调用素材或商品详情接口
                "createTime": "1970-01-01"
            }
        ],
        "count": 15
    }
}
  ```



---



## 1.2. DIY详情
### a.  根据素材编号查询

**请求:**

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/diy/detail/byres`
- Content-Type: `application/x-www-form-urlencoded`

```conf
account:200701150318519815
id:453891
```
demo:

```js
POST /api/v2/diy/detail/bygoods?appKey=100000&timestamp=666&nonce=1234&sign=ABC HTTP/1.1
Host: x.x.x.x:19722
Content-Type: application/x-www-form-urlencoded
Authorization: Bearer xxxxxx

account=200701150318519815&id=453891
```



**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "resGuid": 22701,
        "resourceName": "惊讶",
        "thumbName": "T1vy_TBXAT1RCvBVdK",
        "fans": 0,
        "stars": 0,
        "state": 2,
        "restype": 17,
        "price": 0,
        "machine": 0,
        "authGuid": "29989886",
        "authorName": "ReWorld",
        "authorCover": "T1FyKTByE_1RCvBVdK",
        "timestamp": "2019-10-15",
        "resContext": "",
        "bought": 0,
        "myCollect": 0,
        "canDiscuss": false,
        "official": 1,
        "goodsId": 449728,
        "editTimestamp": "2020-05-27",
        "resTags": 109,
        "FileName": "T1oy_TB_ZT1RCvBVdK",
        "sales": 69,
        "discuss": 4,
        "isDiscuss": 0,
        "isMy": 0,
        "giveLike": 0,
        "giveLikeCount": 4,
        "giveNoLikeCount": 0,
        "isHot": 0,
        "isCart": 0,
        "AddCartTime": "",
        "storeGuid": 1,
        "priceType": 1,
        "isTrial": 0,
        "comId": 449728,
        "haveOther": 0
    }
}
```
_字段参考说明：_
```go
`json:"resGuid"`         //资源guid
`json:"resourceName"`    //资源名字
`json:"thumbName"`       //缩略图名字
`json:"fans"`            //收藏的次数
`json:"stars"`           //星级评分
`json:"state"`           //资源状态，等于2则是在售状态
`json:"restype"`         //资源类型
`json:"price"`           //资源价格
`json:"machine"`         //适用机型
`json:"authorGuid"`      //作者ID
`json:"authorName"`      //发布者
`json:"authorCover"`     //作者头像
`json:"timestamp"`       //资源发布时间
`json:"resContext"`      //资源描述
`json:"bought"`          //我是否购买了该资源(0 没有购买过, 1 购买过)
`json:"myCollect"`       //我是否收藏了该资源(0,没有收藏. 1,收藏了)
`json:"canDiscuss"`      //是否可以评论
`json:"official"`        //是否官方数据
`json:"goodsId"`         //商品guid
`json:"editTimestamp"`   //资源修改时间
`json:"resTags"`         //资源标签
`json:"FileName"`        //文件名字
`json:"sales"`           //购买人数
`json:"discuss"`         //评价人数
`json:"isDiscuss"`       //是否评论过 0 为未评论 1 为评论过
`json:"isMy"`            //是否是自己的 0 为不是 1 为是
`json:"giveLike"`        //我是否点赞了该资源(0 没有点赞过, 1 点赞过)
`json:"giveLikeCount"`   //点赞人数
`json:"giveNoLikeCount"` //不点赞人数
`json:"isHot"`           //是否是热门 1 是
`json:"isCart"`          //是否加入到了购物车0 否 1 是
`json:"AddCartTime"`     //添加到购物车的时间
`json:"storeGuid"`       //资源所属商店guid
`json:"priceType"`       //价格类型 0 worny币 1 魔方
`json:"isTrial"`         //是否试用
`json:"comId"`           //id
`json:"haveOther"`       //是否有别人的资源, 0 没有; 1 有
```

### b.  根据商品编号查询
**请求:**

- Method：Post
- URI：`http://x.x.x.x:19722/api/v2/diy/detail/bygoods`
- Content-Type: `application/x-www-form-urlencoded`

> 请求和响应参数同　[a.根据资源编号查询](### a.  根据资源编号查询)。



### c.  官方素材详情

**请求:**

- Method：Post
- URI：`http://x.x.x.x:19722/api/v2/diy/detail/plat`
- Content-Type: `application/x-www-form-urlencoded`

> 请求和响应参数同　[a.根据资源编号查询](### a.  根据资源编号查询)。



---




## 1.3. 删除DIY
**请求:**
- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/delete`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_
```json
account:200701150318519815
token:
resid:123
```
**响应**:
- Content-Type: `application/json`

_响应参数示例：_
```json
{
     "code": "1", 		
     "message": "OK",　
}
```

---



## 1.4. 发布Step.1－创建资源
**请求:**
- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/pub/create`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_
```json
account:xxxxxx
resName:new004
fileContext:描述信息
resTags:1059
isShow:true
haveOther:0
onsale:1    // 0.关闭销售，1.打开销售
price:22    // 价格
```
_服务端参考字段：_
```go
`json:"account" form:"account"`
`json:"resName" form:"resName"`         //素材名称
`json:"fileContext" form:"fileContext"` //素材描述
`json:"fType" form:"fType"`             //素材类型
`json:"resTags" form:"resTags"`         //标签
`json:"isShow" form:"isShow"`           // 显示
`json:"haveOther" form:"haveOther"`     //有没有包含他人的资源

`json:"onsale" form:"onsale"` // 0.关闭销售，1.打开销售
`json:"price" form:"price"`
```
_素材类型：_
```shell
enum ResourceType {
	None 			= 	0;  //错误类型
	GameObject 		= 	1;  //对象类型
	Mesh 			= 	2;  //Mesh
	Texture2D 		= 	3;  //2D图片
	String 			= 	4;  //字符串
	Avatar 			= 	5;  //角色
	AudioClip 		= 	6;  //音效
	MeshPart 		= 	7;  //部件
	Model 			= 	8;  //模型
	AnimationClip   = 	9;  //动作
    Particle 		= 	10; //特效
	Accessory		=	12; //配饰
	Tool			=	13; //工具
	Shirt 			= 	14; //衣服
    Pants 			= 	15; /裤子
    Shoes 			= 	16; //鞋
    Face 			= 	17; //脸部
	DanceAnimation  =   18; //舞蹈动作
	FrameEmoji		=   19; //泡泡
    ModelParticle 	= 	20; //组合特效
}
```

**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 0,
    "message": "",　
     data:{
   		"resName": "abc", 		
   		"resGuid": 111,
   		"goodsGuid": 222　　　
	}
}
```



## 1.6. 发布Step.2－上传资源

**说明:** 上传原图与缩略图（视频无须上传缩略图）

**请求:**

- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/pub/upload`
- Content-Type: `application/data`

_请求参数示例：_
```json
{
     "account": "", 		
     "resguid": 0,
     "resName": "",
     "fileStr": "base64码",  //兼容文件流
     "isThumb":0  // 是否是缩略图,0：原图，1：缩略图
}
```

_参考：_
```shell
curl --location --request POST 'http://x.x.x.x:19722/api/v2/diy/pub/upload?appKey=100000&timestamp=666&nonce=1234&sign=ABC' \
--form 'account=200701150318519815' \
--form 'resGuid=3384932042801411' \
--form 'resName=zzzzz22222' \
--form 'isThumb=1' \
--form 'uploadfile=@/D:/工作/本地配置.md'
```



**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 0,
    "message": "",　
     data:{
   		"fileguid": 0, 		
   		"fileName": "",
        "resName": "",
        "size": 0
	}
}
```



## 1.7. 发布Step.3－完成创建

**说明**：通知服务器：完成了`素材创建`与`文件上传`，请将素材置为可用状态。

**请求:**

- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/pub/done`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:xxxxxx      // 账号
resGuid:123456      // 素材编号
actionId:111		// 小K动作ID
```
**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 1,
    "message": "OK"
}
```



## 1.8. 更新资源

**请求:**
- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/pub/update`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_
```json
account:200701150318519815
resguid:4118175956336899
resName:333
resContext:shssgdgdsgdccc
resTags:119
```
```shell
string `json:"account" form:"account"`       // 账号
int64  `json:"resguid" form:"resguid"`       // 资源编号
string `json:"resName" form:"resName"`       // 素材名称
string `json:"resContext" form:"resContext"` // 素材描述
int32  `json:"resTags" form:"resTags"`       // 标签
bool   `json:"onSale" form:"onSale"`         // 是否上下架
int32  `json:"price" form:"price"`           // 价格
int32  `json:"currency" form:"currency"`     // 币种(略)
```
**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 0,
    "message": ""
}
```



## 1.9. 是否启动引导页

**请求:**

- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/guide`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_
```json
{
     "account": "xxx"
}
```
**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "guide": false,  // 是否加载引导页
        "staged": false  // 是否有小K暂存数据
    }
}
```




## 1.10. 获取所有标签

**请求:**

- Method：GET/POST
- URI：`http://x.x.x.x:19722/api/v2/res/tags`
- Content-Type: `application/x-www-form-urlencoded`

_请求示例：_
```shell
curl --location --request GET 'http://x.x.x.x:19722/api/v2/res/tags?appKey=100000&timestamp=666&nonce=1234&sign=ABC'
```

---

---





# 2.客态DIY

## 2.1. DIY列表

**说明：** 客态的列表均为商品

**请求:**

- Method：Post
- URI：`http://x.x.x.x:19722/api/v2/diy/list/commodity`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:200701150318519815
resTags:119
pageNum:1    //分页编号(从1开始)
pageSize:10  //分页尺寸
```
_字段参考说明：_
```go
account       //账号id
resTags       //资源标签，全部0、 上装110、下装111、鞋112、表情109、舞蹈127、表情泡129
```

Demo:

```html
POST /api/v2/diy/list/commodity?appKey=100000&timestamp=666&nonce=1234&sign=ABC HTTP/1.1
Host: x.x.x.x:19722
Content-Type: application/x-www-form-urlencoded
Authorization: Bearer xxxxxx

account=200701150318519815&resTags=119&pageIndex=0&pageNum=5
```



**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "list": [
            {
                "id": 435802,
                "name": "CS_dao guang3",
                "priceType": 1,
                "Price": 0,
                "thumbName": "T1waYTB5CT1RCvBVdK",
                "onsale": true,
                "sales": 133,
                "createTime": "2019-06-01"
            }
        ],
        "count": 1298
    }
}
  ```



## 2.2. DIY详情

> 同我的DIY详情




## 2.4. 穿戴列表

**请求:**

- Method：Post
- URI：`http://x.x.x.x:19722/api/v2/diy/player/suit`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_
```json
{
     "account": ""
}
```
**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "nickName": "shs1-6",
        "figure": "T1dtbTBKJv1RCvBVdK",
        "figureUrl": "http://192.168.39.8:19334/file/get?filename=T1dtbTBKJv1RCvBVdK",
        "suits": [{
            resId：123
            officialType：2  // 2表示官方素材(官方素材详情，调`1.2.c`，非官方调`1.2.a`)
            thumbFileName：
        }]
    }
}

```



## 2.5. 魔方币购买

**请求:**

- Method：POST
- URI：`http://x.x.x.x:19722/api/v2/diy/buy`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_
```json
account:200701150318519815
gid:123
```
**响应**:

- Content-Type: `application/json`

_响应参数示例：_
```json
{
    "code": 0,
    "message": "",
    "data": {
        "goodsGuid": 0,
        "resName": "",	
        "userDiamond": 0
    }
}
```

---


# ３. 素材审核

## ３.1. 审核列表
**请求:**

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/diy/check/list`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:200701150318519815
checkState:0
pageNum:1    //分页编号(从1开始)
pageSize:10  //分页尺寸
```

_审核状态：_

	0  //审核失败
	1  //审核通过
	2  //拒绝驳回
**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "list": [
            {
                "resGuid": 10459795546964227,
                "resName": "xxx0079933wwww",
                "resType": 18,
                "fileName": "",
                "status": 0,　　／／素材状态
                "resTags": 127,
                "resTagsTier": "5-127",
                "resTagsTierTitle": "动作-舞蹈动作",
                "thumbName": "http://192.168.39.8:19334/file/get?filename=T1da_TBQKT1RCvBVdK",
                "createTime": "1970-01-01",
                "editTime": "2020-08-06",
                "checkState": 0　　／／审核状态
            }
        ],
        "count": 33
    }
}
  ```
## ３.２.审核详情
**请求:**

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/diy/check/detail`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```shell
account:200701150318519815
id:10305319364198659  //素材ID
```

```shell
＃请求示例：
curl --location --request POST 'http://192.168.39.223:19722/api/v2/diy/check/detail?appKey=100000&timestamp=666&nonce=1234&sign=ABC' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'account=200701150318519815' \
--data-urlencode 'id=10305319364198659'
```

**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "resGuid": 14643371221975299,
        "resourceName": "自动售卖-008",
        "thumbName": "http://192.168.39.8:19334/file/get?filename=T1atbTBgLT1RCvBVdK",
        "state": 0,
        "restype": 0,
        "price": 200,
        "authorGuid": "200701150318519815",
        "authorName": "End",
        "authorCover": "",
        "timestamp": "2020-08-20",
        "resContext": "自动售卖-008",
        "canDiscuss": false,
        "official": 0,
        "editTimestamp": "2020-08-20",
        "resTags": 1059,
        "resTagsTier": "4-112-1059",
        "resTagsTierTitle": "化装-鞋子-皮鞋",
        "FileName": "http://192.168.39.8:19334/file/get?filename=T15y_TBshv1RCvBVdK",
        "AddCartTime": "",
        "storeGuid": 0,
        "priceType": 0,
        "haveOther": 0,
        "checkState": 0,
        "checkMark": "不想填审核说明20200820201533",
        "deleted": false,
        "rejectReason": "[20,212]",
        "rejectReasonTitles": [
            "存在垃圾内容",
            "内购项的介绍存在欺骗"
        ],
        "autoSale": 0
    }
}
  ```
# 4. 小K动捕服务

## 4.1. 获取暂存列表
**请求:**

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/xk/stage`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:123456  // OpenId
pageNum:1       //分页编号(从1开始)
pageSize:10     //分页尺寸
```

**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "count": 4,
        "list": [
            {
                "action_id": 863,
                "file_id": "filename",
                "timestamp": 1598343511,
                "code": 1,
                "msg": "form data",
                "process_level": 2,
                "video_url": "http://video_url.com",
                "action_url": "http://action_url.com",
                "action_args": "200701150318519815",
                "thumb": "T1dy_TB4xg1RCvBVdK",
                "duration": "2.5",
                "thumbUrl": "http://192...1dy_TB4xg1RCvBVdK"
            }
        ]
    }
}
  ```
_响应参数参考说明：_  
```json
string `json:"game_key,omitempty"` // 游戏对应的key
int64  `json:"action_id"`          // 动作⽂文件id,上传视频成功返回的action_id
string `json:"file_id"`            // 游戏平台对应的⽂文件id
int64  `json:"timestamp"`          // 时间戳（毫秒）
int32  `json:"code"`      // 处理理状态码， = 0 正常； <0 玩家上传视频内容敏敏感（涉及多种情况）；>
string `json:"msg"`              // 描述字符串串，succeed 、 error
int32  `json:"process_level"`     // 处理理等级，玩家上传视频动捕效果评级：1-优； 2-良； 3-中；4-差
string `json:"video_url"`         // 合成视频的下载地址
string `json:"action_url"`        // ⽣生成的动作⽂文件下载地址
string `json:"action_args"`       // 穿透参数,(指向Account)
string `json:"thumb" form:"thumb"`  	  //缩略图
string `json:"thumbUrl" form:"thumbUrl"`  //缩略图连接
string `json:"duration" form:"duration"`  //时长
```

## 4.2. 上传缩略图
**请求:**    (`参考1.6 上传资源`)

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/xk/upload`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:200701150318519815
actionId:1234        // 动作⽂文件id
thumb:iVBORw0KGgoAAAANSUh...4M7v9oU9FUHwx6psUxL+3gCAo9Q/bPr    //资源文件
```

**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 1,
    "message": "OK",
    "data": {
        "fileName": "T1dy_TB4xg1RCvBVdK",
        "size": 2145
    }
}
  ```

## 4.3. 删除暂存数据

**请求:**    (`参考1.6 上传资源`)

- Method：Post 
- URI：`http://x.x.x.x:19722/api/v2/xk/dalete`
- Content-Type: `application/x-www-form-urlencoded`

_请求参数示例：_

```json
account:200701150318519815
actionId:1234        // 动作⽂文件id
```

**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "code": 0,
    "message": ""
} 
  ```
## 4.4. 小K回调(第三方接口)
**请求:**

- Method：Post 
- URI：`http://10.0.0.16:19722/api/xk/callback`
- Content-Type: `application/json`

_请求参数示例：_

```shell
{
"game_key":"reworld",
"action_id":{{$randomInt}},
"file_id":"filename",
"timestamp":{{$timestamp}},
"code":1,
"msg":"hello world1",
"process_level":2,
"video_url":"http://video_url.com",
"action_url":"http://action_url.com",
"action_args":"{\"openId\":\"123456\",\"duration\":\"2.5\"}"
}
```

**响应**:

- Content-Type: `application/json`

_响应参数示例：_

  ```json
{
    "ret": 0
}
  ```



# 附A. 文档说明

## A.1. `AppClient` 授权验证

>  适用于服务端(即服务提供方)主动授权的`App应用客户端`安全验证。服务端会向应用客户端授权`AppKey`和`AppSecret`信息。C端客户的安全由`App应用`代理完成。

| App应用 | 认证方式  | AppKey | AppSecret       |
| ------- | --------- | ------ | --------------- |
| 安卓    | AppClient | 100011 | KSJ0&LHGUWN***  |
| IOS     | AppClient | 100012 | 8JHT6HNJ5BF4*** |
| 编辑器  | AppClient | 100013 | WLPOUN2H58P***  |
| Ｗeb    | WebClient |        | Token           |

**安全验证内容：**

1. 身份验证（√）：即验证是授权客户端的可靠访问。
2. 参数校验（×）：即验证调用参数不被劫持和篡改。非开放应用，此验证可完全信任。

3. 防重放攻击（×）：即防止黑客恶意重复提交攻击。建议赋值（注意不要在提交操作时动态生成，一般为页面全局变量）。

4. 非对称验证（×）：即主调方和被调方互交换非对称公钥，实现完全非信任的加密互调。



**App客户端调用示例:**  

```json
POST /api/v2/diy/demo?appKey=100000&timestamp=666&nonce=1234&sign=ABC HTTP/1.1
Content-Type: application/x-www-form-urlencoded
Accept: */*
Cache-Control: no-cache
Host: x.x.x.x:19722
Accept-Encoding: gzip, deflate, br
Content-Length: 20
Connection: keep-alive

userName=jack&age=22

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 20 Jul 2020 06:50:44 GMT
Content-Length: 62

{"code":1,"message":"OK","data":{"userName":"jack","age":22}}
```



**签名算法：**

 sign=`SHA256(appKey.timestamp.nonce.securet)`

> 如：appKey=ABC，timestamp=123456，nonce=99999，securet=XYZ，则sign=SHA256(ABC.123456.99999.XYZ)



## A.2. WebClient：授权验证

> 适用于C端客户自授权的安全验证，需要客户在浏览器端密码登录，并获取JWT的Token信息。可参考[JWT授权验证](xxxxxx)。



---



# 附：版本记录

**2020.08.27**
1. 废弃原来 `/api/v2/diy/mycount`接口，增加  `接口1.9`引导页接口（详见具体章节）
2. 废弃原`/api/v2/diy/pub/onsale(上架商品)`接口，将售卖参数(`onsale`和`price`）移到`接口1.4`中。
3. 添加`接口1.7`, 以通知当前素材创建完成。
4. 列表页面添加了`详情页导航辅助字段"auxDetailLink"`(与原功能兼容)。详见接口`1.1`,`2.1`,`2.4`,`3.1`。
5. 列表分页参数进行了规范化命名和处理。详见接口`1.1`,`2.1`,`3.1`,`4.1`。



**2020.08.26**

1. 增加小K数据暂存服务接口

2. `接口2.4`穿戴列表根据素材来源（官方平台和非官方平台）分别调用详情

3. `接口1.7`增加`actionId`字段，即`接口4.1`中的动作⽂文件id

   










