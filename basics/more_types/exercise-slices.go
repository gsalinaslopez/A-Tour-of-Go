package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic_matrix := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        pic_row := make([]uint8, dx)
        for x := 0; x < dx; x++ {
            pic_row[x] = uint8(x^y)
        }
        pic_matrix[y] = pic_row
    }
    return pic_matrix
}

func main() {
    pic.Show(Pic)
}
