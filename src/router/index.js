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