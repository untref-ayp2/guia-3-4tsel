package main

import (
	"fmt"
	"guia3/linkedlist"
	// "guia3/slicelist"
)

func main() {
/* 	l := linkedlist.NewLinkedList[int]()
	fmt.Println("Agregamos 1, 2 y 3 al final de la lista")
	fmt.Println("El tamaño debe ser 0 antes del primer append: ", l.Size())
	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Println(l)
	fmt.Println("Agregamos 0 al inicio de la lista")
	l.Prepend(0)
	fmt.Println(l)
	fmt.Println("Buscamos el numero 3")
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))

	l.Remove(2)
	l.Remove(0)
	l.Remove(3)
	fmt.Println("Buscamos el numero 3 luego de removerlo de la lista")
	fmt.Println(l)
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))
	l.Remove(1)
	l.Remove(1) //No deberia hacer nada
	l.Remove(1) //No deberia hacer nada
	fmt.Println(l)

	fmt.Println("Agregamos 0, 1 , 3, 4 y 5 al final de la lista")
	l.Append(0)
	l.Append(1)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	fmt.Println(l)
	fmt.Println("Agregamos 2 en la posicion 2")
	l.InsertAt(2, 2)

	fmt.Println(l) */

	/* sl := slicelist.NewSliceList[int]()
	fmt.Println("Agregamos 0 y 1 al slicelist")
	sl.Append(0)
	sl.Append(1)
	fmt.Println(sl)
	sl.Append(3)
	fmt.Println(sl)
	fmt.Println("Agregamos 0 al inicio de la lista")
	sl.Prepend(0)
	fmt.Println(sl)
	fmt.Println("Buscamos el numero 3")
	fmt.Println("Se encuentra en la posicion: ", sl.Search(3))

	sl.Remove(2)
	sl.Remove(0)
	sl.Remove(3)
	fmt.Println("Buscamos el numero 3 luego de removerlo de la lista")
	fmt.Println(sl)
	fmt.Println("Se encuentra en la posicion: ", sl.Search(3))
	sl.Remove(1)
	sl.Remove(1) //No deberia hacer nada
	sl.Remove(1) //No deberia hacer nada
	fmt.Println(sl)

	fmt.Println("Agregamos 0, 1 , 3, 4 y 5 al final de la lista")
	sl.Append(0)
	sl.Append(1)
	sl.Append(3)
	sl.Append(4)
	sl.Append(5)

	fmt.Println(sl)
	fmt.Println("Agregamos 2 en la posicion 2")
	sl.InsertAt(2, 2)

	fmt.Println(sl)
 */

	l1 := linkedlist.NewLinkedList[int]()
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	fmt.Println(l1.String())

	l2 := linkedlist.NewLinkedList[int]()
	l2.Append(4)
	l2.Append(5)
	l2.Append(6)
	l1.Append(8)
	fmt.Println(l2.String())

	la := linkedlist.NewLinkedList[int]()


	fmt.Println(la.Alternate(l1, l2))


}
