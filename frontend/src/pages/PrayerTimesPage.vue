<script setup>
import { useRoute } from 'vue-router';
import axios from 'axios';
import {onMounted, onUnmounted, ref} from 'vue';
import router from "@/router";
import Swal from "sweetalert2";

const route = useRoute();
const cityName = route.query.city;

const timings = ref({});
const hijriDate = ref({});
const gregorianDate = ref({});

const now = new Date();
let today = now.getDate()

const countdown = ref('');
let NEXTPRAYER = ref('');

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

const getCurrentTime = () => {
  const now = new Date();
  return {
    hours: now.getHours(),
    minutes: now.getMinutes(),
  };
};

const timeToMinutes = (timeString) => {
  if (!timeString) return 0; // Geçersiz zaman dizesi durumunda 0 dakika döndür

  // Zaman dizesinden saat dilimini ayırma
  const [timePart] = timeString.split(' (');

  // Saat ve dakika bilgilerini ayırma
  const [hours, minutes] = timePart.split(':').map(Number);

  return hours * 60 + minutes;
};

const calculateCountdown = () => {
  const current = getCurrentTime();
  const currentMinutes = current.hours * 60 + current.minutes;

  if (!timings.value || Object.keys(timings.value).length === 0) {
    countdown.value = 'Veri Yüklenmedi';
    console.log('Namaz vakitleri verisi boş');
    return;
  }

  const times = [
    { name: 'İmsak', time: timings.value.imsak },
    { name: 'Sabah', time: timings.value.sunrise },
    { name: 'Öğle', time: timings.value.dhuhr },
    { name: 'İkindi', time: timings.value.asr },
    { name: 'Akşam', time: timings.value.maghrib },
    { name: 'Yatsı', time: timings.value.isha },
  ];

  let nextPrayerFound = false;
  for (let i = 0; i < times.length; i++) {
    const currentPrayer = times[i];
    const nextPrayer = times[i + 1] || null;

    const currentPrayerMinutes = timeToMinutes(currentPrayer.time);
    const nextPrayerMinutes = nextPrayer ? timeToMinutes(nextPrayer.time) : Infinity;

    console.log(`Namaz: ${currentPrayer.name}, Dakika: ${currentPrayerMinutes}`);
    console.log('Sonraki namazın dakikaları:', nextPrayerMinutes);

    if (currentMinutes >= currentPrayerMinutes && currentMinutes < nextPrayerMinutes) {
      nextPrayerFound = true;
      const nextPrayerTime = nextPrayer ? timeToMinutes(nextPrayer.time) : null;
      if (nextPrayerTime) {
        const minutesUntilNextPrayer = nextPrayerTime - currentMinutes;
        const hoursUntilNextPrayer = Math.floor(minutesUntilNextPrayer / 60);
        const minutesRemaining = minutesUntilNextPrayer % 60;
        countdown.value = `${hoursUntilNextPrayer} saat ${minutesRemaining} dakika`;
        console.log('Kalan süre:', countdown.value);
      } else {
        countdown.value = 'Gün Sonu';
      }
      NEXTPRAYER.value = nextPrayer.name
      break;
    }
  }

  // Eğer mevcut vakitler arasında geçerli bir vakit bulunamadıysa sonraki günün imsağına göre ayarla sayacı
  if (!nextPrayerFound) {
    const nextDayTimes = [
      { name: 'İmsak', time: timings.value.imsak },
    ];

    if (nextDayTimes.length > 0) {
      const nextDayPrayerTime = timeToMinutes(nextDayTimes[0].time);
      const minutesUntilNextDayPrayer = (1440 - currentMinutes) + nextDayPrayerTime; // 1440 dakika = 24 saat
      const hoursUntilNextDayPrayer = Math.floor(minutesUntilNextDayPrayer / 60);
      const minutesRemaining = minutesUntilNextDayPrayer % 60;
      countdown.value = `${hoursUntilNextDayPrayer} saat ${minutesRemaining} dakika`;
    } else {
      countdown.value = 'Gün Sonu';
    }
  }
};

// Fonksiyonları başlat
const init = async () => {
  if (!cityName) {
    Swal.fire({
      title: 'Uyarı',
      text: "Lütfen Önce Şehir Seçiniz!",
      icon: 'warning',
      confirmButtonText: 'Tamam',
    });
    await router.push({name: 'cities'}); // /cities sayfasına yönlendirin
    return;
  }
  await fetchPrayerTimes();
  calculateCountdown();
};

// Sayfa yüklendiğinde ve her 30 saniyede bir yenileyin
onMounted(() => {
  init();
  const interval = setInterval(() => {
    fetchPrayerTimes();
    calculateCountdown();
  }, 30000); // 30 saniyede bir yenile

  // Bileşen kaldırıldığında interval'i temizle
  onUnmounted(() => clearInterval(interval));
});

</script>

<template>
  <div class="background">
    <v-container>
      <v-card class="transparent-card" >
        <v-card-title class="title">Namaz Vakitleri</v-card-title>
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
                  <strong style="margin-bottom: 10px; color: white;  font-size: 18px;">{{ NEXTPRAYER }} Ezanına Kalan Süre</strong>
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
</style>