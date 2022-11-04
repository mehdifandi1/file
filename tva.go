package main

import "fmt"

func main() {
	const taux tva int = 0.2
	var pu,remise,totalht,montanttva,totalttc float32
	var qte int
 fmt.Println("prix unitair de l article ?")
 fmt.Scanln(&pu)
fmt.Print("quantite de l article ?")
fmt.Scanln(&qte)
fmt.Println("taux de la remise ?")
fmt.Scanln(&tauxremise)
total=pu*float32((qte))
remise=total*float32(tauxremise)/100
totalht=total-remise
montanttva=totalht*float32(tauxtva)/100
totalttc=totalht+montanttva
fmt.Println("total avant remise :",total)
fmt.Println("total apres remise :",totalht)
fmt.Println("montant de la tva :",montanttva)
fmt.Println("total ttc :",total)

}