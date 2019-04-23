package sql_builder

import (
	"edp-admin-console/models"
	"fmt"
)

const (
	SelectAllApplications = "select distinct on (\"name\") cb.name, cb.language, cb.build_tool, al.event as status_name " +
		"from codebase cb " +
		"		left join codebase_action_log cal on cb.id = cal.codebase_id " +
		"		left join action_log al on cal.action_log_id = al.id " +
		"%s" +
		"order by name, al.updated_at desc;"
	SelectAllApplicationsWithReleaseBranches = "select c.name as app_name, cb.name as branch_name, al.event " +
		"from codebase c " +
		"		left join codebase_branch cb on c.id = cb.codebase_id " +
		"		left join codebase_branch_action_log cbal on cb.id = cbal.codebase_branch_id " +
		"		left join action_log al on cbal.action_log_id = al.id " +
		"where cb.name is not null %s ;"
)

type ApplicationQueryBuilder struct {
}

func (this *ApplicationQueryBuilder) GetAllApplicationsQuery(filterCriteria models.ApplicationCriteria) string {
	if filterCriteria.Status == nil {
		return fmt.Sprintf(SelectAllApplications, "")
	}
	if *filterCriteria.Status == "active" {
		return fmt.Sprintf(SelectAllApplications, " where al.event = 'created' ")
	}
	return fmt.Sprintf(SelectAllApplications, " where al.event != 'created' ")
}

func (this *ApplicationQueryBuilder) GetAllApplicationsWithReleaseBranchesQuery(filterCriteria models.ApplicationCriteria) string {
	if filterCriteria.Status == nil {
		return fmt.Sprintf(SelectAllApplicationsWithReleaseBranches, "")
	}
	if *filterCriteria.Status == "active" {
		return fmt.Sprintf(SelectAllApplicationsWithReleaseBranches, " and al.event = 'created' ")
	}
	return fmt.Sprintf(SelectAllApplicationsWithReleaseBranches, " and al.event != 'created' ")
}