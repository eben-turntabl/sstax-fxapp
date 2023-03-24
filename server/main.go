package main

import (
	// "net/http"

	// "github.com/gin-gonic/gin"
	"context"
	"example/server/controllers"
	"example/server/routes"
	"example/server/services"
	"example/server/utils"
	"fmt"
	"log"
	"reflect"

	"github.com/gin-contrib/cors"
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


	exchangeService services.ExchangeService
	userService	services.ClientService
	authService	services.AuthService
	orderService	services.OrderService

	AuthController controllers.AuthController
	clientController controllers.ClientController

	AuthRouteController routes.AuthRouteController
	ClientRouteController routes.ClientRouteController
	authCollection *mongo.Collection
	orderCollection *mongo.Collection
)
func init(){

serverAPI := options.ServerAPI(options.ServerAPIVersion1)

opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI).SetRegistry(bson.NewRegistryBuilder().
RegisterDecoder(reflect.TypeOf(""),utils.NullawareStrDecoder{}).
Build(),)



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

orderCollection =(client.Database("sstax_db").Collection("orders"))

nameee, errw := authCollection.Indexes().CreateOne(context.Background(),(mongo.IndexModel{Keys: bson.M{"email":1},Options: options.Index().SetUnique(true)}))	
	fmt.Println(nameee, "yolo", errw)
if errw != nil{
		fmt.Println(err.Error())
		
	}
fmt.Println("Connected to MongoDB")



userService = services.NewClientService(authCollection,ctx)
orderService = services.NewOrderService(orderCollection,ctx,authCollection)
authService = services.NewAuthService(authCollection,ctx)
exchangeService = services.NewExchangeService(authCollection,context.Background())

AuthController = controllers.NewAuthController(authService, userService)
clientController = controllers.NewClientController(userService,exchangeService,orderService)
AuthRouteController =routes.NewAuthRouteController(AuthController)
ClientRouteController=routes.NewClientRouteController(clientController)
server = gin.Default()
server.Use(cors.Default())
}
func main(){
router := server.Group("/api")


AuthRouteController.AuthRoute(router)
ClientRouteController.ClientRoute(router)
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
