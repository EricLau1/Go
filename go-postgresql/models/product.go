package models

type Product struct {
  Id int
  Description string
  Price float64
  Quantity int
  CreatedAt string
}

func NewProduct(product Product) (bool, error) {

  con := Connect()

  sql := "insert into products (description, price, quantity) values ($1, $2, $3)"

  stmt, err := con.Prepare(sql)

  if err != nil {
    
    return false, err

  }

  _, err = stmt.Exec(product.Description, product.Price, product.Quantity)

  if err != nil {
  
    return false, err

  }
  

  defer stmt.Close()
  defer con.Close()

  return true, nil

}

func GetProducts() ([]Product, error) {

  con := Connect()

  sql := "select * from products"

  rs, err := con.Query(sql)

  if err != nil {
  
    return nil, err

  }

  var products []Product

  for rs.Next() {

    var product Product
  
    err := rs.Scan(&product.Id, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)

    if err != nil {
      return nil, err
    }

    products = append(products, product)
  }

  defer rs.Close()
  defer con.Close()

  return products, nil
}

func GetProduct(id int)(Product, error) {

  con := Connect()

  sql := "select * from products where id = $1"

  rs, err := con.Query(sql, id)

  if err != nil {
    
    return Product{}, err 
  
  }

  var product Product

  if rs.Next() {
  
    err := rs.Scan(&product.Id, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)

    if err != nil {
      
      return Product{}, err
    
    }
  }

  defer con.Close()
  defer rs.Close()

  return product, nil
}

func UpdateProduct(product Product) (int64, error ) {

  con := Connect()

  sql := "update products set description = $1, price = $2, quantity = $3 where id = $4"

  stmt, err := con.Prepare(sql)

  if err != nil {

    return 0, err
  
  }

  rs, err := stmt.Exec(product.Description, product.Price, product.Quantity, product.Id)
  
  if err != nil {

    return 0, err
  
  }

  rows, err := rs.RowsAffected()
  
  if err != nil {

    return 0, err
  
  }

  defer stmt.Close()
  defer con.Close()

  return rows, nil

}

func DeleteProduct(id int) (int64, error) {

  con := Connect()

  sql := "delete from products where id = $1"

  stmt, err := con.Prepare(sql)

  if err != nil {
  
    return 0, err

  }

  rs, err := stmt.Exec(id)
  
  if err != nil {
  
    return 0, err

  }

  rows, err := rs.RowsAffected()
  
  if err != nil {
  
    return 0, err

  }
  
  defer stmt.Close()
  defer con.Close()

  return rows, nil
}

func CountProducts()(int, error) {

  con := Connect()

  sql := "select count(*) from products"

  var count int

  err := con.QueryRow(sql).Scan(&count)

  if err != nil {
    
    return 0, err
  
  }

  defer con.Close()

  return count, nil

}
