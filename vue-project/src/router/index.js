import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import StudentsView from "../views/Students/View.vue";
import StudentsCreate from "../views/Students/Create.vue";
import StudentsEdit from "../views/Students/Edit.vue";
import StudentsOne from "../views/Students/Viewone.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/about",
      name: "about",
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/students",
      name: "students",
      component: StudentsView,
    },
    {
      path: "/students/create",
      name: "StudentsCreate",
      component: StudentsCreate,
    },
    {
      path: "/students/:id/edit",
      name: "StudentsEdit",
      component: StudentsEdit,
    },
    {
      path: "/students/:id",
      name: "StudentsOne",
      component: StudentsOne,
    },
  ],
});

export default router;
