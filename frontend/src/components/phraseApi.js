import axios from 'axios';

export const fetchPhraseByID = async (id) => {
    try {
        const response = await axios.get(`https://prayer-times-gc0c.onrender.com/phrases/${id}`);
        return response.data;
    } catch (error) {
        console.error('Data alınırken bir hata oluştu:', error);
        throw error;
    }
};