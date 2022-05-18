yars-run:
	go run \
		yars.go \
			--db-host='localhost' \
			--db-user='user' \
			--db-password='password' \
			--db-name='db' \
			--db-port='5432' \
		run \
			--folder='./projects' \

yars-load-units:
	go run \
		yars.go \
			--db-host='localhost' \
			--db-user='user' \
			--db-password='password' \
			--db-name='db' \
			--db-port='5432' \
		load \
			--file='./movielens100k/user.json' \
			--unit='user' \
			--path-project='projects/movielens' \
			--file-format='json' \
	&& \
	go run \
		yars.go \
			--db-host='localhost' \
			--db-user='user' \
			--db-password='password' \
			--db-name='db' \
			--db-port='5432' \
		load \
			--file='./movielens100k/item.json' \
			--unit='movie' \
			--path-project='projects/movielens' \
			--file-format='json'
