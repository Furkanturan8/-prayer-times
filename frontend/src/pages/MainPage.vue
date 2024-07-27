<script setup>

import { ref, onMounted } from 'vue';
import {fetchPhraseByID} from "@/components/phraseApi";
import {updateTime} from "@/components/getTime";

const phraseData = ref('');
const title = ref('');
const phrase= ref('')

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

try {
  const data = await fetchPhraseByID(1);
  phraseData.value = data
  title.value = phraseData.value.phrases.title
  phrase.value = phraseData.value.phrases.phrase

} catch (error){
  console.error('Özlü sözler alınırken hata oluştu!')
}

onMounted(() => {
  time();
  setInterval(time, 1000);
});
</script>


<template>
  <div class="background">

     <div class="phrases">
       <div>{{title}}</div>
       <div>{{phrase}}</div>
     </div>

    <div class="container">
      <div class="analog-clock">
        <svg class="clock" viewBox="0 0 100 100">
          <circle cx="50" cy="50" r="48" />
          <line :x1="50" :y1="50" :x2="50 + 45 * Math.sin(currentSecondAngle)" :y2="50 - 45 * Math.cos(currentSecondAngle)" class="second-hand" />
          <line :x1="50" :y1="50" :x2="50 + 40 * Math.sin(currentMinuteAngle)" :y2="50 - 40 * Math.cos(currentMinuteAngle)" class="minute-hand" />
          <line :x1="50" :y1="50" :x2="50 + 35 * Math.sin(currentHourAngle)" :y2="50 - 35 * Math.cos(currentHourAngle)" class="hour-hand" />
        </svg>
      </div>
      <div class="digital-clock">{{ currentTime }}</div>
    </div>
  </div>
</template>

<style>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.background {
  background-image: url('@/assets/cami2.jpg');;
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-position: center;
  height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  box-shadow: inset 0 0 450px rgba(1, 1, 1, 1); /* İçten gölge efekti */
}

.container {
  text-align: center;
  color: white;
}

.phrases {
  background-color: rgba(0, 0, 0, 0.5);
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 10px;
  text-align: center;
  font-size: 30px;
  font-weight: bold;
  color : white;
}

.digital-clock {
  font-size: 3em;
  margin-bottom: 20px;
}

.analog-clock .clock {
  width: 200px;
  height: 200px;
}

.analog-clock .clock circle {
  fill: none;
  stroke: white;
  stroke-width: 4;
}

.analog-clock .clock .second-hand {
  stroke: red;
  stroke-width: 2;
}

.analog-clock .clock .minute-hand {
  stroke: white;
  stroke-width: 4;
}

.analog-clock .clock .hour-hand {
  stroke: white;
  stroke-width: 6;
}
@media (max-width: 768px) {
  .background {
    padding: 10px;
  }

  .phrases {
    font-size: 20px;
  }

  .digital-clock {
    font-size: 2em;
  }

  .analog-clock .clock {
    width: 150px;
    height: 150px;
  }
}
</style>
