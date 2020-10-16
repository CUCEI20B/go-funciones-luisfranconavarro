package main

func Burbuja(s []int64)  {
	elemAComparar := len(s)
	for i := 0; i < elemAComparar; i++ {
		if i+1 < len(s){
			if s[i] > s[i+1]{
				aux := s[i]
				s[i] = s[i+1]
				s[i+1] = aux
			}
			if i+1 == elemAComparar-1{
				i = -1
				elemAComparar--	
			}
		}
	}
}

func main()  {

}