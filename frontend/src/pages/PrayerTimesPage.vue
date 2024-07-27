<script setup>
import { useRoute } from 'vue-router';
import {onMounted, onUnmounted, ref} from 'vue';
import router from "@/router";
import Swal from "sweetalert2";
import {fetchPrayerTimesByDay} from "@/components/prayerTimesApi";
import {calculateCountdown} from "@/components/countdownUtils";
import {updateTime} from "@/components/getTime";

const route = useRoute();
const cityName = route.query.city;

const timings = ref({});
const hijriDate = ref({});
const gregorianDate = ref({});

const nowPrayer = ref('') // içinde bulundupumuz vakit
const countdown = ref(''); // sayacımız
let NEXTPRAYER = ref(''); // sonraki vakit

const currentTime = ref('');
const currentSecondAngle = ref(0);
const currentMinuteAngle = ref(0);
const currentHourAngle = ref(0);

const time = () => {
  const timeData = updateTime();
  currentTime.value = timeData.time;
  currentSecondAngle.value = timeData.secondAngle;
  currentMinuteAngle.value = timeData.minuteAngle;
  currentHourAngle.value = timeData.hourAngle;
};

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
  const interval2 = setInterval(()=>{
    time();
  },1000)
  onUnmounted(() => {
    clearInterval(interval);
    clearInterval(interval2);
  });
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
      <v-card class="transparent-card">
        <v-row class="title-container">
          <v-col class="left-column" cols="12">
            <span class="title" style="margin-left: 15px">Namaz Vakitleri: {{ cityName }}</span><hr><br>
            <span class="subtitle"><span style="color: blue">Miladi Takvim: </span>{{ gregorianDate.date }}</span><br>
            <span class="subtitle"><span style="color: blue">Hicri Takvim:  </span>{{ hijriDate.date }}</span><hr><br>
            <v-btn class="prayerTimes-btn" color="primary" @click="goToPrayerTimes">1 AYLIK NAMAZ VAKİTLERİ</v-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="6" class="countdown-container">
            <v-card class="countdown-card">
              <v-card-text class="countdown">
                <strong style="margin-bottom: 10px; color: white;  font-size: 18px;">{{ NEXTPRAYER }} Vaktine Kalan Süre</strong>
                <div>{{ countdown}}</div>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="12" sm="6" class="time-container">
            <v-card class="time-card">
              <v-card-text class="time">
                <strong>Bilgisayar Saati</strong>
                <div>{{ currentTime}}</div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
        <div class="table-container">
          <v-table class="transparent-table text">
            <thead>
            <tr>
              <th>İmsak</th>
              <th>Güneş</th>
              <th>Öğle</th>
              <th>İkindi</th>
              <th>Akşam</th>
              <th>Yatsı</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <td>{{timings.imsak}}</td>
              <td>{{timings.sunrise}}</td>
              <td>{{timings.dhuhr}}</td>
              <td>{{timings.asr}}</td>
              <td>{{timings.maghrib}}</td>
              <td>{{timings.isha}}</td>
            </tr>
            </tbody>
          </v-table>
        </div>
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
  background-color: #00796b;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  align-items: center;
  overflow-y: auto;
}

.transparent-card {
  background-color: rgba(255, 255, 255, 0.7) !important;
  box-shadow: none !important;
  border: none !important;
}

.title-container {
  text-align: center;
}

.title {
  font-size: 24px;
  font-weight: bold;
  color: black;
}

.subtitle {
  font-size: 20px;
  color: red;
  font-weight: bold;
}

.countdown-container, .time-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.countdown-card, .time-card {
  background-color: rgba(0, 100, 107, 1);
  padding: 20px;
  border-radius: 30px;
  text-align: center;
  color: white;
}

.countdown, .time {
  font-size: 2em;
  font-weight: bold;
}

.table-container {
  overflow-x: auto;
  margin-top: 10px;
  padding: 20px;
}

th, td {
  font-size: 18px;
  padding: 10px;
  color: #ffffff;
  background-color: rgba(0, 121, 107, 0.8);
}

.transparent-table {
  width: 100%;
  border-collapse: collapse;
}

@media (max-width: 1024px) {
  .background {
    padding: 10px;
    height: 100%;
    width: 100%;
    box-shadow: none;
    overflow-y: auto;
  }

  .title {
    font-size: 20px;
  }

  .subtitle {
    font-size: 18px;
    margin-bottom:2px;
  }

  .countdown, .time {
    font-size: 1em;
  }

  .countdown-card, .time-card {
    padding: 0px;
    border-radius: 10px;
  }

  th, td {
    font-size: 13px;
  }
}
</style>
