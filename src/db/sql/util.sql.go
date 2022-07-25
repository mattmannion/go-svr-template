package sql

const Util_truncate_users_query string = `truncate users;`

const Util_reset_users_id_query string = `alter sequence users_id_seq restart;`

const Util_insert_default_users_query string = `
insert into users(firstname, lastname)
	values
	('matt', 'mannion'),
	('mack', 'gr'),
	('khris', 'rhodes');
`
