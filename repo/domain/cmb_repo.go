
package domain

import (
	"errors"
	"fmt"
	"github.com/gogap/logrus"
	"mafool.com/go/deeplink/config"
	"mafool.com/go/deeplink/repo"
	"mafool.com/go/deeplink/repo/models"
)

// SAAS模式，根据租户操作不同的数据库
type CmbRepository interface {
	ScanSerialByPage(tablename string, offset, limit int) (devices []models.TbDevice, err error)
	//UploadSerials(tablename int64, serials []string) error
	InsertHitSerials(serials []models.TbHit) error
	SyncStatusReset() error
	InsertSyncLog(log *models.LogSync) error
	GetAllDeviceTable() ([]string, error)
}

type CmbRepo struct {
}

func CmpDefault() *CmbRepo {
	return &CmbRepo{}
}

func (p *CmbRepo) ScanSerialByPage(tablename string, offset, limit int) (devices []models.TbDevice, err error) {
	sql := fmt.Sprintf("select id,serial from %s limit %d,%d", tablename, offset, limit)
	if err := repo.DB.SQL(sql).Find(&devices); err != nil {
		return nil, err
	}
	return
}

func (p *CmbRepo) InsertHitSerials(hits []models.TbHit) (err error) {
	if len(hits) == 0 {
		return errors.New("设备号不能为空")
	}
	if _, err := repo.DB.Insert(hits); err != nil {
		return err
	}
	return nil

	//var sql = fmt.Sprintf("INSERT INTO `tb_hit`(id,`serial`) VALUES ", channelId)
	//var values string
	//for _, v := range hits {
	//	values = fmt.Sprintf(`%s,(next value for MYCATSEQ_HIT,'%s')`, values, v) //Hit自增变量
	//}
	//values = strings.TrimLeft(values, ",")
	//sql = fmt.Sprintf("%s %s", sql, values)
	//_, err = repo.DB.Exec(sql)
	//return err
}

//func (p *CmbRepo) UploadSerials(tablename string, serials []string) (err error) {
//	if len(serials) == 0 {
//		return errors.New("设备号不能为空")
//	}
//	var sql = fmt.Sprintf("/*!mycat: schema=DEEPLINK_%d*/ INSERT INTO `tb_device`(`id`, `serial`) VALUES ", channelId)
//	var values string
//	for _, v := range serials {
//		//values = fmt.Sprintf(`%s,(next value for MYCATSEQ_LOG,'%s')`, values, v)
//		values = fmt.Sprintf(`%s,(next value for MYCATSEQ_DEEPLINK_%d,'%s')`, values, channelId, v)
//	}
//	values = strings.TrimLeft(values, ",")
//	sql = fmt.Sprintf("%s %s", sql, values)
//	_, err = repo.DB.Exec(sql)
//	return err
//}

// 同步之前将状态归零
func (p *CmbRepo) SyncStatusReset() error {
	if _, err := repo.DB.Exec("truncate tb_hit;"); err != nil {
		return err
	}
	return nil
}

func (p *CmbRepo) InsertSyncLog(log *models.LogSync) error {
	if _, err := repo.DB.Insert(log); err != nil {
		return err
	}
	return nil
}

func (p *CmbRepo) GetAllDeviceTable() (ts []string, err error) {
	sql := fmt.Sprintf("select table_name from information_schema.tables where table_schema=%q and table_name like 'tb_device%%';", config.Mysql.Schema)
	rt, err:= repo.DB.SQL(sql).Query()
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	for _, value := range rt {
		ts=append(ts,string(value["table_name"]))
	}
	return ts, nil
}

//

//命中数据
//explain
///*!mycat: schema=DEEPLINK_1000020*/ select * from tb_hit;
///*!mycat: schema=DEEPLINK_1000020*/ INSERT INTO `tb_hit`( id,`uid`, `serial`) VALUES (next value for MYCATSEQ_GLOBAL,1000017, 'twretet');
///*!mycat: schema=DEEPLINK_1000020*/ truncate tb_hit;
