package main

import (
	"fmt"
)

func main() {
	fmt.Println("Principal")

	// testes()
}

// func testes() {
// 	////////////////////////////////////////////////////////////////
// 	// TESTE dao.CarregaPessoa
// 	////////////////////////////////////////////////////////////////
// 	var db = dao.GetDB()
// 	listaPessoas, err := dao.CarregaPessoas(db)
// 	fmt.Printf("MAIN\nerro: %v\ntipo: %T\n", err, listaPessoas)
// 	for n, p := range listaPessoas {
// 		fmt.Printf("%3d [%T]: %v\n", n, p, p)
// 	}
// 	///////////////////////////////////////////////////////////////

// 	///////////////////////////////////////////////////////////////
// 	// TESTE dao.AdicionaPessoa
// 	///////////////////////////////////////////////////////////////
// 	var db = dao.GetDB()
// 	p, _ := pessoa.GetPessoaTest()
// 	p.Cpf = "38674832680"
// 	p.Usuario = "teste_inclusao"
// 	p, err := dao.AdicionaPessoa(db, p)
// 	fmt.Println(p, err)
// 	///////////////////////////////////////////////////////////////

// 	///////////////////////////////////////////////////////////////
// 	// TESTE dao.ProcuraPessoa
// 	///////////////////////////////////////////////////////////////
// 	var db = dao.GetDB()
// 	cpf := "38674832680"
// 	p, err := dao.ProcuraPessoa(db, cpf)
// 	fmt.Println(p, err)
// 	//////////////////////////////////////////////////////////////

// 	//////////////////////////////////////////////////////////////
// 	// TESTE dao.RemovePessoa
// 	//////////////////////////////////////////////////////////////
// 	db := dao.GetDB()
// 	cpf := "38674832680"
// 	err := dao.RemovePessoa(db, cpf)
// 	fmt.Println(err)
// 	//////////////////////////////////////////////////////////////

// }
