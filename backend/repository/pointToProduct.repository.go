package repository

import (
	"github.com/tomy11/line-api/db"
	"github.com/tomy11/line-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const PointToProductCollection = "pointToProduct"

type PointToProductRepository interface {
	Save(ptpo *models.PointToProduct) error
	Update(ptpo *models.PointToProduct) error
	GetById(id string) (ptpo *models.PointToProduct, err error)
	GetAll() (ptpo []*models.PointToProduct, err error)
	Delete(id string) error
}

type pointToProductRepository struct {
	c *mgo.Collection
}

func NewPointToProductRepository(conn db.Connection) PointToProductRepository {
	return &pointToProductRepository{conn.DB().C(PointToProductCollection)}
}

func (r *pointToProductRepository) Save(ptpo *models.PointToProduct) error {
	return r.c.Insert(ptpo)
}

func (r *pointToProductRepository) Update(ptpo *models.PointToProduct) error {
	return r.c.UpdateId(ptpo.Id, ptpo)
}

func (r *pointToProductRepository) GetById(id string) (ptpo *models.PointToProduct, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&ptpo)
	return ptpo, err
}

func (r *pointToProductRepository) GetAll() (ptpo []*models.PointToProduct, err error) {
	err = r.c.Find(bson.M{}).All(&ptpo)
	return ptpo, err
}

func (r *pointToProductRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
