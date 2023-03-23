package main

import (
	// "net/http"

	// "github.com/gin-gonic/gin"
	"context"
	"example/server/controllers"
	"example/server/routes"
	"example/server/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "errors"
)

const dbUri ="mongodb://mongoadmin:password@localhost:27017"
var(
	server *gin.Engine
	ctx	context.Context

	userService	services.ClientService
	authService	services.AuthService
	AuthController controllers.AuthController
	AuthRouteController routes.AuthRouteController
	authCollection *mongo.Collection
)
func init(){

serverAPI := options.ServerAPI(options.ServerAPIVersion1)

opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI)

client,err := mongo.Connect(context.Background(),opts)

if err != nil {
	panic(err)
}
// defer func() {
// 	if err = client.Disconnect(context.TODO());err != nil{
// 		panic(err)
// 	}
// }()


if err := client.Ping (ctx , readpref.Primary()); err != nil{ 
	panic(err)
}

authCollection =(client.Database("sstax_db").Collection("users"))
nameee, errw := authCollection.Indexes().CreateOne(context.Background(),(mongo.IndexModel{Keys: bson.M{"email":1},Options: options.Index().SetUnique(true)}))	
	fmt.Println(nameee, "yolo", errw)
if errw != nil{
		fmt.Println(err.Error())
		
	}
fmt.Println("Connected to MongoDB")



userService = services.NewClientServiceImpl(authCollection,ctx)
authService = services.NewAuthService(authCollection,ctx)
AuthController = controllers.NewAuthController(authService, userService)
AuthRouteController =routes.NewAuthRouteController(AuthController)
server = gin.Default()
}
func main(){
router := server.Group("/api")


AuthRouteController.AuthRoute(router)
log.Fatal(server.Run(":8080"))

}









































































// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run()
// }
