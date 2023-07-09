curl -s https://learn.zone01dakar.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"ssambadi\"}}){id}}"}' | cut -c 24-27 
