package dao

import (
	"controle_pessoal_de_financas/API/v1/model/tipo_conta"
	"database/sql"
	"errors"
)

var (
	tipoContaDB = map[string]string{
		"tabela":           "tipo_conta",
		"nome":             "nome",
		"descricaoDebito":  "descricao_debito",
		"descricaoCredito": "descricao_credito",
		"dataCriacao":      "data_criacao",
		"dataModificacao":  "data_modificacao",
		"estado":           "estado"}
)

func CarregaTiposConta(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = true
`
	query := getTemplateQuery("CarregaTipoConta", tipoContaDB, sql)

	return carregaTiposConta(db, query)
}

func CarregaTiposContaInativa(db *sql.DB) (tiposContas tipo_conta.TiposConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE
	{{.estado}} = false
`
	query := getTemplateQuery("CarregaTipoContaInativa", tipoContaDB, sql)

	return carregaTiposConta(db, query)
}

func AdicionaTipoConta(db *sql.DB, novoTipoConta *tipo_conta.TipoConta) (tc *tipo_conta.TipoConta, err error) {
	tc, err = tipo_conta.NewTipoConta(novoTipoConta.Nome, novoTipoConta.DescricaoDebito, novoTipoConta.DescricaoCredito)
	if err != nil {
		return
	}

	sql := `
INSERT INTO {{.tabela}}(
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}})
VALUES($1, $2, $3, $4, $5, $6)
`
	query := getTemplateQuery("AdicionaTipoConta", tipoContaDB, sql)

	return adicionaTipoConta(db, tc, query)
}

func ProcuraTipoConta(db *sql.DB, nome string) (tc *tipo_conta.TipoConta, err error) {
	sql := `
SELECT
	{{.nome}}, {{.descricaoDebito}}, {{.descricaoCredito}}, {{.dataCriacao}}, {{.dataModificacao}}, {{.estado}}
FROM
	{{.tabela}}
WHERE {{.nome}} = $1
`
	query := getTemplateQuery("ProcuraPessoa", tipoContaDB, sql)

	tiposConta, err := carregaTiposConta(db, query, nome)
	if len(tiposConta) == 1 {
		tc = tiposConta[0]
	} else {
		err = errors.New("Não foi encontrado um registro com o nome " + nome)
	}

	return
}

func AtivaTipoConta(db *sql.DB, nome string) (tc *tipo_conta.TipoConta, err error) {
	tipoContaBanco, err := ProcuraTipoConta(db, nome)
	if err != nil {
		return
	}

	tipoContaBanco.Ativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.nome}} = $3
`

	query := getTemplateQuery("AtivaTipoConta", tipoContaDB, sql)

	return estadoTipoConta(db, tipoContaBanco, query, nome)
}

func InativaTipoConta(db *sql.DB, nome string) (tc *tipo_conta.TipoConta, err error) {
	tipoContaBanco, err := ProcuraTipoConta(db, nome)
	if err != nil {
		return
	}

	tipoContaBanco.Inativa()

	sql := `
UPDATE {{.tabela}}
SET {{.estado}} = $1, {{.dataModificacao}} = $2
WHERE {{.nome}} = $3
`

	query := getTemplateQuery("InativaTipoConta", tipoContaDB, sql)

	return estadoTipoConta(db, tipoContaBanco, query, nome)
}

func AlteraTipoConta(db *sql.DB, nome string, tipoContaAlteracao *tipo_conta.TipoConta) (tc *tipo_conta.TipoConta, err error) {
	tipoContaBanco, err := ProcuraTipoConta(db, nome)
	if err != nil {
		return
	}

	err = tipoContaBanco.Altera(tipoContaAlteracao.Nome, tipoContaAlteracao.DescricaoDebito, tipoContaAlteracao.DescricaoCredito)
	if err != nil {
		return
	}

	sql := `
UPDATE {{.tabela}}
SET {{.nome}} = $1, {{.descricaoDebito}} = $2, {{.descricaoCredito}} = $3, {{.dataModificacao}} = $4
WHERE {{.nome}} = $5
`

	query := getTemplateQuery("AlteraTipoConta", tipoContaDB, sql)

	return alteraTipoConta(db, tipoContaBanco, query, nome)
}

func RemoveTipoConta(db *sql.DB, nome string) (err error) {
	sql := `
DELETE FROM
	{{.tabela}}
WHERE {{.nome}} = $1
`
	query := getTemplateQuery("RemoveTipoConta", tipoContaDB, sql)

	tc, err := ProcuraTipoConta(db, nome)
	if tc != nil {
		err = remove(db, tc.Nome, query)
	}

	return
}

func carregaTiposConta(db *sql.DB, query string, args ...interface{}) (tiposConta tipo_conta.TiposConta, err error) {
	registros, err := carrega(db, query, registrosTipoConta01, args...)

	tiposConta = converteEmTiposConta(registros)

	return
}

func adicionaTipoConta(db *sql.DB, novoTipoConta *tipo_conta.TipoConta, query string) (tc *tipo_conta.TipoConta, err error) {
	resultado, err := adiciona(db, novoTipoConta, query, setValoresTipoConta01)
	tipoContaTemp, ok := resultado.(*tipo_conta.TipoConta)
	if ok {
		tc = tipoContaTemp
	}
	return
}

func alteraTipoConta(db *sql.DB, tipoContaBanco *tipo_conta.TipoConta, query, chave string) (p *tipo_conta.TipoConta, err error) {
	resultado, err := altera(db, tipoContaBanco, query, setValoresTipoConta03, chave)
	tipoContaTemp, ok := resultado.(*tipo_conta.TipoConta)
	if ok {
		p = tipoContaTemp
	}
	return
}

func estadoTipoConta(db *sql.DB, tipoContaBanco *tipo_conta.TipoConta, query, chave string) (tc *tipo_conta.TipoConta, err error) {
	resultado, err := altera(db, tipoContaBanco, query, setValoresTipoConta02, chave)
	tipoContaTemp, ok := resultado.(*tipo_conta.TipoConta)
	if ok {
		tc = tipoContaTemp
	}
	return
}

func setValoresTipoConta01(stmt *sql.Stmt, novoRegistro interface{}) (r sql.Result, err error) {

	novoTipoConta, ok := novoRegistro.(*tipo_conta.TipoConta)

	if ok {
		r, err = stmt.Exec(
			novoTipoConta.Nome,
			novoTipoConta.DescricaoDebito,
			novoTipoConta.DescricaoCredito,
			novoTipoConta.DataCriacao,
			novoTipoConta.DataModificacao,
			novoTipoConta.Estado)
	}
	return
}

func setValoresTipoConta02(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {
	novoTipoConta, ok := novoRegistro.(*tipo_conta.TipoConta)

	if ok {
		r, err = stmt.Exec(
			novoTipoConta.Estado,
			novoTipoConta.DataModificacao,
			chave)
	}
	return
}

func setValoresTipoConta03(stmt *sql.Stmt, novoRegistro interface{}, chave string) (r sql.Result, err error) {

	novoTipoConta, ok := novoRegistro.(*tipo_conta.TipoConta)

	if ok {
		r, err = stmt.Exec(
			novoTipoConta.Nome,
			novoTipoConta.DescricaoDebito,
			novoTipoConta.DescricaoCredito,
			novoTipoConta.DataModificacao,
			chave)
	}
	return
}

func registrosTipoConta01(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error) {
	tipoContaAtual := new(tipo_conta.TipoConta)
	err = scanTipoConta01(rows, tipoContaAtual)
	if err != nil {
		return
	}
	novosRegistros = append(registros, tipoContaAtual)

	return
}

func scanTipoConta01(rows *sql.Rows, tipoContaAtual *tipo_conta.TipoConta) error {
	return rows.Scan(
		&tipoContaAtual.Nome,
		&tipoContaAtual.DescricaoDebito,
		&tipoContaAtual.DescricaoCredito,
		&tipoContaAtual.DataCriacao,
		&tipoContaAtual.DataModificacao,
		&tipoContaAtual.Estado)
}

func converteEmTiposConta(registros []interface{}) (tiposConta tipo_conta.TiposConta) {
	for _, r := range registros {
		tc, ok := r.(*tipo_conta.TipoConta)
		if ok {
			tiposConta = append(tiposConta, tc)
		}
	}

	return
}
