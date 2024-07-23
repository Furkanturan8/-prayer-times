<script setup>
import { useRoute } from 'vue-router';
import axios from 'axios';
import { onMounted, ref } from 'vue';

const route = useRoute();
const cityName = route.query.city;

const timings = ref({});
const hijriDate = ref({});
const gregorianDate = ref({});

const now = new Date();
let today = now.getDate()

console.log('City Name:', cityName);

// API'den veri çekme
const fetchPrayerTimes = async () => {
  try {
    const response = await axios.get(`http://localhost:3000/prayer-times/${cityName}/${today}`);
    const { timings: t,hijri_date: hijri_date,gregorian_date: gregorian_date } = response.data;

    // Verileri ref'lere ayırma
    timings.value = t;
    hijriDate.value = hijri_date;
    gregorianDate.value = gregorian_date;

  } catch (error) {
    console.error('Vakitler alınırken bir hata oluştu:', error);
  }
};

onMounted(fetchPrayerTimes);


console.log('Timings:', JSON.stringify(timings.value, null, 2));
console.log('Hijri Date:', JSON.stringify(hijriDate.value, null, 2));
console.log('Gregorian Date:', JSON.stringify(gregorianDate.value, null, 2));

</script>

<template>
  <div class="background">
    <v-container>
      <v-card class="transparent-card" >
        <v-card-title class="title">Namaz Vakitleri</v-card-title>
        <v-card-subtitle class="subtitle">Şehir : {{ cityName }}</v-card-subtitle>

        <v-card class="transparent-card">
          <v-card-text class="text">
            <strong>Miladi Takvim:</strong> {{ gregorianDate.date }}
          </v-card-text>
          <v-card-text class="text">
            <strong>Hicri Takvim:</strong> {{ hijriDate.date }}
          </v-card-text>
        </v-card>

        <v-table class="transparent-table text">
          <thead>
          <tr>
            <th>Namaz</th>
            <th>Vakit</th>
          </tr>
          </thead>
          <tbody>
          <tr>
            <td>İmsak</td>
            <td>{{ timings.imsak }}</td>
          </tr>
          <tr>
            <td>Sabah</td>
            <td>{{ timings.sunrise }}</td>
          </tr>
          <tr>
            <td>Öğle</td>
            <td>{{ timings.dhuhr }}</td>
          </tr>
          <tr>
            <td>İkindi</td>
            <td>{{ timings.asr }}</td>
          </tr>
          <tr>
            <td>Akşam</td>
            <td>{{ timings.maghrib }}</td>
          </tr>
          <tr>
            <td>Yatsı</td>
            <td>{{ timings.isha }}</td>
          </tr>
          </tbody>
        </v-table>
      </v-card>
    </v-container>
  </div>
</template>

<style scoped>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.background {
  background-image: url('@/assets/cami2.jpg');
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-position: center;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  align-items: center;
  box-shadow: inset 0 0 450px rgba(1, 1, 1, 1); /* Gölge efekti */
}

.transparent-card {
  background-color: rgba(255, 255, 255, 0.7) !important;
  box-shadow: none !important;
  border: none !important;
}

.transparent-table {
  background-color: rgba(255, 255, 255, 0.7) !important;
  box-shadow: none !important;
  border: none !important;
}

.title {
  font-size: 24px;
  font-weight: bold;
}

.subtitle {
  font-size: 20px;
  margin-bottom: 20px;
}

.text {
  font-size: 18px;
}

th, td {
  font-size: 18px;
  padding: 10px;
}
</style>