package beego_player

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// HearthStoneCard is a dumb model of HearthStone card
type HearthStoneCard struct {
	Id          int
	Name        string
	Attack      int
	Life        int
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
	// set debug to on
	orm.Debug = true
}

func main() {
	o := orm.NewOrm()

	err := initRarity(o)
	fmt.Printf("Rarity initilised. err: %v\n", err)

	card := HearthStoneCard{Name: "Mal'ganis", Attack: 9, Life: 7, Description: "I AM MAL'GANIS! I AM ETERNAL!", Rarity: &CardRarity{Description: "Lengendary"}}

	// insert
	id, err := o.Insert(&card)
	fmt.Printf("[%d]: %v\n", id, err)

	// update
	card.Description = "I AM TURTLE!!!"
	num, err := o.Update(&card)
	fmt.Printf("num = %d, err: %v\n", num, err)

	// select
	dummy := HearthStoneCard{Id: card.Id}
	err = o.Read(&dummy)
	fmt.Printf("err: %v\n", err)

	// delete
	num, err = o.Delete(&dummy)
	fmt.Printf("num = %d, err: %v\n", num, err)
}

func initRarity(o orm.Ormer) error {
	
	Rarities := []CardRarity {
		CardRarity{Description: "Free"},
		CardRarity{Description: "Common"},
		CardRarity{Description: "Rare"},
		CardRarity{Description: "Epic"},
		CardRarity{Description: "Legendary"} }
	
	for _, v := range Rarities {
		_, err := o.Insert(&v)
		if err != nil {
			return err
		}
	}

	return nil
}