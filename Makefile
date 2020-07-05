run1:
	docker-compose -f docker-compose.client1.yaml pull
	docker-compose -f docker-compose.client1.yaml up

run2:
	docker-compose -f docker-compose.client2.yaml pull
	docker-compose -f docker-compose.client2.yaml

run3:
	docker-compose -f docker-compose.client3.yaml pull
	docker-compose -f docker-compose.client3.yaml

run4:
	docker-compose -f docker-compose.client4.yaml pull
	docker-compose -f docker-compose.client4.yaml