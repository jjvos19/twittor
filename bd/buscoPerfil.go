package bd

import (
	"context"
	"fmt"
	"github.com/jjvos19/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"log"
)

/*
   BuscoPerfil busca el perfil en la BD.
*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()
	
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	
	var perfil models.Usuario

	objID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
	
		log.Print("-----> error al convertir el ID " + err.Error())
		//fmt.Println("No se puede convertir el ID " + opcion)
		return perfil, err
	}
	
	log.Print("ID: " + ID + "\n objID: ")
	log.Println(objID)
	condicion := bson.M{
		"_id": objID,
	}

	err = col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
