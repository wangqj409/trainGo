package dao

type Barticle struct {
	ID       uint   `json:"id"`
	CatId    uint   `json:"cat_id"`
	Title    string `json:"title" gorm:"size:20"`
	SubTitle string `json:"sub_title" gorm:"size:10"`
	Content  string `json:"content"`
}

func (a *Barticle) FindOne(id string) (result *Barticle) {
	Blogdb.Find(&result, id)
	return
}

func (a *Barticle) List(offset int, limit int) (results []Barticle) {
	Blogdb.Limit(limit).Offset(offset).Find(&results)
	return
}

func (bc *Barticle) Create() int64 {
	Blogdb.Create(bc)
	return Blogdb.RowsAffected
}
