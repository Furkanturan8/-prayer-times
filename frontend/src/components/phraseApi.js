import axios from 'axios';

export const fetchPhraseByID = async (id) => {
    try {
        const response = await axios.get(`http://localhost:3000/phrases/${id}`);
        return response.data;
    } catch (error) {
        console.error('Data alınırken bir hata oluştu:', error);
        throw error;
    }
};