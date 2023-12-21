# CounterApp

To run the app:

## With Docker

Start docker, then:

```
docker-compose up --build  
```
Note: you might have to run this command in sudo mode. 

## Manually
Start the backend application:
```
cd Backend
go run .
```

Then the frontend application:
```
cd Frontend/counter-app
npm install
npm run dev
```
