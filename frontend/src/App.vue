<template>
  <v-app>
    <v-main>
      <app-navbar></app-navbar>
        <router-view/>
      <app-footer></app-footer>
    </v-main>
  </v-app>
</template>

<script>

import AppNavbar from "@/components/AppNavbar.vue";
import AppFooter from "@/components/AppFooter.vue";

export default {
  name: 'App',

  components: {
    AppFooter,
    AppNavbar,
  },

  data: () => ({
    //
  }),
}
</script>
