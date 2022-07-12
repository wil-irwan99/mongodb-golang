package controller

import (
	"golang-mongodb/dto"
	"golang-mongodb/model"
	"golang-mongodb/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	router              *gin.Engine
	productRegisUseCase usecase.ProductRegistrationUseCase
	findAllPrdtUseCase  usecase.FindAllProductUseCase
	findPrdtByID        usecase.FindProductByIdUseCase
	findPrdtByCat       usecase.FindProductByCategoryUseCase
	updtPrdtById        usecase.ProductUpdateUseCase
	dltPrdtbyId         usecase.ProductDeleteUseCase
}

func (pc *ProductController) registerNewProduct(ctx *gin.Context) {
	var newProduct model.Product
	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = pc.productRegisUseCase.Register(&newProduct)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    newProduct,
	})

}

func (pc *ProductController) findAllProduct(ctx *gin.Context) {
	var pageData dto.Paging
	err := ctx.ShouldBindJSON(&pageData)
	if err != nil {
		log.Println(err.Error())
		return
	}

	convertPage, _ := strconv.Atoi(pageData.Page)
	convertLimit, _ := strconv.Atoi(pageData.Limit)

	result, err := pc.findAllPrdtUseCase.FindAll(int64(convertPage), int64(convertLimit))
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    result,
	})
}

func (pc *ProductController) findProductById(ctx *gin.Context) {
	var idInput dto.IdInput
	err := ctx.ShouldBindJSON(&idInput)
	if err != nil {
		log.Println(err.Error())
		return
	}

	objID, err := primitive.ObjectIDFromHex(idInput.Id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	result, err := pc.findPrdtByID.FindPrdtById(objID)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    result,
	})

}

func (pc *ProductController) findProductByCategory(ctx *gin.Context) {
	var catInput dto.CategoryInput
	err := ctx.ShouldBindJSON(&catInput)
	if err != nil {
		log.Println(err.Error())
		return
	}

	result, err := pc.findPrdtByCat.FindPrdtByCat(catInput.CatInput)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    result,
	})

}

func (pc *ProductController) updateProductById(ctx *gin.Context) {
	var dataInput dto.UpdateForm
	err := ctx.ShouldBindJSON(&dataInput)
	if err != nil {
		log.Println(err.Error())
		return
	}

	objID, err := primitive.ObjectIDFromHex(dataInput.Id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	value := dataInput.Value["value"]

	result, err := pc.updtPrdtById.UpdateProduct(objID, bson.D{{dataInput.Table, value}})
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":            "SUCCESS",
		"total data updated": result,
	})

}

func (pc *ProductController) deleteProductById(ctx *gin.Context) {
	var idInput dto.IdInput
	err := ctx.ShouldBindJSON(&idInput)
	if err != nil {
		log.Println(err.Error())
		return
	}

	objID, err := primitive.ObjectIDFromHex(idInput.Id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	result, err := pc.dltPrdtbyId.DeleteProduct(objID)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "SUCCESS",
		"data":    result,
	})

}

func NewProductController(router *gin.Engine, productRegisUseCase usecase.ProductRegistrationUseCase, findAllPrdtUseCase usecase.FindAllProductUseCase, findPrdtByID usecase.FindProductByIdUseCase, findPrdtByCat usecase.FindProductByCategoryUseCase, updtPrdtById usecase.ProductUpdateUseCase, dltPrdtbyId usecase.ProductDeleteUseCase) *ProductController {
	controller := ProductController{
		router:              router,
		productRegisUseCase: productRegisUseCase,
		findAllPrdtUseCase:  findAllPrdtUseCase,
		findPrdtByID:        findPrdtByID,
		findPrdtByCat:       findPrdtByCat,
		updtPrdtById:        updtPrdtById,
		dltPrdtbyId:         dltPrdtbyId,
	}

	rPrdt := router.Group("/product")
	rPrdt.POST("/add", controller.registerNewProduct)
	rPrdt.GET("/all", controller.findAllProduct)
	rPrdt.GET("/find_id", controller.findProductById)
	rPrdt.GET("/find_cat", controller.findProductByCategory)
	rPrdt.PUT("/update", controller.updateProductById)
	rPrdt.DELETE("/delete", controller.deleteProductById)

	return &controller
}
