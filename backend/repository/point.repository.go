package repository

import (
	"github.com/tomy11/line-api/db"
	"github.com/tomy11/line-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const PointCollection = "point"

type PointRepository interface {
	Save(point *models.Point) error
	Update(point *models.Point) error
	GetById(id string) (point *models.Point, err error)
	GetAll() (point []*models.Point, err error)
	Delete(id string) error
}

type pointRepository struct {
	c *mgo.Collection
}

func NewPointRepository(conn db.Connection) PointRepository {
	return &pointRepository{conn.DB().C(PointCollection)}
}

func (r *pointRepository) Save(point *models.Point) error {
	return r.c.Insert(point)
}

func (r *pointRepository) Update(point *models.Point) error {
	return r.c.UpdateId(point.Id, point)
}

func (r *pointRepository) GetById(id string) (point *models.Point, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&point)
	return point, err
}

func (r *pointRepository) GetAll() (point []*models.Point, err error) {
	err = r.c.Find(bson.M{}).All(&point)
	return point, err
}

func (r *pointRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
