# Golang Whatsapp Broker

[Click here to see the repository](https://github.com/renatoviolin/golang-whatsapp-broker)

This is an Open-Source project with the goal to build a whatsapp broker that connect direct with [Whatsapp Business Platform](https://developers.facebook.com/docs/whatsapp)

The motivation started from the fact that Whatsapp Business Platform is now open to "any" company to connect without the use of a 3rd party payed service.

Hope you can enjoy it and help-me to improve this project.

[**Preview in youtube**](https://youtu.be/JE50HTfeKgw)

## Main goals
<ol> 
<li>Build a <b>webhook</b> to connect with Whatsapp Business Plataform without any 3rd party company. </li>
<li>With one whatsapp business number, be able to connect several agents, for example:</li>
<ul>
    <li>agent to deal with consumer service</li>
    <li>agent to deal with sales</li>
    <li>agent to detal with help desk</li>
    <li>...</li>
</ul>
<li> Build a <b>router</b> to route messages from users to agents. </li>
<li> Build a <b>admin</b> service to register agents in the router </li>
<li> Build a demo <b>agent</b> to show how to connect with broker. </li>
<li> Build a simple <b>web interface</b> to simulate an person answering from the given agent. </li>
</ol>


## Tools
Golang: all backend microservices were written in Golang. It is a good choice since it is very optimized and fast for microservices. In a enterprise environment these services will work with a high load.

Kafka: to stream messages from Whatsapp for the agents. Each agent will have three topics (in, out, error).

Redis: to cache the information about what users is connected with a given agent, and also handle the 24h time window required by Whatsapp to deliver messages since the last interaction from the user.

Postgresql: to store information about all agents.

MongoDB: used in the demo-app to store the messages from kafka topic in the specifc agent storage database.

ReactJS: used to build the simple chatbot interface


## Architecture
![](assets/broker.png)
The system has the following microservices:
- [webhook](/webhook): expose the /webhook endpoint that is needed by Whatsapp Business Plafatorm. There's a GET endpoint which is used to validade the identity (by a token you provied) and the POST endpoint which receives all the payload from Whatsapp. All payload is stores in the <b>webhook</b> topic.

- [router](/router): this is the core application, which consume all messages from <b>webhook</b> topic and deliver for each associated agent. The router also generate the first message asking the user to select an available agent.

- [admin](/admin): this is a CRUD API to create and mange the agents. Here is the [Postman collection](assets/Admin.postman_collection.json).

- [agent](/agent/backend): this is the demo application that show how to consume the messages from the *-in and *-error topics, store each message in it's own mongodb and connect via websockt with the [frontend](/agent/frontend).


## Steps to execute
For now, the microservices is not available as containers. This is a further task.

Please note that each service has a **config** folder with an **env.example** file that needs to be update with your credentials and renamed to **env** before build the given service.

The [docker-compose](docker-compose.yaml) starts all the infrastructure services:
- Kafka and Zookeeper
- Kafka-UI
- Redis
- Postgres
- MongoDB


### Step 1. Run the containers
You need to start all services in docker-compose:
```bash
docker-compose up
```
It may take some time to start Kafka. 

You can check if it is online accessing Kafka-UI: http://localhost:8080


### Step 2. Start the webhook services
In the webhook folder adjust the **config/env** file. You need to define a TOKEN that will be used to Whatsapp register and validate your webhook. 

Run the service (by default will start at port 5000):
```bash
go build src/main.go
```

### Step 3. Configure a tunnel service
This step is needed if you want to work in your local machine and don't have a domain or public IP. In this case you can use the [cloudfared tunnel](https://github.com/cloudflare/cloudflared) to expose your webhook from your localhost. After install, run it:
```bash
cloudflared tunnel --url http://localhost:5000
```

You will expose the webhook service from step 2 in random generated hostname provided by cloudflare.


### Step 4. Create an Facebook Developer Account, Whatsapp Application and Configure Webhook
You need to follow the instructions here: https://developers.facebook.com/docs/whatsapp/cloud-api/get-started

When you achieve the part where you need to configure the webhook, you need to provided the name generate by step 3 and the token used in step 2.

Once configured you will see Whatsapp Business calling your webook with provided token and validating.


### Step 5. Get the temporary Whatsapp Token
With a developer account you can send up to 1000 messages within 24h time window, and register up to 5 whatsapp numbers that can send message to your whatsapp businness account.

The provided token will be used in the **router** service.


### Step 6. Starting router service
You need to place the token into **router/config/env** before start the service.
In the router/src folder:
```bash
go run cmd/main.go
```
It will connect to postgresql, read the registered agents (by default only Agent A), prepare the kafka topics and start consuming the webhook. You must be sure that kafka is running.


### Step 7.1 Starting Agent A demo - backend
In this step you need to configure the whatsapp phone number you registerd in step 4. Fill the variable **DEMO_WAID** with the number. In a production application you may get the number from the payload instead of fix it here. This was done this way to be easy to test the core application.

In the agent/backend/src folder:
```bash
go run cmd/main.go
```
It will connect to mongoDB and start consuming the **agent_A-in** topic.


### Step 7.2 Starting Agent A demo - frontend
This web application was build with react, so you need node.js installed to build.

In the agent/frontend folder:
```bash
npm start
```

It will start the frontend application in http://localhost:3000

If you type something and send, you'll receive a message telling you are not connected with any user. This is a receptive application, so the user must start a conversation with the agent.


### Step 8 Talking
From your whatsapp application, send a message to the number you get from step 4. It must start with 1 (555) *****.

Any message send to that number will trigger your webhook.

The router will consume the message and will:

1) check if the user is associated with any agent.

2) if yes, publish the message in the agent_X-in topic.

3) if no, display the list options with all registered agents (you can register a new agent using the agent service)

4) anytime you can type **#exit** to disconnect from the agent and return to main menu.


# Next steps
Here's the next steps by priority:

1) configure a docker-compose to start all the services (webhook, router, agent and admin).
2) improve the agent to be able to handle different users (different whatsapp numbers connections)
3) improve the agent front end.


# Disclaimer
This projected started as a personal project to learn about the whatsapp API and improve the my knowledge about the tools used here. 

This is not a production ready application, but the core services (**webhook** and **router**) is almost there.

