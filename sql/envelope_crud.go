package sql

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Envelope struct {
	ID         int64 `json:"envelope_id"`
	UID        int64 `json:"uid"`
	Opened     bool  `json:"opened"`
	Value      int64 `json:"value"`
	SnatchTime int64 `json:"snatch_time"`
}

type PublicOpenedEnvelope struct {
	*Envelope           // 匿名嵌套
	UID       *struct{} `json:"uid,omitempty"`
}

type PublicClosedEnvelope struct {
	*Envelope           // 匿名嵌套
	UID       *struct{} `json:"uid,omitempty"`
	Value     *struct{} `json:"value,omitempty"`
}

func (Envelope) TableName() string {
	return "red_envelope"
}

func GetAllEnvelopesByUID(uid int64) ([]*Envelope, error) {
	db, err := gorm.Open("mysql", "root:zyz123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	var envelopes []*Envelope
	conditions := map[string]interface{}{
		"uid": uid,
	}
	if err := db.Table(Envelope{}.TableName()).Where(conditions).Find(&envelopes).Error; err != nil {
		return nil, err
	}
	return envelopes, nil
}

func GetEnvelopeByEevelopeID(envelope_id int64) (envelope Envelope) {
	db, err := gorm.Open("mysql", "root:zyz123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.Where("id = ?", envelope_id).First(&envelope)
	return
}

func CreateEnvelope(user User) (envelope Envelope) {
	db, err := gorm.Open("mysql", "root:zyz123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	snatch_time := time.Now().UnixNano()
	var a int64 = 10
	envelope = Envelope{UID: user.ID, Opened: false, Value: a, SnatchTime: snatch_time}
	db.Create(&envelope)
	return envelope
}

func UpdateState(envelope_id int64) (envelope Envelope) {

	db, err := gorm.Open("mysql", "root:zyz123456@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	//查询条件
	envelope.ID = envelope_id
	db.Model(&envelope).Update("opened", true)
	return
}
