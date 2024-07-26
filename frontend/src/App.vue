<template>
  <v-app>
    <app-navbar></app-navbar>
    <v-main>
        <router-view/>
    </v-main>
    <AppFooter></AppFooter>
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
