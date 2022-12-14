import axios from "axios";

const axiosInstance = axios.create({ baseURL: 'http://localhost:4000/' });

export const testRequest = () => {
    const url = '/test';
    return axiosInstance.get(url)
        .then(response => { return response.data; });
};