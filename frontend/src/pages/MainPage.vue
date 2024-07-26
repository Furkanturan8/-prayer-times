<script setup>

import { ref, onMounted } from 'vue';
import {fetchPhraseByID} from "@/components/phraseApi";

const time = ref('');
const secondAngle = ref(0);
const minuteAngle = ref(0);
const hourAngle = ref(0);
const phraseData = ref('');
const title = ref('');
const phrase= ref('')

try {
  const data = await fetchPhraseByID(1);
  phraseData.value = data
  title.value = phraseData.value.phrases.title
  phrase.value = phraseData.value.phrases.phrase

} catch (error){
  console.error('Özlü sözler alınırken hata oluştu!')
}

const updateTime = () => {
  const now = new Date();
  time.value = now.toLocaleTimeString();
  secondAngle.value = (now.getSeconds() / 60) * 2 * Math.PI;
  minuteAngle.value = ((now.getMinutes() + now.getSeconds() / 60) / 60) * 2 * Math.PI;
  hourAngle.value = ((now.getHours() % 12 + now.getMinutes() / 60) / 12) * 2 * Math.PI;
};

onMounted(() => {
  updateTime();
  setInterval(updateTime, 1000);
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
          <line :x1="50" :y1="50" :x2="50 + 45 * Math.sin(secondAngle)" :y2="50 - 45 * Math.cos(secondAngle)" class="second-hand" />
          <line :x1="50" :y1="50" :x2="50 + 40 * Math.sin(minuteAngle)" :y2="50 - 40 * Math.cos(minuteAngle)" class="minute-hand" />
          <line :x1="50" :y1="50" :x2="50 + 35 * Math.sin(hourAngle)" :y2="50 - 35 * Math.cos(hourAngle)" class="hour-hand" />
        </svg>
      </div>
      <div class="digital-clock">{{ time }}</div>
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
</style>
