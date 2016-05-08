package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type MyCollection struct {
	*mgo.Collection
}

func (mc *MyCollection) findAllGoods() (response []goods) {
	err := mc.Find(bson.M{}).Sort("price").Sort("collection").All(&response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response", len(response), " doc")
	return response
}

func (mc *MyCollection) findAllGoodsByCollection(goodsCollection string) (response []goods) {
	mc.EnsureIndexKey("collection")
	err := mc.Find(bson.M{"collection": goodsCollection}).Sort("name").All(&response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response", len(response), " doc")
	return response
}

func (mc *MyCollection) findOneGoodsById(goodsId string) (response goods) {
	mc.EnsureIndexKey("goodsId")
	err := mc.FindId(bson.ObjectIdHex(goodsId)).One(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

type goods struct {
	Id         bson.ObjectId "_id"
	Collection string
	Name       string
	Code       string
	Price      float32
	Desc       string
	Detail     string
	ShipDetail string
	Option     []option
	Voted      int
	Imgsrc     string
}

type option struct {
	Name   string
	Price  float32
	Choice []int
}

func main() {

	// init
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Static("/js", "./js")
	router.Static("/css", "./css")
	router.LoadHTMLGlob("html/*")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Comming Soon",
		})
	})

	router.GET("/findAll", func(c *gin.Context) {
		//size := c.DefaultQuery("size", ".")

		session, err := mgo.Dial("128.199.245.83")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		prod := session.DB("prod")
		prod.Login("admin", "Ball0931372529")
		mc := &MyCollection{prod.C("goods")}
		c.JSON(http.StatusOK, mc.findAllGoods())

	})

	router.GET("/findAllGoodsByCollection/:goodsCollection", func(c *gin.Context) {
		goodsCollection := c.Param("goodsCollection")

		session, err := mgo.Dial("128.199.245.83")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		prod := session.DB("prod")
		prod.Login("admin", "Ball0931372529")
		mc := &MyCollection{prod.C("goods")}
		c.JSON(http.StatusOK, mc.findAllGoodsByCollection(goodsCollection))

	})

	router.GET("/findOneGoodsById/:goodsId", func(c *gin.Context) {
		goodsId := c.Param("goodsId")
		session, err := mgo.Dial("128.199.245.83")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		prod := session.DB("prod")
		prod.Login("admin", "Ball0931372529")
		mc := &MyCollection{prod.C("goods")}
		c.JSON(http.StatusOK, mc.findOneGoodsById(goodsId))

	})
	router.Run(":8081")

}

//db.C(Collection).Find(Query).Limit(limit).Sort(sortBy...)
