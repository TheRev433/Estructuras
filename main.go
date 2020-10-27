package main

import "fmt"

type Multimedia interface {
	mostrar()
}

type Imagen struct {
	titulo, formato, canales string
}

type Audio struct {
	titulo, formato string
	duracion        float64
}

type Video struct {
	titulo, formato string
	frames          int64
}

type ContenidoWeb struct {
	multSlice *[]Multimedia
}

func (c *ContenidoWeb) mosrtar() {
	for _, a := range *c.multSlice {
		a.mostrar()
		fmt.Println("")
	}
}

func (i *Imagen) mostrar() {
	fmt.Println("Titulo: ", i.titulo)
	fmt.Println("Formato: ", i.formato)
	fmt.Println("Canales: ", i.canales)
}

func (a *Audio) mostrar() {
	fmt.Println("Titulo: ", a.titulo)
	fmt.Println("Formato: ", a.formato)
	fmt.Println("Duracion: ", a.duracion)
}

func (v *Video) mostrar() {
	fmt.Println("Titulo: ", v.titulo)
	fmt.Println("Formato: ", v.formato)
	fmt.Println("Frames: ", v.frames)
}

func main() {
	var titulo, formato, canales string
	var duracion float64
	var frames int64
	var opt uint64
	
	s := []Multimedia{}
	c1 := ContenidoWeb{&s}

	for opt != 5 {
		fmt.Println("Seleccione la opcion deseada")
		fmt.Println("")
		fmt.Println("1.- Agregar una Imagen")
		fmt.Println("2.- Agregar un Audio")
		fmt.Println("3.- Agregar un Video")
		fmt.Println("4.- Mostrar el contenido agregado")
		fmt.Println("5.- Agregar un Video")

		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Println("Ingrese el titulo")
			fmt.Scanln(&titulo)
			fmt.Println("Ingrese el formato")
			fmt.Scanln(&formato)
			fmt.Println("Ingrese los canales")
			fmt.Scanln(&canales)
			i1:= Imagen{titulo, formato, canales}
			s = append(s,&i1)
		case 2:
			fmt.Println("Ingrese el titulo")
			fmt.Scanln(&titulo)
			fmt.Println("Ingrese el formato")
			fmt.Scanln(&formato)
			fmt.Println("Ingrese la duracion")
			fmt.Scanln(&duracion)
			a1:= Audio{titulo, formato, duracion}
			s = append(s,&a1)
		case 3:
			fmt.Println("Ingrese el titulo")
			fmt.Scanln(&titulo)
			fmt.Println("Ingrese el formato")
			fmt.Scanln(&formato)
			fmt.Println("Ingrese los frames")
			fmt.Scanln(&frames)
			v1:= Video{titulo, formato, frames}
			s = append(s,&v1)
		case 4:
			c1.mosrtar()
		}
	}
	
}
