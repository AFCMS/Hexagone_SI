# Exercice 1

```shell
docker run -d \
  --name hexagone_si_ex1_db \
  -e POSTGRES_USER=my_user \
  -e POSTGRES_PASSWORD=my_password \
  -e POSTGRES_DB=tprevision \
  postgres:18-alpine

docker start hexagone_si_ex1_db
docker stop hexagone_si_ex1_db

docker network create hexagone_si_ex1
docker network rm hexagone_si_ex1

docker network connect hexagone_si_ex1 hexagone_si_ex1_db --alias db
docker network disconnect hexagone_si_ex1 hexagone_si_ex1_db

docker run --rm -it --network hexagone_si_ex1 \
    postgres:18-alpine \
    psql -h db -U my_user -d tprevision
```