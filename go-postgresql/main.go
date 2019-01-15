package main

import(
  "fmt"
  "log"
  "./models"
)

func main(){

  models.TestConnection()
  
  //TestCreate()
  //TestReadAll()
  //TestUpdate()
  //TestReadById()
  TestCount()
  TestDelete()
  TestCount()
}

func TestCreate() {
  
  var product models.Product
  //product.Description = "Laranja"
  //product.Price = 2.35
  //product.Quantity = 200

  fmt.Print("Description: ")
  fmt.Scan(&product.Description)
 
  fmt.Print("Price: ")
  fmt.Scan(&product.Price)

  fmt.Print("Quantity: ")
  fmt.Scan(&product.Quantity)

  _, err := models.NewProduct(product)

  checkErr(err) 

  fmt.Println("Created successfully!")

}

func TestReadAll() {

  products, err := models.GetProducts()

  checkErr(err)

  fmt.Println("\nproducts: [")
  for _, product := range products {
  
    fmt.Printf("{\n\"id\": %d\n\"description\":%s\n\"price\": %.2f\n\"quantity\": %d\n\"createdAt\": %s\n},\n", product.Id, product.Description, product.Price, product.Quantity, product.CreatedAt)
  
  }
  
  fmt.Println("]")

}

func TestReadById() {

  var id int

  fmt.Print("\nId Product? ")
  fmt.Scan(&id)
  
  product, err := models.GetProduct(id)

  checkErr(err)

  if product.Id == 0 {
  
    fmt.Println("Produto não existe!")
    return

  }

  fmt.Println(product)

}

func TestUpdate() {
  
  var product models.Product

  fmt.Print("Description: ")
  fmt.Scan(&product.Description)
 
  fmt.Print("Price: ")
  fmt.Scan(&product.Price)

  fmt.Print("Quantity: ")
  fmt.Scan(&product.Quantity)

  fmt.Print("Id Update? ")
  fmt.Scan(&product.Id)

  rows, err := models.UpdateProduct(product)

  checkErr(err) 

  fmt.Printf("%d linhas foram modificadas", rows)

}

func TestDelete() {
  
  var id int
  fmt.Print("Delete Id? ")
  fmt.Scan(&id)

  rows, err := models.DeleteProduct(id)

  checkErr(err)

  fmt.Printf("%d linhas foram deletadas", rows)
}

func TestCount() {
  
  count, err := models.CountProducts()

  checkErr(err)

  fmt.Printf("\nTotal de produtos: %d\n", count)

}

func checkErr(err error) {

  if err != nil {
    log.Fatal(err)
  }

}
