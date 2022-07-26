package bd

import (
	"context"
	"fmt"
	"github.com/jjvos19/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*
   BuscoPerfil busca el perfil en la BD.
*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	db := MongoCN.Database("twittors")
	col := db.Collection("usuarios")
	fmt.Println(cancel)

	var perfil models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
