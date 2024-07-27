<script >
export default {
  data() {
    return {
      drawer: false, // Hamburger menü durumu
      mobileMenuOpen: false, // Tost menünün durumu
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
          <v-img src="@/assets/logo.png" max-height="50" contain></v-img>
        </v-col>
        <v-col cols="auto">
          <v-toolbar-title class="title">Namaz Vakitleri</v-toolbar-title>
        </v-col>
        <!-- Tost menüsü butonu -->
        <v-col class="d-lg-none" cols="auto">
          <v-btn icon @click.stop="mobileMenuOpen = !mobileMenuOpen">
            <v-icon>{{ mobileMenuOpen ? 'mdi-close' : 'mdi-menu' }}</v-icon>
          </v-btn>
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

  <!-- Tost menüsü -->
  <v-navigation-drawer class="tost-menu" v-model="mobileMenuOpen" app temporary>
    <v-list>
      <v-list-item v-for="(item, index) in menuItems" :key="index" :to="item.route">
        <v-list-item-icon>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>{{ item.title }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
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

@media (max-width: 1024px) {
  .menu-items {
    margin-left: 50px;
  }
}

@media (max-width: 768px) {
  .menu-items {
    margin-left: 20px;
  }
}

@media (max-width: 768px) {
  .v-navigation-drawer {
    width: 200px; /* Küçük ekranlarda tost menüsünün genişliği */
  }
}
</style>