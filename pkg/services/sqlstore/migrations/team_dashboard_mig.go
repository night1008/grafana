package migrations

import . "github.com/grafana/grafana/pkg/services/sqlstore/migrator"

func addTeamDashboardMigrations(mg *Migrator) {
	teamDashboardV1 := Table{
		Name: "team_dashboard",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "org_id", Type: DB_BigInt},
			{Name: "team_id", Type: DB_BigInt},
			{Name: "dashboard_id", Type: DB_BigInt},
		},
		Indices: []*Index{
			{Cols: []string{"org_id", "team_id"}},
		},
	}

	mg.AddMigration("create team_dashbaord table", NewAddTableMigration(teamDashboardV1))

	//-------  indexes ------------------
	mg.AddMigration("add index team_dashbaord.org_id_team_id", NewAddIndexMigration(teamDashboardV1, teamDashboardV1.Indices[0]))
}
