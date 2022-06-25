package mediaorganization

import (
	"errors"
	"time"

	"github.com/jinghaijun.com/advertisement-management/db"
	"github.com/jinghaijun.com/advertisement-management/models/area"
)

type Mediaorganization struct {
	ID          int
	Mediaorname string    `json:"name"`
	Type        int       `json:"type"`
	AreaID      int       `json:"areaid"`
	Created_at  time.Time `json:"create_at"`
	Area        area.Area `json:"area"`
}

func (m *Mediaorganization) Create() error {
	db := db.Get_DB()
	request := db.Where(map[string]interface{}{"mediaorname": m.Mediaorname}).First(&m)
	if request.Error == nil {
		return errors.New("已经存在相同的")
	}
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
