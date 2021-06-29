package Model

type T1 struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func (b *T1) TableName() string {
	return "t1"
}
