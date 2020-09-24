# Hello Vite

### Step 1 Create 
```sh
$ yarn create vite-app <Project>
```

### Step 2 
```sh
$ yarn dev
```

### Step 3 install router

```sh
yarn add vue-router@v4.0.0-alpha.11
```
```js
# router/index.js
import { createWebHistory, createRouter } from "vue-router";
import Home from "../components/HelloWorld.vue";
import Hello from "../views/Hello.vue";
import About from "../views/About.vue";
const history = createWebHistory();
const routes = [
  { path: "/", component: Home },
  { path: "/hello/", component: Hello },
  { path: "/about/", component: About },
];
const router = createRouter({ history, routes });
export default router;
```
```js
# main.js
import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import router from "./router";

createApp(App).use(router).mount('#app')
```

### Step 4 Build
```sh
$ yarn build
```

### Step 5 run 

```sh
$ go run .
```

---

