alter table codebase add jenkins_slave_id integer constraint jenkins_slave_fk references jenkins_slave;