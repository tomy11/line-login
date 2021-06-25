package repository

import (
	"github.com/tomy11/line-api/db"
	"github.com/tomy11/line-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const SlipUserCollection = "slipuser"

type SlipUserRepository interface {
	Save(slip *models.SlipUser) error
	Update(slip *models.SlipUser) error
	GetById(id string) (user *models.SlipUser, err error)
	GetAll() (slip []*models.SlipUser, err error)
	Delete(id string) error
}

type slipUserRepository struct {
	c *mgo.Collection
}

func NewSlipUserRepository(conn db.Connection) SlipUserRepository {
	return &slipUserRepository{conn.DB().C(SlipUserCollection)}
}

func (r *slipUserRepository) Save(slip *models.SlipUser) error {
	return r.c.Insert(slip)
}

func (r *slipUserRepository) Update(slip *models.SlipUser) error {
	return r.c.UpdateId(slip.Id, slip)
}

func (r *slipUserRepository) GetById(id string) (slip *models.SlipUser, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&slip)
	return slip, err
}

func (r *slipUserRepository) GetAll() (slips []*models.SlipUser, err error) {
	err = r.c.Find(bson.M{}).All(&slips)
	return slips, err
}

func (r *slipUserRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
