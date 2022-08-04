package models

/*
   Tweet, captura del Body, el mensaje que llega.
*/
type Tweet struct {
	Mensaje string  `bson:"mensaje" json:"mensaje"`
}
