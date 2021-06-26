package repository

import (
	"github.com/tomy11/line-api/db"
	"github.com/tomy11/line-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const ProductCollection = "product"

type ProductRepository interface {
	Save(slip *models.Product) error
	Update(slip *models.Product) error
	GetById(id string) (user *models.Product, err error)
	GetAll() (slip []*models.Product, err error)
	Delete(id string) error
}

type productRepository struct {
	c *mgo.Collection
}

func NewProductRepository(conn db.Connection) ProductRepository {
	return &productRepository{conn.DB().C(ProductCollection)}
}

func (r *productRepository) Save(user *models.Product) error {
	return r.c.Insert(user)
}

func (r *productRepository) Update(user *models.Product) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *productRepository) GetById(id string) (user *models.Product, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *productRepository) GetAll() (Product []*models.Product, err error) {
	err = r.c.Find(bson.M{}).All(&Product)
	return Product, err
}

func (r *productRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
