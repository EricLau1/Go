package main

import (
	"fmt"
)

type Title struct {
	Level int
	Description string
}

var titles = []Title{
	Title{1,   "Iniciante"},
	Title{10,  "Língua de ferro"},
	Title{20,  "Língua de aço"},
	Title{30,  "Língua de bronze"},
	Title{40,  "Língua de prata"},
	Title{50,  "Língua de ouro"},
	Title{60,  "Língua de diamante"},
	Title{70,  "Mestre"},
	Title{80,  "Mestre Superior"},
	Title{90,  "Mestre Supremo"},
	Title{100, "Grande Mestre Supremo das Línguas"},
}

var exp = []int{1, 10, 20, 30 , 40, 50, 60, 70, 85, 100, 
				   120, 140, 160, 180, 200, 225, 250, 275, 
				   300, 350, 400, 450, 500, 550, 600, 700, 
				   800, 900, 1000, 1150, 1300, 1450, 1600, 
				   1750, 2000, 2200, 3400, 3600, 3800, 4100, 
				   4400, 4700, 5000, 5400, 5800, 6200, 6700, 
				   7200, 7700, 8200, 8700, 9200, 9800, 10400, 
				   11000, 11600, 12200, 12800, 13500, 14200, 
				   15000, 16000, 17000, 18000, 20000, 22000, 
				   24000, 26000, 28000, 30000, 32500, 35000, 
				   37500, 40000, 42500, 45000, 47500, 50000, 
				   52500, 55000, 58000, 61000, 64000, 67000, 
				   70000, 73500, 77000, 80050, 83550, 90000, 
				   95000, 100000, 105000, 110000, 120000, 
				   130000, 140000, 150000, 175000, 200000}



