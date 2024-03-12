
import {axiosInstanceWithJWT} from '../api/axios';
import React, { useState, useEffect } from 'react';


function LicenseKeys() {

    const [items, setItems] = useState([]); // Для хранения данных из API
    const [loading, setLoading] = useState(true); // Для отслеживания состояния загрузки

    useEffect(() => {
        // Объявление асинхронной функции внутри эффекта
        const fetchData = async () => {
            try {
                const response = await axiosInstanceWithJWT.get('/api/cheats/');
                const { cheats } = response.data;
                console.log(cheats)
                setItems(cheats);
                setLoading(false);
            } catch (error) {
                console.error("Ошибка при получении данных: ", error);
                setLoading(false);
            }
        };

        // Вызов функции
        fetchData();
    }, []);



    if (loading) {
        return <div>Загрузка...</div>;
    }

    return (
        <select>
            {items.map(item => (
                <option key={item.id} value={item.name}>
                    {item.name}
                </option>
            ))}
        </select>
    );
}

export default LicenseKeys;