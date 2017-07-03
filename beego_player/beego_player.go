package beego_player

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// HearthStoneCard is a dumb model of HearthStone card
type HearthStoneCard struct {
	Id          int
	Name        string
	Attack      int
	Health      int
	Description string
	Rarity		*CardRarity `orm:"rel(one)"`
}

// CardRarity is a const dumb model for rarity :
// Free, Common, Rare, Epic, Legendary
type CardRarity struct {
	Id 			int
	Description string
}

var alias = "default"

func init() {
	// set default database
	orm.RegisterDataBase(alias, "mysql", "root:zxasqwxcsdwe@tcp(localhost:3377)/go_db?charset=utf8", 30)
	// regisiter defined model
	orm.RegisterModel(new(HearthStoneCard), new(CardRarity))
	// create table
	orm.RunSyncdb(alias, false, true)
}