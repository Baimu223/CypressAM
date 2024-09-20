package commons

type ParentModel[D any] struct {
	ID D `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
}
