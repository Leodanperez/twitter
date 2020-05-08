package bd

import (
  "context"
  "log"
   
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion al BD */
var MongoCN=ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://LeoPerez:LeoperezSofware2020@cluster0-uec1w.mongodb.net/test?retryWrites=true&w=majority")
//var clientOptions = options.Client().ApplyURI("mongodb+srv://root:EntrenamientoTwitter@twitter-kb2wp.mongodb.net/test?retryWrites=true&w=majority")


/*ConectarBD es la funcion que me permite conectar la BD*/
func ConectarBD() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
      log.Fatal(err.Error())
      return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
      log.Fatal(err.Error())
      return client
    }
    log.Println("Conexion Exitosa con la BD")
    return client
}

/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
  err := MongoCN.Ping(context.TODO(), nil)
  if err != nil {
    return 0
  }
  return 1
}