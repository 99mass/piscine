#HERO_ID=$1
curl https://learn.zone01dakar.sn/assets/superhero/all.json | jq  '.[] | select (.id =='$HERO_ID') | .connections.relatives ' | tr -d '""'
