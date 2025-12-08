package models

type Dimension struct {
	Wi  int `json:"wi"`
	Hi  int `json:"hi"`
	Dep int `json:"dep"`
}
type Product struct {
	PID               int    `gorm:"primaryKey;column:pid" json:"pid"`
	PName             string `gorm:"column:pname;not null" json:"pname"`
	PDescription      string `gorm:"column:pdescription" json:"pdescription"`
	PlanDetail        string `gorm:"column:plandetial" json:"plandetial"`
	WhatIsIncluded    string `gorm:"column:whatisincluded" json:"whatisincluded"`
	WhatIsNotIncluded string `gorm:"column:whatisnotincluded" json:"whatisnotincluded"`
	Price             int    `gorm:"column:price" json:"price"`

	Dem Dimension `gorm:"column:dem;type:dimension" json:"dem"`

	// Relations
	ArchiStyles     []ArchiStyle     `gorm:"foreignKey:PID" json:"archistyles"`
	Ceilings        []Ceiling        `gorm:"foreignKey:PID" json:"ceilings"`
	Garages         []Garage         `gorm:"foreignKey:PID" json:"garages"`
	RoofDetails     []RoofDetail     `gorm:"foreignKey:PID" json:"roofdetails"`
	SpecialFeatures []SpecialFeature `gorm:"foreignKey:PID" json:"specialfeatures"`
}

func (Product) TableName() string { return "product" }

type ArchiStyle struct {
	AID   int    `gorm:"primaryKey;column:aid" json:"aid"`
	PID   *int   `gorm:"column:pid" json:"pid"`
	AName string `gorm:"column:aname" json:"aname"`
}

func (ArchiStyle) TableName() string { return "archistyle" }

type Ceiling struct {
	CID   int    `gorm:"primaryKey;column:cid" json:"cid"`
	PID   *int   `gorm:"column:pid" json:"pid"`
	CName string `gorm:"column:cname;not null" json:"cname"`
	CType string `gorm:"column:ctype" json:"ctype"`
	Hi    int    `gorm:"column:hi" json:"hi"`
}

func (Ceiling) TableName() string { return "ceiling" }

type Garage struct {
	GID           int    `gorm:"primaryKey;column:gid" json:"gid"`
	PID           *int   `gorm:"column:pid" json:"pid"`
	Ty            string `gorm:"column:ty" json:"ty"`
	EntryLocation string `gorm:"column:entrylocation" json:"entrylocation"`
	GArea         int    `gorm:"column:garea" json:"garea"`
	Car           int    `gorm:"column:car" json:"car"`
}

func (Garage) TableName() string { return "garage" }

type RoofDetail struct {
	RID    int    `gorm:"primaryKey;column:rid" json:"rid"`
	PID    *int   `gorm:"column:pid" json:"pid"`
	Detail string `gorm:"column:detail" json:"detail"`
	Val    string `gorm:"column:val" json:"val"`
}

func (RoofDetail) TableName() string { return "roofdetail" }

type SpecialFeature struct {
	SID   int    `gorm:"primaryKey;column:sid" json:"sid"`
	PID   *int   `gorm:"column:pid" json:"pid"`
	SName string `gorm:"column:sname" json:"sname"`
}

func (SpecialFeature) TableName() string { return "specialfeature" }
