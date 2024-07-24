import axios from 'axios';

export const fetchPrayerTimes = async (cityName) => {
    const now = new Date();
    let today = now.getDate()
    try {
        const response = await axios.get(`http://localhost:3000/prayer-times/${cityName}/${today}`);
        return response.data;
    } catch (error) {
        console.error('Vakitler alınırken bir hata oluştu:', error);
        throw error;
    }
};
