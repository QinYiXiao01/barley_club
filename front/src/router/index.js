import { createRouter, createWebHistory } from 'vue-router'
import MainContent from '../components/MainContent.vue'
import BeerDiary from '../views/BeerDiary.vue'
import FindBar from '../views/FindBar.vue'
import MyBeer from '../views/MyBeer.vue'
import BeerWiki from '../views/BeerWiki.vue'



const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: MainContent
    },
    {
      path: '/diary',
      name: 'diary',
      component: BeerDiary
    },
    {
      path: '/find',
      name: 'find-bar',
      component: FindBar
    },
    {
      path: '/my',
      name: 'my-beer',
      component: MyBeer
    },
    {
      path: '/wiki',
      name: 'wiki',
      component: BeerWiki
    },

    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue')
    // }
  ]
})

export default router
