<script setup>
import { ref } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import axios from "axios";
import Swal from "sweetalert2";

const name = ref('');
const surname = ref('');
const email = ref('');
const message = ref('');

const handleSubmit = () => {
  const formData = {
    name: name.value,
    surname:surname.value,
    email: email.value,
    message: message.value,
  };

  createMessage(formData);

  name.value = '';
  surname.value = '';
  email.value = '';
  message.value = '';
};

const createMessage = async (data) => {
  try {
    await axios.post(`http://localhost:3000/contact`,data, {
      headers: {
        'Content-Type': 'application/json',
      },
    });
    await Swal.fire({
      title: 'Mesaj Gönderimi Başarılı',
      text: "Mesajınız Admin Paneline Gönderildi!",
      icon: 'success',
      confirmButtonText: 'Tamam',
    });
  }catch (error){
    await Swal.fire({
      title: 'Hata',
      text: "Mesaj Gönderimi Başarısız!",
      icon: 'error',
      confirmButtonText: 'Tamam',
    });
    console.error("Mesaj gönderilirken hata oluştu!",error)
  }
}

</script>

<template>
  <div class="main">
    <div class="contact">
      <h1>Bize Ulaşın!</h1>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="name">Adınız:</label>
          <input type="text" id="name" v-model="name" required />
        </div>
        <div class="form-group">
          <label for="surname">Soyadınız:</label>
          <input type="text" id="surname" v-model="surname" required />
        </div>
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" id="email" v-model="email" required />
        </div>
        <div class="form-group">
          <label for="message">Mesajınız:</label>
          <textarea id="message" v-model="message" required></textarea>
        </div>
        <button type="submit">Gönder</button>
      </form>
      <div class="social-links">
        <a href="https://github.com/furkanturan8" target="_blank" aria-label="Github">
          <font-awesome-icon :icon="['fab', 'github']" />
          <span>Github</span>
        </a>
        <a href="https://twitter.com/" target="_blank" aria-label="Twitter">
          <font-awesome-icon :icon="['fab', 'twitter']" />
          <span>Twitter</span>
        </a>
        <a href="https://instagram.com/trnfurkan7" target="_blank" aria-label="Instagram">
          <font-awesome-icon :icon="['fab', 'instagram']" />
          <span>Instagram</span>
        </a>
        <a href="https://wa.me/" target="_blank" aria-label="WhatsApp">
          <font-awesome-icon :icon="['fab', 'whatsapp']" />
          <span>Whatsapp</span>
        </a>
      </div>
    </div>

  </div>
</template>

<style scoped>
.main {
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

.contact {
  max-width: 600px;
  flex-direction: column;
  width: 100%;
  padding: 70px;
  background-color: rgba(0,0,0,0.9);
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  margin-bottom: 8em;
}

.social-links {
  margin-top: 40px;
  text-align: center;
}

h1 {
  color: white;
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  color: white;
  display: block;
  margin-bottom: 5px;
}

input,
textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  color: white;
}

button {
  display: block;
  width: 100%;
  padding: 10px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

button:hover {
  background-color: #45a049;
}

.social-links {
  display: flex;
  justify-content: center;
  gap: 20px;
  color: white;
}

.social-links a {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: inherit;
  font-size: 1.2em;
}

.social-links a:hover {
  color: #007bff;
}

.social-links font-awesome-icon {
  margin-right: 8px;
}
.social-links span {
  margin-left: 8px;
}

@media (max-width: 768px) {
  .main {
    padding: 10px;
  }

  .contact {
    padding: 20px;
  }

  h1 {
    font-size: 1.5em;
  }

  .form-group {
    margin-bottom: 10px;
  }

  button {
    font-size: 14px;
  }

  .social-links a {
    font-size: 1em;
  }
}
</style>
