<script setup>
import {computed, onMounted, ref, watch} from "vue";
import {fetchPrayerTimesByMonth} from "@/components/prayerTimesApi";
import {useRoute} from "vue-router";

const route = useRoute();
const cityName = route.query.city;
const loading = ref(true);
const DATA = ref([]);
const itemsPerPage = 10;
const currentPage = ref(1);
const totalPages = computed(() => Math.ceil(DATA.value.length / itemsPerPage));

onMounted(async () => {
  loading.value = false;
  try {
    const data = await fetchPrayerTimesByMonth(cityName);
    DATA.value = data;
  } catch (error) {
    console.error('Vakitler alınırken bir hata oluştu:', error);
  } finally {
    loading.value = false;
  }
});

const paginatedData = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return DATA.value.slice(start, end);
});

watch(currentPage, (newPage) => {
  console.log("Current page changed to:", newPage);
});
</script>

<template>
  <div class="background">
  <v-container>
    <v-card class="card-height">
      <v-card-title class="title">
        Aylık Namaz Vakitleri - {{ cityName }}
      </v-card-title>
      <v-progress-circular v-if="loading" indeterminate color="primary" class="ma-4">
        Veriler Yükleniyor...
      </v-progress-circular>

      <div class="table-container" v-else>
      <table class="table table-striped table-dark">
        <thead>
        <tr>
          <th scope="col">Miladi Takvim</th>
          <th scope="col">Hicri Takvim</th>
          <th scope="col">İmsak</th>
          <th scope="col">Güneş</th>
          <th scope="col">Öğle</th>
          <th scope="col">İkindi</th>
          <th scope="col">Akşam</th>
          <th scope="col">Yatsı</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(item,index) in paginatedData" :key="index">
          <td>{{ item.gregorian_date.date }}</td>
          <td>{{ item.hijri_date.date }}</td>
          <td>{{ item.timings.imsak }}</td>
          <td>{{ item.timings.sunrise }}</td>
          <td>{{ item.timings.dhuhr }}</td>
          <td>{{ item.timings.asr }}</td>
          <td>{{ item.timings.maghrib }}</td>
          <td>{{ item.timings.isha }}</td>
        </tr>
        </tbody>
      </table>
      </div>
      <div class="pagination">
        <button @click="currentPage = currentPage - 1" :disabled="currentPage === 1">
          &lt;
        </button>
        <button
            v-for="page in totalPages"
            :key="page"
            @click="currentPage = page"
            :class="{ active: page === currentPage }"
        >
          {{ page }}
        </button>
        <button @click="currentPage = currentPage + 1" :disabled="currentPage === totalPages">
          &gt;
        </button>
      </div>
        <v-progress-circular
          v-if="loading"
          indeterminate
          color="primary"
          class="ma-4"
      ></v-progress-circular>
    </v-card>
  </v-container>
  </div>
</template>

<style scoped>
html, body {
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
  height: 90vh;
  width: 100vw;
  justify-content: center;
  align-items: center;
  box-shadow: inset 0 0 450px rgba(1, 1, 1, 1); /* Gölge efekti */
  overflow-y: auto;
}

.title{
  background-color: #343a40;
  color: white;
}
table {
  width: 100%;
  border-collapse: collapse;
  background-color: #343a40;
  color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.table-container {
  overflow-x: auto;
}
th, td {
  padding: 15px;
  text-align: left;
  word-wrap: break-word; /* Metin taşmaları için */
}

thead {
  background-color: #454d55;
}

th {
  text-transform: uppercase;
}

tbody tr:nth-child(even) {
  background-color: #3e444a;
}

tbody tr:nth-child(odd) {
  background-color: #343a40;
}

tbody tr:hover {
  background-color: #495057;
}

.card-height {
  height: auto;
  max-height: 100vh;
  overflow-y: auto;
  margin-bottom: 50px;
  margin-top: 120px;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 10px;
  margin-bottom: 10px;
}

.pagination button {
  margin: 0 5px;
  padding: 10px 15px;
  background-color: #343a40;
  color: white;
  border: none;
  cursor: pointer;
}

.pagination button.active {
  background-color: #007bff;
}

.pagination button:hover {
  background-color: #495057;
}
.pagination button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}
@media (max-width: 768px) {
  .title {
    font-size: 20px;
    padding: 8px;
  }

  th, td {
    padding: 8px;
    font-size: 14px;
  }

  .pagination button {
    padding: 8px 12px;
  }
  table {
    font-size: 12px;
  }
}
</style>