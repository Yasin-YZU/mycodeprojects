package main

import "github.com/01-edu/z01"

func PrintComb(){
	for a :='0'; a <= '9'; a++ {
    for b := a + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {
				if a < b {
					if b < c {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						if a == '7' && b == '8' && c == '9' {
							z01.PrintRune(10) //ascii10 is linefeed
						} else {
							z01.PrintRune(44) //ascii 44 is comma
							z01.PrintRune(32) // asci 32 is space 
						}
					}
				}
			}
		}
	}
}

func main(){
	PrintComb()

}