
import {axiosInstanceWithJWT} from '../api/axios';
import React, { useState, useEffect } from 'react';


function LicenseKeys() {

    const [cheats, setCheats] = useState([]); // Данные для cheats
    const [users, setUsers] = useState([]); // Данные для user logins and roles
    const [loading, setLoading] = useState(true);
    const [subsCount, setSubsCount] = useState(1);
    const [daysCount, setDaysCount] = useState(25);
    const [selectedCheat, setSelectedCheat] = useState('');
    const [selectedUser, setSelectedUser] = useState('');
    const [keys, setKeys] = useState([]);
    const [isLoading, setIsLoading] = useState(false);

    const handleCheatChange = (event) => {
        setSelectedCheat(event.target.value);
    };

    const handleUserChange = (event) => {
        setSelectedUser(event.target.value);
    };


    const handleSubsChange = (event) => {
        setSubsCount(event.target.value);
    };

    const handleDaysChange = (event) => {
        setDaysCount(event.target.value);
    };

    const handleSubmit = async () => {

        setIsLoading(true);
        const dataToSend = {
            count_keys: parseInt(subsCount),
            ttl_cheat: parseInt(daysCount),
            cheat:selectedCheat,
            holder:selectedUser
        };


        try {
            const [licenseResponse] = await Promise.all([
                axiosInstanceWithJWT.post('/api/license-keys/',dataToSend),
            ]);
            setKeys(licenseResponse.data.keys)
        } catch (error) {
            console.error("Ошибка при отправке данных: ", error);
        } finally {
            setIsLoading(false);
        }

    };

    useEffect(() => {
        const fetchDatas = async () => {
            try {
                const [cheatsResponse, usersResponse] = await Promise.all([
                    axiosInstanceWithJWT.get('/api/cheats/'),
                    axiosInstanceWithJWT.get('/api/users/sw')
                ]);

                //const { cheats } = response.data;
                setCheats(cheatsResponse.data.cheats);
                setUsers(usersResponse.data.users); // Убедитесь, что здесь правильно обрабатывается ответ
            } catch (error) {
                console.error("Ошибка при получении данных: ", error);
            } finally {
                setLoading(false);
            }
        };

        // Вызов функции
        fetchDatas();
    }, []);



    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h2>Cheat</h2>
            <select value={selectedCheat} onChange={handleCheatChange}>
                <option value="">select cheat</option>
                {cheats.map((cheat) => (
                    <option key={cheat.id} value={cheat.name}>
                        {cheat.name}
                    </option>
                ))}
            </select>

            <h2>User Login</h2>
            <select value={selectedUser} onChange={handleUserChange}>
                <option value="">select user</option>
                {users.map((user) => (
                    <option key={user.id} value={user.login}>
                        {user.login} - {user.role} {/* Пример, предполагая, что у вас есть 'role' */}
                    </option>
                ))}
            </select>
            <br></br>
            <h2>count subscribtion</h2>
            <div>
                <label htmlFor="subsCount">subs: {subsCount}</label>
                <input
                    id="subsCount"
                    type="range"
                    min="0"
                    max="50"
                    value={subsCount}
                    onChange={handleSubsChange}
                />
            </div>
            <h2>count days</h2>
            <div>
                <label htmlFor="daysCount">days: {daysCount}</label>
                <input
                    id="daysCount"
                    type="range"
                    min="0"
                    max="30"
                    value={daysCount}
                    onChange={handleDaysChange}
                />
            </div>
            <br></br>
            <button onClick={handleSubmit}>gen subscribtion</button>

            <div>
                <h2>Keys list:</h2>
                {isLoading ? <p>Loading...</p> : (
                    keys.map((key, index) => (
                        <div key={index}>
                            <p>{key.license_key}</p>
                        </div>
                    ))
                )}

            </div>
        </div>
    );
}

export default LicenseKeys;