package dao

import (
	"testing"
	"time"

	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"
	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/pessoa"
)

var (
	testLancamento01  *lancamento.Lancamento
	testPessoaAdmin02 *pessoa.Pessoa
)

func TestAdicionaLancamento02(t *testing.T) {
	var err error

	p := getPessoaAdmin2()
	testPessoaAdmin02, err = AdicionaPessoa02(db2, p)
	if err != nil {
		t.Error(err)
	}

	l := getLancamento2(testPessoaAdmin02)
	testLancamento01, err = AdicionaLancamento02(db2, l)
	if err != nil {
		t.Error(err)
	}

	cpfEsperado := l.CpfPessoa
	cpfObtido := testLancamento01.CpfPessoa
	id := testLancamento01.ID
	if cpfEsperado != cpfObtido {
		t.Errorf("CPF de pessoa em lancamento %d diferente do esperado. Esperado: '%s', obtido: '%s'", id, cpfEsperado, cpfObtido)
	}
}

func TestCarregaLancamentos02(t *testing.T) {
	lancamentos, err := CarregaLancamentos02(db2)
	if err != nil {
		t.Error(err)
	}

	quantLancamentos := len(lancamentos)
	quantEsperada := 1
	if quantLancamentos != quantEsperada {
		t.Errorf("consulta de Lancamentos retornou uma quantidade de registros diferente do esperado. Esperado: %d, obtido: %d", quantEsperada, quantLancamentos)
	}

	cpf := testLancamento01.CpfPessoa
	id := testLancamento01.ID
	for _, l := range lancamentos {
		if l.CpfPessoa != cpf {
			t.Errorf("registro de lancamento(CpfPessoa) encontrado diferente do esperado. Esperado: '%s', obtido: '%s'", cpf, l.CpfPessoa)
		}

		if l.ID != id {
			t.Errorf("registro de lancamento(ID) encontrado diferente do esperado. Esperado: %d, obtido: %d", id, l.ID)
		}
	}
}

func TestAlteraLancamento02(t *testing.T) {
	id := testLancamento01.ID
	novoNumero := "Ln1234"
	novaData := time.Now()
	novaDescricao := "NOVA descrição Lanc Ln1234"

	testLancamento01.Numero = novoNumero
	testLancamento01.Data = novaData
	testLancamento01.Descricao = novaDescricao

	transacao := db2.Begin()
	l, err := AlteraLancamento02(db2, transacao, id, testLancamento01)
	transacao.Commit()
	if err != nil {
		t.Error(err)
	}

	numero := l.Numero
	if numero != novoNumero {
		t.Errorf("alteração de lancamento com ID %d retornou um 'Número' diferente do esperado. Esperado: '%s', obtido: '%s'", id, novoNumero, numero)
	}

	data := l.Data
	if data.Unix() != novaData.Unix() {
		t.Errorf("alteração de lancamento com ID %d retornou uma 'Data' diferente do esperado. Esperado: '%s', obtido: '%s'", id, novaData, data)
	}

	descricao := l.Descricao
	if descricao != novaDescricao {
		t.Errorf("alteração de lancamento com ID %d retornou uma 'Descrição' diferente do esperado. Esperado: '%s', obtido: '%s'", id, novaDescricao, novoNumero)
	}
}

func TestInativaLancamento02(t *testing.T) {
	id := testLancamento01.ID
	l, err := InativaLancamento02(db2, id)
	if err != nil {
		t.Error(err)
	}

	if l != nil {
		idObtido := l.ID
		if idObtido != id {
			t.Errorf("inativação de lançamento retornou um lançamento com ID diferente do esperado. Esperado: %d, obtido: %d", id, idObtido)
		}

		estadoObtido := l.Estado
		estadoEsperado := false
		if estadoObtido != estadoEsperado {
			t.Errorf("inativação de lançamento retornou um lançamento com estado diferente do esperado. Esperado: %t, obtido: %t", estadoEsperado, estadoObtido)
		}
	} else {
		t.Errorf("func InativaLancamento02(db2, %d) retornou um lançamento nulo(nil)", id)
	}
}

func TestCarregaLancamentosInativo02(t *testing.T) {
	lancamentos, err := CarregaLancamentosInativo02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(lancamentos)
	quantEsperada := 1
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos inativos retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}
}

func TestCarregaLancamentosInativoPorCPF02(t *testing.T) {
	cpf := testLancamento01.CpfPessoa
	lancamentos, err := CarregaLancamentosInativoPorCPF02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 1
	quantObtida := lancamentos.Len()
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos inativos por CPF '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, quantEsperada, quantObtida)
	}

	for _, lanc := range lancamentos {
		cpfObtido := lanc.CpfPessoa
		if lanc.CpfPessoa != cpf {
			t.Errorf("consulta de lançamentos inativos por CPF '%s' retornou um cpf diferente do esperado. Esperado '%[1]s', obtido: '%s'", cpf, cpfObtido)
		}
	}
}

func TestAtivaLancamento02(t *testing.T) {
	id := testLancamento01.ID
	l, err := AtivaLancamento02(db2, id)
	if err != nil {
		t.Error(err)
	}

	if l != nil {
		idObtido := l.ID
		if idObtido != id {
			t.Errorf("ativação de lançamento retornou um lançamento com ID diferente do esperado. Esperado: %d, obtido: %d", id, idObtido)
		}

		estadoObtido := l.Estado
		estadoEsperado := true
		if estadoObtido != estadoEsperado {
			t.Errorf("ativação de lançamento retornou um lançamento com estado diferente do esperado. Esperado: %t, obtido: %t", estadoEsperado, estadoObtido)
		}
	} else {
		t.Errorf("func AtivaLancamento02(db2, %d) retornou um lançamento nulo(nil)", id)
	}
}

func TestCarregaLancamentosAtivo02(t *testing.T) {
	lancamentos, err := CarregaLancamentosAtivo02(db2)
	if err != nil {
		t.Error(err)
	}

	quantObtida := len(lancamentos)
	quantEsperada := 1
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos ativos retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", quantEsperada, quantObtida)
	}
}

func TestCarregaLancamentosAtivoPorCPF02(t *testing.T) {
	cpf := testLancamento01.CpfPessoa
	lancamentos, err := CarregaLancamentosAtivoPorCPF02(db2, cpf)
	if err != nil {
		t.Error(err)
	}

	quantEsperada := 1
	quantObtida := lancamentos.Len()
	if quantEsperada != quantObtida {
		t.Errorf("consulta de lançamentos ativos por CPF '%s' retornou uma quantidade de registros incorreta. Esperado: %d, obtido: %d", cpf, quantEsperada, quantObtida)
	}

	for _, lanc := range lancamentos {
		cpfObtido := lanc.CpfPessoa
		if lanc.CpfPessoa != cpf {
			t.Errorf("consulta de lançamentos ativos por CPF '%s' retornou um cpf diferente do esperado. Esperado '%[1]s', obtido: '%s'", cpf, cpfObtido)
		}
	}
}

func TestRemoveLancamento02(t *testing.T) {
	err := RemoveLancamento02(db2, testLancamento01.ID)
	if err != nil {
		t.Error(err)
	}

	err = RemovePessoa02(db2, testPessoaAdmin02.Cpf)
	if err != nil {
		t.Error(err)
	}
}

// TESTES ANTIGOS
// import (
// 	"github.com/paulocsilvajr/controle_pessoal_de_financas/API/v1/model/lancamento"
// 	"testing"
// 	"time"
// )

// var (
// 	numLanc01, numLanc02, numLanc03 int
// )

// func TestAdicionaLancamento(t *testing.T) {
// 	TestAdicionaPessoa(t)

// 	l1 := lancamento.GetLancamentoTest()
// 	l1.CpfPessoa = cpf

// 	l2 := lancamento.GetLancamentoTest()
// 	l2.Numero = "5678A"
// 	l2.Descricao = "Lançamento teste 02 - parcela 1"
// 	l2.CpfPessoa = cpf

// 	l3 := lancamento.GetLancamentoTest()
// 	l3.Numero = "5678B"
// 	l3.Descricao = "Lançamento teste 02 - parcela 2"
// 	l3.CpfPessoa = cpf

// 	l5, err := AdicionaLancamento(db, l1)
// 	numLanc01 = l5.ID
// 	if err != nil {
// 		t.Error(err, l5)
// 	}

// 	l6, err := AdicionaLancamento(db, l2)
// 	numLanc02 = l6.ID
// 	if err != nil {
// 		t.Error(err, l6)
// 	}

// 	l7, err := AdicionaLancamento(db, l3)
// 	numLanc03 = l7.ID
// 	if err != nil {
// 		t.Error(err, l7)
// 	}

// 	l4 := lancamento.New(0, "", time.Date(2001, 1, 12, 12, 31, 0, 0, new(time.Location)), "", "")

// 	l8, err := AdicionaLancamento(db, l4)
// 	if err.Error() != "Tamanho de campo CPF inválido[0 caracter(es)]" {
// 		t.Error(err, l8)
// 	}

// 	l4.CpfPessoa = cpf
// 	l8, err = AdicionaLancamento(db, l4)
// 	if err.Error() != "Tamanho de campo Descrição inválido[0 caracter(es)]" {
// 		t.Error(err, l8)
// 	}

// 	l4.Descricao = "Lançamento teste 03"
// 	l8, err = AdicionaLancamento(db, l4)
// 	if err != nil {
// 		t.Error(err, l8)
// 	}
// }

// func TestInativaLancamentoECarregaLancamentosInativos(t *testing.T) {
// 	l1, err := InativaLancamento(db, numLanc01)
// 	if err != nil {
// 		t.Error(err, l1)
// 	}

// 	l2, err := InativaLancamento(db, numLanc02)
// 	if err != nil {
// 		t.Error(err, l2)
// 	}

// 	l3, err := InativaLancamento(db, numLanc03)
// 	if err != nil {
// 		t.Error(err, l3)
// 	}

// 	l4, err := InativaLancamento(db, 0)
// 	if err.Error() != "Não foi encontrado um registro com o ID 0" {
// 		t.Error(err, l4)
// 	}

// 	if l1.Estado != false {
// 		t.Error("Estado do lancamento inválido, deveria ser false", l1)
// 	}

// 	if l2.Estado != false {
// 		t.Error("Estado do lancamento inválido, deveria ser false", l2)
// 	}

// 	if l3.Estado != false {
// 		t.Error("Estado do lancamento inválido, deveria ser false", l3)
// 	}

// 	lancamentos, err := CarregaLancamentosInativo(db)
// 	if err != nil {
// 		t.Error(err, lancamentos)
// 	}

// 	if len(lancamentos) == 0 {
// 		t.Error(lancamentos)
// 	}

// 	if len(lancamentos) < 3 {
// 		t.Error(lancamentos)
// 	}

// 	lancamentos, err = CarregaLancamentosInativoPorCpf(db, cpf)
// 	if err != nil {
// 		t.Error(err, lancamentos)
// 	}

// 	if len(lancamentos) == 0 {
// 		t.Error(lancamentos)
// 	}

// 	if len(lancamentos) < 3 {
// 		t.Error(lancamentos)
// 	}
// }

// func TestAtivaLancamentoECarregaLancamentosAtivos(t *testing.T) {
// 	l1, err := AtivaLancamento(db, numLanc01)
// 	if err != nil {
// 		t.Error(err, l1)
// 	}

// 	l2, err := AtivaLancamento(db, numLanc02)
// 	if err != nil {
// 		t.Error(err, l2)
// 	}

// 	l3, err := AtivaLancamento(db, numLanc03)
// 	if err != nil {
// 		t.Error(err, l3)
// 	}

// 	l4, err := AtivaLancamento(db, 0)
// 	if err.Error() != "Não foi encontrado um registro com o ID 0" {
// 		t.Error(err, l4)
// 	}

// 	if l1.Estado != true {
// 		t.Error("Estado do lancamento inválido, deveria ser false", l1)
// 	}

// 	if l2.Estado != true {
// 		t.Error("Estado do lancamento inválido, deveria ser false", l2)
// 	}

// 	if l3.Estado != true {
// 		t.Error("Estado do lancamento inválido, deveria ser false", l3)
// 	}

// 	lancamentos, err := CarregaLancamentosAtivo(db)
// 	if err != nil {
// 		t.Error(err, lancamentos)
// 	}

// 	if len(lancamentos) == 0 {
// 		t.Error(lancamentos)
// 	}

// 	if len(lancamentos) < 4 {
// 		t.Error(lancamentos)
// 	}

// 	lancamentos, err = CarregaLancamentosAtivoPorCpf(db, cpf)
// 	if err != nil {
// 		t.Error(err, lancamentos)
// 	}

// 	if len(lancamentos) == 0 {
// 		t.Error(lancamentos)
// 	}

// 	if len(lancamentos) < 4 {
// 		t.Error(lancamentos)
// 	}
// }

// func TestCarregaLancamentos(t *testing.T) {
// 	listaLancamentos, err := CarregaLancamentos(db)

// 	if err != nil {
// 		t.Error(err, listaLancamentos)
// 	}

// 	if len(listaLancamentos) == 0 {
// 		t.Error(listaLancamentos)
// 	}

// 	if len(listaLancamentos) < 4 {
// 		t.Error(listaLancamentos)
// 	}
// }

// func TestCarregaLancamentosPorCpf(t *testing.T) {
// 	listaLancamentos, err := CarregaLancamentosPorCpf(db, cpf)

// 	if err != nil {
// 		t.Error(err, listaLancamentos)
// 	}

// 	if len(listaLancamentos) == 0 {
// 		t.Error(listaLancamentos)
// 	}

// 	if len(listaLancamentos) < 4 {
// 		t.Error(listaLancamentos)
// 	}
// }

// func TestProcuraLancamento(t *testing.T) {
// 	l1, err := ProcuraLancamento(db, numLanc01)
// 	if err != nil {
// 		t.Error(err, l1)
// 	}

// 	l2, err := ProcuraLancamento(db, numLanc02)
// 	if err != nil {
// 		t.Error(err, l2)
// 	}

// 	l3, err := ProcuraLancamento(db, numLanc03)
// 	if err != nil {
// 		t.Error(err, l3)
// 	}
// }

// func TestAlteraLancamento(t *testing.T) {
// 	novoLancamento := lancamento.GetLancamentoTest()
// 	novoLancamento.CpfPessoa = cpf

// 	l1, err := AlteraLancamento(db, numLanc01, novoLancamento)
// 	if err != nil {
// 		t.Error(err, l1)
// 	}

// 	if l1.CpfPessoa != novoLancamento.CpfPessoa ||
// 		l1.Data.Unix() != novoLancamento.Data.Unix() ||
// 		l1.Descricao != novoLancamento.Descricao ||
// 		l1.Numero != novoLancamento.Numero {
// 		t.Error("Erro na alteração de lancamento(ID ou CpfPessoa ou Data ou Numero ou Descricao)", l1, novoLancamento)
// 	}
// }

// func TestRemoveLancamento(t *testing.T) {
// 	err := RemoveLancamento(db, numLanc01)
// 	if err != nil {
// 		t.Error(err, numLanc01)
// 	}

// 	err = RemoveLancamento(db, numLanc02)
// 	if err != nil {
// 		t.Error(err, numLanc02)
// 	}

// 	err = RemoveLancamento(db, numLanc03)
// 	if err != nil {
// 		t.Error(err, numLanc03)
// 	}

// 	// Lancamento com a descrição "Lançamento teste 03"(var l4 de TestAdicionaLancamento(...)) é excluida automaticamente ao excluir a pessoa com cpf da variável cpf, por causa da CONSTRAINT pessoa_lancamento_fk com ON DELETE CASCADE na tabela lancamento do DB
// 	TestRemovePessoa(t)
// }
