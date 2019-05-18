package dao

import (
	"bytes"
	"controle_pessoal_de_financas/API/v1/config"
	"database/sql"
	"fmt"
	"html/template"
	"log"

	_ "github.com/lib/pq"
)

// GetDB retorna uma conexão com o banco de dados de acordo com as informações obtida de configurações
func GetDB() *sql.DB {
	config := config.AbrirConfiguracoes()
	connStr := getStringConexao(config)
	db, err := sql.Open(config["DB"], connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getStringConexao(config config.Configuracoes) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config["DBhost"], config["DBporta"], config["DBusuario"], config["DBsenha"], config["DBnome"])
}

func getTemplateQuery(nome string, campos map[string]string, sql string) string {
	t := template.Must(template.New(nome).Parse(sql))
	query := new(bytes.Buffer)
	t.Execute(query, campos)

	return query.String()
}

func carrega(db *sql.DB, query string, appendRegistros func(rows *sql.Rows, registros []interface{}) (novosRegistros []interface{}, err error), args ...interface{}) (registros []interface{}, err error) {
	queryStmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	rows, err := queryStmt.Query(args...)
	defer queryStmt.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		registros, err = appendRegistros(rows, registros)
		if err != nil {
			return
		}

	}
	err = rows.Err()
	if err != nil {
		registros = nil
		return
	}

	return
}

func adiciona(db *sql.DB, novoRegistro interface{}, query string, setValores func(*sql.Stmt, interface{}) (sql.Result, error)) (r interface{}, err error) {

	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	r = novoRegistro

	return
}

func remove(db *sql.DB, chavePrimaria interface{}, query string) (err error) {
	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(chavePrimaria)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	return
}

func altera(db *sql.DB, novoRegistro interface{}, query string, setValores func(*sql.Stmt, interface{}, string) (sql.Result, error), chave string) (r interface{}, err error) {

	transacao, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := transacao.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = setValores(stmt, novoRegistro, chave)
	if err != nil {
		return
	}

	err = transacao.Commit()
	if err != nil {
		return
	}

	r = novoRegistro

	return
}
