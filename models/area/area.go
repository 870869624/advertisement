package area

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinghaijun.com/advertisement-management/db"
)

type Area struct {
	ID        int       `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	Pid       int       `json:"pid"`
	Left_     int       `json:"left"`
	Right_    int       `json:"right"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Area) Validate() error {
	if a.Name == "" {
		return errors.New("参数错误")
	}
	return nil
}
func (a *Area) DoesSameExists() bool {
	var count int64
	db := db.Get_DB()
	db.Table("areas").Where("name = ?", a.Name).Count(&count)
	return count > 0

}
func (a *Area) Create() error {
	var Parent Area
	if a.DoesSameExists() {
		return errors.New("已经存在相同的")
	}
	db := db.Get_DB()
	if a.Pid == 0 {
		a.Right_ = a.Right_ + 1
		fmt.Println(1111111)
		fmt.Println(a)
		result2 := db.Create(&a)
		if result2.Error != nil {
			return result2.Error
		}
		return nil
	}
	e := db.Raw("select * from areas where id = ?", a.Pid).Scan(&Parent)
	if e.Error != nil {
		return e.Error
	}
	a.Level = Parent.Level + 1
	fmt.Println(Parent)
	db.Exec("update areas set right_= right_ + 2 where right_ >= ?", Parent.Right_)
	db.Exec("update areas set left_ = left_ + 2 where left_ >= ?", Parent.Right_)
	a.Left_ = Parent.Right_
	a.Right_ = a.Left_ + 1
	result2 := db.Create(&a)
	if result2.Error != nil {
		return result2.Error
	}
	return nil
}

type Request struct {
	Left  int //左区间
	Right int //右区间
	Level int //行政等级
}
type Order struct {
	By []string
}
type AreaQuery struct {
	Request
	Order
}
type AreaResponse struct {
	Request
	Result []Area
}

func List(query *AreaQuery) (*AreaResponse, error) {
	reponse := &AreaResponse{
		Request: query.Request,
	}
	db := db.Get_DB()
	request := db.Table("areas")
	if query.Left != 0 {
		request = request.Where("left_ >= ?", query.Left)
	}
	if query.Right != 0 {
		request = request.Where("right_ <= ?", query.Right)
	}
	if query.Level != 0 {
		request = request.Where("level = ?", query.Level)
	}
	for _, v := range query.By {
		request = request.Order(v)
	}
	r := request.Find(&reponse.Result)
	return reponse, r.Error
}

func (a *Area) Delete() error {
	if !a.DoesSameExists() {
		return errors.New("不存在这条记录")
	}
	fmt.Println(a.Name)
	a.Validate()
	db := db.Get_DB()
	e := db.Exec("delete from areas where left_ >= ? and right_ <= ?", a.Left_, a.Right_)
	if e.Error != nil {
		return e.Error
	}
	newnumbber := a.Right_ - a.Left_ + 1
	db.Exec("update areas set right_= right_ - ? where right_ > ?", newnumbber, a.Right_)
	db.Exec("update areas set left_ = left_ - ? where left_ > ?", newnumbber, a.Right_)
	return nil
}

type New struct {
	ID        int       `json:"id" gorm:"primarykey"`
	NewName   string    `json:"name"`
	NewLevel  int       `json:"level"`
	NewPid    int       `json:"pid"`
	NewLeft_  int       `json:"left"`
	NewRight_ int       `json:"right"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *New) CHange() error {
	var (
		Parent   Area
		DataNow  Area
		DataPast Area
		qqqq     Area
		aaaa     []Area
	)
	db := db.Get_DB()
	e := db.Table("areas").Where("id = ?", n.NewPid).First(&Parent)
	if e.Error != nil {
		return errors.New("父级不存在！")
	}
	e1 := db.Table("areas").Where("id = ?", n.ID).First(&DataPast)
	if e1.Error != nil {
		return errors.New("不存在这条记录")
	}
	if DataPast.Level == 1 {
		db.Exec("update areas set level = level + ? where left_ >= ? and right_ <= ?", Parent.Level, DataPast.Left_, DataPast.Right_)
	} else {
		db.Exec("update areas set level = level + ? where left_ >= ? and right_ <= ?", Parent.Level-1, DataPast.Left_, DataPast.Right_)
	}

	//计算差值
	c := DataPast.Right_ - DataPast.Left_
	a := DataPast.Right_ - DataPast.Left_ + 1
	// d := DataPast.Right_ - DataPast.Left_ + 2
	// //修改需要的那一条记录
	// fmt.Println(a)
	if DataPast.Left_ > Parent.Right_ {
		//如果从后往前移动
		db.Exec("update areas set left_ = left_ + ?  where left_ > ?", a, Parent.Right_)
		db.Exec("update areas set right_ = right_ + ? where right_ >?", a, Parent.Right_)
		db.Table("areas").Where("id = ?", n.ID).First(&DataNow)
		e := DataNow.Left_ - Parent.Right_
		db.Exec("update areas set left_ =  ? where id = ?", Parent.Right_, n.ID)
		db.Exec("update areas set right_ = ? where id = ?", Parent.Right_+c, n.ID)
		db.Exec("update areas set right_ = ? where id = ?", Parent.Right_+a, n.NewPid)
		//往前移动需要的全部区间
		db.Exec("update areas set left_ = left_ - ? where left_ > ? and right_ < ?", e, DataNow.Left_, DataNow.Right_)
		db.Exec("update areas set right_ = right_ - ? where right_ > ? and right_ < ?", e, DataNow.Left_, DataNow.Right_)
		db.Exec("update areas set left_ = left_ - ?  where left_ > ?", a, DataNow.Right_)
		db.Exec("update areas set right_ = right_ - ? where right_ >?", a, DataNow.Right_)
	}

	if DataPast.Left_ < Parent.Right_ {
		// if DataPast.Level == 1 {
		// 	db.Exec("update areas set ")
		// }
		//如果前往后移动
		db.Exec("update areas set right_ = right_ + ? where right_ > ? and left_ < ?", a, Parent.Right_, Parent.Right_)
		db.Exec("update areas set left_ = left_ + ? where left_ > ?", a, Parent.Right_)
		db.Exec("update areas set right_ = right_ + ? where left_ > ?", a, Parent.Right_)
		db.Raw("select * from areas").Scan(&aaaa)
		fmt.Println(aaaa)
		//扩展需要移动的当前级
		db.Exec("update areas set left_ =  ? where id = ?", Parent.Right_, n.ID)
		db.Exec("update areas set right_ = ? where id = ?", Parent.Right_+c, n.ID)
		db.Table("areas").Where("id = ?", n.ID).First(&DataNow)
		b := DataNow.Left_ - DataPast.Left_
		fmt.Println(DataPast, DataNow)
		//更新父级右值
		db.Exec("update areas set right_ = ? where id = ?", Parent.Right_+a, n.NewPid)
		// //修改记录内部的记录
		db.Exec("update areas set left_ = left_ + ? where left_ > ? and right_ < ?", b, DataPast.Left_, DataPast.Right_)
		db.Exec("update areas set right_ = right_ + ? where left_ > ? and right_ < ?", b, DataPast.Left_, DataPast.Right_)
		db.Table("areas").Where("id = ?", 1).First(&qqqq)
		fmt.Println(qqqq)
		db.Exec("update areas set left_ = left_ - ?  where left_ >= ?", a, DataPast.Right_)
		db.Exec("update areas set right_ = right_ - ? where right_ >= ?", a, DataPast.Right_)
	}
	return nil
}
