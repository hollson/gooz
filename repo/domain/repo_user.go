// -------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.cn>
// @ Date: 2020-05-07
// @ Version: 1.0.0
//
// 产品数据仓储
// -------------------------------------------------------------------------------------

package domain

// type IUser interface {
// 	GetUser(id int64) (*models.User, error)
// 	GetUserByPage(offset, limit int) ([]models.User, error)
// }
//
// type UserRepo struct{}
//
// // 查询单条记录
// func (p *UserRepo) GetUser(id int64) (ret *models.User, err error) {
// 	u := new(models.User)
// 	logrus.Infoln("id=", id)
// 	if has, err := db.ID(id).Get(u); err != nil {
// 		logrus.Errorln(has, err) // todo warp包装
// 		return nil, err
// 	} else {
// 		return u, nil
// 	}
// }
//
// // 分页查询
// // Deprecated:启用
// func (p *UserRepo) GetUserByPage(offset, limit int) ([]models.User, error) {
// 	var us []models.User
// 	if err := db.Where("1=1").Limit(limit, offset).Find(us); err != nil {
// 		logrus.Errorln("GetUserByPage Err：", err)
// 		return nil, err
// 	}
// 	return us, nil
// }
//
// // Join查询
//
// // 分库查询
// // explain
// // /*!mycat: schema=DEEPLINK_1000020*/ select * from tb_hit;
// // /*!mycat: schema=DEEPLINK_1000020*/ INSERT INTO `tb_hit`( id,`uid`, `serial`) VALUES (next value for MYCATSEQ_GLOBAL,1000017, 'twretet');
// // /*!mycat: schema=DEEPLINK_1000020*/ truncate tb_hit;
