package commons

import (
	"fmt"
	"gorm.io/gorm"
	"xkginweb/global"
)

// 父类
type BaseService[D any, T any] struct{}

// 明细
func (service *BaseService[D, T]) GetByID(id uint) (data *T, err error) {
	p := ParentModel[D]{}
	d := p.ID
	fmt.Println(d)
	err = global.KSD_DB.Where("id = ?", id).First(&data).Error
	return
}

// 无逻辑条件
func (service *BaseService[D, T]) UnGetByID(id D) (data *T, err error) {
	err = global.KSD_DB.Unscoped().Where("id = ?", id).First(&data).Error
	return
}

// 删除
func (service *BaseService[D, T]) DeleteByID(id D) (bool, int64) {
	var data T
	rows := global.KSD_DB.Where("id = ?", id).Delete(&data).RowsAffected
	return rows > 0, rows
}

// 忽略逻辑删除
func (service *BaseService[D, T]) UnDeleteByID(id D) (bool, int64) {
	var data T
	rows := global.KSD_DB.Unscoped().Where("id = ?", id).Delete(&data).RowsAffected
	return rows > 0, rows
}

// 保存和更新
func (service *BaseService[D, T]) Save(data *T) (dbData *T, err error) {
	err = global.KSD_DB.Create(&data).Error
	return data, err
}

// 批量保存
func (service *BaseService[D, T]) SaveBatch(datas []T) (bool, int64) {
	affected := global.KSD_DB.Create(&datas).RowsAffected
	return affected > 0, affected
}

// 更新忽略物理条件
func (service *BaseService[D, T]) UnUpdateByID(data T) (dbData *T, err error) {
	err = global.KSD_DB.Unscoped().Model(dbData).Updates(&data).Error
	return &data, err
}

// 更新逻辑更新
func (service *BaseService[D, T]) UpdateByID(data T) (dbData *T, err error) {
	err = global.KSD_DB.Model(dbData).Updates(&data).Error
	return &data, err
}

// 更新逻辑更新
//func (service *BaseService[D, T]) UpdateMap(data T) (dbData *T, err error) {
//	err = global.KSD_DB.Model(dbData).Updates(&data).Error
//	return &data, err
//}

// 状态更新
func (service *BaseService[D, T]) UpdateStatus(id D, field string, fieldValue any) (bool, int64) {
	var data T
	affected := global.KSD_DB.Model(data).Where("id = ?", id).Update(field, fieldValue).RowsAffected
	return affected > 0, affected
}

// 状态更新-忽略isdelete
func (service *BaseService[D, T]) UnUpdateStatus(id D, field string, fieldValue any) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ?", id).Update(field, fieldValue).RowsAffected
	return affected > 0, affected
}

// 自增
func (service *BaseService[D, T]) IncrById(id D, field string) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" >= 0", id).Update(field, gorm.Expr(field+" + 1")).RowsAffected
	return affected > 0, affected
}

// 指定步长自增
func (service *BaseService[D, T]) IncrByIdNum(id D, field string, fieldValue int) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" >= 0", id).Update(field, gorm.Expr(field+" + ?", fieldValue)).RowsAffected
	return affected > 0, affected
}

// 自减
func (service *BaseService[D, T]) DecrById(id D, field string) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" > 0", id).Update(field, gorm.Expr(field+" - 1")).RowsAffected
	return affected > 0, affected
}

// 指定步长自减
func (service *BaseService[D, T]) DecrByIdNum(id D, field string, fieldValue int) (bool, int64) {
	var data T
	affected := global.KSD_DB.Unscoped().Model(data).Where("id = ? and "+field+" > 0", id).Update(field, gorm.Expr(field+" - ?", fieldValue)).RowsAffected
	return affected > 0, affected
}

// 批量自增 + 1
func (service *BaseService[D, T]) Incrs(ids []D, field string) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" >= 0", ids).
		Update(field, gorm.Expr(field+" + 1")).RowsAffected
	return affected > 0, affected
}

// 批量自减 + 1
func (service *BaseService[D, T]) Decrs(ids []D, field string) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" > 0", ids).
		Update(field, gorm.Expr(field+" - 1")).RowsAffected
	return affected > 0, affected
}

// 批量自增 + num
func (service *BaseService[D, T]) IncrsByNum(ids []D, field string, num int) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" >= 0", ids).
		Update(field, gorm.Expr(field+" + ?", num)).RowsAffected
	return affected > 0, affected
}

// 批量自减 - num
func (service *BaseService[D, T]) DecrsByNum(ids []D, field string, num int) (bool, int64) {
	var model T
	affected := global.KSD_DB.Unscoped().Model(&model).
		Where("id in ? and "+field+" > 0", ids).
		Update(field, gorm.Expr(field+" - ?", num)).RowsAffected
	return affected > 0, affected
}

// count
// count(Where.gt("age",30)----age > 30)
// list
// list(where)
// page
// page(where)
// update(where)
// delete(where)
