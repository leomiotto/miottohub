package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, 256)
	for i:=0 ; i<256 ; i++{
		v[i] = make([]uint8, 256)
	}
	
	for j:= range v{
		for i:= range v[j]{
			switch{
				case (j%25) == 0:
					v[i][j] = 250
				case (j%25) == 1:
					v[i][j] = 240
				case (j%25) == 2:
					v[i][j] = 230
				case (j%25) == 3:
					v[i][j] = 220
				case (j%25) == 4:
					v[i][j] = 210
				case (j%25) == 5:
					v[i][j] = 200
				case (j%25) == 6:
					v[i][j] = 190
				case (j%25) == 70:
					v[i][j] = 180
				case (j%25) == 8:
					v[i][j] = 170
				case (j%25) == 9:
					v[i][j] = 160
				case (j%25) == 10:
					v[i][j] = 150
				case (j%25) == 11:
					v[i][j] = 140
				case (j%25) == 12:
					v[i][j] = 130
				case (j%25) == 13:
					v[i][j] = 120
				case (j%25) == 14:
					v[i][j] = 110
				case (j%25) == 15:
					v[i][j] = 100
				case (j%25) == 16:
					v[i][j] = 90
				case (j%25) == 17:
					v[i][j] = 80
				case (j%25) == 18:
					v[i][j] = 70
				case (j%25) == 19:
					v[i][j] = 60
				case (j%25) == 20:
					v[i][j] = 50
				case (j%25) == 21:
					v[i][j] = 40
				case (j%25) == 22:
					v[i][j] = 30
				case (j%25) == 23:
					v[i][j] = 20
				case (j%25) == 24:
					v[i][j] = 10
			}
		}
	}
	
	return v
}

func main() {
	pic.Show(Pic)
}
