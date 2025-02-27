<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import router from "@/router";

const searchQuery = ref("");
const cities = ref([]);

// backenddeki cities e ulaşıyoruz axios yardımıyla
const fetchCities = async () => {
  try {
    const response = await axios.get('http://localhost:3000/cities'); // API URL'nizi buraya ekleyin
    cities.value = response.data.cities; // API'den gelen şehirler verisini alıyoruz
  } catch (error) {
    console.error('Şehirler alınırken bir hata oluştu:', error);
  }
};

// Bileşen yüklendiğinde şehirleri almak için fetchCities fonksiyonunu çağırıyoruz
onMounted(fetchCities);

const capitalizeFirstLetter = (str) => {
  return str.charAt(0).toUpperCase() + str.slice(1).toLowerCase();
};

// Türkçe karakterleri İngilizce karşılıklarına çeviren fonksiyon
const convertTurkishChars = (str) => {
  const map = {
    'ı': 'i', 'İ': 'I', 'ç': 'c', 'Ç': 'C', 'ş': 's', 'Ş': 'S', 'ğ': 'g', 'Ğ': 'G', 'ü': 'u', 'Ü': 'U', 'ö': 'o', 'Ö': 'O'
  };
  return str.split('').map(char => map[char] || char).join('');
};

const filteredCities = computed(() => {
  const query = convertTurkishChars(searchQuery.value.toLowerCase());
  return cities.value.filter(city =>
      convertTurkishChars(city.city.toLowerCase()).includes(query)
  );
});

const goToPrayerTimes = (cityName) => {
  if (cityName) {
    router.push({ name: 'prayer-times', query: { city: cityName } });
  } else {
    console.error('City name is undefined');
  }
};
</script>

<template>
  <div class="background">
    <v-container>
      <v-text-field
          v-model="searchQuery"
          label="Şehir Ara"
          solo
          clearable
          class="text-field"
      ></v-text-field>

      <v-list
          class="city-list"
          dense
      >
        <v-list-item
            v-for="city in filteredCities"
            :key="city.id"
            class="city-list-item"
            @click="goToPrayerTimes(city.city)"
        >
          <v-list-item-content>
            <v-list-item-title>{{ capitalizeFirstLetter(city.city) }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-container>  </div>
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

.text-field {
  color: white;
  max-width: 1180px;
}

.city-list {
  max-height: 300px;
  max-width: 1180px;
  overflow-y: auto;
  border: 1px solid #ccc;
  border-radius: 4px;
}

@media (max-width: 768px) {
  .text-field, .city-list {
    max-width: 90%;
  }
}

@media (max-width: 480px) {
  .text-field, .city-list {
    max-width: 100%;
  }

  .city-list {
    max-height: 200px;
  }
}

</style>