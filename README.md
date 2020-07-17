# app-survey
Monorepo with frontend and backend applications

Contains 2 React frontends, one in TS and the other is JS

## Docker compose
Build
```bash
docker-compose build --no-cache
```

Run 
```bash
docker-compose up -d
```

Tear down
```bash
docker-compose down
```

## Endpoints
ReactJS frontend
```bash
curl localhost
```

ReactTS frontend
```bash
curl localhost/webappts
```

Backend
```bash
curl localhost/api/
```

## Creating a new application
Using [create-react-app](https://reactjs.org/docs/create-a-new-react-app.html)

```bash
npx create-react-app webapp
cd webapp
npm start
```

Or for typescript
```bash
npx create-react-app webapp-ts --template typescript
cd webapp-ts
npm start
```
