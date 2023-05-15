package dataBase

// функция нахождения ключа по значению
func mapkey(mapa map[string]int, value int) (key string, ok bool) {
	for k, v := range mapa {
	  if v == value { 
		key = k
		ok = true
		return
	  }
	}
	return
}