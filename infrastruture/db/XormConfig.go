package db

type XormConfig struct {
	// type of datasource,eg sqlite3,mysql
	DatasourceType string
	// connection string of the datasource
	DatasourceName string
}