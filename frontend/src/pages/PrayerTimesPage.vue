<script setup>
import { useRoute } from 'vue-router';
import {onMounted, onUnmounted, ref} from 'vue';
import router from "@/router";
import Swal from "sweetalert2";
import {fetchPrayerTimesByDay} from "@/components/prayerTimesApi";
import {calculateCountdown} from "@/components/countdownUtils";

const route = useRoute();
const cityName = route.query.city;

const timings = ref({});
const hijriDate = ref({});
const gregorianDate = ref({});

const nowPrayer = ref('') // içinde bulundupumuz vakit
const countdown = ref(''); // sayacımız
let NEXTPRAYER = ref(''); // sonraki vakit

const init = async () => {
  if (!cityName) {
    Swal.fire({
      title: 'Uyarı',
      text: "Lütfen Önce Şehir Seçiniz!",
      icon: 'warning',
      confirmButtonText: 'Tamam',
    });
  await router.push({ name: 'cities' });
  return;
}

try {
  const data = await fetchPrayerTimesByDay(cityName);
  timings.value = data.timings;
  hijriDate.value = data.hijri_date;
  gregorianDate.value = data.gregorian_date;

  const countdownData = calculateCountdown(timings.value);
  countdown.value = countdownData.countdown;
  nowPrayer.value = countdownData.nowPrayer;
  NEXTPRAYER.value = countdownData.nextPrayer;
} catch (error) {
  console.error('Vakitler alınırken bir hata oluştu:', error);
}
};

onMounted(() => {
  init();
  const interval = setInterval(() => {
    init();
  }, 30000);

  onUnmounted(() => clearInterval(interval));
});

const goToPrayerTimes = async () => {
  const url = router.resolve({
    name: 'prayer-times-month',
    params: {
      month: gregorianDate.value.month_name
    },
    query: {
      city: cityName
    }
  }).href;

  window.open(url, '_blank');
}
</script>

<template>
  <div class="background">
    <v-container>
      <v-card class="transparent-card" >
        <v-row class="title-container">
          <v-col>
            <span class="title" style="margin-left: 15px">Namaz Vakitleri</span>
          </v-col>
          <v-col class="right-align">
            <v-btn class="prayerTimes-btn" color="primary" @click="goToPrayerTimes">1 AYLIK NAMAZ VAKİTLERİ</v-btn>
          </v-col>
        </v-row>
        <v-card-subtitle class="subtitle">Şehir : {{ cityName }}</v-card-subtitle>
        <v-card class="transparent-card">
          <v-row>
            <v-col class="left-column">
              <v-card-text class="text">
                <strong>Miladi Takvim:</strong> {{ gregorianDate.date }}
              </v-card-text>
              <v-card-text class="text">
                <strong>Hicri Takvim:</strong> {{ hijriDate.date }}
              </v-card-text>
            </v-col>
            <v-col class="right-column" cols="auto">
              <v-card class="countdown-card" >
                <v-card-text class="countdown">
                  <strong style="margin-bottom: 10px; color: white;  font-size: 18px;">{{ NEXTPRAYER }} Vaktine Kalan Süre</strong>
                  <div>{{ countdown}}</div>
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
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
  margin-bottom: 20px;
}

.left-column {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

.right-column {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  margin-right: 5px;
  margin-top: 5px;
}

.countdown-card {
  background-color: grey;
  padding: 20px;
  border-radius: 10px;
  text-align: center;
  margin: auto;
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

.countdown {
  font-size: 2em;
  font-weight: bold;
  color: lightskyblue;
}

th, td {
  font-size: 18px;
  padding: 10px;
}

.prayerTimes-btn {
  margin-top: 5px;
  margin-right: 5px;
  margin-left: auto;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 315px;
}
</style>