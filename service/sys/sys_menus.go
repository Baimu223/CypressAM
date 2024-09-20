package sys

import (
	"xkginweb/global"
	"xkginweb/model/entity/sys"
	"xkginweb/service/commons"
)

// 定义bbs的service提供BbsCategory的数据curd的操作

type SysMenusService struct {
	commons.BaseService[uint, sys.SysMenus]
}

/*
*
数据复制
*/
func (service *SysMenusService) CopyData(id uint) (dbData *sys.SysMenus, err error) {
	// 2: 查询id数据
	sysMenusData, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}
	// 3: 开始复制
	sysMenusData.ID = 0
	sysMenusData.Path = ""
	// 4: 保存入库
	data, err := service.Save(sysMenusData)

	return data, err
}

/*
查询父级菜单
*/
func (service *SysMenusService) FinMenusRoot() (sysMenus []*sys.SysMenus, err error) {
	err = global.KSD_DB.Where("parent_id = ? ", 0).Order("sort asc").Find(&sysMenus).Error
	return sysMenus, err
}

/*
*

  - 查询菜单形成tree数据

  - 格式： {
    id: 2,
    date: '2016-05-04',
    name: 'wangxiaohu',
    address: 'No. 189, Grove St, Los Angeles',
    },
    {
    id: 3,
    date: '2016-05-01',
    name: 'wangxiaohu',
    address: 'No. 189, Grove St, Los Angeles',
    children: [
    {
    id: 31,
    date: '2016-05-01',
    name: 'wangxiaohu',
    address: 'No. 189, Grove St, Los Angeles',
    },
    {
    id: 32,
    date: '2016-05-01',
    name: 'wangxiaohu',
    address: 'No. 189, Grove St, Los Angeles',
    },
    ],
    },
*/
func (service *SysMenusService) FinMenus(keyword string) (sysMenus []*sys.SysMenus, err error) {
	db := global.KSD_DB.Unscoped().Order("sort asc")
	if len(keyword) > 0 {
		db.Where("title like ?", "%"+keyword+"%")
	}
	err = db.Find(&sysMenus).Error
	return sysMenus, err
}

/**
*   开始把数据进行编排--递归
*   Tree(all,0)
 */
func (service *SysMenusService) Tree(allSysMenus []*sys.SysMenus, parentId uint) []*sys.SysMenus {
	var nodes []*sys.SysMenus
	for _, dbMenu := range allSysMenus {
		if dbMenu.ParentId == parentId {
			childrensMenu := service.Tree(allSysMenus, dbMenu.ID)
			if len(childrensMenu) > 0 {
				dbMenu.Children = append(dbMenu.Children, childrensMenu...)
			}
			nodes = append(nodes, dbMenu)
		}
	}
	return nodes
}

// 添加
func (service *SysMenusService) SaveSysMenus(sysMenus *sys.SysMenus) (err error) {
	err = global.KSD_DB.Create(sysMenus).Error
	return err
}

// 按照map的方式过呢更新
func (service *SysMenusService) UpdateSysMenusMap(sysMenus *sys.SysMenus, sysMenusMap *map[string]any) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysMenus).Updates(sysMenusMap).Error
	return err
}

// 删除
func (service *SysMenusService) DelSysMenusById(id uint) (err error) {
	var sysMenus sys.SysMenus
	err = global.KSD_DB.Where("id = ?", id).Delete(&sysMenus).Error
	return err
}

// 批量删除
func (service *SysMenusService) DeleteSysMenussByIds(sysMenuss []sys.SysMenus) (err error) {
	err = global.KSD_DB.Delete(&sysMenuss).Error
	return err
}

// 根据id查询信息
func (service *SysMenusService) GetSysMenusByID(id uint) (sysMenuss *sys.SysMenus, err error) {
	err = global.KSD_DB.Unscoped().Omit("created_at", "updated_at").Where("id = ?", id).First(&sysMenuss).Error
	return
}
