package main

import "fmt"

func main(){
	fmt.Println("------------------------ EXERCISE 1 ------------------------")
  slc1 := []string{"Hello", "Hola", "नमस्कार ", "こんにちは", "Привіт"}
  
  sub1 := slc1[:3]
  sub2 := slc1[1:1]
  sub3 := slc1[3:]

  fmt.Printf("Slice original: %s \n", slc1)
  fmt.Printf("Sub-slice con los primeros dos índices: %s \n", sub1)
  fmt.Printf("Sub-slice con los tres índices de la mitad: %s \n", sub2)
  fmt.Printf("Sub-slice con los últimos dos índices: %s \n", sub3)
   
  fmt.Println("------------------------ EXERCISE 2 ------------------------")
  message := "Hi 👨 and 👩"
  
  r1 := message[:2]
  r2 := message[3:8]
  r4 := message[12:]

  var slc []byte
  slc = append(slc, message[8])
  slc = append(slc, message[9])
  slc = append(slc, message[10])

  r3 := string(slc)

  fmt.Printf("Mensaje original: %s \n", message)
  fmt.Printf("Primer caracter: %s \n", r1)
  fmt.Printf("Segundo caracter: %s \n", r2)
  fmt.Printf("Tercer caracter: %v \n", r3)
  fmt.Printf("Cuarto caracter: %s \n", r4)
  
	fmt.Println("------------------------ EXERCISE 3 ------------------------")
  type Teammate struct{
    firstName string
    lastName string
    id int
  }

  Nekros1x := Teammate{
    "Sergio",
    "Cabrera",
    0x29A9C02EB105,
  }

  Chavsi := Teammate{
    firstName: "Mijael",
    lastName: "Mardonez",
    id: 0x3FA905C483A2D,
  }

  var WatoDev Teammate

  WatoDev.firstName = "Miguel"
  WatoDev.lastName = "Lopéz"
  WatoDev.id = 0x30A908AE0D2

  fmt.Printf("Primera instancia: %+v \n", Nekros1x)
  fmt.Printf("Segunda instancia: %+v \n", Chavsi)
  fmt.Printf("Tercera instancia: %+v \n", WatoDev)
}
