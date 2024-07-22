<script >
export default {
  data() {
    return {
      menuItems: [ // Menü öğelerini tanımla
        {title: 'Anasayfa', icon: 'mdi-home', route: '/'},
        {title: 'Şehir Seç', icon: 'mdi-city', route: '/cities'},
        {title: 'Namaz Vakitleri', icon:'mdi-clock',route: '/prayer-times'},
        {title: 'İletişim', icon: 'mdi-email', route: '/contact'},
      ]
    };
  },
};
</script>

<template>
  <v-app-bar app color="primary">
    <v-container>
      <v-row align="center">

        <v-col cols="auto">
          <v-toolbar-title class="title">Namaz Vakitleri</v-toolbar-title>
        </v-col>

        <!-- Var olan menü -->
        <v-col class="d-none d-lg-flex" cols="auto">
          <v-toolbar-items class="menu-items">
            <div class="menu-item" v-for="(item, index) in menuItems" :key="index">
              <router-link :to="item.route">
                <v-btn text size="large" style="color:#ADD8E6">
                  <v-icon>{{ item.icon }}</v-icon>

                  {{ item.title }}
                </v-btn>
              </router-link>
            </div>
          </v-toolbar-items>
        </v-col>
      </v-row>
    </v-container>
  </v-app-bar>


</template>

<style>
.menu-item {
  display: flex;
  align-items: center;
  margin-left: 5px;
}

.menu-items {
  margin-left: 200px;
}
</style>