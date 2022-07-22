package bd

import (
	"context"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
   MongoCN es el objeto de coneccion a la base de datos.
 */
var MongoCN = conectarBD()
var clientOptions = optinos.Client().ApplyURI("mongodb+srv://twittor:twittor2022@program.nr0u1.mongodb.net/test")

/*
   ConectarBD es la funcion que me permite conectar con la BD
 */
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

	log.Printfln("Conexion Exitosa con la DB")
	return client
}

/*
   ChequeoConnection es el Ping a la base de datos
 */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
