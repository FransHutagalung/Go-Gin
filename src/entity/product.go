package entity

type Product struct{
	Id int 
	Nama string
	Harga int
	Stock int
}

func (p Product) StockStatus() string{
	if p.Stock < 3 {
		return "Stock Hampir Habis"
	} else if p.Stock < 10 {
		 return "Stock Terbatas"
	}
	return "Stock OK"
}