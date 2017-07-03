package beego_player

import (
	"testing"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func TestBeegoPlayer(t *testing.T) {
	o := orm.NewOrm()

	err := initRarity(o)
	if err != nil {
		t.Error(err)
	}

	card := HearthStoneCard{Name: "Mal'ganis", Attack: 9, Health: 7, Description: "I AM MAL'GANIS! I AM ETERNAL!", Rarity: &CardRarity{Description: "Lengendary"}}

	// insert
	_, err = o.Insert(&card)
	if err != nil {
		t.Error(err)
	}

	// update
	card.Description = "I AM TURTLE!!!"
	num, err := o.Update(&card)
	if err != nil || num != 1 {
		t.Error(err)
	}

	printCards(o);

	// select
	dummy := HearthStoneCard{Id: card.Id}
	err = o.Read(&dummy)
	if err != nil {
		t.Error(err)
	}

	// delete
	num, err = o.Delete(&dummy)
	if err != nil || num != 1 {
		t.Error(err)
	}
}

func initRarity(o orm.Ormer) error {
	
	Rarities := []CardRarity {
		CardRarity{Description: "Free"},
		CardRarity{Description: "Common"},
		CardRarity{Description: "Rare"},
		CardRarity{Description: "Epic"},
		CardRarity{Description: "Legendary"} }
	
	qs := o.QueryTable(new(CardRarity))

	count, err := qs.Count()
	if err != nil || count != 0 {
		return err
	}

	for _, v := range Rarities {
		_, err := o.Insert(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func printCards(o orm.Ormer) error {
	var cards []HearthStoneCard
	o.QueryTable(new(HearthStoneCard)).All(&cards)

	for _, v := range cards {
		fmt.Printf("ID:%d, Name:%v, Attack:%d, Health:%d\n[%v]:%v\n", v.Id, v.Name, v.Attack, v.Health, v.Rarity.Description, v.Description)
	}

	return nil
}