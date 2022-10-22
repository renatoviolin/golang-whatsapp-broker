kafka-topics.sh --create --topic myTopic --bootstrap-server localhost:9092
kafka-console-producer.sh --topic myTopic --bootstrap-server localhost:9092
kafka-console-consumer.sh --topic myTopic --from-beginning --bootstrap-server localhost:9092

docker exec -it mysql sh -c 'mysql -uroot -proot < scripts/script.sql'



kafka-console-consumer.sh --topic dev-msk-cea-br-ceatalk-bot-mensagem-enviada --from-beginning --bootstrap-server localhost:9092