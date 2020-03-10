import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      isPublic: true
    }
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/Register.vue"),
    meta: {
      isPublic: true
    }
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/Login.vue"),
    meta: {
      isPublic: true
    }
  },
  {
    path: "/course",
    name: "Course",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Course.vue"),
    meta: {
      isPublic: true
    }
  },
  {
    path: "/article",
    component: () => import("../views/Article.vue"),
    children: [
      {
        path: "",
        name: "Article",
        component: () => import("../views/Article/Index.vue"),
        meta: {
          isPublic: true
        }
      },
      {
        path: "create",
        name: "Article/Create",
        component: () => import("../views/Article/Edit.vue")
      },
      {
        path: "edit/:id",
        name: "Article/Edit",
        component: () => import("../views/Article/Edit.vue")
      }
    ]
  },
  {
    path: "/profile",
    component: () => import("../views/Profile.vue"),
    children: [
      {
        path: "",
        name: "Profile",
        component: () => import("../views/Profile/Index.vue")
      }
    ]
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
    meta: {
      isPublic: true
    }
  }
];

const router = new VueRouter({
  routes
});

// 路由器守卫
router.beforeEach((to, from, next) => {
  if (!to.meta.isPublic && !localStorage.getItem("isLogin")) {
    next({ name: "Login" });
  }

  next();
});

export default router;
