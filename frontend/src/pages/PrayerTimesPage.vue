<script setup>
import { useRoute } from 'vue-router';

const route = useRoute();
const cityName = route.query.city;

console.log('City Name:', cityName); // Debugging: check the city name


</script>

<template>
  <v-container>
    <v-card>
      <v-card-title>Şehir Detayları</v-card-title>
      <v-card-subtitle>ID: {{ cityName }}</v-card-subtitle>
    </v-card>
  </v-container>
</template>

<style scoped>

</style>