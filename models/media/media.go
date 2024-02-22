package media

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinghaijun.com/advertisement-management/db"
)

type mediatype int

const (
	全部类型 mediatype = 0
	电视   mediatype = 1
	广播   mediatype = 2
	报纸杂志 mediatype = 3
)

type Media struct {
	ID                int       `gorm:"primarykey"`
	Medianame         string    `json:"medianame" gorm:"index;comment:媒体名称"`
	Mediatype         int       `json:"mediatype" gorm:"index;comment:媒体类型"`
	Testor            string    `json:"testor" gorm:"index;comment:检测机构"`
	Jurisdiction      string    `json:"jurisdiction" gorm:"comment:管辖机构"`
	Division          string    `json:"division" gorm:"comment:行政区域"`
	Mediaorganization string    `json:"mediaorganization" gorm:"comment:媒介机构"`
	Distributor       string    `json:"distributor" gorm:"comment:派发人"`
	Deadline          string    `json:"deadline" gorm:"comment:截止时间"`
	CreatedAt         time.Time `jaon:"created_at"`
}

// 检验数据完整
func (m *Media) Validate() error {
	if m.Medianame == "" || m.Testor == "" || m.Jurisdiction == "" || m.Division == "" ||
		m.Mediaorganization == "" || m.Deadline == "" {
		return errors.New("参数错误")
	}
	return nil
}
func (media *Media) Create() error {
	if e := media.Validate(); e != nil {
		return e
	}
	connection := db.Get_DB()
	request := connection.Where(map[string]interface{}{"medianame": media.Medianame, "mediatype": media.Mediatype}).First(&media)
	fmt.Println(request.RowsAffected)
	if request.Error == nil {
		return errors.New("已经存在相同的")
	}
	result := connection.Create(media)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (media *Media) Update() error {
	var ad Media
	if e := media.Validate(); e != nil {
		return e
	}
	connection := db.Get_DB()
	find := connection.First(&ad)
	if find.Error != nil {
		return find.Error
	}
	result := find.Updates(media)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type Pagination struct {
	Size    int `json:"size"`
	Current int `json:"current"`
}
type Order struct {
	Order_by []string
}
type Query struct {
	Pagination
	Order
	Title     string
	Mediatype int
}
type Response struct {
	Pagination
	Result []Media
}

func List(query *Query) (*Response, error) {
	resopnse := &Response{
		Pagination: query.Pagination,
	}
	db := db.Get_DB()
	result := db.Table("media")
	if query.Title != "" {
		result = result.Where("medianame like ?", fmt.Sprintf("%%%s%%", query.Title))
	}
	if query.Mediatype != 0 {
		result = result.Where("mediatype = ?", query.Mediatype)
	}
	offset := query.Size * (query.Current - 1)
	result = result.Offset(offset).Limit(query.Size)
	for _, v := range query.Order_by {
		result = result.Order(v)
	}
	e := result.Find(&resopnse.Result)
	if e.Error != nil {
		return nil, e.Error
	}
	return resopnse, nil
}
