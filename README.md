# RCJV App

With the raising number of teams and referees,
and teams missing their runs or games we decided
to develop an app, so each team can filter based on
league, school or team for their runs and matches and 
can configure all kinds of push notifications.

# How to host
The backend and web frontends are all in one docker
image and are completely stateless, so our we run them
in my docker swarm cluster (1 manager, 2 workers) and
I have the databases on another system.
You could run all services in one docker compose file, there 
is an example file in the repo.

# Questions
Feel free to open an issue, if you have a question.

# Devs
 - Backend @EliasB-NU
 - AdminSite @EliasB-NU
 - WebView @EliasB-NU
 - App @strassbcarla

# Architecture

 - Backend [golang](https://go.dev)
   - [Fiber](https://gofiber.io/)
   - [Gorm](https://gorm.io)
   - [goValkey](https://github.com/valkey-io/valkey-go)
 - Databases
   - [Postgresql 17](https://www.postgresql.org/)
   - [Valkey 8](https://valkey.io/)
 - Frontend [VueJS 3](https://vuejs.org/)
   - [VueRouter](https://router.vuejs.org/)
   - [TailwindCSS](https://tailwindcss.com/)
 - App [Flutter](https://flutter.dev/)
   - [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging)

# ToDo
 - [ ] Admin View
 - [ ] Public View
 - [ ] Backend
   - [ ] Fetch Soccer stuff
   - [ ] Websocket connection for soccer stuff
   - [ ] Push notifications
 - [ ] App