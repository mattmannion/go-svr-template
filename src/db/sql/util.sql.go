package sql

const Util_truncate_users_query string = `truncate users;`

const Util_reset_users_id_query string = `alter sequence users_id_seq restart;`

const Util_insert_default_users_query string = `
insert into users(firstname, lastname, email, username, password)
	values
	('matt', 'mannion', 'mm@mm.com', 'mm', 'mm'),
	('mack', 'gr', 'mgr@mgr.com', 'mgr', 'mgr'),
	('khris', 'rhodes', 'kr@kr.com', 'kr', 'kr');
`
