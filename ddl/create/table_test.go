package create

import (
	"log"
	"testing"

	"github.com/Snehal1112/QueryBuilder/constrain"
	"github.com/Snehal1112/QueryBuilder/datatype"
)

func TestNewTable(t *testing.T) {
	createCategories := NewTable("categories", nil)
	createCategories.Field("categoryId", datatype.INT,50, []int{constrain.NOTNULL, constrain.AI, constrain.PK})
	createCategories.Field("categoryName", datatype.VARCHAR, 225, []int{})

	log.Println(createCategories.prepareQuery())

	creatProducts := NewTable("products", nil)
	creatProducts.Field("productId", datatype.INT, 50, []int{constrain.AI, constrain.PK})
	creatProducts.Field("productName", datatype.VARCHAR, 225, []int{constrain.NOTNULL})
	creatProducts.Field("categoryId", datatype.INT, 50, []int{})
	creatProducts.NewForeignKeyConstrain("fk_category", "categoryId", "categories")
	creatProducts.SetForeignKey(constrain.Cascade, constrain.Cascade)

	log.Println(creatProducts.prepareQuery())
}
